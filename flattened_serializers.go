package manta

import (
	"encoding/json"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

// Field is always filled, table only for sub-tables
type dt_property struct {
	Field *dt_field
	Table *dt
}

// A datatable field
type dt_field struct {
	Name  string
	Type  string
	Index int32

	Flags     *int32
	BitCount  *int32
	LowValue  *float32
	HighValue *float32

	Version    *int32
	Serializer *PropertySerializer `json:"-"`
}

// A single datatable
type dt struct {
	Name       string
	Flags      *int32
	Version    int32
	Properties []*dt_property
}

// The flattened serializers object
type flattened_serializers struct {
	Serializers map[string]map[int32]*dt // serializer name -> [versions]
	proto       *dota.CSVCMsg_FlattenedSerializer
	pst         *PropertySerializerTable
}

// Dumps a flattened table as json
func (sers *flattened_serializers) dump_json(name string) string {
	// Can't marshal map[int32]x
	type jContainer struct {
		Version int32
		Data    *dt
	}

	j := make([]jContainer, 0)
	for i, o := range sers.Serializers[name] {
		j = append(j, jContainer{i, o})
	}

	str, _ := json.MarshalIndent(j, "", "  ") // two space ident
	return string(str)
}

// Fills properties for a data table
func (sers *flattened_serializers) recurse_table(cur *dota.ProtoFlattenedSerializerT) *dt {
	// Basic table structure
	table := &dt{
		Name:       sers.proto.GetSymbols()[cur.GetSerializerNameSym()],
		Version:    cur.GetSerializerVersion(),
		Properties: make([]*dt_property, 0),
	}

	props := sers.proto.GetFields()

	// Append all the properties
	for _, idx := range cur.GetFieldsIndex() {
		pField := props[idx]
		prop := &dt_property{nil, nil}

		// Field can always be set
		prop.Field = &dt_field{
			Name:  sers.proto.GetSymbols()[pField.GetVarNameSym()],
			Index: -1,

			Flags:     pField.EncodeFlags,
			BitCount:  pField.BitCount,
			LowValue:  pField.LowValue,
			HighValue: pField.HighValue,

			Type:       (sers.proto.GetSymbols()[pField.GetVarTypeSym()]),
			Version:    pField.FieldSerializerVersion,
			Serializer: sers.pst.GetPropertySerializerByName(sers.proto.GetSymbols()[pField.GetVarTypeSym()]),
		}

		// Optional: Attach the serializer version for the property if applicable
		if pField.FieldSerializerNameSym != nil {
			pFieldName := sers.proto.GetSymbols()[pField.GetFieldSerializerNameSym()]
			pFieldVersion := pField.GetFieldSerializerVersion()
			pSerializer := sers.Serializers[pFieldName][pFieldVersion]

			if pSerializer == nil {
				_panicf("Error: Serializer version %d for %s hasn't been added yet.", pFieldVersion, pFieldName)
			}

			prop.Table = pSerializer
		}

		// Optional: Adjust array fields
		if prop.Field.Serializer.IsArray {
			// Add our own temp table for the array
			tmpDt := &dt{
				Name:       prop.Field.Name,
				Flags:      nil,
				Version:    0,
				Properties: make([]*dt_property, 0),
			}

			// Add each array field to the table
			for i := uint32(0); i < prop.Field.Serializer.Length; i++ {
				tmpDt.Properties = append(tmpDt.Properties, &dt_property{
					Field: &dt_field{
						Name:       _sprintf("%04d", i),
						Type:       prop.Field.Serializer.Name,
						Index:      int32(i),
						Flags:      prop.Field.Flags,
						BitCount:   prop.Field.BitCount,
						LowValue:   prop.Field.LowValue,
						HighValue:  prop.Field.HighValue,
						Version:    prop.Field.Version,
						Serializer: prop.Field.Serializer.ArraySerializer,
					},
					Table: prop.Table, // This carries on the actual table instead of overriding it
				})

				// Copy parent prop to rename it's name according to the array index
				if prop.Table != nil {
					nTable := *prop.Table
					nTable.Name = _sprintf("%04d", i)
					tmpDt.Properties[len(tmpDt.Properties)-1].Table = &nTable
				}
			}

			prop.Table = tmpDt
		}

		table.Properties = append(
			table.Properties,
			prop,
		)
	}

	return table
}

// Parses a CDemoSendTables packet
func ParseSendTables(m *dota.CDemoSendTables, pst *PropertySerializerTable) *flattened_serializers {
	// This packet just contains a single large buffer
	r := NewReader(m.GetData())

	// The buffer starts with a varint encoded length
	size := int(r.readVarUint32())
	if size != r.remBytes() {
		_panicf("expected %d additional bytes, got %d", size, r.remBytes())
	}

	// Read the rest of the buffer as a CSVCMsg_FlattenedSerializer.
	buf := r.readBytes(size)
	msg := &dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, msg); err != nil {
		_panicf("cannot decode proto: %s", err)
	}

	// Create the flattened_serializers object and fill it
	fs := &flattened_serializers{
		Serializers: make(map[string]map[int32]*dt),
		proto:       msg,
		pst:         pst,
	}

	// Iterate through all flattened serializers and fill their properties
	for _, o := range msg.GetSerializers() {
		sName := msg.GetSymbols()[o.GetSerializerNameSym()]
		sVer := o.GetSerializerVersion()

		if fs.Serializers[sName] == nil {
			fs.Serializers[sName] = make(map[int32]*dt)
		}

		fs.Serializers[sName][sVer] = fs.recurse_table(o)
	}

	return fs
}

// Internal callback for OnCDemoSendTables.
func (p *Parser) onCDemoSendTables(m *dota.CDemoSendTables) error {
	p.Serializers = ParseSendTables(m, GetDefaultPropertySerializerTable()).Serializers
	return nil
}
