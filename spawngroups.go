package manta

import (
	"github.com/dotabuff/manta/dota"
)

// Represents the state of a spawngroup
type spawnGroup struct {
    worldName     string
    entLumpName   string
    entFilterName string
    handle        uint32
    ownerHandle   uint32
    manifest      string
    flags         uint32
    tickcount     int32
    localName     string
    parentName    string
    complete      bool
}

func (p *Parser) onCNETMsg_SpawnGroup_Load(m *dota.CNETMsg_SpawnGroup_Load) error {
    _debugf("Loading")
	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_ManifestUpdate(m *dota.CNETMsg_SpawnGroup_ManifestUpdate) error {
    _debugf("Update")
	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_SetCreationTick(m *dota.CNETMsg_SpawnGroup_SetCreationTick) error {
    _debugf("SetCreationTick")
	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_Unload(m *dota.CNETMsg_SpawnGroup_Unload) error {
    _debugf("Unload")
	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_LoadCompleted(m *dota.CNETMsg_SpawnGroup_LoadCompleted) error {
    _debugf("Load Complete")
	return nil
}