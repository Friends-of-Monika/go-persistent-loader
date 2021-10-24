package types

func init() {
	registerType("renpy.python", "RevertableList", &RenpyPythonRevertableList{})
}

type RenpyPythonRevertableList struct {
	content []interface{}
}

func (r RenpyPythonRevertableList) PyNew(_ ...interface{}) (interface{}, error) {
	return &RenpyPythonRevertableList{make([]interface{}, 0)}, nil
}

func (r RenpyPythonRevertableList) PyDictSet(key, value interface{}) error {
	panic("implement me")
}

func (r *RenpyPythonRevertableList) Append(v interface{}) {
	r.content = append(r.content, v)
}
