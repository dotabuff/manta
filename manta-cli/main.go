package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/manta"
	"github.com/golang/protobuf/proto"
)

var pp = spew.Dump

func main() {
	manta.DEBUG = false
	for _, arg := range os.Args[1:] {
		parser := manta.NewParserFromFile(arg)
		parser.Start()
	}
}

func dbg(m proto.Message) {
	fmt.Printf("%T\n", m)
}
