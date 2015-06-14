package manta

import (
	"github.com/dotabuff/manta/dota"
)

// Internal parser for callback dota.CDemoClassInfo.
func (p *Parser) onCDemoClassInfo(classInfo *dota.CDemoClassInfo) error {
	// Iterate through items, storing the mapping in the parser state
	for _, class := range classInfo.GetClasses() {
		p.classInfo[class.GetClassId()] = class.GetNetworkName()

		if _, ok := p.sendTables.getTableByName(class.GetNetworkName()); !ok {
			_panicf("unable to find table for class %d (%s)", class.GetClassId, class.GetNetworkName())
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
	if !p.hasClassInfo {
		_debugf("skipping updateInstanceBaseline: no class info")
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
		sendTable, ok := p.sendTables.getTableByName(className)
		if !ok {
			_panicf("unable to find send table %s for instancebaseline key %d", className, classId)
		}

		// Parse the properties out of the string table buffer and store
		// them as the class baseline in the Parser.
		if len(item.value) > 0 {
			p.classBaseline[classId] = parsePropsValues(item.value, sendTable)
		}
	}
}

// Extracts key-value pairs from a buffer using a given send table.
func parsePropsValues(buf []byte, table *sendTable) map[string]interface{} {
	result := make(map[string]interface{})

	// XXX TODO: read the buffer. We need to link type information to the send tables
	// before we can approach this with any reasonable hope of success.

	return result
}
