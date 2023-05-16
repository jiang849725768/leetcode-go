package answer

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		x := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			y := num2[j] - '0'
			res[i+j+1] += int(x * y)
		}
	}
	for i := len(res) - 1; i > 0; i-- {
		res[i-1] += res[i] / 10
		res[i] %= 10
	}
	if res[0] == 0 {
		res = res[1:]
	}
	rres := ""
	for i := range res {
		rres += fmt.Sprintf("%d", res[i])
	}

	return rres
}

func (sol *Solution) Title43() {
	num1 := "123"
	num2 := "456"
	result := multiply(num1, num2)
	fmt.Printf("%s * %s = %s", num1, num2, result)
	fmt.Printf(", and correct answer is %d\n", 123*456)
	num1 = "2"
	num2 = "3"
	result = multiply(num1, num2)
	fmt.Printf("%s * %s = %s", num1, num2, result)
	fmt.Printf(", and correct answer is %d\n", 2*3)

}
