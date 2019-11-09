package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (list IntList) Foldr(f binFunc, init int) int {
	for _, x := range list.Reverse() {
		init = f(x, init)
	}

	return init
}

func (list IntList) Foldl(f binFunc, init int) int {
	for _, x := range list {
		init = f(init, x)
	}

	return init
}

func (list IntList) Filter(f predFunc) IntList {
	return list
}

func (list IntList) Length() int {
	return len(list)
}

func (list IntList) Map(f unaryFunc) IntList {
	return list
}

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

func (list IntList) Append(x IntList) IntList {
	return list
}

func (list IntList) Concat(x []IntList) IntList {
	return list
}
