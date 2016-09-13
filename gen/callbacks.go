package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

// dotaMessage provides metadata to link an enum type and value to a struct type
type dotaMessage struct {
	// typeRe provides a regular expression for the matching type
	// example: /^CDemo/ to match CDemoPacket
	typeRe *regexp.Regexp

	// enumName provides the type name for the matching enum
	// example: EDemoCommands
	enumName string

	// enumValues contain discovered values for this enum as a name -> int map
	// ex: EDemoCommands_DEM_Packet = 1
	enumValues map[string]int

	// enumToType provides a function to convert an enum name (ex.
	// EDemoCommands_DEM_SignonPacket) to a type name (ex. CDemoPacket). This can
	// typically be accomplished with string replacement but often needs
	// custom rules for non-conventionally named types
	enumToType func(s string) (string, bool)

	// enumToCallback provides a function to convert an enum name (ex.
	// EDemoCommands_DEM_SignonPacket) to a callback name (ex. CDemoSignonPacket).
	// This can typically be done using the value from enumToType, but provides a
	// mechanism to provide distinct callbacks for messages that contain the same
	// type.
	enumToCallback func(s string) (string, bool)

	// isPacket determines whether or not this message is a packet or demo type.
	isPacket bool
}

// messageTypes are constant defined message type mappings, edit as necessary.
var messageTypes = []*dotaMessage{
	&dotaMessage{
		typeRe:     regexp.MustCompile("^CDemo"),
		enumName:   "EDemoCommands",
		enumValues: map[string]int{},
		enumToType: func(s string) (string, bool) {
			switch s {
			case "EDemoCommands_DEM_SignonPacket": // uses CDemoPacket type
				return "CDemoPacket", true
			case "EDemoCommands_DEM_Error": // not a command
				return "", false
			case "EDemoCommands_DEM_Max": // not a command
				return "", false
			case "EDemoCommands_DEM_IsCompressed": // not a command
				return "", false
			}
			return strings.Replace(s, "EDemoCommands_DEM_", "CDemo", 1), true
		},
		enumToCallback: func(s string) (string, bool) {
			switch s {
			case "EDemoCommands_DEM_SignonPacket": // uses CDemoPacket type
				return "CDemoSignonPacket", true
			}
			return "", false
		},
		isPacket: false,
	},
	&dotaMessage{
		typeRe:     regexp.MustCompile("^CNETMsg_"),
		enumName:   "NET_Messages",
		enumValues: map[string]int{},
		enumToType: func(s string) (string, bool) {
			return strings.Replace(s, "NET_Messages_net_", "CNETMsg_", 1), true
		},
		enumToCallback: func(s string) (string, bool) {
			return "", false
		},
		isPacket: true,
	},
	&dotaMessage{
		typeRe:     regexp.MustCompile("^CSVCMsg_"),
		enumName:   "SVC_Messages",
		enumValues: map[string]int{},
		enumToType: func(s string) (string, bool) {
			return strings.Replace(s, "SVC_Messages_svc_", "CSVCMsg_", 1), true
		},
		enumToCallback: func(s string) (string, bool) {
			return "", false
		},
		isPacket: true,
	},
	&dotaMessage{
		typeRe:     regexp.MustCompile("^CUserMessage"),
		enumName:   "EBaseUserMessages",
		enumValues: map[string]int{},
		enumToType: func(s string) (string, bool) {
			switch s {
			case "EBaseUserMessages_UM_ParticleManager":
				return "", false
			case "EBaseUserMessages_UM_HudError":
				return "", false // CUserMsg_HudError ??
			case "EBaseUserMessages_UM_CustomGameEvent":
				return "", false // CClientMsg_CustomGameEvent ?
			case "EBaseUserMessages_UM_MAX_BASE":
				return "", false
			}
			return strings.Replace(s, "EBaseUserMessages_UM_", "CUserMessage", 1), true
		},
		enumToCallback: func(s string) (string, bool) {
			return "", false
		},
		isPacket: true,
	},
	&dotaMessage{
		typeRe:     regexp.MustCompile("^CEntityMessage"),
		enumName:   "EBaseEntityMessages",
		enumValues: map[string]int{},
		enumToType: func(s string) (string, bool) {
			return strings.Replace(s, "EBaseEntityMessages_EM_", "CEntityMessage", 1), true
		},
		enumToCallback: func(s string) (string, bool) {
			return "", false
		},
		isPacket: true,
	},
	&dotaMessage{
		typeRe:     regexp.MustCompile("^CMsg"),
		enumName:   "EBaseGameEvents",
		enumValues: map[string]int{},
		enumToType: func(s string) (string, bool) {
			return strings.Replace(s, "EBaseGameEvents_GE_", "CMsg", 1), true
		},
		enumToCallback: func(s string) (string, bool) {
			return "", false
		},
		isPacket: true,
	},
	&dotaMessage{
		typeRe:     regexp.MustCompile("^CDOTAUserMsg"),
		enumName:   "EDotaUserMessages",
		enumValues: map[string]int{},
		enumToType: func(s string) (string, bool) {
			switch s {
			case "EDotaUserMessages_DOTA_UM_AddUnitToSelection":
				return "", false
			case "EDotaUserMessages_DOTA_UM_CombatLogData":
				return "", false
			case "EDotaUserMessages_DOTA_UM_CharacterSpeakConcept":
				return "", false
			case "EDotaUserMessages_DOTA_UM_TournamentDrop":
				return "CMsgGCToClientTournamentItemDrop", true
			case "EDotaUserMessages_DOTA_UM_StatsHeroDetails":
				return "CDOTAUserMsg_StatsHeroMinuteDetails", true
			case "EDotaUserMessages_DOTA_UM_CombatLogDataHLTV":
				return "CMsgDOTACombatLogEntry", true
			case "EDotaUserMessages_DOTA_UM_MatchMetadata":
				return "CDOTAMatchMetadataFile", true
			case "EDotaUserMessages_DOTA_UM_MatchDetails":
				return "", false
			}
			return strings.Replace(s, "EDotaUserMessages_DOTA_UM_", "CDOTAUserMsg_", 1), true
		},
		enumToCallback: func(s string) (string, bool) {
			return "", false
		},
		isPacket: true,
	},
}

