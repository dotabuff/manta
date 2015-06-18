package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

/*
type CSVCMsg_GameEventKeyT struct {
  Type             *int32   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
  ValString        *string  `protobuf:"bytes,2,opt,name=val_string" json:"val_string,omitempty"`
  ValFloat         *float32 `protobuf:"fixed32,3,opt,name=val_float" json:"val_float,omitempty"`
  ValLong          *int32   `protobuf:"varint,4,opt,name=val_long" json:"val_long,omitempty"`
  ValShort         *int32   `protobuf:"varint,5,opt,name=val_short" json:"val_short,omitempty"`
  ValByte          *int32   `protobuf:"varint,6,opt,name=val_byte" json:"val_byte,omitempty"`
  ValBool          *bool    `protobuf:"varint,7,opt,name=val_bool" json:"val_bool,omitempty"`
  ValUint64        *uint64  `protobuf:"varint,8,opt,name=val_uint64" json:"val_uint64,omitempty"`
  XXX_unrecognized []byte   `json:"-"`
}
*/

func geTypeName(t int32) (string, string) {
	switch t {
	case 1:
		return "string", "GetValString"
	case 2:
		return "float32", "GetValFloat"
	case 3:
		return "int32", "GetValLong"
	case 4:
		return "int32", "GetValShort"
	case 5:
		return "int32", "GetValByte"
	case 6:
		return "bool", "GetValBool"
	case 7:
		return "uint64", "GetValUint64"
	}

	panic(fmt.Sprintf("unknown ge type %d", t))
	return "", ""
}

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	m := &dota.CSVCMsg_GameEventList{}
	if err := proto.Unmarshal(buf, m); err != nil {
		panic(err)
	}

	constOut := "const (\n"
	structOut := ""
	geStructOut := "type GameEvents struct {\n"
	registerOut := ""
	handlerOut := `
    func (ge *GameEvents) onCMsgSource1LegacyGameEvent(m *dota.CMsgSource1LegacyGameEvent) error {
      switch m.GetEventid() {
      `

	constmap := make(map[int]string)
	constidx := make([]int, 0)

	for _, d := range m.GetDescriptors() {
		eventId := int(d.GetEventid())
		eventName := d.GetName()

		// Internal callback, ex. onSomeEvent
		cbInt := camelCase("on_"+eventName, false)

		// External callback, ex. OnSomeEvent
		cbExt := camelCase("on_"+eventName, true)

		// Const type, ex. EGameEvent_SomeEvent
		constSig := fmt.Sprintf("EGameEvent_%s", camelCase(eventName, true))

		// Type signature, ex. GameEventSomeEvent
		typeSig := camelCase("game_event_"+eventName, true)

		// Handler function signature, ex. func (*GameEventSomeEvent) error
		fnSig := fmt.Sprintf("func (*%s) error", typeSig)

		constOut += fmt.Sprintf("\t%s = %d\n", constSig, eventId)

		structOut += fmt.Sprintf("type %s struct {\n", typeSig)

		handlerOut += fmt.Sprintf(`
      case %d: // %s
        if cbs := ge.%s; cbs != nil {
          msg := &%s{}
          `, eventId, constSig, cbInt, typeSig)

		for i, k := range d.GetKeys() {
			fieldName := camelCase(k.GetName(), true)
			keyType, getterName := geTypeName(k.GetType())
			structOut += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, keyType, k.GetName())
			handlerOut += fmt.Sprintf("\tmsg.%s = m.GetKeys()[%d].%s()\n", fieldName, i, getterName)
		}
		structOut += "}\n\n"

		geStructOut += fmt.Sprintf("\t%s []%s\n", cbInt, fnSig)

		registerOut += fmt.Sprintf(`
      func (ge *GameEvents) %s(fn %s) {
        if ge.%s == nil {
          ge.%s = make([]%s, 0)
        }
        ge.%s = append(ge.%s, fn)
      }`+"\n\n", cbExt, fnSig, cbInt, cbInt, fnSig, cbInt, cbInt)

		handlerOut += `
          for _, fn := range cbs {
            if err := fn(msg); err != nil {
              return err
            }
          }
        }
        return nil` + "\n"

		constmap[eventId] = eventName
		constidx = append(constidx, eventId)
	}

	constOut += ")\n"
	geStructOut += "}\n"
	handlerOut += `
    }

    _panicf("unknown message %d", m.GetEventid())
    return nil
  }`

	out := fmt.Sprintf("//go:generate go run gen/game_event.go %s %s\n\n", os.Args[1], os.Args[2])
	out += "package manta\n\n"
	out += "import (\n\t\"github.com/dotabuff/manta/dota\"\n)\n"
	out += constOut
	out += structOut
	out += geStructOut
	out += registerOut
	out += handlerOut

	outbuf, err := format.Source([]byte(out))
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}

	ioutil.WriteFile(os.Args[2], outbuf, 0644)
}

func camelCase(src string, export bool) string {
	re := regexp.MustCompile("[0-9A-Za-z]+")
	byteSrc := []byte(src)
	chunks := re.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 || export {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}
