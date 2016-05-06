package manta

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

// Fetch a bool by key.
func (p *Properties) FetchBool(k string) (bool, bool) {
	if v, ok := p.KV[k]; ok {
		if x, ok := v.(bool); ok {
			return x, true
		}
	}
	return false, false
}

// Fetch an int32 by key.
func (p *Properties) FetchInt32(k string) (int32, bool) {
	if v, ok := p.KV[k]; ok {
		if x, ok := v.(int32); ok {
			return x, true
		}
	}
	return 0, false
}

// Fetch a uint32 by key.
func (p *Properties) FetchUint32(k string) (uint32, bool) {
	if v, ok := p.KV[k]; ok {
		if x, ok := v.(uint32); ok {
			return x, true
		}
	}
	return 0, false
}

// Fetch a uint64 by key.
func (p *Properties) FetchUint64(k string) (uint64, bool) {
	if v, ok := p.KV[k]; ok {
		if x, ok := v.(uint64); ok {
			return x, true
		}
	}
	return 0, false
}

// Fetch a float32 by key.
func (p *Properties) FetchFloat32(k string) (float32, bool) {
	if v, ok := p.KV[k]; ok {
		if x, ok := v.(float32); ok {
			return x, true
		}
	}
	return 0.0, false
}

// Fetch a string by key.
func (p *Properties) FetchString(k string) (string, bool) {
	if v, ok := p.KV[k]; ok {
		if x, ok := v.(string); ok {
			return x, true
		}
	}
	return "", false
}

// Reads properties into p using the given reader and serializer.
func (p *Properties) readProperties(r *reader, ser *dt) {
	// Create fieldpath
	fieldPath := newFieldpath(ser)

	// Get a list of the included fields
	fieldPath.walk(r)

	// iterate all the fields and set their corresponding values
	for _, f := range fieldPath.fields {
		if v(6) {
			_debugf("decoding pos=%d name=%s type=%s encoder=%s", r.pos, f.Name, f.Field.Type, f.Field.Encoder)
		}

		if f.Field.Serializer.DecodeContainer != nil {
			p.KV[f.Name] = f.Field.Serializer.DecodeContainer(r, f.Field)
		} else if f.Field.Serializer.Decode == nil {
			p.KV[f.Name] = r.readVarUint32()
			continue
		} else {
			p.KV[f.Name] = f.Field.Serializer.Decode(r, f.Field)
		}

		if v(6) {
			_debugf("decoding pos=%d name=%s type=%s value=%v", r.pos, f.Name, f.Field.Type, p.KV[f.Name])
		}
	}
}
