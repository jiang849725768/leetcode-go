package answer

import (
	"fmt"
)

func canConstruct(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for i := range ransomNote {
		cnt[ransomNote[i]-'a']++
	}
	for i := range magazine {
		if cnt[magazine[i]-'a'] > 0 {
			cnt[magazine[i]-'a']--
		}
	}
	for i := range cnt {
		if cnt[i] > 0 {
			return false
		}
	}
	return true
}

func (sol *Solution) Title383() {
	ransomNote := "a"
	magazine := "b"
	result := canConstruct(ransomNote, magazine)
	fmt.Println(result)
	ransomNote = "aa"
	magazine = "ab"
	result = canConstruct(ransomNote, magazine)
	fmt.Println(result)
	ransomNote = "aa"
	magazine = "aab"
	result = canConstruct(ransomNote, magazine)
	fmt.Println(result)
}
