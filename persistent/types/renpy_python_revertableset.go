package types

func init() {
	registerType("renpy.python", "RevertableSet", &RenpyPythonRevertableSet{})
}

type RenpyPythonRevertableSet map[interface{}]struct{}

func (r RenpyPythonRevertableSet) PyNew(_ ...interface{}) (interface{}, error) {
	return make(RenpyPythonRevertableSet), nil
}
