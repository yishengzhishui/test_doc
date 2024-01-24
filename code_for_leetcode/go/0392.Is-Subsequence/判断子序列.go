package leetcode

func isSubsequence1(s string, t string) bool {
	for len(s) > 0 && len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}

func isSubsequenceV1(s string, t string) bool {
	// 如果 s 为空字符串，则返回 true
	if len(s) == 0 {
		return true
	}

	// 如果 t 为空字符串，则返回 false
	if len(t) == 0 {
		return false
	}

	// 使用两个指针，一个用于遍历 s，一个用于遍历 t
	j := 0
	for i := 0; i < len(t); i++ {
		// 如果当前字符匹配，移动 s 的指针
		if t[i] == s[j] {
			j++
		}

		// 如果 s 的指针达到末尾，说明 s 是 t 的子序列
		if j == len(s) {
			return true
		}
	}

	// 遍历完整个 t 后仍然没有匹配到 s 的所有字符
	return false
}
