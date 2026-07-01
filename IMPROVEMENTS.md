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
| P4 baseline (re-measured at fbca7ed) | 0.727 s | 388.2 MiB | 4.41M | PASS |
| P4.1 deep-copy mutable baseline leaves | 0.755 s (~) | 389.4 MiB (+0.3%) | 4.478M (+1.5%) | PASS |
| P4.2 clear reused tuple/pending buffers | 0.765 s (~thermal) | 389.4 MiB (flat) | 4.478M (flat) | PASS |
| P4.3 guard skipBits underflow | 0.763 s (flat) | 389.4 MiB (flat) | 4.478M (flat) | PASS |
| P4.4 fix debug position | 0.761 s (flat) | 389.4 MiB (flat) | 4.478M (flat) | PASS |
| P4.5 lock value-changing decoders (tests+docs) | 0.761 s (flat) | 389.4 MiB (flat) | 4.478M (flat) | PASS |
| P4.6 clear buffers on error paths (follow-up) | 0.732 s (flat) | 389.4 MiB (flat) | 4.478M (flat) | PASS |
| **End of Phase 4 vs P3** (cooled re-measure) | 0.748 s (~, p=0.218) | 389.4 MiB (+0.3%) | 4.478M (+1.5%) | PASS |
| **End of Phase 4 vs P0** | **−50.9%** | **−50.8%** | **−78.4%** | PASS |

**Phase 4 cost (cooled, deterministic):** the only real cost of the post-review correctness/robustness
fixes is **+1.54% allocs (+68K)** from P4.1's per-entity baseline mutable-leaf isolation; B/op +0.33%,
sec/op flat (the intermediate sec wobble was thermal drift across the back-to-back bench run, confirmed
by a cooled re-measure: 744→748 ms, p=0.218). The skipBits guard, buffer clearing, debug fix, and
decoder-lock tests are all free.

**Phase 2 (correctness) notes:** all 12 goals landed, full suite green with identical golden assertions,
and no perf regression (sec/op, B/op, allocs/op all statistically flat vs end of Phase 1). Highlights:
P2.1 fixed the real string-table additive-index bug; P2.6–P2.9 added forward-compat decoders
(CUtlBinaryBlock, Quaternion, int64-64bit, HSequence/HeroID_t/BloodType aligned to clarity, QAngle
precise/noscale); P2.2/P2.3/P2.11/P2.12 hardened error paths. P2.7/P2.8 change live field *values* to
match clarity; these are owned as intentional API changes (Decision A) and locked by
decoder-representation tests in Phase 4 (P4.5).

---

## Phase 4 — post-review fixes

Addresses external review feedback on the branch. P4.1/P4.2/P4.3 fix real regressions the branch
introduced; P4.4 is a debug-only cleanup; P4.5 documents and locks the deliberate API changes.
Re-baselined at `fbca7ed`; each step benchstat'd vs the previous so the *cost* of these correctness
fixes is explicit.

### P4.1 — deep-copy mutable baseline leaves in clone()
- **Issue:** `clone()` (field_state.go) shallow-copies cells and only deep-copies nested `*fieldState`.
  Cells whose `ref` holds `[]float32` (vectors) or `[]byte` (binary blobs) are shared across every entity
  cloned from the cached class baseline (introduced by P1.14 decode-once + carried through P3.1). A caller
  mutating a slice from `Entity.Get`/`Map` can corrupt the baseline template and sibling entities. Before
  the branch, baselines were re-decoded per entity, so this aliasing didn't exist.
- **Fix:** type-switch in `clone()` and deep-copy `[]float32`/`[]byte` leaves; strings and boxed 64-bit
  scalars are immutable and stay shared. Only fires for baseline vector/blob leaves, at entity-create.
- **Result:** **cost** allocs/op 4.410M→4.478M (+1.54%, +68K — the per-entity baseline leaf copies),
  B/op +0.33%, sec/op ~ (p=0.105). go test green, identical goldens (deep-copy yields equal slices). ✅

