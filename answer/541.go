package answer

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	res := []byte(s)
	setNum := len(s) / (2 * k)
	restNum := len(s) % (2 * k)
	for i := 0; i < setNum; i++ {
		for j := 0; j < k/2; j++ {
			res[i*2*k+j], res[i*2*k+k-j-1] = res[i*2*k+k-j-1], res[i*2*k+j]
		}
	}
	i := setNum
	m := k
	if restNum < k {
		m = restNum
	}
	for j := 0; j < m/2; j++ {
		res[i*2*k+j], res[i*2*k+m-j-1] = res[i*2*k+m-j-1], res[i*2*k+j]
	}

	return string(res)
}

func (sol *Solution) Title541() {
	//s := "abcdefg"
	//k := 2
	//result := reverseStr(s, k)
	//fmt.Println(result)
	//s = "abcd"
	//k = 2
	//result = reverseStr(s, k)
	//fmt.Println(result)
	s := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	k := 39
	result := reverseStr(s, k)
	fmt.Println(result)
}
