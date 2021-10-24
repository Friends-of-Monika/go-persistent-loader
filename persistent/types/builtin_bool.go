package types

func init() {
	registerType("__builtin__", "bool", false)
}

type BuiltinBool bool