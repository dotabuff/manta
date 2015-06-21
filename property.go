package manta

// Reads properties using a reader and send table.
// Note: this is a work in progress and is almost certainly completely wrong.
func readProperties(r *reader, t *sendTable) (result map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			_debugf("recovered: %s", err)
		}
		_debugf("--")
	}()

	// The result we'll return.
	result = make(map[string]interface{})

	// We may need to calculate indexBits based on the max number of properties.
	// So far this has not shown any better results, but it's here as a reminder
	// that we've considered it (and continue to).
	indexBits := log2(len(t.props))

	// Just debugging.
	_debugf("table %s has %d props, %d bits, indexBits=%d",
		t.name, len(t.props), r.size, indexBits)

	// If we're in debugging mode, this simply dumps analysis of the entire buffer.
	r.dumpRemaining()

	// We create an ordered map so that we can inspect the key->value as they
	// are read from the buffer. This is for debugging (it's easier to inspect
	// an ordered map than a non-ordered map) and will likely be removed later.
	om := newOmap()

	// We currently have a reasonable handle on what's near the end of a buffer,
	// but the beginning is quite the mystery. This provides a facility to skip
	// the first N bits of a buffer before reading what we think we know, in order
	// to document the relationships between structures. We'll need to figure out
	// the magic here, but for now this lets us further our progress and learn
	// more about the structure.
	seekBits := 0
	switch t.name {
	case "CDOTATeam": // 15 fields, 15 entries. 1 unknown/array type
		seekBits = 9
	case "CRagdollManager": // 1 field, 1 entry (uint8)
		seekBits = 2
	case "CDOTAFogOfWarWasVisible": // 1 field, 1024 entries (uint64[1024])
		seekBits = 16
	case "CDOTA_DataDire", "CDOTA_DataRadiant", "CDOTA_DataCustomTeam":
		// 5 fields, 307 entries
		seekBits = 384 - 307 // 384 total header - 307 field bits.
	}
	if seekBits > 0 {
		_debugf("seeking %d bits for table %s", seekBits, t.name)
		r.seekBits(seekBits)
	}

	// It appears that the header of a buffer grows with the number of elements
	// in it. For now, let's just throw away 1 bit for each field.
	throwBits := make([]uint, 0)
	for _, prop := range t.props {
		_, propCount, err := prop.typeInfo()
		if err != nil {
			_panicf("unable to decode type info for %s: %s", prop.dtName, err)
		}
		for i := 0; i < propCount; i++ {
			throwBits = append(throwBits, r.readBits(1))
		}
	}
	_debugf("read %d field index bits", len(throwBits))

	// Once we're past the header, the data seems to be serialized contiguously.
	// This loop iterates over properties we expect to be present, reading them
	// as needed. This will likely need to change to use the information found
	// in the header, as we shouldn't expect all PacketEntity updates to include
	// all fields. It seems to work OK for instancebaseline, though.
	var k string
	var v interface{}

	// Iterate through the props in the sendtable.
	for _, prop := range t.props {
		// Extract type information from the sendprop.
		propType, propCount, err := prop.typeInfo()
		if err != nil {
			_panicf("unable to decode type info for %s: %s", prop.dtName, err)
		}

		// While debugging, print the next 8 bits ahead of us.
		r.dumpBits(8)

		// Iterate through the fields in the prop. Many props have multiple
		// entries. For example, a uint32[8] is read 8 times.
		for i := 0; i < propCount; i++ {
			// Determine the key based on the prop count
			if propCount > 1 {
				// Multiple entries have the key %NAME.%N, where N is a zero based index.
				k = _sprintf("%s.%d", prop.varName, i)
			} else {
				// Single properties have the key %NAME.
				k = prop.varName
			}

			// Just debugging help.
			_debugf("reading %s from position %d/%d", prop.Describe(), r.pos, r.size)
			pos := r.pos

			// Read the property off based on the type. This is hackey, and while I
			// hope we can get better property information, it appears that there
			// may simply be an understanding of types mapped right in the game
			// engine. That would mean that we need to recognize types by name (or id)
			// which is currently what we're doing. Let's hope we can do better here.
			switch propType {
			case "float32":
				// This will be tricky as floats can be encoded in oh-so-many ways.

				// We haven't yet determined how to read a float with a bitcount.
				if prop.bitCount != nil && *prop.bitCount != 32 {
					_panicf("unhandled bitcount: %s", prop.Describe())
				}

				// This just reads a fixed length IEEE 754 float32. It might be 100%
				// wrong, and it will at least be wrong in cases where we have
				// flags, lowVal, highVal or bitcount.
				v = r.readFloat32()

			case "int32":
				// Signed integers appear to be varints.
				v = r.readVarInt32()

			case "int8":
				// So far there's *some* evidence to suggest that this is read as
				// a varint. That doesn't make much sense though, so.... be skeptical.
				v = int8(r.readVarInt32())

			case "uint32":
				// uint32's appear to be varints.
				v = r.readVarUint32()

			case "uint64":
				// uint64's appear to be varints.
				v = r.readVarUint64()

			case "uint8":
				// uint8 appears to be read as a byte.
				v = uint8(r.readBits(8))

			case "uint16":
				// There's *some* evidence that uint16 is read as fixed-length.
				v = uint16(r.readBits(16))

			case "char":
				// A char[N] type appears to be a null terminated string, so
				// this will usually be reasonable.
				v = r.readString()

			case "CUtlSymbolLarge":
				// This appears to be a C++ type that provides some optimization,
				// but so far simply gets serialized as a string with N entries.
				// Example: CUtlSymbolLarge[6] would be 6 strings. It may or may not
				// have an outer element.
				v = r.readString()

			case "bool":
				// Seems reasonable so far.
				v = r.readBoolean()

			case "CUtlVector< CHandle< CBasePlayer > >":
				// XXX TODO: this is just wrong. This is some FML stuff.
				v = r.readBits(1)

			case "Vector":
				// So far we've seen XYZ types represented as Vector, so we're simply
				// reading 3 IEEE 754 float32's in here. It's probably wrong, and may
				// be quite complex. See float32 above for more details.
				v = []float32{r.readFloat32(), r.readFloat32(), r.readFloat32()}

			default:
				// Read unknown types as a varint, which seems to be the most popular
				// way to read most entries.
				_debugf("WARN: reading %s (%s) as varint32", k, propType)
				v = r.readVarInt32()
			}

			// Debugging
			_debugf("read %s = %v in %d bits", k, v, r.pos-pos)

			// Add the entry to the ordered map.
			om.add(k, v)

			// Set the result to the omap value, just in case we panic early.
			result = om.toMap()
		}
	}

	// Dump how many bits are left and print out the omap items.
	_debugf("%d bits left", r.remBits())
	om.print()

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
