package manta

func ReadProperties(r *Reader, ser *dt) (result map[string]interface{}) {
	// Return type
	result = make(map[string]interface{})

	// Generate the huffman tree and fieldpath
	huf := newFieldpathHuffman()
	fieldPath := newFieldpath(ser, &huf)

	// Get a list of the included fields
	fieldPath.walk(r)

	// iterate all the fields and set their corresponding values
	for _, f := range fieldPath.fields {
		if f.Field.Serializer.Decode == nil {
			result[f.Name] = r.readVarUint32()
			_debugfl(6, "Decoded default: %d %s %s %v", r.pos, f.Name, f.Field.Type, result[f.Name])
			continue
		}

		if f.Field.Serializer.DecodeContainer != nil {
			result[f.Name] = f.Field.Serializer.DecodeContainer(r, f.Field)
		} else {
			result[f.Name] = f.Field.Serializer.Decode(r, f.Field)
		}

		_debugfl(6, "Decoded: %d %s %s %v", r.pos, f.Name, f.Field.Type, result[f.Name])
	}

	return result
}

// An ordered map
type omap struct {
	size   int
	keys   []string
	values []interface{}
}

// Creates a new ordered map.
func newOmap() *omap {
	return &omap{
		size:   0,
		keys:   []string{},
		values: []interface{}{},
	}
}

// Adds an element to the ordered map.
func (o *omap) add(k string, v interface{}) {
	o.keys = append(o.keys, k)
	o.values = append(o.values, v)
	o.size++
}

// Prints the ordered map, in order.
func (o *omap) print() {
	for i := 0; i < o.size; i++ {
		_debugf("[%d] %s = %v", i, o.keys[i], o.values[i])
	}
}

// Converts the ordered map to map[string]interface{}
func (o *omap) toMap() map[string]interface{} {
	result := make(map[string]interface{})
	for i := 0; i < o.size; i++ {
		result[o.keys[i]] = o.values[i]
	}
	return result
}
