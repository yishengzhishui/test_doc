package leetcode

// 寻找第一个不重复字符函数
func dismantlingAction(arr string) byte {
	// 使用 map 存储字符是否重复的信息
	dic := make(map[rune]int)

	// 遍历字符串 s 中的每个字符，记录是否重复
	// 如果 dic 中不存在该字符，将该字符添加到 dic 中，并设置对应的值为 true，表示该字符第一次出现。
	for _, char := range arr {
		dic[char]++
	}

	// 再次遍历字符串 s，找到第一个不重复的字符并返回
	for _, char := range arr {
		if dic[char] == 1 {
			return byte(char)
		}
	}

	// 如果没有不重复的字符，返回空格字符
	return ' '
}
// 优化一下空间
func dismantlingActionV1(arr string) byte {
	hash := [26]int{}
	for i := 0; i < len(arr); i++ {
		hash[arr[i]-'a']++
	}
	for i := 0; i < len(arr); i++ {
		if hash[arr[i]-'a'] == 1 {
			return arr[i]
		}
	}
	return ' '
}
