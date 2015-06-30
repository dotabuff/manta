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
	rTypes := reader2.readBits(16) // number of different ressource types in the data

	if rTypes > 1 {
		rStrings := reader2.readBits(16) // number of models / particles / ressources to load
		_ = reader2.readBits(16) // currently not known, probably the size for another field

		for i := 0; uint32(i) < rTypes; i++ {
			_ = reader2.readString() // e.g. vmdl, vmat, vpcf
		}

		for i := 0; uint32(i) < rStrings; i++ {
			_ = reader2.readString() // e.g. models/items/rubick/peculiar_prestidigitator_shoulders/
		}
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
	sg.parse()
	return nil
}

func (p *Parser) onCNETMsg_SpawnGroup_ManifestUpdate(m *dota.CNETMsg_SpawnGroup_ManifestUpdate) error {
	sg, ok := p.spawnGroups[m.GetSpawngrouphandle()]
	if !ok {
		_panicf("Unable to find spawngroup %d for update %d", m.GetSpawngrouphandle(), p.Tick)
	}

	// Invoke the parse method, the data should be added to the spawnGroup variable
	// when it is fully decyphered

	sg.manifest = m.GetSpawngroupmanifest()
	sg.complete = !m.GetManifestincomplete()
	sg.parse()

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