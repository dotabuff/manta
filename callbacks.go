package manta

import (
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

// Callbacks decodes and routes replay events to callback functions
type Callbacks struct {
	onCDemoStop                                            []func(*dota.CDemoStop) error
	onCDemoFileHeader                                      []func(*dota.CDemoFileHeader) error
	onCDemoFileInfo                                        []func(*dota.CDemoFileInfo) error
	onCDemoSyncTick                                        []func(*dota.CDemoSyncTick) error
	onCDemoSendTables                                      []func(*dota.CDemoSendTables) error
	onCDemoClassInfo                                       []func(*dota.CDemoClassInfo) error
	onCDemoStringTables                                    []func(*dota.CDemoStringTables) error
	onCDemoPacket                                          []func(*dota.CDemoPacket) error
	onCDemoSignonPacket                                    []func(*dota.CDemoPacket) error
	onCDemoConsoleCmd                                      []func(*dota.CDemoConsoleCmd) error
	onCDemoCustomData                                      []func(*dota.CDemoCustomData) error
	onCDemoCustomDataCallbacks                             []func(*dota.CDemoCustomDataCallbacks) error
	onCDemoUserCmd                                         []func(*dota.CDemoUserCmd) error
	onCDemoFullPacket                                      []func(*dota.CDemoFullPacket) error
	onCDemoSaveGame                                        []func(*dota.CDemoSaveGame) error
	onCDemoSpawnGroups                                     []func(*dota.CDemoSpawnGroups) error
	onCDemoAnimationData                                   []func(*dota.CDemoAnimationData) error
	onCDemoAnimationHeader                                 []func(*dota.CDemoAnimationHeader) error
	onCDemoRecovery                                        []func(*dota.CDemoRecovery) error
	onCNETMsg_NOP                                          []func(*dota.CNETMsg_NOP) error
	onCNETMsg_SplitScreenUser                              []func(*dota.CNETMsg_SplitScreenUser) error
	onCNETMsg_Tick                                         []func(*dota.CNETMsg_Tick) error
	onCNETMsg_StringCmd                                    []func(*dota.CNETMsg_StringCmd) error
	onCNETMsg_SetConVar                                    []func(*dota.CNETMsg_SetConVar) error
	onCNETMsg_SignonState                                  []func(*dota.CNETMsg_SignonState) error
	onCNETMsg_SpawnGroup_Load                              []func(*dota.CNETMsg_SpawnGroup_Load) error
	onCNETMsg_SpawnGroup_ManifestUpdate                    []func(*dota.CNETMsg_SpawnGroup_ManifestUpdate) error
	onCNETMsg_SpawnGroup_SetCreationTick                   []func(*dota.CNETMsg_SpawnGroup_SetCreationTick) error
	onCNETMsg_SpawnGroup_Unload                            []func(*dota.CNETMsg_SpawnGroup_Unload) error
	onCNETMsg_SpawnGroup_LoadCompleted                     []func(*dota.CNETMsg_SpawnGroup_LoadCompleted) error
	onCNETMsg_DebugOverlay                                 []func(*dota.CNETMsg_DebugOverlay) error
	onCSVCMsg_ServerInfo                                   []func(*dota.CSVCMsg_ServerInfo) error
	onCSVCMsg_FlattenedSerializer                          []func(*dota.CSVCMsg_FlattenedSerializer) error
	onCSVCMsg_ClassInfo                                    []func(*dota.CSVCMsg_ClassInfo) error
	onCSVCMsg_SetPause                                     []func(*dota.CSVCMsg_SetPause) error
	onCSVCMsg_CreateStringTable                            []func(*dota.CSVCMsg_CreateStringTable) error
	onCSVCMsg_UpdateStringTable                            []func(*dota.CSVCMsg_UpdateStringTable) error
	onCSVCMsg_VoiceInit                                    []func(*dota.CSVCMsg_VoiceInit) error
	onCSVCMsg_VoiceData                                    []func(*dota.CSVCMsg_VoiceData) error
	onCSVCMsg_Print                                        []func(*dota.CSVCMsg_Print) error
	onCSVCMsg_Sounds                                       []func(*dota.CSVCMsg_Sounds) error
	onCSVCMsg_SetView                                      []func(*dota.CSVCMsg_SetView) error
	onCSVCMsg_ClearAllStringTables                         []func(*dota.CSVCMsg_ClearAllStringTables) error
	onCSVCMsg_CmdKeyValues                                 []func(*dota.CSVCMsg_CmdKeyValues) error
	onCSVCMsg_BSPDecal                                     []func(*dota.CSVCMsg_BSPDecal) error
	onCSVCMsg_SplitScreen                                  []func(*dota.CSVCMsg_SplitScreen) error
	onCSVCMsg_PacketEntities                               []func(*dota.CSVCMsg_PacketEntities) error
	onCSVCMsg_Prefetch                                     []func(*dota.CSVCMsg_Prefetch) error
	onCSVCMsg_Menu                                         []func(*dota.CSVCMsg_Menu) error
	onCSVCMsg_GetCvarValue                                 []func(*dota.CSVCMsg_GetCvarValue) error
	onCSVCMsg_StopSound                                    []func(*dota.CSVCMsg_StopSound) error
	onCSVCMsg_PeerList                                     []func(*dota.CSVCMsg_PeerList) error
	onCSVCMsg_PacketReliable                               []func(*dota.CSVCMsg_PacketReliable) error
	onCSVCMsg_HLTVStatus                                   []func(*dota.CSVCMsg_HLTVStatus) error
	onCSVCMsg_ServerSteamID                                []func(*dota.CSVCMsg_ServerSteamID) error
	onCSVCMsg_FullFrameSplit                               []func(*dota.CSVCMsg_FullFrameSplit) error
	onCSVCMsg_RconServerDetails                            []func(*dota.CSVCMsg_RconServerDetails) error
	onCSVCMsg_UserMessage                                  []func(*dota.CSVCMsg_UserMessage) error
	onCSVCMsg_Broadcast_Command                            []func(*dota.CSVCMsg_Broadcast_Command) error
	onCSVCMsg_HltvFixupOperatorStatus                      []func(*dota.CSVCMsg_HltvFixupOperatorStatus) error
	onCUserMessageAchievementEvent                         []func(*dota.CUserMessageAchievementEvent) error
	onCUserMessageCloseCaption                             []func(*dota.CUserMessageCloseCaption) error
	onCUserMessageCloseCaptionDirect                       []func(*dota.CUserMessageCloseCaptionDirect) error
	onCUserMessageCurrentTimescale                         []func(*dota.CUserMessageCurrentTimescale) error
	onCUserMessageDesiredTimescale                         []func(*dota.CUserMessageDesiredTimescale) error
	onCUserMessageFade                                     []func(*dota.CUserMessageFade) error
	onCUserMessageGameTitle                                []func(*dota.CUserMessageGameTitle) error
	onCUserMessageHudMsg                                   []func(*dota.CUserMessageHudMsg) error
	onCUserMessageHudText                                  []func(*dota.CUserMessageHudText) error
	onCUserMessageColoredText                              []func(*dota.CUserMessageColoredText) error
	onCUserMessageRequestState                             []func(*dota.CUserMessageRequestState) error
	onCUserMessageResetHUD                                 []func(*dota.CUserMessageResetHUD) error
	onCUserMessageRumble                                   []func(*dota.CUserMessageRumble) error
	onCUserMessageSayText                                  []func(*dota.CUserMessageSayText) error
	onCUserMessageSayText2                                 []func(*dota.CUserMessageSayText2) error
	onCUserMessageSayTextChannel                           []func(*dota.CUserMessageSayTextChannel) error
	onCUserMessageShake                                    []func(*dota.CUserMessageShake) error
	onCUserMessageShakeDir                                 []func(*dota.CUserMessageShakeDir) error
	onCUserMessageWaterShake                               []func(*dota.CUserMessageWaterShake) error
	onCUserMessageTextMsg                                  []func(*dota.CUserMessageTextMsg) error
	onCUserMessageScreenTilt                               []func(*dota.CUserMessageScreenTilt) error
	onCUserMessageVoiceMask                                []func(*dota.CUserMessageVoiceMask) error
	onCUserMessageSendAudio                                []func(*dota.CUserMessageSendAudio) error
	onCUserMessageItemPickup                               []func(*dota.CUserMessageItemPickup) error
	onCUserMessageAmmoDenied                               []func(*dota.CUserMessageAmmoDenied) error
	onCUserMessageShowMenu                                 []func(*dota.CUserMessageShowMenu) error
	onCUserMessageCreditsMsg                               []func(*dota.CUserMessageCreditsMsg) error
	onCEntityMessagePlayJingle                             []func(*dota.CEntityMessagePlayJingle) error
	onCEntityMessageScreenOverlay                          []func(*dota.CEntityMessageScreenOverlay) error
	onCEntityMessageRemoveAllDecals                        []func(*dota.CEntityMessageRemoveAllDecals) error
	onCEntityMessagePropagateForce                         []func(*dota.CEntityMessagePropagateForce) error
	onCEntityMessageDoSpark                                []func(*dota.CEntityMessageDoSpark) error
	onCEntityMessageFixAngle                               []func(*dota.CEntityMessageFixAngle) error
	onCUserMessageCloseCaptionPlaceholder                  []func(*dota.CUserMessageCloseCaptionPlaceholder) error
	onCUserMessageCameraTransition                         []func(*dota.CUserMessageCameraTransition) error
	onCUserMessageAudioParameter                           []func(*dota.CUserMessageAudioParameter) error
	onCUserMessageHapticsManagerPulse                      []func(*dota.CUserMessageHapticsManagerPulse) error
	onCUserMessageHapticsManagerEffect                     []func(*dota.CUserMessageHapticsManagerEffect) error
	onCUserMessageUpdateCssClasses                         []func(*dota.CUserMessageUpdateCssClasses) error
	onCUserMessageServerFrameTime                          []func(*dota.CUserMessageServerFrameTime) error
	onCUserMessageLagCompensationError                     []func(*dota.CUserMessageLagCompensationError) error
	onCUserMessageRequestDllStatus                         []func(*dota.CUserMessageRequestDllStatus) error
	onCUserMessageRequestUtilAction                        []func(*dota.CUserMessageRequestUtilAction) error
	onCUserMessageRequestInventory                         []func(*dota.CUserMessageRequestInventory) error
	onCUserMessageRequestDiagnostic                        []func(*dota.CUserMessageRequestDiagnostic) error
	onCMsgVDebugGameSessionIDEvent                         []func(*dota.CMsgVDebugGameSessionIDEvent) error
	onCMsgPlaceDecalEvent                                  []func(*dota.CMsgPlaceDecalEvent) error
	onCMsgClearWorldDecalsEvent                            []func(*dota.CMsgClearWorldDecalsEvent) error
	onCMsgClearEntityDecalsEvent                           []func(*dota.CMsgClearEntityDecalsEvent) error
	onCMsgClearDecalsForSkeletonInstanceEvent              []func(*dota.CMsgClearDecalsForSkeletonInstanceEvent) error
	onCMsgSource1LegacyGameEventList                       []func(*dota.CMsgSource1LegacyGameEventList) error
	onCMsgSource1LegacyListenEvents                        []func(*dota.CMsgSource1LegacyListenEvents) error
	onCMsgSource1LegacyGameEvent                           []func(*dota.CMsgSource1LegacyGameEvent) error
	onCMsgSosStartSoundEvent                               []func(*dota.CMsgSosStartSoundEvent) error
	onCMsgSosStopSoundEvent                                []func(*dota.CMsgSosStopSoundEvent) error
	onCMsgSosSetSoundEventParams                           []func(*dota.CMsgSosSetSoundEventParams) error
	onCMsgSosSetLibraryStackFields                         []func(*dota.CMsgSosSetLibraryStackFields) error
	onCMsgSosStopSoundEventHash                            []func(*dota.CMsgSosStopSoundEventHash) error
	onCDOTAUserMsg_AIDebugLine                             []func(*dota.CDOTAUserMsg_AIDebugLine) error
	onCDOTAUserMsg_ChatEvent                               []func(*dota.CDOTAUserMsg_ChatEvent) error
	onCDOTAUserMsg_CombatHeroPositions                     []func(*dota.CDOTAUserMsg_CombatHeroPositions) error
	onCDOTAUserMsg_CombatLogBulkData                       []func(*dota.CDOTAUserMsg_CombatLogBulkData) error
	onCDOTAUserMsg_CreateLinearProjectile                  []func(*dota.CDOTAUserMsg_CreateLinearProjectile) error
	onCDOTAUserMsg_DestroyLinearProjectile                 []func(*dota.CDOTAUserMsg_DestroyLinearProjectile) error
	onCDOTAUserMsg_DodgeTrackingProjectiles                []func(*dota.CDOTAUserMsg_DodgeTrackingProjectiles) error
	onCDOTAUserMsg_GlobalLightColor                        []func(*dota.CDOTAUserMsg_GlobalLightColor) error
	onCDOTAUserMsg_GlobalLightDirection                    []func(*dota.CDOTAUserMsg_GlobalLightDirection) error
	onCDOTAUserMsg_InvalidCommand                          []func(*dota.CDOTAUserMsg_InvalidCommand) error
	onCDOTAUserMsg_LocationPing                            []func(*dota.CDOTAUserMsg_LocationPing) error
	onCDOTAUserMsg_MapLine                                 []func(*dota.CDOTAUserMsg_MapLine) error
	onCDOTAUserMsg_MiniKillCamInfo                         []func(*dota.CDOTAUserMsg_MiniKillCamInfo) error
	onCDOTAUserMsg_MinimapDebugPoint                       []func(*dota.CDOTAUserMsg_MinimapDebugPoint) error
	onCDOTAUserMsg_MinimapEvent                            []func(*dota.CDOTAUserMsg_MinimapEvent) error
	onCDOTAUserMsg_NevermoreRequiem                        []func(*dota.CDOTAUserMsg_NevermoreRequiem) error
	onCDOTAUserMsg_OverheadEvent                           []func(*dota.CDOTAUserMsg_OverheadEvent) error
	onCDOTAUserMsg_SetNextAutobuyItem                      []func(*dota.CDOTAUserMsg_SetNextAutobuyItem) error
	onCDOTAUserMsg_SharedCooldown                          []func(*dota.CDOTAUserMsg_SharedCooldown) error
	onCDOTAUserMsg_SpectatorPlayerClick                    []func(*dota.CDOTAUserMsg_SpectatorPlayerClick) error
	onCDOTAUserMsg_TutorialTipInfo                         []func(*dota.CDOTAUserMsg_TutorialTipInfo) error
	onCDOTAUserMsg_UnitEvent                               []func(*dota.CDOTAUserMsg_UnitEvent) error
	onCDOTAUserMsg_BotChat                                 []func(*dota.CDOTAUserMsg_BotChat) error
	onCDOTAUserMsg_HudError                                []func(*dota.CDOTAUserMsg_HudError) error
	onCDOTAUserMsg_ItemPurchased                           []func(*dota.CDOTAUserMsg_ItemPurchased) error
	onCDOTAUserMsg_Ping                                    []func(*dota.CDOTAUserMsg_Ping) error
	onCDOTAUserMsg_ItemFound                               []func(*dota.CDOTAUserMsg_ItemFound) error
	onCDOTAUserMsg_SwapVerify                              []func(*dota.CDOTAUserMsg_SwapVerify) error
	onCDOTAUserMsg_WorldLine                               []func(*dota.CDOTAUserMsg_WorldLine) error
	onCMsgGCToClientTournamentItemDrop                     []func(*dota.CMsgGCToClientTournamentItemDrop) error
	onCDOTAUserMsg_ItemAlert                               []func(*dota.CDOTAUserMsg_ItemAlert) error
	onCDOTAUserMsg_HalloweenDrops                          []func(*dota.CDOTAUserMsg_HalloweenDrops) error
	onCDOTAUserMsg_ChatWheel                               []func(*dota.CDOTAUserMsg_ChatWheel) error
	onCDOTAUserMsg_ReceivedXmasGift                        []func(*dota.CDOTAUserMsg_ReceivedXmasGift) error
	onCDOTAUserMsg_UpdateSharedContent                     []func(*dota.CDOTAUserMsg_UpdateSharedContent) error
	onCDOTAUserMsg_TutorialRequestExp                      []func(*dota.CDOTAUserMsg_TutorialRequestExp) error
	onCDOTAUserMsg_TutorialPingMinimap                     []func(*dota.CDOTAUserMsg_TutorialPingMinimap) error
	onCDOTAUserMsg_GamerulesStateChanged                   []func(*dota.CDOTAUserMsg_GamerulesStateChanged) error
	onCDOTAUserMsg_ShowSurvey                              []func(*dota.CDOTAUserMsg_ShowSurvey) error
	onCDOTAUserMsg_TutorialFade                            []func(*dota.CDOTAUserMsg_TutorialFade) error
	onCDOTAUserMsg_AddQuestLogEntry                        []func(*dota.CDOTAUserMsg_AddQuestLogEntry) error
	onCDOTAUserMsg_SendStatPopup                           []func(*dota.CDOTAUserMsg_SendStatPopup) error
	onCDOTAUserMsg_TutorialFinish                          []func(*dota.CDOTAUserMsg_TutorialFinish) error
	onCDOTAUserMsg_SendRoshanPopup                         []func(*dota.CDOTAUserMsg_SendRoshanPopup) error
	onCDOTAUserMsg_SendGenericToolTip                      []func(*dota.CDOTAUserMsg_SendGenericToolTip) error
	onCDOTAUserMsg_SendFinalGold                           []func(*dota.CDOTAUserMsg_SendFinalGold) error
	onCDOTAUserMsg_CustomMsg                               []func(*dota.CDOTAUserMsg_CustomMsg) error
	onCDOTAUserMsg_CoachHUDPing                            []func(*dota.CDOTAUserMsg_CoachHUDPing) error
	onCDOTAUserMsg_ClientLoadGridNav                       []func(*dota.CDOTAUserMsg_ClientLoadGridNav) error
	onCDOTAUserMsg_TE_Projectile                           []func(*dota.CDOTAUserMsg_TE_Projectile) error
	onCDOTAUserMsg_TE_ProjectileLoc                        []func(*dota.CDOTAUserMsg_TE_ProjectileLoc) error
	onCDOTAUserMsg_TE_DotaBloodImpact                      []func(*dota.CDOTAUserMsg_TE_DotaBloodImpact) error
	onCDOTAUserMsg_TE_UnitAnimation                        []func(*dota.CDOTAUserMsg_TE_UnitAnimation) error
	onCDOTAUserMsg_TE_UnitAnimationEnd                     []func(*dota.CDOTAUserMsg_TE_UnitAnimationEnd) error
	onCDOTAUserMsg_AbilityPing                             []func(*dota.CDOTAUserMsg_AbilityPing) error
	onCDOTAUserMsg_ShowGenericPopup                        []func(*dota.CDOTAUserMsg_ShowGenericPopup) error
	onCDOTAUserMsg_VoteStart                               []func(*dota.CDOTAUserMsg_VoteStart) error
	onCDOTAUserMsg_VoteUpdate                              []func(*dota.CDOTAUserMsg_VoteUpdate) error
	onCDOTAUserMsg_VoteEnd                                 []func(*dota.CDOTAUserMsg_VoteEnd) error
	onCDOTAUserMsg_BoosterState                            []func(*dota.CDOTAUserMsg_BoosterState) error
	onCDOTAUserMsg_WillPurchaseAlert                       []func(*dota.CDOTAUserMsg_WillPurchaseAlert) error
	onCDOTAUserMsg_TutorialMinimapPosition                 []func(*dota.CDOTAUserMsg_TutorialMinimapPosition) error
	onCDOTAUserMsg_AbilitySteal                            []func(*dota.CDOTAUserMsg_AbilitySteal) error
	onCDOTAUserMsg_CourierKilledAlert                      []func(*dota.CDOTAUserMsg_CourierKilledAlert) error
	onCDOTAUserMsg_EnemyItemAlert                          []func(*dota.CDOTAUserMsg_EnemyItemAlert) error
	onCDOTAUserMsg_StatsMatchDetails                       []func(*dota.CDOTAUserMsg_StatsMatchDetails) error
	onCDOTAUserMsg_MiniTaunt                               []func(*dota.CDOTAUserMsg_MiniTaunt) error
	onCDOTAUserMsg_BuyBackStateAlert                       []func(*dota.CDOTAUserMsg_BuyBackStateAlert) error
	onCDOTAUserMsg_SpeechBubble                            []func(*dota.CDOTAUserMsg_SpeechBubble) error
	onCDOTAUserMsg_CustomHeaderMessage                     []func(*dota.CDOTAUserMsg_CustomHeaderMessage) error
	onCDOTAUserMsg_QuickBuyAlert                           []func(*dota.CDOTAUserMsg_QuickBuyAlert) error
	onCDOTAUserMsg_StatsHeroMinuteDetails                  []func(*dota.CDOTAUserMsg_StatsHeroMinuteDetails) error
	onCDOTAUserMsg_ModifierAlert                           []func(*dota.CDOTAUserMsg_ModifierAlert) error
	onCDOTAUserMsg_HPManaAlert                             []func(*dota.CDOTAUserMsg_HPManaAlert) error
	onCDOTAUserMsg_GlyphAlert                              []func(*dota.CDOTAUserMsg_GlyphAlert) error
	onCDOTAUserMsg_BeastChat                               []func(*dota.CDOTAUserMsg_BeastChat) error
	onCDOTAUserMsg_SpectatorPlayerUnitOrders               []func(*dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error
	onCDOTAUserMsg_CustomHudElement_Create                 []func(*dota.CDOTAUserMsg_CustomHudElement_Create) error
	onCDOTAUserMsg_CustomHudElement_Modify                 []func(*dota.CDOTAUserMsg_CustomHudElement_Modify) error
	onCDOTAUserMsg_CustomHudElement_Destroy                []func(*dota.CDOTAUserMsg_CustomHudElement_Destroy) error
	onCDOTAUserMsg_CompendiumState                         []func(*dota.CDOTAUserMsg_CompendiumState) error
	onCDOTAUserMsg_ProjectionAbility                       []func(*dota.CDOTAUserMsg_ProjectionAbility) error
	onCDOTAUserMsg_ProjectionEvent                         []func(*dota.CDOTAUserMsg_ProjectionEvent) error
	onCMsgDOTACombatLogEntry                               []func(*dota.CMsgDOTACombatLogEntry) error
	onCDOTAUserMsg_XPAlert                                 []func(*dota.CDOTAUserMsg_XPAlert) error
	onCDOTAUserMsg_UpdateQuestProgress                     []func(*dota.CDOTAUserMsg_UpdateQuestProgress) error
	onCDOTAMatchMetadataFile                               []func(*dota.CDOTAMatchMetadataFile) error
	onCDOTAUserMsg_QuestStatus                             []func(*dota.CDOTAUserMsg_QuestStatus) error
	onCDOTAUserMsg_SuggestHeroPick                         []func(*dota.CDOTAUserMsg_SuggestHeroPick) error
	onCDOTAUserMsg_SuggestHeroRole                         []func(*dota.CDOTAUserMsg_SuggestHeroRole) error
	onCDOTAUserMsg_KillcamDamageTaken                      []func(*dota.CDOTAUserMsg_KillcamDamageTaken) error
	onCDOTAUserMsg_SelectPenaltyGold                       []func(*dota.CDOTAUserMsg_SelectPenaltyGold) error
	onCDOTAUserMsg_RollDiceResult                          []func(*dota.CDOTAUserMsg_RollDiceResult) error
	onCDOTAUserMsg_FlipCoinResult                          []func(*dota.CDOTAUserMsg_FlipCoinResult) error
	onCDOTAUserMsg_SendRoshanSpectatorPhase                []func(*dota.CDOTAUserMsg_SendRoshanSpectatorPhase) error
	onCDOTAUserMsg_ChatWheelCooldown                       []func(*dota.CDOTAUserMsg_ChatWheelCooldown) error
	onCDOTAUserMsg_DismissAllStatPopups                    []func(*dota.CDOTAUserMsg_DismissAllStatPopups) error
	onCDOTAUserMsg_TE_DestroyProjectile                    []func(*dota.CDOTAUserMsg_TE_DestroyProjectile) error
	onCDOTAUserMsg_HeroRelicProgress                       []func(*dota.CDOTAUserMsg_HeroRelicProgress) error
	onCDOTAUserMsg_AbilityDraftRequestAbility              []func(*dota.CDOTAUserMsg_AbilityDraftRequestAbility) error
	onCDOTAUserMsg_ItemSold                                []func(*dota.CDOTAUserMsg_ItemSold) error
	onCDOTAUserMsg_DamageReport                            []func(*dota.CDOTAUserMsg_DamageReport) error
	onCDOTAUserMsg_SalutePlayer                            []func(*dota.CDOTAUserMsg_SalutePlayer) error
	onCDOTAUserMsg_TipAlert                                []func(*dota.CDOTAUserMsg_TipAlert) error
	onCDOTAUserMsg_ReplaceQueryUnit                        []func(*dota.CDOTAUserMsg_ReplaceQueryUnit) error
	onCDOTAUserMsg_EmptyTeleportAlert                      []func(*dota.CDOTAUserMsg_EmptyTeleportAlert) error
	onCDOTAUserMsg_MarsArenaOfBloodAttack                  []func(*dota.CDOTAUserMsg_MarsArenaOfBloodAttack) error
	onCDOTAUserMsg_ESArcanaCombo                           []func(*dota.CDOTAUserMsg_ESArcanaCombo) error
	onCDOTAUserMsg_ESArcanaComboSummary                    []func(*dota.CDOTAUserMsg_ESArcanaComboSummary) error
	onCDOTAUserMsg_HighFiveLeftHanging                     []func(*dota.CDOTAUserMsg_HighFiveLeftHanging) error
	onCDOTAUserMsg_HighFiveCompleted                       []func(*dota.CDOTAUserMsg_HighFiveCompleted) error
	onCDOTAUserMsg_ShovelUnearth                           []func(*dota.CDOTAUserMsg_ShovelUnearth) error
	onCDOTAUserMsg_RadarAlert                              []func(*dota.CDOTAUserMsg_RadarAlert) error
	onCDOTAUserMsg_AllStarEvent                            []func(*dota.CDOTAUserMsg_AllStarEvent) error
	onCDOTAUserMsg_TalentTreeAlert                         []func(*dota.CDOTAUserMsg_TalentTreeAlert) error
	onCDOTAUserMsg_QueuedOrderRemoved                      []func(*dota.CDOTAUserMsg_QueuedOrderRemoved) error
	onCDOTAUserMsg_DebugChallenge                          []func(*dota.CDOTAUserMsg_DebugChallenge) error
	onCDOTAUserMsg_OMArcanaCombo                           []func(*dota.CDOTAUserMsg_OMArcanaCombo) error
	onCDOTAUserMsg_FoundNeutralItem                        []func(*dota.CDOTAUserMsg_FoundNeutralItem) error
	onCDOTAUserMsg_OutpostCaptured                         []func(*dota.CDOTAUserMsg_OutpostCaptured) error
	onCDOTAUserMsg_OutpostGrantedXP                        []func(*dota.CDOTAUserMsg_OutpostGrantedXP) error
	onCDOTAUserMsg_MoveCameraToUnit                        []func(*dota.CDOTAUserMsg_MoveCameraToUnit) error
	onCDOTAUserMsg_PauseMinigameData                       []func(*dota.CDOTAUserMsg_PauseMinigameData) error
	onCDOTAUserMsg_VersusScene_PlayerBehavior              []func(*dota.CDOTAUserMsg_VersusScene_PlayerBehavior) error
	onCDOTAUserMsg_QoP_ArcanaSummary                       []func(*dota.CDOTAUserMsg_QoP_ArcanaSummary) error
	onCDOTAUserMsg_HotPotato_Created                       []func(*dota.CDOTAUserMsg_HotPotato_Created) error
	onCDOTAUserMsg_HotPotato_Exploded                      []func(*dota.CDOTAUserMsg_HotPotato_Exploded) error
	onCDOTAUserMsg_WK_Arcana_Progress                      []func(*dota.CDOTAUserMsg_WK_Arcana_Progress) error
	onCDOTAUserMsg_GuildChallenge_Progress                 []func(*dota.CDOTAUserMsg_GuildChallenge_Progress) error
	onCDOTAUserMsg_WRArcanaProgress                        []func(*dota.CDOTAUserMsg_WRArcanaProgress) error
	onCDOTAUserMsg_WRArcanaSummary                         []func(*dota.CDOTAUserMsg_WRArcanaSummary) error
	onCDOTAUserMsg_EmptyItemSlotAlert                      []func(*dota.CDOTAUserMsg_EmptyItemSlotAlert) error
	onCDOTAUserMsg_AghsStatusAlert                         []func(*dota.CDOTAUserMsg_AghsStatusAlert) error
	onCDOTAUserMsg_PingConfirmation                        []func(*dota.CDOTAUserMsg_PingConfirmation) error
	onCDOTAUserMsg_MutedPlayers                            []func(*dota.CDOTAUserMsg_MutedPlayers) error
	onCDOTAUserMsg_ContextualTip                           []func(*dota.CDOTAUserMsg_ContextualTip) error
	onCDOTAUserMsg_ChatMessage                             []func(*dota.CDOTAUserMsg_ChatMessage) error
	onCDOTAUserMsg_NeutralCampAlert                        []func(*dota.CDOTAUserMsg_NeutralCampAlert) error
	onCDOTAUserMsg_RockPaperScissorsStarted                []func(*dota.CDOTAUserMsg_RockPaperScissorsStarted) error
	onCDOTAUserMsg_RockPaperScissorsFinished               []func(*dota.CDOTAUserMsg_RockPaperScissorsFinished) error
	onCDOTAUserMsg_DuelOpponentKilled                      []func(*dota.CDOTAUserMsg_DuelOpponentKilled) error
	onCDOTAUserMsg_DuelAccepted                            []func(*dota.CDOTAUserMsg_DuelAccepted) error
	onCDOTAUserMsg_DuelRequested                           []func(*dota.CDOTAUserMsg_DuelRequested) error
	onCDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled []func(*dota.CDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled) error
	onCDOTAUserMsg_PlayerDraftSuggestPick                  []func(*dota.CDOTAUserMsg_PlayerDraftSuggestPick) error
	onCDOTAUserMsg_PlayerDraftPick                         []func(*dota.CDOTAUserMsg_PlayerDraftPick) error
	onCDOTAUserMsg_UpdateLinearProjectileCPData            []func(*dota.CDOTAUserMsg_UpdateLinearProjectileCPData) error
	onCDOTAUserMsg_GiftPlayer                              []func(*dota.CDOTAUserMsg_GiftPlayer) error
	onCDOTAUserMsg_FacetPing                               []func(*dota.CDOTAUserMsg_FacetPing) error
	onCDOTAUserMsg_InnatePing                              []func(*dota.CDOTAUserMsg_InnatePing) error
	onCDOTAUserMsg_RoshanTimer                             []func(*dota.CDOTAUserMsg_RoshanTimer) error
	onCDOTAUserMsg_NeutralCraftAvailable                   []func(*dota.CDOTAUserMsg_NeutralCraftAvailable) error
	onCDOTAUserMsg_TimerAlert                              []func(*dota.CDOTAUserMsg_TimerAlert) error
	onCDOTAUserMsg_MadstoneAlert                           []func(*dota.CDOTAUserMsg_MadstoneAlert) error

	pb *proto.Buffer
}

func newCallbacks() *Callbacks {
	return &Callbacks{
		pb: &proto.Buffer{},
	}
}

// OnCDemoStop registers a callback EDemoCommands_DEM_Stop
func (c *Callbacks) OnCDemoStop(fn func(*dota.CDemoStop) error) {
	c.onCDemoStop = append(c.onCDemoStop, fn)
}

// OnCDemoFileHeader registers a callback EDemoCommands_DEM_FileHeader
func (c *Callbacks) OnCDemoFileHeader(fn func(*dota.CDemoFileHeader) error) {
	c.onCDemoFileHeader = append(c.onCDemoFileHeader, fn)
}

// OnCDemoFileInfo registers a callback EDemoCommands_DEM_FileInfo
func (c *Callbacks) OnCDemoFileInfo(fn func(*dota.CDemoFileInfo) error) {
	c.onCDemoFileInfo = append(c.onCDemoFileInfo, fn)
}

// OnCDemoSyncTick registers a callback EDemoCommands_DEM_SyncTick
func (c *Callbacks) OnCDemoSyncTick(fn func(*dota.CDemoSyncTick) error) {
	c.onCDemoSyncTick = append(c.onCDemoSyncTick, fn)
}

// OnCDemoSendTables registers a callback EDemoCommands_DEM_SendTables
func (c *Callbacks) OnCDemoSendTables(fn func(*dota.CDemoSendTables) error) {
	c.onCDemoSendTables = append(c.onCDemoSendTables, fn)
}

// OnCDemoClassInfo registers a callback EDemoCommands_DEM_ClassInfo
func (c *Callbacks) OnCDemoClassInfo(fn func(*dota.CDemoClassInfo) error) {
	c.onCDemoClassInfo = append(c.onCDemoClassInfo, fn)
}

// OnCDemoStringTables registers a callback EDemoCommands_DEM_StringTables
func (c *Callbacks) OnCDemoStringTables(fn func(*dota.CDemoStringTables) error) {
	c.onCDemoStringTables = append(c.onCDemoStringTables, fn)
}

// OnCDemoPacket registers a callback EDemoCommands_DEM_Packet
func (c *Callbacks) OnCDemoPacket(fn func(*dota.CDemoPacket) error) {
	c.onCDemoPacket = append(c.onCDemoPacket, fn)
}

// OnCDemoSignonPacket registers a callback EDemoCommands_DEM_SignonPacket
func (c *Callbacks) OnCDemoSignonPacket(fn func(*dota.CDemoPacket) error) {
	c.onCDemoSignonPacket = append(c.onCDemoSignonPacket, fn)
}

// OnCDemoConsoleCmd registers a callback EDemoCommands_DEM_ConsoleCmd
func (c *Callbacks) OnCDemoConsoleCmd(fn func(*dota.CDemoConsoleCmd) error) {
	c.onCDemoConsoleCmd = append(c.onCDemoConsoleCmd, fn)
}

// OnCDemoCustomData registers a callback EDemoCommands_DEM_CustomData
func (c *Callbacks) OnCDemoCustomData(fn func(*dota.CDemoCustomData) error) {
	c.onCDemoCustomData = append(c.onCDemoCustomData, fn)
}

// OnCDemoCustomDataCallbacks registers a callback EDemoCommands_DEM_CustomDataCallbacks
func (c *Callbacks) OnCDemoCustomDataCallbacks(fn func(*dota.CDemoCustomDataCallbacks) error) {
	c.onCDemoCustomDataCallbacks = append(c.onCDemoCustomDataCallbacks, fn)
}

// OnCDemoUserCmd registers a callback EDemoCommands_DEM_UserCmd
func (c *Callbacks) OnCDemoUserCmd(fn func(*dota.CDemoUserCmd) error) {
	c.onCDemoUserCmd = append(c.onCDemoUserCmd, fn)
}

// OnCDemoFullPacket registers a callback EDemoCommands_DEM_FullPacket
func (c *Callbacks) OnCDemoFullPacket(fn func(*dota.CDemoFullPacket) error) {
	c.onCDemoFullPacket = append(c.onCDemoFullPacket, fn)
}

// OnCDemoSaveGame registers a callback EDemoCommands_DEM_SaveGame
func (c *Callbacks) OnCDemoSaveGame(fn func(*dota.CDemoSaveGame) error) {
	c.onCDemoSaveGame = append(c.onCDemoSaveGame, fn)
}

// OnCDemoSpawnGroups registers a callback EDemoCommands_DEM_SpawnGroups
func (c *Callbacks) OnCDemoSpawnGroups(fn func(*dota.CDemoSpawnGroups) error) {
	c.onCDemoSpawnGroups = append(c.onCDemoSpawnGroups, fn)
}

// OnCDemoAnimationData registers a callback EDemoCommands_DEM_AnimationData
func (c *Callbacks) OnCDemoAnimationData(fn func(*dota.CDemoAnimationData) error) {
	c.onCDemoAnimationData = append(c.onCDemoAnimationData, fn)
}

// OnCDemoAnimationHeader registers a callback EDemoCommands_DEM_AnimationHeader
func (c *Callbacks) OnCDemoAnimationHeader(fn func(*dota.CDemoAnimationHeader) error) {
	c.onCDemoAnimationHeader = append(c.onCDemoAnimationHeader, fn)
}

// OnCDemoRecovery registers a callback EDemoCommands_DEM_Recovery
func (c *Callbacks) OnCDemoRecovery(fn func(*dota.CDemoRecovery) error) {
	c.onCDemoRecovery = append(c.onCDemoRecovery, fn)
}

// OnCNETMsg_NOP registers a callback for NET_Messages_net_NOP
func (c *Callbacks) OnCNETMsg_NOP(fn func(*dota.CNETMsg_NOP) error) {
	c.onCNETMsg_NOP = append(c.onCNETMsg_NOP, fn)
}

// OnCNETMsg_SplitScreenUser registers a callback for NET_Messages_net_SplitScreenUser
func (c *Callbacks) OnCNETMsg_SplitScreenUser(fn func(*dota.CNETMsg_SplitScreenUser) error) {
	c.onCNETMsg_SplitScreenUser = append(c.onCNETMsg_SplitScreenUser, fn)
}

// OnCNETMsg_Tick registers a callback for NET_Messages_net_Tick
func (c *Callbacks) OnCNETMsg_Tick(fn func(*dota.CNETMsg_Tick) error) {
	c.onCNETMsg_Tick = append(c.onCNETMsg_Tick, fn)
}

// OnCNETMsg_StringCmd registers a callback for NET_Messages_net_StringCmd
func (c *Callbacks) OnCNETMsg_StringCmd(fn func(*dota.CNETMsg_StringCmd) error) {
	c.onCNETMsg_StringCmd = append(c.onCNETMsg_StringCmd, fn)
}

// OnCNETMsg_SetConVar registers a callback for NET_Messages_net_SetConVar
func (c *Callbacks) OnCNETMsg_SetConVar(fn func(*dota.CNETMsg_SetConVar) error) {
	c.onCNETMsg_SetConVar = append(c.onCNETMsg_SetConVar, fn)
}

// OnCNETMsg_SignonState registers a callback for NET_Messages_net_SignonState
func (c *Callbacks) OnCNETMsg_SignonState(fn func(*dota.CNETMsg_SignonState) error) {
	c.onCNETMsg_SignonState = append(c.onCNETMsg_SignonState, fn)
}

// OnCNETMsg_SpawnGroup_Load registers a callback for NET_Messages_net_SpawnGroup_Load
func (c *Callbacks) OnCNETMsg_SpawnGroup_Load(fn func(*dota.CNETMsg_SpawnGroup_Load) error) {
	c.onCNETMsg_SpawnGroup_Load = append(c.onCNETMsg_SpawnGroup_Load, fn)
}

// OnCNETMsg_SpawnGroup_ManifestUpdate registers a callback for NET_Messages_net_SpawnGroup_ManifestUpdate
func (c *Callbacks) OnCNETMsg_SpawnGroup_ManifestUpdate(fn func(*dota.CNETMsg_SpawnGroup_ManifestUpdate) error) {
	c.onCNETMsg_SpawnGroup_ManifestUpdate = append(c.onCNETMsg_SpawnGroup_ManifestUpdate, fn)
}

// OnCNETMsg_SpawnGroup_SetCreationTick registers a callback for NET_Messages_net_SpawnGroup_SetCreationTick
func (c *Callbacks) OnCNETMsg_SpawnGroup_SetCreationTick(fn func(*dota.CNETMsg_SpawnGroup_SetCreationTick) error) {
	c.onCNETMsg_SpawnGroup_SetCreationTick = append(c.onCNETMsg_SpawnGroup_SetCreationTick, fn)
}

// OnCNETMsg_SpawnGroup_Unload registers a callback for NET_Messages_net_SpawnGroup_Unload
func (c *Callbacks) OnCNETMsg_SpawnGroup_Unload(fn func(*dota.CNETMsg_SpawnGroup_Unload) error) {
	c.onCNETMsg_SpawnGroup_Unload = append(c.onCNETMsg_SpawnGroup_Unload, fn)
}

// OnCNETMsg_SpawnGroup_LoadCompleted registers a callback for NET_Messages_net_SpawnGroup_LoadCompleted
func (c *Callbacks) OnCNETMsg_SpawnGroup_LoadCompleted(fn func(*dota.CNETMsg_SpawnGroup_LoadCompleted) error) {
	c.onCNETMsg_SpawnGroup_LoadCompleted = append(c.onCNETMsg_SpawnGroup_LoadCompleted, fn)
}

// OnCNETMsg_DebugOverlay registers a callback for NET_Messages_net_DebugOverlay
func (c *Callbacks) OnCNETMsg_DebugOverlay(fn func(*dota.CNETMsg_DebugOverlay) error) {
	c.onCNETMsg_DebugOverlay = append(c.onCNETMsg_DebugOverlay, fn)
}

// OnCSVCMsg_ServerInfo registers a callback for SVC_Messages_svc_ServerInfo
func (c *Callbacks) OnCSVCMsg_ServerInfo(fn func(*dota.CSVCMsg_ServerInfo) error) {
	c.onCSVCMsg_ServerInfo = append(c.onCSVCMsg_ServerInfo, fn)
}

// OnCSVCMsg_FlattenedSerializer registers a callback for SVC_Messages_svc_FlattenedSerializer
func (c *Callbacks) OnCSVCMsg_FlattenedSerializer(fn func(*dota.CSVCMsg_FlattenedSerializer) error) {
	c.onCSVCMsg_FlattenedSerializer = append(c.onCSVCMsg_FlattenedSerializer, fn)
}

// OnCSVCMsg_ClassInfo registers a callback for SVC_Messages_svc_ClassInfo
func (c *Callbacks) OnCSVCMsg_ClassInfo(fn func(*dota.CSVCMsg_ClassInfo) error) {
	c.onCSVCMsg_ClassInfo = append(c.onCSVCMsg_ClassInfo, fn)
}

// OnCSVCMsg_SetPause registers a callback for SVC_Messages_svc_SetPause
func (c *Callbacks) OnCSVCMsg_SetPause(fn func(*dota.CSVCMsg_SetPause) error) {
	c.onCSVCMsg_SetPause = append(c.onCSVCMsg_SetPause, fn)
}

// OnCSVCMsg_CreateStringTable registers a callback for SVC_Messages_svc_CreateStringTable
func (c *Callbacks) OnCSVCMsg_CreateStringTable(fn func(*dota.CSVCMsg_CreateStringTable) error) {
	c.onCSVCMsg_CreateStringTable = append(c.onCSVCMsg_CreateStringTable, fn)
}

// OnCSVCMsg_UpdateStringTable registers a callback for SVC_Messages_svc_UpdateStringTable
func (c *Callbacks) OnCSVCMsg_UpdateStringTable(fn func(*dota.CSVCMsg_UpdateStringTable) error) {
	c.onCSVCMsg_UpdateStringTable = append(c.onCSVCMsg_UpdateStringTable, fn)
}

// OnCSVCMsg_VoiceInit registers a callback for SVC_Messages_svc_VoiceInit
func (c *Callbacks) OnCSVCMsg_VoiceInit(fn func(*dota.CSVCMsg_VoiceInit) error) {
	c.onCSVCMsg_VoiceInit = append(c.onCSVCMsg_VoiceInit, fn)
}

// OnCSVCMsg_VoiceData registers a callback for SVC_Messages_svc_VoiceData
func (c *Callbacks) OnCSVCMsg_VoiceData(fn func(*dota.CSVCMsg_VoiceData) error) {
	c.onCSVCMsg_VoiceData = append(c.onCSVCMsg_VoiceData, fn)
}

// OnCSVCMsg_Print registers a callback for SVC_Messages_svc_Print
func (c *Callbacks) OnCSVCMsg_Print(fn func(*dota.CSVCMsg_Print) error) {
	c.onCSVCMsg_Print = append(c.onCSVCMsg_Print, fn)
}

// OnCSVCMsg_Sounds registers a callback for SVC_Messages_svc_Sounds
func (c *Callbacks) OnCSVCMsg_Sounds(fn func(*dota.CSVCMsg_Sounds) error) {
	c.onCSVCMsg_Sounds = append(c.onCSVCMsg_Sounds, fn)
}

// OnCSVCMsg_SetView registers a callback for SVC_Messages_svc_SetView
func (c *Callbacks) OnCSVCMsg_SetView(fn func(*dota.CSVCMsg_SetView) error) {
	c.onCSVCMsg_SetView = append(c.onCSVCMsg_SetView, fn)
}

// OnCSVCMsg_ClearAllStringTables registers a callback for SVC_Messages_svc_ClearAllStringTables
func (c *Callbacks) OnCSVCMsg_ClearAllStringTables(fn func(*dota.CSVCMsg_ClearAllStringTables) error) {
	c.onCSVCMsg_ClearAllStringTables = append(c.onCSVCMsg_ClearAllStringTables, fn)
}

// OnCSVCMsg_CmdKeyValues registers a callback for SVC_Messages_svc_CmdKeyValues
func (c *Callbacks) OnCSVCMsg_CmdKeyValues(fn func(*dota.CSVCMsg_CmdKeyValues) error) {
	c.onCSVCMsg_CmdKeyValues = append(c.onCSVCMsg_CmdKeyValues, fn)
}

// OnCSVCMsg_BSPDecal registers a callback for SVC_Messages_svc_BSPDecal
func (c *Callbacks) OnCSVCMsg_BSPDecal(fn func(*dota.CSVCMsg_BSPDecal) error) {
	c.onCSVCMsg_BSPDecal = append(c.onCSVCMsg_BSPDecal, fn)
}

// OnCSVCMsg_SplitScreen registers a callback for SVC_Messages_svc_SplitScreen
func (c *Callbacks) OnCSVCMsg_SplitScreen(fn func(*dota.CSVCMsg_SplitScreen) error) {
	c.onCSVCMsg_SplitScreen = append(c.onCSVCMsg_SplitScreen, fn)
}

// OnCSVCMsg_PacketEntities registers a callback for SVC_Messages_svc_PacketEntities
func (c *Callbacks) OnCSVCMsg_PacketEntities(fn func(*dota.CSVCMsg_PacketEntities) error) {
	c.onCSVCMsg_PacketEntities = append(c.onCSVCMsg_PacketEntities, fn)
}

// OnCSVCMsg_Prefetch registers a callback for SVC_Messages_svc_Prefetch
func (c *Callbacks) OnCSVCMsg_Prefetch(fn func(*dota.CSVCMsg_Prefetch) error) {
	c.onCSVCMsg_Prefetch = append(c.onCSVCMsg_Prefetch, fn)
}

// OnCSVCMsg_Menu registers a callback for SVC_Messages_svc_Menu
func (c *Callbacks) OnCSVCMsg_Menu(fn func(*dota.CSVCMsg_Menu) error) {
	c.onCSVCMsg_Menu = append(c.onCSVCMsg_Menu, fn)
}

// OnCSVCMsg_GetCvarValue registers a callback for SVC_Messages_svc_GetCvarValue
func (c *Callbacks) OnCSVCMsg_GetCvarValue(fn func(*dota.CSVCMsg_GetCvarValue) error) {
	c.onCSVCMsg_GetCvarValue = append(c.onCSVCMsg_GetCvarValue, fn)
}

// OnCSVCMsg_StopSound registers a callback for SVC_Messages_svc_StopSound
func (c *Callbacks) OnCSVCMsg_StopSound(fn func(*dota.CSVCMsg_StopSound) error) {
	c.onCSVCMsg_StopSound = append(c.onCSVCMsg_StopSound, fn)
}

// OnCSVCMsg_PeerList registers a callback for SVC_Messages_svc_PeerList
func (c *Callbacks) OnCSVCMsg_PeerList(fn func(*dota.CSVCMsg_PeerList) error) {
	c.onCSVCMsg_PeerList = append(c.onCSVCMsg_PeerList, fn)
}

// OnCSVCMsg_PacketReliable registers a callback for SVC_Messages_svc_PacketReliable
func (c *Callbacks) OnCSVCMsg_PacketReliable(fn func(*dota.CSVCMsg_PacketReliable) error) {
	c.onCSVCMsg_PacketReliable = append(c.onCSVCMsg_PacketReliable, fn)
}

// OnCSVCMsg_HLTVStatus registers a callback for SVC_Messages_svc_HLTVStatus
func (c *Callbacks) OnCSVCMsg_HLTVStatus(fn func(*dota.CSVCMsg_HLTVStatus) error) {
	c.onCSVCMsg_HLTVStatus = append(c.onCSVCMsg_HLTVStatus, fn)
}

// OnCSVCMsg_ServerSteamID registers a callback for SVC_Messages_svc_ServerSteamID
func (c *Callbacks) OnCSVCMsg_ServerSteamID(fn func(*dota.CSVCMsg_ServerSteamID) error) {
	c.onCSVCMsg_ServerSteamID = append(c.onCSVCMsg_ServerSteamID, fn)
}

// OnCSVCMsg_FullFrameSplit registers a callback for SVC_Messages_svc_FullFrameSplit
func (c *Callbacks) OnCSVCMsg_FullFrameSplit(fn func(*dota.CSVCMsg_FullFrameSplit) error) {
	c.onCSVCMsg_FullFrameSplit = append(c.onCSVCMsg_FullFrameSplit, fn)
}

// OnCSVCMsg_RconServerDetails registers a callback for SVC_Messages_svc_RconServerDetails
func (c *Callbacks) OnCSVCMsg_RconServerDetails(fn func(*dota.CSVCMsg_RconServerDetails) error) {
	c.onCSVCMsg_RconServerDetails = append(c.onCSVCMsg_RconServerDetails, fn)
}

// OnCSVCMsg_UserMessage registers a callback for SVC_Messages_svc_UserMessage
func (c *Callbacks) OnCSVCMsg_UserMessage(fn func(*dota.CSVCMsg_UserMessage) error) {
	c.onCSVCMsg_UserMessage = append(c.onCSVCMsg_UserMessage, fn)
}

// OnCSVCMsg_Broadcast_Command registers a callback for SVC_Messages_svc_Broadcast_Command
func (c *Callbacks) OnCSVCMsg_Broadcast_Command(fn func(*dota.CSVCMsg_Broadcast_Command) error) {
	c.onCSVCMsg_Broadcast_Command = append(c.onCSVCMsg_Broadcast_Command, fn)
}

// OnCSVCMsg_HltvFixupOperatorStatus registers a callback for SVC_Messages_svc_HltvFixupOperatorStatus
func (c *Callbacks) OnCSVCMsg_HltvFixupOperatorStatus(fn func(*dota.CSVCMsg_HltvFixupOperatorStatus) error) {
	c.onCSVCMsg_HltvFixupOperatorStatus = append(c.onCSVCMsg_HltvFixupOperatorStatus, fn)
}

// OnCUserMessageAchievementEvent registers a callback for EBaseUserMessages_UM_AchievementEvent
func (c *Callbacks) OnCUserMessageAchievementEvent(fn func(*dota.CUserMessageAchievementEvent) error) {
	c.onCUserMessageAchievementEvent = append(c.onCUserMessageAchievementEvent, fn)
}

// OnCUserMessageCloseCaption registers a callback for EBaseUserMessages_UM_CloseCaption
func (c *Callbacks) OnCUserMessageCloseCaption(fn func(*dota.CUserMessageCloseCaption) error) {
	c.onCUserMessageCloseCaption = append(c.onCUserMessageCloseCaption, fn)
}

// OnCUserMessageCloseCaptionDirect registers a callback for EBaseUserMessages_UM_CloseCaptionDirect
func (c *Callbacks) OnCUserMessageCloseCaptionDirect(fn func(*dota.CUserMessageCloseCaptionDirect) error) {
	c.onCUserMessageCloseCaptionDirect = append(c.onCUserMessageCloseCaptionDirect, fn)
}

// OnCUserMessageCurrentTimescale registers a callback for EBaseUserMessages_UM_CurrentTimescale
func (c *Callbacks) OnCUserMessageCurrentTimescale(fn func(*dota.CUserMessageCurrentTimescale) error) {
	c.onCUserMessageCurrentTimescale = append(c.onCUserMessageCurrentTimescale, fn)
}

// OnCUserMessageDesiredTimescale registers a callback for EBaseUserMessages_UM_DesiredTimescale
func (c *Callbacks) OnCUserMessageDesiredTimescale(fn func(*dota.CUserMessageDesiredTimescale) error) {
	c.onCUserMessageDesiredTimescale = append(c.onCUserMessageDesiredTimescale, fn)
}

// OnCUserMessageFade registers a callback for EBaseUserMessages_UM_Fade
func (c *Callbacks) OnCUserMessageFade(fn func(*dota.CUserMessageFade) error) {
	c.onCUserMessageFade = append(c.onCUserMessageFade, fn)
}

// OnCUserMessageGameTitle registers a callback for EBaseUserMessages_UM_GameTitle
func (c *Callbacks) OnCUserMessageGameTitle(fn func(*dota.CUserMessageGameTitle) error) {
	c.onCUserMessageGameTitle = append(c.onCUserMessageGameTitle, fn)
}

// OnCUserMessageHudMsg registers a callback for EBaseUserMessages_UM_HudMsg
func (c *Callbacks) OnCUserMessageHudMsg(fn func(*dota.CUserMessageHudMsg) error) {
	c.onCUserMessageHudMsg = append(c.onCUserMessageHudMsg, fn)
}

// OnCUserMessageHudText registers a callback for EBaseUserMessages_UM_HudText
func (c *Callbacks) OnCUserMessageHudText(fn func(*dota.CUserMessageHudText) error) {
	c.onCUserMessageHudText = append(c.onCUserMessageHudText, fn)
}

// OnCUserMessageColoredText registers a callback for EBaseUserMessages_UM_ColoredText
func (c *Callbacks) OnCUserMessageColoredText(fn func(*dota.CUserMessageColoredText) error) {
	c.onCUserMessageColoredText = append(c.onCUserMessageColoredText, fn)
}

// OnCUserMessageRequestState registers a callback for EBaseUserMessages_UM_RequestState
func (c *Callbacks) OnCUserMessageRequestState(fn func(*dota.CUserMessageRequestState) error) {
	c.onCUserMessageRequestState = append(c.onCUserMessageRequestState, fn)
}

// OnCUserMessageResetHUD registers a callback for EBaseUserMessages_UM_ResetHUD
func (c *Callbacks) OnCUserMessageResetHUD(fn func(*dota.CUserMessageResetHUD) error) {
	c.onCUserMessageResetHUD = append(c.onCUserMessageResetHUD, fn)
}

// OnCUserMessageRumble registers a callback for EBaseUserMessages_UM_Rumble
func (c *Callbacks) OnCUserMessageRumble(fn func(*dota.CUserMessageRumble) error) {
	c.onCUserMessageRumble = append(c.onCUserMessageRumble, fn)
}

// OnCUserMessageSayText registers a callback for EBaseUserMessages_UM_SayText
func (c *Callbacks) OnCUserMessageSayText(fn func(*dota.CUserMessageSayText) error) {
	c.onCUserMessageSayText = append(c.onCUserMessageSayText, fn)
}

// OnCUserMessageSayText2 registers a callback for EBaseUserMessages_UM_SayText2
func (c *Callbacks) OnCUserMessageSayText2(fn func(*dota.CUserMessageSayText2) error) {
	c.onCUserMessageSayText2 = append(c.onCUserMessageSayText2, fn)
}

// OnCUserMessageSayTextChannel registers a callback for EBaseUserMessages_UM_SayTextChannel
func (c *Callbacks) OnCUserMessageSayTextChannel(fn func(*dota.CUserMessageSayTextChannel) error) {
	c.onCUserMessageSayTextChannel = append(c.onCUserMessageSayTextChannel, fn)
}

// OnCUserMessageShake registers a callback for EBaseUserMessages_UM_Shake
func (c *Callbacks) OnCUserMessageShake(fn func(*dota.CUserMessageShake) error) {
	c.onCUserMessageShake = append(c.onCUserMessageShake, fn)
}

// OnCUserMessageShakeDir registers a callback for EBaseUserMessages_UM_ShakeDir
func (c *Callbacks) OnCUserMessageShakeDir(fn func(*dota.CUserMessageShakeDir) error) {
	c.onCUserMessageShakeDir = append(c.onCUserMessageShakeDir, fn)
}

// OnCUserMessageWaterShake registers a callback for EBaseUserMessages_UM_WaterShake
func (c *Callbacks) OnCUserMessageWaterShake(fn func(*dota.CUserMessageWaterShake) error) {
	c.onCUserMessageWaterShake = append(c.onCUserMessageWaterShake, fn)
}

// OnCUserMessageTextMsg registers a callback for EBaseUserMessages_UM_TextMsg
func (c *Callbacks) OnCUserMessageTextMsg(fn func(*dota.CUserMessageTextMsg) error) {
	c.onCUserMessageTextMsg = append(c.onCUserMessageTextMsg, fn)
}

// OnCUserMessageScreenTilt registers a callback for EBaseUserMessages_UM_ScreenTilt
func (c *Callbacks) OnCUserMessageScreenTilt(fn func(*dota.CUserMessageScreenTilt) error) {
	c.onCUserMessageScreenTilt = append(c.onCUserMessageScreenTilt, fn)
}

// OnCUserMessageVoiceMask registers a callback for EBaseUserMessages_UM_VoiceMask
func (c *Callbacks) OnCUserMessageVoiceMask(fn func(*dota.CUserMessageVoiceMask) error) {
	c.onCUserMessageVoiceMask = append(c.onCUserMessageVoiceMask, fn)
}

// OnCUserMessageSendAudio registers a callback for EBaseUserMessages_UM_SendAudio
func (c *Callbacks) OnCUserMessageSendAudio(fn func(*dota.CUserMessageSendAudio) error) {
	c.onCUserMessageSendAudio = append(c.onCUserMessageSendAudio, fn)
}

// OnCUserMessageItemPickup registers a callback for EBaseUserMessages_UM_ItemPickup
func (c *Callbacks) OnCUserMessageItemPickup(fn func(*dota.CUserMessageItemPickup) error) {
	c.onCUserMessageItemPickup = append(c.onCUserMessageItemPickup, fn)
}

// OnCUserMessageAmmoDenied registers a callback for EBaseUserMessages_UM_AmmoDenied
func (c *Callbacks) OnCUserMessageAmmoDenied(fn func(*dota.CUserMessageAmmoDenied) error) {
	c.onCUserMessageAmmoDenied = append(c.onCUserMessageAmmoDenied, fn)
}

// OnCUserMessageShowMenu registers a callback for EBaseUserMessages_UM_ShowMenu
func (c *Callbacks) OnCUserMessageShowMenu(fn func(*dota.CUserMessageShowMenu) error) {
	c.onCUserMessageShowMenu = append(c.onCUserMessageShowMenu, fn)
}

// OnCUserMessageCreditsMsg registers a callback for EBaseUserMessages_UM_CreditsMsg
func (c *Callbacks) OnCUserMessageCreditsMsg(fn func(*dota.CUserMessageCreditsMsg) error) {
	c.onCUserMessageCreditsMsg = append(c.onCUserMessageCreditsMsg, fn)
}

// OnCEntityMessagePlayJingle registers a callback for EBaseEntityMessages_EM_PlayJingle
func (c *Callbacks) OnCEntityMessagePlayJingle(fn func(*dota.CEntityMessagePlayJingle) error) {
	c.onCEntityMessagePlayJingle = append(c.onCEntityMessagePlayJingle, fn)
}

// OnCEntityMessageScreenOverlay registers a callback for EBaseEntityMessages_EM_ScreenOverlay
func (c *Callbacks) OnCEntityMessageScreenOverlay(fn func(*dota.CEntityMessageScreenOverlay) error) {
	c.onCEntityMessageScreenOverlay = append(c.onCEntityMessageScreenOverlay, fn)
}

// OnCEntityMessageRemoveAllDecals registers a callback for EBaseEntityMessages_EM_RemoveAllDecals
func (c *Callbacks) OnCEntityMessageRemoveAllDecals(fn func(*dota.CEntityMessageRemoveAllDecals) error) {
	c.onCEntityMessageRemoveAllDecals = append(c.onCEntityMessageRemoveAllDecals, fn)
}

// OnCEntityMessagePropagateForce registers a callback for EBaseEntityMessages_EM_PropagateForce
func (c *Callbacks) OnCEntityMessagePropagateForce(fn func(*dota.CEntityMessagePropagateForce) error) {
	c.onCEntityMessagePropagateForce = append(c.onCEntityMessagePropagateForce, fn)
}

// OnCEntityMessageDoSpark registers a callback for EBaseEntityMessages_EM_DoSpark
func (c *Callbacks) OnCEntityMessageDoSpark(fn func(*dota.CEntityMessageDoSpark) error) {
	c.onCEntityMessageDoSpark = append(c.onCEntityMessageDoSpark, fn)
}

// OnCEntityMessageFixAngle registers a callback for EBaseEntityMessages_EM_FixAngle
func (c *Callbacks) OnCEntityMessageFixAngle(fn func(*dota.CEntityMessageFixAngle) error) {
	c.onCEntityMessageFixAngle = append(c.onCEntityMessageFixAngle, fn)
}

// OnCUserMessageCloseCaptionPlaceholder registers a callback for EBaseUserMessages_UM_CloseCaptionPlaceholder
func (c *Callbacks) OnCUserMessageCloseCaptionPlaceholder(fn func(*dota.CUserMessageCloseCaptionPlaceholder) error) {
	c.onCUserMessageCloseCaptionPlaceholder = append(c.onCUserMessageCloseCaptionPlaceholder, fn)
}

// OnCUserMessageCameraTransition registers a callback for EBaseUserMessages_UM_CameraTransition
func (c *Callbacks) OnCUserMessageCameraTransition(fn func(*dota.CUserMessageCameraTransition) error) {
	c.onCUserMessageCameraTransition = append(c.onCUserMessageCameraTransition, fn)
}

// OnCUserMessageAudioParameter registers a callback for EBaseUserMessages_UM_AudioParameter
func (c *Callbacks) OnCUserMessageAudioParameter(fn func(*dota.CUserMessageAudioParameter) error) {
	c.onCUserMessageAudioParameter = append(c.onCUserMessageAudioParameter, fn)
}

// OnCUserMessageHapticsManagerPulse registers a callback for EBaseUserMessages_UM_HapticsManagerPulse
func (c *Callbacks) OnCUserMessageHapticsManagerPulse(fn func(*dota.CUserMessageHapticsManagerPulse) error) {
	c.onCUserMessageHapticsManagerPulse = append(c.onCUserMessageHapticsManagerPulse, fn)
}

// OnCUserMessageHapticsManagerEffect registers a callback for EBaseUserMessages_UM_HapticsManagerEffect
func (c *Callbacks) OnCUserMessageHapticsManagerEffect(fn func(*dota.CUserMessageHapticsManagerEffect) error) {
	c.onCUserMessageHapticsManagerEffect = append(c.onCUserMessageHapticsManagerEffect, fn)
}

// OnCUserMessageUpdateCssClasses registers a callback for EBaseUserMessages_UM_UpdateCssClasses
func (c *Callbacks) OnCUserMessageUpdateCssClasses(fn func(*dota.CUserMessageUpdateCssClasses) error) {
	c.onCUserMessageUpdateCssClasses = append(c.onCUserMessageUpdateCssClasses, fn)
}

// OnCUserMessageServerFrameTime registers a callback for EBaseUserMessages_UM_ServerFrameTime
func (c *Callbacks) OnCUserMessageServerFrameTime(fn func(*dota.CUserMessageServerFrameTime) error) {
	c.onCUserMessageServerFrameTime = append(c.onCUserMessageServerFrameTime, fn)
}

// OnCUserMessageLagCompensationError registers a callback for EBaseUserMessages_UM_LagCompensationError
func (c *Callbacks) OnCUserMessageLagCompensationError(fn func(*dota.CUserMessageLagCompensationError) error) {
	c.onCUserMessageLagCompensationError = append(c.onCUserMessageLagCompensationError, fn)
}

// OnCUserMessageRequestDllStatus registers a callback for EBaseUserMessages_UM_RequestDllStatus
func (c *Callbacks) OnCUserMessageRequestDllStatus(fn func(*dota.CUserMessageRequestDllStatus) error) {
	c.onCUserMessageRequestDllStatus = append(c.onCUserMessageRequestDllStatus, fn)
}

// OnCUserMessageRequestUtilAction registers a callback for EBaseUserMessages_UM_RequestUtilAction
func (c *Callbacks) OnCUserMessageRequestUtilAction(fn func(*dota.CUserMessageRequestUtilAction) error) {
	c.onCUserMessageRequestUtilAction = append(c.onCUserMessageRequestUtilAction, fn)
}

// OnCUserMessageRequestInventory registers a callback for EBaseUserMessages_UM_RequestInventory
func (c *Callbacks) OnCUserMessageRequestInventory(fn func(*dota.CUserMessageRequestInventory) error) {
	c.onCUserMessageRequestInventory = append(c.onCUserMessageRequestInventory, fn)
}

// OnCUserMessageRequestDiagnostic registers a callback for EBaseUserMessages_UM_RequestDiagnostic
func (c *Callbacks) OnCUserMessageRequestDiagnostic(fn func(*dota.CUserMessageRequestDiagnostic) error) {
	c.onCUserMessageRequestDiagnostic = append(c.onCUserMessageRequestDiagnostic, fn)
}

// OnCMsgVDebugGameSessionIDEvent registers a callback for EBaseGameEvents_GE_VDebugGameSessionIDEvent
func (c *Callbacks) OnCMsgVDebugGameSessionIDEvent(fn func(*dota.CMsgVDebugGameSessionIDEvent) error) {
	c.onCMsgVDebugGameSessionIDEvent = append(c.onCMsgVDebugGameSessionIDEvent, fn)
}

// OnCMsgPlaceDecalEvent registers a callback for EBaseGameEvents_GE_PlaceDecalEvent
func (c *Callbacks) OnCMsgPlaceDecalEvent(fn func(*dota.CMsgPlaceDecalEvent) error) {
	c.onCMsgPlaceDecalEvent = append(c.onCMsgPlaceDecalEvent, fn)
}

// OnCMsgClearWorldDecalsEvent registers a callback for EBaseGameEvents_GE_ClearWorldDecalsEvent
func (c *Callbacks) OnCMsgClearWorldDecalsEvent(fn func(*dota.CMsgClearWorldDecalsEvent) error) {
	c.onCMsgClearWorldDecalsEvent = append(c.onCMsgClearWorldDecalsEvent, fn)
}

// OnCMsgClearEntityDecalsEvent registers a callback for EBaseGameEvents_GE_ClearEntityDecalsEvent
func (c *Callbacks) OnCMsgClearEntityDecalsEvent(fn func(*dota.CMsgClearEntityDecalsEvent) error) {
	c.onCMsgClearEntityDecalsEvent = append(c.onCMsgClearEntityDecalsEvent, fn)
}

// OnCMsgClearDecalsForSkeletonInstanceEvent registers a callback for EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
func (c *Callbacks) OnCMsgClearDecalsForSkeletonInstanceEvent(fn func(*dota.CMsgClearDecalsForSkeletonInstanceEvent) error) {
	c.onCMsgClearDecalsForSkeletonInstanceEvent = append(c.onCMsgClearDecalsForSkeletonInstanceEvent, fn)
}

// OnCMsgSource1LegacyGameEventList registers a callback for EBaseGameEvents_GE_Source1LegacyGameEventList
func (c *Callbacks) OnCMsgSource1LegacyGameEventList(fn func(*dota.CMsgSource1LegacyGameEventList) error) {
	c.onCMsgSource1LegacyGameEventList = append(c.onCMsgSource1LegacyGameEventList, fn)
}

// OnCMsgSource1LegacyListenEvents registers a callback for EBaseGameEvents_GE_Source1LegacyListenEvents
func (c *Callbacks) OnCMsgSource1LegacyListenEvents(fn func(*dota.CMsgSource1LegacyListenEvents) error) {
	c.onCMsgSource1LegacyListenEvents = append(c.onCMsgSource1LegacyListenEvents, fn)
}

// OnCMsgSource1LegacyGameEvent registers a callback for EBaseGameEvents_GE_Source1LegacyGameEvent
func (c *Callbacks) OnCMsgSource1LegacyGameEvent(fn func(*dota.CMsgSource1LegacyGameEvent) error) {
	c.onCMsgSource1LegacyGameEvent = append(c.onCMsgSource1LegacyGameEvent, fn)
}

// OnCMsgSosStartSoundEvent registers a callback for EBaseGameEvents_GE_SosStartSoundEvent
func (c *Callbacks) OnCMsgSosStartSoundEvent(fn func(*dota.CMsgSosStartSoundEvent) error) {
	c.onCMsgSosStartSoundEvent = append(c.onCMsgSosStartSoundEvent, fn)
}

// OnCMsgSosStopSoundEvent registers a callback for EBaseGameEvents_GE_SosStopSoundEvent
func (c *Callbacks) OnCMsgSosStopSoundEvent(fn func(*dota.CMsgSosStopSoundEvent) error) {
	c.onCMsgSosStopSoundEvent = append(c.onCMsgSosStopSoundEvent, fn)
}

// OnCMsgSosSetSoundEventParams registers a callback for EBaseGameEvents_GE_SosSetSoundEventParams
func (c *Callbacks) OnCMsgSosSetSoundEventParams(fn func(*dota.CMsgSosSetSoundEventParams) error) {
	c.onCMsgSosSetSoundEventParams = append(c.onCMsgSosSetSoundEventParams, fn)
}

// OnCMsgSosSetLibraryStackFields registers a callback for EBaseGameEvents_GE_SosSetLibraryStackFields
func (c *Callbacks) OnCMsgSosSetLibraryStackFields(fn func(*dota.CMsgSosSetLibraryStackFields) error) {
	c.onCMsgSosSetLibraryStackFields = append(c.onCMsgSosSetLibraryStackFields, fn)
}

// OnCMsgSosStopSoundEventHash registers a callback for EBaseGameEvents_GE_SosStopSoundEventHash
func (c *Callbacks) OnCMsgSosStopSoundEventHash(fn func(*dota.CMsgSosStopSoundEventHash) error) {
	c.onCMsgSosStopSoundEventHash = append(c.onCMsgSosStopSoundEventHash, fn)
}

// OnCDOTAUserMsg_AIDebugLine registers a callback for EDotaUserMessages_DOTA_UM_AIDebugLine
func (c *Callbacks) OnCDOTAUserMsg_AIDebugLine(fn func(*dota.CDOTAUserMsg_AIDebugLine) error) {
	c.onCDOTAUserMsg_AIDebugLine = append(c.onCDOTAUserMsg_AIDebugLine, fn)
}

// OnCDOTAUserMsg_ChatEvent registers a callback for EDotaUserMessages_DOTA_UM_ChatEvent
func (c *Callbacks) OnCDOTAUserMsg_ChatEvent(fn func(*dota.CDOTAUserMsg_ChatEvent) error) {
	c.onCDOTAUserMsg_ChatEvent = append(c.onCDOTAUserMsg_ChatEvent, fn)
}

// OnCDOTAUserMsg_CombatHeroPositions registers a callback for EDotaUserMessages_DOTA_UM_CombatHeroPositions
func (c *Callbacks) OnCDOTAUserMsg_CombatHeroPositions(fn func(*dota.CDOTAUserMsg_CombatHeroPositions) error) {
	c.onCDOTAUserMsg_CombatHeroPositions = append(c.onCDOTAUserMsg_CombatHeroPositions, fn)
}

// OnCDOTAUserMsg_CombatLogBulkData registers a callback for EDotaUserMessages_DOTA_UM_CombatLogBulkData
func (c *Callbacks) OnCDOTAUserMsg_CombatLogBulkData(fn func(*dota.CDOTAUserMsg_CombatLogBulkData) error) {
	c.onCDOTAUserMsg_CombatLogBulkData = append(c.onCDOTAUserMsg_CombatLogBulkData, fn)
}

// OnCDOTAUserMsg_CreateLinearProjectile registers a callback for EDotaUserMessages_DOTA_UM_CreateLinearProjectile
func (c *Callbacks) OnCDOTAUserMsg_CreateLinearProjectile(fn func(*dota.CDOTAUserMsg_CreateLinearProjectile) error) {
	c.onCDOTAUserMsg_CreateLinearProjectile = append(c.onCDOTAUserMsg_CreateLinearProjectile, fn)
}

// OnCDOTAUserMsg_DestroyLinearProjectile registers a callback for EDotaUserMessages_DOTA_UM_DestroyLinearProjectile
func (c *Callbacks) OnCDOTAUserMsg_DestroyLinearProjectile(fn func(*dota.CDOTAUserMsg_DestroyLinearProjectile) error) {
	c.onCDOTAUserMsg_DestroyLinearProjectile = append(c.onCDOTAUserMsg_DestroyLinearProjectile, fn)
}

// OnCDOTAUserMsg_DodgeTrackingProjectiles registers a callback for EDotaUserMessages_DOTA_UM_DodgeTrackingProjectiles
func (c *Callbacks) OnCDOTAUserMsg_DodgeTrackingProjectiles(fn func(*dota.CDOTAUserMsg_DodgeTrackingProjectiles) error) {
	c.onCDOTAUserMsg_DodgeTrackingProjectiles = append(c.onCDOTAUserMsg_DodgeTrackingProjectiles, fn)
}

// OnCDOTAUserMsg_GlobalLightColor registers a callback for EDotaUserMessages_DOTA_UM_GlobalLightColor
func (c *Callbacks) OnCDOTAUserMsg_GlobalLightColor(fn func(*dota.CDOTAUserMsg_GlobalLightColor) error) {
	c.onCDOTAUserMsg_GlobalLightColor = append(c.onCDOTAUserMsg_GlobalLightColor, fn)
}

// OnCDOTAUserMsg_GlobalLightDirection registers a callback for EDotaUserMessages_DOTA_UM_GlobalLightDirection
func (c *Callbacks) OnCDOTAUserMsg_GlobalLightDirection(fn func(*dota.CDOTAUserMsg_GlobalLightDirection) error) {
	c.onCDOTAUserMsg_GlobalLightDirection = append(c.onCDOTAUserMsg_GlobalLightDirection, fn)
}

// OnCDOTAUserMsg_InvalidCommand registers a callback for EDotaUserMessages_DOTA_UM_InvalidCommand
func (c *Callbacks) OnCDOTAUserMsg_InvalidCommand(fn func(*dota.CDOTAUserMsg_InvalidCommand) error) {
	c.onCDOTAUserMsg_InvalidCommand = append(c.onCDOTAUserMsg_InvalidCommand, fn)
}

// OnCDOTAUserMsg_LocationPing registers a callback for EDotaUserMessages_DOTA_UM_LocationPing
func (c *Callbacks) OnCDOTAUserMsg_LocationPing(fn func(*dota.CDOTAUserMsg_LocationPing) error) {
	c.onCDOTAUserMsg_LocationPing = append(c.onCDOTAUserMsg_LocationPing, fn)
}

// OnCDOTAUserMsg_MapLine registers a callback for EDotaUserMessages_DOTA_UM_MapLine
func (c *Callbacks) OnCDOTAUserMsg_MapLine(fn func(*dota.CDOTAUserMsg_MapLine) error) {
	c.onCDOTAUserMsg_MapLine = append(c.onCDOTAUserMsg_MapLine, fn)
}

// OnCDOTAUserMsg_MiniKillCamInfo registers a callback for EDotaUserMessages_DOTA_UM_MiniKillCamInfo
func (c *Callbacks) OnCDOTAUserMsg_MiniKillCamInfo(fn func(*dota.CDOTAUserMsg_MiniKillCamInfo) error) {
	c.onCDOTAUserMsg_MiniKillCamInfo = append(c.onCDOTAUserMsg_MiniKillCamInfo, fn)
}

// OnCDOTAUserMsg_MinimapDebugPoint registers a callback for EDotaUserMessages_DOTA_UM_MinimapDebugPoint
func (c *Callbacks) OnCDOTAUserMsg_MinimapDebugPoint(fn func(*dota.CDOTAUserMsg_MinimapDebugPoint) error) {
	c.onCDOTAUserMsg_MinimapDebugPoint = append(c.onCDOTAUserMsg_MinimapDebugPoint, fn)
}

// OnCDOTAUserMsg_MinimapEvent registers a callback for EDotaUserMessages_DOTA_UM_MinimapEvent
func (c *Callbacks) OnCDOTAUserMsg_MinimapEvent(fn func(*dota.CDOTAUserMsg_MinimapEvent) error) {
	c.onCDOTAUserMsg_MinimapEvent = append(c.onCDOTAUserMsg_MinimapEvent, fn)
}

// OnCDOTAUserMsg_NevermoreRequiem registers a callback for EDotaUserMessages_DOTA_UM_NevermoreRequiem
func (c *Callbacks) OnCDOTAUserMsg_NevermoreRequiem(fn func(*dota.CDOTAUserMsg_NevermoreRequiem) error) {
	c.onCDOTAUserMsg_NevermoreRequiem = append(c.onCDOTAUserMsg_NevermoreRequiem, fn)
}

// OnCDOTAUserMsg_OverheadEvent registers a callback for EDotaUserMessages_DOTA_UM_OverheadEvent
func (c *Callbacks) OnCDOTAUserMsg_OverheadEvent(fn func(*dota.CDOTAUserMsg_OverheadEvent) error) {
	c.onCDOTAUserMsg_OverheadEvent = append(c.onCDOTAUserMsg_OverheadEvent, fn)
}

// OnCDOTAUserMsg_SetNextAutobuyItem registers a callback for EDotaUserMessages_DOTA_UM_SetNextAutobuyItem
func (c *Callbacks) OnCDOTAUserMsg_SetNextAutobuyItem(fn func(*dota.CDOTAUserMsg_SetNextAutobuyItem) error) {
	c.onCDOTAUserMsg_SetNextAutobuyItem = append(c.onCDOTAUserMsg_SetNextAutobuyItem, fn)
}

// OnCDOTAUserMsg_SharedCooldown registers a callback for EDotaUserMessages_DOTA_UM_SharedCooldown
func (c *Callbacks) OnCDOTAUserMsg_SharedCooldown(fn func(*dota.CDOTAUserMsg_SharedCooldown) error) {
	c.onCDOTAUserMsg_SharedCooldown = append(c.onCDOTAUserMsg_SharedCooldown, fn)
}

// OnCDOTAUserMsg_SpectatorPlayerClick registers a callback for EDotaUserMessages_DOTA_UM_SpectatorPlayerClick
func (c *Callbacks) OnCDOTAUserMsg_SpectatorPlayerClick(fn func(*dota.CDOTAUserMsg_SpectatorPlayerClick) error) {
	c.onCDOTAUserMsg_SpectatorPlayerClick = append(c.onCDOTAUserMsg_SpectatorPlayerClick, fn)
}

// OnCDOTAUserMsg_TutorialTipInfo registers a callback for EDotaUserMessages_DOTA_UM_TutorialTipInfo
func (c *Callbacks) OnCDOTAUserMsg_TutorialTipInfo(fn func(*dota.CDOTAUserMsg_TutorialTipInfo) error) {
	c.onCDOTAUserMsg_TutorialTipInfo = append(c.onCDOTAUserMsg_TutorialTipInfo, fn)
}

// OnCDOTAUserMsg_UnitEvent registers a callback for EDotaUserMessages_DOTA_UM_UnitEvent
func (c *Callbacks) OnCDOTAUserMsg_UnitEvent(fn func(*dota.CDOTAUserMsg_UnitEvent) error) {
	c.onCDOTAUserMsg_UnitEvent = append(c.onCDOTAUserMsg_UnitEvent, fn)
}

// OnCDOTAUserMsg_BotChat registers a callback for EDotaUserMessages_DOTA_UM_BotChat
func (c *Callbacks) OnCDOTAUserMsg_BotChat(fn func(*dota.CDOTAUserMsg_BotChat) error) {
	c.onCDOTAUserMsg_BotChat = append(c.onCDOTAUserMsg_BotChat, fn)
}

// OnCDOTAUserMsg_HudError registers a callback for EDotaUserMessages_DOTA_UM_HudError
func (c *Callbacks) OnCDOTAUserMsg_HudError(fn func(*dota.CDOTAUserMsg_HudError) error) {
	c.onCDOTAUserMsg_HudError = append(c.onCDOTAUserMsg_HudError, fn)
}

// OnCDOTAUserMsg_ItemPurchased registers a callback for EDotaUserMessages_DOTA_UM_ItemPurchased
func (c *Callbacks) OnCDOTAUserMsg_ItemPurchased(fn func(*dota.CDOTAUserMsg_ItemPurchased) error) {
	c.onCDOTAUserMsg_ItemPurchased = append(c.onCDOTAUserMsg_ItemPurchased, fn)
}

// OnCDOTAUserMsg_Ping registers a callback for EDotaUserMessages_DOTA_UM_Ping
func (c *Callbacks) OnCDOTAUserMsg_Ping(fn func(*dota.CDOTAUserMsg_Ping) error) {
	c.onCDOTAUserMsg_Ping = append(c.onCDOTAUserMsg_Ping, fn)
}

// OnCDOTAUserMsg_ItemFound registers a callback for EDotaUserMessages_DOTA_UM_ItemFound
func (c *Callbacks) OnCDOTAUserMsg_ItemFound(fn func(*dota.CDOTAUserMsg_ItemFound) error) {
	c.onCDOTAUserMsg_ItemFound = append(c.onCDOTAUserMsg_ItemFound, fn)
}

// OnCDOTAUserMsg_SwapVerify registers a callback for EDotaUserMessages_DOTA_UM_SwapVerify
func (c *Callbacks) OnCDOTAUserMsg_SwapVerify(fn func(*dota.CDOTAUserMsg_SwapVerify) error) {
	c.onCDOTAUserMsg_SwapVerify = append(c.onCDOTAUserMsg_SwapVerify, fn)
}

// OnCDOTAUserMsg_WorldLine registers a callback for EDotaUserMessages_DOTA_UM_WorldLine
func (c *Callbacks) OnCDOTAUserMsg_WorldLine(fn func(*dota.CDOTAUserMsg_WorldLine) error) {
	c.onCDOTAUserMsg_WorldLine = append(c.onCDOTAUserMsg_WorldLine, fn)
}

// OnCMsgGCToClientTournamentItemDrop registers a callback for EDotaUserMessages_DOTA_UM_TournamentDrop
func (c *Callbacks) OnCMsgGCToClientTournamentItemDrop(fn func(*dota.CMsgGCToClientTournamentItemDrop) error) {
	c.onCMsgGCToClientTournamentItemDrop = append(c.onCMsgGCToClientTournamentItemDrop, fn)
}

// OnCDOTAUserMsg_ItemAlert registers a callback for EDotaUserMessages_DOTA_UM_ItemAlert
func (c *Callbacks) OnCDOTAUserMsg_ItemAlert(fn func(*dota.CDOTAUserMsg_ItemAlert) error) {
	c.onCDOTAUserMsg_ItemAlert = append(c.onCDOTAUserMsg_ItemAlert, fn)
}

// OnCDOTAUserMsg_HalloweenDrops registers a callback for EDotaUserMessages_DOTA_UM_HalloweenDrops
func (c *Callbacks) OnCDOTAUserMsg_HalloweenDrops(fn func(*dota.CDOTAUserMsg_HalloweenDrops) error) {
	c.onCDOTAUserMsg_HalloweenDrops = append(c.onCDOTAUserMsg_HalloweenDrops, fn)
}

// OnCDOTAUserMsg_ChatWheel registers a callback for EDotaUserMessages_DOTA_UM_ChatWheel
func (c *Callbacks) OnCDOTAUserMsg_ChatWheel(fn func(*dota.CDOTAUserMsg_ChatWheel) error) {
	c.onCDOTAUserMsg_ChatWheel = append(c.onCDOTAUserMsg_ChatWheel, fn)
}

// OnCDOTAUserMsg_ReceivedXmasGift registers a callback for EDotaUserMessages_DOTA_UM_ReceivedXmasGift
func (c *Callbacks) OnCDOTAUserMsg_ReceivedXmasGift(fn func(*dota.CDOTAUserMsg_ReceivedXmasGift) error) {
	c.onCDOTAUserMsg_ReceivedXmasGift = append(c.onCDOTAUserMsg_ReceivedXmasGift, fn)
}

// OnCDOTAUserMsg_UpdateSharedContent registers a callback for EDotaUserMessages_DOTA_UM_UpdateSharedContent
func (c *Callbacks) OnCDOTAUserMsg_UpdateSharedContent(fn func(*dota.CDOTAUserMsg_UpdateSharedContent) error) {
	c.onCDOTAUserMsg_UpdateSharedContent = append(c.onCDOTAUserMsg_UpdateSharedContent, fn)
}

// OnCDOTAUserMsg_TutorialRequestExp registers a callback for EDotaUserMessages_DOTA_UM_TutorialRequestExp
func (c *Callbacks) OnCDOTAUserMsg_TutorialRequestExp(fn func(*dota.CDOTAUserMsg_TutorialRequestExp) error) {
	c.onCDOTAUserMsg_TutorialRequestExp = append(c.onCDOTAUserMsg_TutorialRequestExp, fn)
}

// OnCDOTAUserMsg_TutorialPingMinimap registers a callback for EDotaUserMessages_DOTA_UM_TutorialPingMinimap
func (c *Callbacks) OnCDOTAUserMsg_TutorialPingMinimap(fn func(*dota.CDOTAUserMsg_TutorialPingMinimap) error) {
	c.onCDOTAUserMsg_TutorialPingMinimap = append(c.onCDOTAUserMsg_TutorialPingMinimap, fn)
}

// OnCDOTAUserMsg_GamerulesStateChanged registers a callback for EDotaUserMessages_DOTA_UM_GamerulesStateChanged
func (c *Callbacks) OnCDOTAUserMsg_GamerulesStateChanged(fn func(*dota.CDOTAUserMsg_GamerulesStateChanged) error) {
	c.onCDOTAUserMsg_GamerulesStateChanged = append(c.onCDOTAUserMsg_GamerulesStateChanged, fn)
}

// OnCDOTAUserMsg_ShowSurvey registers a callback for EDotaUserMessages_DOTA_UM_ShowSurvey
func (c *Callbacks) OnCDOTAUserMsg_ShowSurvey(fn func(*dota.CDOTAUserMsg_ShowSurvey) error) {
	c.onCDOTAUserMsg_ShowSurvey = append(c.onCDOTAUserMsg_ShowSurvey, fn)
}

// OnCDOTAUserMsg_TutorialFade registers a callback for EDotaUserMessages_DOTA_UM_TutorialFade
func (c *Callbacks) OnCDOTAUserMsg_TutorialFade(fn func(*dota.CDOTAUserMsg_TutorialFade) error) {
	c.onCDOTAUserMsg_TutorialFade = append(c.onCDOTAUserMsg_TutorialFade, fn)
}

// OnCDOTAUserMsg_AddQuestLogEntry registers a callback for EDotaUserMessages_DOTA_UM_AddQuestLogEntry
func (c *Callbacks) OnCDOTAUserMsg_AddQuestLogEntry(fn func(*dota.CDOTAUserMsg_AddQuestLogEntry) error) {
	c.onCDOTAUserMsg_AddQuestLogEntry = append(c.onCDOTAUserMsg_AddQuestLogEntry, fn)
}

// OnCDOTAUserMsg_SendStatPopup registers a callback for EDotaUserMessages_DOTA_UM_SendStatPopup
func (c *Callbacks) OnCDOTAUserMsg_SendStatPopup(fn func(*dota.CDOTAUserMsg_SendStatPopup) error) {
	c.onCDOTAUserMsg_SendStatPopup = append(c.onCDOTAUserMsg_SendStatPopup, fn)
}

// OnCDOTAUserMsg_TutorialFinish registers a callback for EDotaUserMessages_DOTA_UM_TutorialFinish
func (c *Callbacks) OnCDOTAUserMsg_TutorialFinish(fn func(*dota.CDOTAUserMsg_TutorialFinish) error) {
	c.onCDOTAUserMsg_TutorialFinish = append(c.onCDOTAUserMsg_TutorialFinish, fn)
}

// OnCDOTAUserMsg_SendRoshanPopup registers a callback for EDotaUserMessages_DOTA_UM_SendRoshanPopup
func (c *Callbacks) OnCDOTAUserMsg_SendRoshanPopup(fn func(*dota.CDOTAUserMsg_SendRoshanPopup) error) {
	c.onCDOTAUserMsg_SendRoshanPopup = append(c.onCDOTAUserMsg_SendRoshanPopup, fn)
}

// OnCDOTAUserMsg_SendGenericToolTip registers a callback for EDotaUserMessages_DOTA_UM_SendGenericToolTip
func (c *Callbacks) OnCDOTAUserMsg_SendGenericToolTip(fn func(*dota.CDOTAUserMsg_SendGenericToolTip) error) {
	c.onCDOTAUserMsg_SendGenericToolTip = append(c.onCDOTAUserMsg_SendGenericToolTip, fn)
}

// OnCDOTAUserMsg_SendFinalGold registers a callback for EDotaUserMessages_DOTA_UM_SendFinalGold
func (c *Callbacks) OnCDOTAUserMsg_SendFinalGold(fn func(*dota.CDOTAUserMsg_SendFinalGold) error) {
	c.onCDOTAUserMsg_SendFinalGold = append(c.onCDOTAUserMsg_SendFinalGold, fn)
}

// OnCDOTAUserMsg_CustomMsg registers a callback for EDotaUserMessages_DOTA_UM_CustomMsg
func (c *Callbacks) OnCDOTAUserMsg_CustomMsg(fn func(*dota.CDOTAUserMsg_CustomMsg) error) {
	c.onCDOTAUserMsg_CustomMsg = append(c.onCDOTAUserMsg_CustomMsg, fn)
}

// OnCDOTAUserMsg_CoachHUDPing registers a callback for EDotaUserMessages_DOTA_UM_CoachHUDPing
func (c *Callbacks) OnCDOTAUserMsg_CoachHUDPing(fn func(*dota.CDOTAUserMsg_CoachHUDPing) error) {
	c.onCDOTAUserMsg_CoachHUDPing = append(c.onCDOTAUserMsg_CoachHUDPing, fn)
}

// OnCDOTAUserMsg_ClientLoadGridNav registers a callback for EDotaUserMessages_DOTA_UM_ClientLoadGridNav
func (c *Callbacks) OnCDOTAUserMsg_ClientLoadGridNav(fn func(*dota.CDOTAUserMsg_ClientLoadGridNav) error) {
	c.onCDOTAUserMsg_ClientLoadGridNav = append(c.onCDOTAUserMsg_ClientLoadGridNav, fn)
}

// OnCDOTAUserMsg_TE_Projectile registers a callback for EDotaUserMessages_DOTA_UM_TE_Projectile
func (c *Callbacks) OnCDOTAUserMsg_TE_Projectile(fn func(*dota.CDOTAUserMsg_TE_Projectile) error) {
	c.onCDOTAUserMsg_TE_Projectile = append(c.onCDOTAUserMsg_TE_Projectile, fn)
}

// OnCDOTAUserMsg_TE_ProjectileLoc registers a callback for EDotaUserMessages_DOTA_UM_TE_ProjectileLoc
func (c *Callbacks) OnCDOTAUserMsg_TE_ProjectileLoc(fn func(*dota.CDOTAUserMsg_TE_ProjectileLoc) error) {
	c.onCDOTAUserMsg_TE_ProjectileLoc = append(c.onCDOTAUserMsg_TE_ProjectileLoc, fn)
}

// OnCDOTAUserMsg_TE_DotaBloodImpact registers a callback for EDotaUserMessages_DOTA_UM_TE_DotaBloodImpact
func (c *Callbacks) OnCDOTAUserMsg_TE_DotaBloodImpact(fn func(*dota.CDOTAUserMsg_TE_DotaBloodImpact) error) {
	c.onCDOTAUserMsg_TE_DotaBloodImpact = append(c.onCDOTAUserMsg_TE_DotaBloodImpact, fn)
}

// OnCDOTAUserMsg_TE_UnitAnimation registers a callback for EDotaUserMessages_DOTA_UM_TE_UnitAnimation
func (c *Callbacks) OnCDOTAUserMsg_TE_UnitAnimation(fn func(*dota.CDOTAUserMsg_TE_UnitAnimation) error) {
	c.onCDOTAUserMsg_TE_UnitAnimation = append(c.onCDOTAUserMsg_TE_UnitAnimation, fn)
}

// OnCDOTAUserMsg_TE_UnitAnimationEnd registers a callback for EDotaUserMessages_DOTA_UM_TE_UnitAnimationEnd
func (c *Callbacks) OnCDOTAUserMsg_TE_UnitAnimationEnd(fn func(*dota.CDOTAUserMsg_TE_UnitAnimationEnd) error) {
	c.onCDOTAUserMsg_TE_UnitAnimationEnd = append(c.onCDOTAUserMsg_TE_UnitAnimationEnd, fn)
}

// OnCDOTAUserMsg_AbilityPing registers a callback for EDotaUserMessages_DOTA_UM_AbilityPing
func (c *Callbacks) OnCDOTAUserMsg_AbilityPing(fn func(*dota.CDOTAUserMsg_AbilityPing) error) {
	c.onCDOTAUserMsg_AbilityPing = append(c.onCDOTAUserMsg_AbilityPing, fn)
}

// OnCDOTAUserMsg_ShowGenericPopup registers a callback for EDotaUserMessages_DOTA_UM_ShowGenericPopup
func (c *Callbacks) OnCDOTAUserMsg_ShowGenericPopup(fn func(*dota.CDOTAUserMsg_ShowGenericPopup) error) {
	c.onCDOTAUserMsg_ShowGenericPopup = append(c.onCDOTAUserMsg_ShowGenericPopup, fn)
}

// OnCDOTAUserMsg_VoteStart registers a callback for EDotaUserMessages_DOTA_UM_VoteStart
func (c *Callbacks) OnCDOTAUserMsg_VoteStart(fn func(*dota.CDOTAUserMsg_VoteStart) error) {
	c.onCDOTAUserMsg_VoteStart = append(c.onCDOTAUserMsg_VoteStart, fn)
}

// OnCDOTAUserMsg_VoteUpdate registers a callback for EDotaUserMessages_DOTA_UM_VoteUpdate
func (c *Callbacks) OnCDOTAUserMsg_VoteUpdate(fn func(*dota.CDOTAUserMsg_VoteUpdate) error) {
	c.onCDOTAUserMsg_VoteUpdate = append(c.onCDOTAUserMsg_VoteUpdate, fn)
}

// OnCDOTAUserMsg_VoteEnd registers a callback for EDotaUserMessages_DOTA_UM_VoteEnd
func (c *Callbacks) OnCDOTAUserMsg_VoteEnd(fn func(*dota.CDOTAUserMsg_VoteEnd) error) {
	c.onCDOTAUserMsg_VoteEnd = append(c.onCDOTAUserMsg_VoteEnd, fn)
}

// OnCDOTAUserMsg_BoosterState registers a callback for EDotaUserMessages_DOTA_UM_BoosterState
func (c *Callbacks) OnCDOTAUserMsg_BoosterState(fn func(*dota.CDOTAUserMsg_BoosterState) error) {
	c.onCDOTAUserMsg_BoosterState = append(c.onCDOTAUserMsg_BoosterState, fn)
}

// OnCDOTAUserMsg_WillPurchaseAlert registers a callback for EDotaUserMessages_DOTA_UM_WillPurchaseAlert
func (c *Callbacks) OnCDOTAUserMsg_WillPurchaseAlert(fn func(*dota.CDOTAUserMsg_WillPurchaseAlert) error) {
	c.onCDOTAUserMsg_WillPurchaseAlert = append(c.onCDOTAUserMsg_WillPurchaseAlert, fn)
}

// OnCDOTAUserMsg_TutorialMinimapPosition registers a callback for EDotaUserMessages_DOTA_UM_TutorialMinimapPosition
func (c *Callbacks) OnCDOTAUserMsg_TutorialMinimapPosition(fn func(*dota.CDOTAUserMsg_TutorialMinimapPosition) error) {
	c.onCDOTAUserMsg_TutorialMinimapPosition = append(c.onCDOTAUserMsg_TutorialMinimapPosition, fn)
}

// OnCDOTAUserMsg_AbilitySteal registers a callback for EDotaUserMessages_DOTA_UM_AbilitySteal
func (c *Callbacks) OnCDOTAUserMsg_AbilitySteal(fn func(*dota.CDOTAUserMsg_AbilitySteal) error) {
	c.onCDOTAUserMsg_AbilitySteal = append(c.onCDOTAUserMsg_AbilitySteal, fn)
}

// OnCDOTAUserMsg_CourierKilledAlert registers a callback for EDotaUserMessages_DOTA_UM_CourierKilledAlert
func (c *Callbacks) OnCDOTAUserMsg_CourierKilledAlert(fn func(*dota.CDOTAUserMsg_CourierKilledAlert) error) {
	c.onCDOTAUserMsg_CourierKilledAlert = append(c.onCDOTAUserMsg_CourierKilledAlert, fn)
}

// OnCDOTAUserMsg_EnemyItemAlert registers a callback for EDotaUserMessages_DOTA_UM_EnemyItemAlert
func (c *Callbacks) OnCDOTAUserMsg_EnemyItemAlert(fn func(*dota.CDOTAUserMsg_EnemyItemAlert) error) {
	c.onCDOTAUserMsg_EnemyItemAlert = append(c.onCDOTAUserMsg_EnemyItemAlert, fn)
}

// OnCDOTAUserMsg_StatsMatchDetails registers a callback for EDotaUserMessages_DOTA_UM_StatsMatchDetails
func (c *Callbacks) OnCDOTAUserMsg_StatsMatchDetails(fn func(*dota.CDOTAUserMsg_StatsMatchDetails) error) {
	c.onCDOTAUserMsg_StatsMatchDetails = append(c.onCDOTAUserMsg_StatsMatchDetails, fn)
}

// OnCDOTAUserMsg_MiniTaunt registers a callback for EDotaUserMessages_DOTA_UM_MiniTaunt
func (c *Callbacks) OnCDOTAUserMsg_MiniTaunt(fn func(*dota.CDOTAUserMsg_MiniTaunt) error) {
	c.onCDOTAUserMsg_MiniTaunt = append(c.onCDOTAUserMsg_MiniTaunt, fn)
}

// OnCDOTAUserMsg_BuyBackStateAlert registers a callback for EDotaUserMessages_DOTA_UM_BuyBackStateAlert
func (c *Callbacks) OnCDOTAUserMsg_BuyBackStateAlert(fn func(*dota.CDOTAUserMsg_BuyBackStateAlert) error) {
	c.onCDOTAUserMsg_BuyBackStateAlert = append(c.onCDOTAUserMsg_BuyBackStateAlert, fn)
}

// OnCDOTAUserMsg_SpeechBubble registers a callback for EDotaUserMessages_DOTA_UM_SpeechBubble
func (c *Callbacks) OnCDOTAUserMsg_SpeechBubble(fn func(*dota.CDOTAUserMsg_SpeechBubble) error) {
	c.onCDOTAUserMsg_SpeechBubble = append(c.onCDOTAUserMsg_SpeechBubble, fn)
}

// OnCDOTAUserMsg_CustomHeaderMessage registers a callback for EDotaUserMessages_DOTA_UM_CustomHeaderMessage
func (c *Callbacks) OnCDOTAUserMsg_CustomHeaderMessage(fn func(*dota.CDOTAUserMsg_CustomHeaderMessage) error) {
	c.onCDOTAUserMsg_CustomHeaderMessage = append(c.onCDOTAUserMsg_CustomHeaderMessage, fn)
}

// OnCDOTAUserMsg_QuickBuyAlert registers a callback for EDotaUserMessages_DOTA_UM_QuickBuyAlert
func (c *Callbacks) OnCDOTAUserMsg_QuickBuyAlert(fn func(*dota.CDOTAUserMsg_QuickBuyAlert) error) {
	c.onCDOTAUserMsg_QuickBuyAlert = append(c.onCDOTAUserMsg_QuickBuyAlert, fn)
}

// OnCDOTAUserMsg_StatsHeroMinuteDetails registers a callback for EDotaUserMessages_DOTA_UM_StatsHeroDetails
func (c *Callbacks) OnCDOTAUserMsg_StatsHeroMinuteDetails(fn func(*dota.CDOTAUserMsg_StatsHeroMinuteDetails) error) {
	c.onCDOTAUserMsg_StatsHeroMinuteDetails = append(c.onCDOTAUserMsg_StatsHeroMinuteDetails, fn)
}

// OnCDOTAUserMsg_ModifierAlert registers a callback for EDotaUserMessages_DOTA_UM_ModifierAlert
func (c *Callbacks) OnCDOTAUserMsg_ModifierAlert(fn func(*dota.CDOTAUserMsg_ModifierAlert) error) {
	c.onCDOTAUserMsg_ModifierAlert = append(c.onCDOTAUserMsg_ModifierAlert, fn)
}

// OnCDOTAUserMsg_HPManaAlert registers a callback for EDotaUserMessages_DOTA_UM_HPManaAlert
func (c *Callbacks) OnCDOTAUserMsg_HPManaAlert(fn func(*dota.CDOTAUserMsg_HPManaAlert) error) {
	c.onCDOTAUserMsg_HPManaAlert = append(c.onCDOTAUserMsg_HPManaAlert, fn)
}

// OnCDOTAUserMsg_GlyphAlert registers a callback for EDotaUserMessages_DOTA_UM_GlyphAlert
func (c *Callbacks) OnCDOTAUserMsg_GlyphAlert(fn func(*dota.CDOTAUserMsg_GlyphAlert) error) {
	c.onCDOTAUserMsg_GlyphAlert = append(c.onCDOTAUserMsg_GlyphAlert, fn)
}

// OnCDOTAUserMsg_BeastChat registers a callback for EDotaUserMessages_DOTA_UM_BeastChat
func (c *Callbacks) OnCDOTAUserMsg_BeastChat(fn func(*dota.CDOTAUserMsg_BeastChat) error) {
	c.onCDOTAUserMsg_BeastChat = append(c.onCDOTAUserMsg_BeastChat, fn)
}

// OnCDOTAUserMsg_SpectatorPlayerUnitOrders registers a callback for EDotaUserMessages_DOTA_UM_SpectatorPlayerUnitOrders
func (c *Callbacks) OnCDOTAUserMsg_SpectatorPlayerUnitOrders(fn func(*dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error) {
	c.onCDOTAUserMsg_SpectatorPlayerUnitOrders = append(c.onCDOTAUserMsg_SpectatorPlayerUnitOrders, fn)
}

// OnCDOTAUserMsg_CustomHudElement_Create registers a callback for EDotaUserMessages_DOTA_UM_CustomHudElement_Create
func (c *Callbacks) OnCDOTAUserMsg_CustomHudElement_Create(fn func(*dota.CDOTAUserMsg_CustomHudElement_Create) error) {
	c.onCDOTAUserMsg_CustomHudElement_Create = append(c.onCDOTAUserMsg_CustomHudElement_Create, fn)
}

// OnCDOTAUserMsg_CustomHudElement_Modify registers a callback for EDotaUserMessages_DOTA_UM_CustomHudElement_Modify
func (c *Callbacks) OnCDOTAUserMsg_CustomHudElement_Modify(fn func(*dota.CDOTAUserMsg_CustomHudElement_Modify) error) {
	c.onCDOTAUserMsg_CustomHudElement_Modify = append(c.onCDOTAUserMsg_CustomHudElement_Modify, fn)
}

// OnCDOTAUserMsg_CustomHudElement_Destroy registers a callback for EDotaUserMessages_DOTA_UM_CustomHudElement_Destroy
func (c *Callbacks) OnCDOTAUserMsg_CustomHudElement_Destroy(fn func(*dota.CDOTAUserMsg_CustomHudElement_Destroy) error) {
	c.onCDOTAUserMsg_CustomHudElement_Destroy = append(c.onCDOTAUserMsg_CustomHudElement_Destroy, fn)
}

// OnCDOTAUserMsg_CompendiumState registers a callback for EDotaUserMessages_DOTA_UM_CompendiumState
func (c *Callbacks) OnCDOTAUserMsg_CompendiumState(fn func(*dota.CDOTAUserMsg_CompendiumState) error) {
	c.onCDOTAUserMsg_CompendiumState = append(c.onCDOTAUserMsg_CompendiumState, fn)
}

// OnCDOTAUserMsg_ProjectionAbility registers a callback for EDotaUserMessages_DOTA_UM_ProjectionAbility
func (c *Callbacks) OnCDOTAUserMsg_ProjectionAbility(fn func(*dota.CDOTAUserMsg_ProjectionAbility) error) {
	c.onCDOTAUserMsg_ProjectionAbility = append(c.onCDOTAUserMsg_ProjectionAbility, fn)
}

// OnCDOTAUserMsg_ProjectionEvent registers a callback for EDotaUserMessages_DOTA_UM_ProjectionEvent
func (c *Callbacks) OnCDOTAUserMsg_ProjectionEvent(fn func(*dota.CDOTAUserMsg_ProjectionEvent) error) {
	c.onCDOTAUserMsg_ProjectionEvent = append(c.onCDOTAUserMsg_ProjectionEvent, fn)
}

// OnCMsgDOTACombatLogEntry registers a callback for EDotaUserMessages_DOTA_UM_CombatLogDataHLTV
func (c *Callbacks) OnCMsgDOTACombatLogEntry(fn func(*dota.CMsgDOTACombatLogEntry) error) {
	c.onCMsgDOTACombatLogEntry = append(c.onCMsgDOTACombatLogEntry, fn)
}

// OnCDOTAUserMsg_XPAlert registers a callback for EDotaUserMessages_DOTA_UM_XPAlert
func (c *Callbacks) OnCDOTAUserMsg_XPAlert(fn func(*dota.CDOTAUserMsg_XPAlert) error) {
	c.onCDOTAUserMsg_XPAlert = append(c.onCDOTAUserMsg_XPAlert, fn)
}

// OnCDOTAUserMsg_UpdateQuestProgress registers a callback for EDotaUserMessages_DOTA_UM_UpdateQuestProgress
func (c *Callbacks) OnCDOTAUserMsg_UpdateQuestProgress(fn func(*dota.CDOTAUserMsg_UpdateQuestProgress) error) {
	c.onCDOTAUserMsg_UpdateQuestProgress = append(c.onCDOTAUserMsg_UpdateQuestProgress, fn)
}

// OnCDOTAMatchMetadataFile registers a callback for EDotaUserMessages_DOTA_UM_MatchMetadata
func (c *Callbacks) OnCDOTAMatchMetadataFile(fn func(*dota.CDOTAMatchMetadataFile) error) {
	c.onCDOTAMatchMetadataFile = append(c.onCDOTAMatchMetadataFile, fn)
}

// OnCDOTAUserMsg_QuestStatus registers a callback for EDotaUserMessages_DOTA_UM_QuestStatus
func (c *Callbacks) OnCDOTAUserMsg_QuestStatus(fn func(*dota.CDOTAUserMsg_QuestStatus) error) {
	c.onCDOTAUserMsg_QuestStatus = append(c.onCDOTAUserMsg_QuestStatus, fn)
}

// OnCDOTAUserMsg_SuggestHeroPick registers a callback for EDotaUserMessages_DOTA_UM_SuggestHeroPick
func (c *Callbacks) OnCDOTAUserMsg_SuggestHeroPick(fn func(*dota.CDOTAUserMsg_SuggestHeroPick) error) {
	c.onCDOTAUserMsg_SuggestHeroPick = append(c.onCDOTAUserMsg_SuggestHeroPick, fn)
}

// OnCDOTAUserMsg_SuggestHeroRole registers a callback for EDotaUserMessages_DOTA_UM_SuggestHeroRole
func (c *Callbacks) OnCDOTAUserMsg_SuggestHeroRole(fn func(*dota.CDOTAUserMsg_SuggestHeroRole) error) {
	c.onCDOTAUserMsg_SuggestHeroRole = append(c.onCDOTAUserMsg_SuggestHeroRole, fn)
}

// OnCDOTAUserMsg_KillcamDamageTaken registers a callback for EDotaUserMessages_DOTA_UM_KillcamDamageTaken
func (c *Callbacks) OnCDOTAUserMsg_KillcamDamageTaken(fn func(*dota.CDOTAUserMsg_KillcamDamageTaken) error) {
	c.onCDOTAUserMsg_KillcamDamageTaken = append(c.onCDOTAUserMsg_KillcamDamageTaken, fn)
}

// OnCDOTAUserMsg_SelectPenaltyGold registers a callback for EDotaUserMessages_DOTA_UM_SelectPenaltyGold
func (c *Callbacks) OnCDOTAUserMsg_SelectPenaltyGold(fn func(*dota.CDOTAUserMsg_SelectPenaltyGold) error) {
	c.onCDOTAUserMsg_SelectPenaltyGold = append(c.onCDOTAUserMsg_SelectPenaltyGold, fn)
}

// OnCDOTAUserMsg_RollDiceResult registers a callback for EDotaUserMessages_DOTA_UM_RollDiceResult
func (c *Callbacks) OnCDOTAUserMsg_RollDiceResult(fn func(*dota.CDOTAUserMsg_RollDiceResult) error) {
	c.onCDOTAUserMsg_RollDiceResult = append(c.onCDOTAUserMsg_RollDiceResult, fn)
}

// OnCDOTAUserMsg_FlipCoinResult registers a callback for EDotaUserMessages_DOTA_UM_FlipCoinResult
func (c *Callbacks) OnCDOTAUserMsg_FlipCoinResult(fn func(*dota.CDOTAUserMsg_FlipCoinResult) error) {
	c.onCDOTAUserMsg_FlipCoinResult = append(c.onCDOTAUserMsg_FlipCoinResult, fn)
}

// OnCDOTAUserMsg_SendRoshanSpectatorPhase registers a callback for EDotaUserMessages_DOTA_UM_SendRoshanSpectatorPhase
func (c *Callbacks) OnCDOTAUserMsg_SendRoshanSpectatorPhase(fn func(*dota.CDOTAUserMsg_SendRoshanSpectatorPhase) error) {
	c.onCDOTAUserMsg_SendRoshanSpectatorPhase = append(c.onCDOTAUserMsg_SendRoshanSpectatorPhase, fn)
}

// OnCDOTAUserMsg_ChatWheelCooldown registers a callback for EDotaUserMessages_DOTA_UM_ChatWheelCooldown
func (c *Callbacks) OnCDOTAUserMsg_ChatWheelCooldown(fn func(*dota.CDOTAUserMsg_ChatWheelCooldown) error) {
	c.onCDOTAUserMsg_ChatWheelCooldown = append(c.onCDOTAUserMsg_ChatWheelCooldown, fn)
}

// OnCDOTAUserMsg_DismissAllStatPopups registers a callback for EDotaUserMessages_DOTA_UM_DismissAllStatPopups
func (c *Callbacks) OnCDOTAUserMsg_DismissAllStatPopups(fn func(*dota.CDOTAUserMsg_DismissAllStatPopups) error) {
	c.onCDOTAUserMsg_DismissAllStatPopups = append(c.onCDOTAUserMsg_DismissAllStatPopups, fn)
}

// OnCDOTAUserMsg_TE_DestroyProjectile registers a callback for EDotaUserMessages_DOTA_UM_TE_DestroyProjectile
func (c *Callbacks) OnCDOTAUserMsg_TE_DestroyProjectile(fn func(*dota.CDOTAUserMsg_TE_DestroyProjectile) error) {
	c.onCDOTAUserMsg_TE_DestroyProjectile = append(c.onCDOTAUserMsg_TE_DestroyProjectile, fn)
}

// OnCDOTAUserMsg_HeroRelicProgress registers a callback for EDotaUserMessages_DOTA_UM_HeroRelicProgress
func (c *Callbacks) OnCDOTAUserMsg_HeroRelicProgress(fn func(*dota.CDOTAUserMsg_HeroRelicProgress) error) {
	c.onCDOTAUserMsg_HeroRelicProgress = append(c.onCDOTAUserMsg_HeroRelicProgress, fn)
}

// OnCDOTAUserMsg_AbilityDraftRequestAbility registers a callback for EDotaUserMessages_DOTA_UM_AbilityDraftRequestAbility
func (c *Callbacks) OnCDOTAUserMsg_AbilityDraftRequestAbility(fn func(*dota.CDOTAUserMsg_AbilityDraftRequestAbility) error) {
	c.onCDOTAUserMsg_AbilityDraftRequestAbility = append(c.onCDOTAUserMsg_AbilityDraftRequestAbility, fn)
}

// OnCDOTAUserMsg_ItemSold registers a callback for EDotaUserMessages_DOTA_UM_ItemSold
func (c *Callbacks) OnCDOTAUserMsg_ItemSold(fn func(*dota.CDOTAUserMsg_ItemSold) error) {
	c.onCDOTAUserMsg_ItemSold = append(c.onCDOTAUserMsg_ItemSold, fn)
}

// OnCDOTAUserMsg_DamageReport registers a callback for EDotaUserMessages_DOTA_UM_DamageReport
func (c *Callbacks) OnCDOTAUserMsg_DamageReport(fn func(*dota.CDOTAUserMsg_DamageReport) error) {
	c.onCDOTAUserMsg_DamageReport = append(c.onCDOTAUserMsg_DamageReport, fn)
}

// OnCDOTAUserMsg_SalutePlayer registers a callback for EDotaUserMessages_DOTA_UM_SalutePlayer
func (c *Callbacks) OnCDOTAUserMsg_SalutePlayer(fn func(*dota.CDOTAUserMsg_SalutePlayer) error) {
	c.onCDOTAUserMsg_SalutePlayer = append(c.onCDOTAUserMsg_SalutePlayer, fn)
}

// OnCDOTAUserMsg_TipAlert registers a callback for EDotaUserMessages_DOTA_UM_TipAlert
func (c *Callbacks) OnCDOTAUserMsg_TipAlert(fn func(*dota.CDOTAUserMsg_TipAlert) error) {
	c.onCDOTAUserMsg_TipAlert = append(c.onCDOTAUserMsg_TipAlert, fn)
}

// OnCDOTAUserMsg_ReplaceQueryUnit registers a callback for EDotaUserMessages_DOTA_UM_ReplaceQueryUnit
func (c *Callbacks) OnCDOTAUserMsg_ReplaceQueryUnit(fn func(*dota.CDOTAUserMsg_ReplaceQueryUnit) error) {
	c.onCDOTAUserMsg_ReplaceQueryUnit = append(c.onCDOTAUserMsg_ReplaceQueryUnit, fn)
}

// OnCDOTAUserMsg_EmptyTeleportAlert registers a callback for EDotaUserMessages_DOTA_UM_EmptyTeleportAlert
func (c *Callbacks) OnCDOTAUserMsg_EmptyTeleportAlert(fn func(*dota.CDOTAUserMsg_EmptyTeleportAlert) error) {
	c.onCDOTAUserMsg_EmptyTeleportAlert = append(c.onCDOTAUserMsg_EmptyTeleportAlert, fn)
}

// OnCDOTAUserMsg_MarsArenaOfBloodAttack registers a callback for EDotaUserMessages_DOTA_UM_MarsArenaOfBloodAttack
func (c *Callbacks) OnCDOTAUserMsg_MarsArenaOfBloodAttack(fn func(*dota.CDOTAUserMsg_MarsArenaOfBloodAttack) error) {
	c.onCDOTAUserMsg_MarsArenaOfBloodAttack = append(c.onCDOTAUserMsg_MarsArenaOfBloodAttack, fn)
}

// OnCDOTAUserMsg_ESArcanaCombo registers a callback for EDotaUserMessages_DOTA_UM_ESArcanaCombo
func (c *Callbacks) OnCDOTAUserMsg_ESArcanaCombo(fn func(*dota.CDOTAUserMsg_ESArcanaCombo) error) {
	c.onCDOTAUserMsg_ESArcanaCombo = append(c.onCDOTAUserMsg_ESArcanaCombo, fn)
}

// OnCDOTAUserMsg_ESArcanaComboSummary registers a callback for EDotaUserMessages_DOTA_UM_ESArcanaComboSummary
func (c *Callbacks) OnCDOTAUserMsg_ESArcanaComboSummary(fn func(*dota.CDOTAUserMsg_ESArcanaComboSummary) error) {
	c.onCDOTAUserMsg_ESArcanaComboSummary = append(c.onCDOTAUserMsg_ESArcanaComboSummary, fn)
}

// OnCDOTAUserMsg_HighFiveLeftHanging registers a callback for EDotaUserMessages_DOTA_UM_HighFiveLeftHanging
func (c *Callbacks) OnCDOTAUserMsg_HighFiveLeftHanging(fn func(*dota.CDOTAUserMsg_HighFiveLeftHanging) error) {
	c.onCDOTAUserMsg_HighFiveLeftHanging = append(c.onCDOTAUserMsg_HighFiveLeftHanging, fn)
}

// OnCDOTAUserMsg_HighFiveCompleted registers a callback for EDotaUserMessages_DOTA_UM_HighFiveCompleted
func (c *Callbacks) OnCDOTAUserMsg_HighFiveCompleted(fn func(*dota.CDOTAUserMsg_HighFiveCompleted) error) {
	c.onCDOTAUserMsg_HighFiveCompleted = append(c.onCDOTAUserMsg_HighFiveCompleted, fn)
}

// OnCDOTAUserMsg_ShovelUnearth registers a callback for EDotaUserMessages_DOTA_UM_ShovelUnearth
func (c *Callbacks) OnCDOTAUserMsg_ShovelUnearth(fn func(*dota.CDOTAUserMsg_ShovelUnearth) error) {
	c.onCDOTAUserMsg_ShovelUnearth = append(c.onCDOTAUserMsg_ShovelUnearth, fn)
}

// OnCDOTAUserMsg_RadarAlert registers a callback for EDotaUserMessages_DOTA_UM_RadarAlert
func (c *Callbacks) OnCDOTAUserMsg_RadarAlert(fn func(*dota.CDOTAUserMsg_RadarAlert) error) {
	c.onCDOTAUserMsg_RadarAlert = append(c.onCDOTAUserMsg_RadarAlert, fn)
}

// OnCDOTAUserMsg_AllStarEvent registers a callback for EDotaUserMessages_DOTA_UM_AllStarEvent
func (c *Callbacks) OnCDOTAUserMsg_AllStarEvent(fn func(*dota.CDOTAUserMsg_AllStarEvent) error) {
	c.onCDOTAUserMsg_AllStarEvent = append(c.onCDOTAUserMsg_AllStarEvent, fn)
}

// OnCDOTAUserMsg_TalentTreeAlert registers a callback for EDotaUserMessages_DOTA_UM_TalentTreeAlert
func (c *Callbacks) OnCDOTAUserMsg_TalentTreeAlert(fn func(*dota.CDOTAUserMsg_TalentTreeAlert) error) {
	c.onCDOTAUserMsg_TalentTreeAlert = append(c.onCDOTAUserMsg_TalentTreeAlert, fn)
}

// OnCDOTAUserMsg_QueuedOrderRemoved registers a callback for EDotaUserMessages_DOTA_UM_QueuedOrderRemoved
func (c *Callbacks) OnCDOTAUserMsg_QueuedOrderRemoved(fn func(*dota.CDOTAUserMsg_QueuedOrderRemoved) error) {
	c.onCDOTAUserMsg_QueuedOrderRemoved = append(c.onCDOTAUserMsg_QueuedOrderRemoved, fn)
}

// OnCDOTAUserMsg_DebugChallenge registers a callback for EDotaUserMessages_DOTA_UM_DebugChallenge
func (c *Callbacks) OnCDOTAUserMsg_DebugChallenge(fn func(*dota.CDOTAUserMsg_DebugChallenge) error) {
	c.onCDOTAUserMsg_DebugChallenge = append(c.onCDOTAUserMsg_DebugChallenge, fn)
}

// OnCDOTAUserMsg_OMArcanaCombo registers a callback for EDotaUserMessages_DOTA_UM_OMArcanaCombo
func (c *Callbacks) OnCDOTAUserMsg_OMArcanaCombo(fn func(*dota.CDOTAUserMsg_OMArcanaCombo) error) {
	c.onCDOTAUserMsg_OMArcanaCombo = append(c.onCDOTAUserMsg_OMArcanaCombo, fn)
}

// OnCDOTAUserMsg_FoundNeutralItem registers a callback for EDotaUserMessages_DOTA_UM_FoundNeutralItem
func (c *Callbacks) OnCDOTAUserMsg_FoundNeutralItem(fn func(*dota.CDOTAUserMsg_FoundNeutralItem) error) {
	c.onCDOTAUserMsg_FoundNeutralItem = append(c.onCDOTAUserMsg_FoundNeutralItem, fn)
}

// OnCDOTAUserMsg_OutpostCaptured registers a callback for EDotaUserMessages_DOTA_UM_OutpostCaptured
func (c *Callbacks) OnCDOTAUserMsg_OutpostCaptured(fn func(*dota.CDOTAUserMsg_OutpostCaptured) error) {
	c.onCDOTAUserMsg_OutpostCaptured = append(c.onCDOTAUserMsg_OutpostCaptured, fn)
}

// OnCDOTAUserMsg_OutpostGrantedXP registers a callback for EDotaUserMessages_DOTA_UM_OutpostGrantedXP
func (c *Callbacks) OnCDOTAUserMsg_OutpostGrantedXP(fn func(*dota.CDOTAUserMsg_OutpostGrantedXP) error) {
	c.onCDOTAUserMsg_OutpostGrantedXP = append(c.onCDOTAUserMsg_OutpostGrantedXP, fn)
}

// OnCDOTAUserMsg_MoveCameraToUnit registers a callback for EDotaUserMessages_DOTA_UM_MoveCameraToUnit
func (c *Callbacks) OnCDOTAUserMsg_MoveCameraToUnit(fn func(*dota.CDOTAUserMsg_MoveCameraToUnit) error) {
	c.onCDOTAUserMsg_MoveCameraToUnit = append(c.onCDOTAUserMsg_MoveCameraToUnit, fn)
}

// OnCDOTAUserMsg_PauseMinigameData registers a callback for EDotaUserMessages_DOTA_UM_PauseMinigameData
func (c *Callbacks) OnCDOTAUserMsg_PauseMinigameData(fn func(*dota.CDOTAUserMsg_PauseMinigameData) error) {
	c.onCDOTAUserMsg_PauseMinigameData = append(c.onCDOTAUserMsg_PauseMinigameData, fn)
}

// OnCDOTAUserMsg_VersusScene_PlayerBehavior registers a callback for EDotaUserMessages_DOTA_UM_VersusScene_PlayerBehavior
func (c *Callbacks) OnCDOTAUserMsg_VersusScene_PlayerBehavior(fn func(*dota.CDOTAUserMsg_VersusScene_PlayerBehavior) error) {
	c.onCDOTAUserMsg_VersusScene_PlayerBehavior = append(c.onCDOTAUserMsg_VersusScene_PlayerBehavior, fn)
}

// OnCDOTAUserMsg_QoP_ArcanaSummary registers a callback for EDotaUserMessages_DOTA_UM_QoP_ArcanaSummary
func (c *Callbacks) OnCDOTAUserMsg_QoP_ArcanaSummary(fn func(*dota.CDOTAUserMsg_QoP_ArcanaSummary) error) {
	c.onCDOTAUserMsg_QoP_ArcanaSummary = append(c.onCDOTAUserMsg_QoP_ArcanaSummary, fn)
}

// OnCDOTAUserMsg_HotPotato_Created registers a callback for EDotaUserMessages_DOTA_UM_HotPotato_Created
func (c *Callbacks) OnCDOTAUserMsg_HotPotato_Created(fn func(*dota.CDOTAUserMsg_HotPotato_Created) error) {
	c.onCDOTAUserMsg_HotPotato_Created = append(c.onCDOTAUserMsg_HotPotato_Created, fn)
}

// OnCDOTAUserMsg_HotPotato_Exploded registers a callback for EDotaUserMessages_DOTA_UM_HotPotato_Exploded
func (c *Callbacks) OnCDOTAUserMsg_HotPotato_Exploded(fn func(*dota.CDOTAUserMsg_HotPotato_Exploded) error) {
	c.onCDOTAUserMsg_HotPotato_Exploded = append(c.onCDOTAUserMsg_HotPotato_Exploded, fn)
}

// OnCDOTAUserMsg_WK_Arcana_Progress registers a callback for EDotaUserMessages_DOTA_UM_WK_Arcana_Progress
func (c *Callbacks) OnCDOTAUserMsg_WK_Arcana_Progress(fn func(*dota.CDOTAUserMsg_WK_Arcana_Progress) error) {
	c.onCDOTAUserMsg_WK_Arcana_Progress = append(c.onCDOTAUserMsg_WK_Arcana_Progress, fn)
}

// OnCDOTAUserMsg_GuildChallenge_Progress registers a callback for EDotaUserMessages_DOTA_UM_GuildChallenge_Progress
func (c *Callbacks) OnCDOTAUserMsg_GuildChallenge_Progress(fn func(*dota.CDOTAUserMsg_GuildChallenge_Progress) error) {
	c.onCDOTAUserMsg_GuildChallenge_Progress = append(c.onCDOTAUserMsg_GuildChallenge_Progress, fn)
}

// OnCDOTAUserMsg_WRArcanaProgress registers a callback for EDotaUserMessages_DOTA_UM_WRArcanaProgress
func (c *Callbacks) OnCDOTAUserMsg_WRArcanaProgress(fn func(*dota.CDOTAUserMsg_WRArcanaProgress) error) {
	c.onCDOTAUserMsg_WRArcanaProgress = append(c.onCDOTAUserMsg_WRArcanaProgress, fn)
}

// OnCDOTAUserMsg_WRArcanaSummary registers a callback for EDotaUserMessages_DOTA_UM_WRArcanaSummary
func (c *Callbacks) OnCDOTAUserMsg_WRArcanaSummary(fn func(*dota.CDOTAUserMsg_WRArcanaSummary) error) {
	c.onCDOTAUserMsg_WRArcanaSummary = append(c.onCDOTAUserMsg_WRArcanaSummary, fn)
}

// OnCDOTAUserMsg_EmptyItemSlotAlert registers a callback for EDotaUserMessages_DOTA_UM_EmptyItemSlotAlert
func (c *Callbacks) OnCDOTAUserMsg_EmptyItemSlotAlert(fn func(*dota.CDOTAUserMsg_EmptyItemSlotAlert) error) {
	c.onCDOTAUserMsg_EmptyItemSlotAlert = append(c.onCDOTAUserMsg_EmptyItemSlotAlert, fn)
}

// OnCDOTAUserMsg_AghsStatusAlert registers a callback for EDotaUserMessages_DOTA_UM_AghsStatusAlert
func (c *Callbacks) OnCDOTAUserMsg_AghsStatusAlert(fn func(*dota.CDOTAUserMsg_AghsStatusAlert) error) {
	c.onCDOTAUserMsg_AghsStatusAlert = append(c.onCDOTAUserMsg_AghsStatusAlert, fn)
}

// OnCDOTAUserMsg_PingConfirmation registers a callback for EDotaUserMessages_DOTA_UM_PingConfirmation
func (c *Callbacks) OnCDOTAUserMsg_PingConfirmation(fn func(*dota.CDOTAUserMsg_PingConfirmation) error) {
	c.onCDOTAUserMsg_PingConfirmation = append(c.onCDOTAUserMsg_PingConfirmation, fn)
}

// OnCDOTAUserMsg_MutedPlayers registers a callback for EDotaUserMessages_DOTA_UM_MutedPlayers
func (c *Callbacks) OnCDOTAUserMsg_MutedPlayers(fn func(*dota.CDOTAUserMsg_MutedPlayers) error) {
	c.onCDOTAUserMsg_MutedPlayers = append(c.onCDOTAUserMsg_MutedPlayers, fn)
}

// OnCDOTAUserMsg_ContextualTip registers a callback for EDotaUserMessages_DOTA_UM_ContextualTip
func (c *Callbacks) OnCDOTAUserMsg_ContextualTip(fn func(*dota.CDOTAUserMsg_ContextualTip) error) {
	c.onCDOTAUserMsg_ContextualTip = append(c.onCDOTAUserMsg_ContextualTip, fn)
}

// OnCDOTAUserMsg_ChatMessage registers a callback for EDotaUserMessages_DOTA_UM_ChatMessage
func (c *Callbacks) OnCDOTAUserMsg_ChatMessage(fn func(*dota.CDOTAUserMsg_ChatMessage) error) {
	c.onCDOTAUserMsg_ChatMessage = append(c.onCDOTAUserMsg_ChatMessage, fn)
}

// OnCDOTAUserMsg_NeutralCampAlert registers a callback for EDotaUserMessages_DOTA_UM_NeutralCampAlert
func (c *Callbacks) OnCDOTAUserMsg_NeutralCampAlert(fn func(*dota.CDOTAUserMsg_NeutralCampAlert) error) {
	c.onCDOTAUserMsg_NeutralCampAlert = append(c.onCDOTAUserMsg_NeutralCampAlert, fn)
}

// OnCDOTAUserMsg_RockPaperScissorsStarted registers a callback for EDotaUserMessages_DOTA_UM_RockPaperScissorsStarted
func (c *Callbacks) OnCDOTAUserMsg_RockPaperScissorsStarted(fn func(*dota.CDOTAUserMsg_RockPaperScissorsStarted) error) {
	c.onCDOTAUserMsg_RockPaperScissorsStarted = append(c.onCDOTAUserMsg_RockPaperScissorsStarted, fn)
}

// OnCDOTAUserMsg_RockPaperScissorsFinished registers a callback for EDotaUserMessages_DOTA_UM_RockPaperScissorsFinished
func (c *Callbacks) OnCDOTAUserMsg_RockPaperScissorsFinished(fn func(*dota.CDOTAUserMsg_RockPaperScissorsFinished) error) {
	c.onCDOTAUserMsg_RockPaperScissorsFinished = append(c.onCDOTAUserMsg_RockPaperScissorsFinished, fn)
}

// OnCDOTAUserMsg_DuelOpponentKilled registers a callback for EDotaUserMessages_DOTA_UM_DuelOpponentKilled
func (c *Callbacks) OnCDOTAUserMsg_DuelOpponentKilled(fn func(*dota.CDOTAUserMsg_DuelOpponentKilled) error) {
	c.onCDOTAUserMsg_DuelOpponentKilled = append(c.onCDOTAUserMsg_DuelOpponentKilled, fn)
}

// OnCDOTAUserMsg_DuelAccepted registers a callback for EDotaUserMessages_DOTA_UM_DuelAccepted
func (c *Callbacks) OnCDOTAUserMsg_DuelAccepted(fn func(*dota.CDOTAUserMsg_DuelAccepted) error) {
	c.onCDOTAUserMsg_DuelAccepted = append(c.onCDOTAUserMsg_DuelAccepted, fn)
}

// OnCDOTAUserMsg_DuelRequested registers a callback for EDotaUserMessages_DOTA_UM_DuelRequested
func (c *Callbacks) OnCDOTAUserMsg_DuelRequested(fn func(*dota.CDOTAUserMsg_DuelRequested) error) {
	c.onCDOTAUserMsg_DuelRequested = append(c.onCDOTAUserMsg_DuelRequested, fn)
}

// OnCDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled registers a callback for EDotaUserMessages_DOTA_UM_MuertaReleaseEvent_AssignedTargetKilled
func (c *Callbacks) OnCDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled(fn func(*dota.CDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled) error) {
	c.onCDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled = append(c.onCDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled, fn)
}

// OnCDOTAUserMsg_PlayerDraftSuggestPick registers a callback for EDotaUserMessages_DOTA_UM_PlayerDraftSuggestPick
func (c *Callbacks) OnCDOTAUserMsg_PlayerDraftSuggestPick(fn func(*dota.CDOTAUserMsg_PlayerDraftSuggestPick) error) {
	c.onCDOTAUserMsg_PlayerDraftSuggestPick = append(c.onCDOTAUserMsg_PlayerDraftSuggestPick, fn)
}

// OnCDOTAUserMsg_PlayerDraftPick registers a callback for EDotaUserMessages_DOTA_UM_PlayerDraftPick
func (c *Callbacks) OnCDOTAUserMsg_PlayerDraftPick(fn func(*dota.CDOTAUserMsg_PlayerDraftPick) error) {
	c.onCDOTAUserMsg_PlayerDraftPick = append(c.onCDOTAUserMsg_PlayerDraftPick, fn)
}

// OnCDOTAUserMsg_UpdateLinearProjectileCPData registers a callback for EDotaUserMessages_DOTA_UM_UpdateLinearProjectileCPData
func (c *Callbacks) OnCDOTAUserMsg_UpdateLinearProjectileCPData(fn func(*dota.CDOTAUserMsg_UpdateLinearProjectileCPData) error) {
	c.onCDOTAUserMsg_UpdateLinearProjectileCPData = append(c.onCDOTAUserMsg_UpdateLinearProjectileCPData, fn)
}

// OnCDOTAUserMsg_GiftPlayer registers a callback for EDotaUserMessages_DOTA_UM_GiftPlayer
func (c *Callbacks) OnCDOTAUserMsg_GiftPlayer(fn func(*dota.CDOTAUserMsg_GiftPlayer) error) {
	c.onCDOTAUserMsg_GiftPlayer = append(c.onCDOTAUserMsg_GiftPlayer, fn)
}

// OnCDOTAUserMsg_FacetPing registers a callback for EDotaUserMessages_DOTA_UM_FacetPing
func (c *Callbacks) OnCDOTAUserMsg_FacetPing(fn func(*dota.CDOTAUserMsg_FacetPing) error) {
	c.onCDOTAUserMsg_FacetPing = append(c.onCDOTAUserMsg_FacetPing, fn)
}

// OnCDOTAUserMsg_InnatePing registers a callback for EDotaUserMessages_DOTA_UM_InnatePing
func (c *Callbacks) OnCDOTAUserMsg_InnatePing(fn func(*dota.CDOTAUserMsg_InnatePing) error) {
	c.onCDOTAUserMsg_InnatePing = append(c.onCDOTAUserMsg_InnatePing, fn)
}

// OnCDOTAUserMsg_RoshanTimer registers a callback for EDotaUserMessages_DOTA_UM_RoshanTimer
func (c *Callbacks) OnCDOTAUserMsg_RoshanTimer(fn func(*dota.CDOTAUserMsg_RoshanTimer) error) {
	c.onCDOTAUserMsg_RoshanTimer = append(c.onCDOTAUserMsg_RoshanTimer, fn)
}

// OnCDOTAUserMsg_NeutralCraftAvailable registers a callback for EDotaUserMessages_DOTA_UM_NeutralCraftAvailable
func (c *Callbacks) OnCDOTAUserMsg_NeutralCraftAvailable(fn func(*dota.CDOTAUserMsg_NeutralCraftAvailable) error) {
	c.onCDOTAUserMsg_NeutralCraftAvailable = append(c.onCDOTAUserMsg_NeutralCraftAvailable, fn)
}

// OnCDOTAUserMsg_TimerAlert registers a callback for EDotaUserMessages_DOTA_UM_TimerAlert
func (c *Callbacks) OnCDOTAUserMsg_TimerAlert(fn func(*dota.CDOTAUserMsg_TimerAlert) error) {
	c.onCDOTAUserMsg_TimerAlert = append(c.onCDOTAUserMsg_TimerAlert, fn)
}

// OnCDOTAUserMsg_MadstoneAlert registers a callback for EDotaUserMessages_DOTA_UM_MadstoneAlert
func (c *Callbacks) OnCDOTAUserMsg_MadstoneAlert(fn func(*dota.CDOTAUserMsg_MadstoneAlert) error) {
	c.onCDOTAUserMsg_MadstoneAlert = append(c.onCDOTAUserMsg_MadstoneAlert, fn)
}

func (c *Callbacks) callByDemoType(t int32, buf []byte) error {
	switch t {
	case 0: // dota.EDemoCommands_DEM_Stop
		if c.onCDemoStop == nil {
			return nil
		}

		msg := &dota.CDemoStop{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoStop {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 1: // dota.EDemoCommands_DEM_FileHeader
		if c.onCDemoFileHeader == nil {
			return nil
		}

		msg := &dota.CDemoFileHeader{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFileHeader {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 2: // dota.EDemoCommands_DEM_FileInfo
		if c.onCDemoFileInfo == nil {
			return nil
		}

		msg := &dota.CDemoFileInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFileInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 3: // dota.EDemoCommands_DEM_SyncTick
		if c.onCDemoSyncTick == nil {
			return nil
		}

		msg := &dota.CDemoSyncTick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSyncTick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 4: // dota.EDemoCommands_DEM_SendTables
		if c.onCDemoSendTables == nil {
			return nil
		}

		msg := &dota.CDemoSendTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSendTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 5: // dota.EDemoCommands_DEM_ClassInfo
		if c.onCDemoClassInfo == nil {
			return nil
		}

		msg := &dota.CDemoClassInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoClassInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 6: // dota.EDemoCommands_DEM_StringTables
		if c.onCDemoStringTables == nil {
			return nil
		}

		msg := &dota.CDemoStringTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoStringTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 7: // dota.EDemoCommands_DEM_Packet
		if c.onCDemoPacket == nil {
			return nil
		}

		msg := &dota.CDemoPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 8: // dota.EDemoCommands_DEM_SignonPacket
		if c.onCDemoSignonPacket == nil {
			return nil
		}

		msg := &dota.CDemoPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSignonPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 9: // dota.EDemoCommands_DEM_ConsoleCmd
		if c.onCDemoConsoleCmd == nil {
			return nil
		}

		msg := &dota.CDemoConsoleCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoConsoleCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 10: // dota.EDemoCommands_DEM_CustomData
		if c.onCDemoCustomData == nil {
			return nil
		}

		msg := &dota.CDemoCustomData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoCustomData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 11: // dota.EDemoCommands_DEM_CustomDataCallbacks
		if c.onCDemoCustomDataCallbacks == nil {
			return nil
		}

		msg := &dota.CDemoCustomDataCallbacks{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoCustomDataCallbacks {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 12: // dota.EDemoCommands_DEM_UserCmd
		if c.onCDemoUserCmd == nil {
			return nil
		}

		msg := &dota.CDemoUserCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoUserCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 13: // dota.EDemoCommands_DEM_FullPacket
		if c.onCDemoFullPacket == nil {
			return nil
		}

		msg := &dota.CDemoFullPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFullPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 14: // dota.EDemoCommands_DEM_SaveGame
		if c.onCDemoSaveGame == nil {
			return nil
		}

		msg := &dota.CDemoSaveGame{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSaveGame {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 15: // dota.EDemoCommands_DEM_SpawnGroups
		if c.onCDemoSpawnGroups == nil {
			return nil
		}

		msg := &dota.CDemoSpawnGroups{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSpawnGroups {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 16: // dota.EDemoCommands_DEM_AnimationData
		if c.onCDemoAnimationData == nil {
			return nil
		}

		msg := &dota.CDemoAnimationData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoAnimationData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 17: // dota.EDemoCommands_DEM_AnimationHeader
		if c.onCDemoAnimationHeader == nil {
			return nil
		}

		msg := &dota.CDemoAnimationHeader{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoAnimationHeader {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 18: // dota.EDemoCommands_DEM_Recovery
		if c.onCDemoRecovery == nil {
			return nil
		}

		msg := &dota.CDemoRecovery{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoRecovery {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	}

	if v(1) {
		_debugf("warning: no demo type %d found", t)
	}

	return nil
}

func (c *Callbacks) callByPacketType(t int32, buf []byte) error {
	switch t {
	case 0: // dota.NET_Messages_net_NOP
		if c.onCNETMsg_NOP == nil {
			return nil
		}

		msg := &dota.CNETMsg_NOP{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_NOP {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 3: // dota.NET_Messages_net_SplitScreenUser
		if c.onCNETMsg_SplitScreenUser == nil {
			return nil
		}

		msg := &dota.CNETMsg_SplitScreenUser{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SplitScreenUser {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 4: // dota.NET_Messages_net_Tick
		if c.onCNETMsg_Tick == nil {
			return nil
		}

		msg := &dota.CNETMsg_Tick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_Tick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 5: // dota.NET_Messages_net_StringCmd
		if c.onCNETMsg_StringCmd == nil {
			return nil
		}

		msg := &dota.CNETMsg_StringCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_StringCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 6: // dota.NET_Messages_net_SetConVar
		if c.onCNETMsg_SetConVar == nil {
			return nil
		}

		msg := &dota.CNETMsg_SetConVar{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SetConVar {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 7: // dota.NET_Messages_net_SignonState
		if c.onCNETMsg_SignonState == nil {
			return nil
		}

		msg := &dota.CNETMsg_SignonState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SignonState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 8: // dota.NET_Messages_net_SpawnGroup_Load
		if c.onCNETMsg_SpawnGroup_Load == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_Load{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_Load {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 9: // dota.NET_Messages_net_SpawnGroup_ManifestUpdate
		if c.onCNETMsg_SpawnGroup_ManifestUpdate == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_ManifestUpdate{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_ManifestUpdate {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 11: // dota.NET_Messages_net_SpawnGroup_SetCreationTick
		if c.onCNETMsg_SpawnGroup_SetCreationTick == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_SetCreationTick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_SetCreationTick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 12: // dota.NET_Messages_net_SpawnGroup_Unload
		if c.onCNETMsg_SpawnGroup_Unload == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_Unload{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_Unload {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 13: // dota.NET_Messages_net_SpawnGroup_LoadCompleted
		if c.onCNETMsg_SpawnGroup_LoadCompleted == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_LoadCompleted{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_LoadCompleted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 15: // dota.NET_Messages_net_DebugOverlay
		if c.onCNETMsg_DebugOverlay == nil {
			return nil
		}

		msg := &dota.CNETMsg_DebugOverlay{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_DebugOverlay {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 40: // dota.SVC_Messages_svc_ServerInfo
		if c.onCSVCMsg_ServerInfo == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ServerInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ServerInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 41: // dota.SVC_Messages_svc_FlattenedSerializer
		if c.onCSVCMsg_FlattenedSerializer == nil {
			return nil
		}

		msg := &dota.CSVCMsg_FlattenedSerializer{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_FlattenedSerializer {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 42: // dota.SVC_Messages_svc_ClassInfo
		if c.onCSVCMsg_ClassInfo == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ClassInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ClassInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 43: // dota.SVC_Messages_svc_SetPause
		if c.onCSVCMsg_SetPause == nil {
			return nil
		}

		msg := &dota.CSVCMsg_SetPause{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SetPause {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 44: // dota.SVC_Messages_svc_CreateStringTable
		if c.onCSVCMsg_CreateStringTable == nil {
			return nil
		}

		msg := &dota.CSVCMsg_CreateStringTable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_CreateStringTable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 45: // dota.SVC_Messages_svc_UpdateStringTable
		if c.onCSVCMsg_UpdateStringTable == nil {
			return nil
		}

		msg := &dota.CSVCMsg_UpdateStringTable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_UpdateStringTable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 46: // dota.SVC_Messages_svc_VoiceInit
		if c.onCSVCMsg_VoiceInit == nil {
			return nil
		}

		msg := &dota.CSVCMsg_VoiceInit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_VoiceInit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 47: // dota.SVC_Messages_svc_VoiceData
		if c.onCSVCMsg_VoiceData == nil {
			return nil
		}

		msg := &dota.CSVCMsg_VoiceData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_VoiceData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 48: // dota.SVC_Messages_svc_Print
		if c.onCSVCMsg_Print == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Print{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Print {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 49: // dota.SVC_Messages_svc_Sounds
		if c.onCSVCMsg_Sounds == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Sounds{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Sounds {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 50: // dota.SVC_Messages_svc_SetView
		if c.onCSVCMsg_SetView == nil {
			return nil
		}

		msg := &dota.CSVCMsg_SetView{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SetView {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 51: // dota.SVC_Messages_svc_ClearAllStringTables
		if c.onCSVCMsg_ClearAllStringTables == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ClearAllStringTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ClearAllStringTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 52: // dota.SVC_Messages_svc_CmdKeyValues
		if c.onCSVCMsg_CmdKeyValues == nil {
			return nil
		}

		msg := &dota.CSVCMsg_CmdKeyValues{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_CmdKeyValues {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 53: // dota.SVC_Messages_svc_BSPDecal
		if c.onCSVCMsg_BSPDecal == nil {
			return nil
		}

		msg := &dota.CSVCMsg_BSPDecal{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_BSPDecal {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 54: // dota.SVC_Messages_svc_SplitScreen
		if c.onCSVCMsg_SplitScreen == nil {
			return nil
		}

		msg := &dota.CSVCMsg_SplitScreen{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SplitScreen {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 55: // dota.SVC_Messages_svc_PacketEntities
		if c.onCSVCMsg_PacketEntities == nil {
			return nil
		}

		msg := &dota.CSVCMsg_PacketEntities{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PacketEntities {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 56: // dota.SVC_Messages_svc_Prefetch
		if c.onCSVCMsg_Prefetch == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Prefetch{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Prefetch {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 57: // dota.SVC_Messages_svc_Menu
		if c.onCSVCMsg_Menu == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Menu{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Menu {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 58: // dota.SVC_Messages_svc_GetCvarValue
		if c.onCSVCMsg_GetCvarValue == nil {
			return nil
		}

		msg := &dota.CSVCMsg_GetCvarValue{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_GetCvarValue {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 59: // dota.SVC_Messages_svc_StopSound
		if c.onCSVCMsg_StopSound == nil {
			return nil
		}

		msg := &dota.CSVCMsg_StopSound{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_StopSound {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 60: // dota.SVC_Messages_svc_PeerList
		if c.onCSVCMsg_PeerList == nil {
			return nil
		}

		msg := &dota.CSVCMsg_PeerList{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PeerList {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 61: // dota.SVC_Messages_svc_PacketReliable
		if c.onCSVCMsg_PacketReliable == nil {
			return nil
		}

		msg := &dota.CSVCMsg_PacketReliable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PacketReliable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 62: // dota.SVC_Messages_svc_HLTVStatus
		if c.onCSVCMsg_HLTVStatus == nil {
			return nil
		}

		msg := &dota.CSVCMsg_HLTVStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_HLTVStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 63: // dota.SVC_Messages_svc_ServerSteamID
		if c.onCSVCMsg_ServerSteamID == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ServerSteamID{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ServerSteamID {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 70: // dota.SVC_Messages_svc_FullFrameSplit
		if c.onCSVCMsg_FullFrameSplit == nil {
			return nil
		}

		msg := &dota.CSVCMsg_FullFrameSplit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_FullFrameSplit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 71: // dota.SVC_Messages_svc_RconServerDetails
		if c.onCSVCMsg_RconServerDetails == nil {
			return nil
		}

		msg := &dota.CSVCMsg_RconServerDetails{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_RconServerDetails {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 72: // dota.SVC_Messages_svc_UserMessage
		if c.onCSVCMsg_UserMessage == nil {
			return nil
		}

		msg := &dota.CSVCMsg_UserMessage{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_UserMessage {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 74: // dota.SVC_Messages_svc_Broadcast_Command
		if c.onCSVCMsg_Broadcast_Command == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Broadcast_Command{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Broadcast_Command {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 75: // dota.SVC_Messages_svc_HltvFixupOperatorStatus
		if c.onCSVCMsg_HltvFixupOperatorStatus == nil {
			return nil
		}

		msg := &dota.CSVCMsg_HltvFixupOperatorStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_HltvFixupOperatorStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 101: // dota.EBaseUserMessages_UM_AchievementEvent
		if c.onCUserMessageAchievementEvent == nil {
			return nil
		}

		msg := &dota.CUserMessageAchievementEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAchievementEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 102: // dota.EBaseUserMessages_UM_CloseCaption
		if c.onCUserMessageCloseCaption == nil {
			return nil
		}

		msg := &dota.CUserMessageCloseCaption{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaption {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 103: // dota.EBaseUserMessages_UM_CloseCaptionDirect
		if c.onCUserMessageCloseCaptionDirect == nil {
			return nil
		}

		msg := &dota.CUserMessageCloseCaptionDirect{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaptionDirect {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 104: // dota.EBaseUserMessages_UM_CurrentTimescale
		if c.onCUserMessageCurrentTimescale == nil {
			return nil
		}

		msg := &dota.CUserMessageCurrentTimescale{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCurrentTimescale {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 105: // dota.EBaseUserMessages_UM_DesiredTimescale
		if c.onCUserMessageDesiredTimescale == nil {
			return nil
		}

		msg := &dota.CUserMessageDesiredTimescale{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageDesiredTimescale {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 106: // dota.EBaseUserMessages_UM_Fade
		if c.onCUserMessageFade == nil {
			return nil
		}

		msg := &dota.CUserMessageFade{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageFade {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 107: // dota.EBaseUserMessages_UM_GameTitle
		if c.onCUserMessageGameTitle == nil {
			return nil
		}

		msg := &dota.CUserMessageGameTitle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageGameTitle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 110: // dota.EBaseUserMessages_UM_HudMsg
		if c.onCUserMessageHudMsg == nil {
			return nil
		}

		msg := &dota.CUserMessageHudMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHudMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 111: // dota.EBaseUserMessages_UM_HudText
		if c.onCUserMessageHudText == nil {
			return nil
		}

		msg := &dota.CUserMessageHudText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHudText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 113: // dota.EBaseUserMessages_UM_ColoredText
		if c.onCUserMessageColoredText == nil {
			return nil
		}

		msg := &dota.CUserMessageColoredText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageColoredText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 114: // dota.EBaseUserMessages_UM_RequestState
		if c.onCUserMessageRequestState == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 115: // dota.EBaseUserMessages_UM_ResetHUD
		if c.onCUserMessageResetHUD == nil {
			return nil
		}

		msg := &dota.CUserMessageResetHUD{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageResetHUD {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 116: // dota.EBaseUserMessages_UM_Rumble
		if c.onCUserMessageRumble == nil {
			return nil
		}

		msg := &dota.CUserMessageRumble{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRumble {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 117: // dota.EBaseUserMessages_UM_SayText
		if c.onCUserMessageSayText == nil {
			return nil
		}

		msg := &dota.CUserMessageSayText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 118: // dota.EBaseUserMessages_UM_SayText2
		if c.onCUserMessageSayText2 == nil {
			return nil
		}

		msg := &dota.CUserMessageSayText2{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayText2 {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 119: // dota.EBaseUserMessages_UM_SayTextChannel
		if c.onCUserMessageSayTextChannel == nil {
			return nil
		}

		msg := &dota.CUserMessageSayTextChannel{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayTextChannel {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 120: // dota.EBaseUserMessages_UM_Shake
		if c.onCUserMessageShake == nil {
			return nil
		}

		msg := &dota.CUserMessageShake{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShake {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 121: // dota.EBaseUserMessages_UM_ShakeDir
		if c.onCUserMessageShakeDir == nil {
			return nil
		}

		msg := &dota.CUserMessageShakeDir{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShakeDir {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 122: // dota.EBaseUserMessages_UM_WaterShake
		if c.onCUserMessageWaterShake == nil {
			return nil
		}

		msg := &dota.CUserMessageWaterShake{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageWaterShake {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 124: // dota.EBaseUserMessages_UM_TextMsg
		if c.onCUserMessageTextMsg == nil {
			return nil
		}

		msg := &dota.CUserMessageTextMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageTextMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 125: // dota.EBaseUserMessages_UM_ScreenTilt
		if c.onCUserMessageScreenTilt == nil {
			return nil
		}

		msg := &dota.CUserMessageScreenTilt{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageScreenTilt {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 128: // dota.EBaseUserMessages_UM_VoiceMask
		if c.onCUserMessageVoiceMask == nil {
			return nil
		}

		msg := &dota.CUserMessageVoiceMask{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageVoiceMask {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 130: // dota.EBaseUserMessages_UM_SendAudio
		if c.onCUserMessageSendAudio == nil {
			return nil
		}

		msg := &dota.CUserMessageSendAudio{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSendAudio {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 131: // dota.EBaseUserMessages_UM_ItemPickup
		if c.onCUserMessageItemPickup == nil {
			return nil
		}

		msg := &dota.CUserMessageItemPickup{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageItemPickup {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 132: // dota.EBaseUserMessages_UM_AmmoDenied
		if c.onCUserMessageAmmoDenied == nil {
			return nil
		}

		msg := &dota.CUserMessageAmmoDenied{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAmmoDenied {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 134: // dota.EBaseUserMessages_UM_ShowMenu
		if c.onCUserMessageShowMenu == nil {
			return nil
		}

		msg := &dota.CUserMessageShowMenu{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShowMenu {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 135: // dota.EBaseUserMessages_UM_CreditsMsg
		if c.onCUserMessageCreditsMsg == nil {
			return nil
		}

		msg := &dota.CUserMessageCreditsMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCreditsMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 136: // dota.EBaseEntityMessages_EM_PlayJingle
		if c.onCEntityMessagePlayJingle == nil {
			return nil
		}

		msg := &dota.CEntityMessagePlayJingle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessagePlayJingle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 137: // dota.EBaseEntityMessages_EM_ScreenOverlay
		if c.onCEntityMessageScreenOverlay == nil {
			return nil
		}

		msg := &dota.CEntityMessageScreenOverlay{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageScreenOverlay {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 138: // dota.EBaseEntityMessages_EM_RemoveAllDecals
		if c.onCEntityMessageRemoveAllDecals == nil {
			return nil
		}

		msg := &dota.CEntityMessageRemoveAllDecals{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageRemoveAllDecals {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 139: // dota.EBaseEntityMessages_EM_PropagateForce
		if c.onCEntityMessagePropagateForce == nil {
			return nil
		}

		msg := &dota.CEntityMessagePropagateForce{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessagePropagateForce {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 140: // dota.EBaseEntityMessages_EM_DoSpark
		if c.onCEntityMessageDoSpark == nil {
			return nil
		}

		msg := &dota.CEntityMessageDoSpark{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageDoSpark {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 141: // dota.EBaseEntityMessages_EM_FixAngle
		if c.onCEntityMessageFixAngle == nil {
			return nil
		}

		msg := &dota.CEntityMessageFixAngle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageFixAngle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 142: // dota.EBaseUserMessages_UM_CloseCaptionPlaceholder
		if c.onCUserMessageCloseCaptionPlaceholder == nil {
			return nil
		}

		msg := &dota.CUserMessageCloseCaptionPlaceholder{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaptionPlaceholder {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 143: // dota.EBaseUserMessages_UM_CameraTransition
		if c.onCUserMessageCameraTransition == nil {
			return nil
		}

		msg := &dota.CUserMessageCameraTransition{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCameraTransition {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 144: // dota.EBaseUserMessages_UM_AudioParameter
		if c.onCUserMessageAudioParameter == nil {
			return nil
		}

		msg := &dota.CUserMessageAudioParameter{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAudioParameter {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 150: // dota.EBaseUserMessages_UM_HapticsManagerPulse
		if c.onCUserMessageHapticsManagerPulse == nil {
			return nil
		}

		msg := &dota.CUserMessageHapticsManagerPulse{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHapticsManagerPulse {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 151: // dota.EBaseUserMessages_UM_HapticsManagerEffect
		if c.onCUserMessageHapticsManagerEffect == nil {
			return nil
		}

		msg := &dota.CUserMessageHapticsManagerEffect{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHapticsManagerEffect {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 153: // dota.EBaseUserMessages_UM_UpdateCssClasses
		if c.onCUserMessageUpdateCssClasses == nil {
			return nil
		}

		msg := &dota.CUserMessageUpdateCssClasses{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageUpdateCssClasses {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 154: // dota.EBaseUserMessages_UM_ServerFrameTime
		if c.onCUserMessageServerFrameTime == nil {
			return nil
		}

		msg := &dota.CUserMessageServerFrameTime{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageServerFrameTime {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 155: // dota.EBaseUserMessages_UM_LagCompensationError
		if c.onCUserMessageLagCompensationError == nil {
			return nil
		}

		msg := &dota.CUserMessageLagCompensationError{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageLagCompensationError {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 156: // dota.EBaseUserMessages_UM_RequestDllStatus
		if c.onCUserMessageRequestDllStatus == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestDllStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestDllStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 157: // dota.EBaseUserMessages_UM_RequestUtilAction
		if c.onCUserMessageRequestUtilAction == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestUtilAction{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestUtilAction {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 160: // dota.EBaseUserMessages_UM_RequestInventory
		if c.onCUserMessageRequestInventory == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestInventory{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestInventory {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 162: // dota.EBaseUserMessages_UM_RequestDiagnostic
		if c.onCUserMessageRequestDiagnostic == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestDiagnostic{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestDiagnostic {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 200: // dota.EBaseGameEvents_GE_VDebugGameSessionIDEvent
		if c.onCMsgVDebugGameSessionIDEvent == nil {
			return nil
		}

		msg := &dota.CMsgVDebugGameSessionIDEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgVDebugGameSessionIDEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 201: // dota.EBaseGameEvents_GE_PlaceDecalEvent
		if c.onCMsgPlaceDecalEvent == nil {
			return nil
		}

		msg := &dota.CMsgPlaceDecalEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgPlaceDecalEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 202: // dota.EBaseGameEvents_GE_ClearWorldDecalsEvent
		if c.onCMsgClearWorldDecalsEvent == nil {
			return nil
		}

		msg := &dota.CMsgClearWorldDecalsEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearWorldDecalsEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 203: // dota.EBaseGameEvents_GE_ClearEntityDecalsEvent
		if c.onCMsgClearEntityDecalsEvent == nil {
			return nil
		}

		msg := &dota.CMsgClearEntityDecalsEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearEntityDecalsEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 204: // dota.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
		if c.onCMsgClearDecalsForSkeletonInstanceEvent == nil {
			return nil
		}

		msg := &dota.CMsgClearDecalsForSkeletonInstanceEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearDecalsForSkeletonInstanceEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 205: // dota.EBaseGameEvents_GE_Source1LegacyGameEventList
		if c.onCMsgSource1LegacyGameEventList == nil {
			return nil
		}

		msg := &dota.CMsgSource1LegacyGameEventList{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyGameEventList {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 206: // dota.EBaseGameEvents_GE_Source1LegacyListenEvents
		if c.onCMsgSource1LegacyListenEvents == nil {
			return nil
		}

		msg := &dota.CMsgSource1LegacyListenEvents{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyListenEvents {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 207: // dota.EBaseGameEvents_GE_Source1LegacyGameEvent
		if c.onCMsgSource1LegacyGameEvent == nil {
			return nil
		}

		msg := &dota.CMsgSource1LegacyGameEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyGameEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 208: // dota.EBaseGameEvents_GE_SosStartSoundEvent
		if c.onCMsgSosStartSoundEvent == nil {
			return nil
		}

		msg := &dota.CMsgSosStartSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStartSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 209: // dota.EBaseGameEvents_GE_SosStopSoundEvent
		if c.onCMsgSosStopSoundEvent == nil {
			return nil
		}

		msg := &dota.CMsgSosStopSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStopSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 210: // dota.EBaseGameEvents_GE_SosSetSoundEventParams
		if c.onCMsgSosSetSoundEventParams == nil {
			return nil
		}

		msg := &dota.CMsgSosSetSoundEventParams{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosSetSoundEventParams {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 211: // dota.EBaseGameEvents_GE_SosSetLibraryStackFields
		if c.onCMsgSosSetLibraryStackFields == nil {
			return nil
		}

		msg := &dota.CMsgSosSetLibraryStackFields{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosSetLibraryStackFields {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 212: // dota.EBaseGameEvents_GE_SosStopSoundEventHash
		if c.onCMsgSosStopSoundEventHash == nil {
			return nil
		}

		msg := &dota.CMsgSosStopSoundEventHash{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStopSoundEventHash {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 465: // dota.EDotaUserMessages_DOTA_UM_AIDebugLine
		if c.onCDOTAUserMsg_AIDebugLine == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_AIDebugLine{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_AIDebugLine {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 466: // dota.EDotaUserMessages_DOTA_UM_ChatEvent
		if c.onCDOTAUserMsg_ChatEvent == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ChatEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ChatEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 467: // dota.EDotaUserMessages_DOTA_UM_CombatHeroPositions
		if c.onCDOTAUserMsg_CombatHeroPositions == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CombatHeroPositions{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CombatHeroPositions {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 470: // dota.EDotaUserMessages_DOTA_UM_CombatLogBulkData
		if c.onCDOTAUserMsg_CombatLogBulkData == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CombatLogBulkData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CombatLogBulkData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 471: // dota.EDotaUserMessages_DOTA_UM_CreateLinearProjectile
		if c.onCDOTAUserMsg_CreateLinearProjectile == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CreateLinearProjectile{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CreateLinearProjectile {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 472: // dota.EDotaUserMessages_DOTA_UM_DestroyLinearProjectile
		if c.onCDOTAUserMsg_DestroyLinearProjectile == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DestroyLinearProjectile{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DestroyLinearProjectile {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 473: // dota.EDotaUserMessages_DOTA_UM_DodgeTrackingProjectiles
		if c.onCDOTAUserMsg_DodgeTrackingProjectiles == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DodgeTrackingProjectiles{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DodgeTrackingProjectiles {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 474: // dota.EDotaUserMessages_DOTA_UM_GlobalLightColor
		if c.onCDOTAUserMsg_GlobalLightColor == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_GlobalLightColor{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_GlobalLightColor {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 475: // dota.EDotaUserMessages_DOTA_UM_GlobalLightDirection
		if c.onCDOTAUserMsg_GlobalLightDirection == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_GlobalLightDirection{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_GlobalLightDirection {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 476: // dota.EDotaUserMessages_DOTA_UM_InvalidCommand
		if c.onCDOTAUserMsg_InvalidCommand == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_InvalidCommand{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_InvalidCommand {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 477: // dota.EDotaUserMessages_DOTA_UM_LocationPing
		if c.onCDOTAUserMsg_LocationPing == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_LocationPing{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_LocationPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 478: // dota.EDotaUserMessages_DOTA_UM_MapLine
		if c.onCDOTAUserMsg_MapLine == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MapLine{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MapLine {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 479: // dota.EDotaUserMessages_DOTA_UM_MiniKillCamInfo
		if c.onCDOTAUserMsg_MiniKillCamInfo == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MiniKillCamInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MiniKillCamInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 480: // dota.EDotaUserMessages_DOTA_UM_MinimapDebugPoint
		if c.onCDOTAUserMsg_MinimapDebugPoint == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MinimapDebugPoint{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MinimapDebugPoint {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 481: // dota.EDotaUserMessages_DOTA_UM_MinimapEvent
		if c.onCDOTAUserMsg_MinimapEvent == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MinimapEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MinimapEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 482: // dota.EDotaUserMessages_DOTA_UM_NevermoreRequiem
		if c.onCDOTAUserMsg_NevermoreRequiem == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_NevermoreRequiem{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_NevermoreRequiem {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 483: // dota.EDotaUserMessages_DOTA_UM_OverheadEvent
		if c.onCDOTAUserMsg_OverheadEvent == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_OverheadEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_OverheadEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 484: // dota.EDotaUserMessages_DOTA_UM_SetNextAutobuyItem
		if c.onCDOTAUserMsg_SetNextAutobuyItem == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SetNextAutobuyItem{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SetNextAutobuyItem {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 485: // dota.EDotaUserMessages_DOTA_UM_SharedCooldown
		if c.onCDOTAUserMsg_SharedCooldown == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SharedCooldown{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SharedCooldown {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 486: // dota.EDotaUserMessages_DOTA_UM_SpectatorPlayerClick
		if c.onCDOTAUserMsg_SpectatorPlayerClick == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SpectatorPlayerClick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SpectatorPlayerClick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 487: // dota.EDotaUserMessages_DOTA_UM_TutorialTipInfo
		if c.onCDOTAUserMsg_TutorialTipInfo == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TutorialTipInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TutorialTipInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 488: // dota.EDotaUserMessages_DOTA_UM_UnitEvent
		if c.onCDOTAUserMsg_UnitEvent == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_UnitEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_UnitEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 490: // dota.EDotaUserMessages_DOTA_UM_BotChat
		if c.onCDOTAUserMsg_BotChat == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_BotChat{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_BotChat {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 491: // dota.EDotaUserMessages_DOTA_UM_HudError
		if c.onCDOTAUserMsg_HudError == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HudError{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HudError {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 492: // dota.EDotaUserMessages_DOTA_UM_ItemPurchased
		if c.onCDOTAUserMsg_ItemPurchased == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ItemPurchased{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ItemPurchased {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 493: // dota.EDotaUserMessages_DOTA_UM_Ping
		if c.onCDOTAUserMsg_Ping == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_Ping{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_Ping {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 494: // dota.EDotaUserMessages_DOTA_UM_ItemFound
		if c.onCDOTAUserMsg_ItemFound == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ItemFound{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ItemFound {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 496: // dota.EDotaUserMessages_DOTA_UM_SwapVerify
		if c.onCDOTAUserMsg_SwapVerify == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SwapVerify{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SwapVerify {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 497: // dota.EDotaUserMessages_DOTA_UM_WorldLine
		if c.onCDOTAUserMsg_WorldLine == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_WorldLine{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_WorldLine {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 498: // dota.EDotaUserMessages_DOTA_UM_TournamentDrop
		if c.onCMsgGCToClientTournamentItemDrop == nil {
			return nil
		}

		msg := &dota.CMsgGCToClientTournamentItemDrop{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgGCToClientTournamentItemDrop {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 499: // dota.EDotaUserMessages_DOTA_UM_ItemAlert
		if c.onCDOTAUserMsg_ItemAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ItemAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ItemAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 500: // dota.EDotaUserMessages_DOTA_UM_HalloweenDrops
		if c.onCDOTAUserMsg_HalloweenDrops == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HalloweenDrops{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HalloweenDrops {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 501: // dota.EDotaUserMessages_DOTA_UM_ChatWheel
		if c.onCDOTAUserMsg_ChatWheel == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ChatWheel{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ChatWheel {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 502: // dota.EDotaUserMessages_DOTA_UM_ReceivedXmasGift
		if c.onCDOTAUserMsg_ReceivedXmasGift == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ReceivedXmasGift{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ReceivedXmasGift {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 503: // dota.EDotaUserMessages_DOTA_UM_UpdateSharedContent
		if c.onCDOTAUserMsg_UpdateSharedContent == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_UpdateSharedContent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_UpdateSharedContent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 504: // dota.EDotaUserMessages_DOTA_UM_TutorialRequestExp
		if c.onCDOTAUserMsg_TutorialRequestExp == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TutorialRequestExp{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TutorialRequestExp {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 505: // dota.EDotaUserMessages_DOTA_UM_TutorialPingMinimap
		if c.onCDOTAUserMsg_TutorialPingMinimap == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TutorialPingMinimap{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TutorialPingMinimap {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 506: // dota.EDotaUserMessages_DOTA_UM_GamerulesStateChanged
		if c.onCDOTAUserMsg_GamerulesStateChanged == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_GamerulesStateChanged{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_GamerulesStateChanged {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 507: // dota.EDotaUserMessages_DOTA_UM_ShowSurvey
		if c.onCDOTAUserMsg_ShowSurvey == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ShowSurvey{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ShowSurvey {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 508: // dota.EDotaUserMessages_DOTA_UM_TutorialFade
		if c.onCDOTAUserMsg_TutorialFade == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TutorialFade{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TutorialFade {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 509: // dota.EDotaUserMessages_DOTA_UM_AddQuestLogEntry
		if c.onCDOTAUserMsg_AddQuestLogEntry == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_AddQuestLogEntry{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_AddQuestLogEntry {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 510: // dota.EDotaUserMessages_DOTA_UM_SendStatPopup
		if c.onCDOTAUserMsg_SendStatPopup == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SendStatPopup{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SendStatPopup {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 511: // dota.EDotaUserMessages_DOTA_UM_TutorialFinish
		if c.onCDOTAUserMsg_TutorialFinish == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TutorialFinish{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TutorialFinish {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 512: // dota.EDotaUserMessages_DOTA_UM_SendRoshanPopup
		if c.onCDOTAUserMsg_SendRoshanPopup == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SendRoshanPopup{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SendRoshanPopup {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 513: // dota.EDotaUserMessages_DOTA_UM_SendGenericToolTip
		if c.onCDOTAUserMsg_SendGenericToolTip == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SendGenericToolTip{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SendGenericToolTip {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 514: // dota.EDotaUserMessages_DOTA_UM_SendFinalGold
		if c.onCDOTAUserMsg_SendFinalGold == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SendFinalGold{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SendFinalGold {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 515: // dota.EDotaUserMessages_DOTA_UM_CustomMsg
		if c.onCDOTAUserMsg_CustomMsg == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CustomMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CustomMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 516: // dota.EDotaUserMessages_DOTA_UM_CoachHUDPing
		if c.onCDOTAUserMsg_CoachHUDPing == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CoachHUDPing{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CoachHUDPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 517: // dota.EDotaUserMessages_DOTA_UM_ClientLoadGridNav
		if c.onCDOTAUserMsg_ClientLoadGridNav == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ClientLoadGridNav{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ClientLoadGridNav {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 518: // dota.EDotaUserMessages_DOTA_UM_TE_Projectile
		if c.onCDOTAUserMsg_TE_Projectile == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TE_Projectile{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TE_Projectile {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 519: // dota.EDotaUserMessages_DOTA_UM_TE_ProjectileLoc
		if c.onCDOTAUserMsg_TE_ProjectileLoc == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TE_ProjectileLoc{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TE_ProjectileLoc {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 520: // dota.EDotaUserMessages_DOTA_UM_TE_DotaBloodImpact
		if c.onCDOTAUserMsg_TE_DotaBloodImpact == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TE_DotaBloodImpact{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TE_DotaBloodImpact {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 521: // dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimation
		if c.onCDOTAUserMsg_TE_UnitAnimation == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TE_UnitAnimation{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TE_UnitAnimation {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 522: // dota.EDotaUserMessages_DOTA_UM_TE_UnitAnimationEnd
		if c.onCDOTAUserMsg_TE_UnitAnimationEnd == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TE_UnitAnimationEnd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TE_UnitAnimationEnd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 523: // dota.EDotaUserMessages_DOTA_UM_AbilityPing
		if c.onCDOTAUserMsg_AbilityPing == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_AbilityPing{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_AbilityPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 524: // dota.EDotaUserMessages_DOTA_UM_ShowGenericPopup
		if c.onCDOTAUserMsg_ShowGenericPopup == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ShowGenericPopup{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ShowGenericPopup {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 525: // dota.EDotaUserMessages_DOTA_UM_VoteStart
		if c.onCDOTAUserMsg_VoteStart == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_VoteStart{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_VoteStart {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 526: // dota.EDotaUserMessages_DOTA_UM_VoteUpdate
		if c.onCDOTAUserMsg_VoteUpdate == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_VoteUpdate{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_VoteUpdate {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 527: // dota.EDotaUserMessages_DOTA_UM_VoteEnd
		if c.onCDOTAUserMsg_VoteEnd == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_VoteEnd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_VoteEnd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 528: // dota.EDotaUserMessages_DOTA_UM_BoosterState
		if c.onCDOTAUserMsg_BoosterState == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_BoosterState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_BoosterState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 529: // dota.EDotaUserMessages_DOTA_UM_WillPurchaseAlert
		if c.onCDOTAUserMsg_WillPurchaseAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_WillPurchaseAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_WillPurchaseAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 530: // dota.EDotaUserMessages_DOTA_UM_TutorialMinimapPosition
		if c.onCDOTAUserMsg_TutorialMinimapPosition == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TutorialMinimapPosition{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TutorialMinimapPosition {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 532: // dota.EDotaUserMessages_DOTA_UM_AbilitySteal
		if c.onCDOTAUserMsg_AbilitySteal == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_AbilitySteal{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_AbilitySteal {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 533: // dota.EDotaUserMessages_DOTA_UM_CourierKilledAlert
		if c.onCDOTAUserMsg_CourierKilledAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CourierKilledAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CourierKilledAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 534: // dota.EDotaUserMessages_DOTA_UM_EnemyItemAlert
		if c.onCDOTAUserMsg_EnemyItemAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_EnemyItemAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_EnemyItemAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 535: // dota.EDotaUserMessages_DOTA_UM_StatsMatchDetails
		if c.onCDOTAUserMsg_StatsMatchDetails == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_StatsMatchDetails{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_StatsMatchDetails {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 536: // dota.EDotaUserMessages_DOTA_UM_MiniTaunt
		if c.onCDOTAUserMsg_MiniTaunt == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MiniTaunt{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MiniTaunt {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 537: // dota.EDotaUserMessages_DOTA_UM_BuyBackStateAlert
		if c.onCDOTAUserMsg_BuyBackStateAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_BuyBackStateAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_BuyBackStateAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 538: // dota.EDotaUserMessages_DOTA_UM_SpeechBubble
		if c.onCDOTAUserMsg_SpeechBubble == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SpeechBubble{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SpeechBubble {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 539: // dota.EDotaUserMessages_DOTA_UM_CustomHeaderMessage
		if c.onCDOTAUserMsg_CustomHeaderMessage == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CustomHeaderMessage{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CustomHeaderMessage {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 540: // dota.EDotaUserMessages_DOTA_UM_QuickBuyAlert
		if c.onCDOTAUserMsg_QuickBuyAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_QuickBuyAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_QuickBuyAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 541: // dota.EDotaUserMessages_DOTA_UM_StatsHeroDetails
		if c.onCDOTAUserMsg_StatsHeroMinuteDetails == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_StatsHeroMinuteDetails{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_StatsHeroMinuteDetails {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 543: // dota.EDotaUserMessages_DOTA_UM_ModifierAlert
		if c.onCDOTAUserMsg_ModifierAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ModifierAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ModifierAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 544: // dota.EDotaUserMessages_DOTA_UM_HPManaAlert
		if c.onCDOTAUserMsg_HPManaAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HPManaAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HPManaAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 545: // dota.EDotaUserMessages_DOTA_UM_GlyphAlert
		if c.onCDOTAUserMsg_GlyphAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_GlyphAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_GlyphAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 546: // dota.EDotaUserMessages_DOTA_UM_BeastChat
		if c.onCDOTAUserMsg_BeastChat == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_BeastChat{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_BeastChat {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 547: // dota.EDotaUserMessages_DOTA_UM_SpectatorPlayerUnitOrders
		if c.onCDOTAUserMsg_SpectatorPlayerUnitOrders == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SpectatorPlayerUnitOrders{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SpectatorPlayerUnitOrders {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 548: // dota.EDotaUserMessages_DOTA_UM_CustomHudElement_Create
		if c.onCDOTAUserMsg_CustomHudElement_Create == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CustomHudElement_Create{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CustomHudElement_Create {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 549: // dota.EDotaUserMessages_DOTA_UM_CustomHudElement_Modify
		if c.onCDOTAUserMsg_CustomHudElement_Modify == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CustomHudElement_Modify{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CustomHudElement_Modify {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 550: // dota.EDotaUserMessages_DOTA_UM_CustomHudElement_Destroy
		if c.onCDOTAUserMsg_CustomHudElement_Destroy == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CustomHudElement_Destroy{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CustomHudElement_Destroy {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 551: // dota.EDotaUserMessages_DOTA_UM_CompendiumState
		if c.onCDOTAUserMsg_CompendiumState == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_CompendiumState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_CompendiumState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 552: // dota.EDotaUserMessages_DOTA_UM_ProjectionAbility
		if c.onCDOTAUserMsg_ProjectionAbility == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ProjectionAbility{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ProjectionAbility {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 553: // dota.EDotaUserMessages_DOTA_UM_ProjectionEvent
		if c.onCDOTAUserMsg_ProjectionEvent == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ProjectionEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ProjectionEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 554: // dota.EDotaUserMessages_DOTA_UM_CombatLogDataHLTV
		if c.onCMsgDOTACombatLogEntry == nil {
			return nil
		}

		msg := &dota.CMsgDOTACombatLogEntry{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgDOTACombatLogEntry {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 555: // dota.EDotaUserMessages_DOTA_UM_XPAlert
		if c.onCDOTAUserMsg_XPAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_XPAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_XPAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 556: // dota.EDotaUserMessages_DOTA_UM_UpdateQuestProgress
		if c.onCDOTAUserMsg_UpdateQuestProgress == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_UpdateQuestProgress{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_UpdateQuestProgress {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 557: // dota.EDotaUserMessages_DOTA_UM_MatchMetadata
		if c.onCDOTAMatchMetadataFile == nil {
			return nil
		}

		msg := &dota.CDOTAMatchMetadataFile{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAMatchMetadataFile {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 559: // dota.EDotaUserMessages_DOTA_UM_QuestStatus
		if c.onCDOTAUserMsg_QuestStatus == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_QuestStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_QuestStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 560: // dota.EDotaUserMessages_DOTA_UM_SuggestHeroPick
		if c.onCDOTAUserMsg_SuggestHeroPick == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SuggestHeroPick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SuggestHeroPick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 561: // dota.EDotaUserMessages_DOTA_UM_SuggestHeroRole
		if c.onCDOTAUserMsg_SuggestHeroRole == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SuggestHeroRole{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SuggestHeroRole {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 562: // dota.EDotaUserMessages_DOTA_UM_KillcamDamageTaken
		if c.onCDOTAUserMsg_KillcamDamageTaken == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_KillcamDamageTaken{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_KillcamDamageTaken {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 563: // dota.EDotaUserMessages_DOTA_UM_SelectPenaltyGold
		if c.onCDOTAUserMsg_SelectPenaltyGold == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SelectPenaltyGold{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SelectPenaltyGold {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 564: // dota.EDotaUserMessages_DOTA_UM_RollDiceResult
		if c.onCDOTAUserMsg_RollDiceResult == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_RollDiceResult{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_RollDiceResult {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 565: // dota.EDotaUserMessages_DOTA_UM_FlipCoinResult
		if c.onCDOTAUserMsg_FlipCoinResult == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_FlipCoinResult{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_FlipCoinResult {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 568: // dota.EDotaUserMessages_DOTA_UM_SendRoshanSpectatorPhase
		if c.onCDOTAUserMsg_SendRoshanSpectatorPhase == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SendRoshanSpectatorPhase{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SendRoshanSpectatorPhase {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 569: // dota.EDotaUserMessages_DOTA_UM_ChatWheelCooldown
		if c.onCDOTAUserMsg_ChatWheelCooldown == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ChatWheelCooldown{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ChatWheelCooldown {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 570: // dota.EDotaUserMessages_DOTA_UM_DismissAllStatPopups
		if c.onCDOTAUserMsg_DismissAllStatPopups == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DismissAllStatPopups{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DismissAllStatPopups {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 571: // dota.EDotaUserMessages_DOTA_UM_TE_DestroyProjectile
		if c.onCDOTAUserMsg_TE_DestroyProjectile == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TE_DestroyProjectile{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TE_DestroyProjectile {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 572: // dota.EDotaUserMessages_DOTA_UM_HeroRelicProgress
		if c.onCDOTAUserMsg_HeroRelicProgress == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HeroRelicProgress{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HeroRelicProgress {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 573: // dota.EDotaUserMessages_DOTA_UM_AbilityDraftRequestAbility
		if c.onCDOTAUserMsg_AbilityDraftRequestAbility == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_AbilityDraftRequestAbility{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_AbilityDraftRequestAbility {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 574: // dota.EDotaUserMessages_DOTA_UM_ItemSold
		if c.onCDOTAUserMsg_ItemSold == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ItemSold{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ItemSold {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 575: // dota.EDotaUserMessages_DOTA_UM_DamageReport
		if c.onCDOTAUserMsg_DamageReport == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DamageReport{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DamageReport {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 576: // dota.EDotaUserMessages_DOTA_UM_SalutePlayer
		if c.onCDOTAUserMsg_SalutePlayer == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_SalutePlayer{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_SalutePlayer {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 577: // dota.EDotaUserMessages_DOTA_UM_TipAlert
		if c.onCDOTAUserMsg_TipAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TipAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TipAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 578: // dota.EDotaUserMessages_DOTA_UM_ReplaceQueryUnit
		if c.onCDOTAUserMsg_ReplaceQueryUnit == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ReplaceQueryUnit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ReplaceQueryUnit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 579: // dota.EDotaUserMessages_DOTA_UM_EmptyTeleportAlert
		if c.onCDOTAUserMsg_EmptyTeleportAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_EmptyTeleportAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_EmptyTeleportAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 580: // dota.EDotaUserMessages_DOTA_UM_MarsArenaOfBloodAttack
		if c.onCDOTAUserMsg_MarsArenaOfBloodAttack == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MarsArenaOfBloodAttack{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MarsArenaOfBloodAttack {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 581: // dota.EDotaUserMessages_DOTA_UM_ESArcanaCombo
		if c.onCDOTAUserMsg_ESArcanaCombo == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ESArcanaCombo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ESArcanaCombo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 582: // dota.EDotaUserMessages_DOTA_UM_ESArcanaComboSummary
		if c.onCDOTAUserMsg_ESArcanaComboSummary == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ESArcanaComboSummary{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ESArcanaComboSummary {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 583: // dota.EDotaUserMessages_DOTA_UM_HighFiveLeftHanging
		if c.onCDOTAUserMsg_HighFiveLeftHanging == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HighFiveLeftHanging{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HighFiveLeftHanging {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 584: // dota.EDotaUserMessages_DOTA_UM_HighFiveCompleted
		if c.onCDOTAUserMsg_HighFiveCompleted == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HighFiveCompleted{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HighFiveCompleted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 585: // dota.EDotaUserMessages_DOTA_UM_ShovelUnearth
		if c.onCDOTAUserMsg_ShovelUnearth == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ShovelUnearth{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ShovelUnearth {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 587: // dota.EDotaUserMessages_DOTA_UM_RadarAlert
		if c.onCDOTAUserMsg_RadarAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_RadarAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_RadarAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 588: // dota.EDotaUserMessages_DOTA_UM_AllStarEvent
		if c.onCDOTAUserMsg_AllStarEvent == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_AllStarEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_AllStarEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 589: // dota.EDotaUserMessages_DOTA_UM_TalentTreeAlert
		if c.onCDOTAUserMsg_TalentTreeAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TalentTreeAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TalentTreeAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 590: // dota.EDotaUserMessages_DOTA_UM_QueuedOrderRemoved
		if c.onCDOTAUserMsg_QueuedOrderRemoved == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_QueuedOrderRemoved{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_QueuedOrderRemoved {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 591: // dota.EDotaUserMessages_DOTA_UM_DebugChallenge
		if c.onCDOTAUserMsg_DebugChallenge == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DebugChallenge{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DebugChallenge {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 592: // dota.EDotaUserMessages_DOTA_UM_OMArcanaCombo
		if c.onCDOTAUserMsg_OMArcanaCombo == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_OMArcanaCombo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_OMArcanaCombo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 593: // dota.EDotaUserMessages_DOTA_UM_FoundNeutralItem
		if c.onCDOTAUserMsg_FoundNeutralItem == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_FoundNeutralItem{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_FoundNeutralItem {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 594: // dota.EDotaUserMessages_DOTA_UM_OutpostCaptured
		if c.onCDOTAUserMsg_OutpostCaptured == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_OutpostCaptured{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_OutpostCaptured {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 595: // dota.EDotaUserMessages_DOTA_UM_OutpostGrantedXP
		if c.onCDOTAUserMsg_OutpostGrantedXP == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_OutpostGrantedXP{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_OutpostGrantedXP {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 596: // dota.EDotaUserMessages_DOTA_UM_MoveCameraToUnit
		if c.onCDOTAUserMsg_MoveCameraToUnit == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MoveCameraToUnit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MoveCameraToUnit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 597: // dota.EDotaUserMessages_DOTA_UM_PauseMinigameData
		if c.onCDOTAUserMsg_PauseMinigameData == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_PauseMinigameData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_PauseMinigameData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 598: // dota.EDotaUserMessages_DOTA_UM_VersusScene_PlayerBehavior
		if c.onCDOTAUserMsg_VersusScene_PlayerBehavior == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_VersusScene_PlayerBehavior{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_VersusScene_PlayerBehavior {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 600: // dota.EDotaUserMessages_DOTA_UM_QoP_ArcanaSummary
		if c.onCDOTAUserMsg_QoP_ArcanaSummary == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_QoP_ArcanaSummary{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_QoP_ArcanaSummary {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 601: // dota.EDotaUserMessages_DOTA_UM_HotPotato_Created
		if c.onCDOTAUserMsg_HotPotato_Created == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HotPotato_Created{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HotPotato_Created {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 602: // dota.EDotaUserMessages_DOTA_UM_HotPotato_Exploded
		if c.onCDOTAUserMsg_HotPotato_Exploded == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_HotPotato_Exploded{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_HotPotato_Exploded {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 603: // dota.EDotaUserMessages_DOTA_UM_WK_Arcana_Progress
		if c.onCDOTAUserMsg_WK_Arcana_Progress == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_WK_Arcana_Progress{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_WK_Arcana_Progress {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 604: // dota.EDotaUserMessages_DOTA_UM_GuildChallenge_Progress
		if c.onCDOTAUserMsg_GuildChallenge_Progress == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_GuildChallenge_Progress{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_GuildChallenge_Progress {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 605: // dota.EDotaUserMessages_DOTA_UM_WRArcanaProgress
		if c.onCDOTAUserMsg_WRArcanaProgress == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_WRArcanaProgress{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_WRArcanaProgress {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 606: // dota.EDotaUserMessages_DOTA_UM_WRArcanaSummary
		if c.onCDOTAUserMsg_WRArcanaSummary == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_WRArcanaSummary{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_WRArcanaSummary {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 607: // dota.EDotaUserMessages_DOTA_UM_EmptyItemSlotAlert
		if c.onCDOTAUserMsg_EmptyItemSlotAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_EmptyItemSlotAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_EmptyItemSlotAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 608: // dota.EDotaUserMessages_DOTA_UM_AghsStatusAlert
		if c.onCDOTAUserMsg_AghsStatusAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_AghsStatusAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_AghsStatusAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 609: // dota.EDotaUserMessages_DOTA_UM_PingConfirmation
		if c.onCDOTAUserMsg_PingConfirmation == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_PingConfirmation{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_PingConfirmation {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 610: // dota.EDotaUserMessages_DOTA_UM_MutedPlayers
		if c.onCDOTAUserMsg_MutedPlayers == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MutedPlayers{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MutedPlayers {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 611: // dota.EDotaUserMessages_DOTA_UM_ContextualTip
		if c.onCDOTAUserMsg_ContextualTip == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ContextualTip{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ContextualTip {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 612: // dota.EDotaUserMessages_DOTA_UM_ChatMessage
		if c.onCDOTAUserMsg_ChatMessage == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_ChatMessage{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_ChatMessage {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 613: // dota.EDotaUserMessages_DOTA_UM_NeutralCampAlert
		if c.onCDOTAUserMsg_NeutralCampAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_NeutralCampAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_NeutralCampAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 614: // dota.EDotaUserMessages_DOTA_UM_RockPaperScissorsStarted
		if c.onCDOTAUserMsg_RockPaperScissorsStarted == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_RockPaperScissorsStarted{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_RockPaperScissorsStarted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 615: // dota.EDotaUserMessages_DOTA_UM_RockPaperScissorsFinished
		if c.onCDOTAUserMsg_RockPaperScissorsFinished == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_RockPaperScissorsFinished{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_RockPaperScissorsFinished {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 616: // dota.EDotaUserMessages_DOTA_UM_DuelOpponentKilled
		if c.onCDOTAUserMsg_DuelOpponentKilled == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DuelOpponentKilled{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DuelOpponentKilled {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 617: // dota.EDotaUserMessages_DOTA_UM_DuelAccepted
		if c.onCDOTAUserMsg_DuelAccepted == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DuelAccepted{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DuelAccepted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 618: // dota.EDotaUserMessages_DOTA_UM_DuelRequested
		if c.onCDOTAUserMsg_DuelRequested == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_DuelRequested{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_DuelRequested {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 619: // dota.EDotaUserMessages_DOTA_UM_MuertaReleaseEvent_AssignedTargetKilled
		if c.onCDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MuertaReleaseEvent_AssignedTargetKilled {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 620: // dota.EDotaUserMessages_DOTA_UM_PlayerDraftSuggestPick
		if c.onCDOTAUserMsg_PlayerDraftSuggestPick == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_PlayerDraftSuggestPick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_PlayerDraftSuggestPick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 621: // dota.EDotaUserMessages_DOTA_UM_PlayerDraftPick
		if c.onCDOTAUserMsg_PlayerDraftPick == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_PlayerDraftPick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_PlayerDraftPick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 622: // dota.EDotaUserMessages_DOTA_UM_UpdateLinearProjectileCPData
		if c.onCDOTAUserMsg_UpdateLinearProjectileCPData == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_UpdateLinearProjectileCPData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_UpdateLinearProjectileCPData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 623: // dota.EDotaUserMessages_DOTA_UM_GiftPlayer
		if c.onCDOTAUserMsg_GiftPlayer == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_GiftPlayer{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_GiftPlayer {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 624: // dota.EDotaUserMessages_DOTA_UM_FacetPing
		if c.onCDOTAUserMsg_FacetPing == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_FacetPing{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_FacetPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 625: // dota.EDotaUserMessages_DOTA_UM_InnatePing
		if c.onCDOTAUserMsg_InnatePing == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_InnatePing{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_InnatePing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 626: // dota.EDotaUserMessages_DOTA_UM_RoshanTimer
		if c.onCDOTAUserMsg_RoshanTimer == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_RoshanTimer{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_RoshanTimer {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 627: // dota.EDotaUserMessages_DOTA_UM_NeutralCraftAvailable
		if c.onCDOTAUserMsg_NeutralCraftAvailable == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_NeutralCraftAvailable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_NeutralCraftAvailable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 628: // dota.EDotaUserMessages_DOTA_UM_TimerAlert
		if c.onCDOTAUserMsg_TimerAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_TimerAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_TimerAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 629: // dota.EDotaUserMessages_DOTA_UM_MadstoneAlert
		if c.onCDOTAUserMsg_MadstoneAlert == nil {
			return nil
		}

		msg := &dota.CDOTAUserMsg_MadstoneAlert{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDOTAUserMsg_MadstoneAlert {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	}

	return nil
}
