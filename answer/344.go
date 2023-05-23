package answer

import (
	"fmt"
)

func reverseString(s []byte) {
	slen := len(s)
	for i := 0; i < slen/2; i++ {
		s[i], s[slen-i-1] = s[slen-i-1], s[i]
	}
}

func (sol *Solution) Title344() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
	s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	fmt.Println(s)
}
