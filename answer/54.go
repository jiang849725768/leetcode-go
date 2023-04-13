package answer

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	const (
		right int = iota
		down
		left
		up
	)
	l, r, d, u := 0, len(matrix[0])-1, len(matrix)-1, 0
	status := right
	res := make([]int, len(matrix)*len(matrix[0]))
	pos := 0
	for u <= d && l <= r {
		switch status {
		case right:
			for i := l; i <= r; i++ {
				res[pos] = matrix[u][i]
				pos++
			}
			u++
			status = down
		case down:
			for i := u; i <= d; i++ {
				res[pos] = matrix[i][r]
				pos++
			}
			r--
			status = left
		case left:
			for i := r; i >= l; i-- {
				res[pos] = matrix[d][i]
				pos++
			}
			d--
			status = up
		case up:
			for i := d; i >= u; i-- {
				res[pos] = matrix[i][l]
				pos++
			}
			l++
			status = right
		default:
			panic("wrong status")
		}
	}

	return res
}

func (sol *Solution) Title54() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	res := spiralOrder(matrix)
	fmt.Println(res)
}
