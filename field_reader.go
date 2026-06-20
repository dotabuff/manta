package manta

import (
	"strings"
)

func readFields(r *reader, s *serializer, state *fieldState, fpBuf []fieldPath) []fieldPath {
	fpBuf = readFieldPaths(r, fpBuf[:0])

	// Evaluate the debug verbosity once per call rather than twice per field in
	// the hot loop. readFields is called fresh per entity update, so a mid-parse
	// debug-level change (e.g. a debug-tick) still takes effect on the next call.
	dbg := v(6)
	for i := range fpBuf {
		fp := &fpBuf[i]
		decoder := s.getDecoderForFieldPath(fp, 0)

		if dbg {
			name := strings.Join(s.getNameForFieldPath(fp, 0), ".")
			typ := s.getTypeForFieldPath(fp, 0)
			field := s.getFieldForFieldPath(fp, 0)
			_debugf("NEW reading ser=%s path=%s pos=%s name=%s type=%s decoder=%s model=%s", s.name, fp.String(), r.position(), name, typ, _nameof(decoder), field.modelString())
		}

		val := decoder(r)
		state.set(fp, val)

		if dbg {
			name := strings.Join(s.getNameForFieldPath(fp, 0), ".")
			fp2 := newFieldPath()
			b := s.getFieldPathForName(fp2, name)

			if !b {
				_panicf("GOT NO FP: name=%s fp2=%#vv", name, fp2)
			}

			if fp2.String() != fp.String() {
				_panicf("GOT FP MISMATCH: fp=%s fp2=%s", fp, fp2)
			}

			fp2.release()

			_debugf(" => %#v", val.iface())
		}
	}

	return fpBuf
}
