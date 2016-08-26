package diffsquares

import "math"

func SquareOfSums(n int) int {
	var s int
	for i := 0; i <= n; i++ {
		s += i
	}
	return int(math.Pow(float64(s), 2))
}

func SumOfSquares(n int) int {
	var s float64
	for i := 0.; i <= float64(n); i++ {
		s += math.Pow(i, 2)
	}
	return int(s)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
