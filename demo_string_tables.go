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

/*
(table_name:"downloadables" table_flags:0 )
(table_name:"genericprecache" items:<str:"" data:"\000" > table_flags:0 )
(str:"" data:"\000" )
(table_name:"decalprecache" table_flags:0 )
*/
func (p *StringTables) onCDemoStringTables(stringTables *dota.CDemoStringTables) {
	/*
		for _, st := range stringTables.GetTables() {
			st.GetTableFlags()
			// tableName := st.GetTableName()
			PP(st)
			for _, item := range st.GetItems() {
				PP(item)
			}
		}
	*/
}

func (p *StringTables) onUpdateStringTable(ust *dota.CSVCMsg_UpdateStringTable) {
}
