package leetcode

func subtractProductAndSum(n int) int {
	// 初始化和与乘积为0和1
	sum, product := 0, 1

	// 循环，对数字的每一位进行处理
	for ; n > 0; n /= 10 {
		// 取数字的个位数
		x := n % 10
		// 计算各位数的和
		sum += x
		// 计算各位数的乘积
		product *= x
	}

	// 返回各位数的乘积与和之间的差值
	return product - sum
}
