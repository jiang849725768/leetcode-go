package answer

import (
	"fmt"
)

func mySqrt(x int) int {
	res := x
	for {
		if res*res <= x && (res+1)*(res+1) > x {
			break
		}
		res = (res + x/res) / 2
	}

	return res
}

func (sol *Solution) Title69() {
	x := 8
	fmt.Println(mySqrt(x))
}
