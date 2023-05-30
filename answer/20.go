package answer

import (
	"fmt"
)

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for _, bt := range s {
		switch bt {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack = append(stack, byte(bt))
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

func (sol *Solution) Title20() {
	s := "()[]{{}"
	fmt.Println(isValid(s))
}
