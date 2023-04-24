package answer

import (
	"fmt"
)

func minWindow(s string, t string) string {
	l, r := 0, 0                       // 滑动窗口的左右边界
	slen, resl, resr := len(s)+1, 0, 0 // 最小窗口的长度，最小窗口的左右边界
	tmap := make(map[byte]int)         // t中字符的计数
	count := 0                         // 需要的总字符数
	for i := range t {                 // 初始化tmap和count
		tmap[t[i]]++
		count++
	}
	for i := range s { // 滑动窗口
		if _, ok := tmap[s[i]]; ok { // s[i]在t中
			tmap[s[i]]--
			if tmap[s[i]] >= 0 { // s[i]在t中的个数还没用完
				count--
			}
		}
		if count == 0 { // 窗口中包含t中所有字符
			r = i
			if r-l+1 < slen { // 更新最小窗口
				slen = r - l + 1
				resl, resr = l, r
			}
		}
		for count == 0 { // 窗口中包含t中所有字符的情况下，左边界右移
			if _, ok := tmap[s[l]]; ok {
				tmap[s[l]]++
				if tmap[s[l]] > 0 { // l为滑动窗口的临界左边界
					if r-l+1 < slen {
						slen = r - l + 1
						resl, resr = l, r
					}
					count++
				}
			}
			l++
		}
	}
	if slen == len(s)+1 { // 没有找到
		return ""
	}

	return s[resl : resr+1]
}

func (sol *Solution) Title76() {
	//s := "cabwefgewcwaefgcf"
	//t := "cae"
	s := "a"
	t := "aa"
	fmt.Println(minWindow(s, t))
}
