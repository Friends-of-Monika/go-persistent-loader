package types

func init() {
	registerType("collections", "defaultdict", CollectionsDefaultdict{})
}

type CollectionsDefaultdict struct {
	content map[interface{}]interface{}
	def     interface{}
}

func (r CollectionsDefaultdict) Call(args ...interface{}) (interface{}, error) {
	return &CollectionsDefaultdict{make(map[interface{}]interface{}), args[0]}, nil
}

func (r CollectionsDefaultdict) Set(key, value interface{}) {
	r.content[key] = value
}