### P4.2 — clear reused tuple/pending buffers after dispatch
- **Issue:** `p.entityTuples` (entity.go) and `p.pendingMsgBuf` (demo_packet.go) are stored back with full
  length and stale `[len:cap]` entries, retaining `*Entity` pointers (and their de-boxed state) and inner
  packet buffers an extra packet — stale slots from a larger prior packet can pin *deleted* entities.
- **Fix:** `clear()` the used entries and store the slice back at `[:0]`. Clearing the full written length
  each packet keeps `[len:cap]` zero across packets, so no stale refs accumulate.
- **Result:** allocs/op and B/op flat (clear is alloc-free); sec/op +1.4% nominal but consistent with the
  thermal drift seen across the whole P4 run (744→754→765 ms), not the change. go test green. ✅

### P4.3 — guard skipBits underflow + truncated-stream test
- **Issue:** `peekBits` zero-pads at EOF and `skipBits` blindly subtracts, so on truncated/corrupt input a
  lookup entry can consume more bits than are buffered, underflowing `bitCount` (uint32 wraps huge) and
  spinning on garbage ops instead of failing. The old per-bit walk hit a clean panic → parser error.
  Cannot affect well-formed replays (the fast path never under-runs on valid streams). The committed
  huffman test used 8-byte buffers and never hit the `<8 bits remaining` case.
- **Fix:** `skipBits` panics cleanly when `n > bitCount` (caught by the parser recover → error), restoring
  fail-fast. Add a huffman test that truncates an op stream and asserts a clean error.
- **Result:** all metrics flat (sec p=0.631, B/op p=0.971, allocs p=0.057 — the guard is free).
  `TestReadFieldPathsTruncated` confirms a truncated stream now panics cleanly instead of looping
  forever (an all-zero buffer = unbounded PlusOne run). go test green. ✅

### P4.4 — correct debug position() after word refill
- **Issue:** `position()` (reader.go) still assumes `bitCount <= 8`, but the word reader can leave 56/48/…
  Wrong verbose-debug output only; no parsing impact.
- **Fix:** compute the logical bit position as `pos*8 - bitCount`.
- **Result:** flat (debug-only; called only under the `v(6)` guard, off in bench/test). go test green. ✅

### P4.5 — lock value-changing decoders + document Decision A
- **Decision A (owned):** the branch deliberately changes what `.Get` returns for a few fields, to match
  clarity's correct representation. We are keeping these and owning them as intentional behavior changes:
  - `int64` fields: `int32` (truncated >32 bits) → **`int64`** (fixes truncation; type change is forced by
    the fix).
  - `HeroID_t`: `uint32` → **`int32`** (signed, clarity parity).
  - `HSequence`: `uint32` → **`int32`, value − 1** (−1 = "none" handle, clarity parity).
  - `BloodType`: stayed `uint64`; only the *encoding* changed (fixed-8 vs varint), value identical on the
    corpus — not an API change.
  Downstream callers doing concrete type assertions on those specific fields may need updates.
- **Fix:** add decoder-level representation tests asserting the exact dynamic type **and** value for each
  (deterministic, not replay-dependent), so the intended downstream-visible values are locked in CI — this
  is what the review asked for ("prove values are what you intend, not just no-desync"). Also assert the
  inline-vs-boxed uint64 split round-trips, including a `> 2^32` value to prove no truncation.
- **Result:** added `field_decoder_test.go` with `TestDecoderRepresentations` (exact type+value for
  hSequence `int32`−1, HeroID_t signed `int32`, int64 full `int64`, BloodType `uint64` fixed-8, inline
  `uint64` incl. `0xFFFFFFFF`, and a `>2^32` steamid + fixed64 with no truncation) and
  `TestValueChangingDecoderWiring` (locks the `fieldTypeDecoders` map entries). Test-only; bench flat.
  Decision A documented above. go test green. ✅

