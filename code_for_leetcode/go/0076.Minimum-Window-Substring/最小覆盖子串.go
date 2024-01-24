package leetcode

// 滑动窗口
// need中存放当前窗口下所需各元素的数量，如果need[char]<0，代表当前窗口这个元素有多余的(可以剔除)
// missing代表所需元素的总数量， missing==0时，当前窗口包含了所需的全部元素（可能会有多余的元素）
// need[char]>0代表当前窗口这个char是所需元素

// missing == 0时，开始收缩当前的窗口，从i到j中
// [a,b,c,a,b,d]=>[b,c,a,b,d]=>[c,a,b,d]

// minWindow 使用滑动窗口法找到字符串 s 中包含字符串 t 的最小子串
func minWindow(s string, t string) string {
	// 初始化一个 map 用于记录 t 中每个字符的出现次数
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	// missing 用于记录还缺少多少个 t 中的字符
	missing := len(t)
	// start 和 end 记录最小子串的起始和结束位置
	//i 和 j 是两个指针，用于表示滑动窗口的左右边界。
	start, end, i := 0, 0, 0

	// 遍历字符串 s
	for j := 1; j <= len(s); j++ {
		// 获取当前字符
		char := s[j-1]

		// 如果当前字符是 t 中所需的字符，则更新缺失字符的数量
		if need[char] > 0 {
			missing--
		}

		// 更新当前字符在所需字符中的数量(当前char在窗口中了)
		need[char]--

		// 如果所有字符都已经出现，则进入收缩阶段
		if missing == 0 {
			// 收缩窗口左边界，直到不能再收缩为止(need[char]<0，代表当前窗口这个元素有多余的(可以剔除))
			for i < j && need[s[i]] < 0 {
				need[s[i]]++
				i++
			}

			// 更新最小子串的起始和结束位置
			if end == 0 || j-i < end-start {
				start, end = i, j
			}

			// 恢复当前字符在所需字符中的数量，同时增加缺失字符的数量
			// 目的是继续向右移动（寻找下一个满足条件的窗口，为了找到 end-start最小的那个）
			need[s[i]]++
			missing++
			i++
		}
	}

	// 返回最小子串
	return s[start:end]
}
