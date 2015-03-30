package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

func main() {
	for _, arg := range os.Args[1:] {
		parser := manta.NewParserFromFile(arg)
		parser.HookBUM(dota.EBaseUserMessages_UM_SayText2, func(m proto.Message) {
			msg := m.(*dota.CUserMessageSayText2)
			fmt.Printf("%s (%s) | %s: %s\n", filepath.Base(arg), msg.GetMessagename(), msg.GetParam1(), msg.GetParam2())
		})
		parser.Start()
	}
}
