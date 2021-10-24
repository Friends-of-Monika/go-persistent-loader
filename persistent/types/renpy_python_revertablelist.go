package types

func init() {
	registerType("renpy.python", "RevertableList", &RenpyPythonRevertableList{})
}

type RenpyPythonRevertableList []interface{}

func (r RenpyPythonRevertableList) PyNew(_ ...interface{}) (interface{}, error) {
	inst := make(RenpyPythonRevertableList, 0)
	return &inst, nil
}

func (r RenpyPythonRevertableList) PyDictSet(_, _ interface{}) error {
	panic("RenpyPythonRevertableList#PyDictSet is not implemented")
}

func (r *RenpyPythonRevertableList) Append(v interface{}) {
	*r = append(*r, v)
}
