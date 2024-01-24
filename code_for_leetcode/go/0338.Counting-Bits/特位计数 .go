package leetcode

// countBits 计算从 0 到 num 的每个数的二进制表示中 1 的个数
//i&(i-1) 可以将i最右边的1变成0

// 当 i = 5 时，i 的二进制表示为 101，i-1 的二进制表示为 100。因此，i&(i-1) 的结果为 100，将 i 中最右边的 1 变为 0。
//当 i = 8 时，i 的二进制表示为 1000，i-1 的二进制表示为 111。因此，i&(i-1) 的结果为 0000，将 i 中最右边的 1 变为 0。
func countBits(num int) []int {
	// 创建一个长度为 num+1 的整数切片 bits，用于存储结果
	bits := make([]int, num+1)

	// 循环遍历每个数 i，计算其二进制表示中 1 的个数
	for i := 1; i <= num; i++ {
		// 对于当前数 i，通过 i&(i-1) 获取 i 的二进制表示中最右边的 1 及其右边的部分，
		// 然后通过 bits[i&(i-1)] 获取该右边部分中 1 的个数，
		// 最后加上 1 就是当前数 i 的二进制表示中 1 的个数。
		bits[i] = bits[i&(i-1)] + 1
	}

	// 返回存储结果的整数切片 bits
	return bits
}

// countBits 计算从 0 到 num 的每个数的二进制表示中 1 的个数
// 奇数：二进制表示中，奇数一定比前面那个偶数多一个 1，因为多的就是最低位的 1。
// 偶数：二进制表示中，偶数中 1 的个数一定和除以 2 之后的那个数一样多。因为最低位是 0，除以 2 就是右移一位，也就是把那个 0 抹掉而已，所以 1 的个数是不变的。
// x / 2 is x >> 1 and x % 2 is x & 1
func countBitsV1(num int) []int {
	// 创建一个长度为 num+1 的整数切片 result，用于存储结果
	result := make([]int, num+1)
	result[0] = 0

	// 循环遍历每个数 i，计算其二进制表示中 1 的个数
	for i := 1; i <= num; i++ {
		// 如果 i 为奇数，则 result[i] = result[i-1] + 1
		if i%2 == 1 {
			result[i] = result[i-1] + 1
		} else {
			// 如果 i 为偶数，则 result[i] = result[i/2]
			result[i] = result[i/2]
		}
	}

	// 返回存储结果的整数切片 result
	return result
}
