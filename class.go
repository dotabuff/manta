package manta

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dotabuff/manta/dota"
)

var gameBuildRegexp = regexp.MustCompile(`/dota_v(\d+)/`)

type class struct {
	classId    int32
	name       string
	serializer *serializer
}

func (c *class) getNameForFieldPath(fp *fieldPath) string {
	return strings.Join(c.serializer.getNameForFieldPath(fp, 0), ".")
}

func (c *class) getTypeForFieldPath(fp *fieldPath) *fieldType {
	return c.serializer.getTypeForFieldPath(fp, 0)
}

func (c *class) getDecoderForFieldPath(fp *fieldPath) fieldDecoder {
	return c.serializer.getDecoderForFieldPath(fp, 0)
}

func (c *class) getFieldPathForName(fp *fieldPath, name string) bool {
	return c.serializer.getFieldPathForName(fp, name)
}

func (c *class) getFieldPaths(fp *fieldPath, state *fieldState) []*fieldPath {
	return c.serializer.getFieldPaths(fp, state)
}

// Internal callback for OnCSVCMsg_ServerInfo.
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

// Internal callback for OnCDemoClassInfo.
func (p *Parser) onCDemoClassInfo(m *dota.CDemoClassInfo) error {
	for _, c := range m.GetClasses() {
		classId := c.GetClassId()
		networkName := c.GetNetworkName()

		class := &class{
			classId:    classId,
			name:       networkName,
			serializer: p.serializers[networkName],
		}
		p.classesById[class.classId] = class
		p.classesByName[class.name] = class
	}

	p.classInfo = true

	p.updateInstanceBaseline()

	return nil
}

func (p *Parser) updateInstanceBaseline() {
	// We can't update the instancebaseline until we have class info.
	if !p.classInfo {
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
		classId, err := atoi32(item.Key)
		if err != nil {
			_panicf("invalid instancebaseline key '%s': %s", item.Key, err)
		}
		p.classBaselines[classId] = item.Value
	}
}
