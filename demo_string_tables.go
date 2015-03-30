package manta

import "github.com/dotabuff/manta/dota"

func NewStringTables() *StringTables {
	return &StringTables{
		lastIndex: 0,
		tables:    map[int]int{},
	}
}

type StringTables struct {
	lastIndex int
	tables    map[int]int
}

func (p *StringTables) OnCDemoStringTables(stringTables *dota.CDemoStringTables) {
}

func (p *StringTables) OnUpdateStringTable(ust *dota.CSVCMsg_UpdateStringTable) {
}
