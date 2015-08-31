package manta

var huf HuffmanTree

func init() {
	if huf == nil {
		huf = newFieldpathHuffman()
	}
}

type Properties struct {
	KV map[string]interface{}
}

func NewProperties() *Properties {
	return &Properties{
		KV: map[string]interface{}{},
	}
}

func (p *Properties) Merge(p2 *Properties) {
	for k, v := range p2.KV {
		p.KV[k] = v
	}
}

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
