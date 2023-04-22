package answer

func addStrings(num1 string, num2 string) string {
	lnum := len(num1)
	if lnum < len(num2) {
		lnum = len(num2)
	}
	res := make([]byte, lnum+1)
	carry := 0
	for i := 0; i < lnum; i++ {
		sum := carry
		if i < len(num1) {
			sum += int(num1[len(num1)-1-i] - '0')
		}
		if i < len(num2) {
			sum += int(num2[len(num2)-1-i] - '0')
		}
		res[lnum-i] = byte(sum%10 + '0')
		carry = sum / 10
	}
	if carry > 0 {
		res[0] = byte(carry + '0')
	} else {
		res = res[1:]
	}

	return string(res)
}

func (sol *Solution) Title415() {
	num1 := "123"
	num2 := "456"
	println(addStrings(num1, num2))
}
