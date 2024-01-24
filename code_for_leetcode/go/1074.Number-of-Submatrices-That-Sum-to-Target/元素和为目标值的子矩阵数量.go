package leetcode

// 363 和 560相结合
// 前缀和 + hash
// 转化问题
// 问题转化为 对于每一行 计算left到right的元素的和sums
// 随后对这一列中行元素和进行累加计算——前缀和

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	// 前缀和 + 哈希表
	m, n := len(matrix), len(matrix[0])
	result := 0

	// 循环所有的列
	for left := 0; left < n; left++ {
		sums := make([]int, m) //用于存储每一列的前缀和
		for right := left; right < n; right++ {
			// 计算当前列的前缀和
			for i := 0; i < m; i++ {
				sums[i] += matrix[i][right]
			}

			// 利用哈希表记录前缀和的出现次数
			dic := make(map[int]int)
			dic[0] = 1
			cur := 0
			for _, num := range sums {
				cur += num
				result += dic[cur-target] // 查找哈希表中是否存在 cur-target，存在则累加到结果中
				dic[cur]++                // 将当前前缀和 cur 记录到哈希表中
			}
		}
	}
	return result
}
