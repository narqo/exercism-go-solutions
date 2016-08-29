package pascal

// Triangle computes Pascal's triangle up to a given number of rows.
func Triangle(n int) [][]int {
	t := make([][]int, n)
	ncols := 1
	for row := 0; row < n; row++ {
		cols := make([]int, ncols)

		t[row] = cols
		cols[0] = 1
		cols[len(cols)-1] = 1

		for col := 1; col < ncols-1; col++ {
			cols[col] = t[row-1][col-1] + t[row-1][col]
		}
		ncols++
	}
	return t
}
