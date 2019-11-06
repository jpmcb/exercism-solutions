package isogram

import "strings"

// IsIsogram returns true if isogram, false otherwise
// ignores case, hyphines, and spaces
func IsIsogram(i string) bool {
	i = strings.ToLower(i)

	letterMap := make(map[rune]int)
	for _, char := range i {
		_, ok := letterMap[char]
		if !ok {
			letterMap[char] = 1
		} else if char == '-' || char == ' ' {
			letterMap[char]++
		} else if letterMap[char] == 1 {
			return false
		}
	}

	return true
}
