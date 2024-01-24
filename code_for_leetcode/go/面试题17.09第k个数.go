package leetcode

// 动态规划
// dp[i] := getMin(dp[p3]*3, getMin(dp[p5]*5, dp[p7]*7))
func getKthMagicNumber(k int) int {
	// 初始化三个指针和存储魔法数的数组
	p3, p5, p7 := 0, 0, 0
	res := []int{1}

	// 从 1 开始迭代生成魔法数，直到生成第 k 个魔法数
	for i := 1; i < k; i++ {
		// 计算当前位置的最小魔法数
		minRes := getMin(res[p3]*3, getMin(res[p5]*5, res[p7]*7))

		// 根据最小魔法数更新相应的指针,这里用switch会有问题
		// 同时满足 res[p3]*3 res[p5]*5, p3 p5都要加1
		if minRes == res[p3]*3 {
			p3++
		}

		if minRes == res[p5]*5 {
			p5++
		}

		if minRes == res[p7]*7 {
			p7++
		}

		// 将最小魔法数添加到数组中
		res = append(res, minRes)
	}

	// 返回生成的第 k 个魔法数
	return res[len(res)-1]
}

// 获取两个数的最小值的辅助函数
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
