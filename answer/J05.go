package answer

import (
	"fmt"
)

func replaceSpace(s string) string {
	cnt := 0
	for i := range s {
		if s[i] == ' ' {
			cnt++
		}
	}
	sb := make([]byte, len(s)+cnt*2)
	j := 0
	for i := range s {
		if s[i] == ' ' {
			sb[j] = '%'
			sb[j+1] = '2'
			sb[j+2] = '0'
			j += 3
		} else {
			sb[j] = s[i]
			j++
		}
	}

	return string(sb)
}

func (sol *Solution) TitleJ05() {
	s := "We are happy."
	result := replaceSpace(s)
	fmt.Println(result)
}
