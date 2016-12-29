package manta

import (
	"strings"
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

func (c *class) setValueForFieldPath(fp *fieldPath, pos int, data interface{}) {

}
