package answer

import (
	"fmt"
)

//func strStr(haystack string, needle string) int {
//	for i := range haystack {
//		if i+len(needle) > len(haystack) {
//			break
//		}
//		if haystack[i:i+len(needle)] == needle {
//			return i
//		}
//	}
//
//	return -1
//}

func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	j := 0
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	j = 0
	for i := range haystack {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				return i - len(needle) + 1
			}
		}
	}

	return -1
}

func (sol *Solution) Title28() {
	haystack := "sadbutsad"
	needle := "sad"
	result := strStr(haystack, needle)
	fmt.Println(result)
}
