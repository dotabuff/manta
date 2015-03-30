package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/cookieo9/go-misc/slice"
	"github.com/davecgh/go-spew/spew"
)

var pp = spew.Dump

type EnumMap struct {
	Hook, Struct, Enum, Fill string
	EnumNames, StructNames   []string
	Values                   map[string]int
}

func main() {
	protoDir, outFile := os.Args[1], os.Args[2]

	fset := &token.FileSet{}
	pkgs, err := parser.ParseDir(fset, protoDir, nil, 0)
	if err != nil {
		panic(err)
	}

	enums := []*EnumMap{
		&EnumMap{Hook: "DEM", Struct: "CDemo", Enum: "EDemoCommands", Fill: "_DEM_"},
		&EnumMap{Hook: "NET", Struct: "CNETMsg_", Enum: "NET_Messages", Fill: "_net_"},
		&EnumMap{Hook: "SVC", Struct: "CSVCMsg_", Enum: "SVC_Messages", Fill: "_svc_"},
		&EnumMap{Hook: "DUM", Struct: "CDOTAUserMsg_", Enum: "EDotaUserMessages", Fill: "_DOTA_UM_"},
		&EnumMap{Hook: "BEM", Struct: "CEntityMessage", Enum: "EBaseEntityMessages", Fill: "_EM_"},
		&EnumMap{Hook: "BUM", Struct: "CUserMessage", Enum: "EBaseUserMessages", Fill: "_UM_"},
		&EnumMap{Hook: "BGE", Struct: "CMsg", Enum: "EBaseGameEvents", Fill: "_GE_"},
	}

	for _, enum := range enums {
		enum.EnumNames = []string{}
		enum.StructNames = []string{}
		enum.Values = map[string]int{}
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, iDecl := range file.Decls {
				switch decl := iDecl.(type) {
				case *ast.GenDecl:
					for _, iSpec := range decl.Specs {
						switch spec := iSpec.(type) {
						case *ast.ValueSpec:
							switch valueSpecType := spec.Type.(type) {
							case *ast.Ident:
								for _, enum := range enums {
									if enum.Enum == valueSpecType.String() {
										var eValue int
										for _, iValue := range spec.Values {
											switch value := iValue.(type) {
											case *ast.BasicLit:
												if value.Kind == token.INT {
													eValue, _ = strconv.Atoi(value.Value)
												} else {
													panic(spew.Errorf("%v", value))
												}
											case *ast.UnaryExpr:
												switch value.Op {
												case token.SUB:
													switch x := value.X.(type) {
													case *ast.BasicLit:
														if x.Kind == token.INT {
															i, _ := strconv.Atoi(x.Value)
															eValue = -i
														} else {
															panic(spew.Errorf("-%v", x))
														}
													}
												default:
													pp(spec)
													panic("not SUB")
												}
											default:
												pp(spec)
												panic("not basic lit")
											}
										}
										for _, name := range spec.Names {
											enum.EnumNames = append(enum.EnumNames, name.String())
											enum.Values[name.String()] = eValue
										}
									}
								}
							}
						case *ast.TypeSpec:
							name := spec.Name.String()
							for _, enum := range enums {
								if strings.HasPrefix(name, enum.Struct) {
									enum.StructNames = append(enum.StructNames, name)
								}
							}
						}
					}
				}
			}
		}
	}

	file := bytes.NewBufferString(`
package manta
import (
	"fmt"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)
	`)

	values := map[int]string{}
	rawMsg := []string{}
	rawHook := []string{}

	for _, enum := range enums {
		switches := []string{}

		slice.SortTyped(&enum.EnumNames, func(a, b string) bool {
			return enum.Values[a] < enum.Values[b]
		})

		for _, e := range enum.EnumNames {
			matching := ""

			for _, name := range enum.StructNames {
				if name[len(enum.Struct):] == e[len(enum.Enum)+len(enum.Fill):] {
					matching = name
					break
				}
			}

			if matching == "" {
				switch e {
				case "EDotaUserMessages_DOTA_UM_AddUnitToSelection",
					"EDotaUserMessages_DOTA_UM_CombatLogData",
					"EDotaUserMessages_DOTA_UM_CharacterSpeakConcept",
					"EDotaUserMessages_DOTA_UM_TournamentDrop",
					"EBaseUserMessages_UM_MAX_BASE":
					continue
				case "EDemoCommands_DEM_Error", "EDemoCommands_DEM_Max", "EDemoCommands_DEM_IsCompressed":
					continue
				case "EDemoCommands_DEM_SignonPacket":
					matching = "CDemoPacket"
				default:
					pp(e)
				}
			}

			if enum.Hook != "DEM" {
				if prev, found := values[enum.Values[e]]; found {
					pp(matching, e, prev, enum.Values[e])
					panic("dupe")
				} else {
					values[enum.Values[e]] = e
				}
			}

			if matching == "" {
				pp(e, enum.Values[e])
				panic("no matching enum found")
			}

			switches = append(switches,
				spew.Sprintf("case dota.%s: // %d", e, enum.Values[e]),
				spew.Sprintf("  return &dota.%s{}, nil", matching),
			)
		}

		funName := "MessageTypeFor" + enum.Enum
		typName := "dota." + enum.Enum

		file.WriteString(spew.Sprintf(`
func %s(t %s) (proto.Message, error) {
	switch t {
	%s
	}
	return nil, fmt.Errorf("no type found: %v(%%d)", t)
}
		`, funName, typName, strings.Join(switches, "\n"), typName))

		if enum.Hook != "DEM" {
			hookVar := strings.ToLower(enum.Hook)
			rawMsg = append(rawMsg, spew.Sprintf(`
	%s := %s(t)
	if m, err = %s(%s); err == nil {
		if hook, ok := p.hook%s[%s]; ok {
			callHook(b, m, hook)
			return nil
		} else if debug {
			fmt.Printf("ignoring %%T\n", m)
		}
	}
		`, hookVar, typName, funName, hookVar, enum.Hook, hookVar))
		}

		rawHook = append(rawHook,
			spew.Sprintf(`func (p *Parser) Hook%s(t %s, f func(proto.Message)){p.hook%s[t] = f }`,
				enum.Hook, typName, enum.Hook))
	}

	file.WriteString(spew.Sprintf(`
func (p *Parser) HandleRawMessage(t int32, b []byte, debug bool) error {
	var m proto.Message
	var err error

	%s

	return fmt.Errorf("missing handler for %%d", t)
}
	`, strings.Join(rawMsg, "\n")))

	file.WriteString(strings.Join(rawHook, "\n"))

	source, err := format.Source(file.Bytes())
	if err != nil {
		spew.Println(file.String())
		panic(err)
	}

	err = ioutil.WriteFile(outFile, source, 0644)
	if err != nil {
		panic(err)
	}
}
