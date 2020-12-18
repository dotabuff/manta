package manta

import (
	"strings"
)

func readFields(r *reader, s *serializer, state *fieldState) {
	fps := readFieldPaths(r)

	for _, fp := range fps {
		decoder := s.getDecoderForFieldPath(fp, 0)

		if v(6) {
			name := strings.Join(s.getNameForFieldPath(fp, 0), ".")
			typ := s.getTypeForFieldPath(fp, 0)
			field := s.getFieldForFieldPath(fp, 0)
			_debugf("NEW reading ser=%s path=%s pos=%s name=%s type=%s decoder=%s model=%s", s.name, fp.String(), r.position(), name, typ, _nameof(decoder), field.modelString())
		}

		val := decoder(r)
		state.set(fp, val)

		if v(6) {
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

			_debugf(" => %#v", val)
		}

		fp.release()
	}
}
