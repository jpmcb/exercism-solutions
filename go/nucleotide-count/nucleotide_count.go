package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA. Given as a map.
type Histogram map[rune]int

// DNA is a list of nucleotides given as a string
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA string
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, x := range d {
		switch x {
		case 'A':
			h['A']++
		case 'C':
			h['C']++
		case 'G':
			h['G']++
		case 'T':
			h['T']++
		default:
			return nil, fmt.Errorf("Undefined necleotide found: %c", x)
		}
	}

	return h, nil
}
