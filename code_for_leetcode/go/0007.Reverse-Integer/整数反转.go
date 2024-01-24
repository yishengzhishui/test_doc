package leetcode

func reverse7(x int) int {
	// 初始化变量 tmp 用于存储反转后的整数
	tmp := 0

	// 循环直到 x 为 0
	for x != 0 {
		// 反转过程：每次取 x 的末尾数字，并加到 tmp 上
		tmp = tmp*10 + x%10

		// 去掉 x 的末尾数字
		x = x / 10
	}

	// 判断反转后的整数是否溢出，如果溢出则返回 0
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}

	// 返回反转后的整数
	return tmp
}
