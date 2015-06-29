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

// Parse a spawnGroup manifest
// Format: <1 bit IsLZSSCompressed | 24 bit length | Data>
// Data: <8 bit arrayLength | 0 | 8 bit ressourceStrings | 0 | 8 bit unkown | 0 | dataTypes * arrayLength | 0 | ... data ...>
func (sg *spawnGroup) parse() {
    reader := newReader(sg.manifest)

    isCompressed := reader.readBoolean()
    size := reader.readBits(24)
    data := reader.readBytes(int(size))

    if isCompressed {
        dataUnc, err := unlzss(data)
        if err != nil {
            _panicf("Error uncompressing spawnGroup data %s", err)
        }

        data = dataUnc
    }

    reader2 := newReader(data)
    iSize := reader2.readBits(8)
    _ = iSize

    // The actual size reading is probably just readUntilNullByte
    // @Todo: find a package that is larger that 256 entries to investigate this
    
    // null1 := reader2.readBits(8)
    // ressourceStrings := reader2.readBits(8)
    // null2 := reader2.readBits(8)
    // unkown := reader2.readBits(8)
    // null3 := reader2.readBits(8)

    // @Todo: decypher the rest of the package and add it to the spawnGroup struct
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

    // Invoke the parse method, the data should be added to the spawnGroup variable
    // when it is fully decyphered

    if sg.complete {
        sg.manifest = m.GetSpawngroupmanifest()
        sg.complete = !m.GetManifestincomplete()
	    sg.parse()
    } else {
        sg.manifest = m.GetSpawngroupmanifest()
        sg.complete = !m.GetManifestincomplete()
	    sg.parse()
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