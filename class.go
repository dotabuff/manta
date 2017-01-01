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
