package manta

import (
	"reflect"
	"strings"
)

func readFields(r *reader, s *serializer, state *fieldState) {
	fps := readFieldPaths(r)

	for _, fp := range fps {
		decoder := s.getDecoderForFieldPath(fp, 0)

		if v(6) {
			name := strings.Join(s.getNameForFieldPath(fp, 0), ".")
			typ := s.getTypeForFieldPath(fp, 0)
			_debugf("NEW reading ser=%s path=%s pos=%s name=%s type=%s decoder=%s", s.name, fp.String(), r.position(), name, typ, nameOf(decoder))
		}

		val := decoder(r)
		state.set(fp, val)

		if v(6) {
			if val2 := state.get(fp); !reflect.DeepEqual(val, val2) {
				_panicf("WRONG READ: %#v != %#v", val, val2)
			}

			_debugf(" => %#v", val)
		}

		fp.release()
	}
}
