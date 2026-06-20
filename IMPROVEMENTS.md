# Manta parser — correctness & performance improvement plan

Borrowed from the Clarity (Java) reference parser at `../clarity`. Executed on branch
`jcoene/goal-perf`, one goal per commit, in `/goal` mode.

This plan was produced by a 14-agent review (7 subsystem analyses, each adversarially
verified against both codebases with real CPU/alloc profiles of `replays/2159568145.dem`).
27 findings confirmed, 19 adjusted, 1 rejected.

## Ground rules

- **Strictly no public API change.** Internals only — no new exported symbols, no changed
  signatures. The de-box work stays behind the existing `Entity.Get`/`GetFloat32` (box lazily
  on read). The internal `fieldDecoder` type is unexported and is fair game.
- **One optimization per commit.** Each commit message carries its `benchstat` table vs the
  previous commit + `go test ./...` PASS. The branch history is the story.
- **Commits only — do NOT open a PR, do NOT push.** Make local commits on `jcoene/goal-perf` and
  report the cumulative results here for review. The user opens the PR.
- **Sequence:** safe high-ROI wins first, re-baseline, then the invasive typed-state rewrite.

## Verification protocol (per goal)

- **Correctness gate (non-negotiable):** `go test ./...` stays green with **identical** golden
  values (manta_test.go runs ~48 real replays asserting exact `expectEntityEvents` counts,
  `expectHeroEntityMana` floats, combat-log counts). A perf change that alters output is a regression.
- **Perf measure:** `go test -run=XXX -bench=<Bench> -benchmem -benchtime=10x -count=10` before
  and after; compare with `benchstat` (`go install golang.org/x/perf/cmd/benchstat@latest`).
  Report `sec/op`, `B/op`, **`allocs/op`** (the near-deterministic leading indicator).
- **Profile confirmation:** `make memprofile` / `make cpuprofile` — confirm the *targeted* site
  actually moved, not just the aggregate.
- **Canonical bench replay (held fixed for the whole branch):** **`8552595443`** (build 6601 — most
  representative of current real-world replays). The review's profile evidence below was gathered on
  `2159568145`; re-capture the alloc profile on `8552595443` at P0 to confirm the top sources before
  optimizing (relative signal is expected to hold).

### Profile evidence (gathered on `replays/2159568145.dem`, M4 Pro; canonical bench going forward is `8552595443`)
- ~11.0M allocs/op, ~315 MB/op.
- Streaming `*os.File`: 798 ms/op · `bufio`: 658 ms/op · in-memory `bytes.Reader`: 644 ms/op.
- Top alloc sources: `readFieldPaths` slice 56–60% · quantized-float box ~10–12% · QAngle
  `[]float32` ~4.5–5.5% · signed/unsigned int box ~2–3% · `*pendingMessage` ~5% · `*outerMessage` ~5%.

---

## Goal summary

| ID | Goal | Type | Risk | Effort | Expected impact |
|----|------|------|------|--------|-----------------|
| **P0** | In-memory benchmark rig + recent-replay bench | infra | none | small | makes parse wins measurable (removes ~80% syscall noise) |
| **P1.1** | `pendingMessage` value-slice + typed `sort.Stable` (reuse buffer) | perf | low | small | −0.54M allocs, −20 MB; +determinism |
| **P1.2** | `outerMessage` by value | perf | low | small | −0.56M allocs, −17 MB |
| **P1.3** | snappy scratch-buffer reuse | perf | med | small | −9 MB/op |
| **P1.4** | modifier emit: early-return when no handler | perf | low | small | removes modifier unmarshal from hot path |
| **P1.5** | reuse per-packet `tuples` + entity-data reader | perf | low | small | −0.44M allocs |
| **P1.6** | hoist per-entity `fpCache`/`fpNoop` maps to class | perf | low | small | −2 maps/entity; smaller Entity |
| **P1.7** | hoist `v(6)` debug guard out of per-field loop | perf | low | small | −33M branch/calls (CPU) |
| **P1.8** | buffer `stream` IO (bufio when not byte-reader) | perf | low | small | 798→658 ms streaming path |
| **P2.1** | **string-table additive index fix** (real bug) | correctness | low | small | fixes 4,469 wrong indices |
| **P2.2** | string-table: fail loud, don't swallow panics | correctness | low | small | surfaces silent corruption |
| **P2.3** | `getEventKey` off-by-one (`>` → `>=`) | correctness | low | small | prevents panic on skewed events |
| **P2.4** | modifier: skip empty/deleted entries | correctness | low | small | no spurious zero events |
| **P2.5** | combat-log `Type()/TypeName()/String()` descriptor-driven | correctness | low | small | latent landmine fix |
| **P2.6** | decoder forward-compat additions (never-occur types) | correctness | low | small | CUtlBinaryBlock/Quaternion/int64/etc |
| **P2.7** | HSequence / HeroID_t decode parity (value-changing) | correctness | med | small | matches clarity; suite-gated |
| **P2.8** | BloodType fixed-8 decode (risky, suite-gated) | correctness | med | small | only if all goldens stay green |
| **P2.9** | QAngle precise/32/0-bit forward-compat | correctness | low | small | future builds |
| **P2.10** | mana/runetime patch sentinel guards | correctness | low | small | robustness/parity |
| **P2.11** | outer-message size sanity bound | correctness | low | small | OOM hardening on corrupt input |
| **P2.12** | field-path depth-7 guard + comment | correctness | low | small | loud failure vs cryptic panic |
| **P1.9** | **reusable field-path buffer** (value-type, no pool/copy/slice) | perf | med | medium | **−56–60% of all allocs** |
| **P1.10** | word-at-a-time bit reader | perf | med | small | top CPU win (after P0) |
| **P1.11** | varint + `readByte`/`readLeUintX` from accumulator | perf | low | small | compounds P1.10 |
| **P1.12** | flatten huffman tree to int arrays (no dispatch) | perf | low | medium | removes per-bit interface calls |
| **P1.13** | 8-bit field-op lookup table | perf | med | medium | resolves majority of ops in 1 index |
| **P1.14** | decode class baseline once, clone per entity | perf | med | medium | baseline re-decode is ~4× update decode |
| **P3.1** | **typed entity state (eliminate boxing)** | perf | med | large | ~20% of allocs (float+qangle+int) |
| — | string-table Items: map → dense slice | perf | med | medium | locality; depends on P2.1 |
| — | integrate `CDemoStringTables` full dumps | correctness | med | medium | seek/robustness; depends on P2.1 |
| **PARK** | combat-log name-resolution helper / `CombatLogEntry` | correctness | — | — | **blocked: new exported API** |
| **PARK** | S2 HLTV `CMsgDOTACombatLogEntry` path | correctness | — | — | **blocked: new exported API** |
| **PARK** | VTProtobuf zero-reflection unmarshal | perf | med | large | envelope decode only; large effort |

