package manta

import (
	"strings"

	"github.com/dotabuff/manta/dota"
)

type class struct {
	classId    int32
	name       string
	serializer *serializer
}

func (c *class) getNameForFieldPath(fp *fieldPath) string {
	// fmt.Println("c getNameForFieldPath", c.serializer.name, fp.String())
	ss := c.serializer.getNameForFieldPath(fp, 0)
	return strings.Join(ss, ".")
}

func (c *class) getTypeForFieldPath(fp *fieldPath) *fieldType {
	return c.serializer.getTypeForFieldPath(fp, 0)
}

func (c *class) getDecoderForFieldPath(fp *fieldPath) fieldDecoder {
	return c.serializer.getDecoderForFieldPath(fp, 0)
}

func (c *class) getValueForFieldPath(fp *fieldPath, state *fieldState) interface{} {
	return c.serializer.getValueForFieldPath(fp, 0, state)
}

func (c *class) setValueForFieldPath(fp *fieldPath, state *fieldState, v interface{}) {
	c.serializer.setValueForFieldPath(fp, 0, state, v)
}

func (p *Parser) onCDemoClassInfoNew(m *dota.CDemoClassInfo) error {
	for _, c := range m.GetClasses() {
		classId := c.GetClassId()
		networkName := c.GetNetworkName()

		class := &class{
			classId:    classId,
			name:       networkName,
			serializer: p.newSerializers[networkName],
		}
		p.newClassesById[class.classId] = class
		p.newClassesByName[class.name] = class
	}

	p.hasClassInfo = true

	p.updateInstanceBaseline()

	return nil
}
