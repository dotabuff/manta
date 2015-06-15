//go:generate go run gen/game_event.go fixtures/game_events_list.pbmsg game_event_lookup.go

package manta

const (
	EGameEvent_ServerSpawn                        = 0
	EGameEvent_ServerPreShutdown                  = 1
	EGameEvent_ServerShutdown                     = 2
	EGameEvent_ServerCvar                         = 3
	EGameEvent_ServerMessage                      = 4
	EGameEvent_ServerAddban                       = 5
	EGameEvent_ServerRemoveban                    = 6
	EGameEvent_PlayerConnect                      = 7
	EGameEvent_PlayerInfo                         = 8
	EGameEvent_PlayerDisconnect                   = 9
	EGameEvent_PlayerActivate                     = 10
	EGameEvent_PlayerConnectFull                  = 11
	EGameEvent_PlayerSay                          = 12
	EGameEvent_PlayerFullUpdate                   = 13
	EGameEvent_TeamInfo                           = 14
	EGameEvent_TeamScore                          = 15
	EGameEvent_TeamplayBroadcastAudio             = 16
	EGameEvent_PlayerTeam                         = 17
	EGameEvent_PlayerClass                        = 18
	EGameEvent_PlayerDeath                        = 19
	EGameEvent_PlayerHurt                         = 20
	EGameEvent_PlayerChat                         = 21
	EGameEvent_PlayerScore                        = 22
	EGameEvent_PlayerSpawn                        = 23
	EGameEvent_PlayerShoot                        = 24
	EGameEvent_PlayerUse                          = 25
	EGameEvent_PlayerChangename                   = 26
	EGameEvent_PlayerHintmessage                  = 27
	EGameEvent_GameInit                           = 28
	EGameEvent_GameNewmap                         = 29
	EGameEvent_GameStart                          = 30
	EGameEvent_GameEnd                            = 31
	EGameEvent_RoundStart                         = 32
	EGameEvent_RoundEnd                           = 33
	EGameEvent_RoundStartPreEntity                = 34
	EGameEvent_TeamplayRoundStart                 = 35
	EGameEvent_HostnameChanged                    = 36
	EGameEvent_DifficultyChanged                  = 37
	EGameEvent_FinaleStart                        = 38
	EGameEvent_GameMessage                        = 39
	EGameEvent_BreakBreakable                     = 40
	EGameEvent_BreakProp                          = 41
	EGameEvent_NpcSpawned                         = 42
	EGameEvent_NpcReplaced                        = 43
	EGameEvent_EntityKilled                       = 44
	EGameEvent_EntityHurt                         = 45
	EGameEvent_BonusUpdated                       = 46
	EGameEvent_PlayerStatsUpdated                 = 47
	EGameEvent_AchievementEvent                   = 48
	EGameEvent_AchievementEarned                  = 49
	EGameEvent_AchievementWriteFailed             = 50
	EGameEvent_PhysgunPickup                      = 51
	EGameEvent_FlareIgniteNpc                     = 52
	EGameEvent_HelicopterGrenadePuntMiss          = 53
	EGameEvent_UserDataDownloaded                 = 54
	EGameEvent_RagdollDissolved                   = 55
	EGameEvent_GameinstructorDraw                 = 56
	EGameEvent_GameinstructorNodraw               = 57
	EGameEvent_MapTransition                      = 58
	EGameEvent_InstructorServerHintCreate         = 59
	EGameEvent_InstructorServerHintStop           = 60
	EGameEvent_ChatNewMessage                     = 61
	EGameEvent_ChatMembersChanged                 = 62
	EGameEvent_InventoryUpdated                   = 63
	EGameEvent_CartUpdated                        = 64
	EGameEvent_StorePricesheetUpdated             = 65
	EGameEvent_GcConnected                        = 66
	EGameEvent_ItemSchemaInitialized              = 67
	EGameEvent_DropRateModified                   = 68
	EGameEvent_EventTicketModified                = 69
	EGameEvent_ModifierEvent                      = 70
	EGameEvent_DotaPlayerKill                     = 71
	EGameEvent_DotaPlayerDeny                     = 72
	EGameEvent_DotaBarracksKill                   = 73
	EGameEvent_DotaTowerKill                      = 74
	EGameEvent_DotaEffigyKill                     = 75
	EGameEvent_DotaRoshanKill                     = 76
	EGameEvent_DotaCourierLost                    = 77
	EGameEvent_DotaCourierRespawned               = 78
	EGameEvent_DotaGlyphUsed                      = 79
	EGameEvent_DotaSuperCreeps                    = 80
	EGameEvent_DotaItemPurchase                   = 81
	EGameEvent_DotaItemGifted                     = 82
	EGameEvent_DotaRunePickup                     = 83
	EGameEvent_DotaRuneSpotted                    = 84
	EGameEvent_DotaItemSpotted                    = 85
	EGameEvent_DotaNoBattlePoints                 = 86
	EGameEvent_DotaChatInformational              = 87
	EGameEvent_DotaActionItem                     = 88
	EGameEvent_DotaChatBanNotification            = 89
	EGameEvent_DotaChatEvent                      = 90
	EGameEvent_DotaChatTimedReward                = 91
	EGameEvent_DotaPauseEvent                     = 92
	EGameEvent_DotaChatKillStreak                 = 93
	EGameEvent_DotaChatFirstBlood                 = 94
	EGameEvent_DotaChatAssassinAnnounce           = 95
	EGameEvent_DotaChatAssassinDenied             = 96
	EGameEvent_DotaChatAssassinSuccess            = 97
	EGameEvent_DotaPlayerUpdateHeroSelection      = 98
	EGameEvent_DotaPlayerUpdateSelectedUnit       = 99
	EGameEvent_DotaPlayerUpdateQueryUnit          = 100
	EGameEvent_DotaPlayerUpdateKillcamUnit        = 101
	EGameEvent_DotaPlayerTakeTowerDamage          = 102
	EGameEvent_DotaHudErrorMessage                = 103
	EGameEvent_DotaActionSuccess                  = 104
	EGameEvent_DotaStartingPositionChanged        = 105
	EGameEvent_DotaMoneyChanged                   = 106
	EGameEvent_DotaEnemyMoneyChanged              = 107
	EGameEvent_DotaPortraitUnitStatsChanged       = 108
	EGameEvent_DotaPortraitUnitModifiersChanged   = 109
	EGameEvent_DotaForcePortraitUpdate            = 110
	EGameEvent_DotaInventoryChanged               = 111
	EGameEvent_DotaItemPickedUp                   = 112
	EGameEvent_DotaInventoryItemChanged           = 113
	EGameEvent_DotaAbilityChanged                 = 114
	EGameEvent_DotaPortraitAbilityLayoutChanged   = 115
	EGameEvent_DotaInventoryItemAdded             = 116
	EGameEvent_DotaInventoryChangedQueryUnit      = 117
	EGameEvent_DotaLinkClicked                    = 118
	EGameEvent_DotaSetQuickBuy                    = 119
	EGameEvent_DotaQuickBuyChanged                = 120
	EGameEvent_DotaPlayerShopChanged              = 121
	EGameEvent_DotaPlayerShowKillcam              = 122
	EGameEvent_DotaPlayerShowMinikillcam          = 123
	EGameEvent_GcUserSessionCreated               = 124
	EGameEvent_TeamDataUpdated                    = 125
	EGameEvent_GuildDataUpdated                   = 126
	EGameEvent_GuildOpenPartiesUpdated            = 127
	EGameEvent_FantasyUpdated                     = 128
	EGameEvent_FantasyLeagueChanged               = 129
	EGameEvent_FantasyScoreInfoChanged            = 130
	EGameEvent_PlayerInfoUpdated                  = 131
	EGameEvent_PlayerInfoIndividualUpdated        = 132
	EGameEvent_GameRulesStateChange               = 133
	EGameEvent_MatchHistoryUpdated                = 134
	EGameEvent_MatchDetailsUpdated                = 135
	EGameEvent_LiveGamesUpdated                   = 136
	EGameEvent_RecentMatchesUpdated               = 137
	EGameEvent_NewsUpdated                        = 138
	EGameEvent_PersonaUpdated                     = 139
	EGameEvent_TournamentStateUpdated             = 140
	EGameEvent_PartyUpdated                       = 141
	EGameEvent_LobbyUpdated                       = 142
	EGameEvent_DashboardCachesCleared             = 143
	EGameEvent_LastHit                            = 144
	EGameEvent_PlayerCompletedGame                = 145
	EGameEvent_PlayerReconnected                  = 146
	EGameEvent_NommedTree                         = 147
	EGameEvent_DotaRuneActivatedServer            = 148
	EGameEvent_DotaPlayerGainedLevel              = 149
	EGameEvent_DotaPlayerLearnedAbility           = 150
	EGameEvent_DotaPlayerUsedAbility              = 151
	EGameEvent_DotaNonPlayerUsedAbility           = 152
	EGameEvent_DotaPlayerBeginCast                = 153
	EGameEvent_DotaNonPlayerBeginCast             = 154
	EGameEvent_DotaAbilityChannelFinished         = 155
	EGameEvent_DotaHoldoutReviveComplete          = 156
	EGameEvent_DotaPlayerKilled                   = 157
	EGameEvent_BindpanelOpen                      = 158
	EGameEvent_BindpanelClose                     = 159
	EGameEvent_KeybindChanged                     = 160
	EGameEvent_DotaItemDragBegin                  = 161
	EGameEvent_DotaItemDragEnd                    = 162
	EGameEvent_DotaShopItemDragBegin              = 163
	EGameEvent_DotaShopItemDragEnd                = 164
	EGameEvent_DotaItemPurchased                  = 165
	EGameEvent_DotaItemCombined                   = 166
	EGameEvent_DotaItemUsed                       = 167
	EGameEvent_DotaItemAutoPurchase               = 168
	EGameEvent_DotaUnitEvent                      = 169
	EGameEvent_DotaQuestStarted                   = 170
	EGameEvent_DotaQuestCompleted                 = 171
	EGameEvent_GameuiActivated                    = 172
	EGameEvent_GameuiHidden                       = 173
	EGameEvent_PlayerFullyjoined                  = 174
	EGameEvent_DotaSpectateHero                   = 175
	EGameEvent_DotaMatchDone                      = 176
	EGameEvent_DotaMatchDoneClient                = 177
	EGameEvent_SetInstructorGroupEnabled          = 178
	EGameEvent_JoinedChatChannel                  = 179
	EGameEvent_LeftChatChannel                    = 180
	EGameEvent_GcChatChannelListUpdated           = 181
	EGameEvent_TodayMessagesUpdated               = 182
	EGameEvent_FileDownloaded                     = 183
	EGameEvent_PlayerReportCountsUpdated          = 184
	EGameEvent_ScaleformFileDownloadComplete      = 185
	EGameEvent_ItemPurchased                      = 186
	EGameEvent_GcMismatchedVersion                = 187
	EGameEvent_DemoStop                           = 190
	EGameEvent_MapShutdown                        = 191
	EGameEvent_DotaWorkshopFileselected           = 192
	EGameEvent_DotaWorkshopFilecanceled           = 193
	EGameEvent_RichPresenceUpdated                = 194
	EGameEvent_DotaHeroRandom                     = 195
	EGameEvent_DotaRdChatTurn                     = 196
	EGameEvent_DotaFavoriteHeroesUpdated          = 197
	EGameEvent_ProfileOpened                      = 198
	EGameEvent_ProfileClosed                      = 199
	EGameEvent_ItemPreviewClosed                  = 200
	EGameEvent_DashboardSwitchedSection           = 201
	EGameEvent_DotaTournamentItemEvent            = 202
	EGameEvent_DotaHeroSwap                       = 203
	EGameEvent_DotaResetSuggestedItems            = 204
	EGameEvent_HalloweenHighScoreReceived         = 205
	EGameEvent_HalloweenPhaseEnd                  = 206
	EGameEvent_HalloweenHighScoreRequestFailed    = 207
	EGameEvent_DotaHudSkinChanged                 = 208
	EGameEvent_DotaInventoryPlayerGotItem         = 209
	EGameEvent_PlayerIsExperienced                = 210
	EGameEvent_PlayerIsNotexperienced             = 211
	EGameEvent_DotaTutorialLessonStart            = 212
	EGameEvent_DotaTutorialTaskAdvance            = 213
	EGameEvent_DotaTutorialShopToggled            = 214
	EGameEvent_MapLocationUpdated                 = 215
	EGameEvent_RichpresenceCustomUpdated          = 216
	EGameEvent_GameEndVisible                     = 217
	EGameEvent_AntiaddictionUpdate                = 218
	EGameEvent_HighlightHudElement                = 219
	EGameEvent_HideHighlightHudElement            = 220
	EGameEvent_IntroVideoFinished                 = 221
	EGameEvent_MatchmakingStatusVisibilityChanged = 222
	EGameEvent_PracticeLobbyVisibilityChanged     = 223
	EGameEvent_DotaCourierTransferItem            = 224
	EGameEvent_FullUiUnlocked                     = 225
	EGameEvent_HeroSelectorPreviewSet             = 227
	EGameEvent_AntiaddictionToast                 = 228
	EGameEvent_HeroPickerShown                    = 229
	EGameEvent_HeroPickerHidden                   = 230
	EGameEvent_DotaLocalQuickbuyChanged           = 231
	EGameEvent_ShowCenterMessage                  = 232
	EGameEvent_HudFlipChanged                     = 233
	EGameEvent_FrostyPointsUpdated                = 234
	EGameEvent_Defeated                           = 235
	EGameEvent_ResetDefeated                      = 236
	EGameEvent_BoosterStateUpdated                = 237
	EGameEvent_EventPointsUpdated                 = 238
	EGameEvent_LocalPlayerEventPoints             = 239
	EGameEvent_CustomGameDifficulty               = 240
	EGameEvent_TreeCut                            = 241
	EGameEvent_UgcDetailsArrived                  = 242
	EGameEvent_UgcSubscribed                      = 243
	EGameEvent_UgcUnsubscribed                    = 244
	EGameEvent_UgcDownloadRequested               = 245
	EGameEvent_UgcInstalled                       = 246
	EGameEvent_PrizepoolReceived                  = 247
	EGameEvent_MicrotransactionSuccess            = 248
	EGameEvent_DotaRubickAbilitySteal             = 249
	EGameEvent_CompendiumEventActionsLoaded       = 250
	EGameEvent_CompendiumSelectionsLoaded         = 251
	EGameEvent_CompendiumSetSelectionFailed       = 252
	EGameEvent_CompendiumTrophiesLoaded           = 253
	EGameEvent_CommunityCachedNamesUpdated        = 254
	EGameEvent_SpecItemPickup                     = 255
	EGameEvent_SpecAegisReclaimTime               = 256
	EGameEvent_AccountTrophiesChanged             = 257
	EGameEvent_AccountAllHeroChallengeChanged     = 258
	EGameEvent_TeamShowcaseUiUpdate               = 259
	EGameEvent_IngameEventsChanged                = 260
	EGameEvent_DotaMatchSignout                   = 261
	EGameEvent_DotaIllusionsCreated               = 262
	EGameEvent_DotaYearBeastKilled                = 263
	EGameEvent_DotaHeroUndoselection              = 264
	EGameEvent_DotaChallengeSocacheUpdated        = 265
	EGameEvent_PartyInvitesUpdated                = 266
	EGameEvent_LobbyInvitesUpdated                = 267
	EGameEvent_CustomGameModeListUpdated          = 268
	EGameEvent_CustomGameLobbyListUpdated         = 269
	EGameEvent_FriendLobbyListUpdated             = 270
	EGameEvent_DotaTeamPlayerListChanged          = 271
	EGameEvent_DotaPlayerDetailsChanged           = 272
	EGameEvent_PlayerProfileStatsUpdated          = 273
	EGameEvent_CustomGamePlayerCountUpdated       = 274
	EGameEvent_CustomGameFriendsPlayedUpdated     = 275
	EGameEvent_CustomGamesFriendsPlayUpdated      = 276
	EGameEvent_DotaPlayerUpdateAssignedHero       = 277
	EGameEvent_DotaPlayerHeroSelectionDirty       = 278
	EGameEvent_DotaNpcGoalReached                 = 279
	EGameEvent_HltvStatus                         = 280
	EGameEvent_HltvCameraman                      = 281
	EGameEvent_HltvRankCamera                     = 282
	EGameEvent_HltvRankEntity                     = 283
	EGameEvent_HltvFixed                          = 284
	EGameEvent_HltvChase                          = 285
	EGameEvent_HltvMessage                        = 286
	EGameEvent_HltvTitle                          = 287
	EGameEvent_HltvChat                           = 288
	EGameEvent_HltvVersioninfo                    = 289
	EGameEvent_DotaChaseHero                      = 290
	EGameEvent_DotaCombatlog                      = 291
	EGameEvent_DotaGameStateChange                = 292
	EGameEvent_DotaPlayerPickHero                 = 293
	EGameEvent_DotaTeamKillCredit                 = 294
)

type GameEventServerSpawn struct {
	Hostname   string `json:"hostname"`
	Address    string `json:"address"`
	Port       int32  `json:"port"`
	Game       string `json:"game"`
	Mapname    string `json:"mapname"`
	Addonname  string `json:"addonname"`
	Maxplayers int32  `json:"maxplayers"`
	Os         string `json:"os"`
	Dedicated  bool   `json:"dedicated"`
	Password   bool   `json:"password"`
}

type GameEventServerPreShutdown struct {
	Reason string `json:"reason"`
}

type GameEventServerShutdown struct {
	Reason string `json:"reason"`
}

type GameEventServerCvar struct {
	Cvarname  string `json:"cvarname"`
	Cvarvalue string `json:"cvarvalue"`
}

type GameEventServerMessage struct {
	Text string `json:"text"`
}

type GameEventServerAddban struct {
	Name      string `json:"name"`
	Userid    int32  `json:"userid"`
	Networkid string `json:"networkid"`
	Ip        string `json:"ip"`
	Duration  string `json:"duration"`
	By        string `json:"by"`
	Kicked    bool   `json:"kicked"`
}

type GameEventServerRemoveban struct {
	Networkid string `json:"networkid"`
	Ip        string `json:"ip"`
	By        string `json:"by"`
}

type GameEventPlayerConnect struct {
	Name      string `json:"name"`
	Index     int32  `json:"index"`
	Userid    int32  `json:"userid"`
	Networkid string `json:"networkid"`
	Address   string `json:"address"`
}

type GameEventPlayerInfo struct {
	Name      string `json:"name"`
	Index     int32  `json:"index"`
	Userid    int32  `json:"userid"`
	Networkid string `json:"networkid"`
	Bot       bool   `json:"bot"`
}

type GameEventPlayerDisconnect struct {
	Userid    int32  `json:"userid"`
	Reason    int32  `json:"reason"`
	Name      string `json:"name"`
	Networkid string `json:"networkid"`
}

type GameEventPlayerActivate struct {
	Userid int32 `json:"userid"`
}

type GameEventPlayerConnectFull struct {
	Userid int32 `json:"userid"`
	Index  int32 `json:"index"`
}

type GameEventPlayerSay struct {
	Userid int32  `json:"userid"`
	Text   string `json:"text"`
}

type GameEventPlayerFullUpdate struct {
	Userid int32 `json:"userid"`
	Count  int32 `json:"count"`
}

type GameEventTeamInfo struct {
	Teamid   int32  `json:"teamid"`
	Teamname string `json:"teamname"`
}

type GameEventTeamScore struct {
	Teamid int32 `json:"teamid"`
	Score  int32 `json:"score"`
}

type GameEventTeamplayBroadcastAudio struct {
	Team  int32  `json:"team"`
	Sound string `json:"sound"`
}

type GameEventPlayerTeam struct {
	Userid     int32 `json:"userid"`
	Team       int32 `json:"team"`
	Oldteam    int32 `json:"oldteam"`
	Disconnect bool  `json:"disconnect"`
	Autoteam   bool  `json:"autoteam"`
	Silent     bool  `json:"silent"`
}

type GameEventPlayerClass struct {
	Userid int32  `json:"userid"`
	Class  string `json:"class"`
}

type GameEventPlayerDeath struct {
	Userid   int32 `json:"userid"`
	Attacker int32 `json:"attacker"`
}

type GameEventPlayerHurt struct {
	Userid   int32 `json:"userid"`
	Attacker int32 `json:"attacker"`
	Health   int32 `json:"health"`
}

type GameEventPlayerChat struct {
	Teamonly bool   `json:"teamonly"`
	Userid   int32  `json:"userid"`
	Text     string `json:"text"`
}

type GameEventPlayerScore struct {
	Userid int32 `json:"userid"`
	Kills  int32 `json:"kills"`
	Deaths int32 `json:"deaths"`
	Score  int32 `json:"score"`
}

type GameEventPlayerSpawn struct {
	Userid int32 `json:"userid"`
}

type GameEventPlayerShoot struct {
	Userid int32 `json:"userid"`
	Weapon int32 `json:"weapon"`
	Mode   int32 `json:"mode"`
}

type GameEventPlayerUse struct {
	Userid int32 `json:"userid"`
	Entity int32 `json:"entity"`
}

type GameEventPlayerChangename struct {
	Userid  int32  `json:"userid"`
	Oldname string `json:"oldname"`
	Newname string `json:"newname"`
}

type GameEventPlayerHintmessage struct {
	Hintmessage string `json:"hintmessage"`
}

type GameEventGameInit struct {
}

type GameEventGameNewmap struct {
	Mapname string `json:"mapname"`
}

type GameEventGameStart struct {
	Roundslimit int32  `json:"roundslimit"`
	Timelimit   int32  `json:"timelimit"`
	Fraglimit   int32  `json:"fraglimit"`
	Objective   string `json:"objective"`
}

type GameEventGameEnd struct {
	Winner int32 `json:"winner"`
}

type GameEventRoundStart struct {
	Timelimit int32  `json:"timelimit"`
	Fraglimit int32  `json:"fraglimit"`
	Objective string `json:"objective"`
}

type GameEventRoundEnd struct {
	Winner  int32  `json:"winner"`
	Reason  int32  `json:"reason"`
	Message string `json:"message"`
}

type GameEventRoundStartPreEntity struct {
}

type GameEventTeamplayRoundStart struct {
	FullReset bool `json:"full_reset"`
}

type GameEventHostnameChanged struct {
	Hostname string `json:"hostname"`
}

type GameEventDifficultyChanged struct {
	NewDifficulty int32  `json:"newDifficulty"`
	OldDifficulty int32  `json:"oldDifficulty"`
	StrDifficulty string `json:"strDifficulty"`
}

type GameEventFinaleStart struct {
	Rushes int32 `json:"rushes"`
}

type GameEventGameMessage struct {
	Target int32  `json:"target"`
	Text   string `json:"text"`
}

type GameEventBreakBreakable struct {
	Entindex int32 `json:"entindex"`
	Userid   int32 `json:"userid"`
	Material int32 `json:"material"`
}

type GameEventBreakProp struct {
	Entindex int32 `json:"entindex"`
	Userid   int32 `json:"userid"`
}

type GameEventNpcSpawned struct {
	Entindex int32 `json:"entindex"`
}

type GameEventNpcReplaced struct {
	OldEntindex int32 `json:"old_entindex"`
	NewEntindex int32 `json:"new_entindex"`
}

type GameEventEntityKilled struct {
	EntindexKilled    int32 `json:"entindex_killed"`
	EntindexAttacker  int32 `json:"entindex_attacker"`
	EntindexInflictor int32 `json:"entindex_inflictor"`
	Damagebits        int32 `json:"damagebits"`
}

type GameEventEntityHurt struct {
	EntindexKilled    int32 `json:"entindex_killed"`
	EntindexAttacker  int32 `json:"entindex_attacker"`
	EntindexInflictor int32 `json:"entindex_inflictor"`
	Damagebits        int32 `json:"damagebits"`
}

type GameEventBonusUpdated struct {
	Numadvanced int32 `json:"numadvanced"`
	Numbronze   int32 `json:"numbronze"`
	Numsilver   int32 `json:"numsilver"`
	Numgold     int32 `json:"numgold"`
}

type GameEventPlayerStatsUpdated struct {
	Forceupload bool `json:"forceupload"`
}

type GameEventAchievementEvent struct {
	AchievementName string `json:"achievement_name"`
	CurVal          int32  `json:"cur_val"`
	MaxVal          int32  `json:"max_val"`
}

type GameEventAchievementEarned struct {
	Player      int32 `json:"player"`
	Achievement int32 `json:"achievement"`
}

type GameEventAchievementWriteFailed struct {
}

type GameEventPhysgunPickup struct {
	Entindex int32 `json:"entindex"`
}

type GameEventFlareIgniteNpc struct {
	Entindex int32 `json:"entindex"`
}

type GameEventHelicopterGrenadePuntMiss struct {
}

type GameEventUserDataDownloaded struct {
}

type GameEventRagdollDissolved struct {
	Entindex int32 `json:"entindex"`
}

type GameEventGameinstructorDraw struct {
}

type GameEventGameinstructorNodraw struct {
}

type GameEventMapTransition struct {
}

type GameEventInstructorServerHintCreate struct {
	HintName              string  `json:"hint_name"`
	HintReplaceKey        string  `json:"hint_replace_key"`
	HintTarget            int32   `json:"hint_target"`
	HintActivatorUserid   int32   `json:"hint_activator_userid"`
	HintTimeout           int32   `json:"hint_timeout"`
	HintIconOnscreen      string  `json:"hint_icon_onscreen"`
	HintIconOffscreen     string  `json:"hint_icon_offscreen"`
	HintCaption           string  `json:"hint_caption"`
	HintActivatorCaption  string  `json:"hint_activator_caption"`
	HintColor             string  `json:"hint_color"`
	HintIconOffset        float32 `json:"hint_icon_offset"`
	HintRange             float32 `json:"hint_range"`
	HintFlags             int32   `json:"hint_flags"`
	HintBinding           string  `json:"hint_binding"`
	HintAllowNodrawTarget bool    `json:"hint_allow_nodraw_target"`
	HintNooffscreen       bool    `json:"hint_nooffscreen"`
	HintForcecaption      bool    `json:"hint_forcecaption"`
	HintLocalPlayerOnly   bool    `json:"hint_local_player_only"`
}

type GameEventInstructorServerHintStop struct {
	HintName string `json:"hint_name"`
}

type GameEventChatNewMessage struct {
	Channel int32 `json:"channel"`
}

type GameEventChatMembersChanged struct {
	Channel int32 `json:"channel"`
}

type GameEventInventoryUpdated struct {
	Itemdef int32 `json:"itemdef"`
	Itemid  int32 `json:"itemid"`
}

type GameEventCartUpdated struct {
}

type GameEventStorePricesheetUpdated struct {
}

type GameEventGcConnected struct {
}

type GameEventItemSchemaInitialized struct {
}

type GameEventDropRateModified struct {
}

type GameEventEventTicketModified struct {
}

type GameEventModifierEvent struct {
	Eventname string `json:"eventname"`
	Caster    int32  `json:"caster"`
	Ability   int32  `json:"ability"`
}

type GameEventDotaPlayerKill struct {
	VictimUserid  int32 `json:"victim_userid"`
	Killer1Userid int32 `json:"killer1_userid"`
	Killer2Userid int32 `json:"killer2_userid"`
	Killer3Userid int32 `json:"killer3_userid"`
	Killer4Userid int32 `json:"killer4_userid"`
	Killer5Userid int32 `json:"killer5_userid"`
	Bounty        int32 `json:"bounty"`
	Neutral       int32 `json:"neutral"`
	Greevil       int32 `json:"greevil"`
}

type GameEventDotaPlayerDeny struct {
	KillerUserid int32 `json:"killer_userid"`
	VictimUserid int32 `json:"victim_userid"`
}

type GameEventDotaBarracksKill struct {
	BarracksId int32 `json:"barracks_id"`
}

type GameEventDotaTowerKill struct {
	KillerUserid int32 `json:"killer_userid"`
	Teamnumber   int32 `json:"teamnumber"`
	Gold         int32 `json:"gold"`
}

type GameEventDotaEffigyKill struct {
	OwnerUserid int32 `json:"owner_userid"`
}

type GameEventDotaRoshanKill struct {
	Teamnumber int32 `json:"teamnumber"`
	Gold       int32 `json:"gold"`
}

type GameEventDotaCourierLost struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaCourierRespawned struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaGlyphUsed struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaSuperCreeps struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaItemPurchase struct {
	Userid int32 `json:"userid"`
	Itemid int32 `json:"itemid"`
}

type GameEventDotaItemGifted struct {
	Userid   int32 `json:"userid"`
	Itemid   int32 `json:"itemid"`
	Sourceid int32 `json:"sourceid"`
}

type GameEventDotaRunePickup struct {
	Userid int32 `json:"userid"`
	Type   int32 `json:"type"`
	Rune   int32 `json:"rune"`
}

type GameEventDotaRuneSpotted struct {
	Userid int32 `json:"userid"`
	Rune   int32 `json:"rune"`
}

type GameEventDotaItemSpotted struct {
	Userid int32 `json:"userid"`
	Itemid int32 `json:"itemid"`
}

type GameEventDotaNoBattlePoints struct {
	Userid int32 `json:"userid"`
	Reason int32 `json:"reason"`
}

type GameEventDotaChatInformational struct {
	Userid int32 `json:"userid"`
	Type   int32 `json:"type"`
}

type GameEventDotaActionItem struct {
	Reason  int32 `json:"reason"`
	Itemdef int32 `json:"itemdef"`
	Message int32 `json:"message"`
}

type GameEventDotaChatBanNotification struct {
	Userid int32 `json:"userid"`
}

type GameEventDotaChatEvent struct {
	Userid  int32 `json:"userid"`
	Gold    int32 `json:"gold"`
	Message int32 `json:"message"`
}

type GameEventDotaChatTimedReward struct {
	Userid  int32 `json:"userid"`
	Itmedef int32 `json:"itmedef"`
	Message int32 `json:"message"`
}

type GameEventDotaPauseEvent struct {
	Userid  int32 `json:"userid"`
	Value   int32 `json:"value"`
	Message int32 `json:"message"`
}

type GameEventDotaChatKillStreak struct {
	Gold            int32 `json:"gold"`
	KillerId        int32 `json:"killer_id"`
	KillerStreak    int32 `json:"killer_streak"`
	KillerMultikill int32 `json:"killer_multikill"`
	VictimId        int32 `json:"victim_id"`
	VictimStreak    int32 `json:"victim_streak"`
}

type GameEventDotaChatFirstBlood struct {
	Gold     int32 `json:"gold"`
	KillerId int32 `json:"killer_id"`
	VictimId int32 `json:"victim_id"`
}

type GameEventDotaChatAssassinAnnounce struct {
	AssassinId int32 `json:"assassin_id"`
	TargetId   int32 `json:"target_id"`
	Message    int32 `json:"message"`
}

type GameEventDotaChatAssassinDenied struct {
	AssassinId int32 `json:"assassin_id"`
	TargetId   int32 `json:"target_id"`
	Message    int32 `json:"message"`
}

type GameEventDotaChatAssassinSuccess struct {
	AssassinId int32 `json:"assassin_id"`
	TargetId   int32 `json:"target_id"`
	Message    int32 `json:"message"`
}

type GameEventDotaPlayerUpdateHeroSelection struct {
	Tabcycle bool `json:"tabcycle"`
}

type GameEventDotaPlayerUpdateSelectedUnit struct {
}

type GameEventDotaPlayerUpdateQueryUnit struct {
}

type GameEventDotaPlayerUpdateKillcamUnit struct {
}

type GameEventDotaPlayerTakeTowerDamage struct {
	PlayerID int32 `json:"PlayerID"`
	Damage   int32 `json:"damage"`
}

type GameEventDotaHudErrorMessage struct {
	Reason  int32  `json:"reason"`
	Message string `json:"message"`
}

type GameEventDotaActionSuccess struct {
}

type GameEventDotaStartingPositionChanged struct {
}

