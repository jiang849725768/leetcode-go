package answer

import (
	"fmt"
)

//func repeatedSubstringPattern(s string) bool {
//	// 从1开始遍历,子串至少重复一次所以长度小于等于一半
//	for i := 1; i*2 <= len(s); i++ {
//		// 寻找首位与主字符串相同且能整除主字符串的子字符串
//		if s[i] != s[0] || len(s)%i != 0 {
//			continue
//		}
//		subLen := i
//		// 检查子字符串是否能重复组成主字符串
//		for j := 0; j < len(s); j += subLen {
//			if s[:subLen] != s[j:j+subLen] {
//				break
//			}
//			if j+subLen == len(s) {
//				return true
//			}
//		}
//	}
//
//	return false
//}

func repeatedSubstringPattern(s string) bool {
	next := make([]int, len(s))
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	//if next[len(s)-1] != 0 && len(s)%(len(s)-next[len(s)-1]) == 0 {
	//	return true
	//}
	j = 0
	for i := 1; i < 2*len(s)-1; i++ {
		for j > 0 && s[i%len(s)] != s[j] {
			j = next[j-1]
		}
		if s[i%len(s)] == s[j] {
			j++
			if j == len(s) {
				return true
			}
		}
	}

	return false
}

func (sol *Solution) Title459() {
	s := "abab"
	result := repeatedSubstringPattern(s)
	fmt.Println(result)
	s = "aba"
	result = repeatedSubstringPattern(s)
	fmt.Println(result)
	s = "abcabcabcabc"
	result = repeatedSubstringPattern(s)
	fmt.Println(result)
}
