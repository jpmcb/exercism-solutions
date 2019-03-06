// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

const (
	tab   rune = 9
	space rune = 32
	cr    rune = 13
	nl    rune = 10
)

func getLastChar(remark string) byte {
	var lastChar byte

	if len(remark) == 0 {
		return 0x0
	}

	i := len(remark) - 1
	lastChar = remark[i]

	for (lastChar == 9 || lastChar == 32 || lastChar == 13 || lastChar == 10) && i > 0 {
		i--
		lastChar = remark[i]
	}

	return lastChar
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var (
		isAllCapitalized bool
		isEmpty          bool
		hasLetters       bool
		lastChar         byte
	)
	isAllCapitalized = true
	isEmpty = true
	hasLetters = false
	lastChar = getLastChar(remark)

	for _, x := range remark {
		if x >= 65 && x <= 90 {
			hasLetters = true
		}

		if x >= 97 && x <= 122 {
			hasLetters = true
			isAllCapitalized = false
		}

		if x != tab && x != space && x != cr && x != nl {
			isEmpty = false
		}
	}

	if lastChar == 0x0 {
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
