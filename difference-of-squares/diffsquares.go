package diffsquares

func SquareOfSums(n int) int {
	var s int
	for i := 0; i <= n; i++ {
		s += i
	}
	return s * s
}

func SumOfSquares(n int) int {
	var s int
	for i := 0; i <= n; i++ {
		s += i * i
	}
	return s
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
