package manta

import (
    "io/ioutil"
    "fmt"
	"github.com/dotabuff/manta/dota"
)

// Represents the state of a spawngroup
// String parts are almost never filled
// Flags are almost always SPAWN_GROUP_CREATE_CLIENT_ONLY_ENTITIES & SPAWN_GROUP_LOAD_STREAMING_DATA
type spawnGroup struct {
    worldName     string
    entLumpName   string
    entFilterName string
    handle        uint32 // used as an index to refer to spawngroups in non-load packets
    ownerHandle   uint32
    manifest      []byte
    flags         uint32
    tickCount     int32
    localName     string
    parentName    string
    complete      bool
}

func (sg *spawnGroup) writeFixture() {
    // [id]_[isComplete]_sg_manifest.raw
    fname := fmt.Sprintf("%d_%t_sg_manifest.raw", sg.handle, sg.complete);
    err := ioutil.WriteFile(fname, sg.manifest, 0644)
    
    if err != nil {
        _panicf("Error writing spawnGroup fixture, %s", err)
    }
}

func (p *Parser) onCNETMsg_SpawnGroup_Load(m *dota.CNETMsg_SpawnGroup_Load) error {
    sg := &spawnGroup{
        worldName     : m.GetWorldname(),
        entLumpName   : m.GetEntitylumpname(),
        entFilterName : m.GetEntityfiltername(),
        handle        : m.GetSpawngrouphandle(),
        ownerHandle   : m.GetSpawngroupownerhandle(),
        manifest      : m.GetSpawngroupmanifest(),
        flags         : m.GetFlags(),
        tickCount     : m.GetTickcount(),
        localName     : m.GetLocalnamefixup(),
        parentName    : m.GetParentnamefixup(),
        complete      : !m.GetManifestincomplete(),
	}

	p.spawnGroups[m.GetSpawngrouphandle()] = sg
	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_ManifestUpdate(m *dota.CNETMsg_SpawnGroup_ManifestUpdate) error {
    sg, ok := p.spawnGroups[m.GetSpawngrouphandle()]
    if !ok {
        _panicf("Unable to find spawngroup %d for update %d", m.GetSpawngrouphandle(), p.Tick)
    }

    // Current rationale:
    // - If the previous spawn group is complete, overrride it with the new one
    // - If it's incomplete, append the rest of the data
    // - Always update the completion status
    //
    // There might be a delta-update mechanism but the manifest has to be parsed first

    if sg.complete {
        sg.manifest = m.GetSpawngroupmanifest()
        sg.complete = !m.GetManifestincomplete()
    } else {
        sg.manifest = append(sg.manifest, m.GetSpawngroupmanifest()...)
        sg.complete = !m.GetManifestincomplete()
    }

	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_SetCreationTick(m *dota.CNETMsg_SpawnGroup_SetCreationTick) error {
    sg, ok := p.spawnGroups[m.GetSpawngrouphandle()]
    if !ok {
        _panicf("Unable to find spawngroup %d for tick update", m.GetSpawngrouphandle())
    }

    sg.tickCount = m.GetTickcount()

	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_Unload(m *dota.CNETMsg_SpawnGroup_Unload) error {
    _, ok := p.spawnGroups[m.GetSpawngrouphandle()]
    if ok {
        // It seems that spawn groups are not always present when being unloaded
        delete(p.spawnGroups, m.GetSpawngrouphandle())
    }

	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_LoadCompleted(m *dota.CNETMsg_SpawnGroup_LoadCompleted) error {
    _debugf("Info: LoadComplete message found") // This doesn't seem to get called at all
	return nil
}