Ordering note: P1.9–P1.14 are the bigger structural wins; they're listed after the trivial
P1.1–P1.8 alloc trims so we build a clean measured baseline first, exactly as agreed.

---

## P0 — In-memory benchmark rig (foundational, no parser logic)

**Why:** the current `BenchmarkMatch` (manta_test.go:609) streams from `*os.File` inside the timed
loop with no `b.ResetTimer()`; the review measured ~80% of CPU in per-byte-read syscalls (798 ms vs
644 ms in-memory). Parse optimizations are invisible under that noise, and `B/op` is polluted by IO.

**Change:** add `BenchmarkMatch8552595443` (build 6601) as the canonical bench, and rewrite `bench()` to
read the replay into `[]byte` once before the loop, `b.ResetTimer()`, and parse via in-memory
`NewParser(buf)` in the loop. Point the Makefile `memprofile`/`cpuprofile` targets at `8552595443` too.
No parser code touched → cannot affect correctness.

**Verify:** `go test ./...` unchanged; capture the P0 baseline (`sec/op`, `B/op`, `allocs/op`) and a fresh
`make memprofile` on `8552595443` — this is the reference all later commits compare against, and it
reconfirms the top alloc sources (expect `readFieldPaths` to dominate).

**Result:** baseline **1.523 s/op ±1%, 791.5 MiB/op, 20.75M allocs/op** (8552595443, M4 Pro, 10x×10).
Alloc profile confirms `readFieldPaths` #1 at **56.6%**, then quantized-float box 5.4%, `onCDemoPacket`
pointer/tuple allocs 4.6%, noscale/signed boxes ~3% each, QAngle `[]float32` 2.2%.

---

## Phase 1 — safe perf wins

### P1.1 — `pendingMessage` value-slice + typed `sort.Stable`
- **Now:** `onCDemoPacket` builds `[]*pendingMessage`, allocating a pointer per embedded message
  (~5.4M/run), then `sort.Sort` (non-stable) every packet (demo_packet.go:62-89).
- **Change:** make `pendingMessages` a `[]pendingMessage` value slice reusing a parser-level
  `pendingMsgBuf[:0]`; change `priority()` to a value receiver; use **typed `sort.Stable`** (not
  `sort.SliceStable`, which adds ~1.46M reflection allocs). Store the buffer back on both the error
  and success return paths.
- **Clarity:** dispatches embedded messages in file order, no buffering (InputSourceProcessor.java:198-240).
- **Caveat:** manta deliberately reorders across priority buckets — `sort.Stable` gives
  *file-order-within-bucket* determinism; do **not** describe it as full clarity parity, and do **not**
  remove the priority sort (would change dispatch order and break golden counts).
- **Impact (measured):** 11.0M→10.47M allocs, 315→295 MB. **Reentrancy verified safe:** CDemoPacket is
  only dispatched via `callByDemoType`, never nested in `callByPacketType`.
- _This is the change a review agent prototyped + reverted; reimplement cleanly here with a benchmark._
- **Result:** sec/op 1.523→1.514 (−0.62%, p=0.000), B/op 791.5→763.8 MiB (−3.51%), allocs/op
  20.75M→20.06M (−3.33%, −0.69M). go test green. ✅

### P1.2 — `outerMessage` by value
- **Now:** `readOuterMessage` heap-allocates `&outerMessage{}` per message (parser.go:238); single
  consumer is `Start()` (parser.go:142), value never retained past the iteration.
- **Change:** return by value (or reuse a `p.outerMsg` field); update the `var msg *outerMessage` decl.
- **Impact (measured):** −557K allocs, −17 MB. Zero risk.
- **Result:** allocs/op 20.06M→20.03M (−0.19%, −30K), B/op −0.15%; sec/op +0.32% (run-to-run thermal
  noise, +5 ms abs). Smaller than the −557K review estimate — the 70 MB replay has far fewer outer
  messages than 2159568145. go test green. ✅

### P1.3 — snappy scratch-buffer reuse
- **Now:** `snappy.Decode(nil, buf)` per compressed outer message (parser.go:232) allocates fresh.
- **Change:** pass `p.snappyScratch[:cap(...)]`; snappy@v0.0.3 reuses dst when `dLen <= len(dst)`.
- **Risk (keep medium):** the decompressed buffer becomes `msg.data` passed to handlers. If any
  CDemo* handler retains a subslice across messages, reuse corrupts it. **The full `go test` suite is
  the gate** — do not rely on assertion that string tables copy; verify empirically.
