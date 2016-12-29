package manta

import (
	"fmt"
	"os"

	"github.com/dotabuff/manta/dota"
)

const (
	fieldModelSimple = iota
	fieldModelFixedArray
	fieldModelFixedTable
	fieldModelVariableArray
	fieldModelVariableTable
)

type fielder interface {
	getName() string
	getNameForFieldPath(*fieldPath, int) []string
	getTypeForFieldPath(*fieldPath, int) *fieldType
	getDecoderForFieldPath(*fieldPath, int) fieldDecoder
}

type field struct {
	varName           string
	varType           string
	sendNode          string
	serializerName    string
	serializerVersion int32
	encoder           string
	encodeFlags       *int32
	bitCount          *int32
	lowValue          *float32
	highValue         *float32
	fieldType         *fieldType
	serializer        *serializer
	value             interface{}
	model             int

	fixedSubTable  bool
	varSubTable    bool
	fixedArray     bool
	varArray       bool
	fixedArraySize int

	decoder      fieldDecoder
	baseDecoder  fieldDecoder
	childDecoder fieldDecoder
}

func newField(ser *dota.CSVCMsg_FlattenedSerializer, f *dota.ProtoFlattenedSerializerFieldT) *field {
	resolve := func(p *int32) string {
		if p == nil {
			return ""
		}
		return ser.GetSymbols()[*p]
	}

	x := &field{
		varName:           resolve(f.VarNameSym),
		varType:           resolve(f.VarTypeSym),
		sendNode:          resolve(f.SendNodeSym),
		serializerName:    resolve(f.FieldSerializerNameSym),
		serializerVersion: f.GetFieldSerializerVersion(),
		encoder:           resolve(f.VarEncoderSym),
		encodeFlags:       f.EncodeFlags,
		bitCount:          f.BitCount,
		lowValue:          f.LowValue,
		highValue:         f.HighValue,
		model:             fieldModelSimple,
	}

	if x.sendNode == "(root" {
		x.sendNode = ""
	}

	return x
}

func (f *field) isFixedSubtable() {
	f.fixedSubTable = true
}

func (f *field) isVarSubtable() {
	f.varSubTable = true
}

func (f *field) setModel(model int) {
	f.model = model
	switch model {
	case fieldModelFixedArray:
		f.fixedArray = true
	case fieldModelFixedTable:
		f.fixedSubTable = true
		f.baseDecoder = unsignedDecoder
	case fieldModelVariableArray:
		f.varArray = true
		f.baseDecoder = unsignedDecoder
		if f.fieldType.genericType != nil {
			f.childDecoder = findDecoderByBaseType(f.fieldType.genericType.baseType)
		}
	case fieldModelVariableTable:
		f.varSubTable = true
	}
}

func (f *field) getName() string {
	return f.varName
}

func (f *field) getFieldForFieldPath(fp *fieldPath, pos int) *field {
	if f.fixedSubTable {
		if fp.last == pos-1 {
			return f
		}
		return f.serializer.getFieldForFieldPath(fp, pos)
	}

	return f
}

func (f *field) getNameForFieldPath(fp *fieldPath, pos int) []string {
	x := []string{f.varName}

	if f.fixedSubTable {
		// _printf("getNameForFieldPath fixed last=%d pos=%d", fp.last, pos)
		if fp.last >= pos {
			x = append(x, f.serializer.getNameForFieldPath(fp, pos)...)
		}
	} else if f.varSubTable {
		if fp.last != pos-1 {
			x = append(x, fmt.Sprintf("%04d", fp.path[pos]))
			if fp.last != pos {
				x = append(x, f.serializer.getNameForFieldPath(fp, pos+1)...)
			}
		}
	} else if f.fixedArray || f.varArray {
		if fp.last == pos {
			x = append(x, fmt.Sprintf("%04d", fp.path[pos]))
		}
	}

	return x
}

func (f *field) getTypeForFieldPath(fp *fieldPath, pos int) *fieldType {
	if f.fixedSubTable {
		if fp.last == pos-1 {
			return f.fieldType
		}
		return f.serializer.getTypeForFieldPath(fp, pos)
	}
	return f.fieldType
}

func (f *field) getDecoderForFieldPath(fp *fieldPath, pos int) fieldDecoder {
	if f.fixedSubTable {
		// _printf("getDecoderForFieldPath fixed last=%d pos=%d", fp.last, pos)
		if fp.last == pos-1 {
			// _printf("getDecoderForFieldPath fixed last=%d pos=%d BOOLEAN DEFAULT", fp.last, pos)
			return booleanDecoder
		}
		// _printf("getDecoderForFieldPath fixed last=%d pos=%d PUSHING IT OUT", fp.last, pos)
		return f.serializer.getDecoderForFieldPath(fp, pos)
	}

	return f.decoder
}

func (f *field) String() string {
	x := f.varName + " = " + f.fieldType.String()
	if f.serializer != nil {
		x += "(" + f.serializer.id() + ")"
	}
	return x
}

func readFields(r *reader, s *serializer) []interface{} {
	fps := readFieldPaths(r)

	values := make([]interface{}, len(fps))
	for i, fp := range fps {
		name := s.getNameForFieldPath(fp, 0)
		typ := s.getTypeForFieldPath(fp, 0)
		decoder := s.getDecoderForFieldPath(fp, 0)

		if waldnew {
			_printf("NEW reading ser=%s path=%s pos=%d name=%s type=%s decoder=%s", s.name, fp.String(), r.pos, name, typ, nameOf(decoder))
		}

		value := decoder(r)
		values[i] = value

		if waldnew {
			_printf(" => %v", value)
		}
	}

	return values
}

var waldold bool
var waldnew bool

func init() {
	if os.Getenv("WALDOLD") != "" {
		waldold = true
	}
	if os.Getenv("WALDNEW") != "" {
		waldnew = true
	}
}
