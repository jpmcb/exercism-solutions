package acronym

import (
	"strings"
)

// Abbreviate will generate the abbreviation
func Abbreviate(s string) string {
	var result string

	s = strings.Replace(s, "-", " ", -1)

	for i, char := range s {
		if i == 0 {
			result += string(char)
		} else if s[i-1] == ' ' && char != ' ' {
			result += string(char)
		}
	}

	return strings.ToUpper(result)
}
