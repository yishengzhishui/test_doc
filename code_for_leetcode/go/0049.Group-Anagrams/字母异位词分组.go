package leetcode

// 按照计数分类，使用计数数组
// 核心思想是使用一个长度为 26 的整数数组来表示每个字符串中每个字母的出现次数。
//然后，将这个数组作为键存储在 map 中，
//值是具有相同字母异序词的字符串切片。
//最后，将 map 中的值提取出来，形成最终的结果。
func groupAnagramsV1(strs []string) [][]string {
	// 创建一个 map，键是长度为 26 的整数数组，值是字符串切片
	dic := make(map[[26]int][]string)

	// 遍历输入的字符串数组
	for _, s := range strs {
		// 初始化一个长度为 26 的整数数组，用于统计当前字符串中每个字母出现的次数
		count := [26]int{}
		// 遍历当前字符串中的每个字符
		for _, char := range s {
			// 计算字符相对于 'a' 的偏移，将出现的次数记录在 count 数组中
			count[int(char-'a')]++
		}

		// 使用 count 数组作为键，将当前字符串添加到对应的值（字符串切片）中
		key := count
		dic[key] = append(dic[key], s)
	}

	// 初始化一个字符串切片用于存储最终的结果
	result := make([][]string, 0, len(dic))
	// 遍历 map 中的值，将每个字符串切片添加到结果中
	for _, value := range dic {
		result = append(result, value)
	}

	// 返回最终的结果
	return result
}

func groupAnagrams(strs []string) [][]string {
	// 编码到分组的映射
	codeToGroup := make(map[string][]string)
	for _, s := range strs {
		// 对字符串进行编码
		code := encode(s)
		// 把编码相同的字符串放在一起
		codeToGroup[code] = append(codeToGroup[code], s)
	}

	// 获取结果
	res := make([][]string, 0, len(codeToGroup))
	for _, group := range codeToGroup {
		res = append(res, group)
	}

	return res
}

func encode(s string) string {
	// 创建一个长度为26的字节切片，用于存储每个字母的出现次数
	count := make([]byte, 26)

	// 遍历输入字符串 s
	for i := 0; i < len(s); i++ {
		gap := s[i] - 'a' // 计算当前字符与字母 'a' 的差距，得到字母在数组中的索引
		count[gap]++      // 将对应字母的计数加一
	}

	// 将字母出现次数组成的字节切片转换为字符串并返回
	return string(count)
}