### P4.6 — clear reused buffers on error paths too (follow-up review)
- **Issue:** P4.2 only cleared `p.pendingMsgBuf` / `p.entityTuples` after *successful* dispatch. If
  `callByPacketType` or an entity handler returns an error, the parser field still pointed at the
  just-refilled backing array, so packet buffers / entity pointers stayed live. Low severity (the parse
  aborts on error, so it's freed at parser GC; bounded to one packet), but a real inconsistency.
- **Fix:** dispatch into a result variable (`break` / labeled `break dispatch` on error), then a single
  `clear()` + `[:0]` cleanup before a single return — so cleanup runs on success *and* error. Used this
  result-variable pattern rather than a `defer func(){…}()` closure on purpose: the closure captures the
  slice header (also stored on the Parser), forcing it to escape to the heap — a per-packet allocation in a
  path that runs thousands of times. The result-variable form is alloc-free. (Panic mid-dispatch still
  skips cleanup, but that hits the parser's top-level recover and tears the parser down, same as today.)
- **Result:** alloc- and B/op-neutral (4,478,290 allocs / 408.36M B identical to P4.5); sec flat. go test green. ✅

---

## Phase 5 — post-merge profile-driven round (branch `jcoene/goal-perf-phase5`)

Fresh CPU/alloc profiles were captured on master (post-#180) on the canonical replay
`8552595443`. The Phase 1–4 dominators (field-path slice, boxing, huffman walk) are gone;
the profile has a new shape. Same ground rules as before: no public API change, one
optimization per commit with benchstat vs the previous commit, full `go test ./...`
golden gate, commits only (no push, no PR).

### Profile evidence (master @ 0efe7e1, `TestMatchNew8552595443`, M4 Pro)

- **Alloc objects:** unaligned `reader.readBytes` 13.5% (all from `onCDemoPacket` embedded
  messages — the inner stream is bit-shifted after the leading 6-bit `readUBitVar`, so every
  message body takes the `make([]byte,n)` + per-byte loop) · `readBitsAsBytes` 11.6% (cap-0
  append growth per string-table value) · qangle `[]float32` 11.5% (known residual) ·
  protobuf `reflect.New` + `consume*Ptr` + `consumeBytes` ~30% combined · `parseStringTable`
  7.0% flat (`&stringTableItem` per item + items growth) · `fieldState.clone` 6.7% +
  `newFieldState` 2.5% (baseline clone per entity create).
- **Alloc space:** protobuf `consumeBytes` **32.7%** (128 MB — every `bytes` field is copied:
  `CDemoPacket.data`, `PacketEntities.entity_data`, `UpdateStringTable.string_data`) ·
  `fieldState.clone` 19.1% · unaligned `readBytes` 12.7%.
- **CPU:** `fieldState.set` 8.3% · `memclrNoHeapPointers` 8.3% (zeroing the fresh unaligned
  `readBytes` buffers) · `mapaccess1_fast32` ~9% incl. inlined key probe (the `p.entities`
  map in the PacketEntities update loop) · `readBits` 6.2% · decoder-lookup recursion ~3%.

### Goal summary

| ID | Goal | Type | Risk | Effort | Expected impact |
|----|------|------|------|--------|-----------------|
| P5.0 | re-baseline bench on master | infra | none | none | reference point |
| P5.1 | `readBitsAsBytes` exact prealloc + word fill | perf | low | small | −~10% allocs |
| P5.2 | demo-packet arena + word-copy unaligned `readBytes` | perf | med | small | −13% allocs, −12% B/op, CPU |
| P5.3 | `p.entities` map → dense slice | perf | low | small | −5–9% CPU |
| P5.4 | string-table item slab + prealloc + ring history + single lookup | perf | low | small | −~7% allocs |
| P5.5 | reader smalls: `readLeUintX` via accumulator when unaligned, `readString` prealloc | perf | low | small | −~1% allocs |
| P5.6 | hand-rolled envelope decode for internal-only hot messages | perf | med | medium | −25–30% B/op, −10%+ allocs |
| P5.7 | slab-allocated baseline clone | perf | med | medium | −~8% allocs |
| P5.8 | game-event eventid wire-peek, skip unmarshal when unhandled | perf | low | small | workload-dependent |
| P5.9 | prefix cache for `set`/decoder walk across sorted field paths | perf | med | medium | −5–8% CPU (attempt) |
| P5.10 | experiment: inline vec3 cell for qangle/vector residual | perf | med | small | measure; likely reject (P3.1 saw +12.9% B/op at 32B cells) |

### Cumulative results (vs P5.0 baseline) — filled in during execution

| After goal | sec/op | B/op | allocs/op | go test |
|------------|--------|------|-----------|---------|
| P5.0 baseline (master 0efe7e1) | 876.7m ±8% | 389.4 MiB | 4.478M | PASS |
| P5.1 readBitsAsBytes prealloc | 756.3m ±1% (−13.7%*) | 384.4 MiB (−1.3%) | 4.181M (−6.6%) | PASS |
| P5.2 demo-packet arena + word copy | 676.7m ±1% (−22.8%) | 336.3 MiB (−13.6%) | 3.757M (−16.1%) | PASS |
| P5.3 entities map → dense slice | 624.8m ±2% (−28.7%) | 336.3 MiB (−13.6%) | 3.757M (−16.1%) | PASS |
| P5.4 string-table slab + ring history | 618.7m ±4% (−29.4%) | 335.0 MiB (−14.0%) | 3.597M (−19.7%) | PASS |
| P5.5 reader smalls (readLeUintX, readString) | 618.8m ±2% (−29.4%) | 334.5 MiB (−14.1%) | 3.540M (−20.9%) | PASS |
| **P5.6 hand-rolled envelope decode** | **580.5m ±0% (−33.8%)** | **192.7 MiB (−50.5%)** | **2.615M (−41.6%)** | PASS |

\* the P5.0 sec/op baseline was thermally inflated (±8%); treat allocs/B as the
reliable P5.1 signal and 756m ±1% as the true current sec/op level.

### P5.6 — hand-rolled envelope decode for hot internal messages  ⭐ headline win
- **Was:** every CDemoPacket / CSVCMsg_PacketEntities / CSVCMsg_UpdateStringTable /
  CNETMsg_Tick went through the reflective protobuf unmarshal: `reflect.New` per message,
  a pointer alloc per scalar field, and — 32.7% of all allocated bytes — a fresh copy of
  every `bytes` payload. The replay's data was effectively copied twice through the proto
  layer (outer `CDemoPacket.data`, then inner `entity_data`/`string_data`).
- **Change:** `envelope_fast.go` decodes these four envelopes by hand with `protowire`,
  aliasing the payload instead of copying, and calls new scalar-arg core methods
  (`processDemoPacket`/`processPacketEntities`/`processUpdateStringTable`) split out of the
  internal handlers. **Gating:** NewParser registers exactly one internal handler per list
  before returning, so `len(list) == 1` means no user callbacks; any user registration
  reverts that type to the full protobuf path (user-visible messages own their copies, as
  before). Aliasing lifetimes: `entity_data`/`string_data` alias the packet arena (stable
  through the dispatch loop, consumed synchronously — string-table values are copied out by
  `readBitsAsBytes`); `CDemoPacket.data` aliases the outer-message buffer (stable until the
  next `readOuterMessage`; `processDemoPacket` copies bodies into the arena before dispatch).
- **Result:** B/op 334.5→192.7 MiB (**−42.39%**), allocs/op 3.540M→2.615M (**−26.13%,
  −925K**), sec/op 618.8m→580.5m (**−6.18%**, p=0.000). go test green — the goldens
  exercise the fast path (tests register OnEntity/OnGameEvent, not raw message callbacks). ✅

### P5.5 — reader smalls: `readLeUintX` via accumulator, `readString` prealloc
- **Was:** unaligned `readLeUint32/64` allocated through the `readBytes` slow path (24K
  objects via `fixed64Decoder` etc.); `readString` grew a cap-0 buffer byte by byte.
- **Change:** unaligned `readLeUint32` reads `readBits(32)` straight from the accumulator
  (bit-stream LSB-first == LE byte decode); `readLeUint64` uses two accumulator words
  (readBits is capped at 32). Aligned paths keep the zero-copy `readBytes` fast path and
  the exact `pos` contract the reader tests assert. `readString` starts at cap 32 (stack-
  allocatable when it doesn't escape).
- **Result:** allocs/op 3.597M→3.540M (**−1.59%, −57K**), B/op −0.16%, sec ~ (p=0.481).
  go test green. ✅

### P5.4 — string-table item slab + prealloc + ring history + single lookup
- **Was:** `parseStringTable` allocated a `&stringTableItem` per item plus cap-0 `items`
  append growth (217K objects); the 32-entry key history shifted the whole window with a
  `copy` per item once full; the `UpdateStringTable` apply loop did up to four map lookups
  per item.
- **Change:** preallocate `items` and back it with a single `[]stringTableItem` slab
  (bounded at 4096 cap so corrupt `numUpdates` can't OOM; pointer identity survives slab
  regrowth since items are only touched through the taken pointers). Key history becomes a
  fixed ring buffer (`histCount`-based indexing reproduces the shift-window semantics
  exactly). Apply loop hoists to one lookup.
- **Result:** allocs/op 3.757M→3.597M (**−4.26%, −160K**), B/op −0.38%, sec −0.98%
  (p=0.043). go test green (string-table goldens + full corpus). ✅

### P5.3 — `p.entities` map → dense slice
- **Was:** `p.entities` was a `map[int32]*Entity`; the PacketEntities update loop does 1–3
  lookups per entity update (~9% CPU in `mapaccess1_fast32` + inlined key probes). Deletion
  stored `nil` *into* the map, so `FilterEntity` could pass nil entities to user callbacks —
  a latent crash.
- **Change:** dense `[]*Entity` sized `1<<indexBits` (16384 slots, 128 KB per parser —
  entity indices are 14-bit by the handle encoding). `FindEntity` bounds-checks and keeps
  its signature; `FilterEntity` skips nil slots (fixes the nil-callback hazard and makes
  iteration deterministic by index instead of map-random).
- **Result:** sec/op 676.7m→624.8m (**−7.66%**, p=0.000), B/op and allocs exactly flat
  (CPU-only). go test green. ✅

### P5.2 — demo-packet arena + word-copy unaligned `readBytes`
- **Was:** the `onCDemoPacket` inner stream is bit-shifted after the leading 6-bit
  `readUBitVar`, so nearly every embedded message body hit the unaligned `readBytes` slow
  path: a fresh zeroed `make([]byte, n)` (422K objects, 49.5 MB, plus most of the 8%
  `memclrNoHeapPointers` CPU) filled one `readBits(8)` at a time.
- **Change:** (a) new `reader.readBytesInto(dst)` copies unaligned data a 32-bit word at a
  time; `readBytes`'s slow path routes through it. (b) `onCDemoPacket` carves message
  buffers from a single parser-level arena sized to `len(m.GetData())` (headers guarantee
  the payload total fits) and reused across packets. Lifetime matches `pendingMsgBuf`: the
  buffers only live until dispatch, and the protobuf unmarshal copies what it keeps.
- **Result:** sec/op 756.3m→676.7m (**−10.52%**, p=0.000), B/op 384.4→336.3 MiB
  (**−12.52%**), allocs/op 4.181M→3.757M (**−10.13%, −424K**). go test green. ✅

### P5.1 — `readBitsAsBytes` exact prealloc + word fill
- **Was:** `tmp := make([]byte, 0)` grown byte-at-a-time via `append(tmp, r.readByte())` —
  several growth allocations per string-table value (362K objects, 11.6% of allocs).
- **Change:** allocate the exact `(n+7)/8` result up front (callers retain the slice, so a
  fresh allocation per value is required) and fill via `readBits(32)` words, then a byte/bit
  tail. Bit-stream equivalence: bits are consumed LSB-first, so four `readBits(8)` calls and
  one LE-decoded `readBits(32)` yield identical bytes.
- **Result:** allocs/op 4.478M→4.181M (**−6.64%, −297K**), B/op −1.29%, sec −13.7% vs the
  noisy P5.0 run (true level ~756m ±1%). go test green. ✅
