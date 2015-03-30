package manta

import (
	"fmt"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

type Callbacks struct {
	OnCDemoStop                               func(*dota.CDemoStop) error
	OnCDemoFileHeader                         func(*dota.CDemoFileHeader) error
	OnCDemoFileInfo                           func(*dota.CDemoFileInfo) error
	OnCDemoSyncTick                           func(*dota.CDemoSyncTick) error
	OnCDemoSendTables                         func(*dota.CDemoSendTables) error
	OnCDemoClassInfo                          func(*dota.CDemoClassInfo) error
	OnCDemoStringTables                       func(*dota.CDemoStringTables) error
	OnCDemoPacket                             func(*dota.CDemoPacket) error
	OnSignonPacket                            func(*dota.CDemoPacket) error
	OnCDemoConsoleCmd                         func(*dota.CDemoConsoleCmd) error
	OnCDemoCustomData                         func(*dota.CDemoCustomData) error
	OnCDemoCustomDataCallbacks                func(*dota.CDemoCustomDataCallbacks) error
	OnCDemoUserCmd                            func(*dota.CDemoUserCmd) error
	OnCDemoFullPacket                         func(*dota.CDemoFullPacket) error
	OnCDemoSaveGame                           func(*dota.CDemoSaveGame) error
	OnCDemoSpawnGroups                        func(*dota.CDemoSpawnGroups) error
	OnCNETMsg_NOP                             func(*dota.CNETMsg_NOP) error
	OnCNETMsg_Disconnect                      func(*dota.CNETMsg_Disconnect) error
	OnCNETMsg_File                            func(*dota.CNETMsg_File) error
	OnCNETMsg_SplitScreenUser                 func(*dota.CNETMsg_SplitScreenUser) error
	OnCNETMsg_Tick                            func(*dota.CNETMsg_Tick) error
	OnCNETMsg_StringCmd                       func(*dota.CNETMsg_StringCmd) error
	OnCNETMsg_SetConVar                       func(*dota.CNETMsg_SetConVar) error
	OnCNETMsg_SignonState                     func(*dota.CNETMsg_SignonState) error
	OnCNETMsg_SpawnGroup_Load                 func(*dota.CNETMsg_SpawnGroup_Load) error
	OnCNETMsg_SpawnGroup_ManifestUpdate       func(*dota.CNETMsg_SpawnGroup_ManifestUpdate) error
	OnCNETMsg_SpawnGroup_ForceBlockingLoad    func(*dota.CNETMsg_SpawnGroup_ForceBlockingLoad) error
	OnCNETMsg_SpawnGroup_SetCreationTick      func(*dota.CNETMsg_SpawnGroup_SetCreationTick) error
	OnCNETMsg_SpawnGroup_Unload               func(*dota.CNETMsg_SpawnGroup_Unload) error
	OnCNETMsg_SpawnGroup_LoadCompleted        func(*dota.CNETMsg_SpawnGroup_LoadCompleted) error
	OnCSVCMsg_ServerInfo                      func(*dota.CSVCMsg_ServerInfo) error
	OnCSVCMsg_FlattenedSerializer             func(*dota.CSVCMsg_FlattenedSerializer) error
	OnCSVCMsg_ClassInfo                       func(*dota.CSVCMsg_ClassInfo) error
	OnCSVCMsg_SetPause                        func(*dota.CSVCMsg_SetPause) error
	OnCSVCMsg_CreateStringTable               func(*dota.CSVCMsg_CreateStringTable) error
	OnCSVCMsg_UpdateStringTable               func(*dota.CSVCMsg_UpdateStringTable) error
	OnCSVCMsg_VoiceInit                       func(*dota.CSVCMsg_VoiceInit) error
	OnCSVCMsg_VoiceData                       func(*dota.CSVCMsg_VoiceData) error
	OnCSVCMsg_Print                           func(*dota.CSVCMsg_Print) error
	OnCSVCMsg_Sounds                          func(*dota.CSVCMsg_Sounds) error
	OnCSVCMsg_SetView                         func(*dota.CSVCMsg_SetView) error
	OnCSVCMsg_ClearAllStringTables            func(*dota.CSVCMsg_ClearAllStringTables) error
	OnCSVCMsg_CmdKeyValues                    func(*dota.CSVCMsg_CmdKeyValues) error
	OnCSVCMsg_BSPDecal                        func(*dota.CSVCMsg_BSPDecal) error
	OnCSVCMsg_SplitScreen                     func(*dota.CSVCMsg_SplitScreen) error
	OnCSVCMsg_PacketEntities                  func(*dota.CSVCMsg_PacketEntities) error
	OnCSVCMsg_Prefetch                        func(*dota.CSVCMsg_Prefetch) error
	OnCSVCMsg_Menu                            func(*dota.CSVCMsg_Menu) error
	OnCSVCMsg_GetCvarValue                    func(*dota.CSVCMsg_GetCvarValue) error
	OnCSVCMsg_StopSound                       func(*dota.CSVCMsg_StopSound) error
	OnCSVCMsg_PeerList                        func(*dota.CSVCMsg_PeerList) error
	OnCSVCMsg_PacketReliable                  func(*dota.CSVCMsg_PacketReliable) error
	OnCSVCMsg_UserMessage                     func(*dota.CSVCMsg_UserMessage) error
	OnCSVCMsg_SendTable                       func(*dota.CSVCMsg_SendTable) error
	OnCSVCMsg_GameEvent                       func(*dota.CSVCMsg_GameEvent) error
	OnCSVCMsg_TempEntities                    func(*dota.CSVCMsg_TempEntities) error
	OnCSVCMsg_GameEventList                   func(*dota.CSVCMsg_GameEventList) error
	OnCSVCMsg_FullFrameSplit                  func(*dota.CSVCMsg_FullFrameSplit) error
	OnCDOTAUserMsg_AIDebugLine                func(*dota.CDOTAUserMsg_AIDebugLine) error
	OnCDOTAUserMsg_ChatEvent                  func(*dota.CDOTAUserMsg_ChatEvent) error
	OnCDOTAUserMsg_CombatHeroPositions        func(*dota.CDOTAUserMsg_CombatHeroPositions) error
	OnCDOTAUserMsg_CombatLogShowDeath         func(*dota.CDOTAUserMsg_CombatLogShowDeath) error
	OnCDOTAUserMsg_CreateLinearProjectile     func(*dota.CDOTAUserMsg_CreateLinearProjectile) error
	OnCDOTAUserMsg_DestroyLinearProjectile    func(*dota.CDOTAUserMsg_DestroyLinearProjectile) error
	OnCDOTAUserMsg_DodgeTrackingProjectiles   func(*dota.CDOTAUserMsg_DodgeTrackingProjectiles) error
	OnCDOTAUserMsg_GlobalLightColor           func(*dota.CDOTAUserMsg_GlobalLightColor) error
	OnCDOTAUserMsg_GlobalLightDirection       func(*dota.CDOTAUserMsg_GlobalLightDirection) error
	OnCDOTAUserMsg_InvalidCommand             func(*dota.CDOTAUserMsg_InvalidCommand) error
	OnCDOTAUserMsg_LocationPing               func(*dota.CDOTAUserMsg_LocationPing) error
	OnCDOTAUserMsg_MapLine                    func(*dota.CDOTAUserMsg_MapLine) error
	OnCDOTAUserMsg_MiniKillCamInfo            func(*dota.CDOTAUserMsg_MiniKillCamInfo) error
	OnCDOTAUserMsg_MinimapDebugPoint          func(*dota.CDOTAUserMsg_MinimapDebugPoint) error
	OnCDOTAUserMsg_MinimapEvent               func(*dota.CDOTAUserMsg_MinimapEvent) error
	OnCDOTAUserMsg_NevermoreRequiem           func(*dota.CDOTAUserMsg_NevermoreRequiem) error
	OnCDOTAUserMsg_OverheadEvent              func(*dota.CDOTAUserMsg_OverheadEvent) error
	OnCDOTAUserMsg_SetNextAutobuyItem         func(*dota.CDOTAUserMsg_SetNextAutobuyItem) error
	OnCDOTAUserMsg_SharedCooldown             func(*dota.CDOTAUserMsg_SharedCooldown) error
	OnCDOTAUserMsg_SpectatorPlayerClick       func(*dota.CDOTAUserMsg_SpectatorPlayerClick) error
	OnCDOTAUserMsg_TutorialTipInfo            func(*dota.CDOTAUserMsg_TutorialTipInfo) error
	OnCDOTAUserMsg_UnitEvent                  func(*dota.CDOTAUserMsg_UnitEvent) error
	OnCDOTAUserMsg_ParticleManager            func(*dota.CDOTAUserMsg_ParticleManager) error
	OnCDOTAUserMsg_BotChat                    func(*dota.CDOTAUserMsg_BotChat) error
	OnCDOTAUserMsg_HudError                   func(*dota.CDOTAUserMsg_HudError) error
	OnCDOTAUserMsg_ItemPurchased              func(*dota.CDOTAUserMsg_ItemPurchased) error
	OnCDOTAUserMsg_Ping                       func(*dota.CDOTAUserMsg_Ping) error
	OnCDOTAUserMsg_ItemFound                  func(*dota.CDOTAUserMsg_ItemFound) error
	OnCDOTAUserMsg_SwapVerify                 func(*dota.CDOTAUserMsg_SwapVerify) error
	OnCDOTAUserMsg_WorldLine                  func(*dota.CDOTAUserMsg_WorldLine) error
	OnCDOTAUserMsg_ItemAlert                  func(*dota.CDOTAUserMsg_ItemAlert) error
	OnCDOTAUserMsg_HalloweenDrops             func(*dota.CDOTAUserMsg_HalloweenDrops) error
	OnCDOTAUserMsg_ChatWheel                  func(*dota.CDOTAUserMsg_ChatWheel) error
	OnCDOTAUserMsg_ReceivedXmasGift           func(*dota.CDOTAUserMsg_ReceivedXmasGift) error
	OnCDOTAUserMsg_UpdateSharedContent        func(*dota.CDOTAUserMsg_UpdateSharedContent) error
	OnCDOTAUserMsg_TutorialRequestExp         func(*dota.CDOTAUserMsg_TutorialRequestExp) error
	OnCDOTAUserMsg_TutorialPingMinimap        func(*dota.CDOTAUserMsg_TutorialPingMinimap) error
	OnCDOTAUserMsg_GamerulesStateChanged      func(*dota.CDOTAUserMsg_GamerulesStateChanged) error
	OnCDOTAUserMsg_ShowSurvey                 func(*dota.CDOTAUserMsg_ShowSurvey) error
	OnCDOTAUserMsg_TutorialFade               func(*dota.CDOTAUserMsg_TutorialFade) error
	OnCDOTAUserMsg_AddQuestLogEntry           func(*dota.CDOTAUserMsg_AddQuestLogEntry) error
	OnCDOTAUserMsg_SendStatPopup              func(*dota.CDOTAUserMsg_SendStatPopup) error
	OnCDOTAUserMsg_TutorialFinish             func(*dota.CDOTAUserMsg_TutorialFinish) error
	OnCDOTAUserMsg_SendRoshanPopup            func(*dota.CDOTAUserMsg_SendRoshanPopup) error
	OnCDOTAUserMsg_SendGenericToolTip         func(*dota.CDOTAUserMsg_SendGenericToolTip) error
	OnCDOTAUserMsg_SendFinalGold              func(*dota.CDOTAUserMsg_SendFinalGold) error
	OnCDOTAUserMsg_CustomMsg                  func(*dota.CDOTAUserMsg_CustomMsg) error
	OnCDOTAUserMsg_CoachHUDPing               func(*dota.CDOTAUserMsg_CoachHUDPing) error
	OnCDOTAUserMsg_ClientLoadGridNav          func(*dota.CDOTAUserMsg_ClientLoadGridNav) error
	OnCDOTAUserMsg_TE_Projectile              func(*dota.CDOTAUserMsg_TE_Projectile) error
	OnCDOTAUserMsg_TE_ProjectileLoc           func(*dota.CDOTAUserMsg_TE_ProjectileLoc) error
	OnCDOTAUserMsg_TE_DotaBloodImpact         func(*dota.CDOTAUserMsg_TE_DotaBloodImpact) error
	OnCDOTAUserMsg_TE_UnitAnimation           func(*dota.CDOTAUserMsg_TE_UnitAnimation) error
	OnCDOTAUserMsg_TE_UnitAnimationEnd        func(*dota.CDOTAUserMsg_TE_UnitAnimationEnd) error
	OnCDOTAUserMsg_AbilityPing                func(*dota.CDOTAUserMsg_AbilityPing) error
	OnCDOTAUserMsg_ShowGenericPopup           func(*dota.CDOTAUserMsg_ShowGenericPopup) error
	OnCDOTAUserMsg_VoteStart                  func(*dota.CDOTAUserMsg_VoteStart) error
	OnCDOTAUserMsg_VoteUpdate                 func(*dota.CDOTAUserMsg_VoteUpdate) error
	OnCDOTAUserMsg_VoteEnd                    func(*dota.CDOTAUserMsg_VoteEnd) error
	OnCDOTAUserMsg_BoosterState               func(*dota.CDOTAUserMsg_BoosterState) error
	OnCDOTAUserMsg_WillPurchaseAlert          func(*dota.CDOTAUserMsg_WillPurchaseAlert) error
	OnCDOTAUserMsg_TutorialMinimapPosition    func(*dota.CDOTAUserMsg_TutorialMinimapPosition) error
	OnCDOTAUserMsg_PlayerMMR                  func(*dota.CDOTAUserMsg_PlayerMMR) error
	OnCDOTAUserMsg_AbilitySteal               func(*dota.CDOTAUserMsg_AbilitySteal) error
	OnCDOTAUserMsg_CourierKilledAlert         func(*dota.CDOTAUserMsg_CourierKilledAlert) error
	OnCDOTAUserMsg_EnemyItemAlert             func(*dota.CDOTAUserMsg_EnemyItemAlert) error
	OnCDOTAUserMsg_StatsMatchDetails          func(*dota.CDOTAUserMsg_StatsMatchDetails) error
	OnCDOTAUserMsg_MiniTaunt                  func(*dota.CDOTAUserMsg_MiniTaunt) error
	OnCDOTAUserMsg_BuyBackStateAlert          func(*dota.CDOTAUserMsg_BuyBackStateAlert) error
	OnCDOTAUserMsg_SpeechBubble               func(*dota.CDOTAUserMsg_SpeechBubble) error
	OnCDOTAUserMsg_CustomHeaderMessage        func(*dota.CDOTAUserMsg_CustomHeaderMessage) error
	OnCEntityMessagePlayJingle                func(*dota.CEntityMessagePlayJingle) error
	OnCEntityMessageScreenOverlay             func(*dota.CEntityMessageScreenOverlay) error
	OnCEntityMessageRemoveAllDecals           func(*dota.CEntityMessageRemoveAllDecals) error
	OnCEntityMessagePropagateForce            func(*dota.CEntityMessagePropagateForce) error
	OnCEntityMessageDoSpark                   func(*dota.CEntityMessageDoSpark) error
	OnCEntityMessageFixAngle                  func(*dota.CEntityMessageFixAngle) error
	OnCUserMessageAchievementEvent            func(*dota.CUserMessageAchievementEvent) error
	OnCUserMessageCloseCaption                func(*dota.CUserMessageCloseCaption) error
	OnCUserMessageCloseCaptionDirect          func(*dota.CUserMessageCloseCaptionDirect) error
	OnCUserMessageCurrentTimescale            func(*dota.CUserMessageCurrentTimescale) error
	OnCUserMessageDesiredTimescale            func(*dota.CUserMessageDesiredTimescale) error
	OnCUserMessageFade                        func(*dota.CUserMessageFade) error
	OnCUserMessageGameTitle                   func(*dota.CUserMessageGameTitle) error
	OnCUserMessageHintText                    func(*dota.CUserMessageHintText) error
	OnCUserMessageHudMsg                      func(*dota.CUserMessageHudMsg) error
	OnCUserMessageHudText                     func(*dota.CUserMessageHudText) error
	OnCUserMessageKeyHintText                 func(*dota.CUserMessageKeyHintText) error
	OnCUserMessageColoredText                 func(*dota.CUserMessageColoredText) error
	OnCUserMessageRequestState                func(*dota.CUserMessageRequestState) error
	OnCUserMessageResetHUD                    func(*dota.CUserMessageResetHUD) error
	OnCUserMessageRumble                      func(*dota.CUserMessageRumble) error
	OnCUserMessageSayText                     func(*dota.CUserMessageSayText) error
	OnCUserMessageSayText2                    func(*dota.CUserMessageSayText2) error
	OnCUserMessageSayTextChannel              func(*dota.CUserMessageSayTextChannel) error
	OnCUserMessageShake                       func(*dota.CUserMessageShake) error
	OnCUserMessageShakeDir                    func(*dota.CUserMessageShakeDir) error
	OnCUserMessageTextMsg                     func(*dota.CUserMessageTextMsg) error
	OnCUserMessageScreenTilt                  func(*dota.CUserMessageScreenTilt) error
	OnCUserMessageTrain                       func(*dota.CUserMessageTrain) error
	OnCUserMessageVGUIMenu                    func(*dota.CUserMessageVGUIMenu) error
	OnCUserMessageVoiceMask                   func(*dota.CUserMessageVoiceMask) error
	OnCUserMessageVoiceSubtitle               func(*dota.CUserMessageVoiceSubtitle) error
	OnCUserMessageSendAudio                   func(*dota.CUserMessageSendAudio) error
	OnCUserMessageItemPickup                  func(*dota.CUserMessageItemPickup) error
	OnCUserMessageAmmoDenied                  func(*dota.CUserMessageAmmoDenied) error
	OnCUserMessageCrosshairAngle              func(*dota.CUserMessageCrosshairAngle) error
	OnCUserMessageShowMenu                    func(*dota.CUserMessageShowMenu) error
	OnCUserMessageCreditsMsg                  func(*dota.CUserMessageCreditsMsg) error
	OnCUserMessageCloseCaptionPlaceholder     func(*dota.CUserMessageCloseCaptionPlaceholder) error
	OnCUserMessageCameraTransition            func(*dota.CUserMessageCameraTransition) error
	OnCMsgVDebugGameSessionIDEvent            func(*dota.CMsgVDebugGameSessionIDEvent) error
	OnCMsgPlaceDecalEvent                     func(*dota.CMsgPlaceDecalEvent) error
	OnCMsgClearWorldDecalsEvent               func(*dota.CMsgClearWorldDecalsEvent) error
	OnCMsgClearEntityDecalsEvent              func(*dota.CMsgClearEntityDecalsEvent) error
	OnCMsgClearDecalsForSkeletonInstanceEvent func(*dota.CMsgClearDecalsForSkeletonInstanceEvent) error
	OnCMsgSource1LegacyGameEventList          func(*dota.CMsgSource1LegacyGameEventList) error
	OnCMsgSource1LegacyListenEvents           func(*dota.CMsgSource1LegacyListenEvents) error
	OnCMsgSource1LegacyGameEvent              func(*dota.CMsgSource1LegacyGameEvent) error
	OnCMsgSosStartSoundEvent                  func(*dota.CMsgSosStartSoundEvent) error
	OnCMsgSosStopSoundEvent                   func(*dota.CMsgSosStopSoundEvent) error
	OnCMsgSosSetSoundEventParam               func(*dota.CMsgSosSetSoundEventParam) error
	OnCMsgSosSetLibraryStackField             func(*dota.CMsgSosSetLibraryStackField) error
	OnCMsgSosStopSoundEventHash               func(*dota.CMsgSosStopSoundEventHash) error
}

func (p *Parser) CallByDemoType(t int32, raw []byte) error {
	callbacks := p.Callbacks
	switch t {
	case 0: // dota.EDemoCommands_DEM_Stop
		if cb := callbacks.OnCDemoStop; cb != nil {
			msg := &dota.CDemoStop{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 1: // dota.EDemoCommands_DEM_FileHeader
		if cb := callbacks.OnCDemoFileHeader; cb != nil {
			msg := &dota.CDemoFileHeader{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 2: // dota.EDemoCommands_DEM_FileInfo
		if cb := callbacks.OnCDemoFileInfo; cb != nil {
			msg := &dota.CDemoFileInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 3: // dota.EDemoCommands_DEM_SyncTick
		if cb := callbacks.OnCDemoSyncTick; cb != nil {
			msg := &dota.CDemoSyncTick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 4: // dota.EDemoCommands_DEM_SendTables
		if cb := callbacks.OnCDemoSendTables; cb != nil {
			msg := &dota.CDemoSendTables{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 5: // dota.EDemoCommands_DEM_ClassInfo
		if cb := callbacks.OnCDemoClassInfo; cb != nil {
			msg := &dota.CDemoClassInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 6: // dota.EDemoCommands_DEM_StringTables
		if cb := callbacks.OnCDemoStringTables; cb != nil {
			msg := &dota.CDemoStringTables{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 7: // dota.EDemoCommands_DEM_Packet
		if cb := callbacks.OnCDemoPacket; cb != nil {
			msg := &dota.CDemoPacket{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 8: // dota.EDemoCommands_DEM_SignonPacket
		if cb := callbacks.OnSignonPacket; cb != nil {
			msg := &dota.CDemoPacket{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 9: // dota.EDemoCommands_DEM_ConsoleCmd
		if cb := callbacks.OnCDemoConsoleCmd; cb != nil {
			msg := &dota.CDemoConsoleCmd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 10: // dota.EDemoCommands_DEM_CustomData
		if cb := callbacks.OnCDemoCustomData; cb != nil {
			msg := &dota.CDemoCustomData{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 11: // dota.EDemoCommands_DEM_CustomDataCallbacks
		if cb := callbacks.OnCDemoCustomDataCallbacks; cb != nil {
			msg := &dota.CDemoCustomDataCallbacks{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 12: // dota.EDemoCommands_DEM_UserCmd
		if cb := callbacks.OnCDemoUserCmd; cb != nil {
			msg := &dota.CDemoUserCmd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 13: // dota.EDemoCommands_DEM_FullPacket
		if cb := callbacks.OnCDemoFullPacket; cb != nil {
			msg := &dota.CDemoFullPacket{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 14: // dota.EDemoCommands_DEM_SaveGame
		if cb := callbacks.OnCDemoSaveGame; cb != nil {
			msg := &dota.CDemoSaveGame{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 15: // dota.EDemoCommands_DEM_SpawnGroups
		if cb := callbacks.OnCDemoSpawnGroups; cb != nil {
			msg := &dota.CDemoSpawnGroups{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	}
	return fmt.Errorf("no type found: %d", t)
}

func (p *Parser) CallByPacketType(t int32, raw []byte) error {
	callbacks := p.Callbacks
	switch t {
	case 0: // dota.NET_Messages_net_NOP
		if cb := callbacks.OnCNETMsg_NOP; cb != nil {
			msg := &dota.CNETMsg_NOP{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 1: // dota.NET_Messages_net_Disconnect
		if cb := callbacks.OnCNETMsg_Disconnect; cb != nil {
			msg := &dota.CNETMsg_Disconnect{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 2: // dota.NET_Messages_net_File
		if cb := callbacks.OnCNETMsg_File; cb != nil {
			msg := &dota.CNETMsg_File{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 3: // dota.NET_Messages_net_SplitScreenUser
		if cb := callbacks.OnCNETMsg_SplitScreenUser; cb != nil {
			msg := &dota.CNETMsg_SplitScreenUser{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 4: // dota.NET_Messages_net_Tick
		if cb := callbacks.OnCNETMsg_Tick; cb != nil {
			msg := &dota.CNETMsg_Tick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 5: // dota.NET_Messages_net_StringCmd
		if cb := callbacks.OnCNETMsg_StringCmd; cb != nil {
			msg := &dota.CNETMsg_StringCmd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 6: // dota.NET_Messages_net_SetConVar
		if cb := callbacks.OnCNETMsg_SetConVar; cb != nil {
			msg := &dota.CNETMsg_SetConVar{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 7: // dota.NET_Messages_net_SignonState
		if cb := callbacks.OnCNETMsg_SignonState; cb != nil {
			msg := &dota.CNETMsg_SignonState{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 8: // dota.NET_Messages_net_SpawnGroup_Load
		if cb := callbacks.OnCNETMsg_SpawnGroup_Load; cb != nil {
			msg := &dota.CNETMsg_SpawnGroup_Load{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 9: // dota.NET_Messages_net_SpawnGroup_ManifestUpdate
		if cb := callbacks.OnCNETMsg_SpawnGroup_ManifestUpdate; cb != nil {
			msg := &dota.CNETMsg_SpawnGroup_ManifestUpdate{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 10: // dota.NET_Messages_net_SpawnGroup_ForceBlockingLoad
		if cb := callbacks.OnCNETMsg_SpawnGroup_ForceBlockingLoad; cb != nil {
			msg := &dota.CNETMsg_SpawnGroup_ForceBlockingLoad{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 11: // dota.NET_Messages_net_SpawnGroup_SetCreationTick
		if cb := callbacks.OnCNETMsg_SpawnGroup_SetCreationTick; cb != nil {
			msg := &dota.CNETMsg_SpawnGroup_SetCreationTick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 12: // dota.NET_Messages_net_SpawnGroup_Unload
		if cb := callbacks.OnCNETMsg_SpawnGroup_Unload; cb != nil {
			msg := &dota.CNETMsg_SpawnGroup_Unload{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 13: // dota.NET_Messages_net_SpawnGroup_LoadCompleted
		if cb := callbacks.OnCNETMsg_SpawnGroup_LoadCompleted; cb != nil {
			msg := &dota.CNETMsg_SpawnGroup_LoadCompleted{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 40: // dota.SVC_Messages_svc_ServerInfo
		if cb := callbacks.OnCSVCMsg_ServerInfo; cb != nil {
			msg := &dota.CSVCMsg_ServerInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 41: // dota.SVC_Messages_svc_FlattenedSerializer
		if cb := callbacks.OnCSVCMsg_FlattenedSerializer; cb != nil {
			msg := &dota.CSVCMsg_FlattenedSerializer{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 42: // dota.SVC_Messages_svc_ClassInfo
		if cb := callbacks.OnCSVCMsg_ClassInfo; cb != nil {
			msg := &dota.CSVCMsg_ClassInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 43: // dota.SVC_Messages_svc_SetPause
		if cb := callbacks.OnCSVCMsg_SetPause; cb != nil {
			msg := &dota.CSVCMsg_SetPause{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 44: // dota.SVC_Messages_svc_CreateStringTable
		if cb := callbacks.OnCSVCMsg_CreateStringTable; cb != nil {
			msg := &dota.CSVCMsg_CreateStringTable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 45: // dota.SVC_Messages_svc_UpdateStringTable
		if cb := callbacks.OnCSVCMsg_UpdateStringTable; cb != nil {
			msg := &dota.CSVCMsg_UpdateStringTable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 46: // dota.SVC_Messages_svc_VoiceInit
		if cb := callbacks.OnCSVCMsg_VoiceInit; cb != nil {
			msg := &dota.CSVCMsg_VoiceInit{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 47: // dota.SVC_Messages_svc_VoiceData
		if cb := callbacks.OnCSVCMsg_VoiceData; cb != nil {
			msg := &dota.CSVCMsg_VoiceData{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 48: // dota.SVC_Messages_svc_Print
		if cb := callbacks.OnCSVCMsg_Print; cb != nil {
			msg := &dota.CSVCMsg_Print{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 49: // dota.SVC_Messages_svc_Sounds
		if cb := callbacks.OnCSVCMsg_Sounds; cb != nil {
			msg := &dota.CSVCMsg_Sounds{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 50: // dota.SVC_Messages_svc_SetView
		if cb := callbacks.OnCSVCMsg_SetView; cb != nil {
			msg := &dota.CSVCMsg_SetView{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 51: // dota.SVC_Messages_svc_ClearAllStringTables
		if cb := callbacks.OnCSVCMsg_ClearAllStringTables; cb != nil {
			msg := &dota.CSVCMsg_ClearAllStringTables{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 52: // dota.SVC_Messages_svc_CmdKeyValues
		if cb := callbacks.OnCSVCMsg_CmdKeyValues; cb != nil {
			msg := &dota.CSVCMsg_CmdKeyValues{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 53: // dota.SVC_Messages_svc_BSPDecal
		if cb := callbacks.OnCSVCMsg_BSPDecal; cb != nil {
			msg := &dota.CSVCMsg_BSPDecal{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 54: // dota.SVC_Messages_svc_SplitScreen
		if cb := callbacks.OnCSVCMsg_SplitScreen; cb != nil {
			msg := &dota.CSVCMsg_SplitScreen{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 55: // dota.SVC_Messages_svc_PacketEntities
		if cb := callbacks.OnCSVCMsg_PacketEntities; cb != nil {
			msg := &dota.CSVCMsg_PacketEntities{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 56: // dota.SVC_Messages_svc_Prefetch
		if cb := callbacks.OnCSVCMsg_Prefetch; cb != nil {
			msg := &dota.CSVCMsg_Prefetch{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 57: // dota.SVC_Messages_svc_Menu
		if cb := callbacks.OnCSVCMsg_Menu; cb != nil {
			msg := &dota.CSVCMsg_Menu{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 58: // dota.SVC_Messages_svc_GetCvarValue
		if cb := callbacks.OnCSVCMsg_GetCvarValue; cb != nil {
			msg := &dota.CSVCMsg_GetCvarValue{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 59: // dota.SVC_Messages_svc_StopSound
		if cb := callbacks.OnCSVCMsg_StopSound; cb != nil {
			msg := &dota.CSVCMsg_StopSound{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 60: // dota.SVC_Messages_svc_PeerList
		if cb := callbacks.OnCSVCMsg_PeerList; cb != nil {
			msg := &dota.CSVCMsg_PeerList{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 61: // dota.SVC_Messages_svc_PacketReliable
		if cb := callbacks.OnCSVCMsg_PacketReliable; cb != nil {
			msg := &dota.CSVCMsg_PacketReliable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 62: // dota.SVC_Messages_svc_UserMessage
		if cb := callbacks.OnCSVCMsg_UserMessage; cb != nil {
			msg := &dota.CSVCMsg_UserMessage{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 63: // dota.SVC_Messages_svc_SendTable
		if cb := callbacks.OnCSVCMsg_SendTable; cb != nil {
			msg := &dota.CSVCMsg_SendTable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 67: // dota.SVC_Messages_svc_GameEvent
		if cb := callbacks.OnCSVCMsg_GameEvent; cb != nil {
			msg := &dota.CSVCMsg_GameEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 68: // dota.SVC_Messages_svc_TempEntities
		if cb := callbacks.OnCSVCMsg_TempEntities; cb != nil {
			msg := &dota.CSVCMsg_TempEntities{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 69: // dota.SVC_Messages_svc_GameEventList
		if cb := callbacks.OnCSVCMsg_GameEventList; cb != nil {
			msg := &dota.CSVCMsg_GameEventList{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 70: // dota.SVC_Messages_svc_FullFrameSplit
		if cb := callbacks.OnCSVCMsg_FullFrameSplit; cb != nil {
			msg := &dota.CSVCMsg_FullFrameSplit{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 465: // dota.EDotaUserMessages_DOTA_UM_AIDebugLine
		if cb := callbacks.OnCDOTAUserMsg_AIDebugLine; cb != nil {
			msg := &dota.CDOTAUserMsg_AIDebugLine{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 466: // dota.EDotaUserMessages_DOTA_UM_ChatEvent
		if cb := callbacks.OnCDOTAUserMsg_ChatEvent; cb != nil {
			msg := &dota.CDOTAUserMsg_ChatEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 467: // dota.EDotaUserMessages_DOTA_UM_CombatHeroPositions
		if cb := callbacks.OnCDOTAUserMsg_CombatHeroPositions; cb != nil {
			msg := &dota.CDOTAUserMsg_CombatHeroPositions{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 470: // dota.EDotaUserMessages_DOTA_UM_CombatLogShowDeath
		if cb := callbacks.OnCDOTAUserMsg_CombatLogShowDeath; cb != nil {
			msg := &dota.CDOTAUserMsg_CombatLogShowDeath{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 471: // dota.EDotaUserMessages_DOTA_UM_CreateLinearProjectile
		if cb := callbacks.OnCDOTAUserMsg_CreateLinearProjectile; cb != nil {
			msg := &dota.CDOTAUserMsg_CreateLinearProjectile{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 472: // dota.EDotaUserMessages_DOTA_UM_DestroyLinearProjectile
		if cb := callbacks.OnCDOTAUserMsg_DestroyLinearProjectile; cb != nil {
			msg := &dota.CDOTAUserMsg_DestroyLinearProjectile{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 473: // dota.EDotaUserMessages_DOTA_UM_DodgeTrackingProjectiles
		if cb := callbacks.OnCDOTAUserMsg_DodgeTrackingProjectiles; cb != nil {
			msg := &dota.CDOTAUserMsg_DodgeTrackingProjectiles{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 474: // dota.EDotaUserMessages_DOTA_UM_GlobalLightColor
		if cb := callbacks.OnCDOTAUserMsg_GlobalLightColor; cb != nil {
			msg := &dota.CDOTAUserMsg_GlobalLightColor{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 475: // dota.EDotaUserMessages_DOTA_UM_GlobalLightDirection
		if cb := callbacks.OnCDOTAUserMsg_GlobalLightDirection; cb != nil {
			msg := &dota.CDOTAUserMsg_GlobalLightDirection{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 476: // dota.EDotaUserMessages_DOTA_UM_InvalidCommand
		if cb := callbacks.OnCDOTAUserMsg_InvalidCommand; cb != nil {
			msg := &dota.CDOTAUserMsg_InvalidCommand{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 477: // dota.EDotaUserMessages_DOTA_UM_LocationPing
		if cb := callbacks.OnCDOTAUserMsg_LocationPing; cb != nil {
			msg := &dota.CDOTAUserMsg_LocationPing{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 478: // dota.EDotaUserMessages_DOTA_UM_MapLine
		if cb := callbacks.OnCDOTAUserMsg_MapLine; cb != nil {
			msg := &dota.CDOTAUserMsg_MapLine{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 479: // dota.EDotaUserMessages_DOTA_UM_MiniKillCamInfo
		if cb := callbacks.OnCDOTAUserMsg_MiniKillCamInfo; cb != nil {
			msg := &dota.CDOTAUserMsg_MiniKillCamInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 480: // dota.EDotaUserMessages_DOTA_UM_MinimapDebugPoint
		if cb := callbacks.OnCDOTAUserMsg_MinimapDebugPoint; cb != nil {
			msg := &dota.CDOTAUserMsg_MinimapDebugPoint{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 481: // dota.EDotaUserMessages_DOTA_UM_MinimapEvent
		if cb := callbacks.OnCDOTAUserMsg_MinimapEvent; cb != nil {
			msg := &dota.CDOTAUserMsg_MinimapEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 482: // dota.EDotaUserMessages_DOTA_UM_NevermoreRequiem
		if cb := callbacks.OnCDOTAUserMsg_NevermoreRequiem; cb != nil {
			msg := &dota.CDOTAUserMsg_NevermoreRequiem{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 483: // dota.EDotaUserMessages_DOTA_UM_OverheadEvent
		if cb := callbacks.OnCDOTAUserMsg_OverheadEvent; cb != nil {
			msg := &dota.CDOTAUserMsg_OverheadEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 484: // dota.EDotaUserMessages_DOTA_UM_SetNextAutobuyItem
		if cb := callbacks.OnCDOTAUserMsg_SetNextAutobuyItem; cb != nil {
			msg := &dota.CDOTAUserMsg_SetNextAutobuyItem{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 485: // dota.EDotaUserMessages_DOTA_UM_SharedCooldown
		if cb := callbacks.OnCDOTAUserMsg_SharedCooldown; cb != nil {
			msg := &dota.CDOTAUserMsg_SharedCooldown{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 486: // dota.EDotaUserMessages_DOTA_UM_SpectatorPlayerClick
		if cb := callbacks.OnCDOTAUserMsg_SpectatorPlayerClick; cb != nil {
			msg := &dota.CDOTAUserMsg_SpectatorPlayerClick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 487: // dota.EDotaUserMessages_DOTA_UM_TutorialTipInfo
		if cb := callbacks.OnCDOTAUserMsg_TutorialTipInfo; cb != nil {
			msg := &dota.CDOTAUserMsg_TutorialTipInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 488: // dota.EDotaUserMessages_DOTA_UM_UnitEvent
		if cb := callbacks.OnCDOTAUserMsg_UnitEvent; cb != nil {
			msg := &dota.CDOTAUserMsg_UnitEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 489: // dota.EDotaUserMessages_DOTA_UM_ParticleManager
		if cb := callbacks.OnCDOTAUserMsg_ParticleManager; cb != nil {
			msg := &dota.CDOTAUserMsg_ParticleManager{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 490: // dota.EDotaUserMessages_DOTA_UM_BotChat
		if cb := callbacks.OnCDOTAUserMsg_BotChat; cb != nil {
			msg := &dota.CDOTAUserMsg_BotChat{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 491: // dota.EDotaUserMessages_DOTA_UM_HudError
		if cb := callbacks.OnCDOTAUserMsg_HudError; cb != nil {
			msg := &dota.CDOTAUserMsg_HudError{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 492: // dota.EDotaUserMessages_DOTA_UM_ItemPurchased
		if cb := callbacks.OnCDOTAUserMsg_ItemPurchased; cb != nil {
			msg := &dota.CDOTAUserMsg_ItemPurchased{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 493: // dota.EDotaUserMessages_DOTA_UM_Ping
		if cb := callbacks.OnCDOTAUserMsg_Ping; cb != nil {
			msg := &dota.CDOTAUserMsg_Ping{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 494: // dota.EDotaUserMessages_DOTA_UM_ItemFound
		if cb := callbacks.OnCDOTAUserMsg_ItemFound; cb != nil {
			msg := &dota.CDOTAUserMsg_ItemFound{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 496: // dota.EDotaUserMessages_DOTA_UM_SwapVerify
		if cb := callbacks.OnCDOTAUserMsg_SwapVerify; cb != nil {
			msg := &dota.CDOTAUserMsg_SwapVerify{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 497: // dota.EDotaUserMessages_DOTA_UM_WorldLine
		if cb := callbacks.OnCDOTAUserMsg_WorldLine; cb != nil {
			msg := &dota.CDOTAUserMsg_WorldLine{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 499: // dota.EDotaUserMessages_DOTA_UM_ItemAlert
		if cb := callbacks.OnCDOTAUserMsg_ItemAlert; cb != nil {
			msg := &dota.CDOTAUserMsg_ItemAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 500: // dota.EDotaUserMessages_DOTA_UM_HalloweenDrops
		if cb := callbacks.OnCDOTAUserMsg_HalloweenDrops; cb != nil {
			msg := &dota.CDOTAUserMsg_HalloweenDrops{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 501: // dota.EDotaUserMessages_DOTA_UM_ChatWheel
		if cb := callbacks.OnCDOTAUserMsg_ChatWheel; cb != nil {
			msg := &dota.CDOTAUserMsg_ChatWheel{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 502: // dota.EDotaUserMessages_DOTA_UM_ReceivedXmasGift
		if cb := callbacks.OnCDOTAUserMsg_ReceivedXmasGift; cb != nil {
			msg := &dota.CDOTAUserMsg_ReceivedXmasGift{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 503: // dota.EDotaUserMessages_DOTA_UM_UpdateSharedContent
		if cb := callbacks.OnCDOTAUserMsg_UpdateSharedContent; cb != nil {
			msg := &dota.CDOTAUserMsg_UpdateSharedContent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 504: // dota.EDotaUserMessages_DOTA_UM_TutorialRequestExp
		if cb := callbacks.OnCDOTAUserMsg_TutorialRequestExp; cb != nil {
			msg := &dota.CDOTAUserMsg_TutorialRequestExp{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 505: // dota.EDotaUserMessages_DOTA_UM_TutorialPingMinimap
		if cb := callbacks.OnCDOTAUserMsg_TutorialPingMinimap; cb != nil {
			msg := &dota.CDOTAUserMsg_TutorialPingMinimap{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 506: // dota.EDotaUserMessages_DOTA_UM_GamerulesStateChanged
		if cb := callbacks.OnCDOTAUserMsg_GamerulesStateChanged; cb != nil {
			msg := &dota.CDOTAUserMsg_GamerulesStateChanged{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 507: // dota.EDotaUserMessages_DOTA_UM_ShowSurvey
		if cb := callbacks.OnCDOTAUserMsg_ShowSurvey; cb != nil {
			msg := &dota.CDOTAUserMsg_ShowSurvey{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 508: // dota.EDotaUserMessages_DOTA_UM_TutorialFade
		if cb := callbacks.OnCDOTAUserMsg_TutorialFade; cb != nil {
			msg := &dota.CDOTAUserMsg_TutorialFade{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 509: // dota.EDotaUserMessages_DOTA_UM_AddQuestLogEntry
		if cb := callbacks.OnCDOTAUserMsg_AddQuestLogEntry; cb != nil {
			msg := &dota.CDOTAUserMsg_AddQuestLogEntry{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 510: // dota.EDotaUserMessages_DOTA_UM_SendStatPopup
		if cb := callbacks.OnCDOTAUserMsg_SendStatPopup; cb != nil {
			msg := &dota.CDOTAUserMsg_SendStatPopup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 511: // dota.EDotaUserMessages_DOTA_UM_TutorialFinish
		if cb := callbacks.OnCDOTAUserMsg_TutorialFinish; cb != nil {
			msg := &dota.CDOTAUserMsg_TutorialFinish{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 512: // dota.EDotaUserMessages_DOTA_UM_SendRoshanPopup
		if cb := callbacks.OnCDOTAUserMsg_SendRoshanPopup; cb != nil {
			msg := &dota.CDOTAUserMsg_SendRoshanPopup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 513: // dota.EDotaUserMessages_DOTA_UM_SendGenericToolTip
		if cb := callbacks.OnCDOTAUserMsg_SendGenericToolTip; cb != nil {
			msg := &dota.CDOTAUserMsg_SendGenericToolTip{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 514: // dota.EDotaUserMessages_DOTA_UM_SendFinalGold
		if cb := callbacks.OnCDOTAUserMsg_SendFinalGold; cb != nil {
			msg := &dota.CDOTAUserMsg_SendFinalGold{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 515: // dota.EDotaUserMessages_DOTA_UM_CustomMsg
		if cb := callbacks.OnCDOTAUserMsg_CustomMsg; cb != nil {
			msg := &dota.CDOTAUserMsg_CustomMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 516: // dota.EDotaUserMessages_DOTA_UM_CoachHUDPing
		if cb := callbacks.OnCDOTAUserMsg_CoachHUDPing; cb != nil {
			msg := &dota.CDOTAUserMsg_CoachHUDPing{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 517: // dota.EDotaUserMessages_DOTA_UM_ClientLoadGridNav
		if cb := callbacks.OnCDOTAUserMsg_ClientLoadGridNav; cb != nil {
			msg := &dota.CDOTAUserMsg_ClientLoadGridNav{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 518: // dota.EDotaUserMessages_DOTA_UM_TE_Projectile
		if cb := callbacks.OnCDOTAUserMsg_TE_Projectile; cb != nil {
			msg := &dota.CDOTAUserMsg_TE_Projectile{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 519: // dota.EDotaUserMessages_DOTA_UM_TE_ProjectileLoc
		if cb := callbacks.OnCDOTAUserMsg_TE_ProjectileLoc; cb != nil {
			msg := &dota.CDOTAUserMsg_TE_ProjectileLoc{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 520: // dota.EDotaUserMessages_DOTA_UM_TE_DotaBloodImpact
		if cb := callbacks.OnCDOTAUserMsg_TE_DotaBloodImpact; cb != nil {
			msg := &dota.CDOTAUserMsg_TE_DotaBloodImpact{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 521: // dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimation
		if cb := callbacks.OnCDOTAUserMsg_TE_UnitAnimation; cb != nil {
			msg := &dota.CDOTAUserMsg_TE_UnitAnimation{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 522: // dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimationEnd
		if cb := callbacks.OnCDOTAUserMsg_TE_UnitAnimationEnd; cb != nil {
			msg := &dota.CDOTAUserMsg_TE_UnitAnimationEnd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 523: // dota.EDotaUserMessages_DOTA_UM_AbilityPing
		if cb := callbacks.OnCDOTAUserMsg_AbilityPing; cb != nil {
			msg := &dota.CDOTAUserMsg_AbilityPing{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 524: // dota.EDotaUserMessages_DOTA_UM_ShowGenericPopup
		if cb := callbacks.OnCDOTAUserMsg_ShowGenericPopup; cb != nil {
			msg := &dota.CDOTAUserMsg_ShowGenericPopup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 525: // dota.EDotaUserMessages_DOTA_UM_VoteStart
		if cb := callbacks.OnCDOTAUserMsg_VoteStart; cb != nil {
			msg := &dota.CDOTAUserMsg_VoteStart{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 526: // dota.EDotaUserMessages_DOTA_UM_VoteUpdate
		if cb := callbacks.OnCDOTAUserMsg_VoteUpdate; cb != nil {
			msg := &dota.CDOTAUserMsg_VoteUpdate{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 527: // dota.EDotaUserMessages_DOTA_UM_VoteEnd
		if cb := callbacks.OnCDOTAUserMsg_VoteEnd; cb != nil {
			msg := &dota.CDOTAUserMsg_VoteEnd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 528: // dota.EDotaUserMessages_DOTA_UM_BoosterState
		if cb := callbacks.OnCDOTAUserMsg_BoosterState; cb != nil {
			msg := &dota.CDOTAUserMsg_BoosterState{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 529: // dota.EDotaUserMessages_DOTA_UM_WillPurchaseAlert
		if cb := callbacks.OnCDOTAUserMsg_WillPurchaseAlert; cb != nil {
			msg := &dota.CDOTAUserMsg_WillPurchaseAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 530: // dota.EDotaUserMessages_DOTA_UM_TutorialMinimapPosition
		if cb := callbacks.OnCDOTAUserMsg_TutorialMinimapPosition; cb != nil {
			msg := &dota.CDOTAUserMsg_TutorialMinimapPosition{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 531: // dota.EDotaUserMessages_DOTA_UM_PlayerMMR
		if cb := callbacks.OnCDOTAUserMsg_PlayerMMR; cb != nil {
			msg := &dota.CDOTAUserMsg_PlayerMMR{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 532: // dota.EDotaUserMessages_DOTA_UM_AbilitySteal
		if cb := callbacks.OnCDOTAUserMsg_AbilitySteal; cb != nil {
			msg := &dota.CDOTAUserMsg_AbilitySteal{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 533: // dota.EDotaUserMessages_DOTA_UM_CourierKilledAlert
		if cb := callbacks.OnCDOTAUserMsg_CourierKilledAlert; cb != nil {
			msg := &dota.CDOTAUserMsg_CourierKilledAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 534: // dota.EDotaUserMessages_DOTA_UM_EnemyItemAlert
		if cb := callbacks.OnCDOTAUserMsg_EnemyItemAlert; cb != nil {
			msg := &dota.CDOTAUserMsg_EnemyItemAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 535: // dota.EDotaUserMessages_DOTA_UM_StatsMatchDetails
		if cb := callbacks.OnCDOTAUserMsg_StatsMatchDetails; cb != nil {
			msg := &dota.CDOTAUserMsg_StatsMatchDetails{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 536: // dota.EDotaUserMessages_DOTA_UM_MiniTaunt
		if cb := callbacks.OnCDOTAUserMsg_MiniTaunt; cb != nil {
			msg := &dota.CDOTAUserMsg_MiniTaunt{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 537: // dota.EDotaUserMessages_DOTA_UM_BuyBackStateAlert
		if cb := callbacks.OnCDOTAUserMsg_BuyBackStateAlert; cb != nil {
			msg := &dota.CDOTAUserMsg_BuyBackStateAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 538: // dota.EDotaUserMessages_DOTA_UM_SpeechBubble
		if cb := callbacks.OnCDOTAUserMsg_SpeechBubble; cb != nil {
			msg := &dota.CDOTAUserMsg_SpeechBubble{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 539: // dota.EDotaUserMessages_DOTA_UM_CustomHeaderMessage
		if cb := callbacks.OnCDOTAUserMsg_CustomHeaderMessage; cb != nil {
			msg := &dota.CDOTAUserMsg_CustomHeaderMessage{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 136: // dota.EBaseEntityMessages_EM_PlayJingle
		if cb := callbacks.OnCEntityMessagePlayJingle; cb != nil {
			msg := &dota.CEntityMessagePlayJingle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 137: // dota.EBaseEntityMessages_EM_ScreenOverlay
		if cb := callbacks.OnCEntityMessageScreenOverlay; cb != nil {
			msg := &dota.CEntityMessageScreenOverlay{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 138: // dota.EBaseEntityMessages_EM_RemoveAllDecals
		if cb := callbacks.OnCEntityMessageRemoveAllDecals; cb != nil {
			msg := &dota.CEntityMessageRemoveAllDecals{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 139: // dota.EBaseEntityMessages_EM_PropagateForce
		if cb := callbacks.OnCEntityMessagePropagateForce; cb != nil {
			msg := &dota.CEntityMessagePropagateForce{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 140: // dota.EBaseEntityMessages_EM_DoSpark
		if cb := callbacks.OnCEntityMessageDoSpark; cb != nil {
			msg := &dota.CEntityMessageDoSpark{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 141: // dota.EBaseEntityMessages_EM_FixAngle
		if cb := callbacks.OnCEntityMessageFixAngle; cb != nil {
			msg := &dota.CEntityMessageFixAngle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 101: // dota.EBaseUserMessages_UM_AchievementEvent
		if cb := callbacks.OnCUserMessageAchievementEvent; cb != nil {
			msg := &dota.CUserMessageAchievementEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 102: // dota.EBaseUserMessages_UM_CloseCaption
		if cb := callbacks.OnCUserMessageCloseCaption; cb != nil {
			msg := &dota.CUserMessageCloseCaption{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 103: // dota.EBaseUserMessages_UM_CloseCaptionDirect
		if cb := callbacks.OnCUserMessageCloseCaptionDirect; cb != nil {
			msg := &dota.CUserMessageCloseCaptionDirect{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 104: // dota.EBaseUserMessages_UM_CurrentTimescale
		if cb := callbacks.OnCUserMessageCurrentTimescale; cb != nil {
			msg := &dota.CUserMessageCurrentTimescale{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 105: // dota.EBaseUserMessages_UM_DesiredTimescale
		if cb := callbacks.OnCUserMessageDesiredTimescale; cb != nil {
			msg := &dota.CUserMessageDesiredTimescale{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 106: // dota.EBaseUserMessages_UM_Fade
		if cb := callbacks.OnCUserMessageFade; cb != nil {
			msg := &dota.CUserMessageFade{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 107: // dota.EBaseUserMessages_UM_GameTitle
		if cb := callbacks.OnCUserMessageGameTitle; cb != nil {
			msg := &dota.CUserMessageGameTitle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 109: // dota.EBaseUserMessages_UM_HintText
		if cb := callbacks.OnCUserMessageHintText; cb != nil {
			msg := &dota.CUserMessageHintText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 110: // dota.EBaseUserMessages_UM_HudMsg
		if cb := callbacks.OnCUserMessageHudMsg; cb != nil {
			msg := &dota.CUserMessageHudMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 111: // dota.EBaseUserMessages_UM_HudText
		if cb := callbacks.OnCUserMessageHudText; cb != nil {
			msg := &dota.CUserMessageHudText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 112: // dota.EBaseUserMessages_UM_KeyHintText
		if cb := callbacks.OnCUserMessageKeyHintText; cb != nil {
			msg := &dota.CUserMessageKeyHintText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 113: // dota.EBaseUserMessages_UM_ColoredText
		if cb := callbacks.OnCUserMessageColoredText; cb != nil {
			msg := &dota.CUserMessageColoredText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 114: // dota.EBaseUserMessages_UM_RequestState
		if cb := callbacks.OnCUserMessageRequestState; cb != nil {
			msg := &dota.CUserMessageRequestState{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 115: // dota.EBaseUserMessages_UM_ResetHUD
		if cb := callbacks.OnCUserMessageResetHUD; cb != nil {
			msg := &dota.CUserMessageResetHUD{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 116: // dota.EBaseUserMessages_UM_Rumble
		if cb := callbacks.OnCUserMessageRumble; cb != nil {
			msg := &dota.CUserMessageRumble{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 117: // dota.EBaseUserMessages_UM_SayText
		if cb := callbacks.OnCUserMessageSayText; cb != nil {
			msg := &dota.CUserMessageSayText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 118: // dota.EBaseUserMessages_UM_SayText2
		if cb := callbacks.OnCUserMessageSayText2; cb != nil {
			msg := &dota.CUserMessageSayText2{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 119: // dota.EBaseUserMessages_UM_SayTextChannel
		if cb := callbacks.OnCUserMessageSayTextChannel; cb != nil {
			msg := &dota.CUserMessageSayTextChannel{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 120: // dota.EBaseUserMessages_UM_Shake
		if cb := callbacks.OnCUserMessageShake; cb != nil {
			msg := &dota.CUserMessageShake{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 121: // dota.EBaseUserMessages_UM_ShakeDir
		if cb := callbacks.OnCUserMessageShakeDir; cb != nil {
			msg := &dota.CUserMessageShakeDir{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 124: // dota.EBaseUserMessages_UM_TextMsg
		if cb := callbacks.OnCUserMessageTextMsg; cb != nil {
			msg := &dota.CUserMessageTextMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 125: // dota.EBaseUserMessages_UM_ScreenTilt
		if cb := callbacks.OnCUserMessageScreenTilt; cb != nil {
			msg := &dota.CUserMessageScreenTilt{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 126: // dota.EBaseUserMessages_UM_Train
		if cb := callbacks.OnCUserMessageTrain; cb != nil {
			msg := &dota.CUserMessageTrain{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 127: // dota.EBaseUserMessages_UM_VGUIMenu
		if cb := callbacks.OnCUserMessageVGUIMenu; cb != nil {
			msg := &dota.CUserMessageVGUIMenu{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 128: // dota.EBaseUserMessages_UM_VoiceMask
		if cb := callbacks.OnCUserMessageVoiceMask; cb != nil {
			msg := &dota.CUserMessageVoiceMask{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 129: // dota.EBaseUserMessages_UM_VoiceSubtitle
		if cb := callbacks.OnCUserMessageVoiceSubtitle; cb != nil {
			msg := &dota.CUserMessageVoiceSubtitle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 130: // dota.EBaseUserMessages_UM_SendAudio
		if cb := callbacks.OnCUserMessageSendAudio; cb != nil {
			msg := &dota.CUserMessageSendAudio{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 131: // dota.EBaseUserMessages_UM_ItemPickup
		if cb := callbacks.OnCUserMessageItemPickup; cb != nil {
			msg := &dota.CUserMessageItemPickup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 132: // dota.EBaseUserMessages_UM_AmmoDenied
		if cb := callbacks.OnCUserMessageAmmoDenied; cb != nil {
			msg := &dota.CUserMessageAmmoDenied{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 133: // dota.EBaseUserMessages_UM_CrosshairAngle
		if cb := callbacks.OnCUserMessageCrosshairAngle; cb != nil {
			msg := &dota.CUserMessageCrosshairAngle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 134: // dota.EBaseUserMessages_UM_ShowMenu
		if cb := callbacks.OnCUserMessageShowMenu; cb != nil {
			msg := &dota.CUserMessageShowMenu{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 135: // dota.EBaseUserMessages_UM_CreditsMsg
		if cb := callbacks.OnCUserMessageCreditsMsg; cb != nil {
			msg := &dota.CUserMessageCreditsMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 142: // dota.EBaseUserMessages_UM_CloseCaptionPlaceholder
		if cb := callbacks.OnCUserMessageCloseCaptionPlaceholder; cb != nil {
			msg := &dota.CUserMessageCloseCaptionPlaceholder{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 143: // dota.EBaseUserMessages_UM_CameraTransition
		if cb := callbacks.OnCUserMessageCameraTransition; cb != nil {
			msg := &dota.CUserMessageCameraTransition{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 200: // dota.EBaseGameEvents_GE_VDebugGameSessionIDEvent
		if cb := callbacks.OnCMsgVDebugGameSessionIDEvent; cb != nil {
			msg := &dota.CMsgVDebugGameSessionIDEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 201: // dota.EBaseGameEvents_GE_PlaceDecalEvent
		if cb := callbacks.OnCMsgPlaceDecalEvent; cb != nil {
			msg := &dota.CMsgPlaceDecalEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 202: // dota.EBaseGameEvents_GE_ClearWorldDecalsEvent
		if cb := callbacks.OnCMsgClearWorldDecalsEvent; cb != nil {
			msg := &dota.CMsgClearWorldDecalsEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 203: // dota.EBaseGameEvents_GE_ClearEntityDecalsEvent
		if cb := callbacks.OnCMsgClearEntityDecalsEvent; cb != nil {
			msg := &dota.CMsgClearEntityDecalsEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 204: // dota.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
		if cb := callbacks.OnCMsgClearDecalsForSkeletonInstanceEvent; cb != nil {
			msg := &dota.CMsgClearDecalsForSkeletonInstanceEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 205: // dota.EBaseGameEvents_GE_Source1LegacyGameEventList
		if cb := callbacks.OnCMsgSource1LegacyGameEventList; cb != nil {
			msg := &dota.CMsgSource1LegacyGameEventList{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 206: // dota.EBaseGameEvents_GE_Source1LegacyListenEvents
		if cb := callbacks.OnCMsgSource1LegacyListenEvents; cb != nil {
			msg := &dota.CMsgSource1LegacyListenEvents{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 207: // dota.EBaseGameEvents_GE_Source1LegacyGameEvent
		if cb := callbacks.OnCMsgSource1LegacyGameEvent; cb != nil {
			msg := &dota.CMsgSource1LegacyGameEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 208: // dota.EBaseGameEvents_GE_SosStartSoundEvent
		if cb := callbacks.OnCMsgSosStartSoundEvent; cb != nil {
			msg := &dota.CMsgSosStartSoundEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 209: // dota.EBaseGameEvents_GE_SosStopSoundEvent
		if cb := callbacks.OnCMsgSosStopSoundEvent; cb != nil {
			msg := &dota.CMsgSosStopSoundEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 210: // dota.EBaseGameEvents_GE_SosSetSoundEventParam
		if cb := callbacks.OnCMsgSosSetSoundEventParam; cb != nil {
			msg := &dota.CMsgSosSetSoundEventParam{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 211: // dota.EBaseGameEvents_GE_SosSetLibraryStackField
		if cb := callbacks.OnCMsgSosSetLibraryStackField; cb != nil {
			msg := &dota.CMsgSosSetLibraryStackField{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	case 212: // dota.EBaseGameEvents_GE_SosStopSoundEventHash
		if cb := callbacks.OnCMsgSosStopSoundEventHash; cb != nil {
			msg := &dota.CMsgSosStopSoundEventHash{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			return cb(msg)
		}
		return nil
	}
	return fmt.Errorf("no type found: %d", t)
}
