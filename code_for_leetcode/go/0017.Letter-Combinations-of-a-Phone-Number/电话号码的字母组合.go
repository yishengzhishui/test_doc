package leetcode

//回溯
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

// letterCombinations 函数用于生成电话号码的所有可能的字母组合
func letterCombinations(digits string) []string {
	// 存储最终结果的切片
	var result []string

	// 定义递归辅助函数
	var helper func(cur string, index int)
	helper = func(cur string, index int) {
		// 如果当前组合长度等于数字串的长度，将其加入结果切片
		// cur 是当前的字符串
		if index == len(digits) {
			result = append(result, cur)
			return
		}

		// 获取当前数字对应的字母集合
		// digits[index] 直接返回的是一个 Unicode 字符的整数表示
		letters := dict[string(digits[index])]

		// 遍历字母集合，递归调用 helper 函数
		for _, c := range letters {
			helper(cur+c, index+1)
		}
	}

	// 如果数字串为空，直接返回空切片
	if len(digits) == 0 {
		return result
	}

	// 调用 helper 函数，开始递归生成组合
	helper("", 0)
	return result
}
