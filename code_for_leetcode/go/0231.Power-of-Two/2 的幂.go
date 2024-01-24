package leetcode

// 解法一 二进制位操作法
//2的幂的二进制表示中只有一个位为1。
//通过 (num & (num - 1)) == 0 的判断，可以确定 num 的二进制表示中只有一位为1，即 num 是2的幂。
func isPowerOfTwo(num int) bool {
	// 判断 num 是否为正整数且二进制表示中只有一位为1
	return (num > 0 && ((num & (num - 1)) == 0))
}

// 解法二 数论
func isPowerOfTwo1(num int) bool {
	return num > 0 && (1073741824%num == 0)
}

// 解法四 循环
func isPowerOfTwo3(num int) bool {
	// 循环将 num 除以 2，判断是否能一直整除到 1
	for num >= 2 {
		if num%2 == 0 {
			num = num / 2
		} else {
			return false
		}
	}
	// 如果最终 num 等于 1，则是 2 的幂
	return num == 1
}
