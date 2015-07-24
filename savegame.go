package manta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/manta/dota"
	"github.com/dotabuff/manta/vbkv"
)

func ParseCDemoSaveGame(s *dota.CDemoSaveGame) (result map[string]interface{}, err error) {
	buf := bytes.NewBuffer(s.GetData())
	magic := make([]byte, 4)
	magicLen, err := buf.Read(magic)

	if magicLen != len(magic) || !bytes.Equal(magic, []byte("VBKV")) {
		return nil, fmt.Errorf("wrong savegame magic: %v", magic)
	}

	// this seems to be some sort of size value... hasn't been useful.
	buf.Next(4)

	kv, err := vbkv.Parse(buf)

	if debugMode {
		filename := spew.Sprintf("fixtures/savegames/savegame_%v_%v_%v.json", s.GetSignature(), s.GetSteamId(), s.GetVersion())
		j, _ := json.MarshalIndent(kv, "", "  ")
		ioutil.WriteFile(filename, j, 0644)
	}

	return kv, err
}
