package leetcode

//取当前 n 的最后一位：n & 1
//将最后一位移动到对应位置，第一次为 31 位，第二次是 30 位，即：31、30、29... 1、0
func reverseBits(num uint32) uint32 {
	// 用于存储反转后的结果
	var res uint32
	// 迭代32次，对num的二进制表示进行逐位处理
	for i := 0; i < 32; i++ {
		// 将res左移一位，然后将num的最后一位加到res的最右侧
		res = res<<1 | num&1
		// 将num右移一位，以处理下一位
		num >>= 1
	}
	// 返回反转后的结果
	return res
}
