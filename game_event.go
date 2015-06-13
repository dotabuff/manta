package manta

import (
	"fmt"
	"go/format"
	"reflect"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/manta/dota"
)

func (p *Parser) onGameEvent(m *dota.CSVCMsg_GameEvent) error {
	_dump("m", m)
	c := &combatLogParser{stringTables: p.stringTables}
	_dump("parsed m", c.parse(m))
	return nil
}

type CombatLogEntry interface {
	Type() dota.DOTA_COMBATLOG_TYPES
	Timestamp() float32
}

type combatLogParser struct {
	stringTables *stringTables
}

/*
The default is mostly:

1 SourceName
2 TargetName
3 AttackerName
4 InflictorName
5 AttackerIsillusion
6 TargetIsIllusion
7 Value
8 Health
9 Timestamp
10 TargetSourceName
11 TimestampRaw
12 AttackerIsHero
13 TargetIsHero
*/
func (c combatLogParser) parse(obj *dota.CSVCMsg_GameEvent) CombatLogEntry {
	keys := obj.GetKeys()

	var v CombatLogEntry
	t := dota.DOTA_COMBATLOG_TYPES(keys[0].GetValByte())
	switch t {
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_ABILITY:
		v = &CombatLogAbility{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_ABILITY_TRIGGER:
		v = &CombatLogAbilityTrigger{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DAMAGE:
		v = &CombatLogDamage{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DEATH:
		v = &CombatLogDeath{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_GAME_STATE:
		v = &CombatLogGameState{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_GOLD:
		v = &CombatLogGold{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_HEAL:
		v = &CombatLogHeal{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_ITEM:
		v = &CombatLogItem{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_LOCATION:
		v = &CombatLogLocation{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_MODIFIER_ADD:
		v = &CombatLogModifierAdd{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_MODIFIER_REMOVE:
		v = &CombatLogModifierRemove{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_PURCHASE:
		v = &CombatLogPurchase{}
		// printCombatLogKeys(v, keys)
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_XP:
		v = &CombatLogXP{}
	case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_BUYBACK:
		v = &CombatLogBuyback{}
	default:
		_dump("unknown game event", t, keys)
		return nil
	}

	c.assign(v, keys)
	return v
}

type CombatLogBuyback struct {
	PlayerId int     `logIndex:"7"`  //  7:  9
	Time     float32 `logIndex:"9"`  //  9:  2625.6892
	TimeRaw  float32 `logIndex:"11"` // 11:  2666.3
}

func (c CombatLogBuyback) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_BUYBACK
}
func (c CombatLogBuyback) Timestamp() float32 {
	return c.Time
}

type CombatLogItem struct {
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`
	User               string  `logIndex:"3" logTable:"CombatLogNames"`
	Item               string  `logIndex:"4" logTable:"CombatLogNames"`
	AttackerIsIllusion bool    `logIndex:"5"`
	TargetIsIllusion   bool    `logIndex:"6"`
	Time               float32 `logIndex:"9"`
	UserIsHero         bool    `logIndex:"12"`
	TargetIsHero       bool    `logIndex:"13"`
}

func (c CombatLogItem) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_ITEM
}
func (c CombatLogItem) Timestamp() float32 {
	return c.Time
}

// TODO: only observed 2,3,4,7,13 so far, but the others would make sense too.
type CombatLogAbility struct {
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`
	Attacker           string  `logIndex:"3" logTable:"CombatLogNames"`
	Ability            string  `logIndex:"4" logTable:"CombatLogNames"`
	AttackerIsIllusion bool    `logIndex:"5"`
	TargetIsIllusion   bool    `logIndex:"6"`
	IsDebuff           int     `logIndex:"7"` // seen values: 0,1,2
	Time               float32 `logIndex:"9"`
	TargetSource       string  `logIndex:"10" logTable:"CombatLogNames"`
	AttackerIsHero     bool    `logIndex:"12"`
	TargetIsHero       bool    `logIndex:"13"`
}

func (c CombatLogAbility) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_ABILITY
}
func (c CombatLogAbility) Timestamp() float32 {
	return c.Time
}

type CombatLogAbilityTrigger struct {
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`  // 12
	Attacker           string  `logIndex:"3" logTable:"CombatLogNames"`  // 5
	Ability            string  `logIndex:"4" logTable:"CombatLogNames"`  // 47
	AttackerIsIllusion bool    `logIndex:"5"`                            // false
	TargetIsIllusion   bool    `logIndex:"6"`                            // false
	IsDebuff           int     `logIndex:"7"`                            // 3  (seen values: 3)
	Unknown8           int     `logIndex:"8"`                            // 0
	Time               float32 `logIndex:"9"`                            // 1519.1506
	TargetSource       string  `logIndex:"10" logTable:"CombatLogNames"` // 0
	TimeRaw            float32 `logIndex:"11"`                           // 1638.1001
	AttackerIsHero     bool    `logIndex:"12"`                           // true
	TargetIsHero       bool    `logIndex:"13"`                           // true
	Unknown14          bool    `logIndex:"14"`                           // false
	Unknown15          bool    `logIndex:"15"`                           // false
	Unknown16          int     `logIndex:"16"`                           // 4
	Unknown17          int     `logIndex:"17"`                           // 0
	Unknown18          int     `logIndex:"18"`                           // 0
	Unknown19          int     `logIndex:"19"`                           // 0
}

func (c CombatLogAbilityTrigger) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_ABILITY_TRIGGER
}

func (c CombatLogAbilityTrigger) Timestamp() float32 {
	return c.Time
}

type CombatLogDamage struct {
	Source             string  `logIndex:"1" logTable:"CombatLogNames"`  // 3
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`  // 27
	Attacker           string  `logIndex:"3" logTable:"CombatLogNames"`  // 3
	Cause              string  `logIndex:"4" logTable:"CombatLogNames"`  // 0
	AttackerIsIllusion bool    `logIndex:"5"`                            // false
	TargetIsIllusion   bool    `logIndex:"6"`                            // false
	Value              int     `logIndex:"7"`                            // 70
	Health             int     `logIndex:"8"`                            // 429
	Time               float32 `logIndex:"9"`                            // 229.45338
	TargetSource       string  `logIndex:"10" logTable:"CombatLogNames"` // 27
	TimeRaw            float32 `logIndex:"11"`                           // 238.43335
	AttackerIsHero     bool    `logIndex:"12"`                           // true
	TargetIsHero       bool    `logIndex:"13"`                           // false
}

func (c CombatLogDamage) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DAMAGE
}
func (c CombatLogDamage) Timestamp() float32 {
	return c.Time
}

type CombatLogLocation struct {
	Source             string  `logIndex:"1" logTable:"CombatLogNames"`
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`
	Attacker           string  `logIndex:"3" logTable:"CombatLogNames"`
	Modifier           string  `logIndex:"4" logTable:"CombatLogNames"`
	AttackerIsIllusion bool    `logIndex:"5"`
	TargetIsIllusion   bool    `logIndex:"6"`
	Value              int     `logIndex:"7"`
	Health             int     `logIndex:"8"`
	Time               float32 `logIndex:"9"`
	TargetSource       string  `logIndex:"10" logTable:"CombatLogNames"`
	AttackerIsHero     bool    `logIndex:"12"`
	TargetIsHero       bool    `logIndex:"13"`
}

func (c CombatLogLocation) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_LOCATION
}
func (c CombatLogLocation) Timestamp() float32 {
	return c.Time
}

type CombatLogHeal struct {
	Source             string  `logIndex:"1" logTable:"CombatLogNames"`
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`
	Attacker           string  `logIndex:"3" logTable:"CombatLogNames"`
	Modifier           string  `logIndex:"4" logTable:"CombatLogNames"`
	AttackerIsIllusion bool    `logIndex:"5"`
	TargetIsIllusion   bool    `logIndex:"6"`
	Value              int     `logIndex:"7"`
	Health             int     `logIndex:"8"`
	Time               float32 `logIndex:"9"`
	TargetSource       string  `logIndex:"10" logTable:"CombatLogNames"`
	AttackerIsHero     bool    `logIndex:"12"`
	TargetIsHero       bool    `logIndex:"13"`
}

func (c CombatLogHeal) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_HEAL
}
func (c CombatLogHeal) Timestamp() float32 {
	return c.Time
}

type CombatLogModifierAdd struct {
	Source             string  `logIndex:"1" logTable:"CombatLogNames"`
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`
	Attacker           string  `logIndex:"3" logTable:"CombatLogNames"`
	Modifier           string  `logIndex:"4" logTable:"CombatLogNames"`
	AttackerIsIllusion bool    `logIndex:"5"`
	TargetIsIllusion   bool    `logIndex:"6"`
	IsDebuff           bool    `logIndex:"7"`
	Health             int     `logIndex:"8"`
	Time               float32 `logIndex:"9"`
	TargetSource       string  `logIndex:"10" logTable:"CombatLogNames"`
	AttackerIsHero     bool    `logIndex:"12"`
	TargetIsHero       bool    `logIndex:"13"`
}

func (c CombatLogModifierAdd) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_MODIFIER_ADD
}

func (c CombatLogModifierAdd) Timestamp() float32 {
	return c.Time
}

type CombatLogModifierRemove struct {
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`
	Caster             string  `logIndex:"3" logTable:"CombatLogNames"`
	Modifier           string  `logIndex:"4" logTable:"CombatLogNames"`
	AttackerIsIllusion bool    `logIndex:"5"`
	TargetIsIllusion   bool    `logIndex:"6"`
	IsDebuff           bool    `logIndex:"7"`
	Health             int     `logIndex:"8"`
	Time               float32 `logIndex:"9"`
	AttackerIsHero     bool    `logIndex:"12"`
	TargetIsHero       bool    `logIndex:"13"`
}

func (c CombatLogModifierRemove) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_MODIFIER_REMOVE
}
func (c CombatLogModifierRemove) Timestamp() float32 {
	return c.Time
}

type CombatLogDeath struct {
	Source             string  `logIndex:"1" logTable:"CombatLogNames"`
	Target             string  `logIndex:"2" logTable:"CombatLogNames"`
	Attacker           string  `logIndex:"3" logTable:"CombatLogNames"`
	Cause              string  `logIndex:"4" logTable:"CombatLogNames"`
	AttackerIsIllusion bool    `logIndex:"5"`
	TargetIsIllusion   bool    `logIndex:"6"`
	Time               float32 `logIndex:"9"`
	TargetSource       string  `logIndex:"10" logTable:"CombatLogNames"`
	TimeRaw            float32 `logIndex:"11"`
	AttackerIsHero     bool    `logIndex:"12"`
	TargetIsHero       bool    `logIndex:"13"`
}

func (c CombatLogDeath) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DEATH
}
func (c CombatLogDeath) Timestamp() float32 {
	return c.Time
}

type CombatLogPurchase struct {
	Buyer   string  `logIndex:"2" logTable:"CombatLogNames"`
	Item    string  `logIndex:"7" logTable:"CombatLogNames"`
	Time    float32 `logIndex:"9"`
	TimeRaw float32 `logIndex:"11"`
}

func (c CombatLogPurchase) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_PURCHASE
}
func (c CombatLogPurchase) Timestamp() float32 {
	return c.Time
}

type CombatLogGold struct {
	Target  string  `logIndex:"2" logTable:"CombatLogNames"`
	Value   int     `logIndex:"7"`
	Time    float32 `logIndex:"9"`
	TimeRaw float32 `logIndex:"11"`
	Reason  int     `logIndex:"17"`
}

func (c CombatLogGold) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_GOLD
}
func (c CombatLogGold) Timestamp() float32 {
	return c.Time
}

type CombatLogGameState struct {
	State   int     `logIndex:"7"`  //  7: 5 (2,3,4,5,6)
	Time    float32 `logIndex:"9"`  //  9: 505.76474
	TimeRaw float32 `logIndex:"11"` // 11: 597.93335
}

func (c CombatLogGameState) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_GAME_STATE
}
func (c CombatLogGameState) Timestamp() float32 {
	return c.Time
}

type CombatLogXP struct {
	Target  string  `logIndex:"2" logTable:"CombatLogNames"`
	Value   int     `logIndex:"7"`
	Time    float32 `logIndex:"9"`
	TimeRaw float32 `logIndex:"11"`
	Reason  int     `logIndex:"18"`
}

func (c CombatLogXP) Type() dota.DOTA_COMBATLOG_TYPES {
	return dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_XP
}
func (c CombatLogXP) Timestamp() float32 {
	return c.Time
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func (c combatLogParser) assign(v CombatLogEntry, keys []*dota.CSVCMsg_GameEventKeyT) {
	rv := reflect.ValueOf(v).Elem()
	rt := rv.Type()
	fieldIndices := make([]int, rv.NumField())
	for i, _ := range fieldIndices {
		fieldTag := rt.Field(i).Tag
		logIndex := atoi(fieldTag.Get("logIndex"))

		// this allows us to be backwards compatible, but we still need to care in
		// the code that uses the log that some of the newer fields may not be set.
		if logIndex <= 0 || logIndex >= len(keys) {
			continue
		}

		field := rv.Field(i)
		key := keys[logIndex]

		switch key.GetType() {
		case 2:
			field.SetFloat(float64(key.GetValFloat()))
		case 4:
			valShort := key.GetValShort()
			if logTable := fieldTag.Get("logTable"); logTable != "" {
				table := c.stringTables.tables[logTable].items
				entry := table[valShort]
				if entry == nil {
					spew.Printf("no entry %d in %s for %v\n", valShort, logTable, v)
				} else {
					field.SetString(entry.key)
				}
			} else if field.Kind() == reflect.Bool {
				field.SetBool(valShort == 1)
			} else {
				field.SetInt(int64(valShort))
			}
		case 5:
			field.SetInt(int64(key.GetValByte()))
		case 6:
			field.SetBool(key.GetValBool())
		default:
			panic("unknown GameEventKey Type" + spew.Sdump(key) + " in " + spew.Sdump(keys))
		}
	}
}

func printCombatLogKeys(v CombatLogEntry, keys []*dota.CSVCMsg_GameEventKeyT) {
	out := []byte{}

	typeString := v.Type().String()
	parts := strings.Split(typeString, "_")
	parts = parts[2:len(parts)]
	for i, part := range parts {
		parts[i] = strings.Title(strings.ToLower(part))
	}
	typeName := strings.Join(append([]string{"CombatLog"}, parts...), "")

	out = append(out, spew.Sprintf("type %s struct {\n", typeName)...)

	var key interface{}
	var t string

	for i, k := range keys {
		switch k.GetType() {
		case 2:
			key = float64(k.GetValFloat())
			t = "float64"
		case 4:
			key = int64(k.GetValShort())
			t = "int64"
		case 5:
			key = int64(k.GetValByte())
			t = "int64"
		case 6:
			key = k.GetValBool()
			t = "bool"
		}

		name := fmt.Sprintf("Unknown%d", i)

		switch i {
		case 0:
			continue
		case 5:
			name = "AttackerIsillusion"
		case 6:
			name = "TargetIsIllusion"
		case 9:
			name = "Time"
		case 11:
			name = "TimeRaw"
		case 12:
			name = "AttackerIsHero"
		case 13:
			name = "TargetIsHero"
		case 14, 15:
			if key.(bool) {
				_dump("interesting bool key", v, keys)
				panic("found one")
			}
		}

		out = append(out, fmt.Sprintf("%s %s `logIndex:\"%d\"` // %v\n", name, t, i, key)...)
	}

	out = append(out, '}', '\n')
	out = append(out, fmt.Sprintf(`func (c %s) Type() dota.DOTA_COMBATLOG_TYPES {`, typeName)...)
	out = append(out, fmt.Sprintf(`return dota.DOTA_COMBATLOG_TYPES_%s`, typeString)...)
	out = append(out, '}', '\n')
	out = append(out, fmt.Sprintf(`func (c %s) Timestamp() float32 {`, typeName)...)
	out = append(out, `return c.Time`...)
	out = append(out, '}', '\n')

	if formatted, err := format.Source(out); err == nil {
		spew.Println(string(formatted))
	} else {
		panic(err)
	}
}
