package leetcode

// RomanToInt 是将罗马数字转换为整数的函数
func RomanToInt(s string) int {
	// 定义罗马数字与整数的映射关系
	romanNum := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// 初始化总和
	total := 0

	// 遍历字符串，计算总和
	for i := 0; i < len(s)-1; i++ {
		// 如果当前字符对应的数值大于等于下一个字符对应的数值，加上当前字符对应的数值
		if romanNum[s[i]] >= romanNum[s[i+1]] {
			total += romanNum[s[i]]
		} else {
			// 如果当前字符对应的数值小于下一个字符对应的数值，减去当前字符对应的数值
			total -= romanNum[s[i]]
		}
	}

	// 加上最后一个字符对应的数值
	return total + romanNum[s[len(s)-1]]
}
