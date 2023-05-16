package main

import (
	"fmt"
	"reflect"

	"leetcode/answer"
)

func main() {
	var titleNum string
	if _, err := fmt.Scanf("%s", &titleNum); err != nil {
		panic("Title num must be a string")
	}
	titleName := fmt.Sprintf("Title%s", titleNum)

	sol := answer.NewSolution()
	// 获取包函数反射值
	pkg := reflect.ValueOf(sol)
	function := pkg.MethodByName(titleName)

	// 调用函数
	if function.IsValid() && function.Kind() == reflect.Func {
		fmt.Printf("Run solution of %s\n", titleName)
		function.Call(nil)
	} else {
		fmt.Println("Invalid function name")
	}
}
