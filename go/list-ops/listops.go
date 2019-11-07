package listops

type IntList []int

func (list IntList) Foldr(f func(int, int) int, init int) int {
	for _, x := range list {
		init = f(init, x)
	}

	return init
}

// func (*list) Filter() {

// }

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
