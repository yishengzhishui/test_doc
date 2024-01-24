package leetcode

// 与 560 和为 K 的子数组 相似 dic[cur-k]
// 前缀和+hash
// 当前的前缀和cur和之前某个前缀和之间相差若干个K

// 考虑子数组 A[i:j] 的和，其中 0 <= i <= j < len(A)。
//如果 sum[i:j] 对 K 取余为零，即 (sum[j] - sum[i]) % K == 0，那么 sum[j] % K == sum[i] % K。
// 如果两个前缀和对 K 取余相同，那么它们之间的子数组的和就是 K 的倍数。

func subarraysDivByK(A []int, K int) int {
	// 初始化变量
	dic := make(map[int]int)
	dic[0] = 1
	result, cur := 0, 0

	// 遍历数组
	for _, num := range A {
		// 更新当前前缀和
		cur += num

		// 计算当前前缀和对 K 取余的值
		value := (cur%K + K) % K

		// 累加符合条件的子数组个数
		result += dic[value]

		// 更新 map 中当前前缀和的次数
		dic[value]++
	}

	// 返回结果
	return result
}
