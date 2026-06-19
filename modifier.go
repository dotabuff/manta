package manta

import (
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

type ModifierTableEntryHandler func(msg *dota.CDOTAModifierBuffTableEntry) error

// OnModifierTableEntry registers a handler for when a ModifierBuffTableEntry
// is created or updated.
func (p *Parser) OnModifierTableEntry(fn ModifierTableEntryHandler) {
	p.modifierTableEntryHandlers = append(p.modifierTableEntryHandlers, fn)
}

// emitModifierTableEvents emits ModifierBuffTableEntry events
// from the given string table items.
func (p *Parser) emitModifierTableEvents(items []*stringTableItem) error {
	// Nothing to do if no consumer is listening; avoid the per-item proto
	// allocation + unmarshal entirely. This is the common case (e.g. the
	// benchmark and any parse that doesn't register OnModifierTableEntry).
	if len(p.modifierTableEntryHandlers) == 0 {
		return nil
	}

	for _, item := range items {
		// Skip deleted/empty entries (clarity does the same). An empty value
		// would otherwise unmarshal into an all-zero message and be emitted as a
		// spurious modifier event. A real modifier is never zero-length on wire.
		if len(item.Value) == 0 {
			continue
		}

		msg := &dota.CDOTAModifierBuffTableEntry{}
		if err := proto.NewBuffer(item.Value).Unmarshal(msg); err != nil {
			_debugf("unable to unmarshal ModifierBuffTableEntry: %s", err)
			continue
		}

		for _, fn := range p.modifierTableEntryHandlers {
			if err := fn(msg); err != nil {
				return err
			}
		}
	}

	return nil
}
