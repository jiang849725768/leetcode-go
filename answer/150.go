package answer

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	intStack := make([]int, 0)
	getInt := func() (int, int) {
		ls := len(intStack)
		if ls < 2 {
			panic("Expression is invalid.")
		}
		val1, val2 := intStack[ls-2], intStack[ls-1]
		intStack = intStack[:ls-2]
		return val1, val2
	}
	for i := range tokens {
		switch tokens[i] {
		case "+":
			val1, val2 := getInt()
			intStack = append(intStack, val1+val2)
		case "-":
			val1, val2 := getInt()
			intStack = append(intStack, val1-val2)
		case "*":
			val1, val2 := getInt()
			intStack = append(intStack, val1*val2)
		case "/":
			val1, val2 := getInt()
			if val2 == 0 {
				panic("divide number is 0")
			}
			intStack = append(intStack, val1/val2)
		default:
			val, _ := strconv.Atoi(tokens[i])
			intStack = append(intStack, val)
		}
	}
	if len(intStack) != 1 {
		panic("redundant value")
	}

	return intStack[0]
}

func (sol *Solution) Title150() {
	tokens := []string{"2", "1", "+", "3", "*"}
	result := evalRPN(tokens)
	fmt.Println(result)
	tokens = []string{"4", "13", "5", "/", "+"}
	result = evalRPN(tokens)
	fmt.Println(result)
	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	result = evalRPN(tokens)
	fmt.Println(result)
}
