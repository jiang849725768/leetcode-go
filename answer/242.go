package answer

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	smap := make(map[rune]int)
	for i := range s {
		smap[rune(s[i])]++
	}
	for i := range t {
		smap[rune(t[i])]--
	}
	for k := range smap {
		if smap[k] != 0 {
			return false
		}
	}

	return true
}

func (sol *Solution) Title242() {
	s := "anagram"
	t := "nagaram"
	result := isAnagram(s, t)
	fmt.Println(result)
	s = "rat"
	t = "car"
	result = isAnagram(s, t)
	fmt.Println(result)
}