type GameEventDotaMoneyChanged struct {
}

type GameEventDotaEnemyMoneyChanged struct {
}

type GameEventDotaPortraitUnitStatsChanged struct {
}

type GameEventDotaPortraitUnitModifiersChanged struct {
}

type GameEventDotaForcePortraitUpdate struct {
}

type GameEventDotaInventoryChanged struct {
}

type GameEventDotaItemPickedUp struct {
	Itemname        string `json:"itemname"`
	PlayerID        int32  `json:"PlayerID"`
	ItemEntityIndex int32  `json:"ItemEntityIndex"`
	HeroEntityIndex int32  `json:"HeroEntityIndex"`
}

type GameEventDotaInventoryItemChanged struct {
	EntityIndex int32 `json:"entityIndex"`
}

type GameEventDotaAbilityChanged struct {
}

type GameEventDotaPortraitAbilityLayoutChanged struct {
}

type GameEventDotaInventoryItemAdded struct {
	Itemname string `json:"itemname"`
}

type GameEventDotaInventoryChangedQueryUnit struct {
}

type GameEventDotaLinkClicked struct {
	Link    string `json:"link"`
	Nav     bool   `json:"nav"`
	NavBack bool   `json:"nav_back"`
	Recipe  int32  `json:"recipe"`
	Shop    int32  `json:"shop"`
}

type GameEventDotaSetQuickBuy struct {
	Item   string `json:"item"`
	Recipe int32  `json:"recipe"`
	Toggle bool   `json:"toggle"`
}

type GameEventDotaQuickBuyChanged struct {
	Item   string `json:"item"`
	Recipe int32  `json:"recipe"`
}

type GameEventDotaPlayerShopChanged struct {
	Prevshopmask int32 `json:"prevshopmask"`
	Shopmask     int32 `json:"shopmask"`
}

type GameEventDotaPlayerShowKillcam struct {
	Nodes  int32 `json:"nodes"`
	Player int32 `json:"player"`
}

type GameEventDotaPlayerShowMinikillcam struct {
	Nodes  int32 `json:"nodes"`
	Player int32 `json:"player"`
}

type GameEventGcUserSessionCreated struct {
}

type GameEventTeamDataUpdated struct {
}

type GameEventGuildDataUpdated struct {
}

type GameEventGuildOpenPartiesUpdated struct {
}

type GameEventFantasyUpdated struct {
}

type GameEventFantasyLeagueChanged struct {
}

type GameEventFantasyScoreInfoChanged struct {
}

type GameEventPlayerInfoUpdated struct {
}

type GameEventPlayerInfoIndividualUpdated struct {
	AccountId int32 `json:"account_id"`
}

type GameEventGameRulesStateChange struct {
}

type GameEventMatchHistoryUpdated struct {
	SteamID uint64 `json:"SteamID"`
}

type GameEventMatchDetailsUpdated struct {
	MatchID uint64 `json:"matchID"`
	Result  int32  `json:"result"`
}

type GameEventLiveGamesUpdated struct {
}

type GameEventRecentMatchesUpdated struct {
	Page int32 `json:"Page"`
}

type GameEventNewsUpdated struct {
}

type GameEventPersonaUpdated struct {
	SteamID uint64 `json:"SteamID"`
}

type GameEventTournamentStateUpdated struct {
}

type GameEventPartyUpdated struct {
}

type GameEventLobbyUpdated struct {
}

type GameEventDashboardCachesCleared struct {
}

type GameEventLastHit struct {
	PlayerID   int32 `json:"PlayerID"`
	EntKilled  int32 `json:"EntKilled"`
	FirstBlood bool  `json:"FirstBlood"`
	HeroKill   bool  `json:"HeroKill"`
	TowerKill  bool  `json:"TowerKill"`
}

type GameEventPlayerCompletedGame struct {
	PlayerID int32 `json:"PlayerID"`
	Winner   int32 `json:"Winner"`
}

type GameEventPlayerReconnected struct {
	PlayerID int32 `json:"PlayerID"`
}

type GameEventNommedTree struct {
	PlayerID int32 `json:"PlayerID"`
}

type GameEventDotaRuneActivatedServer struct {
	PlayerID int32 `json:"PlayerID"`
	Rune     int32 `json:"rune"`
}

type GameEventDotaPlayerGainedLevel struct {
	PlayerID int32 `json:"PlayerID"`
	Level    int32 `json:"level"`
}

type GameEventDotaPlayerLearnedAbility struct {
	PlayerID    int32  `json:"PlayerID"`
	Abilityname string `json:"abilityname"`
}

type GameEventDotaPlayerUsedAbility struct {
	PlayerID    int32  `json:"PlayerID"`
	Abilityname string `json:"abilityname"`
}

type GameEventDotaNonPlayerUsedAbility struct {
	Abilityname string `json:"abilityname"`
}

type GameEventDotaPlayerBeginCast struct {
	PlayerID    int32  `json:"PlayerID"`
	Abilityname string `json:"abilityname"`
}

type GameEventDotaNonPlayerBeginCast struct {
	Abilityname string `json:"abilityname"`
}

type GameEventDotaAbilityChannelFinished struct {
	Abilityname string `json:"abilityname"`
	Interrupted bool   `json:"interrupted"`
}

type GameEventDotaHoldoutReviveComplete struct {
	Caster int32 `json:"caster"`
	Target int32 `json:"target"`
}

type GameEventDotaPlayerKilled struct {
	PlayerID  int32 `json:"PlayerID"`
	HeroKill  bool  `json:"HeroKill"`
	TowerKill bool  `json:"TowerKill"`
}

type GameEventBindpanelOpen struct {
}

type GameEventBindpanelClose struct {
}

type GameEventKeybindChanged struct {
}

type GameEventDotaItemDragBegin struct {
}

type GameEventDotaItemDragEnd struct {
}

type GameEventDotaShopItemDragBegin struct {
}

type GameEventDotaShopItemDragEnd struct {
}

type GameEventDotaItemPurchased struct {
	PlayerID int32  `json:"PlayerID"`
	Itemname string `json:"itemname"`
	Itemcost int32  `json:"itemcost"`
}

type GameEventDotaItemCombined struct {
	PlayerID int32  `json:"PlayerID"`
	Itemname string `json:"itemname"`
	Itemcost int32  `json:"itemcost"`
}

type GameEventDotaItemUsed struct {
	PlayerID int32  `json:"PlayerID"`
	Itemname string `json:"itemname"`
}

type GameEventDotaItemAutoPurchase struct {
	ItemId int32 `json:"item_id"`
}

type GameEventDotaUnitEvent struct {
	Victim       int32 `json:"victim"`
	Attacker     int32 `json:"attacker"`
	Basepriority int32 `json:"basepriority"`
	Priority     int32 `json:"priority"`
	Eventtype    int32 `json:"eventtype"`
}

type GameEventDotaQuestStarted struct {
	QuestIndex int32 `json:"questIndex"`
}

type GameEventDotaQuestCompleted struct {
	QuestIndex int32 `json:"questIndex"`
}

type GameEventGameuiActivated struct {
}

type GameEventGameuiHidden struct {
}

type GameEventPlayerFullyjoined struct {
	Userid int32  `json:"userid"`
	Name   string `json:"name"`
}

type GameEventDotaSpectateHero struct {
	Entindex int32 `json:"entindex"`
}

type GameEventDotaMatchDone struct {
	Winningteam int32 `json:"winningteam"`
}

type GameEventDotaMatchDoneClient struct {
}

type GameEventSetInstructorGroupEnabled struct {
	Group   string `json:"group"`
	Enabled int32  `json:"enabled"`
}

type GameEventJoinedChatChannel struct {
	ChannelName string `json:"channelName"`
}

type GameEventLeftChatChannel struct {
	ChannelName string `json:"channelName"`
}

type GameEventGcChatChannelListUpdated struct {
}

type GameEventTodayMessagesUpdated struct {
	NumMessages int32 `json:"num_messages"`
}

type GameEventFileDownloaded struct {
	Success       bool   `json:"success"`
	LocalFilename string `json:"local_filename"`
	RemoteUrl     string `json:"remote_url"`
}

type GameEventPlayerReportCountsUpdated struct {
	PositiveRemaining int32 `json:"positive_remaining"`
	NegativeRemaining int32 `json:"negative_remaining"`
	PositiveTotal     int32 `json:"positive_total"`
	NegativeTotal     int32 `json:"negative_total"`
}

type GameEventScaleformFileDownloadComplete struct {
	Success       bool   `json:"success"`
	LocalFilename string `json:"local_filename"`
	RemoteUrl     string `json:"remote_url"`
}

type GameEventItemPurchased struct {
	Itemid int32 `json:"itemid"`
}

type GameEventGcMismatchedVersion struct {
}

type GameEventDemoStop struct {
}

type GameEventMapShutdown struct {
}

type GameEventDotaWorkshopFileselected struct {
	Filename string `json:"filename"`
}

type GameEventDotaWorkshopFilecanceled struct {
}

type GameEventRichPresenceUpdated struct {
}

type GameEventDotaHeroRandom struct {
	Userid int32 `json:"userid"`
	Heroid int32 `json:"heroid"`
}

type GameEventDotaRdChatTurn struct {
	Userid int32 `json:"userid"`
}

type GameEventDotaFavoriteHeroesUpdated struct {
}

type GameEventProfileOpened struct {
}

type GameEventProfileClosed struct {
}

type GameEventItemPreviewClosed struct {
}

type GameEventDashboardSwitchedSection struct {
	Section int32 `json:"section"`
}

type GameEventDotaTournamentItemEvent struct {
	WinnerCount int32 `json:"winner_count"`
	EventType   int32 `json:"event_type"`
}

type GameEventDotaHeroSwap struct {
	Playerid1 int32 `json:"playerid1"`
	Playerid2 int32 `json:"playerid2"`
}

type GameEventDotaResetSuggestedItems struct {
}

type GameEventHalloweenHighScoreReceived struct {
	Round int32 `json:"round"`
}

type GameEventHalloweenPhaseEnd struct {
	Phase int32 `json:"phase"`
	Team  int32 `json:"team"`
}

type GameEventHalloweenHighScoreRequestFailed struct {
	Round int32 `json:"round"`
}

type GameEventDotaHudSkinChanged struct {
	Skin  string `json:"skin"`
	Style int32  `json:"style"`
}

type GameEventDotaInventoryPlayerGotItem struct {
	Itemname string `json:"itemname"`
}

type GameEventPlayerIsExperienced struct {
}

type GameEventPlayerIsNotexperienced struct {
}

type GameEventDotaTutorialLessonStart struct {
}

type GameEventDotaTutorialTaskAdvance struct {
}

type GameEventDotaTutorialShopToggled struct {
	ShopOpened bool `json:"shop_opened"`
}

type GameEventMapLocationUpdated struct {
}

type GameEventRichpresenceCustomUpdated struct {
}

type GameEventGameEndVisible struct {
}

type GameEventAntiaddictionUpdate struct {
}

type GameEventHighlightHudElement struct {
	Elementname string  `json:"elementname"`
	Duration    float32 `json:"duration"`
}

type GameEventHideHighlightHudElement struct {
}

type GameEventIntroVideoFinished struct {
}

type GameEventMatchmakingStatusVisibilityChanged struct {
}

type GameEventPracticeLobbyVisibilityChanged struct {
}

type GameEventDotaCourierTransferItem struct {
}

type GameEventFullUiUnlocked struct {
}

type GameEventHeroSelectorPreviewSet struct {
	Setindex int32 `json:"setindex"`
}

type GameEventAntiaddictionToast struct {
	Message  string  `json:"message"`
	Duration float32 `json:"duration"`
}

type GameEventHeroPickerShown struct {
}

type GameEventHeroPickerHidden struct {
}

type GameEventDotaLocalQuickbuyChanged struct {
}

type GameEventShowCenterMessage struct {
	Message           string  `json:"message"`
	Duration          float32 `json:"duration"`
	ClearMessageQueue bool    `json:"clear_message_queue"`
}

type GameEventHudFlipChanged struct {
	Flipped bool `json:"flipped"`
}

type GameEventFrostyPointsUpdated struct {
}

type GameEventDefeated struct {
	Entindex int32 `json:"entindex"`
}

type GameEventResetDefeated struct {
}

type GameEventBoosterStateUpdated struct {
}

type GameEventEventPointsUpdated struct {
	EventId       int32 `json:"event_id"`
	Points        int32 `json:"points"`
	PremiumPoints int32 `json:"premium_points"`
	Owned         bool  `json:"owned"`
}

type GameEventLocalPlayerEventPoints struct {
	Points         int32 `json:"points"`
	ConversionRate int32 `json:"conversion_rate"`
}

type GameEventCustomGameDifficulty struct {
	Difficulty int32 `json:"difficulty"`
}

type GameEventTreeCut struct {
	TreeX float32 `json:"tree_x"`
	TreeY float32 `json:"tree_y"`
}

type GameEventUgcDetailsArrived struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcSubscribed struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcUnsubscribed struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcDownloadRequested struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcInstalled struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventPrizepoolReceived struct {
	Success   bool   `json:"success"`
	Prizepool uint64 `json:"prizepool"`
	Leagueid  uint64 `json:"leagueid"`
}

type GameEventMicrotransactionSuccess struct {
	Txnid uint64 `json:"txnid"`
}

type GameEventDotaRubickAbilitySteal struct {
	AbilityIndex int32 `json:"abilityIndex"`
	AbilityLevel int32 `json:"abilityLevel"`
}

type GameEventCompendiumEventActionsLoaded struct {
	AccountId      uint64 `json:"account_id"`
	LeagueId       uint64 `json:"league_id"`
	LocalTest      bool   `json:"local_test"`
	OriginalPoints uint64 `json:"original_points"`
}

type GameEventCompendiumSelectionsLoaded struct {
	AccountId uint64 `json:"account_id"`
	LeagueId  uint64 `json:"league_id"`
	LocalTest bool   `json:"local_test"`
}

type GameEventCompendiumSetSelectionFailed struct {
	AccountId uint64 `json:"account_id"`
	LeagueId  uint64 `json:"league_id"`
	LocalTest bool   `json:"local_test"`
}

type GameEventCompendiumTrophiesLoaded struct {
	AccountId uint64 `json:"account_id"`
	LeagueId  uint64 `json:"league_id"`
	LocalTest bool   `json:"local_test"`
}

type GameEventCommunityCachedNamesUpdated struct {
}

type GameEventSpecItemPickup struct {
	PlayerId int32  `json:"player_id"`
	ItemName string `json:"item_name"`
	Purchase bool   `json:"purchase"`
}

type GameEventSpecAegisReclaimTime struct {
	ReclaimTime float32 `json:"reclaim_time"`
}

type GameEventAccountTrophiesChanged struct {
	AccountId uint64 `json:"account_id"`
}

type GameEventAccountAllHeroChallengeChanged struct {
	AccountId uint64 `json:"account_id"`
}

type GameEventTeamShowcaseUiUpdate struct {
	Show            bool   `json:"show"`
	AccountId       uint64 `json:"account_id"`
	HeroEntindex    int32  `json:"hero_entindex"`
	DisplayUiOnLeft bool   `json:"display_ui_on_left"`
}

type GameEventIngameEventsChanged struct {
}

type GameEventDotaMatchSignout struct {
}

type GameEventDotaIllusionsCreated struct {
	OriginalEntindex int32 `json:"original_entindex"`
}

type GameEventDotaYearBeastKilled struct {
	KillerPlayerId int32  `json:"killer_player_id"`
	Message        int32  `json:"message"`
	BeastId        uint64 `json:"beast_id"`
}

type GameEventDotaHeroUndoselection struct {
	Playerid1 int32 `json:"playerid1"`
}

type GameEventDotaChallengeSocacheUpdated struct {
}

type GameEventPartyInvitesUpdated struct {
}

type GameEventLobbyInvitesUpdated struct {
}

type GameEventCustomGameModeListUpdated struct {
}

type GameEventCustomGameLobbyListUpdated struct {
}

type GameEventFriendLobbyListUpdated struct {
}

type GameEventDotaTeamPlayerListChanged struct {
}

type GameEventDotaPlayerDetailsChanged struct {
}

type GameEventPlayerProfileStatsUpdated struct {
	AccountId uint64 `json:"account_id"`
}

type GameEventCustomGamePlayerCountUpdated struct {
	CustomGameId uint64 `json:"custom_game_id"`
}

type GameEventCustomGameFriendsPlayedUpdated struct {
	CustomGameId uint64 `json:"custom_game_id"`
}

type GameEventCustomGamesFriendsPlayUpdated struct {
}

type GameEventDotaPlayerUpdateAssignedHero struct {
}

type GameEventDotaPlayerHeroSelectionDirty struct {
}

type GameEventDotaNpcGoalReached struct {
	NpcEntindex      int32 `json:"npc_entindex"`
	GoalEntindex     int32 `json:"goal_entindex"`
	NextGoalEntindex int32 `json:"next_goal_entindex"`
}

type GameEventHltvStatus struct {
	Clients int32  `json:"clients"`
	Slots   int32  `json:"slots"`
	Proxies int32  `json:"proxies"`
	Master  string `json:"master"`
}

type GameEventHltvCameraman struct {
	Index int32 `json:"index"`
}

type GameEventHltvRankCamera struct {
	Index  int32   `json:"index"`
	Rank   float32 `json:"rank"`
	Target int32   `json:"target"`
}

type GameEventHltvRankEntity struct {
	Index  int32   `json:"index"`
	Rank   float32 `json:"rank"`
	Target int32   `json:"target"`
}

type GameEventHltvFixed struct {
	Posx   int32   `json:"posx"`
	Posy   int32   `json:"posy"`
	Posz   int32   `json:"posz"`
	Theta  int32   `json:"theta"`
	Phi    int32   `json:"phi"`
	Offset int32   `json:"offset"`
	Fov    float32 `json:"fov"`
	Target int32   `json:"target"`
}

type GameEventHltvChase struct {
	Target1  int32 `json:"target1"`
	Target2  int32 `json:"target2"`
	Distance int32 `json:"distance"`
	Theta    int32 `json:"theta"`
	Phi      int32 `json:"phi"`
	Inertia  int32 `json:"inertia"`
	Ineye    int32 `json:"ineye"`
}

type GameEventHltvMessage struct {
	Text string `json:"text"`
}

type GameEventHltvTitle struct {
	Text string `json:"text"`
}

type GameEventHltvChat struct {
	Name    string `json:"name"`
	Text    string `json:"text"`
	SteamID uint64 `json:"steamID"`
}

type GameEventHltvVersioninfo struct {
	Version int32 `json:"version"`
}

type GameEventDotaChaseHero struct {
	Target1         int32   `json:"target1"`
	Target2         int32   `json:"target2"`
	Type            int32   `json:"type"`
	Priority        int32   `json:"priority"`
	Gametime        float32 `json:"gametime"`
	Highlight       bool    `json:"highlight"`
	Target1playerid int32   `json:"target1playerid"`
	Target2playerid int32   `json:"target2playerid"`
	Eventtype       int32   `json:"eventtype"`
}

type GameEventDotaCombatlog struct {
	Type             int32   `json:"type"`
	Sourcename       int32   `json:"sourcename"`
	Targetname       int32   `json:"targetname"`
	Attackername     int32   `json:"attackername"`
	Inflictorname    int32   `json:"inflictorname"`
	Attackerillusion bool    `json:"attackerillusion"`
	Targetillusion   bool    `json:"targetillusion"`
	Value            int32   `json:"value"`
	Health           int32   `json:"health"`
	Timestamp        float32 `json:"timestamp"`
	Targetsourcename int32   `json:"targetsourcename"`
	Timestampraw     float32 `json:"timestampraw"`
	Attackerhero     bool    `json:"attackerhero"`
	Targethero       bool    `json:"targethero"`
	AbilityToggleOn  bool    `json:"ability_toggle_on"`
	AbilityToggleOff bool    `json:"ability_toggle_off"`
	AbilityLevel     int32   `json:"ability_level"`
	GoldReason       int32   `json:"gold_reason"`
	XpReason         int32   `json:"xp_reason"`
}

type GameEventDotaGameStateChange struct {
	OldState int32 `json:"old_state"`
	NewState int32 `json:"new_state"`
}

type GameEventDotaPlayerPickHero struct {
	Player    int32  `json:"player"`
	Heroindex int32  `json:"heroindex"`
	Hero      string `json:"hero"`
}

type GameEventDotaTeamKillCredit struct {
	KillerUserid int32 `json:"killer_userid"`
	VictimUserid int32 `json:"victim_userid"`
	Teamnumber   int32 `json:"teamnumber"`
	Herokills    int32 `json:"herokills"`
}

