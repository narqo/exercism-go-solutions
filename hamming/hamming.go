package hamming

import "fmt"

const testVersion = 4

// Distance returns the Hamming distance between sequencies
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("different length: %d vs %d", len(a), len(b))
	}
	var distance int
	for n := range a {
		if a[n] != b[n] {
			distance++
		}
	}
	return distance, nil
}
