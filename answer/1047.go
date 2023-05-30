package answer

import (
	"fmt"
)

func removeDuplicates(s string) string {
	byteStack := make([]byte, 0)
	for i := range s {
		lb := len(byteStack)
		if lb != 0 && byteStack[lb-1] == s[i] {
			byteStack = byteStack[:lb-1]
		} else {
			byteStack = append(byteStack, s[i])
		}
	}

	return string(byteStack)
}

func (sol *Solution) Title1047() {
	s := "abbaca"
	result := removeDuplicates(s)
	fmt.Println(result)
	s = "azxxzy"
	result = removeDuplicates(s)
	fmt.Println(result)
}
