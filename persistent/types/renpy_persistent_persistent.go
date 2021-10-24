package types

import (
	"fmt"
	"reflect"

	"github.com/nlpodyssey/gopickle/types"
	"mon.icu/pkg/go-persistent-loader/persistent/util"
)

func init() {
	registerType("renpy.persistent", "Persistent", &RenpyPersistentPersistent{})
}

type RenpyPersistentPersistent struct {
	MasAcsPreList                       *RenpyPythonRevertableList
	MasWindowreactsNotifFilters         RenpyPythonRevertableDict
	MasPmIsBullyingVictim               bool
	MasPmHaveFam                        bool
	SetPreferences                      bool
	MasPmTotalHeldMonika                unknown
	Preferences                         *RenpyPreferencesPreferences
	MasLikesHairdown                    bool
	MasDateLastCheckedRain              DatetimeDate
	MasDisableEggs                      bool
	MasChessTimedDisable                unknown
	MasAutoModeEnabled                  bool
	MasNyeDateAffGain                   int
	MasO31WentTrickOrTreatingLonglong   bool
	MasPmDrivingBeenInAccident          unknown
	FilePage                            string
	MasPctadeibe                        unknown
	MasO31CostumeGreetingSeen           bool
	MasPmVolunteerCharity               bool
	MasCurrentBackground                string
	MasCoffeeBrewTime                   unknown
	SeenMonikaInRoom                    bool
	MasPmZoomedIn                       bool
	MasCrashedBefore                    bool
	MasPmPromMonika                     bool
	MasPmPlayJazz                       unknown
	MasPmD25MistletoeKiss               bool
	MasAbsenceTime                      DatetimeTimedelta
	MasCHotchocBrewTime                 unknown
	MasFilereactsJustReacted            bool
	MasShouldRainToday                  bool
	MasPmSeeTherapist                   bool
	MasD25ChibikaSayoriPerformed        bool
	MasO31WentTrickOrTreatingLong       bool
	MasChessQuicksave                   string
	MasPlayerNicknames                  *RenpyPythonRevertableList
	MasPmCaresAboutDokis                bool
	RandomSeen                          int
	Playerxp                            unknown
	MasBdayDateAffectionGained          int
	MasBdayNeedToResetBday              bool
	MasGreetingTypeTimeout              unknown
	SeenGhostMenu                       unknown
	MasForceClothes                     bool
	MasPlayerBdayDateAffGain            int
	MasBdayNoTimeSpent                  bool
	MasO31WentTrickOrTreatingAbort      bool
	MasPmGetsSnow                       bool
	MasPmLikeOtherMusic                 bool
	MasD25InD25Mode                     bool
	Gender                              string
	MasD25StartedUpset                  bool
	MasPmHasFriends                     bool
	MasPmFailedFreshStart               unknown
	EventList                           *RenpyPythonRevertableList
	MasLastSeenNewStory                 RenpyPythonRevertableDict
	MasD25DecoActive                    bool
	MasDockstatCmNoCount                int
	MasPlayerDerandomed                 *RenpyPythonRevertableList
	MasPmPromGood                       unknown
	Steam                               bool
	MasPmLikesTravelling                bool
	GuiPreference                       RenpyPythonRevertableDict
	MasAcsBbhList                       *RenpyPythonRevertableList
	MasD25IntroSeen                     bool
	MasPmLikesNature                    bool
	MasAcsBatList                       *RenpyPythonRevertableList
	MasPmPlaysInstrument                bool
	MasPctaneibe                        unknown
	MasPmLikeRockNRoll                  bool
	MasSettingOcb                       bool
	MasPmReligious                      bool
	DirectorBottom                      bool
	MasO31TrickOrTreatingStartLate      bool
	MasSelsprClothesDb                  RenpyPythonRevertableDict
	MasXpLvl                            int
	MasStripDatesRules                  RenpyPythonRevertableDict
	MasEventInitLockdb                  RenpyPythonRevertableDict
	MasPmHangsD25Lights                 bool
	MasD25SecondChanceUpset             bool
	MasPmLikeOtherMusicHistory          *RenpyPythonRevertableList
	MasD25D25eDateCount                 int
	MasPmListenedToGradSpeech           bool
	MasCHotchocCupDone                  unknown
	MasPmLikesRain                      bool
	MasPmUnitsHeightMetric              bool
	PreferenceDefault                   RenpyPythonRevertableDict
	MasO31TrickOrTreatingAffGain        int
	MasSelsprAcsDb                      RenpyPythonRevertableDict
	MasChessStats                       RenpyPythonRevertableDict
	SeenAudio                           *types.Dict
	MasF14DateCount                     int
	MasPmLikePlayingTennis              unknown
	IdlexpTotal                         int
	MasWindowreactsWindowreactsEnabled  bool
	MonikaReload                        int
	Playername                          string
	MasDockstatCheckinLog               *RenpyPythonRevertableList
	MasBdaySbpReacted                   bool
	MasRandchatFreq                     int
	GuiPreferenceDefault                RenpyPythonRevertableDict
	MasChess3EditSorry                  bool
	MasAffection                        RenpyPythonRevertableDict
	MasF14PreIntroSeen                  bool
	MasCurrentWeather                   string
	MasPmAteLunchTimes                  *RenpyPythonRevertableList
	FarewellDatabase                    RenpyPythonRevertableDict
	MasBdayInBdayMode                   bool
	MasPmBoopStats                      RenpyPythonRevertableDict
	MasFirstCalendarCheck               bool
	MasAbsenceChoice                    string
	EverWon                             RenpyPythonRevertableDict
	MasFilereactsLastReactedDate        DatetimeDate
	MasDisableSorry                     unknown
	VoiceMute                           RenpyPythonRevertableSet
	MasImportedSaves                    bool
	MasPlayerBdaySawSurprise            bool
	MasD25ChibikaSayoriDone             bool
	MasPmLastPromotedD                  unknown
	MasPmAteBreakfastTimes              *RenpyPythonRevertableList
	MasCurrentSeason                    int
	MasF14InF14Mode                     bool
	MasPmMonikaDeletionJustice          bool
	MasPongDifficultyChangeNextGameDate DatetimeDate
	GivenChocolatesBefore               bool
	MonikaSaidTopics                    *RenpyPythonRevertableList
	MonikaTopic                         string
	MasD25GiftsGiven                    *RenpyPythonRevertableList
	MasF14SpentF14                      bool
	MasPmDoSmoke                        bool
	MasFilereactsSpriteReacted          RenpyPythonRevertableDict
	MasMonikaRepeatedHerself            unknown
	MasPmAddedCustomBgm                 bool
	MenuBgM                             unknown
	DateLastGivenRoses                  DatetimeDate
	MasEventClothesMap                  RenpyPythonRevertableDict
	MasPmHadRelationshipsJustOne        bool
	MasTriedGiftRing                    bool
	MasPmHasPianoExperience             unknown
	MasPmHasRpy                         bool
	OpendoorKnockyes                    bool
	MasPmLangJpn                        bool
	SeenSticker                         unknown
	MasAcsBabList                       unknown
	MasPmShavedHair                     unknown
	MasSunset                           int
	MasF14GoneOverF14                   unknown
	MasFilereactsFailedMap              RenpyPythonRevertableDict
	MasPmEatFastFood                    bool
	MasPmLikesPoetry                    bool
	Demo                                bool
	AchievementProgress                 *types.Dict
	Playthrough                         int
	MasLoadInFinalfarewellMode          bool
	MasGradSpeechTimedOut               bool
	MasHistoryArchives                  RenpyPythonRevertableDict
	MasO31Relaunch                      bool
	MasPmHaveFamSibs                    bool
	MasBdayHintFilename                 string
	MasPmHeight                         int
	MasO31CurrentCostume                unknown
	MasUndoActionRules                  RenpyPythonRevertableDict
	MasNyeWentOutNyd                    int
	MasNyeWentOutNye                    int
	MasPmMeditates                      bool
	MasFilereactsStopMap                RenpyPythonRevertableDict
	FileFolder                          int
	MasPmOwnsCar                        bool
	MasPmSwearFrequency                 int
	MasPnmlData                         *RenpyPythonRevertableList
	MasEverWon                          *CollectionsDefaultdict
	MasPmNoProm                         bool
	MasHairChanged                      bool
	MasDmDataVersion                    int
	MasYouChr                           bool
	MasPmLiveSouthHemisphere            bool
	MasPmLongestHeldMonika              DatetimeTimedelta
	MasPmDrivingCanDrive                bool
	MasMoodSick                         bool
	SeenTranslates                      BuiltinSet
	MasMonikaNickname                   string
	MasPmWorksOut                       bool
	MasMonikaLovecountertime            DatetimeDatetime
	MasO31DockstatReturn                bool
	MasIdleData                         RenpyPythonRevertableDict
	MasPmHadPromDate                    unknown
	MasPmDrivingLearning                unknown
	MasPmLikeJazz                       bool
	MasPoemsSeen                        RenpyPythonRevertableDict
	MasPmHadRelationshipsMany           bool
	MasPmHasBeenToAmusementPark         bool
	MasF14TimeSpentSeen                 bool
	Sessions                            RenpyPythonRevertableDict
	ConsoleShort                        bool
	MasPmIsTrans                        bool
	MasPmEverLetMonikaWinOnPurpose      bool
	MasPmAccomplishedResolutions        bool
	MasPmHairLength                     string
	MasO31WentTrickOrTreatingShort      bool
	MasPmLikeMintIceCream               bool
	MasPongDifficulty                   int
	MasPlayerBdayLastSungHbd            DatetimeDate
	MasWeatherMWdata                    RenpyPythonRevertableDict
	MasO31InO31Mode                     bool
	MasNyeNyeDateCount                  int
	MasDockstatCheckoutLog              *RenpyPythonRevertableList
	MasDarkModeEnabled                  bool
	MasBdayDateAffectionFix             bool
	MasPmWouldLikeMtPeak                bool
	MasEnableRandomRepeats              bool
	MasPmLikesSpoops                    bool
	MasAcsBbaList                       *RenpyPythonRevertableList
	MasBdayOnDate                       bool
	FilePageName                        RenpyPythonRevertableDict
	MasPmWouldComeToSpaceroom           bool
	MasMoodCurrent                      string
	MasFilereactsHistoric               RenpyPythonRevertableDict
	FirstRun                            bool
	MasAffectionGoodexpFreeze           bool
	MasAdvancedPyTips                   bool
	MasPlayerBdayLeftOnBday             bool
	MasMonikaReturnedHome               unknown
	MasPmLikesHorror                    bool
	MasCrashedSelf                      bool
	MasFilereactsSpriteGifts            RenpyPythonRevertableDict
	MasAffBeforeFreshStart              unknown
	MasMonikaLovecounter                int
	MasChessSkipFileChecks              bool
	MasBdayVisuals                      bool
	MasPlayerBookmarked                 *RenpyPythonRevertableList
	MasAcsPstList                       *RenpyPythonRevertableList
	MasUnstableMode                     bool
	MasD25WentOutD25e                   int
	MasPmWillChange                     unknown
	MasPmAteDinnerTimes                 *RenpyPythonRevertableList
	MasDevEnablePtods                   bool
	MasBdayGoneOverBday                 bool
	MasFunFactsDatabase                 RenpyPythonRevertableDict
	MasAcsMidList                       *RenpyPythonRevertableList
	MasAcsAseList                       *RenpyPythonRevertableList
	MasGameCrashed                      bool
	MasCoffeeCupDone                    unknown
	MasPmIsFastReader                   bool
	MasPmDonateCharity                  unknown
	MasPmFeelsLonelySometimes           bool
	TriedSkip                           bool
	MasPenname                          unknown
	MasAcsBmhList                       *RenpyPythonRevertableList
	MasPmHasNewYearsRes                 bool
	Achievements                        BuiltinSet
	Changed                             *types.Dict
	CurrentTrack                        unknown
	MasZoomZoomLevel                    int
	StylePreferences                    RenpyPythonRevertableDict
	MasAcsMatList                       *RenpyPythonRevertableList
	MasBdaySbpFoundBanners              bool
	YuriKill                            int
	MasPmGoneToProm                     bool
	MasO31SeenCostumes                  RenpyPythonRevertableDict
	MasPmGotAFreshStart                 unknown
	MasPmMonikaEvilButOk                bool
	MasPmLiveInCity                     bool
	MasD25AlreadyGiftedCookies          bool
	MasWindowreactsNoUnlockList         *RenpyPythonRevertableList
	MasComplimentsDatabase              RenpyPythonRevertableDict
	MasPmDoSmokeQuit                    unknown
	MasAcsBseList                       *RenpyPythonRevertableList
	SeenEyes                            bool
	MasPmHasContributedToMas            unknown
	MasFastbye                          bool
	FlaggedMonikatopic                  string
	ClosedSelf                          bool
	MasBdayConfirmedParty               bool
	MasLastAhogeDt                      DatetimeDatetime
	MasJustFriends                      bool
	MasF14IntroSeen                     bool
	MasPmEyeColor                       string
	MasSubmodVersionData                RenpyPythonRevertableDict
	MasD25WentOutD25                    int
	MasPmMonikaEvil                     bool
	VirtualSize                         *types.Tuple
	MasO31TrickOrTreatingStartEarly     bool
	MasLastKiss                         DatetimeDatetime
	MasPmHasCodeExperience              bool
	MasDockstatCmWaitCount              int
	MasPlayerBdayDate                   int
	MasPmPromShy                        unknown
	MasPongDifficultyChangeNextGame     int
	MasPmNoFamBother                    unknown
	MasPmNoHairNoTalk                   unknown
	MasBdaySbpFoundCake                 bool
	MasDockstatGoingToLeave             bool
	MasForceHair                        bool
	MasPmHasWentBackInTime              bool
	MasApologyReasonUseDb               RenpyPythonRevertableDict
	MasPmHairColor                      string
	MasGivenChocolatesBefore            bool
	MasPmReadYellowWp                   bool
	MasCurrentGiftedRibbons             int
	MasGreetingType                     unknown
	MasFilereactsReactedMap             RenpyPythonRevertableDict
	FirstLoad                           unknown
	MasPmFamLikeMonika                  bool
	MasMonikaBdaySurpriseHintSeen       bool
	MasXpTnl                            float64
	MasO31TtCount                       int
	MasUnseeUnseen                      bool
	MasPmLangOther                      bool
	MasBdayHasDoneBdOutro               bool
	MasD25ChibikaSayori                 unknown
	MasBdayOpenedGame                   bool
	MasPmLikeOrchestralMusic            bool
	MasGrandfatheredNickname            unknown
	MasPmNoTalkFam                      bool
	MasPmDoSmokeQuitSucceededBefore     unknown
	MasPmLikesSingingD25Carols          bool
	MasPmZoomedInMax                    bool
	MasStoryDatabase                    RenpyPythonRevertableDict
	MasChessMangleAll                   bool
	MasAffBackup                        float64
	MasPmHasBulliedPeople               bool
	MasF14NtsSeen                       bool
	MasNyeSpentNyd                      bool
	MasNyeSpentNye                      bool
	MasPmLikePlayingSports              bool
	MasOpenedExtraMenu                  bool
	MonikaBreakup                       int
	EventDatabase                       RenpyPythonRevertableDict
	MasBackgroundMBGdata                RenpyPythonRevertableDict
	MasCanUpdate                        bool
	MasMonikaBreakup                    int
	MasPmSharedAppearance               bool
	MasPmOwnsCarType                    unknown
	MasBdayNoRecognize                  bool
	MasCarrymeChoice                    unknown
	MasO31TrickOrTreatingStartNormal    bool
	SeenEver                            *types.Dict
	MasPmFewFriends                     bool
	MasAffectionLogCounter              int
	RejectedMonika                      unknown
	MasAcsMmhList                       *RenpyPythonRevertableList
	MasHangmanPlayername                bool
	MasPmCalledMoniABadName             bool
	MasTimeconcern                      int
	MasPmLikesBoardGames                bool
	MasO31WentTrickOrTreatingMid        bool
	MasPlayerBdaySpentTime              bool
	MasLastMonikaIly                    unknown
	MasPmWantsToContributeToMas         bool
	MasFunfactfun                       bool
	MasMonikaClothes                    string
	MasFilereactsLastAffGainedResetDate DatetimeDate
	MasPmCurrentlyBullied               bool
	MasCurrentConsumable                RenpyPythonRevertableDict
	MasEventInitLockdbTemplate          *types.Tuple
	MasGivenThermosBefore               bool
	MasPlayerConfirmedBday              bool
	MasBdayDateAffectionLost            int
	MasD25SpentD25                      bool
	CharacterVolume                     RenpyPythonRevertableDict
	MasConsumableMap                    RenpyPythonRevertableDict
	MasPmSocialPersonality              string
	MasPmNoTalkPanties                  bool
	MasD25GoneOverD25                   unknown
	MasPmPromNotInterested              bool
	MasPmTakenMonikaOut                 bool
	MasPmLikesPanties                   bool
	MasPlayerBdayOpenedDoor             bool
	MasFastgreeting                     bool
	MasDdlcReloadCount                  int
	MasPlayerBdayDecor                  bool
	MasPlayerDerandomedSongs            *RenpyPythonRevertableList
	MasLongAbsence                      bool
	MasPmGivenFalseJustice              bool
	MasMonikaDeletionJusticeKidding     unknown
	MasHistoryMhsData                   RenpyPythonRevertableDict
	MasIdleModeWasCrashed               unknown
	MasChangedStartDate                 bool
	MasAcsMabList                       *RenpyPythonRevertableList
	UpdateLastChecked                   RenpyPythonRevertableDict
	MasF14Date                          int
	MasAcsEnablePromisering             bool
	SpecialPoems                        *RenpyPythonRevertableList
	CurrentMonikatopic                  int
	MasD25D25DateCount                  int
	VersionNumber                       string
	MasPmLikeVocaloids                  bool
	MasPianoKeymaps                     RenpyPythonRevertableDict
	MasTimeconcernclose                 bool
	MasFirstKiss                        DatetimeDatetime
	MasMoodDatabase                     RenpyPythonRevertableDict
	Mcname                              string
	UpdateVersion                       RenpyPythonRevertableDict
	MasPmLoveYourself                   bool
	MasO31CostumesWorn                  RenpyPythonRevertableDict
	MasSpritesJsonGiftedSprites         RenpyPythonRevertableDict
	MasNotificationSounds               bool
	MasPmAteLateTimes                   int
	MasPmDrawnArt                       bool
	FirstPoem                           bool
	MasSensitiveMode                    bool
	MasBdaySaidHappybday                bool
	MasMoniChksum                       unknown
	MasPmDrivingPostAccident            unknown
	MasPmLiveNearBeach                  bool
	MasPoolUnlocks                      int
	MasPmZoomedOut                      bool
	MasPmLikedGradSpeech                bool
	Chosen                              *types.Dict
	MasXpRst                            DatetimeDate
	Anticheat                           int
	MonikaKill                          bool
	MasInIdleMode                       bool
	MasCrashedTrynot                    bool
	Clearall                            unknown
	MasLastHold                         DatetimeDate
	MasPlayerBday                       DatetimeDate
	MasPmAteSnackTimes                  *RenpyPythonRevertableList
	MasAcsEnableQuetzalplushie          bool
	MasPmGamedLate                      int
	MasPmMonikaCuteAsNatsuki            bool
	MasDockstatCmYesCount               int
	MasAcsBfhList                       *RenpyPythonRevertableList
	MasF14OnDate                        unknown
	MasApologyDatabase                  RenpyPythonRevertableDict
	MasTimeconcerngraveyard             bool
	MasWindowreactsDatabase             RenpyPythonRevertableDict
	MasLastHoldDt                       DatetimeDatetime
	MasLateFarewell                     bool
	GhostMenu                           bool
	MasXpHrx                            float64
	MasPmWatchMangime                   bool
	MasPctaieibe                        unknown
	MasEnableNotifications              bool
	MasApologyTimeDb                    RenpyPythonRevertableDict
	MasDockstatMoniSize                 int
	Autoload                            string
	MasD25SeenSantaCostume              bool
	MasBdaySbpFoundBalloons             bool
	MasSunrise                          int
	MasChessDlgActions                  RenpyPythonRevertableDict
	MasZzLupdExV                        *RenpyPythonRevertableList
	MasJustUpdated                      bool
	MasOfferedNickname                  bool
	MasEvYearsetBlacklist               RenpyPythonRevertableDict
	MasPmHaveFamMess                    bool
	MasSelsprHairDb                     RenpyPythonRevertableDict
	MasDelayedActionList                *RenpyPythonRevertableList
	MasPmLikeRap                        bool
	MasBdaySbpAffGiven                  int
	MasAcsAfhList                       *RenpyPythonRevertableList
	MasNyeNydDateCount                  int
	IapPurchases                        RenpyPythonRevertableDict
	MasDockstatMoniLog                  RenpyPythonRevertableDict
	MasPmSkinTone                       string
	MasPlayerBdayInPlayerBdayMode       bool
	GreetingDatabase                    RenpyPythonRevertableDict
	MasBdayDateCount                    int
	MasChessDifficulty                  *types.Tuple
	MasTextSpeedEnabled                 bool
	MasPmWearsRing                      bool
	Clear                               *RenpyPythonRevertableList
	MasMonikaHair                       string
	MasGameDatabase                     RenpyPythonRevertableDict
	MasDisableAnimations                bool
	MasO31CostumesAllowed               bool
	OpendoorOpencount                   int
	SeenImages                          *types.Dict
	MasO31WentTrickOrTreatingRight      bool
	MasF14DateAffGain                   int
	MasAffMismatches                    int
	MasPmHaveFamMessBetter              unknown
	MasFilereactsGiftAffGained          int
	MasSongsDatabase                    RenpyPythonRevertableDict
	MasPmAHater                         bool
	MasPmDrinksSoda                     bool
}

func (r RenpyPersistentPersistent) PyNew(_ ...interface{}) (interface{}, error) {
	return &RenpyPersistentPersistent{}, nil
}

func (r *RenpyPersistentPersistent) PyDictSet(key, value interface{}) error {
	normKey := util.SnakeToCamelCase(key.(string))
	field := reflect.ValueOf(r).Elem().FieldByName(normKey)
	if !field.IsValid() {
		return fmt.Errorf("%s %s", reflect.ValueOf(value).Kind().String(), normKey)
	}

	if value != nil {
		field.Set(reflect.ValueOf(value))
	}

	return nil
}
