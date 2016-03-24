package manta

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/dotabuff/manta/dota"
	"github.com/dotabuff/manta/vbkv"
)

func ParseCDemoSaveGame(s *dota.CDemoSaveGame) (save *SaveGame, err error) {
	buf := bytes.NewBuffer(s.GetData())
	magic := buf.Next(4)

	if len(magic) != 4 || !bytes.Equal(magic, []byte("VBKV")) {
		return nil, fmt.Errorf("wrong savegame magic: %v", magic)
	}

	// this seems to be some sort of size value... hasn't been useful.
	buf.Next(4)

	kv, err := vbkv.Parse(buf)
	if err != nil {
		return nil, err
	}

	wbuf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(wbuf)
	err = enc.Encode(kv)
	if err != nil {
		return nil, err
	}

	rbuf := bytes.NewBuffer(wbuf.Bytes())
	dec := json.NewDecoder(rbuf)
	save = &SaveGame{}
	err = dec.Decode(save)

	return save, err
}

type SaveGame struct {
	Heroes               map[string]SaveGameHero                 `json:"Heroes"`
	Players              map[string]SaveGamePlayer               `json:"Players"`
	Roshan               SaveGameRoshan                          `json:"Roshan"`
	StockInfo            map[string]map[string]SaveGameStockInfo `json:"StockInfo"`
	Units                map[string]SaveGameUnit                 `json:"Units"`
	DireDeaths           int                                     `json:"dire_deaths"`
	DireGlyphCooldown    float32                                 `json:"dire_glyph_cooldown"`
	Dotatime             float32                                 `json:"dotatime"`
	Firstblood           int                                     `json:"firstblood"`
	Gameid               uint64                                  `json:"gameid"`
	Gametime             float32                                 `json:"gametime"`
	Matchid              uint64                                  `json:"matchid"`
	RadiantDeaths        int                                     `json:"radiant_deaths"`
	RadiantGlyphCooldown float32                                 `json:"radiant_glyph_cooldown"`
	Runespawntime        float32                                 `json:"runespawntime"`
	Savetime             uint64                                  `json:"savetime"` // this is unix time
	Version              int                                     `json:"version"`
}

type SaveGameRoshan struct {
	Alive       int     `json:"roshan_alive"`
	Kills       int     `json:"roshan_kills"`
	Killtime    float32 `json:"roshan_killtime"`
	Respawntime float32 `json:"roshan_respawntime"`
}

type SaveGameStockInfo struct {
	StockCount int     `json:"StockCount"`
	StockTime  float32 `json:"StockTime"`
}

type SaveGamePlayer struct {
	Assists        int    `json:"assists"`
	Deaths         int    `json:"deaths"`
	Denies         int    `json:"denies"`
	Hero           string `json:"hero"`
	Kills          int    `json:"kills"`
	Lasthits       int    `json:"lasthits"`
	Name           string `json:"name"`
	Reliablegold   int    `json:"reliablegold"`
	Sharemask      int    `json:"sharemask"`
	Steamid        uint64 `json:"steamid"`
	Streak         int    `json:"streak"`
	Team           int    `json:"team"`
	Unreliablegold int    `json:"unreliablegold"`
}

type SaveGameAbility struct {
	Cooldown float64 `json:"cooldown"`
	Hidden   int     `json:"hidden"`
	Level    int     `json:"level"`
	TeamNum  string  `json:"m_iTeamNum"`
}

type SaveGameInventoryItem struct {
	AssembledTime  string  `json:"m_flAssembledTime"`
	Cooldown       float32 `json:"cooldown"`
	CurrentCharges string  `json:"m_iCurrentCharges"`
	Hidden         int     `json:"hidden"`
	Level          int     `json:"level"`
	Owner          string  `json:"owner"`

	PurchaseTime       string `json:"m_flPurchaseTime"`
	PurchasedWhileDead string `json:"m_bPurchasedWhileDead"`
	SecondaryCharges   string `json:"m_iSecondaryCharges"`
	Slot               int    `json:"slot"`
	TeamNum            string `json:"m_iTeamNum"`
}

type SaveGameHero struct {
	Abilities                map[string]SaveGameAbility       `json:"Abilities"`
	AbilityPoints            string                           `json:"m_iAbilityPoints"`
	AccumulatedHeal          string                           `json:"m_flAccumulatedHeal"`
	Agility                  string                           `json:"m_flAgility"`
	BKBCharges               int                              `json:"bkbcharges"`
	BuybackCooldown          float32                          `json:"buybackcooldown"`
	BuybackGoldLimit         float32                          `json:"buybackgoldlimit"`
	CurrentLevel             string                           `json:"m_iCurrentLevel"`
	CurrentXP                string                           `json:"m_iCurrentXP"`
	Health                   string                           `json:"m_iHealth"`
	Intellect                string                           `json:"m_flIntellect"`
	Inventory                map[string]SaveGameInventoryItem `json:"Inventory"`
	IsControllableByPlayer64 string                           `json:"m_iIsControllableByPlayer64"`
	Mana                     string                           `json:"m_flMana"`
	MultipleKillCount        string                           `json:"m_iMultipleKillCount"`
	Name                     string                           `json:"m_iName"`
	Origin                   string                           `json:"m_vecOrigin"`
	RespawnTime              float32                          `json:"respawntime"`
	Rotation                 string                           `json:"m_angRotation"`
	Strength                 string                           `json:"m_flStrength"`
	TeamNum                  string                           `json:"m_iTeamNum"`
	UnitName                 string                           `json:"m_iszUnitName"`
	UseNeutralCreepBehaviour string                           `json:"m_bUseNeutralCreepBehavior"`
}

type SaveGameUnit struct {
	Abilities                map[string]SaveGameAbility       `json:"Abilities"`
	AbilityPoints            string                           `json:"m_iAbilityPoints"`
	AccumulatedHeal          string                           `json:"m_flAccumulatedHeal"`
	Agility                  string                           `json:"m_flAgility"`
	BKBCharges               int                              `json:"bkbcharges"`
	BuybackCooldown          float32                          `json:"buybackcooldown"`
	BuybackGoldLimit         float32                          `json:"buybackgoldlimit"`
	CurrentLevel             string                           `json:"m_iCurrentLevel"`
	CurrentXP                string                           `json:"m_iCurrentXP"`
	Health                   string                           `json:"m_iHealth"`
	Intellect                string                           `json:"m_flIntellect"`
	Inventory                map[string]SaveGameInventoryItem `json:"Inventory"`
	IsControllableByPlayer64 string                           `json:"m_iIsControllableByPlayer64"`
	Mana                     string                           `json:"m_flMana"`
	MultipleKillCount        string                           `json:"m_iMultipleKillCount"`
	Name                     string                           `json:"m_iName"`
	Origin                   string                           `json:"m_vecOrigin"`
	RespawnTime              float32                          `json:"respawntime"`
	Rotation                 string                           `json:"m_angRotation"`
	Strength                 string                           `json:"m_flStrength"`
	TeamNum                  string                           `json:"m_iTeamNum"`
	UnitName                 string                           `json:"m_iszUnitName"`
	UseNeutralCreepBehaviour string                           `json:"m_bUseNeutralCreepBehavior"`
}
