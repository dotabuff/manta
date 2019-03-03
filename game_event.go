package manta

import (
	"bytes"
	"fmt"

	"github.com/dotabuff/manta/dota"
)

const (
	gameEventTypeString = iota + 1
	gameEventTypeFloat
	gameEventTypeLong
	gameEventTypeShort
	gameEventTypeByte
	gameEventTypeBool
	gameEventTypeUint64
)

var gameEventTypeNames = map[int32]string{
	gameEventTypeString: "string",
	gameEventTypeFloat:  "float",
	gameEventTypeLong:   "long",
	gameEventTypeShort:  "short",
	gameEventTypeByte:   "byte",
	gameEventTypeBool:   "bool",
	gameEventTypeUint64: "uint64",
}

// Represents a game event. Includes a type and the actual message.
type GameEvent struct {
	t *gameEventType
	m *dota.CMsgSource1LegacyGameEvent
}

func (ge *GameEvent) TypeName() string {
	return dota.DOTA_COMBATLOG_TYPES_name[ge.m.GetKeys()[0].GetValByte()]
}

func (ge *GameEvent) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES(ge.m.GetKeys()[0].GetValByte())
}

func (ge *GameEvent) String() string {
	keys := ge.m.GetKeys()
	name := dota.DOTA_COMBATLOG_TYPES_name[keys[0].GetValByte()]
	buf := bytes.NewBufferString("\n  " + name + "\n")

	for name, field := range ge.t.fields {
		key := keys[field.i]
		switch key.GetType() {
		case gameEventTypeString:
			fmt.Fprintf(buf, "    %s: %s\n", name, key.GetValString())
		case gameEventTypeFloat:
			fmt.Fprintf(buf, "    %s: %f\n", name, key.GetValFloat())
		case gameEventTypeLong:
			fmt.Fprintf(buf, "    %s: %d\n", name, key.GetValLong())
		case gameEventTypeShort:
			fmt.Fprintf(buf, "    %s: %d\n", name, key.GetValShort())
		case gameEventTypeByte:
			fmt.Fprintf(buf, "    %s: %d\n", name, key.GetValByte())
		case gameEventTypeBool:
			fmt.Fprintf(buf, "    %s: %t\n", name, key.GetValBool())
		case gameEventTypeUint64:
			fmt.Fprintf(buf, "    %s: %d\n", name, key.GetValUint64())
		default:
			_panicf("Unknown type: %s - %d", name, field.i)
		}
	}

	return buf.String()
}

// Gets the string value of a named field.
func (e *GameEvent) GetString(name string) (string, error) {
	// Get the key from the message
	k, err := e.getEventKey(name)
	if err != nil {
		return "", err
	}

	// Make sure it's a string.
	if k.GetType() != gameEventTypeString {
		return "", _errorf("field %s: expected string, got %s", name, gameEventTypeNames[k.GetType()])
	}

	return k.GetValString(), nil
}

// Gets the float value of a named field.
func (e *GameEvent) GetFloat32(name string) (float32, error) {
	// Get the key from the message
	k, err := e.getEventKey(name)
	if err != nil {
		return 0.0, err
	}

	// Make sure it's a bool.
	if k.GetType() != gameEventTypeFloat {
		return 0.0, _errorf("field %s: expected float, got %s", name, gameEventTypeNames[k.GetType()])
	}

	return k.GetValFloat(), nil
}

// Gets the integer value of a named field.
func (e *GameEvent) GetInt32(name string) (int32, error) {
	// Get the key from the message
	k, err := e.getEventKey(name)
	if err != nil {
		return 0, err
	}

	// Return based on the type.
	switch k.GetType() {
	case gameEventTypeLong:
		return k.GetValLong(), nil
	case gameEventTypeShort:
		return k.GetValShort(), nil
	case gameEventTypeByte:
		return k.GetValByte(), nil
	}

	return 0, _errorf("field %s: expected int, got %s", name, gameEventTypeNames[k.GetType()])
}

