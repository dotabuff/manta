package manta

import (
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

var pointerTypes = map[string]bool{
	"PhysicsRagdollPose_t":       true,
	"CBodyComponent":             true,
	"CEntityIdentity":            true,
	"CPhysicsComponent":          true,
	"CRenderComponent":           true,
	"CDOTAGamerules":             true,
	"CDOTAGameManager":           true,
	"CDOTASpectatorGraphManager": true,
	"CPlayerLocalData":           true,
	"CPlayer_CameraServices":     true,
	"CDOTAGameRules":             true,
}

var itemCounts = map[string]int{
	"MAX_ITEM_STOCKS":             8,
	"MAX_ABILITY_DRAFT_ABILITIES": 48,
}

// Internal callback for OnCDemoSendTables.
func (p *Parser) onCDemoSendTables(m *dota.CDemoSendTables) error {
	r := newReader(m.GetData())
	buf := r.readBytes(r.readVarUint32())

	msg := &dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, msg); err != nil {
		return err
	}

	patches := []fieldPatch{}
	for _, h := range fieldPatches {
		if h.shouldApply(p.GameBuild) {
			patches = append(patches, h)
		}
	}

	fields := map[int32]*field{}
	fieldTypes := map[string]*fieldType{}

	for _, s := range msg.GetSerializers() {
		serializer := &serializer{
			name:    msg.GetSymbols()[s.GetSerializerNameSym()],
			version: s.GetSerializerVersion(),
			fields:  []*field{},
		}

		for _, i := range s.GetFieldsIndex() {
			if _, ok := fields[i]; !ok {
				// create a new field
				field := newField(msg, msg.GetFields()[i])

				// patch parent name in builds <= 990
				if p.GameBuild <= 990 {
					field.parentName = serializer.name
				}

				// find or create a field type
				if _, ok := fieldTypes[field.varType]; !ok {
					fieldTypes[field.varType] = newFieldType(field.varType)
				}
				field.fieldType = fieldTypes[field.varType]

				// find associated serializer
				if field.serializerName != "" {
					field.serializer = p.serializers[field.serializerName]
				}

				// apply any build-specific patches to the field
				for _, h := range patches {
					h.patch(field)
				}

				// determine field model
				if field.serializer != nil {
					if field.fieldType.pointer || pointerTypes[field.fieldType.baseType] {
						field.setModel(fieldModelFixedTable)
					} else {
						field.setModel(fieldModelVariableTable)
					}
				} else if field.fieldType.count > 0 && field.fieldType.baseType != "char" {
					field.setModel(fieldModelFixedArray)
				} else if field.fieldType.baseType == "CUtlVector" || field.fieldType.baseType == "CNetworkUtlVectorBase" {
					field.setModel(fieldModelVariableArray)
				} else {
					field.setModel(fieldModelSimple)
				}

				// store the field
				fields[i] = field
			}

			// add the field to the serializer
			serializer.fields = append(serializer.fields, fields[i])
		}

		// store the serializer for field reference
		p.serializers[serializer.name] = serializer

		if _, ok := p.classesByName[serializer.name]; ok {
			p.classesByName[serializer.name].serializer = serializer
		}
	}

	return nil
}
