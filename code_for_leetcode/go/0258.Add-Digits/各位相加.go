package leetcode

//数根是一个整数的各个位数相加，直到结果变成个位数为止。

func addDigits(num int) int {
	// 当数字大于 9 时，继续计算数根
	for num > 9 {
		// cur 用于存储当前数字的各位数之和
		cur := 0
		// 遍历当前数字的各位数
		for num != 0 {
			// 将当前数字的个位数加到 cur 中
			cur += num % 10
			// 去掉当前数字的个位数
			num /= 10
		}
		// 将 cur 赋值给 num，继续下一轮的计算
		num = cur
	}
	// 返回最终计算结果，即数根
	return num
}

func addDigitsV1(num int) int {
	// 如果 num 为 0，则直接返回 0
	if num == 0 {
		return 0
	} else {
		// 如果 num 不为 0，则根据规律计算结果
		return (num-1)%9 + 1
	}
}
