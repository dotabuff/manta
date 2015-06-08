package manta

import "github.com/dotabuff/manta/dota"

func NewStringTables() *StringTables {
	return &StringTables{
		tables: map[string]StringTable{},
	}
}

type StringTables struct {
	tables map[string]StringTable
}

type StringTable struct {
	items map[int]*dota.CDemoStringTablesItemsT
}

func (p *StringTables) onCDemoStringTables(stringTables *dota.CDemoStringTables) error {
	tables := map[string]StringTable{}

	for _, st := range stringTables.GetTables() {
		table := StringTable{items: map[int]*dota.CDemoStringTablesItemsT{}}
		tables[st.GetTableName()] = table
		for n, item := range st.GetItems() {
			table.items[n] = item
		}
	}

	return nil
}
