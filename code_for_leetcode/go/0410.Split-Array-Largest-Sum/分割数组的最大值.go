package leetcode

// dp[i][j] 表示将数组的前 i个数分割为 j段,所能得到的最大连续子数组和的最小值
// 枚举k，前 k个数被分割为 j-1段，而第 k+1 到第 i 个数为第 j段。
// dp[i][j] = min(dp[i][j], max(dp[k][j-1], sum(nums[k+1:i+1])))
// 合法的状态必须保证 i>=j,所以dp初始化为inf
// 另外dp[0][0] = 0
//func splitArray1(nums []int, k int) int {
//	// 获取数组大小
//	size := len(nums)
//
//	// 初始化 dp 数组
//	dp := make([][]int, size+1)
//	for i := range dp {
//		dp[i] = make([]int, k+1)
//		for j := range dp[i] {
//			// 初始化为 int 最大值
//			dp[i][j] = int(^uint(0) >> 1)
//		}
//	}
//
//	// 初始条件：将空数组分割为 0 段得到的最小和为 0
//	dp[0][0] = 0
//
//	// 计算前缀和数组 sub
//	sub := make([]int, size+1)
//	for i, num := range nums {
//		sub[i+1] = sub[i] + num
//	}
//
//	// 动态规划过程
//	for i := 1; i <= size; i++ {
//		for j := 1; j <= getMin(i, k); j++ {
//			for k := 0; k < i; k++ {
//				// 更新状态转移方程：dp[i][j] = min(dp[i][j], max(dp[k][j-1], sub[i]-sub[k]))
//				dp[i][j] = getMin(dp[i][j], getMax(dp[k][j-1], sub[i]-sub[k]))
//			}
//		}
//	}
//
//	return dp[size][k]
//}

//// getMin 函数用于返回两个整数的较小值
//func getMin(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//// getMax 函数用于返回两个整数的较大值
//func getMax(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 我先猜一个mid值，然后遍历数组划分，使每个子数组和都最接近mid（贪心地逼近mid），这样我得到的子数组数一定最少;
// 如果即使这样子数组数量仍旧多于m个，那么明显说明我mid猜小了，因此 left = mid + 1;
// 而如果得到的子数组数量小于等于m个，那么我可能会想，mid是不是有可能更小？值得一试。因此 right = mid;
// 模拟分割 cur 表示已经分割出的子数组的数量（包括当前子数组）
// total表示当前分割子数组的和
// total + num > x表示需要再分割一个数组，并且当前取值作为新子数组的开头

func splitArray(nums []int, k int) int {
	// 定义一个函数 check，用于检查是否可以将数组分割为最多 k 段
	check := func(value int) bool {
		total, cur := 0, 1
		// 遍历数组 nums
		for _, num := range nums {
			// 如果当前段的和加上当前数字超过了目标值 value
			if total+num > value {
				// 将当前数字作为新的一段的起点
				total = num
				// 段数加一
				cur++
			} else {
				// 否则，将当前数字加入当前段
				total += num
			}
		}
		// 判断当前的段数是否小于等于 k
		return cur <= k
	}

	// 初始化二分查找的左右边界
	left, right := getMax(nums), sum(nums)

	// 二分查找
	for left < right {
		// 取中间值
		mid := (left + right) / 2
		// 如果 check(mid) 返回 true，说明当前 mid 可以满足条件，缩小右边界
		if check(mid) {
			right = mid
		} else {
			// 否则，增加左边界
			left = mid + 1
		}
	}

	// 最终返回 left，即为结果
	return left
}

// getMax 函数用于返回数组中的最大值
func getMax(nums []int) int {
	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

// sum 函数用于返回数组的总和
func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