// messageTypeNames will be populated during type discovery, holding names of
// known struct types so that we never write invalid types to the output file
var messageTypeNames = map[string]bool{}

// findMessageByTypeName finds a message type by type name (ex. CDemoPacket)
func findMessageByTypeName(s string) (*dotaMessage, bool) {
	for _, t := range messageTypes {
		if t.typeRe.MatchString(s) {
			return t, true
		}
	}
	return nil, false
}

// findMessageByEnumName finds a message type by enum name (ex. EDemoCommands)
func findMessageByEnumName(s string) (*dotaMessage, bool) {
	for _, t := range messageTypes {
		if s == t.enumName {
			return t, true
		}
	}
	return nil, false
}

func main() {
	discoverTypes("./dota")

	buf, err := ioutil.ReadFile("gen/callbacks.tmpl")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("callbacks").Parse(string(buf))
	if err != nil {
		panic(err)
	}

	c := makeContext()

	bw := bytes.NewBuffer(nil)
	if err := tmpl.Execute(bw, c); err != nil {
		panic(err)
	}

	source, err := format.Source(bw.Bytes())
	if err != nil {
		fmt.Println("gofmt failed!", err)
		fmt.Println(string(bw.Bytes()))
		panic(err)
	}

	if err := ioutil.WriteFile("callbacks.go", source, 0644); err != nil {
		panic(err)
	}
}

// ctxType holds type information for a specific type
type ctxType struct {
	Id       int
	Callback string
	EnumName string
	TypeName string
}

// ctx holds context for the template
type ctx struct {
	DemoTypes   []ctxType
	PacketTypes []ctxType
}

