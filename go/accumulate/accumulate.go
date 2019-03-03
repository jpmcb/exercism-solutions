package accumulate

//Accumulate is a function
func Accumulate(given []string, converter func(string) string) []string {
	var results []string

	for _, word := range given {
		results = append(results, converter(word))
	}

	return results
}
