package manta

var huf HuffmanTree

func init() {
	if huf == nil {
		huf = newFieldpathHuffman()
	}
}

// Properties is an instance of a set of properties containing key-value data.
type Properties struct {
	KV map[string]interface{}
}

// Creates a new instance of Properties.
func NewProperties() *Properties {
	return &Properties{
		KV: map[string]interface{}{},
	}
}

// Merge another set of Properties into an existing instance. Values from the
// other (merging) set overwrite those in the existing instance.
func (p *Properties) Merge(p2 *Properties) {
	for k, v := range p2.KV {
		p.KV[k] = v
	}
}

// Fetch a value by key.
func (p *Properties) Fetch(k string) (interface{}, bool) {
	v, ok := p.KV[k]
	return v, ok
}

// Reads properties using a given reader and serializer.
func ReadProperties(r *Reader, ser *dt) (result *Properties) {
	// Return type
	result = NewProperties()

	// Create fieldpath
	fieldPath := newFieldpath(ser, &huf)

	// Get a list of the included fields
	fieldPath.walk(r)

	// iterate all the fields and set their corresponding values
	for _, f := range fieldPath.fields {
		_debugfl(6, "Decoding field %d %s %s", r.pos, f.Name, f.Field.Type)
		// r.dumpBits(1)

		if f.Field.Serializer.Decode == nil {
			result.KV[f.Name] = r.readVarUint32()
			_debugfl(6, "Decoded default: %d %s %s %v", r.pos, f.Name, f.Field.Type, result.KV[f.Name])
			continue
		}

		if f.Field.Serializer.DecodeContainer != nil {
			result.KV[f.Name] = f.Field.Serializer.DecodeContainer(r, f.Field)
		} else {
			result.KV[f.Name] = f.Field.Serializer.Decode(r, f.Field)
		}

		_debugfl(6, "Decoded: %d %s %s %v", r.pos, f.Name, f.Field.Type, result.KV[f.Name])
	}

	return result
}
