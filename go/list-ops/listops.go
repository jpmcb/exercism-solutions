package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool

func (list IntList) Foldr(f binFunc, init int) int {
	for _, x := range list {
		init = f(init, x)
	}

	return init
}

func (list IntList) Foldl(f binFunc, init int) int {
	for _, x := range list {
		init = f(init, x)
	}

	return init
}

func (list IntList) Filter() {

}

// func (*list) Length() {

// }

// func (*list) Map() {

// }

// func (*list) Reverse() {

// }

// func (*list) Append() {

// }

// func (*list) Concat() {

// }