type GameEvents struct {
	onServerSpawn                        []func(*GameEventServerSpawn) error
	onServerPreShutdown                  []func(*GameEventServerPreShutdown) error
	onServerShutdown                     []func(*GameEventServerShutdown) error
	onServerCvar                         []func(*GameEventServerCvar) error
	onServerMessage                      []func(*GameEventServerMessage) error
	onServerAddban                       []func(*GameEventServerAddban) error
	onServerRemoveban                    []func(*GameEventServerRemoveban) error
	onPlayerConnect                      []func(*GameEventPlayerConnect) error
	onPlayerInfo                         []func(*GameEventPlayerInfo) error
	onPlayerDisconnect                   []func(*GameEventPlayerDisconnect) error
	onPlayerActivate                     []func(*GameEventPlayerActivate) error
	onPlayerConnectFull                  []func(*GameEventPlayerConnectFull) error
	onPlayerSay                          []func(*GameEventPlayerSay) error
	onPlayerFullUpdate                   []func(*GameEventPlayerFullUpdate) error
	onTeamInfo                           []func(*GameEventTeamInfo) error
	onTeamScore                          []func(*GameEventTeamScore) error
	onTeamplayBroadcastAudio             []func(*GameEventTeamplayBroadcastAudio) error
	onPlayerTeam                         []func(*GameEventPlayerTeam) error
	onPlayerClass                        []func(*GameEventPlayerClass) error
	onPlayerDeath                        []func(*GameEventPlayerDeath) error
	onPlayerHurt                         []func(*GameEventPlayerHurt) error
	onPlayerChat                         []func(*GameEventPlayerChat) error
	onPlayerScore                        []func(*GameEventPlayerScore) error
	onPlayerSpawn                        []func(*GameEventPlayerSpawn) error
	onPlayerShoot                        []func(*GameEventPlayerShoot) error
	onPlayerUse                          []func(*GameEventPlayerUse) error
	onPlayerChangename                   []func(*GameEventPlayerChangename) error
	onPlayerHintmessage                  []func(*GameEventPlayerHintmessage) error
	onGameInit                           []func(*GameEventGameInit) error
	onGameNewmap                         []func(*GameEventGameNewmap) error
	onGameStart                          []func(*GameEventGameStart) error
	onGameEnd                            []func(*GameEventGameEnd) error
	onRoundStart                         []func(*GameEventRoundStart) error
	onRoundEnd                           []func(*GameEventRoundEnd) error
	onRoundStartPreEntity                []func(*GameEventRoundStartPreEntity) error
	onTeamplayRoundStart                 []func(*GameEventTeamplayRoundStart) error
	onHostnameChanged                    []func(*GameEventHostnameChanged) error
	onDifficultyChanged                  []func(*GameEventDifficultyChanged) error
	onFinaleStart                        []func(*GameEventFinaleStart) error
	onGameMessage                        []func(*GameEventGameMessage) error
	onBreakBreakable                     []func(*GameEventBreakBreakable) error
	onBreakProp                          []func(*GameEventBreakProp) error
	onNpcSpawned                         []func(*GameEventNpcSpawned) error
	onNpcReplaced                        []func(*GameEventNpcReplaced) error
	onEntityKilled                       []func(*GameEventEntityKilled) error
	onEntityHurt                         []func(*GameEventEntityHurt) error
	onBonusUpdated                       []func(*GameEventBonusUpdated) error
	onPlayerStatsUpdated                 []func(*GameEventPlayerStatsUpdated) error
	onAchievementEvent                   []func(*GameEventAchievementEvent) error
	onAchievementEarned                  []func(*GameEventAchievementEarned) error
	onAchievementWriteFailed             []func(*GameEventAchievementWriteFailed) error
	onPhysgunPickup                      []func(*GameEventPhysgunPickup) error
	onFlareIgniteNpc                     []func(*GameEventFlareIgniteNpc) error
	onHelicopterGrenadePuntMiss          []func(*GameEventHelicopterGrenadePuntMiss) error
	onUserDataDownloaded                 []func(*GameEventUserDataDownloaded) error
	onRagdollDissolved                   []func(*GameEventRagdollDissolved) error
	onGameinstructorDraw                 []func(*GameEventGameinstructorDraw) error
	onGameinstructorNodraw               []func(*GameEventGameinstructorNodraw) error
	onMapTransition                      []func(*GameEventMapTransition) error
	onInstructorServerHintCreate         []func(*GameEventInstructorServerHintCreate) error
	onInstructorServerHintStop           []func(*GameEventInstructorServerHintStop) error
	onChatNewMessage                     []func(*GameEventChatNewMessage) error
	onChatMembersChanged                 []func(*GameEventChatMembersChanged) error
	onInventoryUpdated                   []func(*GameEventInventoryUpdated) error
	onCartUpdated                        []func(*GameEventCartUpdated) error
	onStorePricesheetUpdated             []func(*GameEventStorePricesheetUpdated) error
	onGcConnected                        []func(*GameEventGcConnected) error
	onItemSchemaInitialized              []func(*GameEventItemSchemaInitialized) error
	onDropRateModified                   []func(*GameEventDropRateModified) error
	onEventTicketModified                []func(*GameEventEventTicketModified) error
	onModifierEvent                      []func(*GameEventModifierEvent) error
	onDotaPlayerKill                     []func(*GameEventDotaPlayerKill) error
	onDotaPlayerDeny                     []func(*GameEventDotaPlayerDeny) error
	onDotaBarracksKill                   []func(*GameEventDotaBarracksKill) error
	onDotaTowerKill                      []func(*GameEventDotaTowerKill) error
	onDotaEffigyKill                     []func(*GameEventDotaEffigyKill) error
	onDotaRoshanKill                     []func(*GameEventDotaRoshanKill) error
	onDotaCourierLost                    []func(*GameEventDotaCourierLost) error
	onDotaCourierRespawned               []func(*GameEventDotaCourierRespawned) error
	onDotaGlyphUsed                      []func(*GameEventDotaGlyphUsed) error
	onDotaSuperCreeps                    []func(*GameEventDotaSuperCreeps) error
	onDotaItemPurchase                   []func(*GameEventDotaItemPurchase) error
	onDotaItemGifted                     []func(*GameEventDotaItemGifted) error
	onDotaRunePickup                     []func(*GameEventDotaRunePickup) error
	onDotaRuneSpotted                    []func(*GameEventDotaRuneSpotted) error
	onDotaItemSpotted                    []func(*GameEventDotaItemSpotted) error
	onDotaNoBattlePoints                 []func(*GameEventDotaNoBattlePoints) error
	onDotaChatInformational              []func(*GameEventDotaChatInformational) error
	onDotaActionItem                     []func(*GameEventDotaActionItem) error
	onDotaChatBanNotification            []func(*GameEventDotaChatBanNotification) error
	onDotaChatEvent                      []func(*GameEventDotaChatEvent) error
	onDotaChatTimedReward                []func(*GameEventDotaChatTimedReward) error
	onDotaPauseEvent                     []func(*GameEventDotaPauseEvent) error
	onDotaChatKillStreak                 []func(*GameEventDotaChatKillStreak) error
	onDotaChatFirstBlood                 []func(*GameEventDotaChatFirstBlood) error
	onDotaChatAssassinAnnounce           []func(*GameEventDotaChatAssassinAnnounce) error
	onDotaChatAssassinDenied             []func(*GameEventDotaChatAssassinDenied) error
	onDotaChatAssassinSuccess            []func(*GameEventDotaChatAssassinSuccess) error
	onDotaPlayerUpdateHeroSelection      []func(*GameEventDotaPlayerUpdateHeroSelection) error
	onDotaPlayerUpdateSelectedUnit       []func(*GameEventDotaPlayerUpdateSelectedUnit) error
	onDotaPlayerUpdateQueryUnit          []func(*GameEventDotaPlayerUpdateQueryUnit) error
	onDotaPlayerUpdateKillcamUnit        []func(*GameEventDotaPlayerUpdateKillcamUnit) error
	onDotaPlayerTakeTowerDamage          []func(*GameEventDotaPlayerTakeTowerDamage) error
	onDotaHudErrorMessage                []func(*GameEventDotaHudErrorMessage) error
	onDotaActionSuccess                  []func(*GameEventDotaActionSuccess) error
	onDotaStartingPositionChanged        []func(*GameEventDotaStartingPositionChanged) error
	onDotaMoneyChanged                   []func(*GameEventDotaMoneyChanged) error
	onDotaEnemyMoneyChanged              []func(*GameEventDotaEnemyMoneyChanged) error
	onDotaPortraitUnitStatsChanged       []func(*GameEventDotaPortraitUnitStatsChanged) error
	onDotaPortraitUnitModifiersChanged   []func(*GameEventDotaPortraitUnitModifiersChanged) error
	onDotaForcePortraitUpdate            []func(*GameEventDotaForcePortraitUpdate) error
	onDotaInventoryChanged               []func(*GameEventDotaInventoryChanged) error
	onDotaItemPickedUp                   []func(*GameEventDotaItemPickedUp) error
	onDotaInventoryItemChanged           []func(*GameEventDotaInventoryItemChanged) error
	onDotaAbilityChanged                 []func(*GameEventDotaAbilityChanged) error
	onDotaPortraitAbilityLayoutChanged   []func(*GameEventDotaPortraitAbilityLayoutChanged) error
	onDotaInventoryItemAdded             []func(*GameEventDotaInventoryItemAdded) error
	onDotaInventoryChangedQueryUnit      []func(*GameEventDotaInventoryChangedQueryUnit) error
	onDotaLinkClicked                    []func(*GameEventDotaLinkClicked) error
	onDotaSetQuickBuy                    []func(*GameEventDotaSetQuickBuy) error
	onDotaQuickBuyChanged                []func(*GameEventDotaQuickBuyChanged) error
	onDotaPlayerShopChanged              []func(*GameEventDotaPlayerShopChanged) error
	onDotaPlayerShowKillcam              []func(*GameEventDotaPlayerShowKillcam) error
	onDotaPlayerShowMinikillcam          []func(*GameEventDotaPlayerShowMinikillcam) error
	onGcUserSessionCreated               []func(*GameEventGcUserSessionCreated) error
	onTeamDataUpdated                    []func(*GameEventTeamDataUpdated) error
	onGuildDataUpdated                   []func(*GameEventGuildDataUpdated) error
	onGuildOpenPartiesUpdated            []func(*GameEventGuildOpenPartiesUpdated) error
	onFantasyUpdated                     []func(*GameEventFantasyUpdated) error
	onFantasyLeagueChanged               []func(*GameEventFantasyLeagueChanged) error
	onFantasyScoreInfoChanged            []func(*GameEventFantasyScoreInfoChanged) error
	onPlayerInfoUpdated                  []func(*GameEventPlayerInfoUpdated) error
	onPlayerInfoIndividualUpdated        []func(*GameEventPlayerInfoIndividualUpdated) error
	onGameRulesStateChange               []func(*GameEventGameRulesStateChange) error
	onMatchHistoryUpdated                []func(*GameEventMatchHistoryUpdated) error
	onMatchDetailsUpdated                []func(*GameEventMatchDetailsUpdated) error
	onLiveGamesUpdated                   []func(*GameEventLiveGamesUpdated) error
	onRecentMatchesUpdated               []func(*GameEventRecentMatchesUpdated) error
	onNewsUpdated                        []func(*GameEventNewsUpdated) error
	onPersonaUpdated                     []func(*GameEventPersonaUpdated) error
	onTournamentStateUpdated             []func(*GameEventTournamentStateUpdated) error
	onPartyUpdated                       []func(*GameEventPartyUpdated) error
	onLobbyUpdated                       []func(*GameEventLobbyUpdated) error
	onDashboardCachesCleared             []func(*GameEventDashboardCachesCleared) error
	onLastHit                            []func(*GameEventLastHit) error
	onPlayerCompletedGame                []func(*GameEventPlayerCompletedGame) error
	onPlayerReconnected                  []func(*GameEventPlayerReconnected) error
	onNommedTree                         []func(*GameEventNommedTree) error
	onDotaRuneActivatedServer            []func(*GameEventDotaRuneActivatedServer) error
	onDotaPlayerGainedLevel              []func(*GameEventDotaPlayerGainedLevel) error
	onDotaPlayerLearnedAbility           []func(*GameEventDotaPlayerLearnedAbility) error
	onDotaPlayerUsedAbility              []func(*GameEventDotaPlayerUsedAbility) error
	onDotaNonPlayerUsedAbility           []func(*GameEventDotaNonPlayerUsedAbility) error
	onDotaPlayerBeginCast                []func(*GameEventDotaPlayerBeginCast) error
	onDotaNonPlayerBeginCast             []func(*GameEventDotaNonPlayerBeginCast) error
	onDotaAbilityChannelFinished         []func(*GameEventDotaAbilityChannelFinished) error
	onDotaHoldoutReviveComplete          []func(*GameEventDotaHoldoutReviveComplete) error
	onDotaPlayerKilled                   []func(*GameEventDotaPlayerKilled) error
	onBindpanelOpen                      []func(*GameEventBindpanelOpen) error
	onBindpanelClose                     []func(*GameEventBindpanelClose) error
	onKeybindChanged                     []func(*GameEventKeybindChanged) error
	onDotaItemDragBegin                  []func(*GameEventDotaItemDragBegin) error
	onDotaItemDragEnd                    []func(*GameEventDotaItemDragEnd) error
	onDotaShopItemDragBegin              []func(*GameEventDotaShopItemDragBegin) error
	onDotaShopItemDragEnd                []func(*GameEventDotaShopItemDragEnd) error
	onDotaItemPurchased                  []func(*GameEventDotaItemPurchased) error
	onDotaItemCombined                   []func(*GameEventDotaItemCombined) error
	onDotaItemUsed                       []func(*GameEventDotaItemUsed) error
	onDotaItemAutoPurchase               []func(*GameEventDotaItemAutoPurchase) error
	onDotaUnitEvent                      []func(*GameEventDotaUnitEvent) error
	onDotaQuestStarted                   []func(*GameEventDotaQuestStarted) error
	onDotaQuestCompleted                 []func(*GameEventDotaQuestCompleted) error
	onGameuiActivated                    []func(*GameEventGameuiActivated) error
	onGameuiHidden                       []func(*GameEventGameuiHidden) error
	onPlayerFullyjoined                  []func(*GameEventPlayerFullyjoined) error
	onDotaSpectateHero                   []func(*GameEventDotaSpectateHero) error
	onDotaMatchDone                      []func(*GameEventDotaMatchDone) error
	onDotaMatchDoneClient                []func(*GameEventDotaMatchDoneClient) error
	onSetInstructorGroupEnabled          []func(*GameEventSetInstructorGroupEnabled) error
	onJoinedChatChannel                  []func(*GameEventJoinedChatChannel) error
	onLeftChatChannel                    []func(*GameEventLeftChatChannel) error
	onGcChatChannelListUpdated           []func(*GameEventGcChatChannelListUpdated) error
	onTodayMessagesUpdated               []func(*GameEventTodayMessagesUpdated) error
	onFileDownloaded                     []func(*GameEventFileDownloaded) error
	onPlayerReportCountsUpdated          []func(*GameEventPlayerReportCountsUpdated) error
	onScaleformFileDownloadComplete      []func(*GameEventScaleformFileDownloadComplete) error
	onItemPurchased                      []func(*GameEventItemPurchased) error
	onGcMismatchedVersion                []func(*GameEventGcMismatchedVersion) error
	onDemoStop                           []func(*GameEventDemoStop) error
	onMapShutdown                        []func(*GameEventMapShutdown) error
	onDotaWorkshopFileselected           []func(*GameEventDotaWorkshopFileselected) error
	onDotaWorkshopFilecanceled           []func(*GameEventDotaWorkshopFilecanceled) error
	onRichPresenceUpdated                []func(*GameEventRichPresenceUpdated) error
	onDotaHeroRandom                     []func(*GameEventDotaHeroRandom) error
	onDotaRdChatTurn                     []func(*GameEventDotaRdChatTurn) error
	onDotaFavoriteHeroesUpdated          []func(*GameEventDotaFavoriteHeroesUpdated) error
	onProfileOpened                      []func(*GameEventProfileOpened) error
	onProfileClosed                      []func(*GameEventProfileClosed) error
	onItemPreviewClosed                  []func(*GameEventItemPreviewClosed) error
	onDashboardSwitchedSection           []func(*GameEventDashboardSwitchedSection) error
	onDotaTournamentItemEvent            []func(*GameEventDotaTournamentItemEvent) error
	onDotaHeroSwap                       []func(*GameEventDotaHeroSwap) error
	onDotaResetSuggestedItems            []func(*GameEventDotaResetSuggestedItems) error
	onHalloweenHighScoreReceived         []func(*GameEventHalloweenHighScoreReceived) error
	onHalloweenPhaseEnd                  []func(*GameEventHalloweenPhaseEnd) error
	onHalloweenHighScoreRequestFailed    []func(*GameEventHalloweenHighScoreRequestFailed) error
	onDotaHudSkinChanged                 []func(*GameEventDotaHudSkinChanged) error
	onDotaInventoryPlayerGotItem         []func(*GameEventDotaInventoryPlayerGotItem) error
	onPlayerIsExperienced                []func(*GameEventPlayerIsExperienced) error
	onPlayerIsNotexperienced             []func(*GameEventPlayerIsNotexperienced) error
	onDotaTutorialLessonStart            []func(*GameEventDotaTutorialLessonStart) error
	onDotaTutorialTaskAdvance            []func(*GameEventDotaTutorialTaskAdvance) error
	onDotaTutorialShopToggled            []func(*GameEventDotaTutorialShopToggled) error
	onMapLocationUpdated                 []func(*GameEventMapLocationUpdated) error
	onRichpresenceCustomUpdated          []func(*GameEventRichpresenceCustomUpdated) error
	onGameEndVisible                     []func(*GameEventGameEndVisible) error
	onAntiaddictionUpdate                []func(*GameEventAntiaddictionUpdate) error
	onHighlightHudElement                []func(*GameEventHighlightHudElement) error
	onHideHighlightHudElement            []func(*GameEventHideHighlightHudElement) error
	onIntroVideoFinished                 []func(*GameEventIntroVideoFinished) error
	onMatchmakingStatusVisibilityChanged []func(*GameEventMatchmakingStatusVisibilityChanged) error
	onPracticeLobbyVisibilityChanged     []func(*GameEventPracticeLobbyVisibilityChanged) error
	onDotaCourierTransferItem            []func(*GameEventDotaCourierTransferItem) error
	onFullUiUnlocked                     []func(*GameEventFullUiUnlocked) error
	onHeroSelectorPreviewSet             []func(*GameEventHeroSelectorPreviewSet) error
	onAntiaddictionToast                 []func(*GameEventAntiaddictionToast) error
	onHeroPickerShown                    []func(*GameEventHeroPickerShown) error
	onHeroPickerHidden                   []func(*GameEventHeroPickerHidden) error
	onDotaLocalQuickbuyChanged           []func(*GameEventDotaLocalQuickbuyChanged) error
	onShowCenterMessage                  []func(*GameEventShowCenterMessage) error
	onHudFlipChanged                     []func(*GameEventHudFlipChanged) error
	onFrostyPointsUpdated                []func(*GameEventFrostyPointsUpdated) error
	onDefeated                           []func(*GameEventDefeated) error
	onResetDefeated                      []func(*GameEventResetDefeated) error
	onBoosterStateUpdated                []func(*GameEventBoosterStateUpdated) error
	onEventPointsUpdated                 []func(*GameEventEventPointsUpdated) error
	onLocalPlayerEventPoints             []func(*GameEventLocalPlayerEventPoints) error
	onCustomGameDifficulty               []func(*GameEventCustomGameDifficulty) error
	onTreeCut                            []func(*GameEventTreeCut) error
	onUgcDetailsArrived                  []func(*GameEventUgcDetailsArrived) error
	onUgcSubscribed                      []func(*GameEventUgcSubscribed) error
	onUgcUnsubscribed                    []func(*GameEventUgcUnsubscribed) error
	onUgcDownloadRequested               []func(*GameEventUgcDownloadRequested) error
	onUgcInstalled                       []func(*GameEventUgcInstalled) error
	onPrizepoolReceived                  []func(*GameEventPrizepoolReceived) error
	onMicrotransactionSuccess            []func(*GameEventMicrotransactionSuccess) error
	onDotaRubickAbilitySteal             []func(*GameEventDotaRubickAbilitySteal) error
	onCompendiumEventActionsLoaded       []func(*GameEventCompendiumEventActionsLoaded) error
	onCompendiumSelectionsLoaded         []func(*GameEventCompendiumSelectionsLoaded) error
	onCompendiumSetSelectionFailed       []func(*GameEventCompendiumSetSelectionFailed) error
	onCompendiumTrophiesLoaded           []func(*GameEventCompendiumTrophiesLoaded) error
	onCommunityCachedNamesUpdated        []func(*GameEventCommunityCachedNamesUpdated) error
	onSpecItemPickup                     []func(*GameEventSpecItemPickup) error
	onSpecAegisReclaimTime               []func(*GameEventSpecAegisReclaimTime) error
	onAccountTrophiesChanged             []func(*GameEventAccountTrophiesChanged) error
	onAccountAllHeroChallengeChanged     []func(*GameEventAccountAllHeroChallengeChanged) error
	onTeamShowcaseUiUpdate               []func(*GameEventTeamShowcaseUiUpdate) error
	onIngameEventsChanged                []func(*GameEventIngameEventsChanged) error
	onDotaMatchSignout                   []func(*GameEventDotaMatchSignout) error
	onDotaIllusionsCreated               []func(*GameEventDotaIllusionsCreated) error
	onDotaYearBeastKilled                []func(*GameEventDotaYearBeastKilled) error
	onDotaHeroUndoselection              []func(*GameEventDotaHeroUndoselection) error
	onDotaChallengeSocacheUpdated        []func(*GameEventDotaChallengeSocacheUpdated) error
	onPartyInvitesUpdated                []func(*GameEventPartyInvitesUpdated) error
	onLobbyInvitesUpdated                []func(*GameEventLobbyInvitesUpdated) error
	onCustomGameModeListUpdated          []func(*GameEventCustomGameModeListUpdated) error
	onCustomGameLobbyListUpdated         []func(*GameEventCustomGameLobbyListUpdated) error
	onFriendLobbyListUpdated             []func(*GameEventFriendLobbyListUpdated) error
	onDotaTeamPlayerListChanged          []func(*GameEventDotaTeamPlayerListChanged) error
	onDotaPlayerDetailsChanged           []func(*GameEventDotaPlayerDetailsChanged) error
	onPlayerProfileStatsUpdated          []func(*GameEventPlayerProfileStatsUpdated) error
	onCustomGamePlayerCountUpdated       []func(*GameEventCustomGamePlayerCountUpdated) error
	onCustomGameFriendsPlayedUpdated     []func(*GameEventCustomGameFriendsPlayedUpdated) error
	onCustomGamesFriendsPlayUpdated      []func(*GameEventCustomGamesFriendsPlayUpdated) error
	onDotaPlayerUpdateAssignedHero       []func(*GameEventDotaPlayerUpdateAssignedHero) error
	onDotaPlayerHeroSelectionDirty       []func(*GameEventDotaPlayerHeroSelectionDirty) error
	onDotaNpcGoalReached                 []func(*GameEventDotaNpcGoalReached) error
	onHltvStatus                         []func(*GameEventHltvStatus) error
	onHltvCameraman                      []func(*GameEventHltvCameraman) error
	onHltvRankCamera                     []func(*GameEventHltvRankCamera) error
	onHltvRankEntity                     []func(*GameEventHltvRankEntity) error
	onHltvFixed                          []func(*GameEventHltvFixed) error
	onHltvChase                          []func(*GameEventHltvChase) error
	onHltvMessage                        []func(*GameEventHltvMessage) error
	onHltvTitle                          []func(*GameEventHltvTitle) error
	onHltvChat                           []func(*GameEventHltvChat) error
	onHltvVersioninfo                    []func(*GameEventHltvVersioninfo) error
	onDotaChaseHero                      []func(*GameEventDotaChaseHero) error
	onDotaCombatlog                      []func(*GameEventDotaCombatlog) error
	onDotaGameStateChange                []func(*GameEventDotaGameStateChange) error
	onDotaPlayerPickHero                 []func(*GameEventDotaPlayerPickHero) error
	onDotaTeamKillCredit                 []func(*GameEventDotaTeamKillCredit) error
}

func (ge *GameEvents) OnServerSpawn(fn func(*GameEventServerSpawn) error) {
	if ge.onServerSpawn == nil {
		ge.onServerSpawn = make([]func(*GameEventServerSpawn) error, 0)
	}
	ge.onServerSpawn = append(ge.onServerSpawn, fn)
}

func (ge *GameEvents) OnServerPreShutdown(fn func(*GameEventServerPreShutdown) error) {
	if ge.onServerPreShutdown == nil {
		ge.onServerPreShutdown = make([]func(*GameEventServerPreShutdown) error, 0)
	}
	ge.onServerPreShutdown = append(ge.onServerPreShutdown, fn)
}

func (ge *GameEvents) OnServerShutdown(fn func(*GameEventServerShutdown) error) {
	if ge.onServerShutdown == nil {
		ge.onServerShutdown = make([]func(*GameEventServerShutdown) error, 0)
	}
	ge.onServerShutdown = append(ge.onServerShutdown, fn)
}

func (ge *GameEvents) OnServerCvar(fn func(*GameEventServerCvar) error) {
	if ge.onServerCvar == nil {
		ge.onServerCvar = make([]func(*GameEventServerCvar) error, 0)
	}
	ge.onServerCvar = append(ge.onServerCvar, fn)
}

func (ge *GameEvents) OnServerMessage(fn func(*GameEventServerMessage) error) {
	if ge.onServerMessage == nil {
		ge.onServerMessage = make([]func(*GameEventServerMessage) error, 0)
	}
	ge.onServerMessage = append(ge.onServerMessage, fn)
}

func (ge *GameEvents) OnServerAddban(fn func(*GameEventServerAddban) error) {
	if ge.onServerAddban == nil {
		ge.onServerAddban = make([]func(*GameEventServerAddban) error, 0)
	}
	ge.onServerAddban = append(ge.onServerAddban, fn)
}

func (ge *GameEvents) OnServerRemoveban(fn func(*GameEventServerRemoveban) error) {
	if ge.onServerRemoveban == nil {
		ge.onServerRemoveban = make([]func(*GameEventServerRemoveban) error, 0)
	}
	ge.onServerRemoveban = append(ge.onServerRemoveban, fn)
}

func (ge *GameEvents) OnPlayerConnect(fn func(*GameEventPlayerConnect) error) {
	if ge.onPlayerConnect == nil {
		ge.onPlayerConnect = make([]func(*GameEventPlayerConnect) error, 0)
	}
	ge.onPlayerConnect = append(ge.onPlayerConnect, fn)
}

func (ge *GameEvents) OnPlayerInfo(fn func(*GameEventPlayerInfo) error) {
	if ge.onPlayerInfo == nil {
		ge.onPlayerInfo = make([]func(*GameEventPlayerInfo) error, 0)
	}
	ge.onPlayerInfo = append(ge.onPlayerInfo, fn)
}

func (ge *GameEvents) OnPlayerDisconnect(fn func(*GameEventPlayerDisconnect) error) {
	if ge.onPlayerDisconnect == nil {
		ge.onPlayerDisconnect = make([]func(*GameEventPlayerDisconnect) error, 0)
	}
	ge.onPlayerDisconnect = append(ge.onPlayerDisconnect, fn)
}

func (ge *GameEvents) OnPlayerActivate(fn func(*GameEventPlayerActivate) error) {
	if ge.onPlayerActivate == nil {
		ge.onPlayerActivate = make([]func(*GameEventPlayerActivate) error, 0)
	}
	ge.onPlayerActivate = append(ge.onPlayerActivate, fn)
}

func (ge *GameEvents) OnPlayerConnectFull(fn func(*GameEventPlayerConnectFull) error) {
	if ge.onPlayerConnectFull == nil {
		ge.onPlayerConnectFull = make([]func(*GameEventPlayerConnectFull) error, 0)
	}
	ge.onPlayerConnectFull = append(ge.onPlayerConnectFull, fn)
}

func (ge *GameEvents) OnPlayerSay(fn func(*GameEventPlayerSay) error) {
	if ge.onPlayerSay == nil {
		ge.onPlayerSay = make([]func(*GameEventPlayerSay) error, 0)
	}
	ge.onPlayerSay = append(ge.onPlayerSay, fn)
}

func (ge *GameEvents) OnPlayerFullUpdate(fn func(*GameEventPlayerFullUpdate) error) {
	if ge.onPlayerFullUpdate == nil {
		ge.onPlayerFullUpdate = make([]func(*GameEventPlayerFullUpdate) error, 0)
	}
	ge.onPlayerFullUpdate = append(ge.onPlayerFullUpdate, fn)
}

func (ge *GameEvents) OnTeamInfo(fn func(*GameEventTeamInfo) error) {
	if ge.onTeamInfo == nil {
		ge.onTeamInfo = make([]func(*GameEventTeamInfo) error, 0)
	}
	ge.onTeamInfo = append(ge.onTeamInfo, fn)
}

func (ge *GameEvents) OnTeamScore(fn func(*GameEventTeamScore) error) {
	if ge.onTeamScore == nil {
		ge.onTeamScore = make([]func(*GameEventTeamScore) error, 0)
	}
	ge.onTeamScore = append(ge.onTeamScore, fn)
}

func (ge *GameEvents) OnTeamplayBroadcastAudio(fn func(*GameEventTeamplayBroadcastAudio) error) {
	if ge.onTeamplayBroadcastAudio == nil {
		ge.onTeamplayBroadcastAudio = make([]func(*GameEventTeamplayBroadcastAudio) error, 0)
	}
	ge.onTeamplayBroadcastAudio = append(ge.onTeamplayBroadcastAudio, fn)
}

func (ge *GameEvents) OnPlayerTeam(fn func(*GameEventPlayerTeam) error) {
	if ge.onPlayerTeam == nil {
		ge.onPlayerTeam = make([]func(*GameEventPlayerTeam) error, 0)
	}
	ge.onPlayerTeam = append(ge.onPlayerTeam, fn)
}

func (ge *GameEvents) OnPlayerClass(fn func(*GameEventPlayerClass) error) {
	if ge.onPlayerClass == nil {
		ge.onPlayerClass = make([]func(*GameEventPlayerClass) error, 0)
	}
	ge.onPlayerClass = append(ge.onPlayerClass, fn)
}

func (ge *GameEvents) OnPlayerDeath(fn func(*GameEventPlayerDeath) error) {
	if ge.onPlayerDeath == nil {
		ge.onPlayerDeath = make([]func(*GameEventPlayerDeath) error, 0)
	}
	ge.onPlayerDeath = append(ge.onPlayerDeath, fn)
}

func (ge *GameEvents) OnPlayerHurt(fn func(*GameEventPlayerHurt) error) {
	if ge.onPlayerHurt == nil {
		ge.onPlayerHurt = make([]func(*GameEventPlayerHurt) error, 0)
	}
	ge.onPlayerHurt = append(ge.onPlayerHurt, fn)
}

func (ge *GameEvents) OnPlayerChat(fn func(*GameEventPlayerChat) error) {
	if ge.onPlayerChat == nil {
		ge.onPlayerChat = make([]func(*GameEventPlayerChat) error, 0)
	}
	ge.onPlayerChat = append(ge.onPlayerChat, fn)
}

func (ge *GameEvents) OnPlayerScore(fn func(*GameEventPlayerScore) error) {
	if ge.onPlayerScore == nil {
		ge.onPlayerScore = make([]func(*GameEventPlayerScore) error, 0)
	}
	ge.onPlayerScore = append(ge.onPlayerScore, fn)
}

func (ge *GameEvents) OnPlayerSpawn(fn func(*GameEventPlayerSpawn) error) {
	if ge.onPlayerSpawn == nil {
		ge.onPlayerSpawn = make([]func(*GameEventPlayerSpawn) error, 0)
	}
	ge.onPlayerSpawn = append(ge.onPlayerSpawn, fn)
}

func (ge *GameEvents) OnPlayerShoot(fn func(*GameEventPlayerShoot) error) {
	if ge.onPlayerShoot == nil {
		ge.onPlayerShoot = make([]func(*GameEventPlayerShoot) error, 0)
	}
	ge.onPlayerShoot = append(ge.onPlayerShoot, fn)
}

func (ge *GameEvents) OnPlayerUse(fn func(*GameEventPlayerUse) error) {
	if ge.onPlayerUse == nil {
		ge.onPlayerUse = make([]func(*GameEventPlayerUse) error, 0)
	}
	ge.onPlayerUse = append(ge.onPlayerUse, fn)
}

func (ge *GameEvents) OnPlayerChangename(fn func(*GameEventPlayerChangename) error) {
	if ge.onPlayerChangename == nil {
		ge.onPlayerChangename = make([]func(*GameEventPlayerChangename) error, 0)
	}
	ge.onPlayerChangename = append(ge.onPlayerChangename, fn)
}

func (ge *GameEvents) OnPlayerHintmessage(fn func(*GameEventPlayerHintmessage) error) {
	if ge.onPlayerHintmessage == nil {
		ge.onPlayerHintmessage = make([]func(*GameEventPlayerHintmessage) error, 0)
	}
	ge.onPlayerHintmessage = append(ge.onPlayerHintmessage, fn)
}

func (ge *GameEvents) OnGameInit(fn func(*GameEventGameInit) error) {
	if ge.onGameInit == nil {
		ge.onGameInit = make([]func(*GameEventGameInit) error, 0)
	}
	ge.onGameInit = append(ge.onGameInit, fn)
}

func (ge *GameEvents) OnGameNewmap(fn func(*GameEventGameNewmap) error) {
	if ge.onGameNewmap == nil {
		ge.onGameNewmap = make([]func(*GameEventGameNewmap) error, 0)
	}
	ge.onGameNewmap = append(ge.onGameNewmap, fn)
}

func (ge *GameEvents) OnGameStart(fn func(*GameEventGameStart) error) {
	if ge.onGameStart == nil {
		ge.onGameStart = make([]func(*GameEventGameStart) error, 0)
	}
	ge.onGameStart = append(ge.onGameStart, fn)
}

func (ge *GameEvents) OnGameEnd(fn func(*GameEventGameEnd) error) {
	if ge.onGameEnd == nil {
		ge.onGameEnd = make([]func(*GameEventGameEnd) error, 0)
	}
	ge.onGameEnd = append(ge.onGameEnd, fn)
}

func (ge *GameEvents) OnRoundStart(fn func(*GameEventRoundStart) error) {
	if ge.onRoundStart == nil {
		ge.onRoundStart = make([]func(*GameEventRoundStart) error, 0)
	}
	ge.onRoundStart = append(ge.onRoundStart, fn)
}

func (ge *GameEvents) OnRoundEnd(fn func(*GameEventRoundEnd) error) {
	if ge.onRoundEnd == nil {
		ge.onRoundEnd = make([]func(*GameEventRoundEnd) error, 0)
	}
	ge.onRoundEnd = append(ge.onRoundEnd, fn)
}

func (ge *GameEvents) OnRoundStartPreEntity(fn func(*GameEventRoundStartPreEntity) error) {
	if ge.onRoundStartPreEntity == nil {
		ge.onRoundStartPreEntity = make([]func(*GameEventRoundStartPreEntity) error, 0)
	}
	ge.onRoundStartPreEntity = append(ge.onRoundStartPreEntity, fn)
}

func (ge *GameEvents) OnTeamplayRoundStart(fn func(*GameEventTeamplayRoundStart) error) {
	if ge.onTeamplayRoundStart == nil {
		ge.onTeamplayRoundStart = make([]func(*GameEventTeamplayRoundStart) error, 0)
	}
	ge.onTeamplayRoundStart = append(ge.onTeamplayRoundStart, fn)
}

func (ge *GameEvents) OnHostnameChanged(fn func(*GameEventHostnameChanged) error) {
	if ge.onHostnameChanged == nil {
		ge.onHostnameChanged = make([]func(*GameEventHostnameChanged) error, 0)
	}
	ge.onHostnameChanged = append(ge.onHostnameChanged, fn)
}

func (ge *GameEvents) OnDifficultyChanged(fn func(*GameEventDifficultyChanged) error) {
	if ge.onDifficultyChanged == nil {
		ge.onDifficultyChanged = make([]func(*GameEventDifficultyChanged) error, 0)
	}
	ge.onDifficultyChanged = append(ge.onDifficultyChanged, fn)
}

func (ge *GameEvents) OnFinaleStart(fn func(*GameEventFinaleStart) error) {
	if ge.onFinaleStart == nil {
		ge.onFinaleStart = make([]func(*GameEventFinaleStart) error, 0)
	}
	ge.onFinaleStart = append(ge.onFinaleStart, fn)
}

func (ge *GameEvents) OnGameMessage(fn func(*GameEventGameMessage) error) {
	if ge.onGameMessage == nil {
		ge.onGameMessage = make([]func(*GameEventGameMessage) error, 0)
	}
	ge.onGameMessage = append(ge.onGameMessage, fn)
}

func (ge *GameEvents) OnBreakBreakable(fn func(*GameEventBreakBreakable) error) {
	if ge.onBreakBreakable == nil {
		ge.onBreakBreakable = make([]func(*GameEventBreakBreakable) error, 0)
	}
	ge.onBreakBreakable = append(ge.onBreakBreakable, fn)
}

func (ge *GameEvents) OnBreakProp(fn func(*GameEventBreakProp) error) {
	if ge.onBreakProp == nil {
		ge.onBreakProp = make([]func(*GameEventBreakProp) error, 0)
	}
	ge.onBreakProp = append(ge.onBreakProp, fn)
}

func (ge *GameEvents) OnNpcSpawned(fn func(*GameEventNpcSpawned) error) {
	if ge.onNpcSpawned == nil {
		ge.onNpcSpawned = make([]func(*GameEventNpcSpawned) error, 0)
	}
	ge.onNpcSpawned = append(ge.onNpcSpawned, fn)
}

func (ge *GameEvents) OnNpcReplaced(fn func(*GameEventNpcReplaced) error) {
	if ge.onNpcReplaced == nil {
		ge.onNpcReplaced = make([]func(*GameEventNpcReplaced) error, 0)
	}
	ge.onNpcReplaced = append(ge.onNpcReplaced, fn)
}

func (ge *GameEvents) OnEntityKilled(fn func(*GameEventEntityKilled) error) {
	if ge.onEntityKilled == nil {
		ge.onEntityKilled = make([]func(*GameEventEntityKilled) error, 0)
	}
	ge.onEntityKilled = append(ge.onEntityKilled, fn)
}

func (ge *GameEvents) OnEntityHurt(fn func(*GameEventEntityHurt) error) {
	if ge.onEntityHurt == nil {
		ge.onEntityHurt = make([]func(*GameEventEntityHurt) error, 0)
	}
	ge.onEntityHurt = append(ge.onEntityHurt, fn)
}

func (ge *GameEvents) OnBonusUpdated(fn func(*GameEventBonusUpdated) error) {
	if ge.onBonusUpdated == nil {
		ge.onBonusUpdated = make([]func(*GameEventBonusUpdated) error, 0)
	}
	ge.onBonusUpdated = append(ge.onBonusUpdated, fn)
}

func (ge *GameEvents) OnPlayerStatsUpdated(fn func(*GameEventPlayerStatsUpdated) error) {
	if ge.onPlayerStatsUpdated == nil {
		ge.onPlayerStatsUpdated = make([]func(*GameEventPlayerStatsUpdated) error, 0)
	}
	ge.onPlayerStatsUpdated = append(ge.onPlayerStatsUpdated, fn)
}

func (ge *GameEvents) OnAchievementEvent(fn func(*GameEventAchievementEvent) error) {
	if ge.onAchievementEvent == nil {
		ge.onAchievementEvent = make([]func(*GameEventAchievementEvent) error, 0)
	}
	ge.onAchievementEvent = append(ge.onAchievementEvent, fn)
}

func (ge *GameEvents) OnAchievementEarned(fn func(*GameEventAchievementEarned) error) {
	if ge.onAchievementEarned == nil {
		ge.onAchievementEarned = make([]func(*GameEventAchievementEarned) error, 0)
	}
	ge.onAchievementEarned = append(ge.onAchievementEarned, fn)
}

func (ge *GameEvents) OnAchievementWriteFailed(fn func(*GameEventAchievementWriteFailed) error) {
	if ge.onAchievementWriteFailed == nil {
		ge.onAchievementWriteFailed = make([]func(*GameEventAchievementWriteFailed) error, 0)
	}
	ge.onAchievementWriteFailed = append(ge.onAchievementWriteFailed, fn)
}

func (ge *GameEvents) OnPhysgunPickup(fn func(*GameEventPhysgunPickup) error) {
	if ge.onPhysgunPickup == nil {
		ge.onPhysgunPickup = make([]func(*GameEventPhysgunPickup) error, 0)
	}
	ge.onPhysgunPickup = append(ge.onPhysgunPickup, fn)
}

func (ge *GameEvents) OnFlareIgniteNpc(fn func(*GameEventFlareIgniteNpc) error) {
	if ge.onFlareIgniteNpc == nil {
		ge.onFlareIgniteNpc = make([]func(*GameEventFlareIgniteNpc) error, 0)
	}
	ge.onFlareIgniteNpc = append(ge.onFlareIgniteNpc, fn)
}

func (ge *GameEvents) OnHelicopterGrenadePuntMiss(fn func(*GameEventHelicopterGrenadePuntMiss) error) {
	if ge.onHelicopterGrenadePuntMiss == nil {
		ge.onHelicopterGrenadePuntMiss = make([]func(*GameEventHelicopterGrenadePuntMiss) error, 0)
	}
	ge.onHelicopterGrenadePuntMiss = append(ge.onHelicopterGrenadePuntMiss, fn)
}

func (ge *GameEvents) OnUserDataDownloaded(fn func(*GameEventUserDataDownloaded) error) {
	if ge.onUserDataDownloaded == nil {
		ge.onUserDataDownloaded = make([]func(*GameEventUserDataDownloaded) error, 0)
	}
	ge.onUserDataDownloaded = append(ge.onUserDataDownloaded, fn)
}

func (ge *GameEvents) OnRagdollDissolved(fn func(*GameEventRagdollDissolved) error) {
	if ge.onRagdollDissolved == nil {
		ge.onRagdollDissolved = make([]func(*GameEventRagdollDissolved) error, 0)
	}
	ge.onRagdollDissolved = append(ge.onRagdollDissolved, fn)
}

func (ge *GameEvents) OnGameinstructorDraw(fn func(*GameEventGameinstructorDraw) error) {
	if ge.onGameinstructorDraw == nil {
		ge.onGameinstructorDraw = make([]func(*GameEventGameinstructorDraw) error, 0)
	}
	ge.onGameinstructorDraw = append(ge.onGameinstructorDraw, fn)
}

func (ge *GameEvents) OnGameinstructorNodraw(fn func(*GameEventGameinstructorNodraw) error) {
	if ge.onGameinstructorNodraw == nil {
		ge.onGameinstructorNodraw = make([]func(*GameEventGameinstructorNodraw) error, 0)
	}
	ge.onGameinstructorNodraw = append(ge.onGameinstructorNodraw, fn)
}

