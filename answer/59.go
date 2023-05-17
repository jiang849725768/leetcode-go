package answer

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

func (sol *Solution) Title59() {
	n := 6
	result := generateMatrix(n)
	fmt.Println(result)
	n = 1
	result = generateMatrix(n)
	fmt.Println(result)
}
