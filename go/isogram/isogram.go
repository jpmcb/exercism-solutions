package isogram

import "strings"

// IsIsogram returns true if isogram, false otherwise
// ignores case, hyphines, and spaces
func IsIsogram(i string) bool {
	i = strings.ToLower(i)
	letterMap := make(map[rune]bool)

	for _, char := range i {
		if char == '-' || char == ' ' {
			continue
		}

		if letterMap[char] {
			return false
		}

		letterMap[char] = true
	}

	return true
}
