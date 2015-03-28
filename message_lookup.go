package main

import (
	"fmt"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

func MessageTypeForEDemoCommands(t dota.EDemoCommands) (proto.Message, error) {
	switch t {
	case dota.EDemoCommands_DEM_Stop: // 0
		return &dota.CDemoStop{}, nil
	case dota.EDemoCommands_DEM_FileHeader: // 1
		return &dota.CDemoFileHeader{}, nil
	case dota.EDemoCommands_DEM_FileInfo: // 2
		return &dota.CDemoFileInfo{}, nil
	case dota.EDemoCommands_DEM_SyncTick: // 3
		return &dota.CDemoSyncTick{}, nil
	case dota.EDemoCommands_DEM_SendTables: // 4
		return &dota.CDemoSendTables{}, nil
	case dota.EDemoCommands_DEM_ClassInfo: // 5
		return &dota.CDemoClassInfo{}, nil
	case dota.EDemoCommands_DEM_StringTables: // 6
		return &dota.CDemoStringTables{}, nil
	case dota.EDemoCommands_DEM_Packet: // 7
		return &dota.CDemoPacket{}, nil
	case dota.EDemoCommands_DEM_SignonPacket: // 8
		return &dota.CDemoPacket{}, nil
	case dota.EDemoCommands_DEM_ConsoleCmd: // 9
		return &dota.CDemoConsoleCmd{}, nil
	case dota.EDemoCommands_DEM_CustomData: // 10
		return &dota.CDemoCustomData{}, nil
	case dota.EDemoCommands_DEM_CustomDataCallbacks: // 11
		return &dota.CDemoCustomDataCallbacks{}, nil
	case dota.EDemoCommands_DEM_UserCmd: // 12
		return &dota.CDemoUserCmd{}, nil
	case dota.EDemoCommands_DEM_FullPacket: // 13
		return &dota.CDemoFullPacket{}, nil
	case dota.EDemoCommands_DEM_SaveGame: // 14
		return &dota.CDemoSaveGame{}, nil
	case dota.EDemoCommands_DEM_SpawnGroups: // 15
		return &dota.CDemoSpawnGroups{}, nil
	}
	return nil, fmt.Errorf("no type found: dota.EDemoCommands(%d)", t)
}

func MessageTypeForNET_Messages(t dota.NET_Messages) (proto.Message, error) {
	switch t {
	case dota.NET_Messages_net_NOP: // 0
		return &dota.CNETMsg_NOP{}, nil
	case dota.NET_Messages_net_Disconnect: // 1
		return &dota.CNETMsg_Disconnect{}, nil
	case dota.NET_Messages_net_File: // 2
		return &dota.CNETMsg_File{}, nil
	case dota.NET_Messages_net_SplitScreenUser: // 3
		return &dota.CNETMsg_SplitScreenUser{}, nil
	case dota.NET_Messages_net_Tick: // 4
		return &dota.CNETMsg_Tick{}, nil
	case dota.NET_Messages_net_StringCmd: // 5
		return &dota.CNETMsg_StringCmd{}, nil
	case dota.NET_Messages_net_SetConVar: // 6
		return &dota.CNETMsg_SetConVar{}, nil
	case dota.NET_Messages_net_SignonState: // 7
		return &dota.CNETMsg_SignonState{}, nil
	case dota.NET_Messages_net_SpawnGroup_Load: // 8
		return &dota.CNETMsg_SpawnGroup_Load{}, nil
	case dota.NET_Messages_net_SpawnGroup_ManifestUpdate: // 9
		return &dota.CNETMsg_SpawnGroup_ManifestUpdate{}, nil
	case dota.NET_Messages_net_SpawnGroup_ForceBlockingLoad: // 10
		return &dota.CNETMsg_SpawnGroup_ForceBlockingLoad{}, nil
	case dota.NET_Messages_net_SpawnGroup_SetCreationTick: // 11
		return &dota.CNETMsg_SpawnGroup_SetCreationTick{}, nil
	case dota.NET_Messages_net_SpawnGroup_Unload: // 12
		return &dota.CNETMsg_SpawnGroup_Unload{}, nil
	case dota.NET_Messages_net_SpawnGroup_LoadCompleted: // 13
		return &dota.CNETMsg_SpawnGroup_LoadCompleted{}, nil
	}
	return nil, fmt.Errorf("no type found: dota.NET_Messages(%d)", t)
}

func MessageTypeForSVC_Messages(t dota.SVC_Messages) (proto.Message, error) {
	switch t {
	case dota.SVC_Messages_svc_ServerInfo: // 40
		return &dota.CSVCMsg_ServerInfo{}, nil
	case dota.SVC_Messages_svc_FlattenedSerializer: // 41
		return &dota.CSVCMsg_FlattenedSerializer{}, nil
	case dota.SVC_Messages_svc_ClassInfo: // 42
		return &dota.CSVCMsg_ClassInfo{}, nil
	case dota.SVC_Messages_svc_SetPause: // 43
		return &dota.CSVCMsg_SetPause{}, nil
	case dota.SVC_Messages_svc_CreateStringTable: // 44
		return &dota.CSVCMsg_CreateStringTable{}, nil
	case dota.SVC_Messages_svc_UpdateStringTable: // 45
		return &dota.CSVCMsg_UpdateStringTable{}, nil
	case dota.SVC_Messages_svc_VoiceInit: // 46
		return &dota.CSVCMsg_VoiceInit{}, nil
	case dota.SVC_Messages_svc_VoiceData: // 47
		return &dota.CSVCMsg_VoiceData{}, nil
	case dota.SVC_Messages_svc_Print: // 48
		return &dota.CSVCMsg_Print{}, nil
	case dota.SVC_Messages_svc_Sounds: // 49
		return &dota.CSVCMsg_Sounds{}, nil
	case dota.SVC_Messages_svc_SetView: // 50
		return &dota.CSVCMsg_SetView{}, nil
	case dota.SVC_Messages_svc_ClearAllStringTables: // 51
		return &dota.CSVCMsg_ClearAllStringTables{}, nil
	case dota.SVC_Messages_svc_CmdKeyValues: // 52
		return &dota.CSVCMsg_CmdKeyValues{}, nil
	case dota.SVC_Messages_svc_BSPDecal: // 53
		return &dota.CSVCMsg_BSPDecal{}, nil
	case dota.SVC_Messages_svc_SplitScreen: // 54
		return &dota.CSVCMsg_SplitScreen{}, nil
	case dota.SVC_Messages_svc_PacketEntities: // 55
		return &dota.CSVCMsg_PacketEntities{}, nil
	case dota.SVC_Messages_svc_Prefetch: // 56
		return &dota.CSVCMsg_Prefetch{}, nil
	case dota.SVC_Messages_svc_Menu: // 57
		return &dota.CSVCMsg_Menu{}, nil
	case dota.SVC_Messages_svc_GetCvarValue: // 58
		return &dota.CSVCMsg_GetCvarValue{}, nil
	case dota.SVC_Messages_svc_StopSound: // 59
		return &dota.CSVCMsg_StopSound{}, nil
	case dota.SVC_Messages_svc_PeerList: // 60
		return &dota.CSVCMsg_PeerList{}, nil
	case dota.SVC_Messages_svc_PacketReliable: // 61
		return &dota.CSVCMsg_PacketReliable{}, nil
	case dota.SVC_Messages_svc_UserMessage: // 62
		return &dota.CSVCMsg_UserMessage{}, nil
	case dota.SVC_Messages_svc_SendTable: // 63
		return &dota.CSVCMsg_SendTable{}, nil
	case dota.SVC_Messages_svc_GameEvent: // 67
		return &dota.CSVCMsg_GameEvent{}, nil
	case dota.SVC_Messages_svc_TempEntities: // 68
		return &dota.CSVCMsg_TempEntities{}, nil
	case dota.SVC_Messages_svc_GameEventList: // 69
		return &dota.CSVCMsg_GameEventList{}, nil
	case dota.SVC_Messages_svc_FullFrameSplit: // 70
		return &dota.CSVCMsg_FullFrameSplit{}, nil
	}
	return nil, fmt.Errorf("no type found: dota.SVC_Messages(%d)", t)
}

func MessageTypeForEDotaUserMessages(t dota.EDotaUserMessages) (proto.Message, error) {
	switch t {
	case dota.EDotaUserMessages_DOTA_UM_AIDebugLine: // 465
		return &dota.CDOTAUserMsg_AIDebugLine{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ChatEvent: // 466
		return &dota.CDOTAUserMsg_ChatEvent{}, nil
	case dota.EDotaUserMessages_DOTA_UM_CombatHeroPositions: // 467
		return &dota.CDOTAUserMsg_CombatHeroPositions{}, nil
	case dota.EDotaUserMessages_DOTA_UM_CombatLogShowDeath: // 470
		return &dota.CDOTAUserMsg_CombatLogShowDeath{}, nil
	case dota.EDotaUserMessages_DOTA_UM_CreateLinearProjectile: // 471
		return &dota.CDOTAUserMsg_CreateLinearProjectile{}, nil
	case dota.EDotaUserMessages_DOTA_UM_DestroyLinearProjectile: // 472
		return &dota.CDOTAUserMsg_DestroyLinearProjectile{}, nil
	case dota.EDotaUserMessages_DOTA_UM_DodgeTrackingProjectiles: // 473
		return &dota.CDOTAUserMsg_DodgeTrackingProjectiles{}, nil
	case dota.EDotaUserMessages_DOTA_UM_GlobalLightColor: // 474
		return &dota.CDOTAUserMsg_GlobalLightColor{}, nil
	case dota.EDotaUserMessages_DOTA_UM_GlobalLightDirection: // 475
		return &dota.CDOTAUserMsg_GlobalLightDirection{}, nil
	case dota.EDotaUserMessages_DOTA_UM_InvalidCommand: // 476
		return &dota.CDOTAUserMsg_InvalidCommand{}, nil
	case dota.EDotaUserMessages_DOTA_UM_LocationPing: // 477
		return &dota.CDOTAUserMsg_LocationPing{}, nil
	case dota.EDotaUserMessages_DOTA_UM_MapLine: // 478
		return &dota.CDOTAUserMsg_MapLine{}, nil
	case dota.EDotaUserMessages_DOTA_UM_MiniKillCamInfo: // 479
		return &dota.CDOTAUserMsg_MiniKillCamInfo{}, nil
	case dota.EDotaUserMessages_DOTA_UM_MinimapDebugPoint: // 480
		return &dota.CDOTAUserMsg_MinimapDebugPoint{}, nil
	case dota.EDotaUserMessages_DOTA_UM_MinimapEvent: // 481
		return &dota.CDOTAUserMsg_MinimapEvent{}, nil
	case dota.EDotaUserMessages_DOTA_UM_NevermoreRequiem: // 482
		return &dota.CDOTAUserMsg_NevermoreRequiem{}, nil
	case dota.EDotaUserMessages_DOTA_UM_OverheadEvent: // 483
		return &dota.CDOTAUserMsg_OverheadEvent{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SetNextAutobuyItem: // 484
		return &dota.CDOTAUserMsg_SetNextAutobuyItem{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SharedCooldown: // 485
		return &dota.CDOTAUserMsg_SharedCooldown{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SpectatorPlayerClick: // 486
		return &dota.CDOTAUserMsg_SpectatorPlayerClick{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TutorialTipInfo: // 487
		return &dota.CDOTAUserMsg_TutorialTipInfo{}, nil
	case dota.EDotaUserMessages_DOTA_UM_UnitEvent: // 488
		return &dota.CDOTAUserMsg_UnitEvent{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ParticleManager: // 489
		return &dota.CDOTAUserMsg_ParticleManager{}, nil
	case dota.EDotaUserMessages_DOTA_UM_BotChat: // 490
		return &dota.CDOTAUserMsg_BotChat{}, nil
	case dota.EDotaUserMessages_DOTA_UM_HudError: // 491
		return &dota.CDOTAUserMsg_HudError{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ItemPurchased: // 492
		return &dota.CDOTAUserMsg_ItemPurchased{}, nil
	case dota.EDotaUserMessages_DOTA_UM_Ping: // 493
		return &dota.CDOTAUserMsg_Ping{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ItemFound: // 494
		return &dota.CDOTAUserMsg_ItemFound{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SwapVerify: // 496
		return &dota.CDOTAUserMsg_SwapVerify{}, nil
	case dota.EDotaUserMessages_DOTA_UM_WorldLine: // 497
		return &dota.CDOTAUserMsg_WorldLine{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ItemAlert: // 499
		return &dota.CDOTAUserMsg_ItemAlert{}, nil
	case dota.EDotaUserMessages_DOTA_UM_HalloweenDrops: // 500
		return &dota.CDOTAUserMsg_HalloweenDrops{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ChatWheel: // 501
		return &dota.CDOTAUserMsg_ChatWheel{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ReceivedXmasGift: // 502
		return &dota.CDOTAUserMsg_ReceivedXmasGift{}, nil
	case dota.EDotaUserMessages_DOTA_UM_UpdateSharedContent: // 503
		return &dota.CDOTAUserMsg_UpdateSharedContent{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TutorialRequestExp: // 504
		return &dota.CDOTAUserMsg_TutorialRequestExp{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TutorialPingMinimap: // 505
		return &dota.CDOTAUserMsg_TutorialPingMinimap{}, nil
	case dota.EDotaUserMessages_DOTA_UM_GamerulesStateChanged: // 506
		return &dota.CDOTAUserMsg_GamerulesStateChanged{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ShowSurvey: // 507
		return &dota.CDOTAUserMsg_ShowSurvey{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TutorialFade: // 508
		return &dota.CDOTAUserMsg_TutorialFade{}, nil
	case dota.EDotaUserMessages_DOTA_UM_AddQuestLogEntry: // 509
		return &dota.CDOTAUserMsg_AddQuestLogEntry{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SendStatPopup: // 510
		return &dota.CDOTAUserMsg_SendStatPopup{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TutorialFinish: // 511
		return &dota.CDOTAUserMsg_TutorialFinish{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SendRoshanPopup: // 512
		return &dota.CDOTAUserMsg_SendRoshanPopup{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SendGenericToolTip: // 513
		return &dota.CDOTAUserMsg_SendGenericToolTip{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SendFinalGold: // 514
		return &dota.CDOTAUserMsg_SendFinalGold{}, nil
	case dota.EDotaUserMessages_DOTA_UM_CustomMsg: // 515
		return &dota.CDOTAUserMsg_CustomMsg{}, nil
	case dota.EDotaUserMessages_DOTA_UM_CoachHUDPing: // 516
		return &dota.CDOTAUserMsg_CoachHUDPing{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ClientLoadGridNav: // 517
		return &dota.CDOTAUserMsg_ClientLoadGridNav{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TE_Projectile: // 518
		return &dota.CDOTAUserMsg_TE_Projectile{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TE_ProjectileLoc: // 519
		return &dota.CDOTAUserMsg_TE_ProjectileLoc{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TE_DotaBloodImpact: // 520
		return &dota.CDOTAUserMsg_TE_DotaBloodImpact{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimation: // 521
		return &dota.CDOTAUserMsg_TE_UnitAnimation{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimationEnd: // 522
		return &dota.CDOTAUserMsg_TE_UnitAnimationEnd{}, nil
	case dota.EDotaUserMessages_DOTA_UM_AbilityPing: // 523
		return &dota.CDOTAUserMsg_AbilityPing{}, nil
	case dota.EDotaUserMessages_DOTA_UM_ShowGenericPopup: // 524
		return &dota.CDOTAUserMsg_ShowGenericPopup{}, nil
	case dota.EDotaUserMessages_DOTA_UM_VoteStart: // 525
		return &dota.CDOTAUserMsg_VoteStart{}, nil
	case dota.EDotaUserMessages_DOTA_UM_VoteUpdate: // 526
		return &dota.CDOTAUserMsg_VoteUpdate{}, nil
	case dota.EDotaUserMessages_DOTA_UM_VoteEnd: // 527
		return &dota.CDOTAUserMsg_VoteEnd{}, nil
	case dota.EDotaUserMessages_DOTA_UM_BoosterState: // 528
		return &dota.CDOTAUserMsg_BoosterState{}, nil
	case dota.EDotaUserMessages_DOTA_UM_WillPurchaseAlert: // 529
		return &dota.CDOTAUserMsg_WillPurchaseAlert{}, nil
	case dota.EDotaUserMessages_DOTA_UM_TutorialMinimapPosition: // 530
		return &dota.CDOTAUserMsg_TutorialMinimapPosition{}, nil
	case dota.EDotaUserMessages_DOTA_UM_PlayerMMR: // 531
		return &dota.CDOTAUserMsg_PlayerMMR{}, nil
	case dota.EDotaUserMessages_DOTA_UM_AbilitySteal: // 532
		return &dota.CDOTAUserMsg_AbilitySteal{}, nil
	case dota.EDotaUserMessages_DOTA_UM_CourierKilledAlert: // 533
		return &dota.CDOTAUserMsg_CourierKilledAlert{}, nil
	case dota.EDotaUserMessages_DOTA_UM_EnemyItemAlert: // 534
		return &dota.CDOTAUserMsg_EnemyItemAlert{}, nil
	case dota.EDotaUserMessages_DOTA_UM_StatsMatchDetails: // 535
		return &dota.CDOTAUserMsg_StatsMatchDetails{}, nil
	case dota.EDotaUserMessages_DOTA_UM_MiniTaunt: // 536
		return &dota.CDOTAUserMsg_MiniTaunt{}, nil
	case dota.EDotaUserMessages_DOTA_UM_BuyBackStateAlert: // 537
		return &dota.CDOTAUserMsg_BuyBackStateAlert{}, nil
	case dota.EDotaUserMessages_DOTA_UM_SpeechBubble: // 538
		return &dota.CDOTAUserMsg_SpeechBubble{}, nil
	case dota.EDotaUserMessages_DOTA_UM_CustomHeaderMessage: // 539
		return &dota.CDOTAUserMsg_CustomHeaderMessage{}, nil
	}
	return nil, fmt.Errorf("no type found: dota.EDotaUserMessages(%d)", t)
}

func MessageTypeForEBaseEntityMessages(t dota.EBaseEntityMessages) (proto.Message, error) {
	switch t {
	case dota.EBaseEntityMessages_EM_PlayJingle: // 136
		return &dota.CEntityMessagePlayJingle{}, nil
	case dota.EBaseEntityMessages_EM_ScreenOverlay: // 137
		return &dota.CEntityMessageScreenOverlay{}, nil
	case dota.EBaseEntityMessages_EM_RemoveAllDecals: // 138
		return &dota.CEntityMessageRemoveAllDecals{}, nil
	case dota.EBaseEntityMessages_EM_PropagateForce: // 139
		return &dota.CEntityMessagePropagateForce{}, nil
	case dota.EBaseEntityMessages_EM_DoSpark: // 140
		return &dota.CEntityMessageDoSpark{}, nil
	case dota.EBaseEntityMessages_EM_FixAngle: // 141
		return &dota.CEntityMessageFixAngle{}, nil
	}
	return nil, fmt.Errorf("no type found: dota.EBaseEntityMessages(%d)", t)
}

func MessageTypeForEBaseUserMessages(t dota.EBaseUserMessages) (proto.Message, error) {
	switch t {
	case dota.EBaseUserMessages_UM_AchievementEvent: // 101
		return &dota.CUserMessageAchievementEvent{}, nil
	case dota.EBaseUserMessages_UM_CloseCaption: // 102
		return &dota.CUserMessageCloseCaption{}, nil
	case dota.EBaseUserMessages_UM_CloseCaptionDirect: // 103
		return &dota.CUserMessageCloseCaptionDirect{}, nil
	case dota.EBaseUserMessages_UM_CurrentTimescale: // 104
		return &dota.CUserMessageCurrentTimescale{}, nil
	case dota.EBaseUserMessages_UM_DesiredTimescale: // 105
		return &dota.CUserMessageDesiredTimescale{}, nil
	case dota.EBaseUserMessages_UM_Fade: // 106
		return &dota.CUserMessageFade{}, nil
	case dota.EBaseUserMessages_UM_GameTitle: // 107
		return &dota.CUserMessageGameTitle{}, nil
	case dota.EBaseUserMessages_UM_HintText: // 109
		return &dota.CUserMessageHintText{}, nil
	case dota.EBaseUserMessages_UM_HudMsg: // 110
		return &dota.CUserMessageHudMsg{}, nil
	case dota.EBaseUserMessages_UM_HudText: // 111
		return &dota.CUserMessageHudText{}, nil
	case dota.EBaseUserMessages_UM_KeyHintText: // 112
		return &dota.CUserMessageKeyHintText{}, nil
	case dota.EBaseUserMessages_UM_ColoredText: // 113
		return &dota.CUserMessageColoredText{}, nil
	case dota.EBaseUserMessages_UM_RequestState: // 114
		return &dota.CUserMessageRequestState{}, nil
	case dota.EBaseUserMessages_UM_ResetHUD: // 115
		return &dota.CUserMessageResetHUD{}, nil
	case dota.EBaseUserMessages_UM_Rumble: // 116
		return &dota.CUserMessageRumble{}, nil
	case dota.EBaseUserMessages_UM_SayText: // 117
		return &dota.CUserMessageSayText{}, nil
	case dota.EBaseUserMessages_UM_SayText2: // 118
		return &dota.CUserMessageSayText2{}, nil
	case dota.EBaseUserMessages_UM_SayTextChannel: // 119
		return &dota.CUserMessageSayTextChannel{}, nil
	case dota.EBaseUserMessages_UM_Shake: // 120
		return &dota.CUserMessageShake{}, nil
	case dota.EBaseUserMessages_UM_ShakeDir: // 121
		return &dota.CUserMessageShakeDir{}, nil
	case dota.EBaseUserMessages_UM_TextMsg: // 124
		return &dota.CUserMessageTextMsg{}, nil
	case dota.EBaseUserMessages_UM_ScreenTilt: // 125
		return &dota.CUserMessageScreenTilt{}, nil
	case dota.EBaseUserMessages_UM_Train: // 126
		return &dota.CUserMessageTrain{}, nil
	case dota.EBaseUserMessages_UM_VGUIMenu: // 127
		return &dota.CUserMessageVGUIMenu{}, nil
	case dota.EBaseUserMessages_UM_VoiceMask: // 128
		return &dota.CUserMessageVoiceMask{}, nil
	case dota.EBaseUserMessages_UM_VoiceSubtitle: // 129
		return &dota.CUserMessageVoiceSubtitle{}, nil
	case dota.EBaseUserMessages_UM_SendAudio: // 130
		return &dota.CUserMessageSendAudio{}, nil
	case dota.EBaseUserMessages_UM_ItemPickup: // 131
		return &dota.CUserMessageItemPickup{}, nil
	case dota.EBaseUserMessages_UM_AmmoDenied: // 132
		return &dota.CUserMessageAmmoDenied{}, nil
	case dota.EBaseUserMessages_UM_CrosshairAngle: // 133
		return &dota.CUserMessageCrosshairAngle{}, nil
	case dota.EBaseUserMessages_UM_ShowMenu: // 134
		return &dota.CUserMessageShowMenu{}, nil
	case dota.EBaseUserMessages_UM_CreditsMsg: // 135
		return &dota.CUserMessageCreditsMsg{}, nil
	case dota.EBaseUserMessages_UM_CloseCaptionPlaceholder: // 142
		return &dota.CUserMessageCloseCaptionPlaceholder{}, nil
	case dota.EBaseUserMessages_UM_CameraTransition: // 143
		return &dota.CUserMessageCameraTransition{}, nil
	}
	return nil, fmt.Errorf("no type found: dota.EBaseUserMessages(%d)", t)
}

func MessageTypeForEBaseGameEvents(t dota.EBaseGameEvents) (proto.Message, error) {
	switch t {
	case dota.EBaseGameEvents_GE_VDebugGameSessionIDEvent: // 200
		return &dota.CMsgVDebugGameSessionIDEvent{}, nil
	case dota.EBaseGameEvents_GE_PlaceDecalEvent: // 201
		return &dota.CMsgPlaceDecalEvent{}, nil
	case dota.EBaseGameEvents_GE_ClearWorldDecalsEvent: // 202
		return &dota.CMsgClearWorldDecalsEvent{}, nil
	case dota.EBaseGameEvents_GE_ClearEntityDecalsEvent: // 203
		return &dota.CMsgClearEntityDecalsEvent{}, nil
	case dota.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent: // 204
		return &dota.CMsgClearDecalsForSkeletonInstanceEvent{}, nil
	case dota.EBaseGameEvents_GE_Source1LegacyGameEventList: // 205
		return &dota.CMsgSource1LegacyGameEventList{}, nil
	case dota.EBaseGameEvents_GE_Source1LegacyListenEvents: // 206
		return &dota.CMsgSource1LegacyListenEvents{}, nil
	case dota.EBaseGameEvents_GE_Source1LegacyGameEvent: // 207
		return &dota.CMsgSource1LegacyGameEvent{}, nil
	case dota.EBaseGameEvents_GE_SosStartSoundEvent: // 208
		return &dota.CMsgSosStartSoundEvent{}, nil
	case dota.EBaseGameEvents_GE_SosStopSoundEvent: // 209
		return &dota.CMsgSosStopSoundEvent{}, nil
	case dota.EBaseGameEvents_GE_SosSetSoundEventParam: // 210
		return &dota.CMsgSosSetSoundEventParam{}, nil
	case dota.EBaseGameEvents_GE_SosSetLibraryStackField: // 211
		return &dota.CMsgSosSetLibraryStackField{}, nil
	case dota.EBaseGameEvents_GE_SosStopSoundEventHash: // 212
		return &dota.CMsgSosStopSoundEventHash{}, nil
	}
	return nil, fmt.Errorf("no type found: dota.EBaseGameEvents(%d)", t)
}

func (p *Parser) HandleRawMessage(t int32, b []byte, debug bool) error {
	var m proto.Message
	var err error

	net := dota.NET_Messages(t)
	if m, err = MessageTypeForNET_Messages(net); err == nil {
		if hook, ok := p.hookNET[net]; ok {
			callHook(b, m, hook)
			return nil
		} else if debug {
			fmt.Printf("ignoring %T\n", m)
		}
	}

	svc := dota.SVC_Messages(t)
	if m, err = MessageTypeForSVC_Messages(svc); err == nil {
		if hook, ok := p.hookSVC[svc]; ok {
			callHook(b, m, hook)
			return nil
		} else if debug {
			fmt.Printf("ignoring %T\n", m)
		}
	}

	dum := dota.EDotaUserMessages(t)
	if m, err = MessageTypeForEDotaUserMessages(dum); err == nil {
		if hook, ok := p.hookDUM[dum]; ok {
			callHook(b, m, hook)
			return nil
		} else if debug {
			fmt.Printf("ignoring %T\n", m)
		}
	}

	bem := dota.EBaseEntityMessages(t)
	if m, err = MessageTypeForEBaseEntityMessages(bem); err == nil {
		if hook, ok := p.hookBEM[bem]; ok {
			callHook(b, m, hook)
			return nil
		} else if debug {
			fmt.Printf("ignoring %T\n", m)
		}
	}

	bum := dota.EBaseUserMessages(t)
	if m, err = MessageTypeForEBaseUserMessages(bum); err == nil {
		if hook, ok := p.hookBUM[bum]; ok {
			callHook(b, m, hook)
			return nil
		} else if debug {
			fmt.Printf("ignoring %T\n", m)
		}
	}

	bge := dota.EBaseGameEvents(t)
	if m, err = MessageTypeForEBaseGameEvents(bge); err == nil {
		if hook, ok := p.hookBGE[bge]; ok {
			callHook(b, m, hook)
			return nil
		} else if debug {
			fmt.Printf("ignoring %T\n", m)
		}
	}

	return fmt.Errorf("missing handler for %d", t)
}
func (p *Parser) HookDEM(t dota.EDemoCommands, f func(proto.Message))       { p.hookDEM[t] = f }
func (p *Parser) HookNET(t dota.NET_Messages, f func(proto.Message))        { p.hookNET[t] = f }
func (p *Parser) HookSVC(t dota.SVC_Messages, f func(proto.Message))        { p.hookSVC[t] = f }
func (p *Parser) HookDUM(t dota.EDotaUserMessages, f func(proto.Message))   { p.hookDUM[t] = f }
func (p *Parser) HookBEM(t dota.EBaseEntityMessages, f func(proto.Message)) { p.hookBEM[t] = f }
func (p *Parser) HookBUM(t dota.EBaseUserMessages, f func(proto.Message))   { p.hookBUM[t] = f }
func (p *Parser) HookBGE(t dota.EBaseGameEvents, f func(proto.Message))     { p.hookBGE[t] = f }
