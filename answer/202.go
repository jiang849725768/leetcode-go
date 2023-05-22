package answer

import (
	"fmt"
)

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		if m[n] {
			return false
		}
		m[n] = true
		n2 := 0
		for n > 0 {
			n2 += (n % 10) * (n % 10)
			n /= 10
		}
		n = n2
	}

	return true
}

func (sol *Solution) Title202() {
	n := 19
	result := isHappy(n)
	fmt.Println(result)
	n = 2
	result = isHappy(n)
	fmt.Println(result)
}

// 2 4 16 37 58 89 145 42 20 4
// 3 9 81 65 61 37 58 89 145 42 20 4
