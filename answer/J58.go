package answer

import (
	"fmt"
)

func reverseLeftWords(s string, n int) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s)-n; i++ {
		res[i] = s[i+n]
	}
	for i := len(s) - n; i < len(s); i++ {
		res[i] = s[i-len(s)+n]
	}

	return string(res)
}

func (sol *Solution) TitleJ58() {
	s := "abcdefg"
	k := 2
	result := reverseLeftWords(s, k)
	fmt.Println(result)
	s = "lrloseumgh"
	k = 6
	result = reverseLeftWords(s, k)
	fmt.Println(result)
}
