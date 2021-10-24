package types

import (
	"errors"
	"reflect"

	"github.com/friends-of-monika/go-persistent-loader/persistent/util"
	"github.com/nlpodyssey/gopickle/types"
)

func init() {
	registerType("renpy.preferences", "Preferences", &RenpyPreferencesPreferences{})
}

type RenpyPreferencesPreferences struct {
	SelfVoicing         bool
	SkipUnseen          bool
	GlNpot              bool
	Mute                *types.Dict
	MouseMove           bool
	Renderer            string
	GlPowersave         bool
	Transitions         int
	FontLineSpacing     float64
	Fullscreen          bool
	GlTearing           bool
	VirtualSize         *types.Tuple
	SfxVolume           float64
	ShowEmptyWindow     bool
	EmphasizeAudio      bool
	SkipAfterChoices    bool
	WaitVoice           bool
	MagicVersion        int
	AfmTime             float64
	GlFramerate         unknown
	PerformanceTest     bool
	TextCps             int
	DesktopRollbackSide string
	MobileRollbackSide  string
	FontTransform       unknown
	AfmAfterClick       bool
	VideoImageFallback  bool
	Joymap              *types.Dict
	FontSize            float64
	AfmEnable           bool
	Language            unknown
	PhysicalSize        *types.Tuple
	VoiceSustain        bool
	MusicVolume         float64
	PadEnabled          bool
	Volumes             *types.Dict
	UsingAfmEnable      bool
}

func (r *RenpyPreferencesPreferences) PyNew(_ ...interface{}) (interface{}, error) {
	return &RenpyPreferencesPreferences{}, nil
}

func (r *RenpyPreferencesPreferences) PyDictSet(key, value interface{}) error {
	normKey := util.SnakeToCamelCase(key.(string))
	field := reflect.ValueOf(r).Elem().FieldByName(normKey)
	if !field.IsValid() {
		return errors.New(normKey)
	}

	if value != nil {
		field.Set(reflect.ValueOf(value))
	}

	return nil
}
