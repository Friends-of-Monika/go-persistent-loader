package types

import (
	"github.com/nlpodyssey/gopickle/types"
)

func init() {
	registerType("__builtin__", "set", BuiltinSet{})
}

type BuiltinSet map[interface{}]struct{}

func (_ BuiltinSet) Call(args ...interface{}) (interface{}, error) {
	l := args[0].(*types.List)
	s := make(BuiltinSet, l.Len())
	for i := 0; i < l.Len(); i++ {
		v := l.Get(i)
		if _, exists := s[v]; !exists {
			s[v] = struct{}{}
		}
	}
	return s, nil
}
