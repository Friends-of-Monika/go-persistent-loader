package types

func init() {
	registerType("renpy.python", "RevertableDict", &RenpyPythonRevertableDict{})
}

type RenpyPythonRevertableDict map[interface{}]interface{}

func (r RenpyPythonRevertableDict) PyNew(_ ...interface{}) (interface{}, error) {
	return make(RenpyPythonRevertableDict), nil
}

func (r RenpyPythonRevertableDict) PyDictSet(_, _ interface{}) error {
	panic("implement me")
}

func (r RenpyPythonRevertableDict) Set(key, value interface{}) {
	r[key] = value
}