- **Impact (measured):** −9 MB/op.
- **Result:** B/op 762.6→700.6 MiB (**−8.14%** — snappy full-packet buffers), sec/op 1.519→1.511
  (−0.49%), allocs/op −0.12% (−30K). Full 48-replay suite green → reuse aliasing is safe on the corpus. ✅

### P1.4 — modifier emit early-return when no handler
- **Now:** `emitModifierTableEvents` (modifier.go:18) allocates + unmarshals a proto per ActiveModifiers
  item per snapshot/delta even when `p.modifierTableEntryHandlers` is empty (the bench registers none).
- **Change:** `if len(p.modifierTableEntryHandlers) == 0 { return nil }` at the top of
  `emitModifierTableEvents` (covers both call sites string_table.go:126,171). Also switch
  `proto.NewBuffer(v).Unmarshal` → `proto.Unmarshal` so a reused msg is safe (Buffer.Unmarshal merges
  without reset; only safe on fresh msg).
- **Impact:** removes all modifier unmarshal from the bench hot path. Verified ActiveModifiers-only
  (never touches instancebaseline).
- **Result:** allocs/op 20.00M→18.74M (**−6.32%, −1.26M** — bigger than the profile top-12 suggested;
  the cost was spread across proto internals), B/op −7.51%, sec/op −2.15%. Did the early-return only;
  left per-item fresh-msg unmarshal on the handler path (no consumer retain risk). go test green. ✅

### P1.5 — reuse per-packet `tuples` + entity-data reader
- **Now:** `onCSVCMsg_PacketEntities` allocates `newReader(m.GetEntityData())` and
  `tuples := make([]tuple, 0, updates)` per packet (entity.go:223,244).
- **Change:** parser-level `tuples[:0]` reset each call (preserves append order → preserves the
  handlers-outer/tuples-inner emission order `expectEntityEvents` depends on); `reader.reset(buf)`
  method for the entity-data reader.
- **Impact:** −0.44M allocs. (The baseline reader at entity.go:269 is handled by P1.14, not here.)
- **Result:** B/op 647.9→571.0 MiB (**−11.87%** — the per-packet tuples backing array was large:
  updates×16B over many packets), allocs/op −0.40% (−80K), sec/op +1.03% (run variance; cumulative sec
  still −1.9% vs P0). Hoisted tuple to package-level `entityOpTuple`; reused parser-level reader+tuples. go test green. ✅

### P1.6 — hoist per-entity `fpCache`/`fpNoop` maps to class
- **Now:** `newEntity` allocates two maps per entity (entity.go:69-70); name→fieldPath depends only on
  `class.serializer`, so every entity of a class recomputes/re-stores the identical mapping.
