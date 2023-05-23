package utils

// kmp 算法获取字符串的前缀表
func kmp(s string) []int {
	next := make([]int, len(s))
	next[0] = 0
	j := 0
	for i := 1; i < len(s); i++ {
		// j 回退到上一个相同的位置
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		// 如果相同，j 后移
		if s[i] == s[j] {
			j++
		}
		// j 为当前位置的最长相同前后缀长度
		next[i] = j
	}

	return next
}
