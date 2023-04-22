package answer

func myAtoi(s string) int {
	res := 0
	flag := 1
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) {
		if s[i] == '-' {
			flag = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int(s[i]-'0')
		if res > 1<<31-1 {
			if flag == 1 {
				return 1<<31 - 1
			}
			return -1 << 31
		}
		i++
	}

	return flag * res
}

func (sol *Solution) Title8() {
	// s := "  -42"
	s := "2148943579854789"
	println(myAtoi(s))
}
