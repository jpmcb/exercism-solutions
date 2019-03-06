// Package bob contains constants to define various ASCII values used in
// string matching functions. Further, exposes the Hey function that
// returns the appropriate bob saying
package bob

const (
	null           rune = 0
	tab            rune = 9
	space          rune = 32
	cr             rune = 13
	nl             rune = 10
	upperCaseStart rune = 65
	upperCaseEnd   rune = 90
	lowerCaseStart rune = 97
	lowerCaseEnd   rune = 122
)

func getLastChar(remark string) rune {
	var lastChar rune

	if len(remark) == 0 {
		return null
	}

	i := len(remark) - 1
	lastChar = rune(remark[i])

	for (lastChar == tab || lastChar == space || lastChar == cr || lastChar == nl) && i > 0 {
		i--
		lastChar = rune(remark[i])
	}

	return lastChar
}

func checkStringConditions(remark string) (bool, bool, bool) {
	isAllCapitalized := true
	isEmpty := true
	hasLetters := false

	for _, x := range remark {
		if x >= upperCaseStart && x <= upperCaseEnd {
			hasLetters = true
		}

		if x >= lowerCaseStart && x <= lowerCaseEnd {
			hasLetters = true
			isAllCapitalized = false
		}

		if x != tab && x != space && x != cr && x != nl {
			isEmpty = false
		}
	}

	return isAllCapitalized, isEmpty, hasLetters
}

// Hey performs logic to check the input string and determine it's various properties.
// These properties are then used to return the appropriate statement defined in the readme
func Hey(remark string) string {

	lastChar := getLastChar(remark)
	isAllCapitalized, isEmpty, hasLetters := checkStringConditions(remark)

	if lastChar == null {
		return "Fine. Be that way!"
	} else if isEmpty {
		return "Fine. Be that way!"
	} else if !hasLetters && lastChar == '?' {
		return "Sure."
	} else if !hasLetters {
		return "Whatever."
	} else if isAllCapitalized && lastChar == '!' {
		return "Whoa, chill out!"
	} else if isAllCapitalized && lastChar == '?' {
		return "Calm down, I know what I'm doing!"
	} else if isAllCapitalized {
		return "Whoa, chill out!"
	} else if lastChar == '?' {
		return "Sure."
	}

	return "Whatever."
}
