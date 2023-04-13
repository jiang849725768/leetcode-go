package answer

import (
	"fmt"
)

func longestPalindrome(s string) string {
	var res string
	checkPalindrome := func(left, right int) {
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		if right-left-1 > len(res) {
			res = s[left+1 : right]
		}
	}

	for i := range s {
		if i > 0 && s[i] == s[i-1] {
			checkPalindrome(i-1, i)
		}
		if i > 1 && s[i] == s[i-2] {
			checkPalindrome(i-2, i)
		}
	}

	if len(res) != 0 {
		return res
	}
	return s[0:1]
}

func (sol *Solution) Title5() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
