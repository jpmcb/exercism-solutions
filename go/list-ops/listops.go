package listops

// IntList is a slice of ints
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldr folds in a list from right to left
func (list IntList) Foldr(f binFunc, init int) int {
	for _, x := range list.Reverse() {
		init = f(x, init)
	}

	return init
}

// Foldl folds in a list from left to right
func (list IntList) Foldl(f binFunc, init int) int {
	for _, x := range list {
		init = f(init, x)
	}

	return init
}

// Filter will reduce a list based on argument function
func (list IntList) Filter(f predFunc) IntList {

	if list.Length() == 0 {
		return list
	}

	var newList IntList

	for _, x := range list {
		if f(x) {
			newList = append(newList, x)
		}
	}

	return newList
}

// Length will return the length of the IntList
func (list IntList) Length() int {
	return len(list)
}

// Map performs the argument input function on all list items
func (list IntList) Map(f unaryFunc) IntList {
	for i, x := range list {
		list[i] = f(x)
	}

	return list
}

// Reverse will swap items from right to left
func (list IntList) Reverse() IntList {
	frontRunner := 0
	backRunner := list.Length() - 1
	var temp int

	if backRunner == 0 {
		return list
	}

	for frontRunner <= backRunner {
		temp = list[frontRunner]
		list[frontRunner] = list[backRunner]
		list[backRunner] = temp

		frontRunner++
		backRunner--
	}

	return list
}

// Append will add the argument list to the IntList
func (list IntList) Append(input IntList) IntList {
	for _, x := range input {
		list = append(list, x)
	}
	return list
}

// Concat will combine the slice of IntLists into a single IntList
func (list IntList) Concat(inputs []IntList) IntList {
	for _, x := range inputs {
		list = list.Append(x)
	}

	return list
}
