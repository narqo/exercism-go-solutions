package raindrops

import "strconv"

const testVersion = 2

var nToDrop = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(n int) string {
	var (
		s    string
		prev int
	)
	for _, f := range primeFactors(n) {
		if f == prev {
			continue
		}
		if drop, ok := nToDrop[f]; ok {
			s += drop
			prev = f
		}
	}
	if len(s) == 0 {
		return strconv.Itoa(n)
	}
	return s
}

func primeFactors(n int) []int {
	factors := []int{}
	d := 2
	for n > 1 {
		for ; n%d == 0; n /= d {
			factors = append(factors, d)
		}
		d++
	}
	return factors
}
