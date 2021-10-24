package types

var index = make(map[string]map[string]interface{})

func registerType(module, name string, class interface{}) {
	if _, exists := index[module]; !exists {
		index[module] = make(map[string]interface{}, 1)
	}
	index[module][name] = class
}

func GetRegisteredType(module, name string) interface{} {
	if mod, exists := index[module]; exists {
		return mod[name]
	}

	return nil
}
