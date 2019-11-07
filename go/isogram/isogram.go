package isogram

import "strings"

// IsIsogram returns true if isogram, false otherwise
// ignores case, hyphines, and spaces
func IsIsogram(i string) bool {
	i = strings.ToLower(i)
	letterMap := make(map[rune]bool)

	for _, char := range i {
		_, ok := letterMap[char]
		if !ok {
			letterMap[char] = true
		} else if char != '-' && char != ' ' {
			return false
		}
	}

	return true
}