// makeContext transforms messageTypes information into template context
func makeContext() ctx {
	c := ctx{
		DemoTypes:   []ctxType{},
		PacketTypes: []ctxType{},
	}

	demoIds := []int{}
	demoMap := map[int]ctxType{}
	packetIds := []int{}
	packetMap := map[int]ctxType{}

	for _, t := range messageTypes {
		for vs, n := range t.enumValues {
			// Find an enum type for the enum value
			if ts, ok := t.enumToType(vs); ok {
				// Find a struct type for the enum type
				if messageTypeNames[ts] {
					cs, ok := t.enumToCallback(vs)
					if !ok {
						cs = ts
					}
					if t.isPacket {
						packetIds = append(packetIds, n)
						packetMap[n] = ctxType{
							Id:       n,
							Callback: cs,
							EnumName: vs,
							TypeName: ts,
						}
					} else {
						demoIds = append(demoIds, n)
						demoMap[n] = ctxType{
							Id:       n,
							Callback: cs,
							EnumName: vs,
							TypeName: ts,
						}
					}
				} else {
					fmt.Printf("warning: no matching type %s for enum: type %s %s = %d\n", ts, vs, t.enumName, n)
				}
			} else {
				fmt.Printf("notice: skipped manually excluded enum: type %s %s = %d\n", vs, t.enumName, n)
			}
		}
	}

	sort.Ints(demoIds)
	sort.Ints(packetIds)

	for _, id := range demoIds {
		c.DemoTypes = append(c.DemoTypes, demoMap[id])
	}
	for _, id := range packetIds {
		c.PacketTypes = append(c.PacketTypes, packetMap[id])
	}

	return c
}

// discoverTypes walks the go files in the dota directory and populates the data
// in messageTypes and messageTypeNames.
func discoverTypes(protoPath string) {
	fileset := token.NewFileSet()
	packageMap, err := parser.ParseDir(fileset, protoPath, nil, 0)
	if err != nil {
		panic(err)
	}

	// Iterate over packages
	for _, p := range packageMap {
		// Iterate over files
		for _, f := range p.Files {
			// Iterate over declarations
			for _, d := range f.Decls {
				switch t := d.(type) {
				case *ast.GenDecl:
					// Iterate over specifications
					for _, s := range t.Specs {
						// Handle type and value specs
						switch t := s.(type) {
						case *ast.TypeSpec:
							// Type specs contain types such as CDemoPacket. Extract the type
							// name and mark it as found in the corresponding message type.
							typeName := t.Name.String()
							//if _, ok := findMessageByTypeName(typeName); ok {
							messageTypeNames[typeName] = true
							//}

						case *ast.ValueSpec:
							// Value specs contain enum values such as EDemoCommands_DEM_Packet.
							// Extract the enum name (ex. EDemoCommands), value name
							// (ex. EDemoCommands_DEM_Packet) and real value (ex. 7) and
							// store them with the matching enum type.
							valueName := nameOfValue(t)
							if enumName := typeOfValue(t); enumName != "" {
								if mt, ok := findMessageByEnumName(enumName); ok {
									mt.enumValues[valueName] = valueOfValue(t)
								}
							}
						}
					}
				}
			}
		}
	}
}

// nameOfValue extracts the name of an enum value, ex. EDemoCommands_DEM_Packet
func nameOfValue(t *ast.ValueSpec) string {
	if len(t.Names) != 1 {
		fmt.Println("unexpected names", t.Names)
		panic("unexpected names")
	}
	return t.Names[0].Name
}

// typeOfValue extracts the type of an enum value, ex. EDemoCommands
func typeOfValue(t *ast.ValueSpec) string {
	switch t := t.Type.(type) {
	case *ast.Ident:
		return t.String()
	}
	return ""
}

// valueOfValue extracts the value of an enum value, ex. 7
func valueOfValue(t *ast.ValueSpec) int {
	if len(t.Values) != 1 {
		fmt.Println("unexpected values", t.Values)
		panic("unexpected values")
	}

	switch t := t.Values[0].(type) {
	case *ast.BasicLit:
		n, _ := strconv.Atoi(t.Value)
		return n

	case *ast.UnaryExpr:
		x := t.X.(*ast.BasicLit)
		n, _ := strconv.Atoi(x.Value)
		return n

	default:
		fmt.Println("unexpected type", t)
		panic("unexpected type")
	}

	panic("fell through")
	return 0
}
