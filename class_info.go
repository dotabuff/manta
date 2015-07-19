package manta

import (
	"github.com/dotabuff/manta/dota"
)

// Internal callback for CSVCMsg_ServerInfo.
func (p *Parser) onCSVCMsg_ServerInfo(m *dota.CSVCMsg_ServerInfo) error {
	// This may be needed to parse PacketEntities.
	p.classIdSize = log2(int(m.GetMaxClasses()))
	return nil
}

// Internal callback for CDemoClassInfo.
func (p *Parser) onCDemoClassInfo(m *dota.CDemoClassInfo) error {
	// Iterate through items, storing the mapping in the parser state
	for _, c := range m.GetClasses() {
		p.classInfo[c.GetClassId()] = c.GetNetworkName()

		if _, ok := p.serializers[c.GetNetworkName()]; !ok {
			_panicf("unable to find table for class %d (%s)", c.GetClassId, c.GetNetworkName())
		}
	}

	// Remember that we've gotten the class info
	p.hasClassInfo = true

	// Try to update the instancebaseline
	p.updateInstanceBaseline()

	return nil
}

// Updates the state of instancebaseline
func (p *Parser) updateInstanceBaseline() {
	// We can't update the instancebaseline until we have class info.
	if !p.hasClassInfo {
		return
	}

	stringTable, ok := p.stringTables.getTableByName("instancebaseline")
	if !ok {
		_debugf("skipping updateInstanceBaseline: no instancebaseline string table")
		return
	}

	// Iterate through instancebaseline table items
	for _, item := range stringTable.items {
		// Get the class id for the string table item
		classId, err := atoi32(item.key)
		if err != nil {
			_panicf("invalid instancebaseline key '%s': %s", item.key, err)
		}

		// Get the class name
		className, ok := p.classInfo[classId]
		if !ok {
			_panicf("unable to find class info for instancebaseline key %d", classId)
		}

		// Create an entry in the map if needed
		if _, ok := p.classBaseline[classId]; !ok {
			p.classBaseline[classId] = make(map[string]interface{})
		}

		// Get the send table associated with the class.
		serializer, ok := p.serializers[className]
		if !ok {
			_panicf("unable to find send table %s for instancebaseline key %d", className, classId)
		}

		// Parse the properties out of the string table buffer and store
		// them as the class baseline in the Parser.
		if len(item.value) > 0 {
			if serializer[0].Name == "CIngameEvent_TI5" {
				// This one can't parse because it want's to go two levels into
				// DOTA_PlayerChallengeInfo. That one might be an array (would make sense)
				// but isn't marked as such.
				// @todo: Investigate later
				continue
			}

			_debugf("Parsing entity baseline %v", serializer[0].Name)
			p.classBaseline[classId] = readPropertiesNew(newReader(item.value), serializer[0])
		}
	}
}
