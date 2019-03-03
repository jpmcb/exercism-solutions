package strand

// ToRNA takes a DNA strand and converts it to it's RNA equivelent
func ToRNA(dna string) string {
	var builder string

	for _, strand := range dna {
		switch strand {
		case 'G':
			builder += string('C')
		case 'C':
			builder += string('G')
		case 'T':
			builder += string('A')
		case 'A':
			builder += string('U')
		}
	}

	return builder
}
