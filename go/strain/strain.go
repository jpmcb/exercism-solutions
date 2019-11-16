package strain

// Ints is a slice of integers
type Ints []int

// Lists is a slice of integer slices
type Lists [][]int

// Strings is a slice of strings
type Strings []string

// Keep filters an Ints list based on the input function
func (list Ints) Keep(f func(int) bool) Ints {
	var newList Ints

	for _, x := range list {
		if f(x) {
			newList = append(newList, x)
		}
	}

	return newList
}

// Discard will remove int from list if provided func is true
func (list Ints) Discard(f func(int) bool) Ints {
	var newList Ints

	for _, x := range list {
		if !f(x) {
			newList = append(newList, x)
		}
	}

	return newList
}

// Keep filters Lists based on the input function
func (list Lists) Keep(f func([]int) bool) Lists {
	var newList Lists

	for _, x := range list {
		if f(x) {
			newList = append(newList, x)
		}
	}

	return newList
}

// Keep filters Strings based on the input function
func (list Strings) Keep(f func(string) bool) Strings {
	var newList Strings

	for _, x := range list {
		if f(x) {
			newList = append(newList, x)
		}
	}

	return newList
}