func (ge *GameEvents) OnMapTransition(fn func(*GameEventMapTransition) error) {
	if ge.onMapTransition == nil {
		ge.onMapTransition = make([]func(*GameEventMapTransition) error, 0)
	}
	ge.onMapTransition = append(ge.onMapTransition, fn)
}

func (ge *GameEvents) OnInstructorServerHintCreate(fn func(*GameEventInstructorServerHintCreate) error) {
	if ge.onInstructorServerHintCreate == nil {
		ge.onInstructorServerHintCreate = make([]func(*GameEventInstructorServerHintCreate) error, 0)
	}
	ge.onInstructorServerHintCreate = append(ge.onInstructorServerHintCreate, fn)
}

func (ge *GameEvents) OnInstructorServerHintStop(fn func(*GameEventInstructorServerHintStop) error) {
	if ge.onInstructorServerHintStop == nil {
		ge.onInstructorServerHintStop = make([]func(*GameEventInstructorServerHintStop) error, 0)
	}
	ge.onInstructorServerHintStop = append(ge.onInstructorServerHintStop, fn)
}

func (ge *GameEvents) OnChatNewMessage(fn func(*GameEventChatNewMessage) error) {
	if ge.onChatNewMessage == nil {
		ge.onChatNewMessage = make([]func(*GameEventChatNewMessage) error, 0)
	}
	ge.onChatNewMessage = append(ge.onChatNewMessage, fn)
}

func (ge *GameEvents) OnChatMembersChanged(fn func(*GameEventChatMembersChanged) error) {
	if ge.onChatMembersChanged == nil {
		ge.onChatMembersChanged = make([]func(*GameEventChatMembersChanged) error, 0)
	}
	ge.onChatMembersChanged = append(ge.onChatMembersChanged, fn)
}

func (ge *GameEvents) OnInventoryUpdated(fn func(*GameEventInventoryUpdated) error) {
	if ge.onInventoryUpdated == nil {
		ge.onInventoryUpdated = make([]func(*GameEventInventoryUpdated) error, 0)
	}
	ge.onInventoryUpdated = append(ge.onInventoryUpdated, fn)
}

func (ge *GameEvents) OnCartUpdated(fn func(*GameEventCartUpdated) error) {
	if ge.onCartUpdated == nil {
		ge.onCartUpdated = make([]func(*GameEventCartUpdated) error, 0)
	}
	ge.onCartUpdated = append(ge.onCartUpdated, fn)
}

func (ge *GameEvents) OnStorePricesheetUpdated(fn func(*GameEventStorePricesheetUpdated) error) {
	if ge.onStorePricesheetUpdated == nil {
		ge.onStorePricesheetUpdated = make([]func(*GameEventStorePricesheetUpdated) error, 0)
	}
	ge.onStorePricesheetUpdated = append(ge.onStorePricesheetUpdated, fn)
}

func (ge *GameEvents) OnGcConnected(fn func(*GameEventGcConnected) error) {
	if ge.onGcConnected == nil {
		ge.onGcConnected = make([]func(*GameEventGcConnected) error, 0)
	}
	ge.onGcConnected = append(ge.onGcConnected, fn)
}

func (ge *GameEvents) OnItemSchemaInitialized(fn func(*GameEventItemSchemaInitialized) error) {
	if ge.onItemSchemaInitialized == nil {
		ge.onItemSchemaInitialized = make([]func(*GameEventItemSchemaInitialized) error, 0)
	}
	ge.onItemSchemaInitialized = append(ge.onItemSchemaInitialized, fn)
}

func (ge *GameEvents) OnDropRateModified(fn func(*GameEventDropRateModified) error) {
	if ge.onDropRateModified == nil {
		ge.onDropRateModified = make([]func(*GameEventDropRateModified) error, 0)
	}
	ge.onDropRateModified = append(ge.onDropRateModified, fn)
}

func (ge *GameEvents) OnEventTicketModified(fn func(*GameEventEventTicketModified) error) {
	if ge.onEventTicketModified == nil {
		ge.onEventTicketModified = make([]func(*GameEventEventTicketModified) error, 0)
	}
	ge.onEventTicketModified = append(ge.onEventTicketModified, fn)
}

func (ge *GameEvents) OnModifierEvent(fn func(*GameEventModifierEvent) error) {
	if ge.onModifierEvent == nil {
		ge.onModifierEvent = make([]func(*GameEventModifierEvent) error, 0)
	}
	ge.onModifierEvent = append(ge.onModifierEvent, fn)
}

func (ge *GameEvents) OnDotaPlayerKill(fn func(*GameEventDotaPlayerKill) error) {
	if ge.onDotaPlayerKill == nil {
		ge.onDotaPlayerKill = make([]func(*GameEventDotaPlayerKill) error, 0)
	}
	ge.onDotaPlayerKill = append(ge.onDotaPlayerKill, fn)
}

func (ge *GameEvents) OnDotaPlayerDeny(fn func(*GameEventDotaPlayerDeny) error) {
	if ge.onDotaPlayerDeny == nil {
		ge.onDotaPlayerDeny = make([]func(*GameEventDotaPlayerDeny) error, 0)
	}
	ge.onDotaPlayerDeny = append(ge.onDotaPlayerDeny, fn)
}

func (ge *GameEvents) OnDotaBarracksKill(fn func(*GameEventDotaBarracksKill) error) {
	if ge.onDotaBarracksKill == nil {
		ge.onDotaBarracksKill = make([]func(*GameEventDotaBarracksKill) error, 0)
	}
	ge.onDotaBarracksKill = append(ge.onDotaBarracksKill, fn)
}

func (ge *GameEvents) OnDotaTowerKill(fn func(*GameEventDotaTowerKill) error) {
	if ge.onDotaTowerKill == nil {
		ge.onDotaTowerKill = make([]func(*GameEventDotaTowerKill) error, 0)
	}
	ge.onDotaTowerKill = append(ge.onDotaTowerKill, fn)
}

func (ge *GameEvents) OnDotaEffigyKill(fn func(*GameEventDotaEffigyKill) error) {
	if ge.onDotaEffigyKill == nil {
		ge.onDotaEffigyKill = make([]func(*GameEventDotaEffigyKill) error, 0)
	}
	ge.onDotaEffigyKill = append(ge.onDotaEffigyKill, fn)
}

func (ge *GameEvents) OnDotaRoshanKill(fn func(*GameEventDotaRoshanKill) error) {
	if ge.onDotaRoshanKill == nil {
		ge.onDotaRoshanKill = make([]func(*GameEventDotaRoshanKill) error, 0)
	}
	ge.onDotaRoshanKill = append(ge.onDotaRoshanKill, fn)
}

func (ge *GameEvents) OnDotaCourierLost(fn func(*GameEventDotaCourierLost) error) {
	if ge.onDotaCourierLost == nil {
		ge.onDotaCourierLost = make([]func(*GameEventDotaCourierLost) error, 0)
	}
	ge.onDotaCourierLost = append(ge.onDotaCourierLost, fn)
}

func (ge *GameEvents) OnDotaCourierRespawned(fn func(*GameEventDotaCourierRespawned) error) {
	if ge.onDotaCourierRespawned == nil {
		ge.onDotaCourierRespawned = make([]func(*GameEventDotaCourierRespawned) error, 0)
	}
	ge.onDotaCourierRespawned = append(ge.onDotaCourierRespawned, fn)
}

func (ge *GameEvents) OnDotaGlyphUsed(fn func(*GameEventDotaGlyphUsed) error) {
	if ge.onDotaGlyphUsed == nil {
		ge.onDotaGlyphUsed = make([]func(*GameEventDotaGlyphUsed) error, 0)
	}
	ge.onDotaGlyphUsed = append(ge.onDotaGlyphUsed, fn)
}

func (ge *GameEvents) OnDotaSuperCreeps(fn func(*GameEventDotaSuperCreeps) error) {
	if ge.onDotaSuperCreeps == nil {
		ge.onDotaSuperCreeps = make([]func(*GameEventDotaSuperCreeps) error, 0)
	}
	ge.onDotaSuperCreeps = append(ge.onDotaSuperCreeps, fn)
}

func (ge *GameEvents) OnDotaItemPurchase(fn func(*GameEventDotaItemPurchase) error) {
	if ge.onDotaItemPurchase == nil {
		ge.onDotaItemPurchase = make([]func(*GameEventDotaItemPurchase) error, 0)
	}
	ge.onDotaItemPurchase = append(ge.onDotaItemPurchase, fn)
}

func (ge *GameEvents) OnDotaItemGifted(fn func(*GameEventDotaItemGifted) error) {
	if ge.onDotaItemGifted == nil {
		ge.onDotaItemGifted = make([]func(*GameEventDotaItemGifted) error, 0)
	}
	ge.onDotaItemGifted = append(ge.onDotaItemGifted, fn)
}

func (ge *GameEvents) OnDotaRunePickup(fn func(*GameEventDotaRunePickup) error) {
	if ge.onDotaRunePickup == nil {
		ge.onDotaRunePickup = make([]func(*GameEventDotaRunePickup) error, 0)
	}
	ge.onDotaRunePickup = append(ge.onDotaRunePickup, fn)
}

func (ge *GameEvents) OnDotaRuneSpotted(fn func(*GameEventDotaRuneSpotted) error) {
	if ge.onDotaRuneSpotted == nil {
		ge.onDotaRuneSpotted = make([]func(*GameEventDotaRuneSpotted) error, 0)
	}
	ge.onDotaRuneSpotted = append(ge.onDotaRuneSpotted, fn)
}

func (ge *GameEvents) OnDotaItemSpotted(fn func(*GameEventDotaItemSpotted) error) {
	if ge.onDotaItemSpotted == nil {
		ge.onDotaItemSpotted = make([]func(*GameEventDotaItemSpotted) error, 0)
	}
	ge.onDotaItemSpotted = append(ge.onDotaItemSpotted, fn)
}

func (ge *GameEvents) OnDotaNoBattlePoints(fn func(*GameEventDotaNoBattlePoints) error) {
	if ge.onDotaNoBattlePoints == nil {
		ge.onDotaNoBattlePoints = make([]func(*GameEventDotaNoBattlePoints) error, 0)
	}
	ge.onDotaNoBattlePoints = append(ge.onDotaNoBattlePoints, fn)
}

func (ge *GameEvents) OnDotaChatInformational(fn func(*GameEventDotaChatInformational) error) {
	if ge.onDotaChatInformational == nil {
		ge.onDotaChatInformational = make([]func(*GameEventDotaChatInformational) error, 0)
	}
	ge.onDotaChatInformational = append(ge.onDotaChatInformational, fn)
}

func (ge *GameEvents) OnDotaActionItem(fn func(*GameEventDotaActionItem) error) {
	if ge.onDotaActionItem == nil {
		ge.onDotaActionItem = make([]func(*GameEventDotaActionItem) error, 0)
	}
	ge.onDotaActionItem = append(ge.onDotaActionItem, fn)
}

func (ge *GameEvents) OnDotaChatBanNotification(fn func(*GameEventDotaChatBanNotification) error) {
	if ge.onDotaChatBanNotification == nil {
		ge.onDotaChatBanNotification = make([]func(*GameEventDotaChatBanNotification) error, 0)
	}
	ge.onDotaChatBanNotification = append(ge.onDotaChatBanNotification, fn)
}

func (ge *GameEvents) OnDotaChatEvent(fn func(*GameEventDotaChatEvent) error) {
	if ge.onDotaChatEvent == nil {
		ge.onDotaChatEvent = make([]func(*GameEventDotaChatEvent) error, 0)
	}
	ge.onDotaChatEvent = append(ge.onDotaChatEvent, fn)
}

func (ge *GameEvents) OnDotaChatTimedReward(fn func(*GameEventDotaChatTimedReward) error) {
	if ge.onDotaChatTimedReward == nil {
		ge.onDotaChatTimedReward = make([]func(*GameEventDotaChatTimedReward) error, 0)
	}
	ge.onDotaChatTimedReward = append(ge.onDotaChatTimedReward, fn)
}

func (ge *GameEvents) OnDotaPauseEvent(fn func(*GameEventDotaPauseEvent) error) {
	if ge.onDotaPauseEvent == nil {
		ge.onDotaPauseEvent = make([]func(*GameEventDotaPauseEvent) error, 0)
	}
	ge.onDotaPauseEvent = append(ge.onDotaPauseEvent, fn)
}

func (ge *GameEvents) OnDotaChatKillStreak(fn func(*GameEventDotaChatKillStreak) error) {
	if ge.onDotaChatKillStreak == nil {
		ge.onDotaChatKillStreak = make([]func(*GameEventDotaChatKillStreak) error, 0)
	}
	ge.onDotaChatKillStreak = append(ge.onDotaChatKillStreak, fn)
}

func (ge *GameEvents) OnDotaChatFirstBlood(fn func(*GameEventDotaChatFirstBlood) error) {
	if ge.onDotaChatFirstBlood == nil {
		ge.onDotaChatFirstBlood = make([]func(*GameEventDotaChatFirstBlood) error, 0)
	}
	ge.onDotaChatFirstBlood = append(ge.onDotaChatFirstBlood, fn)
}

func (ge *GameEvents) OnDotaChatAssassinAnnounce(fn func(*GameEventDotaChatAssassinAnnounce) error) {
	if ge.onDotaChatAssassinAnnounce == nil {
		ge.onDotaChatAssassinAnnounce = make([]func(*GameEventDotaChatAssassinAnnounce) error, 0)
	}
	ge.onDotaChatAssassinAnnounce = append(ge.onDotaChatAssassinAnnounce, fn)
}

func (ge *GameEvents) OnDotaChatAssassinDenied(fn func(*GameEventDotaChatAssassinDenied) error) {
	if ge.onDotaChatAssassinDenied == nil {
		ge.onDotaChatAssassinDenied = make([]func(*GameEventDotaChatAssassinDenied) error, 0)
	}
	ge.onDotaChatAssassinDenied = append(ge.onDotaChatAssassinDenied, fn)
}

