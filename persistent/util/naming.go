package util

import (
	"strings"
)

func SnakeToCamelCase(s string) string {
	if strings.HasPrefix(s, "__") && strings.HasSuffix(s, "__") {
		s = "Magic" + strings.Title(s[2:len(s)-2])
	}

	comp := strings.Split(s, "_")
	for i, c := range comp {
		comp[i] = strings.Title(c)
	}

	return strings.Join(comp, "")
}
