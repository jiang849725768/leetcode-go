package answer

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				island := [][]int{{i, j}}
				for len(island) != 0 {
					x, y := island[0][0], island[0][1]
					island = island[1:]
					if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
						grid[x][y] = '2'
						island = append(island, [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}...)
					}
				}
			}
		}
	}
	return res
}

func (sol *Solution) Title200() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
