package vbkv

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	None = iota
	String
	Int32
	Float32
	Pointer
	WideString
	Color
	Uint64
	End
)

type Parser struct {
	buf *bytes.Buffer
}

func ParseBytes(b []byte) (kv map[string]interface{}, err error) {
	kv, err = Parse(bytes.NewBuffer(b))
	if err != nil {
		return kv, err
	}
	return kv, err
}

func Parse(buffer *bytes.Buffer) (kv map[string]interface{}, err error) {
	parser := &Parser{buf: buffer}
	return parser.Parse()
}

func (p *Parser) Parse() (kv map[string]interface{}, err error) {
	entities := map[string]interface{}{}

	defer func() {
		if r := recover(); r != nil {
			kv = entities
			err = fmt.Errorf("%v", r)
		}
	}()

	p.parseString() // "Dota 2 Saved Game"

	for p.buf.Len() > 4 && !bytes.Equal(p.buf.Bytes()[0:4], []byte{0, 0, 0, 0}) {
		k, v := p.parseKV()
		if v == nil {
			return entities, err
		}
		entities[k] = v
	}

	for {
		k, v := p.parseEntity()
		if v == nil {
			return entities, err
		}
		entities[k] = v
	}

	return entities, err
}

func (p *Parser) parseEntity() (name string, v interface{}) {
	if p.buf.Len() == 0 {
		return name, v
	}
	flag := p.parseByte()
	entity := map[string]interface{}{}

	switch flag {
	case 0x0b:
		return name, v
	case None:
		name = p.parseString()
		for {
			k, v := p.parseKV()
			if v == nil {
				break
			}
			entity[k] = v
		}
	default:
		panic(fmt.Errorf("unknown entity flag: %d", flag))
	}

	return name, entity
}

func (p *Parser) parseKV() (k string, v interface{}) {
	if p.buf.Len() == 0 {
		return k, v
	}
	flag := p.parseByte()

	switch flag {
	case 0x0b:
		// start the next entity
		return k, nil
	case None:
		object := map[string]interface{}{}
		name := p.parseString()
		for {
			k, v = p.parseKV()
			if v == nil {
				return name, object
			}
			object[k] = v
		}
		k = name
		v = object
	case String:
		k, v = p.parseString(), p.parseString()
	case Int32:
		k, v = p.parseString(), p.parseInt32()
	case Float32:
		k, v = p.parseString(), p.parseFloat32()
	case Uint64:
		k, v = p.parseString(), p.parseUint64()
	default:
		panic(fmt.Errorf("unknown type: %d", flag))
	}

	return k, v
}

func (p *Parser) parseByte() (b byte) {
	if b, err := p.buf.ReadByte(); err != nil {
		panic(err)
	} else {
		return b
	}
}

func (p *Parser) parseString() string {
	if str, err := p.buf.ReadString(0); err != nil {
		panic(err)
	} else {
		return str[:len(str)-1]
	}
}

func (p *Parser) parseInt32() (v int32) {
	if err := binary.Read(p.buf, binary.LittleEndian, &v); err != nil {
		panic(err)
	}
	return v
}

func (p *Parser) parseUint64() (v uint64) {
	if err := binary.Read(p.buf, binary.LittleEndian, &v); err != nil {
		panic(err)
	}
	return v
}

func (p *Parser) parseFloat32() (v float32) {
	if err := binary.Read(p.buf, binary.LittleEndian, &v); err != nil {
		panic(err)
	}
	return v
}
