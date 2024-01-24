package leetcode

// 路径加密函数
func pathEncryption(path string) string {
	var res []rune

	// 遍历字符串 path 中的每个字符
	for _, char := range path {
		// 如果当前字符是 '.'，则将空格追加到结果字符串
		if char == '.' {
			res = append(res, ' ')
		} else {
			// 如果当前字符不是 '.'，直接将字符追加到结果字符串
			res = append(res, char)
		}
	}
	// 将字符切片转换为字符串并返回
	return string(res)
}