- **Change:** move both to a single shared map on `*class` (single-goroutine parse → plain map, no
  sync.Map). The class-level `*fieldPath` is read-only/never released (matches today's retain-on-hit).
- **Clarity:** resolves name→FieldPath on the DTClass, Entity holds only state (Entity.java:88-114).
- **Impact:** −2 maps/entity + smaller Entity struct. Golden-safe (resolved fieldPath is class-invariant;
  `expectPlayer6Name`/`expectHeroEntityName` guard it).
- **Result:** allocs/op −0.16% (−30K, two maps/entity removed), B/op −0.27%, sec/op ~ (p=0.190, noise).
  Caches now shared on `*class`; single-goroutine parse so safe. go test green. ✅

### P1.7 — hoist `v(6)` debug guard out of the per-field loop
- **Now:** `readFields` calls `v(6)` twice per field path in the hot loop (field_reader.go:13,23).
- **Change:** evaluate `dbg := v(6)` once at the top of `readFields` (per call — **not** higher; debug
  tick mutates `debugLevel` mid-parse and `readFields` is called fresh per entity-update, so per-call
  re-eval preserves behavior).
- **Impact:** −~33M calls/branches (CPU only, ~0 alloc when disabled).
- **Result:** sec/op 1.509→1.476 (**−2.18%**, p=0.002 — `v(6)` was 2 calls/field across millions of
  fields), allocs/B unchanged (CPU-only). go test green. ✅

### P1.8 — buffer `stream` IO
- **Now:** `stream.readByte`→`readBytes(1)`→`io.ReadFull` per byte; 3 varints + payload per outer
  message (stream.go, parser.go:198-227).
- **Change:** in `newStream`, wrap with `bufio.NewReaderSize` **only** when the reader is not already an
  `io.ByteReader` (`*os.File` isn't; `*bytes.Reader` is → `NewParser` stays unwrapped, no double-buffer).
- **Impact:** 798→658 ms on the streaming path. **Note:** the canonical in-memory bench (P0) won't show
  this — it's a real-world streaming-path win; verify with a streaming-reader bench variant.
- **Result:** canonical in-memory bench flat (+0.46%, noise — `bytes.Reader` left unwrapped, no
  regression). Streaming `NewStreamParser(os.File)` path **1.612→1.507 s/op (−6.47%, p=0.002)**, measured
  with a throwaway streaming bench (reverted before commit). go test green. ✅

### P1.9 — reusable field-path buffer  ⭐ biggest single win
- **Now:** `readFieldPaths` (field_path.go:309-337) starts `paths := []*fieldPath{}` (cap 0) and
  `append(paths, fp.copy())` per op; this slice growth is **56–60% of all allocations** (17M objects/op,
  309 MB), independently confirmed by two agents. `fp.copy()` itself allocs ~0 (sync.Pool).
- **Change (do the three coupled findings together):** (a) stop materializing a fresh `[]*fieldPath` per
  call — reuse a buffer reset to len 0 each call; (b) make `fieldPath` carry a fixed `[7]int` value array
  so a "copy" is a cheap value store, not a pool Get/Put; (c) since `readFields` immediately iterates +
  releases, consider fusing decode inline (ops apply in stream order; decoder lookup depends only on the
  current fp value, so interleaving = resolve-all-then-decode — verified equivalent).
- **Plumbing:** `readFieldPaths(r)` / `readFields(r,s,state)` are free functions with no Parser handle —
  thread a buffer param or reuse the existing `fpPool` sync.Pool (already concurrency-safe; manta runs
  independent Parsers, so a package-global mutable buffer is **unsafe**). Keep `*fieldPath` for the cold
  `getFieldPaths()`/`Entity.Map()` paths.
- **Clarity:** one reused `S2ModifiableFieldPath` writing immutable snapshots into a fixed
  `fieldPaths[MAX_PROPERTIES=0x3FFF]` (FieldReader.java:11-14, S2FieldReader.java:50-65).
- **Impact:** allocs/op ~11M → ~4–5M. **Keep the fixed depth 7** (`[7]int`). Golden-safe (only the
  container changes; field-path values + decode order unchanged).
- **Result:** allocs/op 18.63M→10.65M (**−42.86%, −7.98M**), B/op 569.5→398.6 MiB (−30.0%), sec/op
  1.483→1.198 (**−19.22%**). The 56%-of-allocs `readFieldPaths` container is gone. Gotcha: first attempt
  regressed B/op +31% because `&fp` escaped via the indirect closure call — borrowing the accumulator
  from the pool fixed it. go test green. ✅ **(headline win)**

### P1.10 — word-at-a-time bit reader
- **Now:** `readBits` refills the accumulator one byte at a time (reader.go:50-61) with a per-byte
  bounds-check+panic in `nextByte`. `readBits` is the #1 CPU consumer.
- **Change:** refill a full `uint64` at once:
  `r.bitVal |= binary.LittleEndian.Uint64(r.buf[r.pos:]) << r.bitCount; free := (64 - r.bitCount) >> 3;
  r.pos += free; r.bitCount += free*8`, keeping a byte-loop for the final ≤8 bytes (protobuf-owned
  buffers have unknown trailing capacity → tail guard `r.pos+8 <= r.size`).
- **Correction to the naive formula:** advance is `(64 - bitCount) >> 3` whole bytes, **not**
  `8 - bitCount/8` (bitCount transiently reaches ~33–39 after a refill).
- **Safety invariant (empirically verified):** max `n` observed across real replays is exactly 32
  (`readBits(qfd.Bitcount)`, `noscaleDecoder` readBits(32), `readUBitVarFP` readBits(31)), so one uint64
  load always yields ≥32 valid bits. A 5,000-trial fuzz of byte-vs-word produced bit-identical output.
- **Add a build-time assert** that quantized `Bitcount <= 32` (quantizedfloat.go integer-encode loop can
  raise it) so a pathological field can't overflow the accumulator.
- **Impact:** large CPU win (only visible after P0). 0 allocs delta.
- **Result:** sec/op 1.198→1.058 (**−11.64%**), allocs/B flat. Word refill masks before shifting (no
  stale partial-byte bits). Added `realign()` so byte-aligned `readBytes`/`readByte` rewind the
  buffered word and stay zero-copy — this extended the fast path beyond the old `bitCount==0`, making byte
  reads zero-copy *more* often than before (hence no B/op regression). Added quantized `Bitcount<=32`
  assert. go test green. ✅ (also satisfies the P-guard zero-copy invariant; P-guard adds the test.)

### P1.11 — varints + `readByte`/`readLeUintX` from the accumulator
- **Now:** `readVarUint32/64` call `readByte` per byte (byte-aligned fast path → `nextByte`), each a
  separate bounds-check (reader.go:102-140).
- **Change:** after P1.10, route `readByte` through `readBits(8)` so varint bytes share the loaded word;
  fold the now-dead `bitCount==0` fast path; read `readLeUint32/64`/`readFloat` straight from the
  accumulator instead of `binary.LittleEndian.UintXX(readBytes(N))` (the unaligned `readBytes` path
  allocates).
- **Impact:** compounds P1.10. `BenchmarkReadVarUint32/64` (reader_test.go) measures it directly.
- **Guard:** preserve `readBytes` zero-copy aliasing (P-guard below).
- **Result:** **SKIPPED.** Routing `readByte`/`readLeUintX` through `readBits` makes them consume into
  the read-ahead accumulator, leaving `r.pos` ahead of the logical position. The reader unit tests
  (`TestReaderReplayBeginning`, `TestReaderVarints`) assert exact `r.pos` and manually set `r.pos = 0`, and
  `r.pos`/`remBytes` are part of the reader contract. The marginal gain (avoiding a `make` on rare
  *unaligned* `readLeUintX`) wasn't worth rewriting tests and risking the `pos` contract. P1.10's `realign`
  already keeps these byte reads fast and zero-copy with `pos` accurate. Reverted.

### P1.12 — flatten huffman tree to int arrays
- **Now:** `readFieldPaths` walks an interface pointer-tree per bit (`node.Right()/Left()/IsLeaf()/Value()`
  — ~2–3 interface dispatches/bit, huffman.go:9-77, field_path.go:316-332).
- **Change:** flatten to parallel int arrays `tree[node][bit]` (negative = `-ordinal-1` leaf), walked with
  no dispatch. **Build from manta's OWN `huffTree`** (manta's tie-break differs structurally from
  clarity's; codes were verified bit-for-bit identical when built from manta's table — porting clarity's
  numbering would desync).
- **Impact:** removes per-bit interface calls. 0 allocs. **Add a permanent decoder-equivalence test**
  (interface-walk vs flat path → identical ordinal + bits-consumed across all 256 prefixes).
- **Result:** sec/op 1058.5→960.2 ms (**−9.28%**), allocs/B unchanged. Flat `int32` child arrays
  (negative = leaf op) replace the per-bit interface walk; built from manta's own tree. Added
  `TestHuffmanFlatMatchesTree` (flat ≡ interface tree). go test green. ✅ (under 1 s now.)

### P1.13 — 8-bit field-op lookup table
- **Now:** even flattened, each op is a 3–8 iteration bit-walk.
- **Change:** precompute a 256-entry table (`bits0-7` = consumed count, `bits8-15` = op ordinal or
  fall-back node) resolving most ops in one index (clarity FieldOpHuffmanTree.java:19-46,
  BitStream64.java:53-84). Depends on P1.12.
- **Safety-critical correction:** manta's reader is a **consuming** stream — the peek must read only
  available bits and **zero-pad** the rest, never over-read (`FieldPathEncodeFinish` = '10', 2 bits, fires
  near buffer end and would panic `nextByte` past `r.size`). Add a `reader.peekBits(n)` primitive and
  unit-test it at the <8-bits-remaining boundary. Also assert `huffTree` is never mutated
  (`swapNodes`/`addNode`) after the table is built.
- **Note:** the "99.7%" is clarity's runtime figure; manta's static weights resolve ~98.4% within ≤8 bits
  (28/40 ops have >8-bit codes, max 17) — still a large majority. Don't cite 99.7% as verified.
- **Result:** sec/op 960.2→796.0 ms (**−17.10%**), allocs/B unchanged. 256-entry lookup resolves most
  ops in one index; `peekBits` zero-pads (never over-reads), flat-walk fallback for >8-bit codes.
  `TestHuffmanLookupMatchesWalk` (5000 random streams) confirms lookup ≡ walk (op + bits consumed). go test green. ✅

### P1.14 — decode class baseline once, clone per entity
- **Now:** every entity CREATE re-parses raw baseline bytes from scratch
  (`readFields(newReader(baseline), …)`, entity.go:269); profiled at ~2.05% — **~4× the cost of the
  actual update decode** and more than `newEntity` itself.
- **Change:** decode each class's baseline once into a template state; clone into new entities. **Without
  COW** (manta has none today) the clone is a recursive deep-copy of the `*fieldState` tree (must
  faithfully clone nested `*fieldState` — 10 type-assert sites rely on leaf-vs-`*fieldState`; mutating one
  entity must never corrupt the template/siblings; reproduce the `set` rule that won't overwrite a slot
  holding a `*fieldState`).
- **Clarity:** caches decoded baseline per class, `getBaseline().copy()` with copy-on-write
  (Entities.java:655-676, NestedArrayEntityState.java:28-43,219-224).
- **Impact:** net win = clone cheaper than re-decode (it is, especially after P3.1 enables cheap COW).
  Golden-critical: cloned+overlaid state must be value-identical to today's fresh decode.
- **Result:** sec/op 796.0→766.5 ms (−3.71%), B/op 398.6→378.1 MiB (−5.14%), allocs/op 10.65M→10.26M
  (−3.61%, −0.39M). Baseline decoded once per class into a template, `clone()`d per entity (shares
  immutable leaf values, deep-copies nested states). Templates invalidated (`clear`) on instancebaseline
  update for correctness. go test green. ✅

### P-guard — zero-copy `readBytes` invariant (correctness guard, do alongside P1.10/P1.11)
- `readBytes` aligned path returns `r.buf[r.pos-n:r.pos]` aliasing the protobuf buffer (reader.go:81);
  `demo_packet.go` stores these into `pendingMessage` and parses them **later**, and sendtable/string
  reads depend on the aliasing. A clarity-style padded-copy reader would corrupt this. Add an explicit
  invariant + test that byte-aligned `readBytes` stays zero-copy after the word-reader rewrite.
- **Result:** Added `TestReaderReadBytesZeroCopy` — verifies byte-aligned `readBytes` aliases the buffer
  (zero-copy) both at `bitCount==0` and through the word reader's `realign` path (non-zero multiple of 8).
  Test-only; bench unchanged. The realign mechanism itself landed in P1.10. go test green. ✅

---

## Phase 2 — correctness (independent; can interleave with Phase 1)

### P2.1 — string-table additive index fix  ⭐ real latent bug
- **Now:** non-increment branch computes **absolute** `index = readVarUint32() + 1` (string_table.go:226).
- **Change:** **additive** `index += readVarUint32() + 2` (clarity S2StringTableEmitter.java:91; matches
  manta's own entity decoder entity.go:247).
- **Evidence (experiment):** on the bench replay the additive scheme is strictly monotonic within every
  blob; the current absolute scheme produces 2,309 impossible non-increasing steps and 4,469 divergent
  indices (all in ActiveModifiers; e.g. manta=11 where correct=50). Invisible to goldens only because no
  test resolves a delta-updated table by index. Applying the one-line fix: `TestParseStringTable*` and 5
  golden `TestMatch*` all pass with identical values. Init stays `-1`; do **not** switch to `readUBitVar`.
- **Result:** Applied the one-line additive fix. Full `go test ./...` green — all golden values identical
  (no test resolves a delta-updated table by index). Fixes the latent ActiveModifiers mis-indexing. ✅

### P2.2 — string-table: fail loud
- `parseStringTable` defers `recover()` and returns partial `items` silently (string_table.go:181-186);
  callers see success and apply a half-populated table. Propagate the error (callbacks already return
  `error`). `recover` never fires on the bench replay → no healthy-replay regression. Add a truncated-blob
  unit test. (Also reconcile the inconsistency: `UpdateStringTable` `_panicf`s on a missing table id at
  string_table.go:139 — harsher than clarity, which skips unknown ids.)
- **Result:** `parseStringTable` now returns an error (recover → error); both callers propagate it. Added
  `TestParseStringTableTruncated`. go test green (recover never fires on healthy replays). Left the
  `UpdateStringTable` `_panicf`-on-missing-table as-is (separate behavior call). ✅

### P2.3 — `getEventKey` off-by-one
- `if f.i > len(e.m.GetKeys())` lets `f.i == len` pass → `GetKeys()[f.i]` panics (game_event.go:166).
  Change to `>=`. One char; zero golden risk (well-formed descriptors never hit it).
- **Result:** Changed `>` to `>=` (game_event.go). go test green; zero golden change. ✅

### P2.4 — modifier: skip empty/deleted entries
- `emitModifierTableEvents` unmarshals every item including empty (`Value == []byte{}`), raising all-zero
  messages to handlers. Add `if len(item.Value) == 0 { continue }` (modifier.go:19), matching clarity
  `if (value != null)`. A real new modifier is never zero-length on the wire.
- **Result:** Added the empty-value skip. go test green; golden-neutral (the suite registers no modifier
  handler). ✅

### P2.5 — combat-log `Type()/TypeName()/String()` descriptor-driven
- These hardcode `keys[0].GetValByte()` (game_event.go:37,41,46), bypassing the descriptor field map. Resolve
  `type` via the descriptor (clarity S1CombatLogIndices.java:8) and route through `GetInt32`-style dispatch
  (not raw `GetValByte`). On current descriptors `type` is index 0 → output unchanged; these methods aren't
  called by any manta source/test → zero golden risk. **Fix all three** (the review noted `String()` too).
- **Result:** All three now resolve `type` via `GetInt32("type")` (descriptor + typed dispatch). go test
  green — and `GetInt32("type")` is golden-guarded (the combat-log test asserts it returns no error on
  every event), so this is value-identical on current descriptors. ✅

### P2.6 — decoder forward-compat additions (never-occur today → golden-neutral)
Verified zero occurrences across all 39 build replays, so safe insurance:
- `CUtlBinaryBlock`: no decoder → falls to varint (would desync). Add `n:=readVarUint32(); readBytes(n)`
  (clarity CUtlBinaryBlockDecoder).
- `Quaternion`: add to **`fieldTypeFactories`** as `vectorFactory(4)` (128 bits; not `fieldTypeDecoders`).
- `int64`: maps to 32-bit `readVarInt32` (truncates); add `signed64Decoder → readVarInt64` (exists). Note:
  changes stored dynamic type int32→int64 — a `.(int32)` consumer would break (forward-compat, not asserted).
- `ResourceId_t → unsigned64Decoder`, `CGlobalSymbol → stringDecoder`, `HSequence → readVarUint32()-1`
  (see P2.7 for the value-changing nuance). **Drop** `CBaseVRHandAttachmentHandle` (already correct — no-op).
- Place each in the correct map; note `findDecoderByBaseType` (variable-array childDecoder) consults only
  `fieldTypeDecoders`, not factories.
- **Result:** Added `CUtlBinaryBlock` (varint+bytes), `Quaternion` (vectorFactory4), `ResourceId_t`
  (unsigned64), `CGlobalSymbol` (string), and `int64`→`signed64Decoder` (full 64-bit varint). The first
  four never occur on the corpus (zero behavior change). `int64` affects only `m_nTotalDamageTaken`
  (~175k decodes, ≤1 byte today so the value is unchanged, but `Get` now returns `int64` not `int32`).
  Dropped CBaseVRHandAttachmentHandle (already correct). go test green (identical golden values). ✅

### P2.7 — HSequence / HeroID_t decode parity (value-changing, suite-gated)
- `HSequence` is REAL (644k decodes): manta stores `value`, clarity `value-1` → off-by-one on every
  HSequence field today. `HeroID_t` REAL (1536× on builds 6600/6601): clarity uses signed varint (same bits,
  different value for negative ids). Both change consumer-visible output but no golden asserts them. Decide
  storage signedness (HSequence `value-1` underflows if stored unsigned at value 0). **Run the full suite
  incl. 6600/6601 replays; accept only if all green.**
- **Result:** Full suite green (incl. 6600/6601 replays). HSequence → `int32(varuint)-1`, HeroID_t →
  signed varint. Same bits consumed (no desync); neither is asserted, so goldens are identical. Changes
  live consumer-visible values to match clarity. ✅

### P2.8 — BloodType fixed-8 decode (risky)
- `m_nBloodType` (109k× on builds 6600/6601, **in the golden suite**) currently a varint; clarity reads a
  fixed 8-bit. Identical bits only when value < 128; if any ≥127, widths differ → desync → broken goldens.
  **Suite-gated:** apply only if all goldens (esp. 6600/6601) stay green; otherwise drop.
- **Result:** **KEPT** — full suite green, so every `m_nBloodType` in the corpus is < 128 (fixed-8 ≡
  varint, no desync). Now matches clarity's encoding and is correct for future values ≥ 128. ✅

### P2.9 — QAngle precise/32/0-bit forward-compat
- `qangle_pitch_yaw` doesn't handle bitCount ∈ {0,32} (raw float); no `qangle_precise` (20-bit) handling →
  would fall to the coord path and desync. Zero occurrences on the corpus (observed bitcounts {0,8,13}) →
  pure forward-compat. Add the clarity special-cases (QAnglePitchYawOnly/Precise/NoScale decoders).
- **Result:** Rewrote `qangleFactory` from clarity's actual decoders: `qangle_pitch_yaw` bc∈{0,32}→raw
  floats; `qangle_precise`→3 flags + 20-bit angles; general bc==32→raw floats. Current fields (bc {8,13})
  keep identical behavior. go test green; the new branches are dormant on the corpus (forward-compat). ✅

### P2.10 — mana/runetime patch sentinel guards
- Mana (builds ≤954) and simtime/runetime (all builds) patches apply unconditionally (field_patch.go:51-78);
  clarity guards on the sentinel `±3.4028235e38` bounds. Adding the guards is a no-op on the corpus
  (robustness/parity). **Keep manta's bespoke 4-bit `runeTimeDecoder`** (Outlanders case) and the `/30`
  simtime API — guards only; touching the runetime decode math risks goldens.
- **Result:** Guarded the mana and runetime patches on the `±MaxFloat32` sentinel bounds (clarity-style);
  kept the 4-bit runetime decoder and `/30` simtime. These fields always carry the sentinel, so the
  guards always fire → go test green, identical goldens. Robustness/parity for edge-of-range fields. ✅

### P2.11 — outer-message size sanity bound
- `readOuterMessage` passes the size varint straight to `stream.readBytes` which does `make([]byte, n)` with
  no bound (parser.go:219, stream.go:26) → a corrupt/huge varint can OOM before `io.ReadFull` errors. Add a
  max-size guard (safely above the largest legitimate full packet) returning an error. Golden-neutral.
- **Result:** Added a 256 MiB `maxOuterMessageSize` guard before `readBytes` (every corpus replay is
  ≤70 MB total, so no single message approaches it). go test green; only rejects corrupt sizes. ✅

### P2.12 — field-path depth-7 guard + comment
- `Push*` ops do `fp.last++` then index `fp.path[fp.last]` with no guard → depth>6 panics with a raw
  out-of-range (field_path.go). Clarity bounds at 7 and fails loudly (S2LongFieldPathFormat.java:7-58). Keep
  7, add a cheap descriptive guard + comment. No behavior change (nothing exceeds 7 today). **Must stay `[7]int`
  when P1.9 lands.**
- **Result:** Added `maxFieldPathDepth = 7` const + comment citing `S2LongFieldPathFormat`, used for the
  `[maxFieldPathDepth]int` path array. The fixed array already fails loudly (recovered bounds-check) on
  overflow, so no hot-path guard was added. No behavior change. go test green. ✅

---

## Phase 3 — invasive (after Phase 1 re-baseline)

### P3.1 — typed entity state (eliminate per-field boxing)  ⭐
- **Now:** `fieldState.state` is `[]interface{}`; every decoder returns `interface{}` (field_decoder.go:7).
  In Go, boxing float32 / `[]float32` / large uint64 / int32 (>255) into `interface{}` **always** heap-allocates,
  on the hot write path. Measured: quantized float box ~10–12%, QAngle `[]float32` ~4.5–5.5%, signed int box
  ~2%, unsigned int box — together **~20% of all allocs**.
- **Change (internal only):** store scalars unboxed — e.g. a tagged-union cell `{kind uint8; f float32;
  i uint64; ref interface{}}` (strings/sub-state use `ref`); decoders write into the typed lane; box **lazily**
  in `Entity.Get`/`GetFloat32` (reads are rare vs writes). For vectors, store per-element float32 cells in the
  nested `fieldState` and reassemble the slice on `Get`.
- **Scope reality (verifier):** `*fieldState` shares the same `[]interface{}` slots as leaf values, distinguished
  by `.(*fieldState)` in 10 sites (field.go, field_state.go) plus `Map()/getFieldPaths` — all must be reworked.
  This is **not** a clarity mirror (clarity's `Object[]` also boxes; the JVM just amortizes via TLAB/escape
  analysis) — it's a manta-specific optimization. risk medium, effort large.
- **Golden-critical:** `Get("m_flMaxMana").(float32)` (manta_test.go:694) must stay bit-identical — `Get`/
  `GetFloat32` must still return a boxed `float32`. Folds in P-decoder findings (quantized/qangle/int boxing)
  which deliver 0 allocs alone and explicitly depend on this. Also unlocks cheap COW for P1.14.
- **Result:** Implemented as a **24-byte tagged-union `cell`**
  `{ref interface{}; num uint32; kind cellKind}`. Scalars (float32/int32/uint32/bool/≤32-bit uint64
  handles) live inline in `num` — zero write-path alloc; reference values (string/`[]float32`/`[]byte`) and
  the rare genuinely-64-bit ints (CStrongHandle/fixed64/int64/steamids) go in `ref`; nested tables in `ref`
  as `*fieldState`. Decoders now return `cell`; values box lazily in `cell.iface()` only on `Get` (rare).
  Public API unchanged — `Get` still returns `interface{}` with the exact dynamic type (entity_test's
  `int32`/`uint64`/`bool`/`string` assertions and `expectHeroEntityMana` float all pass; the >32-bit steamid
  proves 64-bit values aren't truncated). **vs P2: allocs −57.0% (10.26M→4.41M), sec −4.2%, B/op +2.6%**
  (residual is the qangle/Vector `[]float32` backing arrays, left boxed to respect no-API-change on `Get`).
  Tried a 32-byte `uint64`-num cell first (no 64-bit boxing) but it cost +12.9% B/op; the 24-byte variant
  recovers that for just +1% allocs since 64-bit values are rare (profile-confirmed). go test ./... green. ✅

---

## Deferred / parked

- **string-table Items map → dense slice** (perf, depends P2.1): indices become dense after the additive fix;
  replicate clarity's `setValueForIndex(index<len)` vs `append(index==len)` and partial-update semantics
  exactly. Modest (string tables ~0.3–0.5% of allocs); mostly locality.
- **integrate `CDemoStringTables` full dumps** (correctness, depends P2.1): manta drops them
  (string_table.go:65); clarity reconciles (BaseStringTableEmitter.java:110-145). Implement as a **silent
  reconcile/extend only** (no event re-emission, else modifier events double-fire on every full dump).
  Useful for seeking/robustness.
- **PARK — combat-log name-resolution helper / `CombatLogEntry`**: the single biggest combat-log
  usability/correctness gap (resolve `*_name` indices via `CombatLogNames`; clarity's `idx==0 → null`). **Blocked
  by strictly-no-API-change** (adds `combat_log.go` + new exported type). **Punted for this effort (decided)
  — out of scope; revisit only in a future, intentionally API-bumping effort.**
- **PARK — S2 HLTV `CMsgDOTACombatLogEntry` path**: clarity unifies S1/S2/bulk (CombatLog.java:48-81); manta
  has the proto slots but wires no consumer. Blocked (new exported hook); also S2 uses `hasIndex` name
  resolution, not the S1 `idx==0` rule.
- **PARK — VTProtobuf zero-reflection unmarshal**: legacy `golang/protobuf` already delegates to the V2
  reflective decoder (library swap = zero gain, verified). VTProtobuf shrinks only the **envelope** decode
  (the PacketEntities envelope is small; its 1.7 GB cost is the downstream entity-field decoder, not the
  unmarshal). Large effort, medium payoff — prototype on `CSVCMsg_PacketEntities`+`CNETMsg_Tick` first if pursued.

## Cumulative results (vs P0 baseline) — filled in during execution

| After goal | sec/op | B/op | allocs/op | go test |
|------------|--------|------|-----------|---------|
| P0 baseline | 1.523 s ±1% | 791.5 MiB | 20.75M | PASS |
| P1.1 pendingMessage value-slice | 1.514 s (−0.6%) | 763.8 MiB (−3.5%) | 20.06M (−3.3%) | PASS |
| P1.2 outerMessage by value | 1.519 s (−0.3%) | 762.6 MiB (−3.7%) | 20.03M (−3.5%) | PASS |
| P1.3 snappy scratch reuse | 1.511 s (−0.8%) | 700.6 MiB (−11.5%) | 20.00M (−3.6%) | PASS |
| P1.4 modifier early-return | 1.479 s (−2.9%) | 647.9 MiB (−18.1%) | 18.74M (−9.7%) | PASS |
| P1.5 reuse tuples + reader | 1.494 s (−1.9%) | 571.0 MiB (−27.9%) | 18.66M (−10.1%) | PASS |
| P1.6 hoist fp caches to class | 1.509 s (−0.9%) | 569.5 MiB (−28.1%) | 18.63M (−10.2%) | PASS |
| P1.7 hoist v(6) debug guard | 1.476 s (−3.1%) | 569.5 MiB (−28.1%) | 18.63M (−10.2%) | PASS |
| P1.8 buffer stream IO (streaming −6.5%) | 1.483 s (flat, in-mem) | 569.5 MiB | 18.63M | PASS |
| **P1.9 reusable field-path buffer** | **1.198 s (−21.3%)** | **398.6 MiB (−49.6%)** | **10.65M (−48.7%)** | PASS |
| **P1.10 word-at-a-time reader** | **1.058 s (−30.5%)** | 398.6 MiB (−49.6%) | 10.65M (−48.7%) | PASS |
| P1.11 varints from accumulator | _skipped (breaks reader pos contract)_ | — | — | — |
| **P1.12 flatten huffman tree** | **0.960 s (−37.0%)** | 398.6 MiB (−49.6%) | 10.65M (−48.7%) | PASS |
| **P1.13 8-bit op lookup table** | **0.796 s (−47.7%)** | 398.6 MiB (−49.6%) | 10.65M (−48.7%) | PASS |
| **P1.14 baseline decode-once clone** | **0.766 s (−49.7%)** | **378.1 MiB (−52.2%)** | **10.26M (−50.6%)** | PASS |
| P-guard zero-copy readBytes test | test-only (no perf change) | — | — | PASS |
| **Phase 1 total vs P0** | **−49.7%** | **−52.2%** | **−50.6%** | PASS |
| Phase 2 (P2.1–P2.12, correctness) | ~ flat | ~ flat | ~ flat | PASS |
| **End of Phase 2 vs P0** | **−49.7%** | **−52.2%** | **−50.6%** | PASS |
| P3.1 typed entity state (de-box) | 0.734 s (−4.2% vs P2) | 388.2 MiB (+2.6% vs P2) | 4.41M (−57.0% vs P2) | PASS |
| **End of Phase 3 vs P0** | **−51.8%** | **−51.0%** | **−78.8%** | PASS |

**Phase 2 (correctness) notes:** all 12 goals landed, full suite green with identical golden assertions,
and no perf regression (sec/op, B/op, allocs/op all statistically flat vs end of Phase 1). Highlights:
P2.1 fixed the real string-table additive-index bug; P2.6–P2.9 added forward-compat decoders
(CUtlBinaryBlock, Quaternion, int64-64bit, HSequence/HeroID_t/BloodType aligned to clarity, QAngle
precise/noscale); P2.2/P2.3/P2.11/P2.12 hardened error paths. P2.7/P2.8 change live field *values* to
match clarity (not asserted by goldens) — flagged for review.
