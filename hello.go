package hello

import (
	"strings"
)

// takes a slice of strings
func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return "Hello, " + strings.Join(names, ", ")
}