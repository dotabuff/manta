package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

var pp = spew.Dump

func main() {
	manta.DEBUG = true
	for _, arg := range os.Args[1:] {
		parser := manta.NewParserFromFile(arg)
		parser.HookNET(dota.NET_Messages_net_Tick, dbg)
		parser.HookSVC(dota.SVC_Messages_svc_PacketEntities, dbg)
		parser.HookSVC(dota.SVC_Messages_svc_Print, dbg)
		parser.Start()
	}
}

func dbg(m proto.Message) {
	pp(m)
}
