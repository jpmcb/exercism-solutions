// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

const (
	tab   rune = 9
	space rune = 32
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var (
		isCapitalized bool
		isEmpty       bool
		hasLetters    bool
		lastChar      byte
	)
	isCapitalized = true
	isEmpty = true
	hasLetters = false

	if len(remark) > 0 {
		lastChar = remark[len(remark)-1]
	} else {
		lastChar = 0x0
	}

	for _, x := range remark {
		if x >= 65 && x <= 90 {
			hasLetters = true
		}

		if x >= 97 && x <= 122 {
			hasLetters = true
			isCapitalized = false
		}

		if x != tab || x != space {
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
	} else if isCapitalized && lastChar == '!' {
		return "Whoa, chill out!"
	} else if isCapitalized && lastChar == '?' {
		return "Calm down, I know what I'm doing!"
	} else if isCapitalized {
		return "Whoa, chill out!"
	} else if lastChar == '?' {
		return "Sure."
	}

	return "Whatever."
}
