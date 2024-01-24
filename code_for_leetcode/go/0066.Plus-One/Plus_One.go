package leetcode

func plusOne(digits []int) []int {
	// 分类: 末尾为9 和其他 +1
	// 如果数组中全是9，才会走到最后的return

	// 从数组末尾开始循环
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果当前元素不是9，直接加1，并返回结果数组
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		// 如果当前元素是9，将其变为0，继续处理前一位
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
