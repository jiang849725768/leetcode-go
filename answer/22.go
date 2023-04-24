package answer

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res := make([][]string, n+1)
	res[0] = []string{""}
	res[1] = []string{"()"}
	var generate func(i int)
	generate = func(i int) {
		if len(res[i]) > 0 {
			return
		}
		for p := 0; p < i; p++ {
			q := i - p - 1
			if len(res[p]) == 0 {
				generate(p)
			}
			if len(res[q]) == 0 {
				generate(q)
			}
			for _, r1 := range res[p] {
				for _, r2 := range res[q] {
					res[i] = append(res[i], "("+r1+")"+r2)
				}
			}
		}
	}
	generate(n)

	return res[n]
}

func (sol *Solution) Title22() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
