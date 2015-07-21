package manta

import (
	"github.com/dotabuff/manta/dota"
)

const (
	gameEventTypeString = 1
	gameEventTypeFloat  = 2
	gameEventTypeLong   = 3
	gameEventTypeShort  = 4
	gameEventTypeByte   = 5
	gameEventTypeBool   = 6
	gameEventTypeUint64 = 7
)

var gameEventTypeNames = map[int32]string{
	1: "string",
	2: "float",
	3: "long",
	4: "short",
	5: "byte",
	6: "bool",
	7: "uint64",
}

// Represents a game event. Includes a type and the actual message.
type GameEvent struct {
	t *gameEventType
	m *dota.CMsgSource1LegacyGameEvent
}

// Gets the string value of a named field.
func (e *GameEvent) GetString(name string) (string, error) {
	// Get the index of the field.
	tf, ok := e.t.fields[name]
	if !ok {
		return "", _errorf("field %s: missing", name)
	}

	// Get the field.
	f := e.m.GetKeys()[tf.i]

	// Make sure it's a string.
	if f.GetType() != gameEventTypeString {
		return "", _errorf("field %s: expected string, got %s", gameEventTypeNames[f.GetType()])
	}

	return f.GetValString(), nil
}

// Gets the float value of a named field.
func (e *GameEvent) GetFloat32(name string) (float32, error) {
	// Get the index of the field.
	tf, ok := e.t.fields[name]
	if !ok {
		return 0.0, _errorf("field %s: missing", name)
	}

	// Get the field.
	f := e.m.GetKeys()[tf.i]

	// Make sure it's a bool.
	if f.GetType() != gameEventTypeFloat {
		return 0.0, _errorf("field %s: expected float, got %s", gameEventTypeNames[f.GetType()])
	}

	return f.GetValFloat(), nil
}

// Gets the integer value of a named field.
func (e *GameEvent) GetInt32(name string) (int32, error) {
	// Get the index of the field.
	tf, ok := e.t.fields[name]
	if !ok {
		return 0, _errorf("field %s: missing", name)
	}

	// Get the field.
	f := e.m.GetKeys()[tf.i]

	// Return based on the type.
	switch f.GetType() {
	case gameEventTypeLong:
		return f.GetValLong(), nil
	case gameEventTypeShort:
		return f.GetValShort(), nil
	case gameEventTypeByte:
		return f.GetValByte(), nil
	}

	return 0, _errorf("field %s: expected int, got %s", gameEventTypeNames[f.GetType()])
}

// Gets the bool value of a named field.
func (e *GameEvent) GetBool(name string) (bool, error) {
	// Get the index of the field.
	tf, ok := e.t.fields[name]
	if !ok {
		return false, _errorf("field %s: missing", name)
	}

	// Get the field.
	f := e.m.GetKeys()[tf.i]

	// Make sure it's a bool.
	if f.GetType() != gameEventTypeBool {
		return false, _errorf("field %s: expected bool, got %s", gameEventTypeNames[f.GetType()])
	}

	return f.GetValBool(), nil
}

// Gets the uint64 value of a named field.
func (e *GameEvent) GetUint64(name string) (uint64, error) {
	// Get the index of the field.
	tf, ok := e.t.fields[name]
	if !ok {
		return 0, _errorf("field %s: missing", name)
	}

	// Get the field.
	f := e.m.GetKeys()[tf.i]

	// Make sure it's a uint64.
	if f.GetType() != gameEventTypeUint64 {
		return 0, _errorf("field %s: expected uint64, got %s", gameEventTypeNames[f.GetType()])
	}

	return f.GetValUint64(), nil
}

// A function that can handle a game event.
type gameEventHandler func(*GameEvent) error

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
		_debugf("registering game event type %s: %v", d.GetName(), t)
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

// Registers a new game event handler.
func (p *Parser) OnGameEvent(name string, fn gameEventHandler) {
	p.gameEventHandlers[name] = append(p.gameEventHandlers[name], fn)
}
