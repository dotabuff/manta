package manta

var INDEXBITS int

// Reads properties using a reader and send table.
// Note: this is a work in progress and is almost certainly completely wrong.
func readProperties(r *reader, t *sendTable) (result map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			_debugf("recovered: %s", err)
		}
		_debugf("--")
	}()

	// The result we'll return
	result = make(map[string]interface{})

	// We may need to calculate indexBits based on the max number of properties.
	// So far this has not shown any better results, but it's here as a reminder
	// that we've considered it (and continue to).
	indexBits := log2(len(t.props))

	_debugf("reading properties for sendtable %s, props %d, buffer bits %d, indexBits %d", t.name, len(t.props), r.size, indexBits)

	// Naiive implementation from Soure 1, not working.
	// Iterate through, collecting indexes of props that we have data for.
	// This seems to be wrong. Maybe its prop->value->repeat?
	index := -1
	props := []int{}
	for {
		if r.readBoolean() {
			index += 1
			_debugf("index incr to %d", index)
		} else {
			n := int(r.readVarUint32())
			if n == 0 {
				_debugf("read 0 break")
				break
			}

			index += n + 1
			_debugf("index read %d to %d", n, index)
		}
		props = append(props, index)
	}

	// Print out a list of prop indexes
	_debugf("props: %v", props)

	// Iterate through properties, reading their data and storing in the result.
	var key string
	var prop *sendProp
	for _, index = range props {
		// Verify the index is in range.
		if index > len(t.props) {
			_panicf("prop index %d out of range (max %d)", index, len(t.props))
		}

		// Get the prop
		prop = t.props[index]
		key = prop.varName

		_debugf("reading property %s (%s)", prop.varName, prop.dtName)

		// XXX TODO: We don't know how to read anything because we don't have
		// type or bit sizes on any of the sendprops.
		switch prop.dtName {
		case "float32":
			result[key] = r.readBits(32)
		default:
			result[key] = r.readVarUint32()
		}

		_debugf("%s = %v", key, result[key])
	}

	// Write out the number of remaining bits, assuming we read short.
	_debugf("remaining bits: %d", r.remBits())

	return result
}
