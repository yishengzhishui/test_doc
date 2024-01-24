package leetcode

// 解法一
func isAnagram(s string, t string) bool {
	// isAnagram 函数用于判断两个字符串是否是字母异位词（由相同的字母重新排列而成，但顺序不同）。
	// 利用数组 alphabet 记录每个字母在字符串中出现的次数，遍历字符串 s 时，将对应字母的计数加一；
	// 遍历字符串 t 时，将对应字母的计数减一。
	// 如果两个字符串是字母异位词，那么最终 alphabet 中所有字母的计数都应该为零。
	// 最后，遍历 alphabet 数组，如果存在计数不为零的字母，说明两个字符串不是字母异位词，返回 false；
	// 否则，返回 true。

	alphabet := make([]int, 26) // 用于记录字母出现次数的数组，初始化为全零
	sBytes := []byte(s)         // 将字符串 s 转换为字节数组
	tBytes := []byte(t)         // 将字符串 t 转换为字节数组

	// 如果字符串 s 和 t 的长度不相等，说明它们不可能是字母异位词，直接返回 false
	if len(sBytes) != len(tBytes) {
		return false
	}

	// 遍历字符串 s，更新字母计数数组 alphabet
	// 字符 sBytes[i] 转换为相对于小写字母 'a' 的偏移量。
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}

	// 遍历字符串 t，更新字母计数数组 alphabet
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}

	// 遍历字母计数数组 alphabet，如果存在计数不为零的字母，返回 false
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}

	return true // 返回 true，表示两个字符串是字母异位词
}

// 解法二
func isAnagram1(s string, t string) bool {
	hash := map[rune]int{}
	for _, value := range s {
		hash[value]++
	}
	for _, value := range t {
		hash[value]--
	}
	for _, value := range hash {
		if value != 0 {
			return false
		}
	}
	return true
}
