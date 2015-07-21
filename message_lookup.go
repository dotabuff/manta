//go:generate go run gen/message_lookup.go dota message_lookup.go
package manta

import (
	"fmt"

	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

var packetNames = map[int32]string{
	0:   "NET_Messages_net_NOP",
	1:   "NET_Messages_net_Disconnect",
	2:   "NET_Messages_net_File",
	3:   "NET_Messages_net_SplitScreenUser",
	4:   "NET_Messages_net_Tick",
	5:   "NET_Messages_net_StringCmd",
	6:   "NET_Messages_net_SetConVar",
	7:   "NET_Messages_net_SignonState",
	8:   "NET_Messages_net_SpawnGroup_Load",
	9:   "NET_Messages_net_SpawnGroup_ManifestUpdate",
	11:  "NET_Messages_net_SpawnGroup_SetCreationTick",
	12:  "NET_Messages_net_SpawnGroup_Unload",
	13:  "NET_Messages_net_SpawnGroup_LoadCompleted",
	14:  "NET_Messages_net_ReliableMessageEndMarker",
	40:  "SVC_Messages_svc_ServerInfo",
	41:  "SVC_Messages_svc_FlattenedSerializer",
	42:  "SVC_Messages_svc_ClassInfo",
	43:  "SVC_Messages_svc_SetPause",
	44:  "SVC_Messages_svc_CreateStringTable",
	45:  "SVC_Messages_svc_UpdateStringTable",
	46:  "SVC_Messages_svc_VoiceInit",
	47:  "SVC_Messages_svc_VoiceData",
	48:  "SVC_Messages_svc_Print",
	49:  "SVC_Messages_svc_Sounds",
	50:  "SVC_Messages_svc_SetView",
	51:  "SVC_Messages_svc_ClearAllStringTables",
	52:  "SVC_Messages_svc_CmdKeyValues",
	53:  "SVC_Messages_svc_BSPDecal",
	54:  "SVC_Messages_svc_SplitScreen",
	55:  "SVC_Messages_svc_PacketEntities",
	56:  "SVC_Messages_svc_Prefetch",
	57:  "SVC_Messages_svc_Menu",
	58:  "SVC_Messages_svc_GetCvarValue",
	59:  "SVC_Messages_svc_StopSound",
	60:  "SVC_Messages_svc_PeerList",
	61:  "SVC_Messages_svc_PacketReliable",
	62:  "SVC_Messages_svc_UserMessage",
	63:  "SVC_Messages_svc_SendTable",
	67:  "SVC_Messages_svc_GameEvent",
	68:  "SVC_Messages_svc_TempEntities",
	69:  "SVC_Messages_svc_GameEventList",
	70:  "SVC_Messages_svc_FullFrameSplit",
	101: "EBaseUserMessages_UM_AchievementEvent",
	102: "EBaseUserMessages_UM_CloseCaption",
	103: "EBaseUserMessages_UM_CloseCaptionDirect",
	104: "EBaseUserMessages_UM_CurrentTimescale",
	105: "EBaseUserMessages_UM_DesiredTimescale",
	106: "EBaseUserMessages_UM_Fade",
	107: "EBaseUserMessages_UM_GameTitle",
	109: "EBaseUserMessages_UM_HintText",
	110: "EBaseUserMessages_UM_HudMsg",
	111: "EBaseUserMessages_UM_HudText",
	112: "EBaseUserMessages_UM_KeyHintText",
	113: "EBaseUserMessages_UM_ColoredText",
	114: "EBaseUserMessages_UM_RequestState",
	115: "EBaseUserMessages_UM_ResetHUD",
	116: "EBaseUserMessages_UM_Rumble",
	117: "EBaseUserMessages_UM_SayText",
	118: "EBaseUserMessages_UM_SayText2",
	119: "EBaseUserMessages_UM_SayTextChannel",
	120: "EBaseUserMessages_UM_Shake",
	121: "EBaseUserMessages_UM_ShakeDir",
	124: "EBaseUserMessages_UM_TextMsg",
	125: "EBaseUserMessages_UM_ScreenTilt",
	126: "EBaseUserMessages_UM_Train",
	127: "EBaseUserMessages_UM_VGUIMenu",
	128: "EBaseUserMessages_UM_VoiceMask",
	129: "EBaseUserMessages_UM_VoiceSubtitle",
	130: "EBaseUserMessages_UM_SendAudio",
	131: "EBaseUserMessages_UM_ItemPickup",
	132: "EBaseUserMessages_UM_AmmoDenied",
	133: "EBaseUserMessages_UM_CrosshairAngle",
	134: "EBaseUserMessages_UM_ShowMenu",
	135: "EBaseUserMessages_UM_CreditsMsg",
	136: "EBaseEntityMessages_EM_PlayJingle",
	137: "EBaseEntityMessages_EM_ScreenOverlay",
	138: "EBaseEntityMessages_EM_RemoveAllDecals",
	139: "EBaseEntityMessages_EM_PropagateForce",
	140: "EBaseEntityMessages_EM_DoSpark",
	141: "EBaseEntityMessages_EM_FixAngle",
	142: "EBaseUserMessages_UM_CloseCaptionPlaceholder",
	143: "EBaseUserMessages_UM_CameraTransition",
	144: "EBaseUserMessages_UM_AudioParameter",
	145: "EBaseUserMessages_UM_ParticleManager",
	146: "EBaseUserMessages_UM_HudError",
	147: "EBaseUserMessages_UM_CustomGameEvent_ClientToServer",
	148: "EBaseUserMessages_UM_CustomGameEvent_ServerToClient",
	149: "EBaseUserMessages_UM_TrackedControllerInput_ClientToServer",
	200: "EBaseGameEvents_GE_VDebugGameSessionIDEvent",
	201: "EBaseGameEvents_GE_PlaceDecalEvent",
	202: "EBaseGameEvents_GE_ClearWorldDecalsEvent",
	203: "EBaseGameEvents_GE_ClearEntityDecalsEvent",
	204: "EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent",
	205: "EBaseGameEvents_GE_Source1LegacyGameEventList",
	206: "EBaseGameEvents_GE_Source1LegacyListenEvents",
	207: "EBaseGameEvents_GE_Source1LegacyGameEvent",
	208: "EBaseGameEvents_GE_SosStartSoundEvent",
	209: "EBaseGameEvents_GE_SosStopSoundEvent",
	210: "EBaseGameEvents_GE_SosSetSoundEventParams",
	211: "EBaseGameEvents_GE_SosSetLibraryStackFields",
	212: "EBaseGameEvents_GE_SosStopSoundEventHash",
	465: "EDotaUserMessages_DOTA_UM_AIDebugLine",
	466: "EDotaUserMessages_DOTA_UM_ChatEvent",
	467: "EDotaUserMessages_DOTA_UM_CombatHeroPositions",
	470: "EDotaUserMessages_DOTA_UM_CombatLogShowDeath",
	471: "EDotaUserMessages_DOTA_UM_CreateLinearProjectile",
	472: "EDotaUserMessages_DOTA_UM_DestroyLinearProjectile",
	473: "EDotaUserMessages_DOTA_UM_DodgeTrackingProjectiles",
	474: "EDotaUserMessages_DOTA_UM_GlobalLightColor",
	475: "EDotaUserMessages_DOTA_UM_GlobalLightDirection",
	476: "EDotaUserMessages_DOTA_UM_InvalidCommand",
	477: "EDotaUserMessages_DOTA_UM_LocationPing",
	478: "EDotaUserMessages_DOTA_UM_MapLine",
	479: "EDotaUserMessages_DOTA_UM_MiniKillCamInfo",
	480: "EDotaUserMessages_DOTA_UM_MinimapDebugPoint",
	481: "EDotaUserMessages_DOTA_UM_MinimapEvent",
	482: "EDotaUserMessages_DOTA_UM_NevermoreRequiem",
	483: "EDotaUserMessages_DOTA_UM_OverheadEvent",
	484: "EDotaUserMessages_DOTA_UM_SetNextAutobuyItem",
	485: "EDotaUserMessages_DOTA_UM_SharedCooldown",
	486: "EDotaUserMessages_DOTA_UM_SpectatorPlayerClick",
	487: "EDotaUserMessages_DOTA_UM_TutorialTipInfo",
	488: "EDotaUserMessages_DOTA_UM_UnitEvent",
	489: "EDotaUserMessages_DOTA_UM_ParticleManager",
	490: "EDotaUserMessages_DOTA_UM_BotChat",
	491: "EDotaUserMessages_DOTA_UM_HudError",
	492: "EDotaUserMessages_DOTA_UM_ItemPurchased",
	493: "EDotaUserMessages_DOTA_UM_Ping",
	494: "EDotaUserMessages_DOTA_UM_ItemFound",
	496: "EDotaUserMessages_DOTA_UM_SwapVerify",
	497: "EDotaUserMessages_DOTA_UM_WorldLine",
	499: "EDotaUserMessages_DOTA_UM_ItemAlert",
	500: "EDotaUserMessages_DOTA_UM_HalloweenDrops",
	501: "EDotaUserMessages_DOTA_UM_ChatWheel",
	502: "EDotaUserMessages_DOTA_UM_ReceivedXmasGift",
	503: "EDotaUserMessages_DOTA_UM_UpdateSharedContent",
	504: "EDotaUserMessages_DOTA_UM_TutorialRequestExp",
	505: "EDotaUserMessages_DOTA_UM_TutorialPingMinimap",
	506: "EDotaUserMessages_DOTA_UM_GamerulesStateChanged",
	507: "EDotaUserMessages_DOTA_UM_ShowSurvey",
	508: "EDotaUserMessages_DOTA_UM_TutorialFade",
	509: "EDotaUserMessages_DOTA_UM_AddQuestLogEntry",
	510: "EDotaUserMessages_DOTA_UM_SendStatPopup",
	511: "EDotaUserMessages_DOTA_UM_TutorialFinish",
	512: "EDotaUserMessages_DOTA_UM_SendRoshanPopup",
	513: "EDotaUserMessages_DOTA_UM_SendGenericToolTip",
	514: "EDotaUserMessages_DOTA_UM_SendFinalGold",
	515: "EDotaUserMessages_DOTA_UM_CustomMsg",
	516: "EDotaUserMessages_DOTA_UM_CoachHUDPing",
	517: "EDotaUserMessages_DOTA_UM_ClientLoadGridNav",
	518: "EDotaUserMessages_DOTA_UM_TE_Projectile",
	519: "EDotaUserMessages_DOTA_UM_TE_ProjectileLoc",
	520: "EDotaUserMessages_DOTA_UM_TE_DotaBloodImpact",
	521: "EDotaUserMessages_DOTA_UM_TE_UnitAnimation",
	522: "EDotaUserMessages_DOTA_UM_TE_UnitAnimationEnd",
	523: "EDotaUserMessages_DOTA_UM_AbilityPing",
	524: "EDotaUserMessages_DOTA_UM_ShowGenericPopup",
	525: "EDotaUserMessages_DOTA_UM_VoteStart",
	526: "EDotaUserMessages_DOTA_UM_VoteUpdate",
	527: "EDotaUserMessages_DOTA_UM_VoteEnd",
	528: "EDotaUserMessages_DOTA_UM_BoosterState",
	529: "EDotaUserMessages_DOTA_UM_WillPurchaseAlert",
	530: "EDotaUserMessages_DOTA_UM_TutorialMinimapPosition",
	531: "EDotaUserMessages_DOTA_UM_PlayerMMR",
	532: "EDotaUserMessages_DOTA_UM_AbilitySteal",
	533: "EDotaUserMessages_DOTA_UM_CourierKilledAlert",
	534: "EDotaUserMessages_DOTA_UM_EnemyItemAlert",
	535: "EDotaUserMessages_DOTA_UM_StatsMatchDetails",
	536: "EDotaUserMessages_DOTA_UM_MiniTaunt",
	537: "EDotaUserMessages_DOTA_UM_BuyBackStateAlert",
	538: "EDotaUserMessages_DOTA_UM_SpeechBubble",
	539: "EDotaUserMessages_DOTA_UM_CustomHeaderMessage",
	540: "EDotaUserMessages_DOTA_UM_QuickBuyAlert",
	541: "EDotaUserMessages_DOTA_UM_StatsHeroDetails",
	542: "EDotaUserMessages_DOTA_UM_PredictionResult",
	543: "EDotaUserMessages_DOTA_UM_ModifierAlert",
	544: "EDotaUserMessages_DOTA_UM_HPManaAlert",
	545: "EDotaUserMessages_DOTA_UM_GlyphAlert",
	546: "EDotaUserMessages_DOTA_UM_BeastChat",
	547: "EDotaUserMessages_DOTA_UM_SpectatorPlayerUnitOrders",
	548: "EDotaUserMessages_DOTA_UM_CustomHudElement_Create",
	549: "EDotaUserMessages_DOTA_UM_CustomHudElement_Modify",
	550: "EDotaUserMessages_DOTA_UM_CustomHudElement_Destroy",
	551: "EDotaUserMessages_DOTA_UM_CompendiumState",
}

type Callbacks struct {
	onCDemoStop                               []func(*dota.CDemoStop) error
	onCDemoFileHeader                         []func(*dota.CDemoFileHeader) error
	onCDemoFileInfo                           []func(*dota.CDemoFileInfo) error
	onCDemoSyncTick                           []func(*dota.CDemoSyncTick) error
	onCDemoSendTables                         []func(*dota.CDemoSendTables) error
	onCDemoClassInfo                          []func(*dota.CDemoClassInfo) error
	onCDemoStringTables                       []func(*dota.CDemoStringTables) error
	onCDemoPacket                             []func(*dota.CDemoPacket) error
	onCDemoSignonPacket                       []func(*dota.CDemoPacket) error
	onCDemoConsoleCmd                         []func(*dota.CDemoConsoleCmd) error
	onCDemoCustomData                         []func(*dota.CDemoCustomData) error
	onCDemoCustomDataCallbacks                []func(*dota.CDemoCustomDataCallbacks) error
	onCDemoUserCmd                            []func(*dota.CDemoUserCmd) error
	onCDemoFullPacket                         []func(*dota.CDemoFullPacket) error
	onCDemoSaveGame                           []func(*dota.CDemoSaveGame) error
	onCDemoSpawnGroups                        []func(*dota.CDemoSpawnGroups) error
	onCNETMsg_NOP                             []func(*dota.CNETMsg_NOP) error
	onCNETMsg_Disconnect                      []func(*dota.CNETMsg_Disconnect) error
	onCNETMsg_File                            []func(*dota.CNETMsg_File) error
	onCNETMsg_SplitScreenUser                 []func(*dota.CNETMsg_SplitScreenUser) error
	onCNETMsg_Tick                            []func(*dota.CNETMsg_Tick) error
	onCNETMsg_StringCmd                       []func(*dota.CNETMsg_StringCmd) error
	onCNETMsg_SetConVar                       []func(*dota.CNETMsg_SetConVar) error
	onCNETMsg_SignonState                     []func(*dota.CNETMsg_SignonState) error
	onCNETMsg_SpawnGroup_Load                 []func(*dota.CNETMsg_SpawnGroup_Load) error
	onCNETMsg_SpawnGroup_ManifestUpdate       []func(*dota.CNETMsg_SpawnGroup_ManifestUpdate) error
	onCNETMsg_SpawnGroup_SetCreationTick      []func(*dota.CNETMsg_SpawnGroup_SetCreationTick) error
	onCNETMsg_SpawnGroup_Unload               []func(*dota.CNETMsg_SpawnGroup_Unload) error
	onCNETMsg_SpawnGroup_LoadCompleted        []func(*dota.CNETMsg_SpawnGroup_LoadCompleted) error
	onCNETMsg_ReliableMessageEndMarker        []func(*dota.CNETMsg_ReliableMessageEndMarker) error
	onCSVCMsg_ServerInfo                      []func(*dota.CSVCMsg_ServerInfo) error
	onCSVCMsg_FlattenedSerializer             []func(*dota.CSVCMsg_FlattenedSerializer) error
	onCSVCMsg_ClassInfo                       []func(*dota.CSVCMsg_ClassInfo) error
	onCSVCMsg_SetPause                        []func(*dota.CSVCMsg_SetPause) error
	onCSVCMsg_CreateStringTable               []func(*dota.CSVCMsg_CreateStringTable) error
	onCSVCMsg_UpdateStringTable               []func(*dota.CSVCMsg_UpdateStringTable) error
	onCSVCMsg_VoiceInit                       []func(*dota.CSVCMsg_VoiceInit) error
	onCSVCMsg_VoiceData                       []func(*dota.CSVCMsg_VoiceData) error
	onCSVCMsg_Print                           []func(*dota.CSVCMsg_Print) error
	onCSVCMsg_Sounds                          []func(*dota.CSVCMsg_Sounds) error
	onCSVCMsg_SetView                         []func(*dota.CSVCMsg_SetView) error
	onCSVCMsg_ClearAllStringTables            []func(*dota.CSVCMsg_ClearAllStringTables) error
	onCSVCMsg_CmdKeyValues                    []func(*dota.CSVCMsg_CmdKeyValues) error
	onCSVCMsg_BSPDecal                        []func(*dota.CSVCMsg_BSPDecal) error
	onCSVCMsg_SplitScreen                     []func(*dota.CSVCMsg_SplitScreen) error
	onCSVCMsg_PacketEntities                  []func(*dota.CSVCMsg_PacketEntities) error
	onCSVCMsg_Prefetch                        []func(*dota.CSVCMsg_Prefetch) error
	onCSVCMsg_Menu                            []func(*dota.CSVCMsg_Menu) error
	onCSVCMsg_GetCvarValue                    []func(*dota.CSVCMsg_GetCvarValue) error
	onCSVCMsg_StopSound                       []func(*dota.CSVCMsg_StopSound) error
	onCSVCMsg_PeerList                        []func(*dota.CSVCMsg_PeerList) error
	onCSVCMsg_PacketReliable                  []func(*dota.CSVCMsg_PacketReliable) error
	onCSVCMsg_UserMessage                     []func(*dota.CSVCMsg_UserMessage) error
	onCSVCMsg_SendTable                       []func(*dota.CSVCMsg_SendTable) error
	onCSVCMsg_GameEvent                       []func(*dota.CSVCMsg_GameEvent) error
	onCSVCMsg_TempEntities                    []func(*dota.CSVCMsg_TempEntities) error
	onCSVCMsg_GameEventList                   []func(*dota.CSVCMsg_GameEventList) error
	onCSVCMsg_FullFrameSplit                  []func(*dota.CSVCMsg_FullFrameSplit) error
	onCUserMessageAchievementEvent            []func(*dota.CUserMessageAchievementEvent) error
	onCUserMessageCloseCaption                []func(*dota.CUserMessageCloseCaption) error
	onCUserMessageCloseCaptionDirect          []func(*dota.CUserMessageCloseCaptionDirect) error
	onCUserMessageCurrentTimescale            []func(*dota.CUserMessageCurrentTimescale) error
	onCUserMessageDesiredTimescale            []func(*dota.CUserMessageDesiredTimescale) error
	onCUserMessageFade                        []func(*dota.CUserMessageFade) error
	onCUserMessageGameTitle                   []func(*dota.CUserMessageGameTitle) error
	onCUserMessageHintText                    []func(*dota.CUserMessageHintText) error
	onCUserMessageHudMsg                      []func(*dota.CUserMessageHudMsg) error
	onCUserMessageHudText                     []func(*dota.CUserMessageHudText) error
	onCUserMessageKeyHintText                 []func(*dota.CUserMessageKeyHintText) error
	onCUserMessageColoredText                 []func(*dota.CUserMessageColoredText) error
	onCUserMessageRequestState                []func(*dota.CUserMessageRequestState) error
	onCUserMessageResetHUD                    []func(*dota.CUserMessageResetHUD) error
	onCUserMessageRumble                      []func(*dota.CUserMessageRumble) error
	onCUserMessageSayText                     []func(*dota.CUserMessageSayText) error
	onCUserMessageSayText2                    []func(*dota.CUserMessageSayText2) error
	onCUserMessageSayTextChannel              []func(*dota.CUserMessageSayTextChannel) error
	onCUserMessageShake                       []func(*dota.CUserMessageShake) error
	onCUserMessageShakeDir                    []func(*dota.CUserMessageShakeDir) error
	onCUserMessageTextMsg                     []func(*dota.CUserMessageTextMsg) error
	onCUserMessageScreenTilt                  []func(*dota.CUserMessageScreenTilt) error
	onCUserMessageTrain                       []func(*dota.CUserMessageTrain) error
	onCUserMessageVGUIMenu                    []func(*dota.CUserMessageVGUIMenu) error
	onCUserMessageVoiceMask                   []func(*dota.CUserMessageVoiceMask) error
	onCUserMessageVoiceSubtitle               []func(*dota.CUserMessageVoiceSubtitle) error
	onCUserMessageSendAudio                   []func(*dota.CUserMessageSendAudio) error
	onCUserMessageItemPickup                  []func(*dota.CUserMessageItemPickup) error
	onCUserMessageAmmoDenied                  []func(*dota.CUserMessageAmmoDenied) error
	onCUserMessageCrosshairAngle              []func(*dota.CUserMessageCrosshairAngle) error
	onCUserMessageShowMenu                    []func(*dota.CUserMessageShowMenu) error
	onCUserMessageCreditsMsg                  []func(*dota.CUserMessageCreditsMsg) error
	onCUserMessageCloseCaptionPlaceholder     []func(*dota.CUserMessageCloseCaptionPlaceholder) error
	onCUserMessageCameraTransition            []func(*dota.CUserMessageCameraTransition) error
	onCUserMessageAudioParameter              []func(*dota.CUserMessageAudioParameter) error
	onCEntityMessagePlayJingle                []func(*dota.CEntityMessagePlayJingle) error
	onCEntityMessageScreenOverlay             []func(*dota.CEntityMessageScreenOverlay) error
	onCEntityMessageRemoveAllDecals           []func(*dota.CEntityMessageRemoveAllDecals) error
	onCEntityMessagePropagateForce            []func(*dota.CEntityMessagePropagateForce) error
	onCEntityMessageDoSpark                   []func(*dota.CEntityMessageDoSpark) error
	onCEntityMessageFixAngle                  []func(*dota.CEntityMessageFixAngle) error
	onCMsgVDebugGameSessionIDEvent            []func(*dota.CMsgVDebugGameSessionIDEvent) error
	onCMsgPlaceDecalEvent                     []func(*dota.CMsgPlaceDecalEvent) error
	onCMsgClearWorldDecalsEvent               []func(*dota.CMsgClearWorldDecalsEvent) error
	onCMsgClearEntityDecalsEvent              []func(*dota.CMsgClearEntityDecalsEvent) error
	onCMsgClearDecalsForSkeletonInstanceEvent []func(*dota.CMsgClearDecalsForSkeletonInstanceEvent) error
	onCMsgSource1LegacyGameEventList          []func(*dota.CMsgSource1LegacyGameEventList) error
	onCMsgSource1LegacyListenEvents           []func(*dota.CMsgSource1LegacyListenEvents) error
	onCMsgSource1LegacyGameEvent              []func(*dota.CMsgSource1LegacyGameEvent) error
	onCMsgSosStartSoundEvent                  []func(*dota.CMsgSosStartSoundEvent) error
	onCMsgSosStopSoundEvent                   []func(*dota.CMsgSosStopSoundEvent) error
	onCMsgSosSetSoundEventParams              []func(*dota.CMsgSosSetSoundEventParams) error
	onCMsgSosSetLibraryStackFields            []func(*dota.CMsgSosSetLibraryStackFields) error
	onCMsgSosStopSoundEventHash               []func(*dota.CMsgSosStopSoundEventHash) error
	onCDOTAUserMsg_AIDebugLine                []func(*dota.CDOTAUserMsg_AIDebugLine) error
	onCDOTAUserMsg_ChatEvent                  []func(*dota.CDOTAUserMsg_ChatEvent) error
	onCDOTAUserMsg_CombatHeroPositions        []func(*dota.CDOTAUserMsg_CombatHeroPositions) error
	onCDOTAUserMsg_CombatLogShowDeath         []func(*dota.CDOTAUserMsg_CombatLogShowDeath) error
	onCDOTAUserMsg_CreateLinearProjectile     []func(*dota.CDOTAUserMsg_CreateLinearProjectile) error
	onCDOTAUserMsg_DestroyLinearProjectile    []func(*dota.CDOTAUserMsg_DestroyLinearProjectile) error
	onCDOTAUserMsg_DodgeTrackingProjectiles   []func(*dota.CDOTAUserMsg_DodgeTrackingProjectiles) error
	onCDOTAUserMsg_GlobalLightColor           []func(*dota.CDOTAUserMsg_GlobalLightColor) error
	onCDOTAUserMsg_GlobalLightDirection       []func(*dota.CDOTAUserMsg_GlobalLightDirection) error
	onCDOTAUserMsg_InvalidCommand             []func(*dota.CDOTAUserMsg_InvalidCommand) error
	onCDOTAUserMsg_LocationPing               []func(*dota.CDOTAUserMsg_LocationPing) error
	onCDOTAUserMsg_MapLine                    []func(*dota.CDOTAUserMsg_MapLine) error
	onCDOTAUserMsg_MiniKillCamInfo            []func(*dota.CDOTAUserMsg_MiniKillCamInfo) error
	onCDOTAUserMsg_MinimapDebugPoint          []func(*dota.CDOTAUserMsg_MinimapDebugPoint) error
	onCDOTAUserMsg_MinimapEvent               []func(*dota.CDOTAUserMsg_MinimapEvent) error
	onCDOTAUserMsg_NevermoreRequiem           []func(*dota.CDOTAUserMsg_NevermoreRequiem) error
	onCDOTAUserMsg_OverheadEvent              []func(*dota.CDOTAUserMsg_OverheadEvent) error
	onCDOTAUserMsg_SetNextAutobuyItem         []func(*dota.CDOTAUserMsg_SetNextAutobuyItem) error
	onCDOTAUserMsg_SharedCooldown             []func(*dota.CDOTAUserMsg_SharedCooldown) error
	onCDOTAUserMsg_SpectatorPlayerClick       []func(*dota.CDOTAUserMsg_SpectatorPlayerClick) error
	onCDOTAUserMsg_TutorialTipInfo            []func(*dota.CDOTAUserMsg_TutorialTipInfo) error
	onCDOTAUserMsg_UnitEvent                  []func(*dota.CDOTAUserMsg_UnitEvent) error
	onCDOTAUserMsg_ParticleManager            []func(*dota.CDOTAUserMsg_ParticleManager) error
	onCDOTAUserMsg_BotChat                    []func(*dota.CDOTAUserMsg_BotChat) error
	onCDOTAUserMsg_HudError                   []func(*dota.CDOTAUserMsg_HudError) error
	onCDOTAUserMsg_ItemPurchased              []func(*dota.CDOTAUserMsg_ItemPurchased) error
	onCDOTAUserMsg_Ping                       []func(*dota.CDOTAUserMsg_Ping) error
	onCDOTAUserMsg_ItemFound                  []func(*dota.CDOTAUserMsg_ItemFound) error
	onCDOTAUserMsg_SwapVerify                 []func(*dota.CDOTAUserMsg_SwapVerify) error
	onCDOTAUserMsg_WorldLine                  []func(*dota.CDOTAUserMsg_WorldLine) error
	onCDOTAUserMsg_ItemAlert                  []func(*dota.CDOTAUserMsg_ItemAlert) error
	onCDOTAUserMsg_HalloweenDrops             []func(*dota.CDOTAUserMsg_HalloweenDrops) error
	onCDOTAUserMsg_ChatWheel                  []func(*dota.CDOTAUserMsg_ChatWheel) error
	onCDOTAUserMsg_ReceivedXmasGift           []func(*dota.CDOTAUserMsg_ReceivedXmasGift) error
	onCDOTAUserMsg_UpdateSharedContent        []func(*dota.CDOTAUserMsg_UpdateSharedContent) error
	onCDOTAUserMsg_TutorialRequestExp         []func(*dota.CDOTAUserMsg_TutorialRequestExp) error
	onCDOTAUserMsg_TutorialPingMinimap        []func(*dota.CDOTAUserMsg_TutorialPingMinimap) error
	onCDOTAUserMsg_GamerulesStateChanged      []func(*dota.CDOTAUserMsg_GamerulesStateChanged) error
	onCDOTAUserMsg_ShowSurvey                 []func(*dota.CDOTAUserMsg_ShowSurvey) error
	onCDOTAUserMsg_TutorialFade               []func(*dota.CDOTAUserMsg_TutorialFade) error
	onCDOTAUserMsg_AddQuestLogEntry           []func(*dota.CDOTAUserMsg_AddQuestLogEntry) error
	onCDOTAUserMsg_SendStatPopup              []func(*dota.CDOTAUserMsg_SendStatPopup) error
	onCDOTAUserMsg_TutorialFinish             []func(*dota.CDOTAUserMsg_TutorialFinish) error
	onCDOTAUserMsg_SendRoshanPopup            []func(*dota.CDOTAUserMsg_SendRoshanPopup) error
	onCDOTAUserMsg_SendGenericToolTip         []func(*dota.CDOTAUserMsg_SendGenericToolTip) error
	onCDOTAUserMsg_SendFinalGold              []func(*dota.CDOTAUserMsg_SendFinalGold) error
	onCDOTAUserMsg_CustomMsg                  []func(*dota.CDOTAUserMsg_CustomMsg) error
	onCDOTAUserMsg_CoachHUDPing               []func(*dota.CDOTAUserMsg_CoachHUDPing) error
	onCDOTAUserMsg_ClientLoadGridNav          []func(*dota.CDOTAUserMsg_ClientLoadGridNav) error
	onCDOTAUserMsg_TE_Projectile              []func(*dota.CDOTAUserMsg_TE_Projectile) error
	onCDOTAUserMsg_TE_ProjectileLoc           []func(*dota.CDOTAUserMsg_TE_ProjectileLoc) error
	onCDOTAUserMsg_TE_DotaBloodImpact         []func(*dota.CDOTAUserMsg_TE_DotaBloodImpact) error
	onCDOTAUserMsg_TE_UnitAnimation           []func(*dota.CDOTAUserMsg_TE_UnitAnimation) error
	onCDOTAUserMsg_TE_UnitAnimationEnd        []func(*dota.CDOTAUserMsg_TE_UnitAnimationEnd) error
	onCDOTAUserMsg_AbilityPing                []func(*dota.CDOTAUserMsg_AbilityPing) error
	onCDOTAUserMsg_ShowGenericPopup           []func(*dota.CDOTAUserMsg_ShowGenericPopup) error
	onCDOTAUserMsg_VoteStart                  []func(*dota.CDOTAUserMsg_VoteStart) error
	onCDOTAUserMsg_VoteUpdate                 []func(*dota.CDOTAUserMsg_VoteUpdate) error
	onCDOTAUserMsg_VoteEnd                    []func(*dota.CDOTAUserMsg_VoteEnd) error
	onCDOTAUserMsg_BoosterState               []func(*dota.CDOTAUserMsg_BoosterState) error
	onCDOTAUserMsg_WillPurchaseAlert          []func(*dota.CDOTAUserMsg_WillPurchaseAlert) error
	onCDOTAUserMsg_TutorialMinimapPosition    []func(*dota.CDOTAUserMsg_TutorialMinimapPosition) error
	onCDOTAUserMsg_PlayerMMR                  []func(*dota.CDOTAUserMsg_PlayerMMR) error
	onCDOTAUserMsg_AbilitySteal               []func(*dota.CDOTAUserMsg_AbilitySteal) error
	onCDOTAUserMsg_CourierKilledAlert         []func(*dota.CDOTAUserMsg_CourierKilledAlert) error
	onCDOTAUserMsg_EnemyItemAlert             []func(*dota.CDOTAUserMsg_EnemyItemAlert) error
	onCDOTAUserMsg_StatsMatchDetails          []func(*dota.CDOTAUserMsg_StatsMatchDetails) error
	onCDOTAUserMsg_MiniTaunt                  []func(*dota.CDOTAUserMsg_MiniTaunt) error
	onCDOTAUserMsg_BuyBackStateAlert          []func(*dota.CDOTAUserMsg_BuyBackStateAlert) error
	onCDOTAUserMsg_SpeechBubble               []func(*dota.CDOTAUserMsg_SpeechBubble) error
	onCDOTAUserMsg_CustomHeaderMessage        []func(*dota.CDOTAUserMsg_CustomHeaderMessage) error
	onCDOTAUserMsg_QuickBuyAlert              []func(*dota.CDOTAUserMsg_QuickBuyAlert) error
	onCDOTAUserMsg_PredictionResult           []func(*dota.CDOTAUserMsg_PredictionResult) error
	onCDOTAUserMsg_ModifierAlert              []func(*dota.CDOTAUserMsg_ModifierAlert) error
	onCDOTAUserMsg_HPManaAlert                []func(*dota.CDOTAUserMsg_HPManaAlert) error
	onCDOTAUserMsg_GlyphAlert                 []func(*dota.CDOTAUserMsg_GlyphAlert) error
	onCDOTAUserMsg_BeastChat                  []func(*dota.CDOTAUserMsg_BeastChat) error
	onCDOTAUserMsg_SpectatorPlayerUnitOrders  []func(*dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error
	onCDOTAUserMsg_CustomHudElement_Create    []func(*dota.CDOTAUserMsg_CustomHudElement_Create) error
	onCDOTAUserMsg_CustomHudElement_Modify    []func(*dota.CDOTAUserMsg_CustomHudElement_Modify) error
	onCDOTAUserMsg_CustomHudElement_Destroy   []func(*dota.CDOTAUserMsg_CustomHudElement_Destroy) error
	onCDOTAUserMsg_CompendiumState            []func(*dota.CDOTAUserMsg_CompendiumState) error
}

func (c *Callbacks) OnCDemoStop(fn func(*dota.CDemoStop) error) {
	c.onCDemoStop = append(c.onCDemoStop, fn)
}
func (c *Callbacks) OnCDemoFileHeader(fn func(*dota.CDemoFileHeader) error) {
	c.onCDemoFileHeader = append(c.onCDemoFileHeader, fn)
}
func (c *Callbacks) OnCDemoFileInfo(fn func(*dota.CDemoFileInfo) error) {
	c.onCDemoFileInfo = append(c.onCDemoFileInfo, fn)
}
func (c *Callbacks) OnCDemoSyncTick(fn func(*dota.CDemoSyncTick) error) {
	c.onCDemoSyncTick = append(c.onCDemoSyncTick, fn)
}
func (c *Callbacks) OnCDemoSendTables(fn func(*dota.CDemoSendTables) error) {
	c.onCDemoSendTables = append(c.onCDemoSendTables, fn)
}
func (c *Callbacks) OnCDemoClassInfo(fn func(*dota.CDemoClassInfo) error) {
	c.onCDemoClassInfo = append(c.onCDemoClassInfo, fn)
}
func (c *Callbacks) OnCDemoStringTables(fn func(*dota.CDemoStringTables) error) {
	c.onCDemoStringTables = append(c.onCDemoStringTables, fn)
}
func (c *Callbacks) OnCDemoPacket(fn func(*dota.CDemoPacket) error) {
	c.onCDemoPacket = append(c.onCDemoPacket, fn)
}
func (c *Callbacks) OnCDemoSignonPacket(fn func(*dota.CDemoPacket) error) {
	c.onCDemoSignonPacket = append(c.onCDemoSignonPacket, fn)
}
func (c *Callbacks) OnCDemoConsoleCmd(fn func(*dota.CDemoConsoleCmd) error) {
	c.onCDemoConsoleCmd = append(c.onCDemoConsoleCmd, fn)
}
func (c *Callbacks) OnCDemoCustomData(fn func(*dota.CDemoCustomData) error) {
	c.onCDemoCustomData = append(c.onCDemoCustomData, fn)
}
func (c *Callbacks) OnCDemoCustomDataCallbacks(fn func(*dota.CDemoCustomDataCallbacks) error) {
	c.onCDemoCustomDataCallbacks = append(c.onCDemoCustomDataCallbacks, fn)
}
func (c *Callbacks) OnCDemoUserCmd(fn func(*dota.CDemoUserCmd) error) {
	c.onCDemoUserCmd = append(c.onCDemoUserCmd, fn)
}
func (c *Callbacks) OnCDemoFullPacket(fn func(*dota.CDemoFullPacket) error) {
	c.onCDemoFullPacket = append(c.onCDemoFullPacket, fn)
}
func (c *Callbacks) OnCDemoSaveGame(fn func(*dota.CDemoSaveGame) error) {
	c.onCDemoSaveGame = append(c.onCDemoSaveGame, fn)
}
func (c *Callbacks) OnCDemoSpawnGroups(fn func(*dota.CDemoSpawnGroups) error) {
	c.onCDemoSpawnGroups = append(c.onCDemoSpawnGroups, fn)
}
func (c *Callbacks) OnCNETMsg_NOP(fn func(*dota.CNETMsg_NOP) error) {
	c.onCNETMsg_NOP = append(c.onCNETMsg_NOP, fn)
}
func (c *Callbacks) OnCNETMsg_Disconnect(fn func(*dota.CNETMsg_Disconnect) error) {
	c.onCNETMsg_Disconnect = append(c.onCNETMsg_Disconnect, fn)
}
func (c *Callbacks) OnCNETMsg_File(fn func(*dota.CNETMsg_File) error) {
	c.onCNETMsg_File = append(c.onCNETMsg_File, fn)
}
func (c *Callbacks) OnCNETMsg_SplitScreenUser(fn func(*dota.CNETMsg_SplitScreenUser) error) {
	c.onCNETMsg_SplitScreenUser = append(c.onCNETMsg_SplitScreenUser, fn)
}
func (c *Callbacks) OnCNETMsg_Tick(fn func(*dota.CNETMsg_Tick) error) {
	c.onCNETMsg_Tick = append(c.onCNETMsg_Tick, fn)
}
func (c *Callbacks) OnCNETMsg_StringCmd(fn func(*dota.CNETMsg_StringCmd) error) {
	c.onCNETMsg_StringCmd = append(c.onCNETMsg_StringCmd, fn)
}
func (c *Callbacks) OnCNETMsg_SetConVar(fn func(*dota.CNETMsg_SetConVar) error) {
	c.onCNETMsg_SetConVar = append(c.onCNETMsg_SetConVar, fn)
}
func (c *Callbacks) OnCNETMsg_SignonState(fn func(*dota.CNETMsg_SignonState) error) {
	c.onCNETMsg_SignonState = append(c.onCNETMsg_SignonState, fn)
}
func (c *Callbacks) OnCNETMsg_SpawnGroup_Load(fn func(*dota.CNETMsg_SpawnGroup_Load) error) {
	c.onCNETMsg_SpawnGroup_Load = append(c.onCNETMsg_SpawnGroup_Load, fn)
}
func (c *Callbacks) OnCNETMsg_SpawnGroup_ManifestUpdate(fn func(*dota.CNETMsg_SpawnGroup_ManifestUpdate) error) {
	c.onCNETMsg_SpawnGroup_ManifestUpdate = append(c.onCNETMsg_SpawnGroup_ManifestUpdate, fn)
}
func (c *Callbacks) OnCNETMsg_SpawnGroup_SetCreationTick(fn func(*dota.CNETMsg_SpawnGroup_SetCreationTick) error) {
	c.onCNETMsg_SpawnGroup_SetCreationTick = append(c.onCNETMsg_SpawnGroup_SetCreationTick, fn)
}
func (c *Callbacks) OnCNETMsg_SpawnGroup_Unload(fn func(*dota.CNETMsg_SpawnGroup_Unload) error) {
	c.onCNETMsg_SpawnGroup_Unload = append(c.onCNETMsg_SpawnGroup_Unload, fn)
}
func (c *Callbacks) OnCNETMsg_SpawnGroup_LoadCompleted(fn func(*dota.CNETMsg_SpawnGroup_LoadCompleted) error) {
	c.onCNETMsg_SpawnGroup_LoadCompleted = append(c.onCNETMsg_SpawnGroup_LoadCompleted, fn)
}
func (c *Callbacks) OnCNETMsg_ReliableMessageEndMarker(fn func(*dota.CNETMsg_ReliableMessageEndMarker) error) {
	c.onCNETMsg_ReliableMessageEndMarker = append(c.onCNETMsg_ReliableMessageEndMarker, fn)
}
func (c *Callbacks) OnCSVCMsg_ServerInfo(fn func(*dota.CSVCMsg_ServerInfo) error) {
	c.onCSVCMsg_ServerInfo = append(c.onCSVCMsg_ServerInfo, fn)
}
func (c *Callbacks) OnCSVCMsg_FlattenedSerializer(fn func(*dota.CSVCMsg_FlattenedSerializer) error) {
	c.onCSVCMsg_FlattenedSerializer = append(c.onCSVCMsg_FlattenedSerializer, fn)
}
func (c *Callbacks) OnCSVCMsg_ClassInfo(fn func(*dota.CSVCMsg_ClassInfo) error) {
	c.onCSVCMsg_ClassInfo = append(c.onCSVCMsg_ClassInfo, fn)
}
func (c *Callbacks) OnCSVCMsg_SetPause(fn func(*dota.CSVCMsg_SetPause) error) {
	c.onCSVCMsg_SetPause = append(c.onCSVCMsg_SetPause, fn)
}
func (c *Callbacks) OnCSVCMsg_CreateStringTable(fn func(*dota.CSVCMsg_CreateStringTable) error) {
	c.onCSVCMsg_CreateStringTable = append(c.onCSVCMsg_CreateStringTable, fn)
}
func (c *Callbacks) OnCSVCMsg_UpdateStringTable(fn func(*dota.CSVCMsg_UpdateStringTable) error) {
	c.onCSVCMsg_UpdateStringTable = append(c.onCSVCMsg_UpdateStringTable, fn)
}
func (c *Callbacks) OnCSVCMsg_VoiceInit(fn func(*dota.CSVCMsg_VoiceInit) error) {
	c.onCSVCMsg_VoiceInit = append(c.onCSVCMsg_VoiceInit, fn)
}
func (c *Callbacks) OnCSVCMsg_VoiceData(fn func(*dota.CSVCMsg_VoiceData) error) {
	c.onCSVCMsg_VoiceData = append(c.onCSVCMsg_VoiceData, fn)
}
func (c *Callbacks) OnCSVCMsg_Print(fn func(*dota.CSVCMsg_Print) error) {
	c.onCSVCMsg_Print = append(c.onCSVCMsg_Print, fn)
}
func (c *Callbacks) OnCSVCMsg_Sounds(fn func(*dota.CSVCMsg_Sounds) error) {
	c.onCSVCMsg_Sounds = append(c.onCSVCMsg_Sounds, fn)
}
func (c *Callbacks) OnCSVCMsg_SetView(fn func(*dota.CSVCMsg_SetView) error) {
	c.onCSVCMsg_SetView = append(c.onCSVCMsg_SetView, fn)
}
func (c *Callbacks) OnCSVCMsg_ClearAllStringTables(fn func(*dota.CSVCMsg_ClearAllStringTables) error) {
	c.onCSVCMsg_ClearAllStringTables = append(c.onCSVCMsg_ClearAllStringTables, fn)
}
func (c *Callbacks) OnCSVCMsg_CmdKeyValues(fn func(*dota.CSVCMsg_CmdKeyValues) error) {
	c.onCSVCMsg_CmdKeyValues = append(c.onCSVCMsg_CmdKeyValues, fn)
}
func (c *Callbacks) OnCSVCMsg_BSPDecal(fn func(*dota.CSVCMsg_BSPDecal) error) {
	c.onCSVCMsg_BSPDecal = append(c.onCSVCMsg_BSPDecal, fn)
}
func (c *Callbacks) OnCSVCMsg_SplitScreen(fn func(*dota.CSVCMsg_SplitScreen) error) {
	c.onCSVCMsg_SplitScreen = append(c.onCSVCMsg_SplitScreen, fn)
}
func (c *Callbacks) OnCSVCMsg_PacketEntities(fn func(*dota.CSVCMsg_PacketEntities) error) {
	c.onCSVCMsg_PacketEntities = append(c.onCSVCMsg_PacketEntities, fn)
}
func (c *Callbacks) OnCSVCMsg_Prefetch(fn func(*dota.CSVCMsg_Prefetch) error) {
	c.onCSVCMsg_Prefetch = append(c.onCSVCMsg_Prefetch, fn)
}
func (c *Callbacks) OnCSVCMsg_Menu(fn func(*dota.CSVCMsg_Menu) error) {
	c.onCSVCMsg_Menu = append(c.onCSVCMsg_Menu, fn)
}
func (c *Callbacks) OnCSVCMsg_GetCvarValue(fn func(*dota.CSVCMsg_GetCvarValue) error) {
	c.onCSVCMsg_GetCvarValue = append(c.onCSVCMsg_GetCvarValue, fn)
}
func (c *Callbacks) OnCSVCMsg_StopSound(fn func(*dota.CSVCMsg_StopSound) error) {
	c.onCSVCMsg_StopSound = append(c.onCSVCMsg_StopSound, fn)
}
func (c *Callbacks) OnCSVCMsg_PeerList(fn func(*dota.CSVCMsg_PeerList) error) {
	c.onCSVCMsg_PeerList = append(c.onCSVCMsg_PeerList, fn)
}
func (c *Callbacks) OnCSVCMsg_PacketReliable(fn func(*dota.CSVCMsg_PacketReliable) error) {
	c.onCSVCMsg_PacketReliable = append(c.onCSVCMsg_PacketReliable, fn)
}
func (c *Callbacks) OnCSVCMsg_UserMessage(fn func(*dota.CSVCMsg_UserMessage) error) {
	c.onCSVCMsg_UserMessage = append(c.onCSVCMsg_UserMessage, fn)
}
func (c *Callbacks) OnCSVCMsg_SendTable(fn func(*dota.CSVCMsg_SendTable) error) {
	c.onCSVCMsg_SendTable = append(c.onCSVCMsg_SendTable, fn)
}
func (c *Callbacks) OnCSVCMsg_GameEvent(fn func(*dota.CSVCMsg_GameEvent) error) {
	c.onCSVCMsg_GameEvent = append(c.onCSVCMsg_GameEvent, fn)
}
func (c *Callbacks) OnCSVCMsg_TempEntities(fn func(*dota.CSVCMsg_TempEntities) error) {
	c.onCSVCMsg_TempEntities = append(c.onCSVCMsg_TempEntities, fn)
}
func (c *Callbacks) OnCSVCMsg_GameEventList(fn func(*dota.CSVCMsg_GameEventList) error) {
	c.onCSVCMsg_GameEventList = append(c.onCSVCMsg_GameEventList, fn)
}
func (c *Callbacks) OnCSVCMsg_FullFrameSplit(fn func(*dota.CSVCMsg_FullFrameSplit) error) {
	c.onCSVCMsg_FullFrameSplit = append(c.onCSVCMsg_FullFrameSplit, fn)
}
func (c *Callbacks) OnCUserMessageAchievementEvent(fn func(*dota.CUserMessageAchievementEvent) error) {
	c.onCUserMessageAchievementEvent = append(c.onCUserMessageAchievementEvent, fn)
}
func (c *Callbacks) OnCUserMessageCloseCaption(fn func(*dota.CUserMessageCloseCaption) error) {
	c.onCUserMessageCloseCaption = append(c.onCUserMessageCloseCaption, fn)
}
func (c *Callbacks) OnCUserMessageCloseCaptionDirect(fn func(*dota.CUserMessageCloseCaptionDirect) error) {
	c.onCUserMessageCloseCaptionDirect = append(c.onCUserMessageCloseCaptionDirect, fn)
}
func (c *Callbacks) OnCUserMessageCurrentTimescale(fn func(*dota.CUserMessageCurrentTimescale) error) {
	c.onCUserMessageCurrentTimescale = append(c.onCUserMessageCurrentTimescale, fn)
}
func (c *Callbacks) OnCUserMessageDesiredTimescale(fn func(*dota.CUserMessageDesiredTimescale) error) {
	c.onCUserMessageDesiredTimescale = append(c.onCUserMessageDesiredTimescale, fn)
}
func (c *Callbacks) OnCUserMessageFade(fn func(*dota.CUserMessageFade) error) {
	c.onCUserMessageFade = append(c.onCUserMessageFade, fn)
}
func (c *Callbacks) OnCUserMessageGameTitle(fn func(*dota.CUserMessageGameTitle) error) {
	c.onCUserMessageGameTitle = append(c.onCUserMessageGameTitle, fn)
}
func (c *Callbacks) OnCUserMessageHintText(fn func(*dota.CUserMessageHintText) error) {
	c.onCUserMessageHintText = append(c.onCUserMessageHintText, fn)
}
func (c *Callbacks) OnCUserMessageHudMsg(fn func(*dota.CUserMessageHudMsg) error) {
	c.onCUserMessageHudMsg = append(c.onCUserMessageHudMsg, fn)
}
func (c *Callbacks) OnCUserMessageHudText(fn func(*dota.CUserMessageHudText) error) {
	c.onCUserMessageHudText = append(c.onCUserMessageHudText, fn)
}
func (c *Callbacks) OnCUserMessageKeyHintText(fn func(*dota.CUserMessageKeyHintText) error) {
	c.onCUserMessageKeyHintText = append(c.onCUserMessageKeyHintText, fn)
}
func (c *Callbacks) OnCUserMessageColoredText(fn func(*dota.CUserMessageColoredText) error) {
	c.onCUserMessageColoredText = append(c.onCUserMessageColoredText, fn)
}
func (c *Callbacks) OnCUserMessageRequestState(fn func(*dota.CUserMessageRequestState) error) {
	c.onCUserMessageRequestState = append(c.onCUserMessageRequestState, fn)
}
func (c *Callbacks) OnCUserMessageResetHUD(fn func(*dota.CUserMessageResetHUD) error) {
	c.onCUserMessageResetHUD = append(c.onCUserMessageResetHUD, fn)
}
func (c *Callbacks) OnCUserMessageRumble(fn func(*dota.CUserMessageRumble) error) {
	c.onCUserMessageRumble = append(c.onCUserMessageRumble, fn)
}
func (c *Callbacks) OnCUserMessageSayText(fn func(*dota.CUserMessageSayText) error) {
	c.onCUserMessageSayText = append(c.onCUserMessageSayText, fn)
}
func (c *Callbacks) OnCUserMessageSayText2(fn func(*dota.CUserMessageSayText2) error) {
	c.onCUserMessageSayText2 = append(c.onCUserMessageSayText2, fn)
}
func (c *Callbacks) OnCUserMessageSayTextChannel(fn func(*dota.CUserMessageSayTextChannel) error) {
	c.onCUserMessageSayTextChannel = append(c.onCUserMessageSayTextChannel, fn)
}
func (c *Callbacks) OnCUserMessageShake(fn func(*dota.CUserMessageShake) error) {
	c.onCUserMessageShake = append(c.onCUserMessageShake, fn)
}
func (c *Callbacks) OnCUserMessageShakeDir(fn func(*dota.CUserMessageShakeDir) error) {
	c.onCUserMessageShakeDir = append(c.onCUserMessageShakeDir, fn)
}
func (c *Callbacks) OnCUserMessageTextMsg(fn func(*dota.CUserMessageTextMsg) error) {
	c.onCUserMessageTextMsg = append(c.onCUserMessageTextMsg, fn)
}
func (c *Callbacks) OnCUserMessageScreenTilt(fn func(*dota.CUserMessageScreenTilt) error) {
	c.onCUserMessageScreenTilt = append(c.onCUserMessageScreenTilt, fn)
}
func (c *Callbacks) OnCUserMessageTrain(fn func(*dota.CUserMessageTrain) error) {
	c.onCUserMessageTrain = append(c.onCUserMessageTrain, fn)
}
func (c *Callbacks) OnCUserMessageVGUIMenu(fn func(*dota.CUserMessageVGUIMenu) error) {
	c.onCUserMessageVGUIMenu = append(c.onCUserMessageVGUIMenu, fn)
}
func (c *Callbacks) OnCUserMessageVoiceMask(fn func(*dota.CUserMessageVoiceMask) error) {
	c.onCUserMessageVoiceMask = append(c.onCUserMessageVoiceMask, fn)
}
func (c *Callbacks) OnCUserMessageVoiceSubtitle(fn func(*dota.CUserMessageVoiceSubtitle) error) {
	c.onCUserMessageVoiceSubtitle = append(c.onCUserMessageVoiceSubtitle, fn)
}
func (c *Callbacks) OnCUserMessageSendAudio(fn func(*dota.CUserMessageSendAudio) error) {
	c.onCUserMessageSendAudio = append(c.onCUserMessageSendAudio, fn)
}
func (c *Callbacks) OnCUserMessageItemPickup(fn func(*dota.CUserMessageItemPickup) error) {
	c.onCUserMessageItemPickup = append(c.onCUserMessageItemPickup, fn)
}
func (c *Callbacks) OnCUserMessageAmmoDenied(fn func(*dota.CUserMessageAmmoDenied) error) {
	c.onCUserMessageAmmoDenied = append(c.onCUserMessageAmmoDenied, fn)
}
func (c *Callbacks) OnCUserMessageCrosshairAngle(fn func(*dota.CUserMessageCrosshairAngle) error) {
	c.onCUserMessageCrosshairAngle = append(c.onCUserMessageCrosshairAngle, fn)
}
func (c *Callbacks) OnCUserMessageShowMenu(fn func(*dota.CUserMessageShowMenu) error) {
	c.onCUserMessageShowMenu = append(c.onCUserMessageShowMenu, fn)
}
func (c *Callbacks) OnCUserMessageCreditsMsg(fn func(*dota.CUserMessageCreditsMsg) error) {
	c.onCUserMessageCreditsMsg = append(c.onCUserMessageCreditsMsg, fn)
}
func (c *Callbacks) OnCUserMessageCloseCaptionPlaceholder(fn func(*dota.CUserMessageCloseCaptionPlaceholder) error) {
	c.onCUserMessageCloseCaptionPlaceholder = append(c.onCUserMessageCloseCaptionPlaceholder, fn)
}
func (c *Callbacks) OnCUserMessageCameraTransition(fn func(*dota.CUserMessageCameraTransition) error) {
	c.onCUserMessageCameraTransition = append(c.onCUserMessageCameraTransition, fn)
}
func (c *Callbacks) OnCUserMessageAudioParameter(fn func(*dota.CUserMessageAudioParameter) error) {
	c.onCUserMessageAudioParameter = append(c.onCUserMessageAudioParameter, fn)
}
func (c *Callbacks) OnCEntityMessagePlayJingle(fn func(*dota.CEntityMessagePlayJingle) error) {
	c.onCEntityMessagePlayJingle = append(c.onCEntityMessagePlayJingle, fn)
}
func (c *Callbacks) OnCEntityMessageScreenOverlay(fn func(*dota.CEntityMessageScreenOverlay) error) {
	c.onCEntityMessageScreenOverlay = append(c.onCEntityMessageScreenOverlay, fn)
}
func (c *Callbacks) OnCEntityMessageRemoveAllDecals(fn func(*dota.CEntityMessageRemoveAllDecals) error) {
	c.onCEntityMessageRemoveAllDecals = append(c.onCEntityMessageRemoveAllDecals, fn)
}
func (c *Callbacks) OnCEntityMessagePropagateForce(fn func(*dota.CEntityMessagePropagateForce) error) {
	c.onCEntityMessagePropagateForce = append(c.onCEntityMessagePropagateForce, fn)
}
func (c *Callbacks) OnCEntityMessageDoSpark(fn func(*dota.CEntityMessageDoSpark) error) {
	c.onCEntityMessageDoSpark = append(c.onCEntityMessageDoSpark, fn)
}
func (c *Callbacks) OnCEntityMessageFixAngle(fn func(*dota.CEntityMessageFixAngle) error) {
	c.onCEntityMessageFixAngle = append(c.onCEntityMessageFixAngle, fn)
}
func (c *Callbacks) OnCMsgVDebugGameSessionIDEvent(fn func(*dota.CMsgVDebugGameSessionIDEvent) error) {
	c.onCMsgVDebugGameSessionIDEvent = append(c.onCMsgVDebugGameSessionIDEvent, fn)
}
func (c *Callbacks) OnCMsgPlaceDecalEvent(fn func(*dota.CMsgPlaceDecalEvent) error) {
	c.onCMsgPlaceDecalEvent = append(c.onCMsgPlaceDecalEvent, fn)
}
func (c *Callbacks) OnCMsgClearWorldDecalsEvent(fn func(*dota.CMsgClearWorldDecalsEvent) error) {
	c.onCMsgClearWorldDecalsEvent = append(c.onCMsgClearWorldDecalsEvent, fn)
}
func (c *Callbacks) OnCMsgClearEntityDecalsEvent(fn func(*dota.CMsgClearEntityDecalsEvent) error) {
	c.onCMsgClearEntityDecalsEvent = append(c.onCMsgClearEntityDecalsEvent, fn)
}
func (c *Callbacks) OnCMsgClearDecalsForSkeletonInstanceEvent(fn func(*dota.CMsgClearDecalsForSkeletonInstanceEvent) error) {
	c.onCMsgClearDecalsForSkeletonInstanceEvent = append(c.onCMsgClearDecalsForSkeletonInstanceEvent, fn)
}
func (c *Callbacks) OnCMsgSource1LegacyGameEventList(fn func(*dota.CMsgSource1LegacyGameEventList) error) {
	c.onCMsgSource1LegacyGameEventList = append(c.onCMsgSource1LegacyGameEventList, fn)
}
func (c *Callbacks) OnCMsgSource1LegacyListenEvents(fn func(*dota.CMsgSource1LegacyListenEvents) error) {
	c.onCMsgSource1LegacyListenEvents = append(c.onCMsgSource1LegacyListenEvents, fn)
}
func (c *Callbacks) OnCMsgSource1LegacyGameEvent(fn func(*dota.CMsgSource1LegacyGameEvent) error) {
	c.onCMsgSource1LegacyGameEvent = append(c.onCMsgSource1LegacyGameEvent, fn)
}
func (c *Callbacks) OnCMsgSosStartSoundEvent(fn func(*dota.CMsgSosStartSoundEvent) error) {
	c.onCMsgSosStartSoundEvent = append(c.onCMsgSosStartSoundEvent, fn)
}
func (c *Callbacks) OnCMsgSosStopSoundEvent(fn func(*dota.CMsgSosStopSoundEvent) error) {
	c.onCMsgSosStopSoundEvent = append(c.onCMsgSosStopSoundEvent, fn)
}
func (c *Callbacks) OnCMsgSosSetSoundEventParams(fn func(*dota.CMsgSosSetSoundEventParams) error) {
	c.onCMsgSosSetSoundEventParams = append(c.onCMsgSosSetSoundEventParams, fn)
}
func (c *Callbacks) OnCMsgSosSetLibraryStackFields(fn func(*dota.CMsgSosSetLibraryStackFields) error) {
	c.onCMsgSosSetLibraryStackFields = append(c.onCMsgSosSetLibraryStackFields, fn)
}
func (c *Callbacks) OnCMsgSosStopSoundEventHash(fn func(*dota.CMsgSosStopSoundEventHash) error) {
	c.onCMsgSosStopSoundEventHash = append(c.onCMsgSosStopSoundEventHash, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_AIDebugLine(fn func(*dota.CDOTAUserMsg_AIDebugLine) error) {
	c.onCDOTAUserMsg_AIDebugLine = append(c.onCDOTAUserMsg_AIDebugLine, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ChatEvent(fn func(*dota.CDOTAUserMsg_ChatEvent) error) {
	c.onCDOTAUserMsg_ChatEvent = append(c.onCDOTAUserMsg_ChatEvent, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CombatHeroPositions(fn func(*dota.CDOTAUserMsg_CombatHeroPositions) error) {
	c.onCDOTAUserMsg_CombatHeroPositions = append(c.onCDOTAUserMsg_CombatHeroPositions, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CombatLogShowDeath(fn func(*dota.CDOTAUserMsg_CombatLogShowDeath) error) {
	c.onCDOTAUserMsg_CombatLogShowDeath = append(c.onCDOTAUserMsg_CombatLogShowDeath, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CreateLinearProjectile(fn func(*dota.CDOTAUserMsg_CreateLinearProjectile) error) {
	c.onCDOTAUserMsg_CreateLinearProjectile = append(c.onCDOTAUserMsg_CreateLinearProjectile, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_DestroyLinearProjectile(fn func(*dota.CDOTAUserMsg_DestroyLinearProjectile) error) {
	c.onCDOTAUserMsg_DestroyLinearProjectile = append(c.onCDOTAUserMsg_DestroyLinearProjectile, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_DodgeTrackingProjectiles(fn func(*dota.CDOTAUserMsg_DodgeTrackingProjectiles) error) {
	c.onCDOTAUserMsg_DodgeTrackingProjectiles = append(c.onCDOTAUserMsg_DodgeTrackingProjectiles, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_GlobalLightColor(fn func(*dota.CDOTAUserMsg_GlobalLightColor) error) {
	c.onCDOTAUserMsg_GlobalLightColor = append(c.onCDOTAUserMsg_GlobalLightColor, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_GlobalLightDirection(fn func(*dota.CDOTAUserMsg_GlobalLightDirection) error) {
	c.onCDOTAUserMsg_GlobalLightDirection = append(c.onCDOTAUserMsg_GlobalLightDirection, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_InvalidCommand(fn func(*dota.CDOTAUserMsg_InvalidCommand) error) {
	c.onCDOTAUserMsg_InvalidCommand = append(c.onCDOTAUserMsg_InvalidCommand, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_LocationPing(fn func(*dota.CDOTAUserMsg_LocationPing) error) {
	c.onCDOTAUserMsg_LocationPing = append(c.onCDOTAUserMsg_LocationPing, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_MapLine(fn func(*dota.CDOTAUserMsg_MapLine) error) {
	c.onCDOTAUserMsg_MapLine = append(c.onCDOTAUserMsg_MapLine, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_MiniKillCamInfo(fn func(*dota.CDOTAUserMsg_MiniKillCamInfo) error) {
	c.onCDOTAUserMsg_MiniKillCamInfo = append(c.onCDOTAUserMsg_MiniKillCamInfo, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_MinimapDebugPoint(fn func(*dota.CDOTAUserMsg_MinimapDebugPoint) error) {
	c.onCDOTAUserMsg_MinimapDebugPoint = append(c.onCDOTAUserMsg_MinimapDebugPoint, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_MinimapEvent(fn func(*dota.CDOTAUserMsg_MinimapEvent) error) {
	c.onCDOTAUserMsg_MinimapEvent = append(c.onCDOTAUserMsg_MinimapEvent, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_NevermoreRequiem(fn func(*dota.CDOTAUserMsg_NevermoreRequiem) error) {
	c.onCDOTAUserMsg_NevermoreRequiem = append(c.onCDOTAUserMsg_NevermoreRequiem, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_OverheadEvent(fn func(*dota.CDOTAUserMsg_OverheadEvent) error) {
	c.onCDOTAUserMsg_OverheadEvent = append(c.onCDOTAUserMsg_OverheadEvent, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SetNextAutobuyItem(fn func(*dota.CDOTAUserMsg_SetNextAutobuyItem) error) {
	c.onCDOTAUserMsg_SetNextAutobuyItem = append(c.onCDOTAUserMsg_SetNextAutobuyItem, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SharedCooldown(fn func(*dota.CDOTAUserMsg_SharedCooldown) error) {
	c.onCDOTAUserMsg_SharedCooldown = append(c.onCDOTAUserMsg_SharedCooldown, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SpectatorPlayerClick(fn func(*dota.CDOTAUserMsg_SpectatorPlayerClick) error) {
	c.onCDOTAUserMsg_SpectatorPlayerClick = append(c.onCDOTAUserMsg_SpectatorPlayerClick, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TutorialTipInfo(fn func(*dota.CDOTAUserMsg_TutorialTipInfo) error) {
	c.onCDOTAUserMsg_TutorialTipInfo = append(c.onCDOTAUserMsg_TutorialTipInfo, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_UnitEvent(fn func(*dota.CDOTAUserMsg_UnitEvent) error) {
	c.onCDOTAUserMsg_UnitEvent = append(c.onCDOTAUserMsg_UnitEvent, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ParticleManager(fn func(*dota.CDOTAUserMsg_ParticleManager) error) {
	c.onCDOTAUserMsg_ParticleManager = append(c.onCDOTAUserMsg_ParticleManager, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_BotChat(fn func(*dota.CDOTAUserMsg_BotChat) error) {
	c.onCDOTAUserMsg_BotChat = append(c.onCDOTAUserMsg_BotChat, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_HudError(fn func(*dota.CDOTAUserMsg_HudError) error) {
	c.onCDOTAUserMsg_HudError = append(c.onCDOTAUserMsg_HudError, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ItemPurchased(fn func(*dota.CDOTAUserMsg_ItemPurchased) error) {
	c.onCDOTAUserMsg_ItemPurchased = append(c.onCDOTAUserMsg_ItemPurchased, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_Ping(fn func(*dota.CDOTAUserMsg_Ping) error) {
	c.onCDOTAUserMsg_Ping = append(c.onCDOTAUserMsg_Ping, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ItemFound(fn func(*dota.CDOTAUserMsg_ItemFound) error) {
	c.onCDOTAUserMsg_ItemFound = append(c.onCDOTAUserMsg_ItemFound, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SwapVerify(fn func(*dota.CDOTAUserMsg_SwapVerify) error) {
	c.onCDOTAUserMsg_SwapVerify = append(c.onCDOTAUserMsg_SwapVerify, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_WorldLine(fn func(*dota.CDOTAUserMsg_WorldLine) error) {
	c.onCDOTAUserMsg_WorldLine = append(c.onCDOTAUserMsg_WorldLine, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ItemAlert(fn func(*dota.CDOTAUserMsg_ItemAlert) error) {
	c.onCDOTAUserMsg_ItemAlert = append(c.onCDOTAUserMsg_ItemAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_HalloweenDrops(fn func(*dota.CDOTAUserMsg_HalloweenDrops) error) {
	c.onCDOTAUserMsg_HalloweenDrops = append(c.onCDOTAUserMsg_HalloweenDrops, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ChatWheel(fn func(*dota.CDOTAUserMsg_ChatWheel) error) {
	c.onCDOTAUserMsg_ChatWheel = append(c.onCDOTAUserMsg_ChatWheel, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ReceivedXmasGift(fn func(*dota.CDOTAUserMsg_ReceivedXmasGift) error) {
	c.onCDOTAUserMsg_ReceivedXmasGift = append(c.onCDOTAUserMsg_ReceivedXmasGift, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_UpdateSharedContent(fn func(*dota.CDOTAUserMsg_UpdateSharedContent) error) {
	c.onCDOTAUserMsg_UpdateSharedContent = append(c.onCDOTAUserMsg_UpdateSharedContent, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TutorialRequestExp(fn func(*dota.CDOTAUserMsg_TutorialRequestExp) error) {
	c.onCDOTAUserMsg_TutorialRequestExp = append(c.onCDOTAUserMsg_TutorialRequestExp, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TutorialPingMinimap(fn func(*dota.CDOTAUserMsg_TutorialPingMinimap) error) {
	c.onCDOTAUserMsg_TutorialPingMinimap = append(c.onCDOTAUserMsg_TutorialPingMinimap, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_GamerulesStateChanged(fn func(*dota.CDOTAUserMsg_GamerulesStateChanged) error) {
	c.onCDOTAUserMsg_GamerulesStateChanged = append(c.onCDOTAUserMsg_GamerulesStateChanged, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ShowSurvey(fn func(*dota.CDOTAUserMsg_ShowSurvey) error) {
	c.onCDOTAUserMsg_ShowSurvey = append(c.onCDOTAUserMsg_ShowSurvey, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TutorialFade(fn func(*dota.CDOTAUserMsg_TutorialFade) error) {
	c.onCDOTAUserMsg_TutorialFade = append(c.onCDOTAUserMsg_TutorialFade, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_AddQuestLogEntry(fn func(*dota.CDOTAUserMsg_AddQuestLogEntry) error) {
	c.onCDOTAUserMsg_AddQuestLogEntry = append(c.onCDOTAUserMsg_AddQuestLogEntry, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SendStatPopup(fn func(*dota.CDOTAUserMsg_SendStatPopup) error) {
	c.onCDOTAUserMsg_SendStatPopup = append(c.onCDOTAUserMsg_SendStatPopup, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TutorialFinish(fn func(*dota.CDOTAUserMsg_TutorialFinish) error) {
	c.onCDOTAUserMsg_TutorialFinish = append(c.onCDOTAUserMsg_TutorialFinish, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SendRoshanPopup(fn func(*dota.CDOTAUserMsg_SendRoshanPopup) error) {
	c.onCDOTAUserMsg_SendRoshanPopup = append(c.onCDOTAUserMsg_SendRoshanPopup, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SendGenericToolTip(fn func(*dota.CDOTAUserMsg_SendGenericToolTip) error) {
	c.onCDOTAUserMsg_SendGenericToolTip = append(c.onCDOTAUserMsg_SendGenericToolTip, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SendFinalGold(fn func(*dota.CDOTAUserMsg_SendFinalGold) error) {
	c.onCDOTAUserMsg_SendFinalGold = append(c.onCDOTAUserMsg_SendFinalGold, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CustomMsg(fn func(*dota.CDOTAUserMsg_CustomMsg) error) {
	c.onCDOTAUserMsg_CustomMsg = append(c.onCDOTAUserMsg_CustomMsg, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CoachHUDPing(fn func(*dota.CDOTAUserMsg_CoachHUDPing) error) {
	c.onCDOTAUserMsg_CoachHUDPing = append(c.onCDOTAUserMsg_CoachHUDPing, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ClientLoadGridNav(fn func(*dota.CDOTAUserMsg_ClientLoadGridNav) error) {
	c.onCDOTAUserMsg_ClientLoadGridNav = append(c.onCDOTAUserMsg_ClientLoadGridNav, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TE_Projectile(fn func(*dota.CDOTAUserMsg_TE_Projectile) error) {
	c.onCDOTAUserMsg_TE_Projectile = append(c.onCDOTAUserMsg_TE_Projectile, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TE_ProjectileLoc(fn func(*dota.CDOTAUserMsg_TE_ProjectileLoc) error) {
	c.onCDOTAUserMsg_TE_ProjectileLoc = append(c.onCDOTAUserMsg_TE_ProjectileLoc, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TE_DotaBloodImpact(fn func(*dota.CDOTAUserMsg_TE_DotaBloodImpact) error) {
	c.onCDOTAUserMsg_TE_DotaBloodImpact = append(c.onCDOTAUserMsg_TE_DotaBloodImpact, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TE_UnitAnimation(fn func(*dota.CDOTAUserMsg_TE_UnitAnimation) error) {
	c.onCDOTAUserMsg_TE_UnitAnimation = append(c.onCDOTAUserMsg_TE_UnitAnimation, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TE_UnitAnimationEnd(fn func(*dota.CDOTAUserMsg_TE_UnitAnimationEnd) error) {
	c.onCDOTAUserMsg_TE_UnitAnimationEnd = append(c.onCDOTAUserMsg_TE_UnitAnimationEnd, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_AbilityPing(fn func(*dota.CDOTAUserMsg_AbilityPing) error) {
	c.onCDOTAUserMsg_AbilityPing = append(c.onCDOTAUserMsg_AbilityPing, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ShowGenericPopup(fn func(*dota.CDOTAUserMsg_ShowGenericPopup) error) {
	c.onCDOTAUserMsg_ShowGenericPopup = append(c.onCDOTAUserMsg_ShowGenericPopup, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_VoteStart(fn func(*dota.CDOTAUserMsg_VoteStart) error) {
	c.onCDOTAUserMsg_VoteStart = append(c.onCDOTAUserMsg_VoteStart, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_VoteUpdate(fn func(*dota.CDOTAUserMsg_VoteUpdate) error) {
	c.onCDOTAUserMsg_VoteUpdate = append(c.onCDOTAUserMsg_VoteUpdate, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_VoteEnd(fn func(*dota.CDOTAUserMsg_VoteEnd) error) {
	c.onCDOTAUserMsg_VoteEnd = append(c.onCDOTAUserMsg_VoteEnd, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_BoosterState(fn func(*dota.CDOTAUserMsg_BoosterState) error) {
	c.onCDOTAUserMsg_BoosterState = append(c.onCDOTAUserMsg_BoosterState, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_WillPurchaseAlert(fn func(*dota.CDOTAUserMsg_WillPurchaseAlert) error) {
	c.onCDOTAUserMsg_WillPurchaseAlert = append(c.onCDOTAUserMsg_WillPurchaseAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_TutorialMinimapPosition(fn func(*dota.CDOTAUserMsg_TutorialMinimapPosition) error) {
	c.onCDOTAUserMsg_TutorialMinimapPosition = append(c.onCDOTAUserMsg_TutorialMinimapPosition, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_PlayerMMR(fn func(*dota.CDOTAUserMsg_PlayerMMR) error) {
	c.onCDOTAUserMsg_PlayerMMR = append(c.onCDOTAUserMsg_PlayerMMR, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_AbilitySteal(fn func(*dota.CDOTAUserMsg_AbilitySteal) error) {
	c.onCDOTAUserMsg_AbilitySteal = append(c.onCDOTAUserMsg_AbilitySteal, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CourierKilledAlert(fn func(*dota.CDOTAUserMsg_CourierKilledAlert) error) {
	c.onCDOTAUserMsg_CourierKilledAlert = append(c.onCDOTAUserMsg_CourierKilledAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_EnemyItemAlert(fn func(*dota.CDOTAUserMsg_EnemyItemAlert) error) {
	c.onCDOTAUserMsg_EnemyItemAlert = append(c.onCDOTAUserMsg_EnemyItemAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_StatsMatchDetails(fn func(*dota.CDOTAUserMsg_StatsMatchDetails) error) {
	c.onCDOTAUserMsg_StatsMatchDetails = append(c.onCDOTAUserMsg_StatsMatchDetails, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_MiniTaunt(fn func(*dota.CDOTAUserMsg_MiniTaunt) error) {
	c.onCDOTAUserMsg_MiniTaunt = append(c.onCDOTAUserMsg_MiniTaunt, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_BuyBackStateAlert(fn func(*dota.CDOTAUserMsg_BuyBackStateAlert) error) {
	c.onCDOTAUserMsg_BuyBackStateAlert = append(c.onCDOTAUserMsg_BuyBackStateAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SpeechBubble(fn func(*dota.CDOTAUserMsg_SpeechBubble) error) {
	c.onCDOTAUserMsg_SpeechBubble = append(c.onCDOTAUserMsg_SpeechBubble, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CustomHeaderMessage(fn func(*dota.CDOTAUserMsg_CustomHeaderMessage) error) {
	c.onCDOTAUserMsg_CustomHeaderMessage = append(c.onCDOTAUserMsg_CustomHeaderMessage, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_QuickBuyAlert(fn func(*dota.CDOTAUserMsg_QuickBuyAlert) error) {
	c.onCDOTAUserMsg_QuickBuyAlert = append(c.onCDOTAUserMsg_QuickBuyAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_PredictionResult(fn func(*dota.CDOTAUserMsg_PredictionResult) error) {
	c.onCDOTAUserMsg_PredictionResult = append(c.onCDOTAUserMsg_PredictionResult, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_ModifierAlert(fn func(*dota.CDOTAUserMsg_ModifierAlert) error) {
	c.onCDOTAUserMsg_ModifierAlert = append(c.onCDOTAUserMsg_ModifierAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_HPManaAlert(fn func(*dota.CDOTAUserMsg_HPManaAlert) error) {
	c.onCDOTAUserMsg_HPManaAlert = append(c.onCDOTAUserMsg_HPManaAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_GlyphAlert(fn func(*dota.CDOTAUserMsg_GlyphAlert) error) {
	c.onCDOTAUserMsg_GlyphAlert = append(c.onCDOTAUserMsg_GlyphAlert, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_BeastChat(fn func(*dota.CDOTAUserMsg_BeastChat) error) {
	c.onCDOTAUserMsg_BeastChat = append(c.onCDOTAUserMsg_BeastChat, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_SpectatorPlayerUnitOrders(fn func(*dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error) {
	c.onCDOTAUserMsg_SpectatorPlayerUnitOrders = append(c.onCDOTAUserMsg_SpectatorPlayerUnitOrders, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CustomHudElement_Create(fn func(*dota.CDOTAUserMsg_CustomHudElement_Create) error) {
	c.onCDOTAUserMsg_CustomHudElement_Create = append(c.onCDOTAUserMsg_CustomHudElement_Create, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CustomHudElement_Modify(fn func(*dota.CDOTAUserMsg_CustomHudElement_Modify) error) {
	c.onCDOTAUserMsg_CustomHudElement_Modify = append(c.onCDOTAUserMsg_CustomHudElement_Modify, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CustomHudElement_Destroy(fn func(*dota.CDOTAUserMsg_CustomHudElement_Destroy) error) {
	c.onCDOTAUserMsg_CustomHudElement_Destroy = append(c.onCDOTAUserMsg_CustomHudElement_Destroy, fn)
}
func (c *Callbacks) OnCDOTAUserMsg_CompendiumState(fn func(*dota.CDOTAUserMsg_CompendiumState) error) {
	c.onCDOTAUserMsg_CompendiumState = append(c.onCDOTAUserMsg_CompendiumState, fn)
}
func (p *Parser) CallByDemoType(t int32, raw []byte) error {
	callbacks := p.Callbacks
	switch t {
	case 0: // dota.EDemoCommands_DEM_Stop
		if cbs := callbacks.onCDemoStop; cbs != nil {
			msg := &dota.CDemoStop{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 1: // dota.EDemoCommands_DEM_FileHeader
		if cbs := callbacks.onCDemoFileHeader; cbs != nil {
			msg := &dota.CDemoFileHeader{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 2: // dota.EDemoCommands_DEM_FileInfo
		if cbs := callbacks.onCDemoFileInfo; cbs != nil {
			msg := &dota.CDemoFileInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 3: // dota.EDemoCommands_DEM_SyncTick
		if cbs := callbacks.onCDemoSyncTick; cbs != nil {
			msg := &dota.CDemoSyncTick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 4: // dota.EDemoCommands_DEM_SendTables
		if cbs := callbacks.onCDemoSendTables; cbs != nil {
			msg := &dota.CDemoSendTables{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 5: // dota.EDemoCommands_DEM_ClassInfo
		if cbs := callbacks.onCDemoClassInfo; cbs != nil {
			msg := &dota.CDemoClassInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 6: // dota.EDemoCommands_DEM_StringTables
		if cbs := callbacks.onCDemoStringTables; cbs != nil {
			msg := &dota.CDemoStringTables{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 7: // dota.EDemoCommands_DEM_Packet
		if cbs := callbacks.onCDemoPacket; cbs != nil {
			msg := &dota.CDemoPacket{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 8: // dota.EDemoCommands_DEM_SignonPacket
		if cbs := callbacks.onCDemoSignonPacket; cbs != nil {
			msg := &dota.CDemoPacket{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 9: // dota.EDemoCommands_DEM_ConsoleCmd
		if cbs := callbacks.onCDemoConsoleCmd; cbs != nil {
			msg := &dota.CDemoConsoleCmd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 10: // dota.EDemoCommands_DEM_CustomData
		if cbs := callbacks.onCDemoCustomData; cbs != nil {
			msg := &dota.CDemoCustomData{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 11: // dota.EDemoCommands_DEM_CustomDataCallbacks
		if cbs := callbacks.onCDemoCustomDataCallbacks; cbs != nil {
			msg := &dota.CDemoCustomDataCallbacks{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 12: // dota.EDemoCommands_DEM_UserCmd
		if cbs := callbacks.onCDemoUserCmd; cbs != nil {
			msg := &dota.CDemoUserCmd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 13: // dota.EDemoCommands_DEM_FullPacket
		if cbs := callbacks.onCDemoFullPacket; cbs != nil {
			msg := &dota.CDemoFullPacket{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 14: // dota.EDemoCommands_DEM_SaveGame
		if cbs := callbacks.onCDemoSaveGame; cbs != nil {
			msg := &dota.CDemoSaveGame{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 15: // dota.EDemoCommands_DEM_SpawnGroups
		if cbs := callbacks.onCDemoSpawnGroups; cbs != nil {
			msg := &dota.CDemoSpawnGroups{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return fmt.Errorf("no type found: %d", t)
}

func (p *Parser) CallByPacketType(t int32, raw []byte) error {
	callbacks := p.Callbacks
	switch t {
	case 0: // dota.NET_Messages_net_NOP
		if cbs := callbacks.onCNETMsg_NOP; cbs != nil {
			msg := &dota.CNETMsg_NOP{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 1: // dota.NET_Messages_net_Disconnect
		if cbs := callbacks.onCNETMsg_Disconnect; cbs != nil {
			msg := &dota.CNETMsg_Disconnect{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 2: // dota.NET_Messages_net_File
		if cbs := callbacks.onCNETMsg_File; cbs != nil {
			msg := &dota.CNETMsg_File{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 3: // dota.NET_Messages_net_SplitScreenUser
		if cbs := callbacks.onCNETMsg_SplitScreenUser; cbs != nil {
			msg := &dota.CNETMsg_SplitScreenUser{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 4: // dota.NET_Messages_net_Tick
		if cbs := callbacks.onCNETMsg_Tick; cbs != nil {
			msg := &dota.CNETMsg_Tick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 5: // dota.NET_Messages_net_StringCmd
		if cbs := callbacks.onCNETMsg_StringCmd; cbs != nil {
			msg := &dota.CNETMsg_StringCmd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 6: // dota.NET_Messages_net_SetConVar
		if cbs := callbacks.onCNETMsg_SetConVar; cbs != nil {
			msg := &dota.CNETMsg_SetConVar{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 7: // dota.NET_Messages_net_SignonState
		if cbs := callbacks.onCNETMsg_SignonState; cbs != nil {
			msg := &dota.CNETMsg_SignonState{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 8: // dota.NET_Messages_net_SpawnGroup_Load
		if cbs := callbacks.onCNETMsg_SpawnGroup_Load; cbs != nil {
			msg := &dota.CNETMsg_SpawnGroup_Load{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 9: // dota.NET_Messages_net_SpawnGroup_ManifestUpdate
		if cbs := callbacks.onCNETMsg_SpawnGroup_ManifestUpdate; cbs != nil {
			msg := &dota.CNETMsg_SpawnGroup_ManifestUpdate{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 11: // dota.NET_Messages_net_SpawnGroup_SetCreationTick
		if cbs := callbacks.onCNETMsg_SpawnGroup_SetCreationTick; cbs != nil {
			msg := &dota.CNETMsg_SpawnGroup_SetCreationTick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 12: // dota.NET_Messages_net_SpawnGroup_Unload
		if cbs := callbacks.onCNETMsg_SpawnGroup_Unload; cbs != nil {
			msg := &dota.CNETMsg_SpawnGroup_Unload{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 13: // dota.NET_Messages_net_SpawnGroup_LoadCompleted
		if cbs := callbacks.onCNETMsg_SpawnGroup_LoadCompleted; cbs != nil {
			msg := &dota.CNETMsg_SpawnGroup_LoadCompleted{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 14: // dota.NET_Messages_net_ReliableMessageEndMarker
		if cbs := callbacks.onCNETMsg_ReliableMessageEndMarker; cbs != nil {
			msg := &dota.CNETMsg_ReliableMessageEndMarker{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 40: // dota.SVC_Messages_svc_ServerInfo
		if cbs := callbacks.onCSVCMsg_ServerInfo; cbs != nil {
			msg := &dota.CSVCMsg_ServerInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 41: // dota.SVC_Messages_svc_FlattenedSerializer
		if cbs := callbacks.onCSVCMsg_FlattenedSerializer; cbs != nil {
			msg := &dota.CSVCMsg_FlattenedSerializer{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 42: // dota.SVC_Messages_svc_ClassInfo
		if cbs := callbacks.onCSVCMsg_ClassInfo; cbs != nil {
			msg := &dota.CSVCMsg_ClassInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 43: // dota.SVC_Messages_svc_SetPause
		if cbs := callbacks.onCSVCMsg_SetPause; cbs != nil {
			msg := &dota.CSVCMsg_SetPause{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 44: // dota.SVC_Messages_svc_CreateStringTable
		if cbs := callbacks.onCSVCMsg_CreateStringTable; cbs != nil {
			msg := &dota.CSVCMsg_CreateStringTable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 45: // dota.SVC_Messages_svc_UpdateStringTable
		if cbs := callbacks.onCSVCMsg_UpdateStringTable; cbs != nil {
			msg := &dota.CSVCMsg_UpdateStringTable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 46: // dota.SVC_Messages_svc_VoiceInit
		if cbs := callbacks.onCSVCMsg_VoiceInit; cbs != nil {
			msg := &dota.CSVCMsg_VoiceInit{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 47: // dota.SVC_Messages_svc_VoiceData
		if cbs := callbacks.onCSVCMsg_VoiceData; cbs != nil {
			msg := &dota.CSVCMsg_VoiceData{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 48: // dota.SVC_Messages_svc_Print
		if cbs := callbacks.onCSVCMsg_Print; cbs != nil {
			msg := &dota.CSVCMsg_Print{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 49: // dota.SVC_Messages_svc_Sounds
		if cbs := callbacks.onCSVCMsg_Sounds; cbs != nil {
			msg := &dota.CSVCMsg_Sounds{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 50: // dota.SVC_Messages_svc_SetView
		if cbs := callbacks.onCSVCMsg_SetView; cbs != nil {
			msg := &dota.CSVCMsg_SetView{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 51: // dota.SVC_Messages_svc_ClearAllStringTables
		if cbs := callbacks.onCSVCMsg_ClearAllStringTables; cbs != nil {
			msg := &dota.CSVCMsg_ClearAllStringTables{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 52: // dota.SVC_Messages_svc_CmdKeyValues
		if cbs := callbacks.onCSVCMsg_CmdKeyValues; cbs != nil {
			msg := &dota.CSVCMsg_CmdKeyValues{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 53: // dota.SVC_Messages_svc_BSPDecal
		if cbs := callbacks.onCSVCMsg_BSPDecal; cbs != nil {
			msg := &dota.CSVCMsg_BSPDecal{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 54: // dota.SVC_Messages_svc_SplitScreen
		if cbs := callbacks.onCSVCMsg_SplitScreen; cbs != nil {
			msg := &dota.CSVCMsg_SplitScreen{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 55: // dota.SVC_Messages_svc_PacketEntities
		if cbs := callbacks.onCSVCMsg_PacketEntities; cbs != nil {
			msg := &dota.CSVCMsg_PacketEntities{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 56: // dota.SVC_Messages_svc_Prefetch
		if cbs := callbacks.onCSVCMsg_Prefetch; cbs != nil {
			msg := &dota.CSVCMsg_Prefetch{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 57: // dota.SVC_Messages_svc_Menu
		if cbs := callbacks.onCSVCMsg_Menu; cbs != nil {
			msg := &dota.CSVCMsg_Menu{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 58: // dota.SVC_Messages_svc_GetCvarValue
		if cbs := callbacks.onCSVCMsg_GetCvarValue; cbs != nil {
			msg := &dota.CSVCMsg_GetCvarValue{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 59: // dota.SVC_Messages_svc_StopSound
		if cbs := callbacks.onCSVCMsg_StopSound; cbs != nil {
			msg := &dota.CSVCMsg_StopSound{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 60: // dota.SVC_Messages_svc_PeerList
		if cbs := callbacks.onCSVCMsg_PeerList; cbs != nil {
			msg := &dota.CSVCMsg_PeerList{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 61: // dota.SVC_Messages_svc_PacketReliable
		if cbs := callbacks.onCSVCMsg_PacketReliable; cbs != nil {
			msg := &dota.CSVCMsg_PacketReliable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 62: // dota.SVC_Messages_svc_UserMessage
		if cbs := callbacks.onCSVCMsg_UserMessage; cbs != nil {
			msg := &dota.CSVCMsg_UserMessage{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 63: // dota.SVC_Messages_svc_SendTable
		if cbs := callbacks.onCSVCMsg_SendTable; cbs != nil {
			msg := &dota.CSVCMsg_SendTable{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 67: // dota.SVC_Messages_svc_GameEvent
		if cbs := callbacks.onCSVCMsg_GameEvent; cbs != nil {
			msg := &dota.CSVCMsg_GameEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 68: // dota.SVC_Messages_svc_TempEntities
		if cbs := callbacks.onCSVCMsg_TempEntities; cbs != nil {
			msg := &dota.CSVCMsg_TempEntities{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 69: // dota.SVC_Messages_svc_GameEventList
		if cbs := callbacks.onCSVCMsg_GameEventList; cbs != nil {
			msg := &dota.CSVCMsg_GameEventList{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 70: // dota.SVC_Messages_svc_FullFrameSplit
		if cbs := callbacks.onCSVCMsg_FullFrameSplit; cbs != nil {
			msg := &dota.CSVCMsg_FullFrameSplit{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 101: // dota.EBaseUserMessages_UM_AchievementEvent
		if cbs := callbacks.onCUserMessageAchievementEvent; cbs != nil {
			msg := &dota.CUserMessageAchievementEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 102: // dota.EBaseUserMessages_UM_CloseCaption
		if cbs := callbacks.onCUserMessageCloseCaption; cbs != nil {
			msg := &dota.CUserMessageCloseCaption{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 103: // dota.EBaseUserMessages_UM_CloseCaptionDirect
		if cbs := callbacks.onCUserMessageCloseCaptionDirect; cbs != nil {
			msg := &dota.CUserMessageCloseCaptionDirect{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 104: // dota.EBaseUserMessages_UM_CurrentTimescale
		if cbs := callbacks.onCUserMessageCurrentTimescale; cbs != nil {
			msg := &dota.CUserMessageCurrentTimescale{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 105: // dota.EBaseUserMessages_UM_DesiredTimescale
		if cbs := callbacks.onCUserMessageDesiredTimescale; cbs != nil {
			msg := &dota.CUserMessageDesiredTimescale{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 106: // dota.EBaseUserMessages_UM_Fade
		if cbs := callbacks.onCUserMessageFade; cbs != nil {
			msg := &dota.CUserMessageFade{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 107: // dota.EBaseUserMessages_UM_GameTitle
		if cbs := callbacks.onCUserMessageGameTitle; cbs != nil {
			msg := &dota.CUserMessageGameTitle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 109: // dota.EBaseUserMessages_UM_HintText
		if cbs := callbacks.onCUserMessageHintText; cbs != nil {
			msg := &dota.CUserMessageHintText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 110: // dota.EBaseUserMessages_UM_HudMsg
		if cbs := callbacks.onCUserMessageHudMsg; cbs != nil {
			msg := &dota.CUserMessageHudMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 111: // dota.EBaseUserMessages_UM_HudText
		if cbs := callbacks.onCUserMessageHudText; cbs != nil {
			msg := &dota.CUserMessageHudText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 112: // dota.EBaseUserMessages_UM_KeyHintText
		if cbs := callbacks.onCUserMessageKeyHintText; cbs != nil {
			msg := &dota.CUserMessageKeyHintText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 113: // dota.EBaseUserMessages_UM_ColoredText
		if cbs := callbacks.onCUserMessageColoredText; cbs != nil {
			msg := &dota.CUserMessageColoredText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 114: // dota.EBaseUserMessages_UM_RequestState
		if cbs := callbacks.onCUserMessageRequestState; cbs != nil {
			msg := &dota.CUserMessageRequestState{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 115: // dota.EBaseUserMessages_UM_ResetHUD
		if cbs := callbacks.onCUserMessageResetHUD; cbs != nil {
			msg := &dota.CUserMessageResetHUD{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 116: // dota.EBaseUserMessages_UM_Rumble
		if cbs := callbacks.onCUserMessageRumble; cbs != nil {
			msg := &dota.CUserMessageRumble{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 117: // dota.EBaseUserMessages_UM_SayText
		if cbs := callbacks.onCUserMessageSayText; cbs != nil {
			msg := &dota.CUserMessageSayText{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 118: // dota.EBaseUserMessages_UM_SayText2
		if cbs := callbacks.onCUserMessageSayText2; cbs != nil {
			msg := &dota.CUserMessageSayText2{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 119: // dota.EBaseUserMessages_UM_SayTextChannel
		if cbs := callbacks.onCUserMessageSayTextChannel; cbs != nil {
			msg := &dota.CUserMessageSayTextChannel{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 120: // dota.EBaseUserMessages_UM_Shake
		if cbs := callbacks.onCUserMessageShake; cbs != nil {
			msg := &dota.CUserMessageShake{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 121: // dota.EBaseUserMessages_UM_ShakeDir
		if cbs := callbacks.onCUserMessageShakeDir; cbs != nil {
			msg := &dota.CUserMessageShakeDir{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 124: // dota.EBaseUserMessages_UM_TextMsg
		if cbs := callbacks.onCUserMessageTextMsg; cbs != nil {
			msg := &dota.CUserMessageTextMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 125: // dota.EBaseUserMessages_UM_ScreenTilt
		if cbs := callbacks.onCUserMessageScreenTilt; cbs != nil {
			msg := &dota.CUserMessageScreenTilt{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 126: // dota.EBaseUserMessages_UM_Train
		if cbs := callbacks.onCUserMessageTrain; cbs != nil {
			msg := &dota.CUserMessageTrain{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 127: // dota.EBaseUserMessages_UM_VGUIMenu
		if cbs := callbacks.onCUserMessageVGUIMenu; cbs != nil {
			msg := &dota.CUserMessageVGUIMenu{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 128: // dota.EBaseUserMessages_UM_VoiceMask
		if cbs := callbacks.onCUserMessageVoiceMask; cbs != nil {
			msg := &dota.CUserMessageVoiceMask{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 129: // dota.EBaseUserMessages_UM_VoiceSubtitle
		if cbs := callbacks.onCUserMessageVoiceSubtitle; cbs != nil {
			msg := &dota.CUserMessageVoiceSubtitle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 130: // dota.EBaseUserMessages_UM_SendAudio
		if cbs := callbacks.onCUserMessageSendAudio; cbs != nil {
			msg := &dota.CUserMessageSendAudio{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 131: // dota.EBaseUserMessages_UM_ItemPickup
		if cbs := callbacks.onCUserMessageItemPickup; cbs != nil {
			msg := &dota.CUserMessageItemPickup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 132: // dota.EBaseUserMessages_UM_AmmoDenied
		if cbs := callbacks.onCUserMessageAmmoDenied; cbs != nil {
			msg := &dota.CUserMessageAmmoDenied{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 133: // dota.EBaseUserMessages_UM_CrosshairAngle
		if cbs := callbacks.onCUserMessageCrosshairAngle; cbs != nil {
			msg := &dota.CUserMessageCrosshairAngle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 134: // dota.EBaseUserMessages_UM_ShowMenu
		if cbs := callbacks.onCUserMessageShowMenu; cbs != nil {
			msg := &dota.CUserMessageShowMenu{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 135: // dota.EBaseUserMessages_UM_CreditsMsg
		if cbs := callbacks.onCUserMessageCreditsMsg; cbs != nil {
			msg := &dota.CUserMessageCreditsMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 142: // dota.EBaseUserMessages_UM_CloseCaptionPlaceholder
		if cbs := callbacks.onCUserMessageCloseCaptionPlaceholder; cbs != nil {
			msg := &dota.CUserMessageCloseCaptionPlaceholder{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 143: // dota.EBaseUserMessages_UM_CameraTransition
		if cbs := callbacks.onCUserMessageCameraTransition; cbs != nil {
			msg := &dota.CUserMessageCameraTransition{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 144: // dota.EBaseUserMessages_UM_AudioParameter
		if cbs := callbacks.onCUserMessageAudioParameter; cbs != nil {
			msg := &dota.CUserMessageAudioParameter{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 136: // dota.EBaseEntityMessages_EM_PlayJingle
		if cbs := callbacks.onCEntityMessagePlayJingle; cbs != nil {
			msg := &dota.CEntityMessagePlayJingle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 137: // dota.EBaseEntityMessages_EM_ScreenOverlay
		if cbs := callbacks.onCEntityMessageScreenOverlay; cbs != nil {
			msg := &dota.CEntityMessageScreenOverlay{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 138: // dota.EBaseEntityMessages_EM_RemoveAllDecals
		if cbs := callbacks.onCEntityMessageRemoveAllDecals; cbs != nil {
			msg := &dota.CEntityMessageRemoveAllDecals{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 139: // dota.EBaseEntityMessages_EM_PropagateForce
		if cbs := callbacks.onCEntityMessagePropagateForce; cbs != nil {
			msg := &dota.CEntityMessagePropagateForce{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 140: // dota.EBaseEntityMessages_EM_DoSpark
		if cbs := callbacks.onCEntityMessageDoSpark; cbs != nil {
			msg := &dota.CEntityMessageDoSpark{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 141: // dota.EBaseEntityMessages_EM_FixAngle
		if cbs := callbacks.onCEntityMessageFixAngle; cbs != nil {
			msg := &dota.CEntityMessageFixAngle{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 200: // dota.EBaseGameEvents_GE_VDebugGameSessionIDEvent
		if cbs := callbacks.onCMsgVDebugGameSessionIDEvent; cbs != nil {
			msg := &dota.CMsgVDebugGameSessionIDEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 201: // dota.EBaseGameEvents_GE_PlaceDecalEvent
		if cbs := callbacks.onCMsgPlaceDecalEvent; cbs != nil {
			msg := &dota.CMsgPlaceDecalEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 202: // dota.EBaseGameEvents_GE_ClearWorldDecalsEvent
		if cbs := callbacks.onCMsgClearWorldDecalsEvent; cbs != nil {
			msg := &dota.CMsgClearWorldDecalsEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 203: // dota.EBaseGameEvents_GE_ClearEntityDecalsEvent
		if cbs := callbacks.onCMsgClearEntityDecalsEvent; cbs != nil {
			msg := &dota.CMsgClearEntityDecalsEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 204: // dota.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
		if cbs := callbacks.onCMsgClearDecalsForSkeletonInstanceEvent; cbs != nil {
			msg := &dota.CMsgClearDecalsForSkeletonInstanceEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 205: // dota.EBaseGameEvents_GE_Source1LegacyGameEventList
		if cbs := callbacks.onCMsgSource1LegacyGameEventList; cbs != nil {
			msg := &dota.CMsgSource1LegacyGameEventList{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 206: // dota.EBaseGameEvents_GE_Source1LegacyListenEvents
		if cbs := callbacks.onCMsgSource1LegacyListenEvents; cbs != nil {
			msg := &dota.CMsgSource1LegacyListenEvents{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 207: // dota.EBaseGameEvents_GE_Source1LegacyGameEvent
		if cbs := callbacks.onCMsgSource1LegacyGameEvent; cbs != nil {
			msg := &dota.CMsgSource1LegacyGameEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 208: // dota.EBaseGameEvents_GE_SosStartSoundEvent
		if cbs := callbacks.onCMsgSosStartSoundEvent; cbs != nil {
			msg := &dota.CMsgSosStartSoundEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 209: // dota.EBaseGameEvents_GE_SosStopSoundEvent
		if cbs := callbacks.onCMsgSosStopSoundEvent; cbs != nil {
			msg := &dota.CMsgSosStopSoundEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 210: // dota.EBaseGameEvents_GE_SosSetSoundEventParams
		if cbs := callbacks.onCMsgSosSetSoundEventParams; cbs != nil {
			msg := &dota.CMsgSosSetSoundEventParams{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 211: // dota.EBaseGameEvents_GE_SosSetLibraryStackFields
		if cbs := callbacks.onCMsgSosSetLibraryStackFields; cbs != nil {
			msg := &dota.CMsgSosSetLibraryStackFields{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 212: // dota.EBaseGameEvents_GE_SosStopSoundEventHash
		if cbs := callbacks.onCMsgSosStopSoundEventHash; cbs != nil {
			msg := &dota.CMsgSosStopSoundEventHash{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 465: // dota.EDotaUserMessages_DOTA_UM_AIDebugLine
		if cbs := callbacks.onCDOTAUserMsg_AIDebugLine; cbs != nil {
			msg := &dota.CDOTAUserMsg_AIDebugLine{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 466: // dota.EDotaUserMessages_DOTA_UM_ChatEvent
		if cbs := callbacks.onCDOTAUserMsg_ChatEvent; cbs != nil {
			msg := &dota.CDOTAUserMsg_ChatEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 467: // dota.EDotaUserMessages_DOTA_UM_CombatHeroPositions
		if cbs := callbacks.onCDOTAUserMsg_CombatHeroPositions; cbs != nil {
			msg := &dota.CDOTAUserMsg_CombatHeroPositions{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 470: // dota.EDotaUserMessages_DOTA_UM_CombatLogShowDeath
		if cbs := callbacks.onCDOTAUserMsg_CombatLogShowDeath; cbs != nil {
			msg := &dota.CDOTAUserMsg_CombatLogShowDeath{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 471: // dota.EDotaUserMessages_DOTA_UM_CreateLinearProjectile
		if cbs := callbacks.onCDOTAUserMsg_CreateLinearProjectile; cbs != nil {
			msg := &dota.CDOTAUserMsg_CreateLinearProjectile{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 472: // dota.EDotaUserMessages_DOTA_UM_DestroyLinearProjectile
		if cbs := callbacks.onCDOTAUserMsg_DestroyLinearProjectile; cbs != nil {
			msg := &dota.CDOTAUserMsg_DestroyLinearProjectile{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 473: // dota.EDotaUserMessages_DOTA_UM_DodgeTrackingProjectiles
		if cbs := callbacks.onCDOTAUserMsg_DodgeTrackingProjectiles; cbs != nil {
			msg := &dota.CDOTAUserMsg_DodgeTrackingProjectiles{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 474: // dota.EDotaUserMessages_DOTA_UM_GlobalLightColor
		if cbs := callbacks.onCDOTAUserMsg_GlobalLightColor; cbs != nil {
			msg := &dota.CDOTAUserMsg_GlobalLightColor{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 475: // dota.EDotaUserMessages_DOTA_UM_GlobalLightDirection
		if cbs := callbacks.onCDOTAUserMsg_GlobalLightDirection; cbs != nil {
			msg := &dota.CDOTAUserMsg_GlobalLightDirection{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 476: // dota.EDotaUserMessages_DOTA_UM_InvalidCommand
		if cbs := callbacks.onCDOTAUserMsg_InvalidCommand; cbs != nil {
			msg := &dota.CDOTAUserMsg_InvalidCommand{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 477: // dota.EDotaUserMessages_DOTA_UM_LocationPing
		if cbs := callbacks.onCDOTAUserMsg_LocationPing; cbs != nil {
			msg := &dota.CDOTAUserMsg_LocationPing{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 478: // dota.EDotaUserMessages_DOTA_UM_MapLine
		if cbs := callbacks.onCDOTAUserMsg_MapLine; cbs != nil {
			msg := &dota.CDOTAUserMsg_MapLine{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 479: // dota.EDotaUserMessages_DOTA_UM_MiniKillCamInfo
		if cbs := callbacks.onCDOTAUserMsg_MiniKillCamInfo; cbs != nil {
			msg := &dota.CDOTAUserMsg_MiniKillCamInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 480: // dota.EDotaUserMessages_DOTA_UM_MinimapDebugPoint
		if cbs := callbacks.onCDOTAUserMsg_MinimapDebugPoint; cbs != nil {
			msg := &dota.CDOTAUserMsg_MinimapDebugPoint{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 481: // dota.EDotaUserMessages_DOTA_UM_MinimapEvent
		if cbs := callbacks.onCDOTAUserMsg_MinimapEvent; cbs != nil {
			msg := &dota.CDOTAUserMsg_MinimapEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 482: // dota.EDotaUserMessages_DOTA_UM_NevermoreRequiem
		if cbs := callbacks.onCDOTAUserMsg_NevermoreRequiem; cbs != nil {
			msg := &dota.CDOTAUserMsg_NevermoreRequiem{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 483: // dota.EDotaUserMessages_DOTA_UM_OverheadEvent
		if cbs := callbacks.onCDOTAUserMsg_OverheadEvent; cbs != nil {
			msg := &dota.CDOTAUserMsg_OverheadEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 484: // dota.EDotaUserMessages_DOTA_UM_SetNextAutobuyItem
		if cbs := callbacks.onCDOTAUserMsg_SetNextAutobuyItem; cbs != nil {
			msg := &dota.CDOTAUserMsg_SetNextAutobuyItem{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 485: // dota.EDotaUserMessages_DOTA_UM_SharedCooldown
		if cbs := callbacks.onCDOTAUserMsg_SharedCooldown; cbs != nil {
			msg := &dota.CDOTAUserMsg_SharedCooldown{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 486: // dota.EDotaUserMessages_DOTA_UM_SpectatorPlayerClick
		if cbs := callbacks.onCDOTAUserMsg_SpectatorPlayerClick; cbs != nil {
			msg := &dota.CDOTAUserMsg_SpectatorPlayerClick{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 487: // dota.EDotaUserMessages_DOTA_UM_TutorialTipInfo
		if cbs := callbacks.onCDOTAUserMsg_TutorialTipInfo; cbs != nil {
			msg := &dota.CDOTAUserMsg_TutorialTipInfo{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 488: // dota.EDotaUserMessages_DOTA_UM_UnitEvent
		if cbs := callbacks.onCDOTAUserMsg_UnitEvent; cbs != nil {
			msg := &dota.CDOTAUserMsg_UnitEvent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 489: // dota.EDotaUserMessages_DOTA_UM_ParticleManager
		if cbs := callbacks.onCDOTAUserMsg_ParticleManager; cbs != nil {
			msg := &dota.CDOTAUserMsg_ParticleManager{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 490: // dota.EDotaUserMessages_DOTA_UM_BotChat
		if cbs := callbacks.onCDOTAUserMsg_BotChat; cbs != nil {
			msg := &dota.CDOTAUserMsg_BotChat{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 491: // dota.EDotaUserMessages_DOTA_UM_HudError
		if cbs := callbacks.onCDOTAUserMsg_HudError; cbs != nil {
			msg := &dota.CDOTAUserMsg_HudError{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 492: // dota.EDotaUserMessages_DOTA_UM_ItemPurchased
		if cbs := callbacks.onCDOTAUserMsg_ItemPurchased; cbs != nil {
			msg := &dota.CDOTAUserMsg_ItemPurchased{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 493: // dota.EDotaUserMessages_DOTA_UM_Ping
		if cbs := callbacks.onCDOTAUserMsg_Ping; cbs != nil {
			msg := &dota.CDOTAUserMsg_Ping{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 494: // dota.EDotaUserMessages_DOTA_UM_ItemFound
		if cbs := callbacks.onCDOTAUserMsg_ItemFound; cbs != nil {
			msg := &dota.CDOTAUserMsg_ItemFound{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 496: // dota.EDotaUserMessages_DOTA_UM_SwapVerify
		if cbs := callbacks.onCDOTAUserMsg_SwapVerify; cbs != nil {
			msg := &dota.CDOTAUserMsg_SwapVerify{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 497: // dota.EDotaUserMessages_DOTA_UM_WorldLine
		if cbs := callbacks.onCDOTAUserMsg_WorldLine; cbs != nil {
			msg := &dota.CDOTAUserMsg_WorldLine{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 499: // dota.EDotaUserMessages_DOTA_UM_ItemAlert
		if cbs := callbacks.onCDOTAUserMsg_ItemAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_ItemAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 500: // dota.EDotaUserMessages_DOTA_UM_HalloweenDrops
		if cbs := callbacks.onCDOTAUserMsg_HalloweenDrops; cbs != nil {
			msg := &dota.CDOTAUserMsg_HalloweenDrops{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 501: // dota.EDotaUserMessages_DOTA_UM_ChatWheel
		if cbs := callbacks.onCDOTAUserMsg_ChatWheel; cbs != nil {
			msg := &dota.CDOTAUserMsg_ChatWheel{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 502: // dota.EDotaUserMessages_DOTA_UM_ReceivedXmasGift
		if cbs := callbacks.onCDOTAUserMsg_ReceivedXmasGift; cbs != nil {
			msg := &dota.CDOTAUserMsg_ReceivedXmasGift{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 503: // dota.EDotaUserMessages_DOTA_UM_UpdateSharedContent
		if cbs := callbacks.onCDOTAUserMsg_UpdateSharedContent; cbs != nil {
			msg := &dota.CDOTAUserMsg_UpdateSharedContent{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 504: // dota.EDotaUserMessages_DOTA_UM_TutorialRequestExp
		if cbs := callbacks.onCDOTAUserMsg_TutorialRequestExp; cbs != nil {
			msg := &dota.CDOTAUserMsg_TutorialRequestExp{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 505: // dota.EDotaUserMessages_DOTA_UM_TutorialPingMinimap
		if cbs := callbacks.onCDOTAUserMsg_TutorialPingMinimap; cbs != nil {
			msg := &dota.CDOTAUserMsg_TutorialPingMinimap{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 506: // dota.EDotaUserMessages_DOTA_UM_GamerulesStateChanged
		if cbs := callbacks.onCDOTAUserMsg_GamerulesStateChanged; cbs != nil {
			msg := &dota.CDOTAUserMsg_GamerulesStateChanged{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 507: // dota.EDotaUserMessages_DOTA_UM_ShowSurvey
		if cbs := callbacks.onCDOTAUserMsg_ShowSurvey; cbs != nil {
			msg := &dota.CDOTAUserMsg_ShowSurvey{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 508: // dota.EDotaUserMessages_DOTA_UM_TutorialFade
		if cbs := callbacks.onCDOTAUserMsg_TutorialFade; cbs != nil {
			msg := &dota.CDOTAUserMsg_TutorialFade{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 509: // dota.EDotaUserMessages_DOTA_UM_AddQuestLogEntry
		if cbs := callbacks.onCDOTAUserMsg_AddQuestLogEntry; cbs != nil {
			msg := &dota.CDOTAUserMsg_AddQuestLogEntry{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 510: // dota.EDotaUserMessages_DOTA_UM_SendStatPopup
		if cbs := callbacks.onCDOTAUserMsg_SendStatPopup; cbs != nil {
			msg := &dota.CDOTAUserMsg_SendStatPopup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 511: // dota.EDotaUserMessages_DOTA_UM_TutorialFinish
		if cbs := callbacks.onCDOTAUserMsg_TutorialFinish; cbs != nil {
			msg := &dota.CDOTAUserMsg_TutorialFinish{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 512: // dota.EDotaUserMessages_DOTA_UM_SendRoshanPopup
		if cbs := callbacks.onCDOTAUserMsg_SendRoshanPopup; cbs != nil {
			msg := &dota.CDOTAUserMsg_SendRoshanPopup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 513: // dota.EDotaUserMessages_DOTA_UM_SendGenericToolTip
		if cbs := callbacks.onCDOTAUserMsg_SendGenericToolTip; cbs != nil {
			msg := &dota.CDOTAUserMsg_SendGenericToolTip{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 514: // dota.EDotaUserMessages_DOTA_UM_SendFinalGold
		if cbs := callbacks.onCDOTAUserMsg_SendFinalGold; cbs != nil {
			msg := &dota.CDOTAUserMsg_SendFinalGold{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 515: // dota.EDotaUserMessages_DOTA_UM_CustomMsg
		if cbs := callbacks.onCDOTAUserMsg_CustomMsg; cbs != nil {
			msg := &dota.CDOTAUserMsg_CustomMsg{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 516: // dota.EDotaUserMessages_DOTA_UM_CoachHUDPing
		if cbs := callbacks.onCDOTAUserMsg_CoachHUDPing; cbs != nil {
			msg := &dota.CDOTAUserMsg_CoachHUDPing{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 517: // dota.EDotaUserMessages_DOTA_UM_ClientLoadGridNav
		if cbs := callbacks.onCDOTAUserMsg_ClientLoadGridNav; cbs != nil {
			msg := &dota.CDOTAUserMsg_ClientLoadGridNav{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 518: // dota.EDotaUserMessages_DOTA_UM_TE_Projectile
		if cbs := callbacks.onCDOTAUserMsg_TE_Projectile; cbs != nil {
			msg := &dota.CDOTAUserMsg_TE_Projectile{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 519: // dota.EDotaUserMessages_DOTA_UM_TE_ProjectileLoc
		if cbs := callbacks.onCDOTAUserMsg_TE_ProjectileLoc; cbs != nil {
			msg := &dota.CDOTAUserMsg_TE_ProjectileLoc{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 520: // dota.EDotaUserMessages_DOTA_UM_TE_DotaBloodImpact
		if cbs := callbacks.onCDOTAUserMsg_TE_DotaBloodImpact; cbs != nil {
			msg := &dota.CDOTAUserMsg_TE_DotaBloodImpact{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 521: // dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimation
		if cbs := callbacks.onCDOTAUserMsg_TE_UnitAnimation; cbs != nil {
			msg := &dota.CDOTAUserMsg_TE_UnitAnimation{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 522: // dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimationEnd
		if cbs := callbacks.onCDOTAUserMsg_TE_UnitAnimationEnd; cbs != nil {
			msg := &dota.CDOTAUserMsg_TE_UnitAnimationEnd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 523: // dota.EDotaUserMessages_DOTA_UM_AbilityPing
		if cbs := callbacks.onCDOTAUserMsg_AbilityPing; cbs != nil {
			msg := &dota.CDOTAUserMsg_AbilityPing{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 524: // dota.EDotaUserMessages_DOTA_UM_ShowGenericPopup
		if cbs := callbacks.onCDOTAUserMsg_ShowGenericPopup; cbs != nil {
			msg := &dota.CDOTAUserMsg_ShowGenericPopup{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 525: // dota.EDotaUserMessages_DOTA_UM_VoteStart
		if cbs := callbacks.onCDOTAUserMsg_VoteStart; cbs != nil {
			msg := &dota.CDOTAUserMsg_VoteStart{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 526: // dota.EDotaUserMessages_DOTA_UM_VoteUpdate
		if cbs := callbacks.onCDOTAUserMsg_VoteUpdate; cbs != nil {
			msg := &dota.CDOTAUserMsg_VoteUpdate{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 527: // dota.EDotaUserMessages_DOTA_UM_VoteEnd
		if cbs := callbacks.onCDOTAUserMsg_VoteEnd; cbs != nil {
			msg := &dota.CDOTAUserMsg_VoteEnd{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 528: // dota.EDotaUserMessages_DOTA_UM_BoosterState
		if cbs := callbacks.onCDOTAUserMsg_BoosterState; cbs != nil {
			msg := &dota.CDOTAUserMsg_BoosterState{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 529: // dota.EDotaUserMessages_DOTA_UM_WillPurchaseAlert
		if cbs := callbacks.onCDOTAUserMsg_WillPurchaseAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_WillPurchaseAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 530: // dota.EDotaUserMessages_DOTA_UM_TutorialMinimapPosition
		if cbs := callbacks.onCDOTAUserMsg_TutorialMinimapPosition; cbs != nil {
			msg := &dota.CDOTAUserMsg_TutorialMinimapPosition{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 531: // dota.EDotaUserMessages_DOTA_UM_PlayerMMR
		if cbs := callbacks.onCDOTAUserMsg_PlayerMMR; cbs != nil {
			msg := &dota.CDOTAUserMsg_PlayerMMR{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 532: // dota.EDotaUserMessages_DOTA_UM_AbilitySteal
		if cbs := callbacks.onCDOTAUserMsg_AbilitySteal; cbs != nil {
			msg := &dota.CDOTAUserMsg_AbilitySteal{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 533: // dota.EDotaUserMessages_DOTA_UM_CourierKilledAlert
		if cbs := callbacks.onCDOTAUserMsg_CourierKilledAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_CourierKilledAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 534: // dota.EDotaUserMessages_DOTA_UM_EnemyItemAlert
		if cbs := callbacks.onCDOTAUserMsg_EnemyItemAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_EnemyItemAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 535: // dota.EDotaUserMessages_DOTA_UM_StatsMatchDetails
		if cbs := callbacks.onCDOTAUserMsg_StatsMatchDetails; cbs != nil {
			msg := &dota.CDOTAUserMsg_StatsMatchDetails{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 536: // dota.EDotaUserMessages_DOTA_UM_MiniTaunt
		if cbs := callbacks.onCDOTAUserMsg_MiniTaunt; cbs != nil {
			msg := &dota.CDOTAUserMsg_MiniTaunt{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 537: // dota.EDotaUserMessages_DOTA_UM_BuyBackStateAlert
		if cbs := callbacks.onCDOTAUserMsg_BuyBackStateAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_BuyBackStateAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 538: // dota.EDotaUserMessages_DOTA_UM_SpeechBubble
		if cbs := callbacks.onCDOTAUserMsg_SpeechBubble; cbs != nil {
			msg := &dota.CDOTAUserMsg_SpeechBubble{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 539: // dota.EDotaUserMessages_DOTA_UM_CustomHeaderMessage
		if cbs := callbacks.onCDOTAUserMsg_CustomHeaderMessage; cbs != nil {
			msg := &dota.CDOTAUserMsg_CustomHeaderMessage{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 540: // dota.EDotaUserMessages_DOTA_UM_QuickBuyAlert
		if cbs := callbacks.onCDOTAUserMsg_QuickBuyAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_QuickBuyAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 542: // dota.EDotaUserMessages_DOTA_UM_PredictionResult
		if cbs := callbacks.onCDOTAUserMsg_PredictionResult; cbs != nil {
			msg := &dota.CDOTAUserMsg_PredictionResult{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 543: // dota.EDotaUserMessages_DOTA_UM_ModifierAlert
		if cbs := callbacks.onCDOTAUserMsg_ModifierAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_ModifierAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 544: // dota.EDotaUserMessages_DOTA_UM_HPManaAlert
		if cbs := callbacks.onCDOTAUserMsg_HPManaAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_HPManaAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 545: // dota.EDotaUserMessages_DOTA_UM_GlyphAlert
		if cbs := callbacks.onCDOTAUserMsg_GlyphAlert; cbs != nil {
			msg := &dota.CDOTAUserMsg_GlyphAlert{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 546: // dota.EDotaUserMessages_DOTA_UM_BeastChat
		if cbs := callbacks.onCDOTAUserMsg_BeastChat; cbs != nil {
			msg := &dota.CDOTAUserMsg_BeastChat{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 547: // dota.EDotaUserMessages_DOTA_UM_SpectatorPlayerUnitOrders
		if cbs := callbacks.onCDOTAUserMsg_SpectatorPlayerUnitOrders; cbs != nil {
			msg := &dota.CDOTAUserMsg_SpectatorPlayerUnitOrders{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 548: // dota.EDotaUserMessages_DOTA_UM_CustomHudElement_Create
		if cbs := callbacks.onCDOTAUserMsg_CustomHudElement_Create; cbs != nil {
			msg := &dota.CDOTAUserMsg_CustomHudElement_Create{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 549: // dota.EDotaUserMessages_DOTA_UM_CustomHudElement_Modify
		if cbs := callbacks.onCDOTAUserMsg_CustomHudElement_Modify; cbs != nil {
			msg := &dota.CDOTAUserMsg_CustomHudElement_Modify{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 550: // dota.EDotaUserMessages_DOTA_UM_CustomHudElement_Destroy
		if cbs := callbacks.onCDOTAUserMsg_CustomHudElement_Destroy; cbs != nil {
			msg := &dota.CDOTAUserMsg_CustomHudElement_Destroy{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	case 551: // dota.EDotaUserMessages_DOTA_UM_CompendiumState
		if cbs := callbacks.onCDOTAUserMsg_CompendiumState; cbs != nil {
			msg := &dota.CDOTAUserMsg_CompendiumState{}
			if err := proto.Unmarshal(raw, msg); err != nil {
				return err
			}
			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return fmt.Errorf("no type found: %d", t)
}

func (c *Callbacks) OnAny(all func(interface{}) error) {
	c.OnCDemoStop(func(pkg *dota.CDemoStop) error { return all(pkg) })
	c.OnCDemoFileHeader(func(pkg *dota.CDemoFileHeader) error { return all(pkg) })
	c.OnCDemoFileInfo(func(pkg *dota.CDemoFileInfo) error { return all(pkg) })
	c.OnCDemoSyncTick(func(pkg *dota.CDemoSyncTick) error { return all(pkg) })
	c.OnCDemoSendTables(func(pkg *dota.CDemoSendTables) error { return all(pkg) })
	c.OnCDemoClassInfo(func(pkg *dota.CDemoClassInfo) error { return all(pkg) })
	c.OnCDemoStringTables(func(pkg *dota.CDemoStringTables) error { return all(pkg) })
	c.OnCDemoPacket(func(pkg *dota.CDemoPacket) error { return all(pkg) })
	c.OnCDemoSignonPacket(func(pkg *dota.CDemoPacket) error { return all(pkg) })
	c.OnCDemoConsoleCmd(func(pkg *dota.CDemoConsoleCmd) error { return all(pkg) })
	c.OnCDemoCustomData(func(pkg *dota.CDemoCustomData) error { return all(pkg) })
	c.OnCDemoCustomDataCallbacks(func(pkg *dota.CDemoCustomDataCallbacks) error { return all(pkg) })
	c.OnCDemoUserCmd(func(pkg *dota.CDemoUserCmd) error { return all(pkg) })
	c.OnCDemoFullPacket(func(pkg *dota.CDemoFullPacket) error { return all(pkg) })
	c.OnCDemoSaveGame(func(pkg *dota.CDemoSaveGame) error { return all(pkg) })
	c.OnCDemoSpawnGroups(func(pkg *dota.CDemoSpawnGroups) error { return all(pkg) })
	c.OnCNETMsg_NOP(func(pkg *dota.CNETMsg_NOP) error { return all(pkg) })
	c.OnCNETMsg_Disconnect(func(pkg *dota.CNETMsg_Disconnect) error { return all(pkg) })
	c.OnCNETMsg_File(func(pkg *dota.CNETMsg_File) error { return all(pkg) })
	c.OnCNETMsg_SplitScreenUser(func(pkg *dota.CNETMsg_SplitScreenUser) error { return all(pkg) })
	c.OnCNETMsg_Tick(func(pkg *dota.CNETMsg_Tick) error { return all(pkg) })
	c.OnCNETMsg_StringCmd(func(pkg *dota.CNETMsg_StringCmd) error { return all(pkg) })
	c.OnCNETMsg_SetConVar(func(pkg *dota.CNETMsg_SetConVar) error { return all(pkg) })
	c.OnCNETMsg_SignonState(func(pkg *dota.CNETMsg_SignonState) error { return all(pkg) })
	c.OnCNETMsg_SpawnGroup_Load(func(pkg *dota.CNETMsg_SpawnGroup_Load) error { return all(pkg) })
	c.OnCNETMsg_SpawnGroup_ManifestUpdate(func(pkg *dota.CNETMsg_SpawnGroup_ManifestUpdate) error { return all(pkg) })
	c.OnCNETMsg_SpawnGroup_SetCreationTick(func(pkg *dota.CNETMsg_SpawnGroup_SetCreationTick) error { return all(pkg) })
	c.OnCNETMsg_SpawnGroup_Unload(func(pkg *dota.CNETMsg_SpawnGroup_Unload) error { return all(pkg) })
	c.OnCNETMsg_SpawnGroup_LoadCompleted(func(pkg *dota.CNETMsg_SpawnGroup_LoadCompleted) error { return all(pkg) })
	c.OnCNETMsg_ReliableMessageEndMarker(func(pkg *dota.CNETMsg_ReliableMessageEndMarker) error { return all(pkg) })
	c.OnCSVCMsg_ServerInfo(func(pkg *dota.CSVCMsg_ServerInfo) error { return all(pkg) })
	c.OnCSVCMsg_FlattenedSerializer(func(pkg *dota.CSVCMsg_FlattenedSerializer) error { return all(pkg) })
	c.OnCSVCMsg_ClassInfo(func(pkg *dota.CSVCMsg_ClassInfo) error { return all(pkg) })
	c.OnCSVCMsg_SetPause(func(pkg *dota.CSVCMsg_SetPause) error { return all(pkg) })
	c.OnCSVCMsg_CreateStringTable(func(pkg *dota.CSVCMsg_CreateStringTable) error { return all(pkg) })
	c.OnCSVCMsg_UpdateStringTable(func(pkg *dota.CSVCMsg_UpdateStringTable) error { return all(pkg) })
	c.OnCSVCMsg_VoiceInit(func(pkg *dota.CSVCMsg_VoiceInit) error { return all(pkg) })
	c.OnCSVCMsg_VoiceData(func(pkg *dota.CSVCMsg_VoiceData) error { return all(pkg) })
	c.OnCSVCMsg_Print(func(pkg *dota.CSVCMsg_Print) error { return all(pkg) })
	c.OnCSVCMsg_Sounds(func(pkg *dota.CSVCMsg_Sounds) error { return all(pkg) })
	c.OnCSVCMsg_SetView(func(pkg *dota.CSVCMsg_SetView) error { return all(pkg) })
	c.OnCSVCMsg_ClearAllStringTables(func(pkg *dota.CSVCMsg_ClearAllStringTables) error { return all(pkg) })
	c.OnCSVCMsg_CmdKeyValues(func(pkg *dota.CSVCMsg_CmdKeyValues) error { return all(pkg) })
	c.OnCSVCMsg_BSPDecal(func(pkg *dota.CSVCMsg_BSPDecal) error { return all(pkg) })
	c.OnCSVCMsg_SplitScreen(func(pkg *dota.CSVCMsg_SplitScreen) error { return all(pkg) })
	c.OnCSVCMsg_PacketEntities(func(pkg *dota.CSVCMsg_PacketEntities) error { return all(pkg) })
	c.OnCSVCMsg_Prefetch(func(pkg *dota.CSVCMsg_Prefetch) error { return all(pkg) })
	c.OnCSVCMsg_Menu(func(pkg *dota.CSVCMsg_Menu) error { return all(pkg) })
	c.OnCSVCMsg_GetCvarValue(func(pkg *dota.CSVCMsg_GetCvarValue) error { return all(pkg) })
	c.OnCSVCMsg_StopSound(func(pkg *dota.CSVCMsg_StopSound) error { return all(pkg) })
	c.OnCSVCMsg_PeerList(func(pkg *dota.CSVCMsg_PeerList) error { return all(pkg) })
	c.OnCSVCMsg_PacketReliable(func(pkg *dota.CSVCMsg_PacketReliable) error { return all(pkg) })
	c.OnCSVCMsg_UserMessage(func(pkg *dota.CSVCMsg_UserMessage) error { return all(pkg) })
	c.OnCSVCMsg_SendTable(func(pkg *dota.CSVCMsg_SendTable) error { return all(pkg) })
	c.OnCSVCMsg_GameEvent(func(pkg *dota.CSVCMsg_GameEvent) error { return all(pkg) })
	c.OnCSVCMsg_TempEntities(func(pkg *dota.CSVCMsg_TempEntities) error { return all(pkg) })
	c.OnCSVCMsg_GameEventList(func(pkg *dota.CSVCMsg_GameEventList) error { return all(pkg) })
	c.OnCSVCMsg_FullFrameSplit(func(pkg *dota.CSVCMsg_FullFrameSplit) error { return all(pkg) })
	c.OnCUserMessageAchievementEvent(func(pkg *dota.CUserMessageAchievementEvent) error { return all(pkg) })
	c.OnCUserMessageCloseCaption(func(pkg *dota.CUserMessageCloseCaption) error { return all(pkg) })
	c.OnCUserMessageCloseCaptionDirect(func(pkg *dota.CUserMessageCloseCaptionDirect) error { return all(pkg) })
	c.OnCUserMessageCurrentTimescale(func(pkg *dota.CUserMessageCurrentTimescale) error { return all(pkg) })
	c.OnCUserMessageDesiredTimescale(func(pkg *dota.CUserMessageDesiredTimescale) error { return all(pkg) })
	c.OnCUserMessageFade(func(pkg *dota.CUserMessageFade) error { return all(pkg) })
	c.OnCUserMessageGameTitle(func(pkg *dota.CUserMessageGameTitle) error { return all(pkg) })
	c.OnCUserMessageHintText(func(pkg *dota.CUserMessageHintText) error { return all(pkg) })
	c.OnCUserMessageHudMsg(func(pkg *dota.CUserMessageHudMsg) error { return all(pkg) })
	c.OnCUserMessageHudText(func(pkg *dota.CUserMessageHudText) error { return all(pkg) })
	c.OnCUserMessageKeyHintText(func(pkg *dota.CUserMessageKeyHintText) error { return all(pkg) })
	c.OnCUserMessageColoredText(func(pkg *dota.CUserMessageColoredText) error { return all(pkg) })
	c.OnCUserMessageRequestState(func(pkg *dota.CUserMessageRequestState) error { return all(pkg) })
	c.OnCUserMessageResetHUD(func(pkg *dota.CUserMessageResetHUD) error { return all(pkg) })
	c.OnCUserMessageRumble(func(pkg *dota.CUserMessageRumble) error { return all(pkg) })
	c.OnCUserMessageSayText(func(pkg *dota.CUserMessageSayText) error { return all(pkg) })
	c.OnCUserMessageSayText2(func(pkg *dota.CUserMessageSayText2) error { return all(pkg) })
	c.OnCUserMessageSayTextChannel(func(pkg *dota.CUserMessageSayTextChannel) error { return all(pkg) })
	c.OnCUserMessageShake(func(pkg *dota.CUserMessageShake) error { return all(pkg) })
	c.OnCUserMessageShakeDir(func(pkg *dota.CUserMessageShakeDir) error { return all(pkg) })
	c.OnCUserMessageTextMsg(func(pkg *dota.CUserMessageTextMsg) error { return all(pkg) })
	c.OnCUserMessageScreenTilt(func(pkg *dota.CUserMessageScreenTilt) error { return all(pkg) })
	c.OnCUserMessageTrain(func(pkg *dota.CUserMessageTrain) error { return all(pkg) })
	c.OnCUserMessageVGUIMenu(func(pkg *dota.CUserMessageVGUIMenu) error { return all(pkg) })
	c.OnCUserMessageVoiceMask(func(pkg *dota.CUserMessageVoiceMask) error { return all(pkg) })
	c.OnCUserMessageVoiceSubtitle(func(pkg *dota.CUserMessageVoiceSubtitle) error { return all(pkg) })
	c.OnCUserMessageSendAudio(func(pkg *dota.CUserMessageSendAudio) error { return all(pkg) })
	c.OnCUserMessageItemPickup(func(pkg *dota.CUserMessageItemPickup) error { return all(pkg) })
	c.OnCUserMessageAmmoDenied(func(pkg *dota.CUserMessageAmmoDenied) error { return all(pkg) })
	c.OnCUserMessageCrosshairAngle(func(pkg *dota.CUserMessageCrosshairAngle) error { return all(pkg) })
	c.OnCUserMessageShowMenu(func(pkg *dota.CUserMessageShowMenu) error { return all(pkg) })
	c.OnCUserMessageCreditsMsg(func(pkg *dota.CUserMessageCreditsMsg) error { return all(pkg) })
	c.OnCUserMessageCloseCaptionPlaceholder(func(pkg *dota.CUserMessageCloseCaptionPlaceholder) error { return all(pkg) })
	c.OnCUserMessageCameraTransition(func(pkg *dota.CUserMessageCameraTransition) error { return all(pkg) })
	c.OnCUserMessageAudioParameter(func(pkg *dota.CUserMessageAudioParameter) error { return all(pkg) })
	c.OnCEntityMessagePlayJingle(func(pkg *dota.CEntityMessagePlayJingle) error { return all(pkg) })
	c.OnCEntityMessageScreenOverlay(func(pkg *dota.CEntityMessageScreenOverlay) error { return all(pkg) })
	c.OnCEntityMessageRemoveAllDecals(func(pkg *dota.CEntityMessageRemoveAllDecals) error { return all(pkg) })
	c.OnCEntityMessagePropagateForce(func(pkg *dota.CEntityMessagePropagateForce) error { return all(pkg) })
	c.OnCEntityMessageDoSpark(func(pkg *dota.CEntityMessageDoSpark) error { return all(pkg) })
	c.OnCEntityMessageFixAngle(func(pkg *dota.CEntityMessageFixAngle) error { return all(pkg) })
	c.OnCMsgVDebugGameSessionIDEvent(func(pkg *dota.CMsgVDebugGameSessionIDEvent) error { return all(pkg) })
	c.OnCMsgPlaceDecalEvent(func(pkg *dota.CMsgPlaceDecalEvent) error { return all(pkg) })
	c.OnCMsgClearWorldDecalsEvent(func(pkg *dota.CMsgClearWorldDecalsEvent) error { return all(pkg) })
	c.OnCMsgClearEntityDecalsEvent(func(pkg *dota.CMsgClearEntityDecalsEvent) error { return all(pkg) })
	c.OnCMsgClearDecalsForSkeletonInstanceEvent(func(pkg *dota.CMsgClearDecalsForSkeletonInstanceEvent) error { return all(pkg) })
	c.OnCMsgSource1LegacyGameEventList(func(pkg *dota.CMsgSource1LegacyGameEventList) error { return all(pkg) })
	c.OnCMsgSource1LegacyListenEvents(func(pkg *dota.CMsgSource1LegacyListenEvents) error { return all(pkg) })
	c.OnCMsgSource1LegacyGameEvent(func(pkg *dota.CMsgSource1LegacyGameEvent) error { return all(pkg) })
	c.OnCMsgSosStartSoundEvent(func(pkg *dota.CMsgSosStartSoundEvent) error { return all(pkg) })
	c.OnCMsgSosStopSoundEvent(func(pkg *dota.CMsgSosStopSoundEvent) error { return all(pkg) })
	c.OnCMsgSosSetSoundEventParams(func(pkg *dota.CMsgSosSetSoundEventParams) error { return all(pkg) })
	c.OnCMsgSosSetLibraryStackFields(func(pkg *dota.CMsgSosSetLibraryStackFields) error { return all(pkg) })
	c.OnCMsgSosStopSoundEventHash(func(pkg *dota.CMsgSosStopSoundEventHash) error { return all(pkg) })
	c.OnCDOTAUserMsg_AIDebugLine(func(pkg *dota.CDOTAUserMsg_AIDebugLine) error { return all(pkg) })
	c.OnCDOTAUserMsg_ChatEvent(func(pkg *dota.CDOTAUserMsg_ChatEvent) error { return all(pkg) })
	c.OnCDOTAUserMsg_CombatHeroPositions(func(pkg *dota.CDOTAUserMsg_CombatHeroPositions) error { return all(pkg) })
	c.OnCDOTAUserMsg_CombatLogShowDeath(func(pkg *dota.CDOTAUserMsg_CombatLogShowDeath) error { return all(pkg) })
	c.OnCDOTAUserMsg_CreateLinearProjectile(func(pkg *dota.CDOTAUserMsg_CreateLinearProjectile) error { return all(pkg) })
	c.OnCDOTAUserMsg_DestroyLinearProjectile(func(pkg *dota.CDOTAUserMsg_DestroyLinearProjectile) error { return all(pkg) })
	c.OnCDOTAUserMsg_DodgeTrackingProjectiles(func(pkg *dota.CDOTAUserMsg_DodgeTrackingProjectiles) error { return all(pkg) })
	c.OnCDOTAUserMsg_GlobalLightColor(func(pkg *dota.CDOTAUserMsg_GlobalLightColor) error { return all(pkg) })
	c.OnCDOTAUserMsg_GlobalLightDirection(func(pkg *dota.CDOTAUserMsg_GlobalLightDirection) error { return all(pkg) })
	c.OnCDOTAUserMsg_InvalidCommand(func(pkg *dota.CDOTAUserMsg_InvalidCommand) error { return all(pkg) })
	c.OnCDOTAUserMsg_LocationPing(func(pkg *dota.CDOTAUserMsg_LocationPing) error { return all(pkg) })
	c.OnCDOTAUserMsg_MapLine(func(pkg *dota.CDOTAUserMsg_MapLine) error { return all(pkg) })
	c.OnCDOTAUserMsg_MiniKillCamInfo(func(pkg *dota.CDOTAUserMsg_MiniKillCamInfo) error { return all(pkg) })
	c.OnCDOTAUserMsg_MinimapDebugPoint(func(pkg *dota.CDOTAUserMsg_MinimapDebugPoint) error { return all(pkg) })
	c.OnCDOTAUserMsg_MinimapEvent(func(pkg *dota.CDOTAUserMsg_MinimapEvent) error { return all(pkg) })
	c.OnCDOTAUserMsg_NevermoreRequiem(func(pkg *dota.CDOTAUserMsg_NevermoreRequiem) error { return all(pkg) })
	c.OnCDOTAUserMsg_OverheadEvent(func(pkg *dota.CDOTAUserMsg_OverheadEvent) error { return all(pkg) })
	c.OnCDOTAUserMsg_SetNextAutobuyItem(func(pkg *dota.CDOTAUserMsg_SetNextAutobuyItem) error { return all(pkg) })
	c.OnCDOTAUserMsg_SharedCooldown(func(pkg *dota.CDOTAUserMsg_SharedCooldown) error { return all(pkg) })
	c.OnCDOTAUserMsg_SpectatorPlayerClick(func(pkg *dota.CDOTAUserMsg_SpectatorPlayerClick) error { return all(pkg) })
	c.OnCDOTAUserMsg_TutorialTipInfo(func(pkg *dota.CDOTAUserMsg_TutorialTipInfo) error { return all(pkg) })
	c.OnCDOTAUserMsg_UnitEvent(func(pkg *dota.CDOTAUserMsg_UnitEvent) error { return all(pkg) })
	c.OnCDOTAUserMsg_ParticleManager(func(pkg *dota.CDOTAUserMsg_ParticleManager) error { return all(pkg) })
	c.OnCDOTAUserMsg_BotChat(func(pkg *dota.CDOTAUserMsg_BotChat) error { return all(pkg) })
	c.OnCDOTAUserMsg_HudError(func(pkg *dota.CDOTAUserMsg_HudError) error { return all(pkg) })
	c.OnCDOTAUserMsg_ItemPurchased(func(pkg *dota.CDOTAUserMsg_ItemPurchased) error { return all(pkg) })
	c.OnCDOTAUserMsg_Ping(func(pkg *dota.CDOTAUserMsg_Ping) error { return all(pkg) })
	c.OnCDOTAUserMsg_ItemFound(func(pkg *dota.CDOTAUserMsg_ItemFound) error { return all(pkg) })
	c.OnCDOTAUserMsg_SwapVerify(func(pkg *dota.CDOTAUserMsg_SwapVerify) error { return all(pkg) })
	c.OnCDOTAUserMsg_WorldLine(func(pkg *dota.CDOTAUserMsg_WorldLine) error { return all(pkg) })
	c.OnCDOTAUserMsg_ItemAlert(func(pkg *dota.CDOTAUserMsg_ItemAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_HalloweenDrops(func(pkg *dota.CDOTAUserMsg_HalloweenDrops) error { return all(pkg) })
	c.OnCDOTAUserMsg_ChatWheel(func(pkg *dota.CDOTAUserMsg_ChatWheel) error { return all(pkg) })
	c.OnCDOTAUserMsg_ReceivedXmasGift(func(pkg *dota.CDOTAUserMsg_ReceivedXmasGift) error { return all(pkg) })
	c.OnCDOTAUserMsg_UpdateSharedContent(func(pkg *dota.CDOTAUserMsg_UpdateSharedContent) error { return all(pkg) })
	c.OnCDOTAUserMsg_TutorialRequestExp(func(pkg *dota.CDOTAUserMsg_TutorialRequestExp) error { return all(pkg) })
	c.OnCDOTAUserMsg_TutorialPingMinimap(func(pkg *dota.CDOTAUserMsg_TutorialPingMinimap) error { return all(pkg) })
	c.OnCDOTAUserMsg_GamerulesStateChanged(func(pkg *dota.CDOTAUserMsg_GamerulesStateChanged) error { return all(pkg) })
	c.OnCDOTAUserMsg_ShowSurvey(func(pkg *dota.CDOTAUserMsg_ShowSurvey) error { return all(pkg) })
	c.OnCDOTAUserMsg_TutorialFade(func(pkg *dota.CDOTAUserMsg_TutorialFade) error { return all(pkg) })
	c.OnCDOTAUserMsg_AddQuestLogEntry(func(pkg *dota.CDOTAUserMsg_AddQuestLogEntry) error { return all(pkg) })
	c.OnCDOTAUserMsg_SendStatPopup(func(pkg *dota.CDOTAUserMsg_SendStatPopup) error { return all(pkg) })
	c.OnCDOTAUserMsg_TutorialFinish(func(pkg *dota.CDOTAUserMsg_TutorialFinish) error { return all(pkg) })
	c.OnCDOTAUserMsg_SendRoshanPopup(func(pkg *dota.CDOTAUserMsg_SendRoshanPopup) error { return all(pkg) })
	c.OnCDOTAUserMsg_SendGenericToolTip(func(pkg *dota.CDOTAUserMsg_SendGenericToolTip) error { return all(pkg) })
	c.OnCDOTAUserMsg_SendFinalGold(func(pkg *dota.CDOTAUserMsg_SendFinalGold) error { return all(pkg) })
	c.OnCDOTAUserMsg_CustomMsg(func(pkg *dota.CDOTAUserMsg_CustomMsg) error { return all(pkg) })
	c.OnCDOTAUserMsg_CoachHUDPing(func(pkg *dota.CDOTAUserMsg_CoachHUDPing) error { return all(pkg) })
	c.OnCDOTAUserMsg_ClientLoadGridNav(func(pkg *dota.CDOTAUserMsg_ClientLoadGridNav) error { return all(pkg) })
	c.OnCDOTAUserMsg_TE_Projectile(func(pkg *dota.CDOTAUserMsg_TE_Projectile) error { return all(pkg) })
	c.OnCDOTAUserMsg_TE_ProjectileLoc(func(pkg *dota.CDOTAUserMsg_TE_ProjectileLoc) error { return all(pkg) })
	c.OnCDOTAUserMsg_TE_DotaBloodImpact(func(pkg *dota.CDOTAUserMsg_TE_DotaBloodImpact) error { return all(pkg) })
	c.OnCDOTAUserMsg_TE_UnitAnimation(func(pkg *dota.CDOTAUserMsg_TE_UnitAnimation) error { return all(pkg) })
	c.OnCDOTAUserMsg_TE_UnitAnimationEnd(func(pkg *dota.CDOTAUserMsg_TE_UnitAnimationEnd) error { return all(pkg) })
	c.OnCDOTAUserMsg_AbilityPing(func(pkg *dota.CDOTAUserMsg_AbilityPing) error { return all(pkg) })
	c.OnCDOTAUserMsg_ShowGenericPopup(func(pkg *dota.CDOTAUserMsg_ShowGenericPopup) error { return all(pkg) })
	c.OnCDOTAUserMsg_VoteStart(func(pkg *dota.CDOTAUserMsg_VoteStart) error { return all(pkg) })
	c.OnCDOTAUserMsg_VoteUpdate(func(pkg *dota.CDOTAUserMsg_VoteUpdate) error { return all(pkg) })
	c.OnCDOTAUserMsg_VoteEnd(func(pkg *dota.CDOTAUserMsg_VoteEnd) error { return all(pkg) })
	c.OnCDOTAUserMsg_BoosterState(func(pkg *dota.CDOTAUserMsg_BoosterState) error { return all(pkg) })
	c.OnCDOTAUserMsg_WillPurchaseAlert(func(pkg *dota.CDOTAUserMsg_WillPurchaseAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_TutorialMinimapPosition(func(pkg *dota.CDOTAUserMsg_TutorialMinimapPosition) error { return all(pkg) })
	c.OnCDOTAUserMsg_PlayerMMR(func(pkg *dota.CDOTAUserMsg_PlayerMMR) error { return all(pkg) })
	c.OnCDOTAUserMsg_AbilitySteal(func(pkg *dota.CDOTAUserMsg_AbilitySteal) error { return all(pkg) })
	c.OnCDOTAUserMsg_CourierKilledAlert(func(pkg *dota.CDOTAUserMsg_CourierKilledAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_EnemyItemAlert(func(pkg *dota.CDOTAUserMsg_EnemyItemAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_StatsMatchDetails(func(pkg *dota.CDOTAUserMsg_StatsMatchDetails) error { return all(pkg) })
	c.OnCDOTAUserMsg_MiniTaunt(func(pkg *dota.CDOTAUserMsg_MiniTaunt) error { return all(pkg) })
	c.OnCDOTAUserMsg_BuyBackStateAlert(func(pkg *dota.CDOTAUserMsg_BuyBackStateAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_SpeechBubble(func(pkg *dota.CDOTAUserMsg_SpeechBubble) error { return all(pkg) })
	c.OnCDOTAUserMsg_CustomHeaderMessage(func(pkg *dota.CDOTAUserMsg_CustomHeaderMessage) error { return all(pkg) })
	c.OnCDOTAUserMsg_QuickBuyAlert(func(pkg *dota.CDOTAUserMsg_QuickBuyAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_PredictionResult(func(pkg *dota.CDOTAUserMsg_PredictionResult) error { return all(pkg) })
	c.OnCDOTAUserMsg_ModifierAlert(func(pkg *dota.CDOTAUserMsg_ModifierAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_HPManaAlert(func(pkg *dota.CDOTAUserMsg_HPManaAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_GlyphAlert(func(pkg *dota.CDOTAUserMsg_GlyphAlert) error { return all(pkg) })
	c.OnCDOTAUserMsg_BeastChat(func(pkg *dota.CDOTAUserMsg_BeastChat) error { return all(pkg) })
	c.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(pkg *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error { return all(pkg) })
	c.OnCDOTAUserMsg_CustomHudElement_Create(func(pkg *dota.CDOTAUserMsg_CustomHudElement_Create) error { return all(pkg) })
	c.OnCDOTAUserMsg_CustomHudElement_Modify(func(pkg *dota.CDOTAUserMsg_CustomHudElement_Modify) error { return all(pkg) })
	c.OnCDOTAUserMsg_CustomHudElement_Destroy(func(pkg *dota.CDOTAUserMsg_CustomHudElement_Destroy) error { return all(pkg) })
	c.OnCDOTAUserMsg_CompendiumState(func(pkg *dota.CDOTAUserMsg_CompendiumState) error { return all(pkg) })
}
