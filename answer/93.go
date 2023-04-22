package answer

import (
	"fmt"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string
	ip := make([]string, 4)
	var restoreIpAddress func(i, num int)
	restoreIpAddress = func(i, num int) {
		if len(s)-i > 3*(4-num) || len(s)-i < 4-num || num > 3 {
			return
		}
		if s[i] == '0' {
			ip[num] = "0"
			if num == 3 && i == len(s)-1 {
				res = append(res, strings.Join(ip, "."))
			} else {
				restoreIpAddress(i+1, num+1)
			}
			return
		}
		val := 0
		for j := i; j < len(s); j++ {
			val = val*10 + int(s[j]-'0')
			if val > 255 {
				break
			}
			if num == 3 {
				if j == len(s)-1 {
					ip[num] = s[i : j+1]
					res = append(res, strings.Join(ip, "."))
				}
				continue
			}
			ip[num] = s[i : j+1]
			restoreIpAddress(j+1, num+1)
		}
	}
	restoreIpAddress(0, 0)
	return res
}

func (sol *Solution) Title93() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
