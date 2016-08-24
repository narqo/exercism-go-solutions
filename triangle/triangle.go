package triangle

import "math"

const testVersion = 2

// Kind represents a possible kind of the triangle
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns  if the triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	if !isPossibleTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func isPossibleTriangle(a, b, c float64) bool {
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		return false
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return false
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return false
	}
	return a+b >= c && a+c >= b && b+c >= a
}
