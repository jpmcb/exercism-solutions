package hamming

import "errors"

// Distance returns the hamming distance of two dna strings
func Distance(a, b string) (int, error) {
	if len(a) == len(b) {
		count := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}

		return count, nil
	}

	return 0, errors.New("dna strands unequal lengths")
}
