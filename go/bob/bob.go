// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

// Hey should have a comment documenting it.
func Hey(remark string) string {

	isCapitalized := true
	hasLetters := false
	lastChar := remark[len(remark)-1]

	for _, x := range remark {
		if x >= 65 && x <= 90 {
			hasLetters = true
		}

		if x >= 97 && x <= 122 {
			hasLetters = true
			isCapitalized = false
		}
	}

	if !hasLetters && lastChar == '?' {
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
