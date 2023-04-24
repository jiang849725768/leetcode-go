package answer

import (
	"fmt"
	"sort"
)

func merge2(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := 0
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][1] > intervals[i+1][1] {
				intervals[i+1][1] = intervals[i][1]
			}
			if intervals[i][0] < intervals[i+1][0] {
				intervals[i+1][0] = intervals[i][0]
			}
			intervals[i], intervals[start] = intervals[start], intervals[i]
			start++
		}
	}

	return intervals[start:]
}

func (sol *Solution) Title56() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge2(intervals))
}
