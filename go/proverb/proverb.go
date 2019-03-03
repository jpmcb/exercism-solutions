// Package proverb implements a function that returns the slice of strings
// formatted to be as "For want of ..., For want of..., And all for ..."
package proverb

import "fmt"

// Proverb returns a slice of strings that forms the predefined proverb format
func Proverb(rhyme []string) []string {
	result := []string{}

	for i := range rhyme {
		if i == len(rhyme)-1 {
			result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}

	return result
}