// Gets the bool value of a named field.
func (e *GameEvent) GetBool(name string) (bool, error) {
	// Get the key from the message
	k, err := e.getEventKey(name)
	if err != nil {
		return false, err
	}

	// Make sure it's a bool.
	if k.GetType() != gameEventTypeBool {
		return false, _errorf("field %s: expected bool, got %s", name, gameEventTypeNames[k.GetType()])
	}

	return k.GetValBool(), nil
}

// Gets the uint64 value of a named field.
func (e *GameEvent) GetUint64(name string) (uint64, error) {
	// Get the key from the message
	k, err := e.getEventKey(name)
	if err != nil {
		return 0, err
	}

	// Make sure it's a uint64.
	if k.GetType() != gameEventTypeUint64 {
		return 0, _errorf("field %s: expected uint64, got %s", name, gameEventTypeNames[k.GetType()])
	}

	return k.GetValUint64(), nil
}

// Finds the key in the game event which corresponds to a given name.
func (e *GameEvent) getEventKey(name string) (*dota.CMsgSource1LegacyGameEventKeyT, error) {
	f, ok := e.t.fields[name]
	if !ok {
		return nil, _errorf("field %s: missing", name)
	}

	if f.i > len(e.m.GetKeys()) {
		return nil, _errorf("field %s: %d out of range", name, f.i)
	}

	return e.m.GetKeys()[f.i], nil
}

// GameEventHandler is a function that can receive a game event
type GameEventHandler func(*GameEvent) error

// The type of a game event.
// Has an identifier, name, and ordered fields.
type gameEventType struct {
	eventId int32
	name    string
	fields  map[string]*gameEventField
}

// The type of a field in a game event.
// Has an index, name and type.
type gameEventField struct {
	i int
	n string
	t int32
}

// Internal handler for callback OnCMsgSource1LegacyGameEventList.
// Registers game event names and types with the parser for O(1) lookup later.
func (p *Parser) onCMsgSource1LegacyGameEventList(m *dota.CMsgSource1LegacyGameEventList) error {
	for _, d := range m.GetDescriptors() {
		t := &gameEventType{
			eventId: d.GetEventid(),
			name:    d.GetName(),
			fields:  make(map[string]*gameEventField),
		}
		for i, k := range d.GetKeys() {
			t.fields[k.GetName()] = &gameEventField{
				i: int(i),
				n: k.GetName(),
				t: k.GetType(),
			}
		}
		p.gameEventNames[d.GetEventid()] = d.GetName()
		p.gameEventTypes[d.GetName()] = t
	}

	return nil
}

// Internal handler for callback OnCMsgSource1LegacyGameEvent.
// Looks up the name and type of an event and offers it to registered handlers.
func (p *Parser) onCMsgSource1LegacyGameEvent(m *dota.CMsgSource1LegacyGameEvent) error {
	// Look up the handler name by event id.
	name, ok := p.gameEventNames[m.GetEventid()]
	if !ok {
		return _errorf("unknown event id: %d", m.GetEventid())
	}

	// Get the handlers for the event name. Return early if none.
	handlers := p.gameEventHandlers[name]
	if handlers == nil {
		return nil
	}

	// Get the type for the event.
	t, ok := p.gameEventTypes[name]
	if !ok {
		return _errorf("unknown event type: %s", name)
	}

	// Create a GameEvent, offer to all handlers.
	e := &GameEvent{t: t, m: m}
	for _, h := range handlers {
		if err := h(e); err != nil {
			return err
		}
	}

	return nil
}

// OnGameEvent registers an GameEventHandler that will be called when a
// named GameEvent occurs.
func (p *Parser) OnGameEvent(name string, fn GameEventHandler) {
	p.gameEventHandlers[name] = append(p.gameEventHandlers[name], fn)
}