func (ge *GameEvents) OnDotaChatAssassinSuccess(fn func(*GameEventDotaChatAssassinSuccess) error) {
	if ge.onDotaChatAssassinSuccess == nil {
		ge.onDotaChatAssassinSuccess = make([]func(*GameEventDotaChatAssassinSuccess) error, 0)
	}
	ge.onDotaChatAssassinSuccess = append(ge.onDotaChatAssassinSuccess, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateHeroSelection(fn func(*GameEventDotaPlayerUpdateHeroSelection) error) {
	if ge.onDotaPlayerUpdateHeroSelection == nil {
		ge.onDotaPlayerUpdateHeroSelection = make([]func(*GameEventDotaPlayerUpdateHeroSelection) error, 0)
	}
	ge.onDotaPlayerUpdateHeroSelection = append(ge.onDotaPlayerUpdateHeroSelection, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateSelectedUnit(fn func(*GameEventDotaPlayerUpdateSelectedUnit) error) {
	if ge.onDotaPlayerUpdateSelectedUnit == nil {
		ge.onDotaPlayerUpdateSelectedUnit = make([]func(*GameEventDotaPlayerUpdateSelectedUnit) error, 0)
	}
	ge.onDotaPlayerUpdateSelectedUnit = append(ge.onDotaPlayerUpdateSelectedUnit, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateQueryUnit(fn func(*GameEventDotaPlayerUpdateQueryUnit) error) {
	if ge.onDotaPlayerUpdateQueryUnit == nil {
		ge.onDotaPlayerUpdateQueryUnit = make([]func(*GameEventDotaPlayerUpdateQueryUnit) error, 0)
	}
	ge.onDotaPlayerUpdateQueryUnit = append(ge.onDotaPlayerUpdateQueryUnit, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateKillcamUnit(fn func(*GameEventDotaPlayerUpdateKillcamUnit) error) {
	if ge.onDotaPlayerUpdateKillcamUnit == nil {
		ge.onDotaPlayerUpdateKillcamUnit = make([]func(*GameEventDotaPlayerUpdateKillcamUnit) error, 0)
	}
	ge.onDotaPlayerUpdateKillcamUnit = append(ge.onDotaPlayerUpdateKillcamUnit, fn)
}

func (ge *GameEvents) OnDotaPlayerTakeTowerDamage(fn func(*GameEventDotaPlayerTakeTowerDamage) error) {
	if ge.onDotaPlayerTakeTowerDamage == nil {
		ge.onDotaPlayerTakeTowerDamage = make([]func(*GameEventDotaPlayerTakeTowerDamage) error, 0)
	}
	ge.onDotaPlayerTakeTowerDamage = append(ge.onDotaPlayerTakeTowerDamage, fn)
}

func (ge *GameEvents) OnDotaHudErrorMessage(fn func(*GameEventDotaHudErrorMessage) error) {
	if ge.onDotaHudErrorMessage == nil {
		ge.onDotaHudErrorMessage = make([]func(*GameEventDotaHudErrorMessage) error, 0)
	}
	ge.onDotaHudErrorMessage = append(ge.onDotaHudErrorMessage, fn)
}

func (ge *GameEvents) OnDotaActionSuccess(fn func(*GameEventDotaActionSuccess) error) {
	if ge.onDotaActionSuccess == nil {
		ge.onDotaActionSuccess = make([]func(*GameEventDotaActionSuccess) error, 0)
	}
	ge.onDotaActionSuccess = append(ge.onDotaActionSuccess, fn)
}

func (ge *GameEvents) OnDotaStartingPositionChanged(fn func(*GameEventDotaStartingPositionChanged) error) {
	if ge.onDotaStartingPositionChanged == nil {
		ge.onDotaStartingPositionChanged = make([]func(*GameEventDotaStartingPositionChanged) error, 0)
	}
	ge.onDotaStartingPositionChanged = append(ge.onDotaStartingPositionChanged, fn)
}

func (ge *GameEvents) OnDotaMoneyChanged(fn func(*GameEventDotaMoneyChanged) error) {
	if ge.onDotaMoneyChanged == nil {
		ge.onDotaMoneyChanged = make([]func(*GameEventDotaMoneyChanged) error, 0)
	}
	ge.onDotaMoneyChanged = append(ge.onDotaMoneyChanged, fn)
}

func (ge *GameEvents) OnDotaEnemyMoneyChanged(fn func(*GameEventDotaEnemyMoneyChanged) error) {
	if ge.onDotaEnemyMoneyChanged == nil {
		ge.onDotaEnemyMoneyChanged = make([]func(*GameEventDotaEnemyMoneyChanged) error, 0)
	}
	ge.onDotaEnemyMoneyChanged = append(ge.onDotaEnemyMoneyChanged, fn)
}

func (ge *GameEvents) OnDotaPortraitUnitStatsChanged(fn func(*GameEventDotaPortraitUnitStatsChanged) error) {
	if ge.onDotaPortraitUnitStatsChanged == nil {
		ge.onDotaPortraitUnitStatsChanged = make([]func(*GameEventDotaPortraitUnitStatsChanged) error, 0)
	}
	ge.onDotaPortraitUnitStatsChanged = append(ge.onDotaPortraitUnitStatsChanged, fn)
}

func (ge *GameEvents) OnDotaPortraitUnitModifiersChanged(fn func(*GameEventDotaPortraitUnitModifiersChanged) error) {
	if ge.onDotaPortraitUnitModifiersChanged == nil {
		ge.onDotaPortraitUnitModifiersChanged = make([]func(*GameEventDotaPortraitUnitModifiersChanged) error, 0)
	}
	ge.onDotaPortraitUnitModifiersChanged = append(ge.onDotaPortraitUnitModifiersChanged, fn)
}

func (ge *GameEvents) OnDotaForcePortraitUpdate(fn func(*GameEventDotaForcePortraitUpdate) error) {
	if ge.onDotaForcePortraitUpdate == nil {
		ge.onDotaForcePortraitUpdate = make([]func(*GameEventDotaForcePortraitUpdate) error, 0)
	}
	ge.onDotaForcePortraitUpdate = append(ge.onDotaForcePortraitUpdate, fn)
}

func (ge *GameEvents) OnDotaInventoryChanged(fn func(*GameEventDotaInventoryChanged) error) {
	if ge.onDotaInventoryChanged == nil {
		ge.onDotaInventoryChanged = make([]func(*GameEventDotaInventoryChanged) error, 0)
	}
	ge.onDotaInventoryChanged = append(ge.onDotaInventoryChanged, fn)
}

func (ge *GameEvents) OnDotaItemPickedUp(fn func(*GameEventDotaItemPickedUp) error) {
	if ge.onDotaItemPickedUp == nil {
		ge.onDotaItemPickedUp = make([]func(*GameEventDotaItemPickedUp) error, 0)
	}
	ge.onDotaItemPickedUp = append(ge.onDotaItemPickedUp, fn)
}

func (ge *GameEvents) OnDotaInventoryItemChanged(fn func(*GameEventDotaInventoryItemChanged) error) {
	if ge.onDotaInventoryItemChanged == nil {
		ge.onDotaInventoryItemChanged = make([]func(*GameEventDotaInventoryItemChanged) error, 0)
	}
	ge.onDotaInventoryItemChanged = append(ge.onDotaInventoryItemChanged, fn)
}

func (ge *GameEvents) OnDotaAbilityChanged(fn func(*GameEventDotaAbilityChanged) error) {
	if ge.onDotaAbilityChanged == nil {
		ge.onDotaAbilityChanged = make([]func(*GameEventDotaAbilityChanged) error, 0)
	}
	ge.onDotaAbilityChanged = append(ge.onDotaAbilityChanged, fn)
}

func (ge *GameEvents) OnDotaPortraitAbilityLayoutChanged(fn func(*GameEventDotaPortraitAbilityLayoutChanged) error) {
	if ge.onDotaPortraitAbilityLayoutChanged == nil {
		ge.onDotaPortraitAbilityLayoutChanged = make([]func(*GameEventDotaPortraitAbilityLayoutChanged) error, 0)
	}
	ge.onDotaPortraitAbilityLayoutChanged = append(ge.onDotaPortraitAbilityLayoutChanged, fn)
}

func (ge *GameEvents) OnDotaInventoryItemAdded(fn func(*GameEventDotaInventoryItemAdded) error) {
	if ge.onDotaInventoryItemAdded == nil {
		ge.onDotaInventoryItemAdded = make([]func(*GameEventDotaInventoryItemAdded) error, 0)
	}
	ge.onDotaInventoryItemAdded = append(ge.onDotaInventoryItemAdded, fn)
}

func (ge *GameEvents) OnDotaInventoryChangedQueryUnit(fn func(*GameEventDotaInventoryChangedQueryUnit) error) {
	if ge.onDotaInventoryChangedQueryUnit == nil {
		ge.onDotaInventoryChangedQueryUnit = make([]func(*GameEventDotaInventoryChangedQueryUnit) error, 0)
	}
	ge.onDotaInventoryChangedQueryUnit = append(ge.onDotaInventoryChangedQueryUnit, fn)
}

func (ge *GameEvents) OnDotaLinkClicked(fn func(*GameEventDotaLinkClicked) error) {
	if ge.onDotaLinkClicked == nil {
		ge.onDotaLinkClicked = make([]func(*GameEventDotaLinkClicked) error, 0)
	}
	ge.onDotaLinkClicked = append(ge.onDotaLinkClicked, fn)
}

func (ge *GameEvents) OnDotaSetQuickBuy(fn func(*GameEventDotaSetQuickBuy) error) {
	if ge.onDotaSetQuickBuy == nil {
		ge.onDotaSetQuickBuy = make([]func(*GameEventDotaSetQuickBuy) error, 0)
	}
	ge.onDotaSetQuickBuy = append(ge.onDotaSetQuickBuy, fn)
}

func (ge *GameEvents) OnDotaQuickBuyChanged(fn func(*GameEventDotaQuickBuyChanged) error) {
	if ge.onDotaQuickBuyChanged == nil {
		ge.onDotaQuickBuyChanged = make([]func(*GameEventDotaQuickBuyChanged) error, 0)
	}
	ge.onDotaQuickBuyChanged = append(ge.onDotaQuickBuyChanged, fn)
}

func (ge *GameEvents) OnDotaPlayerShopChanged(fn func(*GameEventDotaPlayerShopChanged) error) {
	if ge.onDotaPlayerShopChanged == nil {
		ge.onDotaPlayerShopChanged = make([]func(*GameEventDotaPlayerShopChanged) error, 0)
	}
	ge.onDotaPlayerShopChanged = append(ge.onDotaPlayerShopChanged, fn)
}

func (ge *GameEvents) OnDotaPlayerShowKillcam(fn func(*GameEventDotaPlayerShowKillcam) error) {
	if ge.onDotaPlayerShowKillcam == nil {
		ge.onDotaPlayerShowKillcam = make([]func(*GameEventDotaPlayerShowKillcam) error, 0)
	}
	ge.onDotaPlayerShowKillcam = append(ge.onDotaPlayerShowKillcam, fn)
}

func (ge *GameEvents) OnDotaPlayerShowMinikillcam(fn func(*GameEventDotaPlayerShowMinikillcam) error) {
	if ge.onDotaPlayerShowMinikillcam == nil {
		ge.onDotaPlayerShowMinikillcam = make([]func(*GameEventDotaPlayerShowMinikillcam) error, 0)
	}
	ge.onDotaPlayerShowMinikillcam = append(ge.onDotaPlayerShowMinikillcam, fn)
}

func (ge *GameEvents) OnGcUserSessionCreated(fn func(*GameEventGcUserSessionCreated) error) {
	if ge.onGcUserSessionCreated == nil {
		ge.onGcUserSessionCreated = make([]func(*GameEventGcUserSessionCreated) error, 0)
	}
	ge.onGcUserSessionCreated = append(ge.onGcUserSessionCreated, fn)
}

func (ge *GameEvents) OnTeamDataUpdated(fn func(*GameEventTeamDataUpdated) error) {
	if ge.onTeamDataUpdated == nil {
		ge.onTeamDataUpdated = make([]func(*GameEventTeamDataUpdated) error, 0)
	}
	ge.onTeamDataUpdated = append(ge.onTeamDataUpdated, fn)
}

func (ge *GameEvents) OnGuildDataUpdated(fn func(*GameEventGuildDataUpdated) error) {
	if ge.onGuildDataUpdated == nil {
		ge.onGuildDataUpdated = make([]func(*GameEventGuildDataUpdated) error, 0)
	}
	ge.onGuildDataUpdated = append(ge.onGuildDataUpdated, fn)
}

func (ge *GameEvents) OnGuildOpenPartiesUpdated(fn func(*GameEventGuildOpenPartiesUpdated) error) {
	if ge.onGuildOpenPartiesUpdated == nil {
		ge.onGuildOpenPartiesUpdated = make([]func(*GameEventGuildOpenPartiesUpdated) error, 0)
	}
	ge.onGuildOpenPartiesUpdated = append(ge.onGuildOpenPartiesUpdated, fn)
}

func (ge *GameEvents) OnFantasyUpdated(fn func(*GameEventFantasyUpdated) error) {
	if ge.onFantasyUpdated == nil {
		ge.onFantasyUpdated = make([]func(*GameEventFantasyUpdated) error, 0)
	}
	ge.onFantasyUpdated = append(ge.onFantasyUpdated, fn)
}

func (ge *GameEvents) OnFantasyLeagueChanged(fn func(*GameEventFantasyLeagueChanged) error) {
	if ge.onFantasyLeagueChanged == nil {
		ge.onFantasyLeagueChanged = make([]func(*GameEventFantasyLeagueChanged) error, 0)
	}
	ge.onFantasyLeagueChanged = append(ge.onFantasyLeagueChanged, fn)
}

func (ge *GameEvents) OnFantasyScoreInfoChanged(fn func(*GameEventFantasyScoreInfoChanged) error) {
	if ge.onFantasyScoreInfoChanged == nil {
		ge.onFantasyScoreInfoChanged = make([]func(*GameEventFantasyScoreInfoChanged) error, 0)
	}
	ge.onFantasyScoreInfoChanged = append(ge.onFantasyScoreInfoChanged, fn)
}

func (ge *GameEvents) OnPlayerInfoUpdated(fn func(*GameEventPlayerInfoUpdated) error) {
	if ge.onPlayerInfoUpdated == nil {
		ge.onPlayerInfoUpdated = make([]func(*GameEventPlayerInfoUpdated) error, 0)
	}
	ge.onPlayerInfoUpdated = append(ge.onPlayerInfoUpdated, fn)
}

func (ge *GameEvents) OnPlayerInfoIndividualUpdated(fn func(*GameEventPlayerInfoIndividualUpdated) error) {
	if ge.onPlayerInfoIndividualUpdated == nil {
		ge.onPlayerInfoIndividualUpdated = make([]func(*GameEventPlayerInfoIndividualUpdated) error, 0)
	}
	ge.onPlayerInfoIndividualUpdated = append(ge.onPlayerInfoIndividualUpdated, fn)
}

func (ge *GameEvents) OnGameRulesStateChange(fn func(*GameEventGameRulesStateChange) error) {
	if ge.onGameRulesStateChange == nil {
		ge.onGameRulesStateChange = make([]func(*GameEventGameRulesStateChange) error, 0)
	}
	ge.onGameRulesStateChange = append(ge.onGameRulesStateChange, fn)
}

func (ge *GameEvents) OnMatchHistoryUpdated(fn func(*GameEventMatchHistoryUpdated) error) {
	if ge.onMatchHistoryUpdated == nil {
		ge.onMatchHistoryUpdated = make([]func(*GameEventMatchHistoryUpdated) error, 0)
	}
	ge.onMatchHistoryUpdated = append(ge.onMatchHistoryUpdated, fn)
}

func (ge *GameEvents) OnMatchDetailsUpdated(fn func(*GameEventMatchDetailsUpdated) error) {
	if ge.onMatchDetailsUpdated == nil {
		ge.onMatchDetailsUpdated = make([]func(*GameEventMatchDetailsUpdated) error, 0)
	}
	ge.onMatchDetailsUpdated = append(ge.onMatchDetailsUpdated, fn)
}

func (ge *GameEvents) OnLiveGamesUpdated(fn func(*GameEventLiveGamesUpdated) error) {
	if ge.onLiveGamesUpdated == nil {
		ge.onLiveGamesUpdated = make([]func(*GameEventLiveGamesUpdated) error, 0)
	}
	ge.onLiveGamesUpdated = append(ge.onLiveGamesUpdated, fn)
}

func (ge *GameEvents) OnRecentMatchesUpdated(fn func(*GameEventRecentMatchesUpdated) error) {
	if ge.onRecentMatchesUpdated == nil {
		ge.onRecentMatchesUpdated = make([]func(*GameEventRecentMatchesUpdated) error, 0)
	}
	ge.onRecentMatchesUpdated = append(ge.onRecentMatchesUpdated, fn)
}

func (ge *GameEvents) OnNewsUpdated(fn func(*GameEventNewsUpdated) error) {
	if ge.onNewsUpdated == nil {
		ge.onNewsUpdated = make([]func(*GameEventNewsUpdated) error, 0)
	}
	ge.onNewsUpdated = append(ge.onNewsUpdated, fn)
}

func (ge *GameEvents) OnPersonaUpdated(fn func(*GameEventPersonaUpdated) error) {
	if ge.onPersonaUpdated == nil {
		ge.onPersonaUpdated = make([]func(*GameEventPersonaUpdated) error, 0)
	}
	ge.onPersonaUpdated = append(ge.onPersonaUpdated, fn)
}

func (ge *GameEvents) OnTournamentStateUpdated(fn func(*GameEventTournamentStateUpdated) error) {
	if ge.onTournamentStateUpdated == nil {
		ge.onTournamentStateUpdated = make([]func(*GameEventTournamentStateUpdated) error, 0)
	}
	ge.onTournamentStateUpdated = append(ge.onTournamentStateUpdated, fn)
}

func (ge *GameEvents) OnPartyUpdated(fn func(*GameEventPartyUpdated) error) {
	if ge.onPartyUpdated == nil {
		ge.onPartyUpdated = make([]func(*GameEventPartyUpdated) error, 0)
	}
	ge.onPartyUpdated = append(ge.onPartyUpdated, fn)
}

func (ge *GameEvents) OnLobbyUpdated(fn func(*GameEventLobbyUpdated) error) {
	if ge.onLobbyUpdated == nil {
		ge.onLobbyUpdated = make([]func(*GameEventLobbyUpdated) error, 0)
	}
	ge.onLobbyUpdated = append(ge.onLobbyUpdated, fn)
}

func (ge *GameEvents) OnDashboardCachesCleared(fn func(*GameEventDashboardCachesCleared) error) {
	if ge.onDashboardCachesCleared == nil {
		ge.onDashboardCachesCleared = make([]func(*GameEventDashboardCachesCleared) error, 0)
	}
	ge.onDashboardCachesCleared = append(ge.onDashboardCachesCleared, fn)
}

func (ge *GameEvents) OnLastHit(fn func(*GameEventLastHit) error) {
	if ge.onLastHit == nil {
		ge.onLastHit = make([]func(*GameEventLastHit) error, 0)
	}
	ge.onLastHit = append(ge.onLastHit, fn)
}

func (ge *GameEvents) OnPlayerCompletedGame(fn func(*GameEventPlayerCompletedGame) error) {
	if ge.onPlayerCompletedGame == nil {
		ge.onPlayerCompletedGame = make([]func(*GameEventPlayerCompletedGame) error, 0)
	}
	ge.onPlayerCompletedGame = append(ge.onPlayerCompletedGame, fn)
}

func (ge *GameEvents) OnPlayerReconnected(fn func(*GameEventPlayerReconnected) error) {
	if ge.onPlayerReconnected == nil {
		ge.onPlayerReconnected = make([]func(*GameEventPlayerReconnected) error, 0)
	}
	ge.onPlayerReconnected = append(ge.onPlayerReconnected, fn)
}

func (ge *GameEvents) OnNommedTree(fn func(*GameEventNommedTree) error) {
	if ge.onNommedTree == nil {
		ge.onNommedTree = make([]func(*GameEventNommedTree) error, 0)
	}
	ge.onNommedTree = append(ge.onNommedTree, fn)
}

func (ge *GameEvents) OnDotaRuneActivatedServer(fn func(*GameEventDotaRuneActivatedServer) error) {
	if ge.onDotaRuneActivatedServer == nil {
		ge.onDotaRuneActivatedServer = make([]func(*GameEventDotaRuneActivatedServer) error, 0)
	}
	ge.onDotaRuneActivatedServer = append(ge.onDotaRuneActivatedServer, fn)
}

func (ge *GameEvents) OnDotaPlayerGainedLevel(fn func(*GameEventDotaPlayerGainedLevel) error) {
	if ge.onDotaPlayerGainedLevel == nil {
		ge.onDotaPlayerGainedLevel = make([]func(*GameEventDotaPlayerGainedLevel) error, 0)
	}
	ge.onDotaPlayerGainedLevel = append(ge.onDotaPlayerGainedLevel, fn)
}

func (ge *GameEvents) OnDotaPlayerLearnedAbility(fn func(*GameEventDotaPlayerLearnedAbility) error) {
	if ge.onDotaPlayerLearnedAbility == nil {
		ge.onDotaPlayerLearnedAbility = make([]func(*GameEventDotaPlayerLearnedAbility) error, 0)
	}
	ge.onDotaPlayerLearnedAbility = append(ge.onDotaPlayerLearnedAbility, fn)
}

func (ge *GameEvents) OnDotaPlayerUsedAbility(fn func(*GameEventDotaPlayerUsedAbility) error) {
	if ge.onDotaPlayerUsedAbility == nil {
		ge.onDotaPlayerUsedAbility = make([]func(*GameEventDotaPlayerUsedAbility) error, 0)
	}
	ge.onDotaPlayerUsedAbility = append(ge.onDotaPlayerUsedAbility, fn)
}

func (ge *GameEvents) OnDotaNonPlayerUsedAbility(fn func(*GameEventDotaNonPlayerUsedAbility) error) {
	if ge.onDotaNonPlayerUsedAbility == nil {
		ge.onDotaNonPlayerUsedAbility = make([]func(*GameEventDotaNonPlayerUsedAbility) error, 0)
	}
	ge.onDotaNonPlayerUsedAbility = append(ge.onDotaNonPlayerUsedAbility, fn)
}

func (ge *GameEvents) OnDotaPlayerBeginCast(fn func(*GameEventDotaPlayerBeginCast) error) {
	if ge.onDotaPlayerBeginCast == nil {
		ge.onDotaPlayerBeginCast = make([]func(*GameEventDotaPlayerBeginCast) error, 0)
	}
	ge.onDotaPlayerBeginCast = append(ge.onDotaPlayerBeginCast, fn)
}

func (ge *GameEvents) OnDotaNonPlayerBeginCast(fn func(*GameEventDotaNonPlayerBeginCast) error) {
	if ge.onDotaNonPlayerBeginCast == nil {
		ge.onDotaNonPlayerBeginCast = make([]func(*GameEventDotaNonPlayerBeginCast) error, 0)
	}
	ge.onDotaNonPlayerBeginCast = append(ge.onDotaNonPlayerBeginCast, fn)
}

func (ge *GameEvents) OnDotaAbilityChannelFinished(fn func(*GameEventDotaAbilityChannelFinished) error) {
	if ge.onDotaAbilityChannelFinished == nil {
		ge.onDotaAbilityChannelFinished = make([]func(*GameEventDotaAbilityChannelFinished) error, 0)
	}
	ge.onDotaAbilityChannelFinished = append(ge.onDotaAbilityChannelFinished, fn)
}

func (ge *GameEvents) OnDotaHoldoutReviveComplete(fn func(*GameEventDotaHoldoutReviveComplete) error) {
	if ge.onDotaHoldoutReviveComplete == nil {
		ge.onDotaHoldoutReviveComplete = make([]func(*GameEventDotaHoldoutReviveComplete) error, 0)
	}
	ge.onDotaHoldoutReviveComplete = append(ge.onDotaHoldoutReviveComplete, fn)
}

func (ge *GameEvents) OnDotaPlayerKilled(fn func(*GameEventDotaPlayerKilled) error) {
	if ge.onDotaPlayerKilled == nil {
		ge.onDotaPlayerKilled = make([]func(*GameEventDotaPlayerKilled) error, 0)
	}
	ge.onDotaPlayerKilled = append(ge.onDotaPlayerKilled, fn)
}

func (ge *GameEvents) OnBindpanelOpen(fn func(*GameEventBindpanelOpen) error) {
	if ge.onBindpanelOpen == nil {
		ge.onBindpanelOpen = make([]func(*GameEventBindpanelOpen) error, 0)
	}
	ge.onBindpanelOpen = append(ge.onBindpanelOpen, fn)
}

func (ge *GameEvents) OnBindpanelClose(fn func(*GameEventBindpanelClose) error) {
	if ge.onBindpanelClose == nil {
		ge.onBindpanelClose = make([]func(*GameEventBindpanelClose) error, 0)
	}
	ge.onBindpanelClose = append(ge.onBindpanelClose, fn)
}

func (ge *GameEvents) OnKeybindChanged(fn func(*GameEventKeybindChanged) error) {
	if ge.onKeybindChanged == nil {
		ge.onKeybindChanged = make([]func(*GameEventKeybindChanged) error, 0)
	}
	ge.onKeybindChanged = append(ge.onKeybindChanged, fn)
}

func (ge *GameEvents) OnDotaItemDragBegin(fn func(*GameEventDotaItemDragBegin) error) {
	if ge.onDotaItemDragBegin == nil {
		ge.onDotaItemDragBegin = make([]func(*GameEventDotaItemDragBegin) error, 0)
	}
	ge.onDotaItemDragBegin = append(ge.onDotaItemDragBegin, fn)
}

func (ge *GameEvents) OnDotaItemDragEnd(fn func(*GameEventDotaItemDragEnd) error) {
	if ge.onDotaItemDragEnd == nil {
		ge.onDotaItemDragEnd = make([]func(*GameEventDotaItemDragEnd) error, 0)
	}
	ge.onDotaItemDragEnd = append(ge.onDotaItemDragEnd, fn)
}

func (ge *GameEvents) OnDotaShopItemDragBegin(fn func(*GameEventDotaShopItemDragBegin) error) {
	if ge.onDotaShopItemDragBegin == nil {
		ge.onDotaShopItemDragBegin = make([]func(*GameEventDotaShopItemDragBegin) error, 0)
	}
	ge.onDotaShopItemDragBegin = append(ge.onDotaShopItemDragBegin, fn)
}

func (ge *GameEvents) OnDotaShopItemDragEnd(fn func(*GameEventDotaShopItemDragEnd) error) {
	if ge.onDotaShopItemDragEnd == nil {
		ge.onDotaShopItemDragEnd = make([]func(*GameEventDotaShopItemDragEnd) error, 0)
	}
	ge.onDotaShopItemDragEnd = append(ge.onDotaShopItemDragEnd, fn)
}

func (ge *GameEvents) OnDotaItemPurchased(fn func(*GameEventDotaItemPurchased) error) {
	if ge.onDotaItemPurchased == nil {
		ge.onDotaItemPurchased = make([]func(*GameEventDotaItemPurchased) error, 0)
	}
	ge.onDotaItemPurchased = append(ge.onDotaItemPurchased, fn)
}

func (ge *GameEvents) OnDotaItemCombined(fn func(*GameEventDotaItemCombined) error) {
	if ge.onDotaItemCombined == nil {
		ge.onDotaItemCombined = make([]func(*GameEventDotaItemCombined) error, 0)
	}
	ge.onDotaItemCombined = append(ge.onDotaItemCombined, fn)
}

func (ge *GameEvents) OnDotaItemUsed(fn func(*GameEventDotaItemUsed) error) {
	if ge.onDotaItemUsed == nil {
		ge.onDotaItemUsed = make([]func(*GameEventDotaItemUsed) error, 0)
	}
	ge.onDotaItemUsed = append(ge.onDotaItemUsed, fn)
}

func (ge *GameEvents) OnDotaItemAutoPurchase(fn func(*GameEventDotaItemAutoPurchase) error) {
	if ge.onDotaItemAutoPurchase == nil {
		ge.onDotaItemAutoPurchase = make([]func(*GameEventDotaItemAutoPurchase) error, 0)
	}
	ge.onDotaItemAutoPurchase = append(ge.onDotaItemAutoPurchase, fn)
}

func (ge *GameEvents) OnDotaUnitEvent(fn func(*GameEventDotaUnitEvent) error) {
	if ge.onDotaUnitEvent == nil {
		ge.onDotaUnitEvent = make([]func(*GameEventDotaUnitEvent) error, 0)
	}
	ge.onDotaUnitEvent = append(ge.onDotaUnitEvent, fn)
}

func (ge *GameEvents) OnDotaQuestStarted(fn func(*GameEventDotaQuestStarted) error) {
	if ge.onDotaQuestStarted == nil {
		ge.onDotaQuestStarted = make([]func(*GameEventDotaQuestStarted) error, 0)
	}
	ge.onDotaQuestStarted = append(ge.onDotaQuestStarted, fn)
}

func (ge *GameEvents) OnDotaQuestCompleted(fn func(*GameEventDotaQuestCompleted) error) {
	if ge.onDotaQuestCompleted == nil {
		ge.onDotaQuestCompleted = make([]func(*GameEventDotaQuestCompleted) error, 0)
	}
	ge.onDotaQuestCompleted = append(ge.onDotaQuestCompleted, fn)
}

func (ge *GameEvents) OnGameuiActivated(fn func(*GameEventGameuiActivated) error) {
	if ge.onGameuiActivated == nil {
		ge.onGameuiActivated = make([]func(*GameEventGameuiActivated) error, 0)
	}
	ge.onGameuiActivated = append(ge.onGameuiActivated, fn)
}

func (ge *GameEvents) OnGameuiHidden(fn func(*GameEventGameuiHidden) error) {
	if ge.onGameuiHidden == nil {
		ge.onGameuiHidden = make([]func(*GameEventGameuiHidden) error, 0)
	}
	ge.onGameuiHidden = append(ge.onGameuiHidden, fn)
}

func (ge *GameEvents) OnPlayerFullyjoined(fn func(*GameEventPlayerFullyjoined) error) {
	if ge.onPlayerFullyjoined == nil {
		ge.onPlayerFullyjoined = make([]func(*GameEventPlayerFullyjoined) error, 0)
	}
	ge.onPlayerFullyjoined = append(ge.onPlayerFullyjoined, fn)
}

func (ge *GameEvents) OnDotaSpectateHero(fn func(*GameEventDotaSpectateHero) error) {
	if ge.onDotaSpectateHero == nil {
		ge.onDotaSpectateHero = make([]func(*GameEventDotaSpectateHero) error, 0)
	}
	ge.onDotaSpectateHero = append(ge.onDotaSpectateHero, fn)
}

func (ge *GameEvents) OnDotaMatchDone(fn func(*GameEventDotaMatchDone) error) {
	if ge.onDotaMatchDone == nil {
		ge.onDotaMatchDone = make([]func(*GameEventDotaMatchDone) error, 0)
	}
	ge.onDotaMatchDone = append(ge.onDotaMatchDone, fn)
}

func (ge *GameEvents) OnDotaMatchDoneClient(fn func(*GameEventDotaMatchDoneClient) error) {
	if ge.onDotaMatchDoneClient == nil {
		ge.onDotaMatchDoneClient = make([]func(*GameEventDotaMatchDoneClient) error, 0)
	}
	ge.onDotaMatchDoneClient = append(ge.onDotaMatchDoneClient, fn)
}

func (ge *GameEvents) OnSetInstructorGroupEnabled(fn func(*GameEventSetInstructorGroupEnabled) error) {
	if ge.onSetInstructorGroupEnabled == nil {
		ge.onSetInstructorGroupEnabled = make([]func(*GameEventSetInstructorGroupEnabled) error, 0)
	}
	ge.onSetInstructorGroupEnabled = append(ge.onSetInstructorGroupEnabled, fn)
}

func (ge *GameEvents) OnJoinedChatChannel(fn func(*GameEventJoinedChatChannel) error) {
	if ge.onJoinedChatChannel == nil {
		ge.onJoinedChatChannel = make([]func(*GameEventJoinedChatChannel) error, 0)
	}
	ge.onJoinedChatChannel = append(ge.onJoinedChatChannel, fn)
}

func (ge *GameEvents) OnLeftChatChannel(fn func(*GameEventLeftChatChannel) error) {
	if ge.onLeftChatChannel == nil {
		ge.onLeftChatChannel = make([]func(*GameEventLeftChatChannel) error, 0)
	}
	ge.onLeftChatChannel = append(ge.onLeftChatChannel, fn)
}

func (ge *GameEvents) OnGcChatChannelListUpdated(fn func(*GameEventGcChatChannelListUpdated) error) {
	if ge.onGcChatChannelListUpdated == nil {
		ge.onGcChatChannelListUpdated = make([]func(*GameEventGcChatChannelListUpdated) error, 0)
	}
	ge.onGcChatChannelListUpdated = append(ge.onGcChatChannelListUpdated, fn)
}

func (ge *GameEvents) OnTodayMessagesUpdated(fn func(*GameEventTodayMessagesUpdated) error) {
	if ge.onTodayMessagesUpdated == nil {
		ge.onTodayMessagesUpdated = make([]func(*GameEventTodayMessagesUpdated) error, 0)
	}
	ge.onTodayMessagesUpdated = append(ge.onTodayMessagesUpdated, fn)
}

func (ge *GameEvents) OnFileDownloaded(fn func(*GameEventFileDownloaded) error) {
	if ge.onFileDownloaded == nil {
		ge.onFileDownloaded = make([]func(*GameEventFileDownloaded) error, 0)
	}
	ge.onFileDownloaded = append(ge.onFileDownloaded, fn)
}

func (ge *GameEvents) OnPlayerReportCountsUpdated(fn func(*GameEventPlayerReportCountsUpdated) error) {
	if ge.onPlayerReportCountsUpdated == nil {
		ge.onPlayerReportCountsUpdated = make([]func(*GameEventPlayerReportCountsUpdated) error, 0)
	}
	ge.onPlayerReportCountsUpdated = append(ge.onPlayerReportCountsUpdated, fn)
}

func (ge *GameEvents) OnScaleformFileDownloadComplete(fn func(*GameEventScaleformFileDownloadComplete) error) {
	if ge.onScaleformFileDownloadComplete == nil {
		ge.onScaleformFileDownloadComplete = make([]func(*GameEventScaleformFileDownloadComplete) error, 0)
	}
	ge.onScaleformFileDownloadComplete = append(ge.onScaleformFileDownloadComplete, fn)
}

func (ge *GameEvents) OnItemPurchased(fn func(*GameEventItemPurchased) error) {
	if ge.onItemPurchased == nil {
		ge.onItemPurchased = make([]func(*GameEventItemPurchased) error, 0)
	}
	ge.onItemPurchased = append(ge.onItemPurchased, fn)
}

func (ge *GameEvents) OnGcMismatchedVersion(fn func(*GameEventGcMismatchedVersion) error) {
	if ge.onGcMismatchedVersion == nil {
		ge.onGcMismatchedVersion = make([]func(*GameEventGcMismatchedVersion) error, 0)
	}
	ge.onGcMismatchedVersion = append(ge.onGcMismatchedVersion, fn)
}

func (ge *GameEvents) OnDemoStop(fn func(*GameEventDemoStop) error) {
	if ge.onDemoStop == nil {
		ge.onDemoStop = make([]func(*GameEventDemoStop) error, 0)
	}
	ge.onDemoStop = append(ge.onDemoStop, fn)
}

func (ge *GameEvents) OnMapShutdown(fn func(*GameEventMapShutdown) error) {
	if ge.onMapShutdown == nil {
		ge.onMapShutdown = make([]func(*GameEventMapShutdown) error, 0)
	}
	ge.onMapShutdown = append(ge.onMapShutdown, fn)
}

func (ge *GameEvents) OnDotaWorkshopFileselected(fn func(*GameEventDotaWorkshopFileselected) error) {
	if ge.onDotaWorkshopFileselected == nil {
		ge.onDotaWorkshopFileselected = make([]func(*GameEventDotaWorkshopFileselected) error, 0)
	}
	ge.onDotaWorkshopFileselected = append(ge.onDotaWorkshopFileselected, fn)
}

func (ge *GameEvents) OnDotaWorkshopFilecanceled(fn func(*GameEventDotaWorkshopFilecanceled) error) {
	if ge.onDotaWorkshopFilecanceled == nil {
		ge.onDotaWorkshopFilecanceled = make([]func(*GameEventDotaWorkshopFilecanceled) error, 0)
	}
	ge.onDotaWorkshopFilecanceled = append(ge.onDotaWorkshopFilecanceled, fn)
}

func (ge *GameEvents) OnRichPresenceUpdated(fn func(*GameEventRichPresenceUpdated) error) {
	if ge.onRichPresenceUpdated == nil {
		ge.onRichPresenceUpdated = make([]func(*GameEventRichPresenceUpdated) error, 0)
	}
	ge.onRichPresenceUpdated = append(ge.onRichPresenceUpdated, fn)
}

func (ge *GameEvents) OnDotaHeroRandom(fn func(*GameEventDotaHeroRandom) error) {
	if ge.onDotaHeroRandom == nil {
		ge.onDotaHeroRandom = make([]func(*GameEventDotaHeroRandom) error, 0)
	}
	ge.onDotaHeroRandom = append(ge.onDotaHeroRandom, fn)
}

func (ge *GameEvents) OnDotaRdChatTurn(fn func(*GameEventDotaRdChatTurn) error) {
	if ge.onDotaRdChatTurn == nil {
		ge.onDotaRdChatTurn = make([]func(*GameEventDotaRdChatTurn) error, 0)
	}
	ge.onDotaRdChatTurn = append(ge.onDotaRdChatTurn, fn)
}

func (ge *GameEvents) OnDotaFavoriteHeroesUpdated(fn func(*GameEventDotaFavoriteHeroesUpdated) error) {
	if ge.onDotaFavoriteHeroesUpdated == nil {
		ge.onDotaFavoriteHeroesUpdated = make([]func(*GameEventDotaFavoriteHeroesUpdated) error, 0)
	}
	ge.onDotaFavoriteHeroesUpdated = append(ge.onDotaFavoriteHeroesUpdated, fn)
}

func (ge *GameEvents) OnProfileOpened(fn func(*GameEventProfileOpened) error) {
	if ge.onProfileOpened == nil {
		ge.onProfileOpened = make([]func(*GameEventProfileOpened) error, 0)
	}
	ge.onProfileOpened = append(ge.onProfileOpened, fn)
}

func (ge *GameEvents) OnProfileClosed(fn func(*GameEventProfileClosed) error) {
	if ge.onProfileClosed == nil {
		ge.onProfileClosed = make([]func(*GameEventProfileClosed) error, 0)
	}
	ge.onProfileClosed = append(ge.onProfileClosed, fn)
}

func (ge *GameEvents) OnItemPreviewClosed(fn func(*GameEventItemPreviewClosed) error) {
	if ge.onItemPreviewClosed == nil {
		ge.onItemPreviewClosed = make([]func(*GameEventItemPreviewClosed) error, 0)
	}
	ge.onItemPreviewClosed = append(ge.onItemPreviewClosed, fn)
}

func (ge *GameEvents) OnDashboardSwitchedSection(fn func(*GameEventDashboardSwitchedSection) error) {
	if ge.onDashboardSwitchedSection == nil {
		ge.onDashboardSwitchedSection = make([]func(*GameEventDashboardSwitchedSection) error, 0)
	}
	ge.onDashboardSwitchedSection = append(ge.onDashboardSwitchedSection, fn)
}

func (ge *GameEvents) OnDotaTournamentItemEvent(fn func(*GameEventDotaTournamentItemEvent) error) {
	if ge.onDotaTournamentItemEvent == nil {
		ge.onDotaTournamentItemEvent = make([]func(*GameEventDotaTournamentItemEvent) error, 0)
	}
	ge.onDotaTournamentItemEvent = append(ge.onDotaTournamentItemEvent, fn)
}

func (ge *GameEvents) OnDotaHeroSwap(fn func(*GameEventDotaHeroSwap) error) {
	if ge.onDotaHeroSwap == nil {
		ge.onDotaHeroSwap = make([]func(*GameEventDotaHeroSwap) error, 0)
	}
	ge.onDotaHeroSwap = append(ge.onDotaHeroSwap, fn)
}

func (ge *GameEvents) OnDotaResetSuggestedItems(fn func(*GameEventDotaResetSuggestedItems) error) {
	if ge.onDotaResetSuggestedItems == nil {
		ge.onDotaResetSuggestedItems = make([]func(*GameEventDotaResetSuggestedItems) error, 0)
	}
	ge.onDotaResetSuggestedItems = append(ge.onDotaResetSuggestedItems, fn)
}

func (ge *GameEvents) OnHalloweenHighScoreReceived(fn func(*GameEventHalloweenHighScoreReceived) error) {
	if ge.onHalloweenHighScoreReceived == nil {
		ge.onHalloweenHighScoreReceived = make([]func(*GameEventHalloweenHighScoreReceived) error, 0)
	}
	ge.onHalloweenHighScoreReceived = append(ge.onHalloweenHighScoreReceived, fn)
}

func (ge *GameEvents) OnHalloweenPhaseEnd(fn func(*GameEventHalloweenPhaseEnd) error) {
	if ge.onHalloweenPhaseEnd == nil {
		ge.onHalloweenPhaseEnd = make([]func(*GameEventHalloweenPhaseEnd) error, 0)
	}
	ge.onHalloweenPhaseEnd = append(ge.onHalloweenPhaseEnd, fn)
}

func (ge *GameEvents) OnHalloweenHighScoreRequestFailed(fn func(*GameEventHalloweenHighScoreRequestFailed) error) {
	if ge.onHalloweenHighScoreRequestFailed == nil {
		ge.onHalloweenHighScoreRequestFailed = make([]func(*GameEventHalloweenHighScoreRequestFailed) error, 0)
	}
	ge.onHalloweenHighScoreRequestFailed = append(ge.onHalloweenHighScoreRequestFailed, fn)
}

func (ge *GameEvents) OnDotaHudSkinChanged(fn func(*GameEventDotaHudSkinChanged) error) {
	if ge.onDotaHudSkinChanged == nil {
		ge.onDotaHudSkinChanged = make([]func(*GameEventDotaHudSkinChanged) error, 0)
	}
	ge.onDotaHudSkinChanged = append(ge.onDotaHudSkinChanged, fn)
}

func (ge *GameEvents) OnDotaInventoryPlayerGotItem(fn func(*GameEventDotaInventoryPlayerGotItem) error) {
	if ge.onDotaInventoryPlayerGotItem == nil {
		ge.onDotaInventoryPlayerGotItem = make([]func(*GameEventDotaInventoryPlayerGotItem) error, 0)
	}
	ge.onDotaInventoryPlayerGotItem = append(ge.onDotaInventoryPlayerGotItem, fn)
}

func (ge *GameEvents) OnPlayerIsExperienced(fn func(*GameEventPlayerIsExperienced) error) {
	if ge.onPlayerIsExperienced == nil {
		ge.onPlayerIsExperienced = make([]func(*GameEventPlayerIsExperienced) error, 0)
	}
	ge.onPlayerIsExperienced = append(ge.onPlayerIsExperienced, fn)
}

func (ge *GameEvents) OnPlayerIsNotexperienced(fn func(*GameEventPlayerIsNotexperienced) error) {
	if ge.onPlayerIsNotexperienced == nil {
		ge.onPlayerIsNotexperienced = make([]func(*GameEventPlayerIsNotexperienced) error, 0)
	}
	ge.onPlayerIsNotexperienced = append(ge.onPlayerIsNotexperienced, fn)
}

func (ge *GameEvents) OnDotaTutorialLessonStart(fn func(*GameEventDotaTutorialLessonStart) error) {
	if ge.onDotaTutorialLessonStart == nil {
		ge.onDotaTutorialLessonStart = make([]func(*GameEventDotaTutorialLessonStart) error, 0)
	}
	ge.onDotaTutorialLessonStart = append(ge.onDotaTutorialLessonStart, fn)
}

func (ge *GameEvents) OnDotaTutorialTaskAdvance(fn func(*GameEventDotaTutorialTaskAdvance) error) {
	if ge.onDotaTutorialTaskAdvance == nil {
		ge.onDotaTutorialTaskAdvance = make([]func(*GameEventDotaTutorialTaskAdvance) error, 0)
	}
	ge.onDotaTutorialTaskAdvance = append(ge.onDotaTutorialTaskAdvance, fn)
}

func (ge *GameEvents) OnDotaTutorialShopToggled(fn func(*GameEventDotaTutorialShopToggled) error) {
	if ge.onDotaTutorialShopToggled == nil {
		ge.onDotaTutorialShopToggled = make([]func(*GameEventDotaTutorialShopToggled) error, 0)
	}
	ge.onDotaTutorialShopToggled = append(ge.onDotaTutorialShopToggled, fn)
}

func (ge *GameEvents) OnMapLocationUpdated(fn func(*GameEventMapLocationUpdated) error) {
	if ge.onMapLocationUpdated == nil {
		ge.onMapLocationUpdated = make([]func(*GameEventMapLocationUpdated) error, 0)
	}
	ge.onMapLocationUpdated = append(ge.onMapLocationUpdated, fn)
}

func (ge *GameEvents) OnRichpresenceCustomUpdated(fn func(*GameEventRichpresenceCustomUpdated) error) {
	if ge.onRichpresenceCustomUpdated == nil {
		ge.onRichpresenceCustomUpdated = make([]func(*GameEventRichpresenceCustomUpdated) error, 0)
	}
	ge.onRichpresenceCustomUpdated = append(ge.onRichpresenceCustomUpdated, fn)
}

func (ge *GameEvents) OnGameEndVisible(fn func(*GameEventGameEndVisible) error) {
	if ge.onGameEndVisible == nil {
		ge.onGameEndVisible = make([]func(*GameEventGameEndVisible) error, 0)
	}
	ge.onGameEndVisible = append(ge.onGameEndVisible, fn)
}

func (ge *GameEvents) OnAntiaddictionUpdate(fn func(*GameEventAntiaddictionUpdate) error) {
	if ge.onAntiaddictionUpdate == nil {
		ge.onAntiaddictionUpdate = make([]func(*GameEventAntiaddictionUpdate) error, 0)
	}
	ge.onAntiaddictionUpdate = append(ge.onAntiaddictionUpdate, fn)
}

func (ge *GameEvents) OnHighlightHudElement(fn func(*GameEventHighlightHudElement) error) {
	if ge.onHighlightHudElement == nil {
		ge.onHighlightHudElement = make([]func(*GameEventHighlightHudElement) error, 0)
	}
	ge.onHighlightHudElement = append(ge.onHighlightHudElement, fn)
}

func (ge *GameEvents) OnHideHighlightHudElement(fn func(*GameEventHideHighlightHudElement) error) {
	if ge.onHideHighlightHudElement == nil {
		ge.onHideHighlightHudElement = make([]func(*GameEventHideHighlightHudElement) error, 0)
	}
	ge.onHideHighlightHudElement = append(ge.onHideHighlightHudElement, fn)
}

func (ge *GameEvents) OnIntroVideoFinished(fn func(*GameEventIntroVideoFinished) error) {
	if ge.onIntroVideoFinished == nil {
		ge.onIntroVideoFinished = make([]func(*GameEventIntroVideoFinished) error, 0)
	}
	ge.onIntroVideoFinished = append(ge.onIntroVideoFinished, fn)
}

func (ge *GameEvents) OnMatchmakingStatusVisibilityChanged(fn func(*GameEventMatchmakingStatusVisibilityChanged) error) {
	if ge.onMatchmakingStatusVisibilityChanged == nil {
		ge.onMatchmakingStatusVisibilityChanged = make([]func(*GameEventMatchmakingStatusVisibilityChanged) error, 0)
	}
	ge.onMatchmakingStatusVisibilityChanged = append(ge.onMatchmakingStatusVisibilityChanged, fn)
}

func (ge *GameEvents) OnPracticeLobbyVisibilityChanged(fn func(*GameEventPracticeLobbyVisibilityChanged) error) {
	if ge.onPracticeLobbyVisibilityChanged == nil {
		ge.onPracticeLobbyVisibilityChanged = make([]func(*GameEventPracticeLobbyVisibilityChanged) error, 0)
	}
	ge.onPracticeLobbyVisibilityChanged = append(ge.onPracticeLobbyVisibilityChanged, fn)
}

func (ge *GameEvents) OnDotaCourierTransferItem(fn func(*GameEventDotaCourierTransferItem) error) {
	if ge.onDotaCourierTransferItem == nil {
		ge.onDotaCourierTransferItem = make([]func(*GameEventDotaCourierTransferItem) error, 0)
	}
	ge.onDotaCourierTransferItem = append(ge.onDotaCourierTransferItem, fn)
}

func (ge *GameEvents) OnFullUiUnlocked(fn func(*GameEventFullUiUnlocked) error) {
	if ge.onFullUiUnlocked == nil {
		ge.onFullUiUnlocked = make([]func(*GameEventFullUiUnlocked) error, 0)
	}
	ge.onFullUiUnlocked = append(ge.onFullUiUnlocked, fn)
}

func (ge *GameEvents) OnHeroSelectorPreviewSet(fn func(*GameEventHeroSelectorPreviewSet) error) {
	if ge.onHeroSelectorPreviewSet == nil {
		ge.onHeroSelectorPreviewSet = make([]func(*GameEventHeroSelectorPreviewSet) error, 0)
	}
	ge.onHeroSelectorPreviewSet = append(ge.onHeroSelectorPreviewSet, fn)
}

func (ge *GameEvents) OnAntiaddictionToast(fn func(*GameEventAntiaddictionToast) error) {
	if ge.onAntiaddictionToast == nil {
		ge.onAntiaddictionToast = make([]func(*GameEventAntiaddictionToast) error, 0)
	}
	ge.onAntiaddictionToast = append(ge.onAntiaddictionToast, fn)
}

func (ge *GameEvents) OnHeroPickerShown(fn func(*GameEventHeroPickerShown) error) {
	if ge.onHeroPickerShown == nil {
		ge.onHeroPickerShown = make([]func(*GameEventHeroPickerShown) error, 0)
	}
	ge.onHeroPickerShown = append(ge.onHeroPickerShown, fn)
}

func (ge *GameEvents) OnHeroPickerHidden(fn func(*GameEventHeroPickerHidden) error) {
	if ge.onHeroPickerHidden == nil {
		ge.onHeroPickerHidden = make([]func(*GameEventHeroPickerHidden) error, 0)
	}
	ge.onHeroPickerHidden = append(ge.onHeroPickerHidden, fn)
}

func (ge *GameEvents) OnDotaLocalQuickbuyChanged(fn func(*GameEventDotaLocalQuickbuyChanged) error) {
	if ge.onDotaLocalQuickbuyChanged == nil {
		ge.onDotaLocalQuickbuyChanged = make([]func(*GameEventDotaLocalQuickbuyChanged) error, 0)
	}
	ge.onDotaLocalQuickbuyChanged = append(ge.onDotaLocalQuickbuyChanged, fn)
}

func (ge *GameEvents) OnShowCenterMessage(fn func(*GameEventShowCenterMessage) error) {
	if ge.onShowCenterMessage == nil {
		ge.onShowCenterMessage = make([]func(*GameEventShowCenterMessage) error, 0)
	}
	ge.onShowCenterMessage = append(ge.onShowCenterMessage, fn)
}

func (ge *GameEvents) OnHudFlipChanged(fn func(*GameEventHudFlipChanged) error) {
	if ge.onHudFlipChanged == nil {
		ge.onHudFlipChanged = make([]func(*GameEventHudFlipChanged) error, 0)
	}
	ge.onHudFlipChanged = append(ge.onHudFlipChanged, fn)
}

func (ge *GameEvents) OnFrostyPointsUpdated(fn func(*GameEventFrostyPointsUpdated) error) {
	if ge.onFrostyPointsUpdated == nil {
		ge.onFrostyPointsUpdated = make([]func(*GameEventFrostyPointsUpdated) error, 0)
	}
	ge.onFrostyPointsUpdated = append(ge.onFrostyPointsUpdated, fn)
}

func (ge *GameEvents) OnDefeated(fn func(*GameEventDefeated) error) {
	if ge.onDefeated == nil {
		ge.onDefeated = make([]func(*GameEventDefeated) error, 0)
	}
	ge.onDefeated = append(ge.onDefeated, fn)
}

func (ge *GameEvents) OnResetDefeated(fn func(*GameEventResetDefeated) error) {
	if ge.onResetDefeated == nil {
		ge.onResetDefeated = make([]func(*GameEventResetDefeated) error, 0)
	}
	ge.onResetDefeated = append(ge.onResetDefeated, fn)
}

func (ge *GameEvents) OnBoosterStateUpdated(fn func(*GameEventBoosterStateUpdated) error) {
	if ge.onBoosterStateUpdated == nil {
		ge.onBoosterStateUpdated = make([]func(*GameEventBoosterStateUpdated) error, 0)
	}
	ge.onBoosterStateUpdated = append(ge.onBoosterStateUpdated, fn)
}

func (ge *GameEvents) OnEventPointsUpdated(fn func(*GameEventEventPointsUpdated) error) {
	if ge.onEventPointsUpdated == nil {
		ge.onEventPointsUpdated = make([]func(*GameEventEventPointsUpdated) error, 0)
	}
	ge.onEventPointsUpdated = append(ge.onEventPointsUpdated, fn)
}

func (ge *GameEvents) OnLocalPlayerEventPoints(fn func(*GameEventLocalPlayerEventPoints) error) {
	if ge.onLocalPlayerEventPoints == nil {
		ge.onLocalPlayerEventPoints = make([]func(*GameEventLocalPlayerEventPoints) error, 0)
	}
	ge.onLocalPlayerEventPoints = append(ge.onLocalPlayerEventPoints, fn)
}

func (ge *GameEvents) OnCustomGameDifficulty(fn func(*GameEventCustomGameDifficulty) error) {
	if ge.onCustomGameDifficulty == nil {
		ge.onCustomGameDifficulty = make([]func(*GameEventCustomGameDifficulty) error, 0)
	}
	ge.onCustomGameDifficulty = append(ge.onCustomGameDifficulty, fn)
}

func (ge *GameEvents) OnTreeCut(fn func(*GameEventTreeCut) error) {
	if ge.onTreeCut == nil {
		ge.onTreeCut = make([]func(*GameEventTreeCut) error, 0)
	}
	ge.onTreeCut = append(ge.onTreeCut, fn)
}

func (ge *GameEvents) OnUgcDetailsArrived(fn func(*GameEventUgcDetailsArrived) error) {
	if ge.onUgcDetailsArrived == nil {
		ge.onUgcDetailsArrived = make([]func(*GameEventUgcDetailsArrived) error, 0)
	}
	ge.onUgcDetailsArrived = append(ge.onUgcDetailsArrived, fn)
}

func (ge *GameEvents) OnUgcSubscribed(fn func(*GameEventUgcSubscribed) error) {
	if ge.onUgcSubscribed == nil {
		ge.onUgcSubscribed = make([]func(*GameEventUgcSubscribed) error, 0)
	}
	ge.onUgcSubscribed = append(ge.onUgcSubscribed, fn)
}

func (ge *GameEvents) OnUgcUnsubscribed(fn func(*GameEventUgcUnsubscribed) error) {
	if ge.onUgcUnsubscribed == nil {
		ge.onUgcUnsubscribed = make([]func(*GameEventUgcUnsubscribed) error, 0)
	}
	ge.onUgcUnsubscribed = append(ge.onUgcUnsubscribed, fn)
}

func (ge *GameEvents) OnUgcDownloadRequested(fn func(*GameEventUgcDownloadRequested) error) {
	if ge.onUgcDownloadRequested == nil {
		ge.onUgcDownloadRequested = make([]func(*GameEventUgcDownloadRequested) error, 0)
	}
	ge.onUgcDownloadRequested = append(ge.onUgcDownloadRequested, fn)
}

func (ge *GameEvents) OnUgcInstalled(fn func(*GameEventUgcInstalled) error) {
	if ge.onUgcInstalled == nil {
		ge.onUgcInstalled = make([]func(*GameEventUgcInstalled) error, 0)
	}
	ge.onUgcInstalled = append(ge.onUgcInstalled, fn)
}

func (ge *GameEvents) OnPrizepoolReceived(fn func(*GameEventPrizepoolReceived) error) {
	if ge.onPrizepoolReceived == nil {
		ge.onPrizepoolReceived = make([]func(*GameEventPrizepoolReceived) error, 0)
	}
	ge.onPrizepoolReceived = append(ge.onPrizepoolReceived, fn)
}

func (ge *GameEvents) OnMicrotransactionSuccess(fn func(*GameEventMicrotransactionSuccess) error) {
	if ge.onMicrotransactionSuccess == nil {
		ge.onMicrotransactionSuccess = make([]func(*GameEventMicrotransactionSuccess) error, 0)
	}
	ge.onMicrotransactionSuccess = append(ge.onMicrotransactionSuccess, fn)
}

func (ge *GameEvents) OnDotaRubickAbilitySteal(fn func(*GameEventDotaRubickAbilitySteal) error) {
	if ge.onDotaRubickAbilitySteal == nil {
		ge.onDotaRubickAbilitySteal = make([]func(*GameEventDotaRubickAbilitySteal) error, 0)
	}
	ge.onDotaRubickAbilitySteal = append(ge.onDotaRubickAbilitySteal, fn)
}

func (ge *GameEvents) OnCompendiumEventActionsLoaded(fn func(*GameEventCompendiumEventActionsLoaded) error) {
	if ge.onCompendiumEventActionsLoaded == nil {
		ge.onCompendiumEventActionsLoaded = make([]func(*GameEventCompendiumEventActionsLoaded) error, 0)
	}
	ge.onCompendiumEventActionsLoaded = append(ge.onCompendiumEventActionsLoaded, fn)
}

func (ge *GameEvents) OnCompendiumSelectionsLoaded(fn func(*GameEventCompendiumSelectionsLoaded) error) {
	if ge.onCompendiumSelectionsLoaded == nil {
		ge.onCompendiumSelectionsLoaded = make([]func(*GameEventCompendiumSelectionsLoaded) error, 0)
	}
	ge.onCompendiumSelectionsLoaded = append(ge.onCompendiumSelectionsLoaded, fn)
}

func (ge *GameEvents) OnCompendiumSetSelectionFailed(fn func(*GameEventCompendiumSetSelectionFailed) error) {
	if ge.onCompendiumSetSelectionFailed == nil {
		ge.onCompendiumSetSelectionFailed = make([]func(*GameEventCompendiumSetSelectionFailed) error, 0)
	}
	ge.onCompendiumSetSelectionFailed = append(ge.onCompendiumSetSelectionFailed, fn)
}

func (ge *GameEvents) OnCompendiumTrophiesLoaded(fn func(*GameEventCompendiumTrophiesLoaded) error) {
	if ge.onCompendiumTrophiesLoaded == nil {
		ge.onCompendiumTrophiesLoaded = make([]func(*GameEventCompendiumTrophiesLoaded) error, 0)
	}
	ge.onCompendiumTrophiesLoaded = append(ge.onCompendiumTrophiesLoaded, fn)
}

func (ge *GameEvents) OnCommunityCachedNamesUpdated(fn func(*GameEventCommunityCachedNamesUpdated) error) {
	if ge.onCommunityCachedNamesUpdated == nil {
		ge.onCommunityCachedNamesUpdated = make([]func(*GameEventCommunityCachedNamesUpdated) error, 0)
	}
	ge.onCommunityCachedNamesUpdated = append(ge.onCommunityCachedNamesUpdated, fn)
}

func (ge *GameEvents) OnSpecItemPickup(fn func(*GameEventSpecItemPickup) error) {
	if ge.onSpecItemPickup == nil {
		ge.onSpecItemPickup = make([]func(*GameEventSpecItemPickup) error, 0)
	}
	ge.onSpecItemPickup = append(ge.onSpecItemPickup, fn)
}

func (ge *GameEvents) OnSpecAegisReclaimTime(fn func(*GameEventSpecAegisReclaimTime) error) {
	if ge.onSpecAegisReclaimTime == nil {
		ge.onSpecAegisReclaimTime = make([]func(*GameEventSpecAegisReclaimTime) error, 0)
	}
	ge.onSpecAegisReclaimTime = append(ge.onSpecAegisReclaimTime, fn)
}

func (ge *GameEvents) OnAccountTrophiesChanged(fn func(*GameEventAccountTrophiesChanged) error) {
	if ge.onAccountTrophiesChanged == nil {
		ge.onAccountTrophiesChanged = make([]func(*GameEventAccountTrophiesChanged) error, 0)
	}
	ge.onAccountTrophiesChanged = append(ge.onAccountTrophiesChanged, fn)
}

func (ge *GameEvents) OnAccountAllHeroChallengeChanged(fn func(*GameEventAccountAllHeroChallengeChanged) error) {
	if ge.onAccountAllHeroChallengeChanged == nil {
		ge.onAccountAllHeroChallengeChanged = make([]func(*GameEventAccountAllHeroChallengeChanged) error, 0)
	}
	ge.onAccountAllHeroChallengeChanged = append(ge.onAccountAllHeroChallengeChanged, fn)
}

func (ge *GameEvents) OnTeamShowcaseUiUpdate(fn func(*GameEventTeamShowcaseUiUpdate) error) {
	if ge.onTeamShowcaseUiUpdate == nil {
		ge.onTeamShowcaseUiUpdate = make([]func(*GameEventTeamShowcaseUiUpdate) error, 0)
	}
	ge.onTeamShowcaseUiUpdate = append(ge.onTeamShowcaseUiUpdate, fn)
}

func (ge *GameEvents) OnIngameEventsChanged(fn func(*GameEventIngameEventsChanged) error) {
	if ge.onIngameEventsChanged == nil {
		ge.onIngameEventsChanged = make([]func(*GameEventIngameEventsChanged) error, 0)
	}
	ge.onIngameEventsChanged = append(ge.onIngameEventsChanged, fn)
}

func (ge *GameEvents) OnDotaMatchSignout(fn func(*GameEventDotaMatchSignout) error) {
	if ge.onDotaMatchSignout == nil {
		ge.onDotaMatchSignout = make([]func(*GameEventDotaMatchSignout) error, 0)
	}
	ge.onDotaMatchSignout = append(ge.onDotaMatchSignout, fn)
}

func (ge *GameEvents) OnDotaIllusionsCreated(fn func(*GameEventDotaIllusionsCreated) error) {
	if ge.onDotaIllusionsCreated == nil {
		ge.onDotaIllusionsCreated = make([]func(*GameEventDotaIllusionsCreated) error, 0)
	}
	ge.onDotaIllusionsCreated = append(ge.onDotaIllusionsCreated, fn)
}

func (ge *GameEvents) OnDotaYearBeastKilled(fn func(*GameEventDotaYearBeastKilled) error) {
	if ge.onDotaYearBeastKilled == nil {
		ge.onDotaYearBeastKilled = make([]func(*GameEventDotaYearBeastKilled) error, 0)
	}
	ge.onDotaYearBeastKilled = append(ge.onDotaYearBeastKilled, fn)
}

func (ge *GameEvents) OnDotaHeroUndoselection(fn func(*GameEventDotaHeroUndoselection) error) {
	if ge.onDotaHeroUndoselection == nil {
		ge.onDotaHeroUndoselection = make([]func(*GameEventDotaHeroUndoselection) error, 0)
	}
	ge.onDotaHeroUndoselection = append(ge.onDotaHeroUndoselection, fn)
}

func (ge *GameEvents) OnDotaChallengeSocacheUpdated(fn func(*GameEventDotaChallengeSocacheUpdated) error) {
	if ge.onDotaChallengeSocacheUpdated == nil {
		ge.onDotaChallengeSocacheUpdated = make([]func(*GameEventDotaChallengeSocacheUpdated) error, 0)
	}
	ge.onDotaChallengeSocacheUpdated = append(ge.onDotaChallengeSocacheUpdated, fn)
}

func (ge *GameEvents) OnPartyInvitesUpdated(fn func(*GameEventPartyInvitesUpdated) error) {
	if ge.onPartyInvitesUpdated == nil {
		ge.onPartyInvitesUpdated = make([]func(*GameEventPartyInvitesUpdated) error, 0)
	}
	ge.onPartyInvitesUpdated = append(ge.onPartyInvitesUpdated, fn)
}

func (ge *GameEvents) OnLobbyInvitesUpdated(fn func(*GameEventLobbyInvitesUpdated) error) {
	if ge.onLobbyInvitesUpdated == nil {
		ge.onLobbyInvitesUpdated = make([]func(*GameEventLobbyInvitesUpdated) error, 0)
	}
	ge.onLobbyInvitesUpdated = append(ge.onLobbyInvitesUpdated, fn)
}

func (ge *GameEvents) OnCustomGameModeListUpdated(fn func(*GameEventCustomGameModeListUpdated) error) {
	if ge.onCustomGameModeListUpdated == nil {
		ge.onCustomGameModeListUpdated = make([]func(*GameEventCustomGameModeListUpdated) error, 0)
	}
	ge.onCustomGameModeListUpdated = append(ge.onCustomGameModeListUpdated, fn)
}

func (ge *GameEvents) OnCustomGameLobbyListUpdated(fn func(*GameEventCustomGameLobbyListUpdated) error) {
	if ge.onCustomGameLobbyListUpdated == nil {
		ge.onCustomGameLobbyListUpdated = make([]func(*GameEventCustomGameLobbyListUpdated) error, 0)
	}
	ge.onCustomGameLobbyListUpdated = append(ge.onCustomGameLobbyListUpdated, fn)
}

func (ge *GameEvents) OnFriendLobbyListUpdated(fn func(*GameEventFriendLobbyListUpdated) error) {
	if ge.onFriendLobbyListUpdated == nil {
		ge.onFriendLobbyListUpdated = make([]func(*GameEventFriendLobbyListUpdated) error, 0)
	}
	ge.onFriendLobbyListUpdated = append(ge.onFriendLobbyListUpdated, fn)
}

func (ge *GameEvents) OnDotaTeamPlayerListChanged(fn func(*GameEventDotaTeamPlayerListChanged) error) {
	if ge.onDotaTeamPlayerListChanged == nil {
		ge.onDotaTeamPlayerListChanged = make([]func(*GameEventDotaTeamPlayerListChanged) error, 0)
	}
	ge.onDotaTeamPlayerListChanged = append(ge.onDotaTeamPlayerListChanged, fn)
}

func (ge *GameEvents) OnDotaPlayerDetailsChanged(fn func(*GameEventDotaPlayerDetailsChanged) error) {
	if ge.onDotaPlayerDetailsChanged == nil {
		ge.onDotaPlayerDetailsChanged = make([]func(*GameEventDotaPlayerDetailsChanged) error, 0)
	}
	ge.onDotaPlayerDetailsChanged = append(ge.onDotaPlayerDetailsChanged, fn)
}

func (ge *GameEvents) OnPlayerProfileStatsUpdated(fn func(*GameEventPlayerProfileStatsUpdated) error) {
	if ge.onPlayerProfileStatsUpdated == nil {
		ge.onPlayerProfileStatsUpdated = make([]func(*GameEventPlayerProfileStatsUpdated) error, 0)
	}
	ge.onPlayerProfileStatsUpdated = append(ge.onPlayerProfileStatsUpdated, fn)
}

func (ge *GameEvents) OnCustomGamePlayerCountUpdated(fn func(*GameEventCustomGamePlayerCountUpdated) error) {
	if ge.onCustomGamePlayerCountUpdated == nil {
		ge.onCustomGamePlayerCountUpdated = make([]func(*GameEventCustomGamePlayerCountUpdated) error, 0)
	}
	ge.onCustomGamePlayerCountUpdated = append(ge.onCustomGamePlayerCountUpdated, fn)
}

func (ge *GameEvents) OnCustomGameFriendsPlayedUpdated(fn func(*GameEventCustomGameFriendsPlayedUpdated) error) {
	if ge.onCustomGameFriendsPlayedUpdated == nil {
		ge.onCustomGameFriendsPlayedUpdated = make([]func(*GameEventCustomGameFriendsPlayedUpdated) error, 0)
	}
	ge.onCustomGameFriendsPlayedUpdated = append(ge.onCustomGameFriendsPlayedUpdated, fn)
}

func (ge *GameEvents) OnCustomGamesFriendsPlayUpdated(fn func(*GameEventCustomGamesFriendsPlayUpdated) error) {
	if ge.onCustomGamesFriendsPlayUpdated == nil {
		ge.onCustomGamesFriendsPlayUpdated = make([]func(*GameEventCustomGamesFriendsPlayUpdated) error, 0)
	}
	ge.onCustomGamesFriendsPlayUpdated = append(ge.onCustomGamesFriendsPlayUpdated, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateAssignedHero(fn func(*GameEventDotaPlayerUpdateAssignedHero) error) {
	if ge.onDotaPlayerUpdateAssignedHero == nil {
		ge.onDotaPlayerUpdateAssignedHero = make([]func(*GameEventDotaPlayerUpdateAssignedHero) error, 0)
	}
	ge.onDotaPlayerUpdateAssignedHero = append(ge.onDotaPlayerUpdateAssignedHero, fn)
}

func (ge *GameEvents) OnDotaPlayerHeroSelectionDirty(fn func(*GameEventDotaPlayerHeroSelectionDirty) error) {
	if ge.onDotaPlayerHeroSelectionDirty == nil {
		ge.onDotaPlayerHeroSelectionDirty = make([]func(*GameEventDotaPlayerHeroSelectionDirty) error, 0)
	}
	ge.onDotaPlayerHeroSelectionDirty = append(ge.onDotaPlayerHeroSelectionDirty, fn)
}

func (ge *GameEvents) OnDotaNpcGoalReached(fn func(*GameEventDotaNpcGoalReached) error) {
	if ge.onDotaNpcGoalReached == nil {
		ge.onDotaNpcGoalReached = make([]func(*GameEventDotaNpcGoalReached) error, 0)
	}
	ge.onDotaNpcGoalReached = append(ge.onDotaNpcGoalReached, fn)
}

func (ge *GameEvents) OnHltvStatus(fn func(*GameEventHltvStatus) error) {
	if ge.onHltvStatus == nil {
		ge.onHltvStatus = make([]func(*GameEventHltvStatus) error, 0)
	}
	ge.onHltvStatus = append(ge.onHltvStatus, fn)
}

func (ge *GameEvents) OnHltvCameraman(fn func(*GameEventHltvCameraman) error) {
	if ge.onHltvCameraman == nil {
		ge.onHltvCameraman = make([]func(*GameEventHltvCameraman) error, 0)
	}
	ge.onHltvCameraman = append(ge.onHltvCameraman, fn)
}

func (ge *GameEvents) OnHltvRankCamera(fn func(*GameEventHltvRankCamera) error) {
	if ge.onHltvRankCamera == nil {
		ge.onHltvRankCamera = make([]func(*GameEventHltvRankCamera) error, 0)
	}
	ge.onHltvRankCamera = append(ge.onHltvRankCamera, fn)
}

func (ge *GameEvents) OnHltvRankEntity(fn func(*GameEventHltvRankEntity) error) {
	if ge.onHltvRankEntity == nil {
		ge.onHltvRankEntity = make([]func(*GameEventHltvRankEntity) error, 0)
	}
	ge.onHltvRankEntity = append(ge.onHltvRankEntity, fn)
}

func (ge *GameEvents) OnHltvFixed(fn func(*GameEventHltvFixed) error) {
	if ge.onHltvFixed == nil {
		ge.onHltvFixed = make([]func(*GameEventHltvFixed) error, 0)
	}
	ge.onHltvFixed = append(ge.onHltvFixed, fn)
}

func (ge *GameEvents) OnHltvChase(fn func(*GameEventHltvChase) error) {
	if ge.onHltvChase == nil {
		ge.onHltvChase = make([]func(*GameEventHltvChase) error, 0)
	}
	ge.onHltvChase = append(ge.onHltvChase, fn)
}

func (ge *GameEvents) OnHltvMessage(fn func(*GameEventHltvMessage) error) {
	if ge.onHltvMessage == nil {
		ge.onHltvMessage = make([]func(*GameEventHltvMessage) error, 0)
	}
	ge.onHltvMessage = append(ge.onHltvMessage, fn)
}

func (ge *GameEvents) OnHltvTitle(fn func(*GameEventHltvTitle) error) {
	if ge.onHltvTitle == nil {
		ge.onHltvTitle = make([]func(*GameEventHltvTitle) error, 0)
	}
	ge.onHltvTitle = append(ge.onHltvTitle, fn)
}

func (ge *GameEvents) OnHltvChat(fn func(*GameEventHltvChat) error) {
	if ge.onHltvChat == nil {
		ge.onHltvChat = make([]func(*GameEventHltvChat) error, 0)
	}
	ge.onHltvChat = append(ge.onHltvChat, fn)
}

func (ge *GameEvents) OnHltvVersioninfo(fn func(*GameEventHltvVersioninfo) error) {
	if ge.onHltvVersioninfo == nil {
		ge.onHltvVersioninfo = make([]func(*GameEventHltvVersioninfo) error, 0)
	}
	ge.onHltvVersioninfo = append(ge.onHltvVersioninfo, fn)
}

func (ge *GameEvents) OnDotaChaseHero(fn func(*GameEventDotaChaseHero) error) {
	if ge.onDotaChaseHero == nil {
		ge.onDotaChaseHero = make([]func(*GameEventDotaChaseHero) error, 0)
	}
	ge.onDotaChaseHero = append(ge.onDotaChaseHero, fn)
}

func (ge *GameEvents) OnDotaCombatlog(fn func(*GameEventDotaCombatlog) error) {
	if ge.onDotaCombatlog == nil {
		ge.onDotaCombatlog = make([]func(*GameEventDotaCombatlog) error, 0)
	}
	ge.onDotaCombatlog = append(ge.onDotaCombatlog, fn)
}

func (ge *GameEvents) OnDotaGameStateChange(fn func(*GameEventDotaGameStateChange) error) {
	if ge.onDotaGameStateChange == nil {
		ge.onDotaGameStateChange = make([]func(*GameEventDotaGameStateChange) error, 0)
	}
	ge.onDotaGameStateChange = append(ge.onDotaGameStateChange, fn)
}

func (ge *GameEvents) OnDotaPlayerPickHero(fn func(*GameEventDotaPlayerPickHero) error) {
	if ge.onDotaPlayerPickHero == nil {
		ge.onDotaPlayerPickHero = make([]func(*GameEventDotaPlayerPickHero) error, 0)
	}
	ge.onDotaPlayerPickHero = append(ge.onDotaPlayerPickHero, fn)
}

func (ge *GameEvents) OnDotaTeamKillCredit(fn func(*GameEventDotaTeamKillCredit) error) {
	if ge.onDotaTeamKillCredit == nil {
		ge.onDotaTeamKillCredit = make([]func(*GameEventDotaTeamKillCredit) error, 0)
	}
	ge.onDotaTeamKillCredit = append(ge.onDotaTeamKillCredit, fn)
}

func (ge *GameEvents) onCMsgSource1LegacyGameEvent(m *wireSource1GameEvent) error {
	switch m.GetEventid() {

	case 0: // EGameEvent_ServerSpawn
		if cbs := ge.onServerSpawn; cbs != nil {
			msg := &GameEventServerSpawn{}
			msg.Hostname = m.GetKeys()[0].GetValString()
			msg.Address = m.GetKeys()[1].GetValString()
			msg.Port = m.GetKeys()[2].GetValShort()
			msg.Game = m.GetKeys()[3].GetValString()
			msg.Mapname = m.GetKeys()[4].GetValString()
			msg.Addonname = m.GetKeys()[5].GetValString()
			msg.Maxplayers = m.GetKeys()[6].GetValLong()
			msg.Os = m.GetKeys()[7].GetValString()
			msg.Dedicated = m.GetKeys()[8].GetValBool()
			msg.Password = m.GetKeys()[9].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 1: // EGameEvent_ServerPreShutdown
		if cbs := ge.onServerPreShutdown; cbs != nil {
			msg := &GameEventServerPreShutdown{}
			msg.Reason = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 2: // EGameEvent_ServerShutdown
		if cbs := ge.onServerShutdown; cbs != nil {
			msg := &GameEventServerShutdown{}
			msg.Reason = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 3: // EGameEvent_ServerCvar
		if cbs := ge.onServerCvar; cbs != nil {
			msg := &GameEventServerCvar{}
			msg.Cvarname = m.GetKeys()[0].GetValString()
			msg.Cvarvalue = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 4: // EGameEvent_ServerMessage
		if cbs := ge.onServerMessage; cbs != nil {
			msg := &GameEventServerMessage{}
			msg.Text = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 5: // EGameEvent_ServerAddban
		if cbs := ge.onServerAddban; cbs != nil {
			msg := &GameEventServerAddban{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Userid = m.GetKeys()[1].GetValShort()
			msg.Networkid = m.GetKeys()[2].GetValString()
			msg.Ip = m.GetKeys()[3].GetValString()
			msg.Duration = m.GetKeys()[4].GetValString()
			msg.By = m.GetKeys()[5].GetValString()
			msg.Kicked = m.GetKeys()[6].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 6: // EGameEvent_ServerRemoveban
		if cbs := ge.onServerRemoveban; cbs != nil {
			msg := &GameEventServerRemoveban{}
			msg.Networkid = m.GetKeys()[0].GetValString()
			msg.Ip = m.GetKeys()[1].GetValString()
			msg.By = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 7: // EGameEvent_PlayerConnect
		if cbs := ge.onPlayerConnect; cbs != nil {
			msg := &GameEventPlayerConnect{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Index = m.GetKeys()[1].GetValByte()
			msg.Userid = m.GetKeys()[2].GetValShort()
			msg.Networkid = m.GetKeys()[3].GetValString()
			msg.Address = m.GetKeys()[4].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 8: // EGameEvent_PlayerInfo
		if cbs := ge.onPlayerInfo; cbs != nil {
			msg := &GameEventPlayerInfo{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Index = m.GetKeys()[1].GetValByte()
			msg.Userid = m.GetKeys()[2].GetValShort()
			msg.Networkid = m.GetKeys()[3].GetValString()
			msg.Bot = m.GetKeys()[4].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 9: // EGameEvent_PlayerDisconnect
		if cbs := ge.onPlayerDisconnect; cbs != nil {
			msg := &GameEventPlayerDisconnect{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Reason = m.GetKeys()[1].GetValShort()
			msg.Name = m.GetKeys()[2].GetValString()
			msg.Networkid = m.GetKeys()[3].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 10: // EGameEvent_PlayerActivate
		if cbs := ge.onPlayerActivate; cbs != nil {
			msg := &GameEventPlayerActivate{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 11: // EGameEvent_PlayerConnectFull
		if cbs := ge.onPlayerConnectFull; cbs != nil {
			msg := &GameEventPlayerConnectFull{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Index = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 12: // EGameEvent_PlayerSay
		if cbs := ge.onPlayerSay; cbs != nil {
			msg := &GameEventPlayerSay{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Text = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 13: // EGameEvent_PlayerFullUpdate
		if cbs := ge.onPlayerFullUpdate; cbs != nil {
			msg := &GameEventPlayerFullUpdate{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Count = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 14: // EGameEvent_TeamInfo
		if cbs := ge.onTeamInfo; cbs != nil {
			msg := &GameEventTeamInfo{}
			msg.Teamid = m.GetKeys()[0].GetValByte()
			msg.Teamname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 15: // EGameEvent_TeamScore
		if cbs := ge.onTeamScore; cbs != nil {
			msg := &GameEventTeamScore{}
			msg.Teamid = m.GetKeys()[0].GetValByte()
			msg.Score = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 16: // EGameEvent_TeamplayBroadcastAudio
		if cbs := ge.onTeamplayBroadcastAudio; cbs != nil {
			msg := &GameEventTeamplayBroadcastAudio{}
			msg.Team = m.GetKeys()[0].GetValByte()
			msg.Sound = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 17: // EGameEvent_PlayerTeam
		if cbs := ge.onPlayerTeam; cbs != nil {
			msg := &GameEventPlayerTeam{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Team = m.GetKeys()[1].GetValByte()
			msg.Oldteam = m.GetKeys()[2].GetValByte()
			msg.Disconnect = m.GetKeys()[3].GetValBool()
			msg.Autoteam = m.GetKeys()[4].GetValBool()
			msg.Silent = m.GetKeys()[5].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 18: // EGameEvent_PlayerClass
		if cbs := ge.onPlayerClass; cbs != nil {
			msg := &GameEventPlayerClass{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Class = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 19: // EGameEvent_PlayerDeath
		if cbs := ge.onPlayerDeath; cbs != nil {
			msg := &GameEventPlayerDeath{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Attacker = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 20: // EGameEvent_PlayerHurt
		if cbs := ge.onPlayerHurt; cbs != nil {
			msg := &GameEventPlayerHurt{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Attacker = m.GetKeys()[1].GetValShort()
			msg.Health = m.GetKeys()[2].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 21: // EGameEvent_PlayerChat
		if cbs := ge.onPlayerChat; cbs != nil {
			msg := &GameEventPlayerChat{}
			msg.Teamonly = m.GetKeys()[0].GetValBool()
			msg.Userid = m.GetKeys()[1].GetValShort()
			msg.Text = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 22: // EGameEvent_PlayerScore
		if cbs := ge.onPlayerScore; cbs != nil {
			msg := &GameEventPlayerScore{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Kills = m.GetKeys()[1].GetValShort()
			msg.Deaths = m.GetKeys()[2].GetValShort()
			msg.Score = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 23: // EGameEvent_PlayerSpawn
		if cbs := ge.onPlayerSpawn; cbs != nil {
			msg := &GameEventPlayerSpawn{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 24: // EGameEvent_PlayerShoot
		if cbs := ge.onPlayerShoot; cbs != nil {
			msg := &GameEventPlayerShoot{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Weapon = m.GetKeys()[1].GetValByte()
			msg.Mode = m.GetKeys()[2].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 25: // EGameEvent_PlayerUse
		if cbs := ge.onPlayerUse; cbs != nil {
			msg := &GameEventPlayerUse{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Entity = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 26: // EGameEvent_PlayerChangename
		if cbs := ge.onPlayerChangename; cbs != nil {
			msg := &GameEventPlayerChangename{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Oldname = m.GetKeys()[1].GetValString()
			msg.Newname = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 27: // EGameEvent_PlayerHintmessage
		if cbs := ge.onPlayerHintmessage; cbs != nil {
			msg := &GameEventPlayerHintmessage{}
			msg.Hintmessage = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 28: // EGameEvent_GameInit
		if cbs := ge.onGameInit; cbs != nil {
			msg := &GameEventGameInit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 29: // EGameEvent_GameNewmap
		if cbs := ge.onGameNewmap; cbs != nil {
			msg := &GameEventGameNewmap{}
			msg.Mapname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 30: // EGameEvent_GameStart
		if cbs := ge.onGameStart; cbs != nil {
			msg := &GameEventGameStart{}
			msg.Roundslimit = m.GetKeys()[0].GetValLong()
			msg.Timelimit = m.GetKeys()[1].GetValLong()
			msg.Fraglimit = m.GetKeys()[2].GetValLong()
			msg.Objective = m.GetKeys()[3].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 31: // EGameEvent_GameEnd
		if cbs := ge.onGameEnd; cbs != nil {
			msg := &GameEventGameEnd{}
			msg.Winner = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 32: // EGameEvent_RoundStart
		if cbs := ge.onRoundStart; cbs != nil {
			msg := &GameEventRoundStart{}
			msg.Timelimit = m.GetKeys()[0].GetValLong()
			msg.Fraglimit = m.GetKeys()[1].GetValLong()
			msg.Objective = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 33: // EGameEvent_RoundEnd
		if cbs := ge.onRoundEnd; cbs != nil {
			msg := &GameEventRoundEnd{}
			msg.Winner = m.GetKeys()[0].GetValByte()
			msg.Reason = m.GetKeys()[1].GetValByte()
			msg.Message = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 34: // EGameEvent_RoundStartPreEntity
		if cbs := ge.onRoundStartPreEntity; cbs != nil {
			msg := &GameEventRoundStartPreEntity{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 35: // EGameEvent_TeamplayRoundStart
		if cbs := ge.onTeamplayRoundStart; cbs != nil {
			msg := &GameEventTeamplayRoundStart{}
			msg.FullReset = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 36: // EGameEvent_HostnameChanged
		if cbs := ge.onHostnameChanged; cbs != nil {
			msg := &GameEventHostnameChanged{}
			msg.Hostname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 37: // EGameEvent_DifficultyChanged
		if cbs := ge.onDifficultyChanged; cbs != nil {
			msg := &GameEventDifficultyChanged{}
			msg.NewDifficulty = m.GetKeys()[0].GetValShort()
			msg.OldDifficulty = m.GetKeys()[1].GetValShort()
			msg.StrDifficulty = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 38: // EGameEvent_FinaleStart
		if cbs := ge.onFinaleStart; cbs != nil {
			msg := &GameEventFinaleStart{}
			msg.Rushes = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 39: // EGameEvent_GameMessage
		if cbs := ge.onGameMessage; cbs != nil {
			msg := &GameEventGameMessage{}
			msg.Target = m.GetKeys()[0].GetValByte()
			msg.Text = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 40: // EGameEvent_BreakBreakable
		if cbs := ge.onBreakBreakable; cbs != nil {
			msg := &GameEventBreakBreakable{}
			msg.Entindex = m.GetKeys()[0].GetValLong()
			msg.Userid = m.GetKeys()[1].GetValShort()
			msg.Material = m.GetKeys()[2].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 41: // EGameEvent_BreakProp
		if cbs := ge.onBreakProp; cbs != nil {
			msg := &GameEventBreakProp{}
			msg.Entindex = m.GetKeys()[0].GetValLong()
			msg.Userid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 42: // EGameEvent_NpcSpawned
		if cbs := ge.onNpcSpawned; cbs != nil {
			msg := &GameEventNpcSpawned{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 43: // EGameEvent_NpcReplaced
		if cbs := ge.onNpcReplaced; cbs != nil {
			msg := &GameEventNpcReplaced{}
			msg.OldEntindex = m.GetKeys()[0].GetValLong()
			msg.NewEntindex = m.GetKeys()[1].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 44: // EGameEvent_EntityKilled
		if cbs := ge.onEntityKilled; cbs != nil {
			msg := &GameEventEntityKilled{}
			msg.EntindexKilled = m.GetKeys()[0].GetValLong()
			msg.EntindexAttacker = m.GetKeys()[1].GetValLong()
			msg.EntindexInflictor = m.GetKeys()[2].GetValLong()
			msg.Damagebits = m.GetKeys()[3].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 45: // EGameEvent_EntityHurt
		if cbs := ge.onEntityHurt; cbs != nil {
			msg := &GameEventEntityHurt{}
			msg.EntindexKilled = m.GetKeys()[0].GetValLong()
			msg.EntindexAttacker = m.GetKeys()[1].GetValLong()
			msg.EntindexInflictor = m.GetKeys()[2].GetValLong()
			msg.Damagebits = m.GetKeys()[3].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 46: // EGameEvent_BonusUpdated
		if cbs := ge.onBonusUpdated; cbs != nil {
			msg := &GameEventBonusUpdated{}
			msg.Numadvanced = m.GetKeys()[0].GetValShort()
			msg.Numbronze = m.GetKeys()[1].GetValShort()
			msg.Numsilver = m.GetKeys()[2].GetValShort()
			msg.Numgold = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 47: // EGameEvent_PlayerStatsUpdated
		if cbs := ge.onPlayerStatsUpdated; cbs != nil {
			msg := &GameEventPlayerStatsUpdated{}
			msg.Forceupload = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 48: // EGameEvent_AchievementEvent
		if cbs := ge.onAchievementEvent; cbs != nil {
			msg := &GameEventAchievementEvent{}
			msg.AchievementName = m.GetKeys()[0].GetValString()
			msg.CurVal = m.GetKeys()[1].GetValShort()
			msg.MaxVal = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 49: // EGameEvent_AchievementEarned
		if cbs := ge.onAchievementEarned; cbs != nil {
			msg := &GameEventAchievementEarned{}
			msg.Player = m.GetKeys()[0].GetValByte()
			msg.Achievement = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 50: // EGameEvent_AchievementWriteFailed
		if cbs := ge.onAchievementWriteFailed; cbs != nil {
			msg := &GameEventAchievementWriteFailed{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 51: // EGameEvent_PhysgunPickup
		if cbs := ge.onPhysgunPickup; cbs != nil {
			msg := &GameEventPhysgunPickup{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 52: // EGameEvent_FlareIgniteNpc
		if cbs := ge.onFlareIgniteNpc; cbs != nil {
			msg := &GameEventFlareIgniteNpc{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 53: // EGameEvent_HelicopterGrenadePuntMiss
		if cbs := ge.onHelicopterGrenadePuntMiss; cbs != nil {
			msg := &GameEventHelicopterGrenadePuntMiss{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 54: // EGameEvent_UserDataDownloaded
		if cbs := ge.onUserDataDownloaded; cbs != nil {
			msg := &GameEventUserDataDownloaded{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 55: // EGameEvent_RagdollDissolved
		if cbs := ge.onRagdollDissolved; cbs != nil {
			msg := &GameEventRagdollDissolved{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 56: // EGameEvent_GameinstructorDraw
		if cbs := ge.onGameinstructorDraw; cbs != nil {
			msg := &GameEventGameinstructorDraw{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 57: // EGameEvent_GameinstructorNodraw
		if cbs := ge.onGameinstructorNodraw; cbs != nil {
			msg := &GameEventGameinstructorNodraw{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 58: // EGameEvent_MapTransition
		if cbs := ge.onMapTransition; cbs != nil {
			msg := &GameEventMapTransition{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 59: // EGameEvent_InstructorServerHintCreate
		if cbs := ge.onInstructorServerHintCreate; cbs != nil {
			msg := &GameEventInstructorServerHintCreate{}
			msg.HintName = m.GetKeys()[0].GetValString()
			msg.HintReplaceKey = m.GetKeys()[1].GetValString()
			msg.HintTarget = m.GetKeys()[2].GetValLong()
			msg.HintActivatorUserid = m.GetKeys()[3].GetValShort()
			msg.HintTimeout = m.GetKeys()[4].GetValShort()
			msg.HintIconOnscreen = m.GetKeys()[5].GetValString()
			msg.HintIconOffscreen = m.GetKeys()[6].GetValString()
			msg.HintCaption = m.GetKeys()[7].GetValString()
			msg.HintActivatorCaption = m.GetKeys()[8].GetValString()
			msg.HintColor = m.GetKeys()[9].GetValString()
			msg.HintIconOffset = m.GetKeys()[10].GetValFloat()
			msg.HintRange = m.GetKeys()[11].GetValFloat()
			msg.HintFlags = m.GetKeys()[12].GetValLong()
			msg.HintBinding = m.GetKeys()[13].GetValString()
			msg.HintAllowNodrawTarget = m.GetKeys()[14].GetValBool()
			msg.HintNooffscreen = m.GetKeys()[15].GetValBool()
			msg.HintForcecaption = m.GetKeys()[16].GetValBool()
			msg.HintLocalPlayerOnly = m.GetKeys()[17].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 60: // EGameEvent_InstructorServerHintStop
		if cbs := ge.onInstructorServerHintStop; cbs != nil {
			msg := &GameEventInstructorServerHintStop{}
			msg.HintName = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 61: // EGameEvent_ChatNewMessage
		if cbs := ge.onChatNewMessage; cbs != nil {
			msg := &GameEventChatNewMessage{}
			msg.Channel = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 62: // EGameEvent_ChatMembersChanged
		if cbs := ge.onChatMembersChanged; cbs != nil {
			msg := &GameEventChatMembersChanged{}
			msg.Channel = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 63: // EGameEvent_InventoryUpdated
		if cbs := ge.onInventoryUpdated; cbs != nil {
			msg := &GameEventInventoryUpdated{}
			msg.Itemdef = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 64: // EGameEvent_CartUpdated
		if cbs := ge.onCartUpdated; cbs != nil {
			msg := &GameEventCartUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 65: // EGameEvent_StorePricesheetUpdated
		if cbs := ge.onStorePricesheetUpdated; cbs != nil {
			msg := &GameEventStorePricesheetUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 66: // EGameEvent_GcConnected
		if cbs := ge.onGcConnected; cbs != nil {
			msg := &GameEventGcConnected{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 67: // EGameEvent_ItemSchemaInitialized
		if cbs := ge.onItemSchemaInitialized; cbs != nil {
			msg := &GameEventItemSchemaInitialized{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 68: // EGameEvent_DropRateModified
		if cbs := ge.onDropRateModified; cbs != nil {
			msg := &GameEventDropRateModified{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 69: // EGameEvent_EventTicketModified
		if cbs := ge.onEventTicketModified; cbs != nil {
			msg := &GameEventEventTicketModified{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 70: // EGameEvent_ModifierEvent
		if cbs := ge.onModifierEvent; cbs != nil {
			msg := &GameEventModifierEvent{}
			msg.Eventname = m.GetKeys()[0].GetValString()
			msg.Caster = m.GetKeys()[1].GetValShort()
			msg.Ability = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 71: // EGameEvent_DotaPlayerKill
		if cbs := ge.onDotaPlayerKill; cbs != nil {
			msg := &GameEventDotaPlayerKill{}
			msg.VictimUserid = m.GetKeys()[0].GetValShort()
			msg.Killer1Userid = m.GetKeys()[1].GetValShort()
			msg.Killer2Userid = m.GetKeys()[2].GetValShort()
			msg.Killer3Userid = m.GetKeys()[3].GetValShort()
			msg.Killer4Userid = m.GetKeys()[4].GetValShort()
			msg.Killer5Userid = m.GetKeys()[5].GetValShort()
			msg.Bounty = m.GetKeys()[6].GetValShort()
			msg.Neutral = m.GetKeys()[7].GetValShort()
			msg.Greevil = m.GetKeys()[8].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 72: // EGameEvent_DotaPlayerDeny
		if cbs := ge.onDotaPlayerDeny; cbs != nil {
			msg := &GameEventDotaPlayerDeny{}
			msg.KillerUserid = m.GetKeys()[0].GetValShort()
			msg.VictimUserid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 73: // EGameEvent_DotaBarracksKill
		if cbs := ge.onDotaBarracksKill; cbs != nil {
			msg := &GameEventDotaBarracksKill{}
			msg.BarracksId = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 74: // EGameEvent_DotaTowerKill
		if cbs := ge.onDotaTowerKill; cbs != nil {
			msg := &GameEventDotaTowerKill{}
			msg.KillerUserid = m.GetKeys()[0].GetValShort()
			msg.Teamnumber = m.GetKeys()[1].GetValShort()
			msg.Gold = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 75: // EGameEvent_DotaEffigyKill
		if cbs := ge.onDotaEffigyKill; cbs != nil {
			msg := &GameEventDotaEffigyKill{}
			msg.OwnerUserid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 76: // EGameEvent_DotaRoshanKill
		if cbs := ge.onDotaRoshanKill; cbs != nil {
			msg := &GameEventDotaRoshanKill{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()
			msg.Gold = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 77: // EGameEvent_DotaCourierLost
		if cbs := ge.onDotaCourierLost; cbs != nil {
			msg := &GameEventDotaCourierLost{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 78: // EGameEvent_DotaCourierRespawned
		if cbs := ge.onDotaCourierRespawned; cbs != nil {
			msg := &GameEventDotaCourierRespawned{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 79: // EGameEvent_DotaGlyphUsed
		if cbs := ge.onDotaGlyphUsed; cbs != nil {
			msg := &GameEventDotaGlyphUsed{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 80: // EGameEvent_DotaSuperCreeps
		if cbs := ge.onDotaSuperCreeps; cbs != nil {
			msg := &GameEventDotaSuperCreeps{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 81: // EGameEvent_DotaItemPurchase
		if cbs := ge.onDotaItemPurchase; cbs != nil {
			msg := &GameEventDotaItemPurchase{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 82: // EGameEvent_DotaItemGifted
		if cbs := ge.onDotaItemGifted; cbs != nil {
			msg := &GameEventDotaItemGifted{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValShort()
			msg.Sourceid = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 83: // EGameEvent_DotaRunePickup
		if cbs := ge.onDotaRunePickup; cbs != nil {
			msg := &GameEventDotaRunePickup{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Type = m.GetKeys()[1].GetValShort()
			msg.Rune = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 84: // EGameEvent_DotaRuneSpotted
		if cbs := ge.onDotaRuneSpotted; cbs != nil {
			msg := &GameEventDotaRuneSpotted{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Rune = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 85: // EGameEvent_DotaItemSpotted
		if cbs := ge.onDotaItemSpotted; cbs != nil {
			msg := &GameEventDotaItemSpotted{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 86: // EGameEvent_DotaNoBattlePoints
		if cbs := ge.onDotaNoBattlePoints; cbs != nil {
			msg := &GameEventDotaNoBattlePoints{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Reason = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 87: // EGameEvent_DotaChatInformational
		if cbs := ge.onDotaChatInformational; cbs != nil {
			msg := &GameEventDotaChatInformational{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Type = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 88: // EGameEvent_DotaActionItem
		if cbs := ge.onDotaActionItem; cbs != nil {
			msg := &GameEventDotaActionItem{}
			msg.Reason = m.GetKeys()[0].GetValShort()
			msg.Itemdef = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 89: // EGameEvent_DotaChatBanNotification
		if cbs := ge.onDotaChatBanNotification; cbs != nil {
			msg := &GameEventDotaChatBanNotification{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 90: // EGameEvent_DotaChatEvent
		if cbs := ge.onDotaChatEvent; cbs != nil {
			msg := &GameEventDotaChatEvent{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Gold = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 91: // EGameEvent_DotaChatTimedReward
		if cbs := ge.onDotaChatTimedReward; cbs != nil {
			msg := &GameEventDotaChatTimedReward{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itmedef = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 92: // EGameEvent_DotaPauseEvent
		if cbs := ge.onDotaPauseEvent; cbs != nil {
			msg := &GameEventDotaPauseEvent{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Value = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 93: // EGameEvent_DotaChatKillStreak
		if cbs := ge.onDotaChatKillStreak; cbs != nil {
			msg := &GameEventDotaChatKillStreak{}
			msg.Gold = m.GetKeys()[0].GetValShort()
			msg.KillerId = m.GetKeys()[1].GetValShort()
			msg.KillerStreak = m.GetKeys()[2].GetValShort()
			msg.KillerMultikill = m.GetKeys()[3].GetValShort()
			msg.VictimId = m.GetKeys()[4].GetValShort()
			msg.VictimStreak = m.GetKeys()[5].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 94: // EGameEvent_DotaChatFirstBlood
		if cbs := ge.onDotaChatFirstBlood; cbs != nil {
			msg := &GameEventDotaChatFirstBlood{}
			msg.Gold = m.GetKeys()[0].GetValShort()
			msg.KillerId = m.GetKeys()[1].GetValShort()
			msg.VictimId = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 95: // EGameEvent_DotaChatAssassinAnnounce
		if cbs := ge.onDotaChatAssassinAnnounce; cbs != nil {
			msg := &GameEventDotaChatAssassinAnnounce{}
			msg.AssassinId = m.GetKeys()[0].GetValShort()
			msg.TargetId = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 96: // EGameEvent_DotaChatAssassinDenied
		if cbs := ge.onDotaChatAssassinDenied; cbs != nil {
			msg := &GameEventDotaChatAssassinDenied{}
			msg.AssassinId = m.GetKeys()[0].GetValShort()
			msg.TargetId = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 97: // EGameEvent_DotaChatAssassinSuccess
		if cbs := ge.onDotaChatAssassinSuccess; cbs != nil {
			msg := &GameEventDotaChatAssassinSuccess{}
			msg.AssassinId = m.GetKeys()[0].GetValShort()
			msg.TargetId = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 98: // EGameEvent_DotaPlayerUpdateHeroSelection
		if cbs := ge.onDotaPlayerUpdateHeroSelection; cbs != nil {
			msg := &GameEventDotaPlayerUpdateHeroSelection{}
			msg.Tabcycle = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 99: // EGameEvent_DotaPlayerUpdateSelectedUnit
		if cbs := ge.onDotaPlayerUpdateSelectedUnit; cbs != nil {
			msg := &GameEventDotaPlayerUpdateSelectedUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 100: // EGameEvent_DotaPlayerUpdateQueryUnit
		if cbs := ge.onDotaPlayerUpdateQueryUnit; cbs != nil {
			msg := &GameEventDotaPlayerUpdateQueryUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 101: // EGameEvent_DotaPlayerUpdateKillcamUnit
		if cbs := ge.onDotaPlayerUpdateKillcamUnit; cbs != nil {
			msg := &GameEventDotaPlayerUpdateKillcamUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 102: // EGameEvent_DotaPlayerTakeTowerDamage
		if cbs := ge.onDotaPlayerTakeTowerDamage; cbs != nil {
			msg := &GameEventDotaPlayerTakeTowerDamage{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Damage = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 103: // EGameEvent_DotaHudErrorMessage
		if cbs := ge.onDotaHudErrorMessage; cbs != nil {
			msg := &GameEventDotaHudErrorMessage{}
			msg.Reason = m.GetKeys()[0].GetValByte()
			msg.Message = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 104: // EGameEvent_DotaActionSuccess
		if cbs := ge.onDotaActionSuccess; cbs != nil {
			msg := &GameEventDotaActionSuccess{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 105: // EGameEvent_DotaStartingPositionChanged
		if cbs := ge.onDotaStartingPositionChanged; cbs != nil {
			msg := &GameEventDotaStartingPositionChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 106: // EGameEvent_DotaMoneyChanged
		if cbs := ge.onDotaMoneyChanged; cbs != nil {
			msg := &GameEventDotaMoneyChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 107: // EGameEvent_DotaEnemyMoneyChanged
		if cbs := ge.onDotaEnemyMoneyChanged; cbs != nil {
			msg := &GameEventDotaEnemyMoneyChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 108: // EGameEvent_DotaPortraitUnitStatsChanged
		if cbs := ge.onDotaPortraitUnitStatsChanged; cbs != nil {
			msg := &GameEventDotaPortraitUnitStatsChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 109: // EGameEvent_DotaPortraitUnitModifiersChanged
		if cbs := ge.onDotaPortraitUnitModifiersChanged; cbs != nil {
			msg := &GameEventDotaPortraitUnitModifiersChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 110: // EGameEvent_DotaForcePortraitUpdate
		if cbs := ge.onDotaForcePortraitUpdate; cbs != nil {
			msg := &GameEventDotaForcePortraitUpdate{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 111: // EGameEvent_DotaInventoryChanged
		if cbs := ge.onDotaInventoryChanged; cbs != nil {
			msg := &GameEventDotaInventoryChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 112: // EGameEvent_DotaItemPickedUp
		if cbs := ge.onDotaItemPickedUp; cbs != nil {
			msg := &GameEventDotaItemPickedUp{}
			msg.Itemname = m.GetKeys()[0].GetValString()
			msg.PlayerID = m.GetKeys()[1].GetValShort()
			msg.ItemEntityIndex = m.GetKeys()[2].GetValShort()
			msg.HeroEntityIndex = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 113: // EGameEvent_DotaInventoryItemChanged
		if cbs := ge.onDotaInventoryItemChanged; cbs != nil {
			msg := &GameEventDotaInventoryItemChanged{}
			msg.EntityIndex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 114: // EGameEvent_DotaAbilityChanged
		if cbs := ge.onDotaAbilityChanged; cbs != nil {
			msg := &GameEventDotaAbilityChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 115: // EGameEvent_DotaPortraitAbilityLayoutChanged
		if cbs := ge.onDotaPortraitAbilityLayoutChanged; cbs != nil {
			msg := &GameEventDotaPortraitAbilityLayoutChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 116: // EGameEvent_DotaInventoryItemAdded
		if cbs := ge.onDotaInventoryItemAdded; cbs != nil {
			msg := &GameEventDotaInventoryItemAdded{}
			msg.Itemname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 117: // EGameEvent_DotaInventoryChangedQueryUnit
		if cbs := ge.onDotaInventoryChangedQueryUnit; cbs != nil {
			msg := &GameEventDotaInventoryChangedQueryUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 118: // EGameEvent_DotaLinkClicked
		if cbs := ge.onDotaLinkClicked; cbs != nil {
			msg := &GameEventDotaLinkClicked{}
			msg.Link = m.GetKeys()[0].GetValString()
			msg.Nav = m.GetKeys()[1].GetValBool()
			msg.NavBack = m.GetKeys()[2].GetValBool()
			msg.Recipe = m.GetKeys()[3].GetValShort()
			msg.Shop = m.GetKeys()[4].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 119: // EGameEvent_DotaSetQuickBuy
		if cbs := ge.onDotaSetQuickBuy; cbs != nil {
			msg := &GameEventDotaSetQuickBuy{}
			msg.Item = m.GetKeys()[0].GetValString()
			msg.Recipe = m.GetKeys()[1].GetValByte()
			msg.Toggle = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 120: // EGameEvent_DotaQuickBuyChanged
		if cbs := ge.onDotaQuickBuyChanged; cbs != nil {
			msg := &GameEventDotaQuickBuyChanged{}
			msg.Item = m.GetKeys()[0].GetValString()
			msg.Recipe = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 121: // EGameEvent_DotaPlayerShopChanged
		if cbs := ge.onDotaPlayerShopChanged; cbs != nil {
			msg := &GameEventDotaPlayerShopChanged{}
			msg.Prevshopmask = m.GetKeys()[0].GetValByte()
			msg.Shopmask = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 122: // EGameEvent_DotaPlayerShowKillcam
		if cbs := ge.onDotaPlayerShowKillcam; cbs != nil {
			msg := &GameEventDotaPlayerShowKillcam{}
			msg.Nodes = m.GetKeys()[0].GetValByte()
			msg.Player = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 123: // EGameEvent_DotaPlayerShowMinikillcam
		if cbs := ge.onDotaPlayerShowMinikillcam; cbs != nil {
			msg := &GameEventDotaPlayerShowMinikillcam{}
			msg.Nodes = m.GetKeys()[0].GetValByte()
			msg.Player = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 124: // EGameEvent_GcUserSessionCreated
		if cbs := ge.onGcUserSessionCreated; cbs != nil {
			msg := &GameEventGcUserSessionCreated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 125: // EGameEvent_TeamDataUpdated
		if cbs := ge.onTeamDataUpdated; cbs != nil {
			msg := &GameEventTeamDataUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 126: // EGameEvent_GuildDataUpdated
		if cbs := ge.onGuildDataUpdated; cbs != nil {
			msg := &GameEventGuildDataUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 127: // EGameEvent_GuildOpenPartiesUpdated
		if cbs := ge.onGuildOpenPartiesUpdated; cbs != nil {
			msg := &GameEventGuildOpenPartiesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 128: // EGameEvent_FantasyUpdated
		if cbs := ge.onFantasyUpdated; cbs != nil {
			msg := &GameEventFantasyUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 129: // EGameEvent_FantasyLeagueChanged
		if cbs := ge.onFantasyLeagueChanged; cbs != nil {
			msg := &GameEventFantasyLeagueChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 130: // EGameEvent_FantasyScoreInfoChanged
		if cbs := ge.onFantasyScoreInfoChanged; cbs != nil {
			msg := &GameEventFantasyScoreInfoChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 131: // EGameEvent_PlayerInfoUpdated
		if cbs := ge.onPlayerInfoUpdated; cbs != nil {
			msg := &GameEventPlayerInfoUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 132: // EGameEvent_PlayerInfoIndividualUpdated
		if cbs := ge.onPlayerInfoIndividualUpdated; cbs != nil {
			msg := &GameEventPlayerInfoIndividualUpdated{}
			msg.AccountId = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 133: // EGameEvent_GameRulesStateChange
		if cbs := ge.onGameRulesStateChange; cbs != nil {
			msg := &GameEventGameRulesStateChange{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 134: // EGameEvent_MatchHistoryUpdated
		if cbs := ge.onMatchHistoryUpdated; cbs != nil {
			msg := &GameEventMatchHistoryUpdated{}
			msg.SteamID = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 135: // EGameEvent_MatchDetailsUpdated
		if cbs := ge.onMatchDetailsUpdated; cbs != nil {
			msg := &GameEventMatchDetailsUpdated{}
			msg.MatchID = m.GetKeys()[0].GetValUint64()
			msg.Result = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 136: // EGameEvent_LiveGamesUpdated
		if cbs := ge.onLiveGamesUpdated; cbs != nil {
			msg := &GameEventLiveGamesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 137: // EGameEvent_RecentMatchesUpdated
		if cbs := ge.onRecentMatchesUpdated; cbs != nil {
			msg := &GameEventRecentMatchesUpdated{}
			msg.Page = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 138: // EGameEvent_NewsUpdated
		if cbs := ge.onNewsUpdated; cbs != nil {
			msg := &GameEventNewsUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 139: // EGameEvent_PersonaUpdated
		if cbs := ge.onPersonaUpdated; cbs != nil {
			msg := &GameEventPersonaUpdated{}
			msg.SteamID = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 140: // EGameEvent_TournamentStateUpdated
		if cbs := ge.onTournamentStateUpdated; cbs != nil {
			msg := &GameEventTournamentStateUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 141: // EGameEvent_PartyUpdated
		if cbs := ge.onPartyUpdated; cbs != nil {
			msg := &GameEventPartyUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 142: // EGameEvent_LobbyUpdated
		if cbs := ge.onLobbyUpdated; cbs != nil {
			msg := &GameEventLobbyUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 143: // EGameEvent_DashboardCachesCleared
		if cbs := ge.onDashboardCachesCleared; cbs != nil {
			msg := &GameEventDashboardCachesCleared{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 144: // EGameEvent_LastHit
		if cbs := ge.onLastHit; cbs != nil {
			msg := &GameEventLastHit{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.EntKilled = m.GetKeys()[1].GetValShort()
			msg.FirstBlood = m.GetKeys()[2].GetValBool()
			msg.HeroKill = m.GetKeys()[3].GetValBool()
			msg.TowerKill = m.GetKeys()[4].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 145: // EGameEvent_PlayerCompletedGame
		if cbs := ge.onPlayerCompletedGame; cbs != nil {
			msg := &GameEventPlayerCompletedGame{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Winner = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 146: // EGameEvent_PlayerReconnected
		if cbs := ge.onPlayerReconnected; cbs != nil {
			msg := &GameEventPlayerReconnected{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 147: // EGameEvent_NommedTree
		if cbs := ge.onNommedTree; cbs != nil {
			msg := &GameEventNommedTree{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 148: // EGameEvent_DotaRuneActivatedServer
		if cbs := ge.onDotaRuneActivatedServer; cbs != nil {
			msg := &GameEventDotaRuneActivatedServer{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Rune = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 149: // EGameEvent_DotaPlayerGainedLevel
		if cbs := ge.onDotaPlayerGainedLevel; cbs != nil {
			msg := &GameEventDotaPlayerGainedLevel{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Level = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 150: // EGameEvent_DotaPlayerLearnedAbility
		if cbs := ge.onDotaPlayerLearnedAbility; cbs != nil {
			msg := &GameEventDotaPlayerLearnedAbility{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Abilityname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 151: // EGameEvent_DotaPlayerUsedAbility
		if cbs := ge.onDotaPlayerUsedAbility; cbs != nil {
			msg := &GameEventDotaPlayerUsedAbility{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Abilityname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 152: // EGameEvent_DotaNonPlayerUsedAbility
		if cbs := ge.onDotaNonPlayerUsedAbility; cbs != nil {
			msg := &GameEventDotaNonPlayerUsedAbility{}
			msg.Abilityname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 153: // EGameEvent_DotaPlayerBeginCast
		if cbs := ge.onDotaPlayerBeginCast; cbs != nil {
			msg := &GameEventDotaPlayerBeginCast{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Abilityname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 154: // EGameEvent_DotaNonPlayerBeginCast
		if cbs := ge.onDotaNonPlayerBeginCast; cbs != nil {
			msg := &GameEventDotaNonPlayerBeginCast{}
			msg.Abilityname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 155: // EGameEvent_DotaAbilityChannelFinished
		if cbs := ge.onDotaAbilityChannelFinished; cbs != nil {
			msg := &GameEventDotaAbilityChannelFinished{}
			msg.Abilityname = m.GetKeys()[0].GetValString()
			msg.Interrupted = m.GetKeys()[1].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 156: // EGameEvent_DotaHoldoutReviveComplete
		if cbs := ge.onDotaHoldoutReviveComplete; cbs != nil {
			msg := &GameEventDotaHoldoutReviveComplete{}
			msg.Caster = m.GetKeys()[0].GetValShort()
			msg.Target = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 157: // EGameEvent_DotaPlayerKilled
		if cbs := ge.onDotaPlayerKilled; cbs != nil {
			msg := &GameEventDotaPlayerKilled{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.HeroKill = m.GetKeys()[1].GetValBool()
			msg.TowerKill = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 158: // EGameEvent_BindpanelOpen
		if cbs := ge.onBindpanelOpen; cbs != nil {
			msg := &GameEventBindpanelOpen{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 159: // EGameEvent_BindpanelClose
		if cbs := ge.onBindpanelClose; cbs != nil {
			msg := &GameEventBindpanelClose{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 160: // EGameEvent_KeybindChanged
		if cbs := ge.onKeybindChanged; cbs != nil {
			msg := &GameEventKeybindChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 161: // EGameEvent_DotaItemDragBegin
		if cbs := ge.onDotaItemDragBegin; cbs != nil {
			msg := &GameEventDotaItemDragBegin{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 162: // EGameEvent_DotaItemDragEnd
		if cbs := ge.onDotaItemDragEnd; cbs != nil {
			msg := &GameEventDotaItemDragEnd{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 163: // EGameEvent_DotaShopItemDragBegin
		if cbs := ge.onDotaShopItemDragBegin; cbs != nil {
			msg := &GameEventDotaShopItemDragBegin{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 164: // EGameEvent_DotaShopItemDragEnd
		if cbs := ge.onDotaShopItemDragEnd; cbs != nil {
			msg := &GameEventDotaShopItemDragEnd{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 165: // EGameEvent_DotaItemPurchased
		if cbs := ge.onDotaItemPurchased; cbs != nil {
			msg := &GameEventDotaItemPurchased{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Itemname = m.GetKeys()[1].GetValString()
			msg.Itemcost = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 166: // EGameEvent_DotaItemCombined
		if cbs := ge.onDotaItemCombined; cbs != nil {
			msg := &GameEventDotaItemCombined{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Itemname = m.GetKeys()[1].GetValString()
			msg.Itemcost = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 167: // EGameEvent_DotaItemUsed
		if cbs := ge.onDotaItemUsed; cbs != nil {
			msg := &GameEventDotaItemUsed{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Itemname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 168: // EGameEvent_DotaItemAutoPurchase
		if cbs := ge.onDotaItemAutoPurchase; cbs != nil {
			msg := &GameEventDotaItemAutoPurchase{}
			msg.ItemId = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 169: // EGameEvent_DotaUnitEvent
		if cbs := ge.onDotaUnitEvent; cbs != nil {
			msg := &GameEventDotaUnitEvent{}
			msg.Victim = m.GetKeys()[0].GetValShort()
			msg.Attacker = m.GetKeys()[1].GetValShort()
			msg.Basepriority = m.GetKeys()[2].GetValShort()
			msg.Priority = m.GetKeys()[3].GetValShort()
			msg.Eventtype = m.GetKeys()[4].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 170: // EGameEvent_DotaQuestStarted
		if cbs := ge.onDotaQuestStarted; cbs != nil {
			msg := &GameEventDotaQuestStarted{}
			msg.QuestIndex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 171: // EGameEvent_DotaQuestCompleted
		if cbs := ge.onDotaQuestCompleted; cbs != nil {
			msg := &GameEventDotaQuestCompleted{}
			msg.QuestIndex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 172: // EGameEvent_GameuiActivated
		if cbs := ge.onGameuiActivated; cbs != nil {
			msg := &GameEventGameuiActivated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 173: // EGameEvent_GameuiHidden
		if cbs := ge.onGameuiHidden; cbs != nil {
			msg := &GameEventGameuiHidden{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 174: // EGameEvent_PlayerFullyjoined
		if cbs := ge.onPlayerFullyjoined; cbs != nil {
			msg := &GameEventPlayerFullyjoined{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Name = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 175: // EGameEvent_DotaSpectateHero
		if cbs := ge.onDotaSpectateHero; cbs != nil {
			msg := &GameEventDotaSpectateHero{}
			msg.Entindex = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 176: // EGameEvent_DotaMatchDone
		if cbs := ge.onDotaMatchDone; cbs != nil {
			msg := &GameEventDotaMatchDone{}
			msg.Winningteam = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 177: // EGameEvent_DotaMatchDoneClient
		if cbs := ge.onDotaMatchDoneClient; cbs != nil {
			msg := &GameEventDotaMatchDoneClient{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 178: // EGameEvent_SetInstructorGroupEnabled
		if cbs := ge.onSetInstructorGroupEnabled; cbs != nil {
			msg := &GameEventSetInstructorGroupEnabled{}
			msg.Group = m.GetKeys()[0].GetValString()
			msg.Enabled = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 179: // EGameEvent_JoinedChatChannel
		if cbs := ge.onJoinedChatChannel; cbs != nil {
			msg := &GameEventJoinedChatChannel{}
			msg.ChannelName = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 180: // EGameEvent_LeftChatChannel
		if cbs := ge.onLeftChatChannel; cbs != nil {
			msg := &GameEventLeftChatChannel{}
			msg.ChannelName = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 181: // EGameEvent_GcChatChannelListUpdated
		if cbs := ge.onGcChatChannelListUpdated; cbs != nil {
			msg := &GameEventGcChatChannelListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 182: // EGameEvent_TodayMessagesUpdated
		if cbs := ge.onTodayMessagesUpdated; cbs != nil {
			msg := &GameEventTodayMessagesUpdated{}
			msg.NumMessages = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 183: // EGameEvent_FileDownloaded
		if cbs := ge.onFileDownloaded; cbs != nil {
			msg := &GameEventFileDownloaded{}
			msg.Success = m.GetKeys()[0].GetValBool()
			msg.LocalFilename = m.GetKeys()[1].GetValString()
			msg.RemoteUrl = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 184: // EGameEvent_PlayerReportCountsUpdated
		if cbs := ge.onPlayerReportCountsUpdated; cbs != nil {
			msg := &GameEventPlayerReportCountsUpdated{}
			msg.PositiveRemaining = m.GetKeys()[0].GetValByte()
			msg.NegativeRemaining = m.GetKeys()[1].GetValByte()
			msg.PositiveTotal = m.GetKeys()[2].GetValShort()
			msg.NegativeTotal = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 185: // EGameEvent_ScaleformFileDownloadComplete
		if cbs := ge.onScaleformFileDownloadComplete; cbs != nil {
			msg := &GameEventScaleformFileDownloadComplete{}
			msg.Success = m.GetKeys()[0].GetValBool()
			msg.LocalFilename = m.GetKeys()[1].GetValString()
			msg.RemoteUrl = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 186: // EGameEvent_ItemPurchased
		if cbs := ge.onItemPurchased; cbs != nil {
			msg := &GameEventItemPurchased{}
			msg.Itemid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 187: // EGameEvent_GcMismatchedVersion
		if cbs := ge.onGcMismatchedVersion; cbs != nil {
			msg := &GameEventGcMismatchedVersion{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 190: // EGameEvent_DemoStop
		if cbs := ge.onDemoStop; cbs != nil {
			msg := &GameEventDemoStop{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 191: // EGameEvent_MapShutdown
		if cbs := ge.onMapShutdown; cbs != nil {
			msg := &GameEventMapShutdown{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 192: // EGameEvent_DotaWorkshopFileselected
		if cbs := ge.onDotaWorkshopFileselected; cbs != nil {
			msg := &GameEventDotaWorkshopFileselected{}
			msg.Filename = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 193: // EGameEvent_DotaWorkshopFilecanceled
		if cbs := ge.onDotaWorkshopFilecanceled; cbs != nil {
			msg := &GameEventDotaWorkshopFilecanceled{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 194: // EGameEvent_RichPresenceUpdated
		if cbs := ge.onRichPresenceUpdated; cbs != nil {
			msg := &GameEventRichPresenceUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 195: // EGameEvent_DotaHeroRandom
		if cbs := ge.onDotaHeroRandom; cbs != nil {
			msg := &GameEventDotaHeroRandom{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Heroid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 196: // EGameEvent_DotaRdChatTurn
		if cbs := ge.onDotaRdChatTurn; cbs != nil {
			msg := &GameEventDotaRdChatTurn{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 197: // EGameEvent_DotaFavoriteHeroesUpdated
		if cbs := ge.onDotaFavoriteHeroesUpdated; cbs != nil {
			msg := &GameEventDotaFavoriteHeroesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 198: // EGameEvent_ProfileOpened
		if cbs := ge.onProfileOpened; cbs != nil {
			msg := &GameEventProfileOpened{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 199: // EGameEvent_ProfileClosed
		if cbs := ge.onProfileClosed; cbs != nil {
			msg := &GameEventProfileClosed{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 200: // EGameEvent_ItemPreviewClosed
		if cbs := ge.onItemPreviewClosed; cbs != nil {
			msg := &GameEventItemPreviewClosed{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 201: // EGameEvent_DashboardSwitchedSection
		if cbs := ge.onDashboardSwitchedSection; cbs != nil {
			msg := &GameEventDashboardSwitchedSection{}
			msg.Section = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 202: // EGameEvent_DotaTournamentItemEvent
		if cbs := ge.onDotaTournamentItemEvent; cbs != nil {
			msg := &GameEventDotaTournamentItemEvent{}
			msg.WinnerCount = m.GetKeys()[0].GetValShort()
			msg.EventType = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 203: // EGameEvent_DotaHeroSwap
		if cbs := ge.onDotaHeroSwap; cbs != nil {
			msg := &GameEventDotaHeroSwap{}
			msg.Playerid1 = m.GetKeys()[0].GetValByte()
			msg.Playerid2 = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 204: // EGameEvent_DotaResetSuggestedItems
		if cbs := ge.onDotaResetSuggestedItems; cbs != nil {
			msg := &GameEventDotaResetSuggestedItems{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 205: // EGameEvent_HalloweenHighScoreReceived
		if cbs := ge.onHalloweenHighScoreReceived; cbs != nil {
			msg := &GameEventHalloweenHighScoreReceived{}
			msg.Round = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 206: // EGameEvent_HalloweenPhaseEnd
		if cbs := ge.onHalloweenPhaseEnd; cbs != nil {
			msg := &GameEventHalloweenPhaseEnd{}
			msg.Phase = m.GetKeys()[0].GetValByte()
			msg.Team = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 207: // EGameEvent_HalloweenHighScoreRequestFailed
		if cbs := ge.onHalloweenHighScoreRequestFailed; cbs != nil {
			msg := &GameEventHalloweenHighScoreRequestFailed{}
			msg.Round = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 208: // EGameEvent_DotaHudSkinChanged
		if cbs := ge.onDotaHudSkinChanged; cbs != nil {
			msg := &GameEventDotaHudSkinChanged{}
			msg.Skin = m.GetKeys()[0].GetValString()
			msg.Style = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 209: // EGameEvent_DotaInventoryPlayerGotItem
		if cbs := ge.onDotaInventoryPlayerGotItem; cbs != nil {
			msg := &GameEventDotaInventoryPlayerGotItem{}
			msg.Itemname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 210: // EGameEvent_PlayerIsExperienced
		if cbs := ge.onPlayerIsExperienced; cbs != nil {
			msg := &GameEventPlayerIsExperienced{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 211: // EGameEvent_PlayerIsNotexperienced
		if cbs := ge.onPlayerIsNotexperienced; cbs != nil {
			msg := &GameEventPlayerIsNotexperienced{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 212: // EGameEvent_DotaTutorialLessonStart
		if cbs := ge.onDotaTutorialLessonStart; cbs != nil {
			msg := &GameEventDotaTutorialLessonStart{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 213: // EGameEvent_DotaTutorialTaskAdvance
		if cbs := ge.onDotaTutorialTaskAdvance; cbs != nil {
			msg := &GameEventDotaTutorialTaskAdvance{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 214: // EGameEvent_DotaTutorialShopToggled
		if cbs := ge.onDotaTutorialShopToggled; cbs != nil {
			msg := &GameEventDotaTutorialShopToggled{}
			msg.ShopOpened = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 215: // EGameEvent_MapLocationUpdated
		if cbs := ge.onMapLocationUpdated; cbs != nil {
			msg := &GameEventMapLocationUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 216: // EGameEvent_RichpresenceCustomUpdated
		if cbs := ge.onRichpresenceCustomUpdated; cbs != nil {
			msg := &GameEventRichpresenceCustomUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 217: // EGameEvent_GameEndVisible
		if cbs := ge.onGameEndVisible; cbs != nil {
			msg := &GameEventGameEndVisible{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 218: // EGameEvent_AntiaddictionUpdate
		if cbs := ge.onAntiaddictionUpdate; cbs != nil {
			msg := &GameEventAntiaddictionUpdate{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 219: // EGameEvent_HighlightHudElement
		if cbs := ge.onHighlightHudElement; cbs != nil {
			msg := &GameEventHighlightHudElement{}
			msg.Elementname = m.GetKeys()[0].GetValString()
			msg.Duration = m.GetKeys()[1].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 220: // EGameEvent_HideHighlightHudElement
		if cbs := ge.onHideHighlightHudElement; cbs != nil {
			msg := &GameEventHideHighlightHudElement{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 221: // EGameEvent_IntroVideoFinished
		if cbs := ge.onIntroVideoFinished; cbs != nil {
			msg := &GameEventIntroVideoFinished{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 222: // EGameEvent_MatchmakingStatusVisibilityChanged
		if cbs := ge.onMatchmakingStatusVisibilityChanged; cbs != nil {
			msg := &GameEventMatchmakingStatusVisibilityChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 223: // EGameEvent_PracticeLobbyVisibilityChanged
		if cbs := ge.onPracticeLobbyVisibilityChanged; cbs != nil {
			msg := &GameEventPracticeLobbyVisibilityChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 224: // EGameEvent_DotaCourierTransferItem
		if cbs := ge.onDotaCourierTransferItem; cbs != nil {
			msg := &GameEventDotaCourierTransferItem{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 225: // EGameEvent_FullUiUnlocked
		if cbs := ge.onFullUiUnlocked; cbs != nil {
			msg := &GameEventFullUiUnlocked{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 227: // EGameEvent_HeroSelectorPreviewSet
		if cbs := ge.onHeroSelectorPreviewSet; cbs != nil {
			msg := &GameEventHeroSelectorPreviewSet{}
			msg.Setindex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 228: // EGameEvent_AntiaddictionToast
		if cbs := ge.onAntiaddictionToast; cbs != nil {
			msg := &GameEventAntiaddictionToast{}
			msg.Message = m.GetKeys()[0].GetValString()
			msg.Duration = m.GetKeys()[1].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 229: // EGameEvent_HeroPickerShown
		if cbs := ge.onHeroPickerShown; cbs != nil {
			msg := &GameEventHeroPickerShown{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 230: // EGameEvent_HeroPickerHidden
		if cbs := ge.onHeroPickerHidden; cbs != nil {
			msg := &GameEventHeroPickerHidden{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 231: // EGameEvent_DotaLocalQuickbuyChanged
		if cbs := ge.onDotaLocalQuickbuyChanged; cbs != nil {
			msg := &GameEventDotaLocalQuickbuyChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 232: // EGameEvent_ShowCenterMessage
		if cbs := ge.onShowCenterMessage; cbs != nil {
			msg := &GameEventShowCenterMessage{}
			msg.Message = m.GetKeys()[0].GetValString()
			msg.Duration = m.GetKeys()[1].GetValFloat()
			msg.ClearMessageQueue = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 233: // EGameEvent_HudFlipChanged
		if cbs := ge.onHudFlipChanged; cbs != nil {
			msg := &GameEventHudFlipChanged{}
			msg.Flipped = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 234: // EGameEvent_FrostyPointsUpdated
		if cbs := ge.onFrostyPointsUpdated; cbs != nil {
			msg := &GameEventFrostyPointsUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 235: // EGameEvent_Defeated
		if cbs := ge.onDefeated; cbs != nil {
			msg := &GameEventDefeated{}
			msg.Entindex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 236: // EGameEvent_ResetDefeated
		if cbs := ge.onResetDefeated; cbs != nil {
			msg := &GameEventResetDefeated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 237: // EGameEvent_BoosterStateUpdated
		if cbs := ge.onBoosterStateUpdated; cbs != nil {
			msg := &GameEventBoosterStateUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 238: // EGameEvent_EventPointsUpdated
		if cbs := ge.onEventPointsUpdated; cbs != nil {
			msg := &GameEventEventPointsUpdated{}
			msg.EventId = m.GetKeys()[0].GetValShort()
			msg.Points = m.GetKeys()[1].GetValShort()
			msg.PremiumPoints = m.GetKeys()[2].GetValShort()
			msg.Owned = m.GetKeys()[3].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 239: // EGameEvent_LocalPlayerEventPoints
		if cbs := ge.onLocalPlayerEventPoints; cbs != nil {
			msg := &GameEventLocalPlayerEventPoints{}
			msg.Points = m.GetKeys()[0].GetValShort()
			msg.ConversionRate = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 240: // EGameEvent_CustomGameDifficulty
		if cbs := ge.onCustomGameDifficulty; cbs != nil {
			msg := &GameEventCustomGameDifficulty{}
			msg.Difficulty = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 241: // EGameEvent_TreeCut
		if cbs := ge.onTreeCut; cbs != nil {
			msg := &GameEventTreeCut{}
			msg.TreeX = m.GetKeys()[0].GetValFloat()
			msg.TreeY = m.GetKeys()[1].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 242: // EGameEvent_UgcDetailsArrived
		if cbs := ge.onUgcDetailsArrived; cbs != nil {
			msg := &GameEventUgcDetailsArrived{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 243: // EGameEvent_UgcSubscribed
		if cbs := ge.onUgcSubscribed; cbs != nil {
			msg := &GameEventUgcSubscribed{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 244: // EGameEvent_UgcUnsubscribed
		if cbs := ge.onUgcUnsubscribed; cbs != nil {
			msg := &GameEventUgcUnsubscribed{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 245: // EGameEvent_UgcDownloadRequested
		if cbs := ge.onUgcDownloadRequested; cbs != nil {
			msg := &GameEventUgcDownloadRequested{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 246: // EGameEvent_UgcInstalled
		if cbs := ge.onUgcInstalled; cbs != nil {
			msg := &GameEventUgcInstalled{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 247: // EGameEvent_PrizepoolReceived
		if cbs := ge.onPrizepoolReceived; cbs != nil {
			msg := &GameEventPrizepoolReceived{}
			msg.Success = m.GetKeys()[0].GetValBool()
			msg.Prizepool = m.GetKeys()[1].GetValUint64()
			msg.Leagueid = m.GetKeys()[2].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 248: // EGameEvent_MicrotransactionSuccess
		if cbs := ge.onMicrotransactionSuccess; cbs != nil {
			msg := &GameEventMicrotransactionSuccess{}
			msg.Txnid = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 249: // EGameEvent_DotaRubickAbilitySteal
		if cbs := ge.onDotaRubickAbilitySteal; cbs != nil {
			msg := &GameEventDotaRubickAbilitySteal{}
			msg.AbilityIndex = m.GetKeys()[0].GetValShort()
			msg.AbilityLevel = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 250: // EGameEvent_CompendiumEventActionsLoaded
		if cbs := ge.onCompendiumEventActionsLoaded; cbs != nil {
			msg := &GameEventCompendiumEventActionsLoaded{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()
			msg.OriginalPoints = m.GetKeys()[3].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 251: // EGameEvent_CompendiumSelectionsLoaded
		if cbs := ge.onCompendiumSelectionsLoaded; cbs != nil {
			msg := &GameEventCompendiumSelectionsLoaded{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 252: // EGameEvent_CompendiumSetSelectionFailed
		if cbs := ge.onCompendiumSetSelectionFailed; cbs != nil {
			msg := &GameEventCompendiumSetSelectionFailed{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 253: // EGameEvent_CompendiumTrophiesLoaded
		if cbs := ge.onCompendiumTrophiesLoaded; cbs != nil {
			msg := &GameEventCompendiumTrophiesLoaded{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 254: // EGameEvent_CommunityCachedNamesUpdated
		if cbs := ge.onCommunityCachedNamesUpdated; cbs != nil {
			msg := &GameEventCommunityCachedNamesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 255: // EGameEvent_SpecItemPickup
		if cbs := ge.onSpecItemPickup; cbs != nil {
			msg := &GameEventSpecItemPickup{}
			msg.PlayerId = m.GetKeys()[0].GetValShort()
			msg.ItemName = m.GetKeys()[1].GetValString()
			msg.Purchase = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 256: // EGameEvent_SpecAegisReclaimTime
		if cbs := ge.onSpecAegisReclaimTime; cbs != nil {
			msg := &GameEventSpecAegisReclaimTime{}
			msg.ReclaimTime = m.GetKeys()[0].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 257: // EGameEvent_AccountTrophiesChanged
		if cbs := ge.onAccountTrophiesChanged; cbs != nil {
			msg := &GameEventAccountTrophiesChanged{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 258: // EGameEvent_AccountAllHeroChallengeChanged
		if cbs := ge.onAccountAllHeroChallengeChanged; cbs != nil {
			msg := &GameEventAccountAllHeroChallengeChanged{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 259: // EGameEvent_TeamShowcaseUiUpdate
		if cbs := ge.onTeamShowcaseUiUpdate; cbs != nil {
			msg := &GameEventTeamShowcaseUiUpdate{}
			msg.Show = m.GetKeys()[0].GetValBool()
			msg.AccountId = m.GetKeys()[1].GetValUint64()
			msg.HeroEntindex = m.GetKeys()[2].GetValShort()
			msg.DisplayUiOnLeft = m.GetKeys()[3].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 260: // EGameEvent_IngameEventsChanged
		if cbs := ge.onIngameEventsChanged; cbs != nil {
			msg := &GameEventIngameEventsChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 261: // EGameEvent_DotaMatchSignout
		if cbs := ge.onDotaMatchSignout; cbs != nil {
			msg := &GameEventDotaMatchSignout{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 262: // EGameEvent_DotaIllusionsCreated
		if cbs := ge.onDotaIllusionsCreated; cbs != nil {
			msg := &GameEventDotaIllusionsCreated{}
			msg.OriginalEntindex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 263: // EGameEvent_DotaYearBeastKilled
		if cbs := ge.onDotaYearBeastKilled; cbs != nil {
			msg := &GameEventDotaYearBeastKilled{}
			msg.KillerPlayerId = m.GetKeys()[0].GetValShort()
			msg.Message = m.GetKeys()[1].GetValShort()
			msg.BeastId = m.GetKeys()[2].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 264: // EGameEvent_DotaHeroUndoselection
		if cbs := ge.onDotaHeroUndoselection; cbs != nil {
			msg := &GameEventDotaHeroUndoselection{}
			msg.Playerid1 = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 265: // EGameEvent_DotaChallengeSocacheUpdated
		if cbs := ge.onDotaChallengeSocacheUpdated; cbs != nil {
			msg := &GameEventDotaChallengeSocacheUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 266: // EGameEvent_PartyInvitesUpdated
		if cbs := ge.onPartyInvitesUpdated; cbs != nil {
			msg := &GameEventPartyInvitesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 267: // EGameEvent_LobbyInvitesUpdated
		if cbs := ge.onLobbyInvitesUpdated; cbs != nil {
			msg := &GameEventLobbyInvitesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 268: // EGameEvent_CustomGameModeListUpdated
		if cbs := ge.onCustomGameModeListUpdated; cbs != nil {
			msg := &GameEventCustomGameModeListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 269: // EGameEvent_CustomGameLobbyListUpdated
		if cbs := ge.onCustomGameLobbyListUpdated; cbs != nil {
			msg := &GameEventCustomGameLobbyListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 270: // EGameEvent_FriendLobbyListUpdated
		if cbs := ge.onFriendLobbyListUpdated; cbs != nil {
			msg := &GameEventFriendLobbyListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 271: // EGameEvent_DotaTeamPlayerListChanged
		if cbs := ge.onDotaTeamPlayerListChanged; cbs != nil {
			msg := &GameEventDotaTeamPlayerListChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 272: // EGameEvent_DotaPlayerDetailsChanged
		if cbs := ge.onDotaPlayerDetailsChanged; cbs != nil {
			msg := &GameEventDotaPlayerDetailsChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 273: // EGameEvent_PlayerProfileStatsUpdated
		if cbs := ge.onPlayerProfileStatsUpdated; cbs != nil {
			msg := &GameEventPlayerProfileStatsUpdated{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 274: // EGameEvent_CustomGamePlayerCountUpdated
		if cbs := ge.onCustomGamePlayerCountUpdated; cbs != nil {
			msg := &GameEventCustomGamePlayerCountUpdated{}
			msg.CustomGameId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 275: // EGameEvent_CustomGameFriendsPlayedUpdated
		if cbs := ge.onCustomGameFriendsPlayedUpdated; cbs != nil {
			msg := &GameEventCustomGameFriendsPlayedUpdated{}
			msg.CustomGameId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 276: // EGameEvent_CustomGamesFriendsPlayUpdated
		if cbs := ge.onCustomGamesFriendsPlayUpdated; cbs != nil {
			msg := &GameEventCustomGamesFriendsPlayUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 277: // EGameEvent_DotaPlayerUpdateAssignedHero
		if cbs := ge.onDotaPlayerUpdateAssignedHero; cbs != nil {
			msg := &GameEventDotaPlayerUpdateAssignedHero{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 278: // EGameEvent_DotaPlayerHeroSelectionDirty
		if cbs := ge.onDotaPlayerHeroSelectionDirty; cbs != nil {
			msg := &GameEventDotaPlayerHeroSelectionDirty{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 279: // EGameEvent_DotaNpcGoalReached
		if cbs := ge.onDotaNpcGoalReached; cbs != nil {
			msg := &GameEventDotaNpcGoalReached{}
			msg.NpcEntindex = m.GetKeys()[0].GetValShort()
			msg.GoalEntindex = m.GetKeys()[1].GetValShort()
			msg.NextGoalEntindex = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 280: // EGameEvent_HltvStatus
		if cbs := ge.onHltvStatus; cbs != nil {
			msg := &GameEventHltvStatus{}
			msg.Clients = m.GetKeys()[0].GetValLong()
			msg.Slots = m.GetKeys()[1].GetValLong()
			msg.Proxies = m.GetKeys()[2].GetValShort()
			msg.Master = m.GetKeys()[3].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 281: // EGameEvent_HltvCameraman
		if cbs := ge.onHltvCameraman; cbs != nil {
			msg := &GameEventHltvCameraman{}
			msg.Index = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 282: // EGameEvent_HltvRankCamera
		if cbs := ge.onHltvRankCamera; cbs != nil {
			msg := &GameEventHltvRankCamera{}
			msg.Index = m.GetKeys()[0].GetValByte()
			msg.Rank = m.GetKeys()[1].GetValFloat()
			msg.Target = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 283: // EGameEvent_HltvRankEntity
		if cbs := ge.onHltvRankEntity; cbs != nil {
			msg := &GameEventHltvRankEntity{}
			msg.Index = m.GetKeys()[0].GetValShort()
			msg.Rank = m.GetKeys()[1].GetValFloat()
			msg.Target = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 284: // EGameEvent_HltvFixed
		if cbs := ge.onHltvFixed; cbs != nil {
			msg := &GameEventHltvFixed{}
			msg.Posx = m.GetKeys()[0].GetValLong()
			msg.Posy = m.GetKeys()[1].GetValLong()
			msg.Posz = m.GetKeys()[2].GetValLong()
			msg.Theta = m.GetKeys()[3].GetValShort()
			msg.Phi = m.GetKeys()[4].GetValShort()
			msg.Offset = m.GetKeys()[5].GetValShort()
			msg.Fov = m.GetKeys()[6].GetValFloat()
			msg.Target = m.GetKeys()[7].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 285: // EGameEvent_HltvChase
		if cbs := ge.onHltvChase; cbs != nil {
			msg := &GameEventHltvChase{}
			msg.Target1 = m.GetKeys()[0].GetValShort()
			msg.Target2 = m.GetKeys()[1].GetValShort()
			msg.Distance = m.GetKeys()[2].GetValShort()
			msg.Theta = m.GetKeys()[3].GetValShort()
			msg.Phi = m.GetKeys()[4].GetValShort()
			msg.Inertia = m.GetKeys()[5].GetValByte()
			msg.Ineye = m.GetKeys()[6].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 286: // EGameEvent_HltvMessage
		if cbs := ge.onHltvMessage; cbs != nil {
			msg := &GameEventHltvMessage{}
			msg.Text = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 287: // EGameEvent_HltvTitle
		if cbs := ge.onHltvTitle; cbs != nil {
			msg := &GameEventHltvTitle{}
			msg.Text = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 288: // EGameEvent_HltvChat
		if cbs := ge.onHltvChat; cbs != nil {
			msg := &GameEventHltvChat{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Text = m.GetKeys()[1].GetValString()
			msg.SteamID = m.GetKeys()[2].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 289: // EGameEvent_HltvVersioninfo
		if cbs := ge.onHltvVersioninfo; cbs != nil {
			msg := &GameEventHltvVersioninfo{}
			msg.Version = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 290: // EGameEvent_DotaChaseHero
		if cbs := ge.onDotaChaseHero; cbs != nil {
			msg := &GameEventDotaChaseHero{}
			msg.Target1 = m.GetKeys()[0].GetValShort()
			msg.Target2 = m.GetKeys()[1].GetValShort()
			msg.Type = m.GetKeys()[2].GetValByte()
			msg.Priority = m.GetKeys()[3].GetValShort()
			msg.Gametime = m.GetKeys()[4].GetValFloat()
			msg.Highlight = m.GetKeys()[5].GetValBool()
			msg.Target1playerid = m.GetKeys()[6].GetValByte()
			msg.Target2playerid = m.GetKeys()[7].GetValByte()
			msg.Eventtype = m.GetKeys()[8].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 291: // EGameEvent_DotaCombatlog
		if cbs := ge.onDotaCombatlog; cbs != nil {
			msg := &GameEventDotaCombatlog{}
			msg.Type = m.GetKeys()[0].GetValByte()
			msg.Sourcename = m.GetKeys()[1].GetValShort()
			msg.Targetname = m.GetKeys()[2].GetValShort()
			msg.Attackername = m.GetKeys()[3].GetValShort()
			msg.Inflictorname = m.GetKeys()[4].GetValShort()
			msg.Attackerillusion = m.GetKeys()[5].GetValBool()
			msg.Targetillusion = m.GetKeys()[6].GetValBool()
			msg.Value = m.GetKeys()[7].GetValShort()
			msg.Health = m.GetKeys()[8].GetValShort()
			msg.Timestamp = m.GetKeys()[9].GetValFloat()
			msg.Targetsourcename = m.GetKeys()[10].GetValShort()
			msg.Timestampraw = m.GetKeys()[11].GetValFloat()
			msg.Attackerhero = m.GetKeys()[12].GetValBool()
			msg.Targethero = m.GetKeys()[13].GetValBool()
			msg.AbilityToggleOn = m.GetKeys()[14].GetValBool()
			msg.AbilityToggleOff = m.GetKeys()[15].GetValBool()
			msg.AbilityLevel = m.GetKeys()[16].GetValShort()
			msg.GoldReason = m.GetKeys()[17].GetValShort()
			msg.XpReason = m.GetKeys()[18].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 292: // EGameEvent_DotaGameStateChange
		if cbs := ge.onDotaGameStateChange; cbs != nil {
			msg := &GameEventDotaGameStateChange{}
			msg.OldState = m.GetKeys()[0].GetValShort()
			msg.NewState = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 293: // EGameEvent_DotaPlayerPickHero
		if cbs := ge.onDotaPlayerPickHero; cbs != nil {
			msg := &GameEventDotaPlayerPickHero{}
			msg.Player = m.GetKeys()[0].GetValShort()
			msg.Heroindex = m.GetKeys()[1].GetValShort()
			msg.Hero = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 294: // EGameEvent_DotaTeamKillCredit
		if cbs := ge.onDotaTeamKillCredit; cbs != nil {
			msg := &GameEventDotaTeamKillCredit{}
			msg.KillerUserid = m.GetKeys()[0].GetValShort()
			msg.VictimUserid = m.GetKeys()[1].GetValShort()
			msg.Teamnumber = m.GetKeys()[2].GetValShort()
			msg.Herokills = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	}

	_panicf("unknown message %d", m.GetEventid())
	return nil
}
