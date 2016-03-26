package manta

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/dotabuff/manta/dota"
)

var gameBuildRegexp = regexp.MustCompile(`/dota_v(\d+)/`)

// Internal callback for CSVCMsg_ServerInfo.
func (p *Parser) onCSVCMsg_ServerInfo(m *dota.CSVCMsg_ServerInfo) error {
	// This may be needed to parse PacketEntities.
	p.classIdSize = uint32(math.Log(float64(m.GetMaxClasses()))/math.Log(2)) + 1

	// Extract the build from the game dir.
	matches := gameBuildRegexp.FindStringSubmatch(m.GetGameDir())
	if len(matches) < 2 {
		return fmt.Errorf("unable to determine game build from '%s'", m.GetGameDir())
	}
	build, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		return err
	}
	p.GameBuild = uint32(build)

	return nil
}

// Internal callback for CDemoClassInfo.
func (p *Parser) onCDemoClassInfo(m *dota.CDemoClassInfo) error {
	// Iterate through items, storing the mapping in the parser state
	for _, c := range m.GetClasses() {
		p.ClassInfo[c.GetClassId()] = c.GetNetworkName()

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

	stringTable, ok := p.stringTables.GetTableByName("instancebaseline")
	if !ok {
		if v(1) {
			_debugf("skipping updateInstanceBaseline: no instancebaseline string table")
		}
		return
	}

	// Iterate through instancebaseline table items
	for _, item := range stringTable.Items {
		p.updateInstanceBaselineItem(item)
	}
}

func (p *Parser) updateInstanceBaselineItem(item *stringTableItem) {
	// Get the class id for the string table item
	classId, err := atoi32(item.Key)
	if err != nil {
		_panicf("invalid instancebaseline key '%s': %s", item.Key, err)
	}

	// Get the class name
	className, ok := p.ClassInfo[classId]
	if !ok {
		_panicf("unable to find class info for instancebaseline key %d", classId)
	}

	// Create an entry in the map if needed
	if _, ok := p.ClassBaselines[classId]; !ok {
		p.ClassBaselines[classId] = NewProperties()
	}

	// Get the send table associated with the class.
	serializer, ok := p.serializers[className]
	if !ok {
		_panicf("unable to find send table %s for instancebaseline key %d", className, classId)
	}

	// Uncomment to dump fixtures
	//_dump_fixture("instancebaseline/1731962898_"+className+".rawbuf", item.Value)

	// Parse the properties out of the string table buffer and store
	// them as the class baseline in the Parser.
	if len(item.Value) > 0 {
		if v(1) {
			_debugf("parsing entity baseline %v", serializer[0].Name)
		}
		r := newReader(item.Value)
		p.ClassBaselines[classId].readProperties(r, serializer[0])
	}
}
