package main

import (
	"fmt"
	"reflect"

	"leetcode/answer"
)

func main() {
	var titleNum int
	if _, err := fmt.Scanf("%d", &titleNum); err != nil {
		fmt.Println("Title num must be a integer")
	}
	titleName := fmt.Sprintf("Title%d", titleNum)

	sol := answer.New()
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
