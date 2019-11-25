package diffsquares

// SquareOfSum is ...
func SquareOfSum(n int) int {
	i := 1
	result := 0
	for i <= n {
		result += i
		i++
	}

	return result * result
}

// Difference is ...
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SumOfSquares is ...
func SumOfSquares(n int) int {
	i := 1
	result := 0
	for i <= n {
		result += (i * i)
		i++
	}

	return result
}
