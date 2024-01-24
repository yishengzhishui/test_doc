package leetcode

func longestCommonPrefix(strs []string) string {
	// 将第一个字符串作为初始前缀
	prefix := strs[0]

	// 从第二个字符串开始遍历
	for index, str := range strs {
		if index == 0 {
			continue
		}
		// 遍历当前前缀的每个字符
		for j := 0; j < len(prefix); j++ {
			//len(str) <= j：检查当前遍历的字符串 str 的长度是否小于等于当前位置 j。
			//如果是，说明 str 已经遍历结束，不再具有前缀的这个位置，因此不需要继续比较，可以截断前缀。

			//str[j] != prefix[j]：检查当前遍历的字符串 str 在位置 j 的字符是否不等于前缀在相同位置 j 的字符。
			//如果是，说明当前位置的字符不匹配，因此需要截断前缀到这个位置。
			if len(str) <= j || str[j] != prefix[j] {
				// 如果是，截取前缀，保留相同的部分
				prefix = prefix[0:j]
				break
			}
		}
	}

	// 返回最终的前缀
	return prefix
}
