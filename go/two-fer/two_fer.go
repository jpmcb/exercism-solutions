// Package twofer gives a simple "one for 'x', one for me" message
package twofer

import (
	"fmt"
)

// ShareWith implements twofer and returns the expected string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
