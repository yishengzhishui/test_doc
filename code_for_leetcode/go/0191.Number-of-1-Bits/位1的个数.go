package leetcode

import "math/bits"

// 解法一
func hammingWeight(num uint32) int {
	// 使用标准库 bits 包的 OnesCount 函数直接统计二进制中1的个数
	return bits.OnesCount(uint(num))
}

// 解法二
func hammingWeight1(num uint32) int {
	// 初始化计数器
	count := 0
	// 当num不等于0时，执行循环
	for num != 0 {
		// 每次将 num 与 (num - 1) 进行与操作，将二进制表示中的最低位的1变为0
		num = num & (num - 1)
		// 计数器加1
		count++
	}
	// 返回计数器的值，即二进制中1的个数
	return count
}
