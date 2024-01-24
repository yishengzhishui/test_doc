package leetcode

// DP dp[i][0,1][0,1,2],当前i层A的个数，结尾处L的个数
// 关注 A ，只关心整个字符串里 有 或 没有
//   若有，则下一层不能再加 A 了
//   若无，下一层可以加 A
// 关注 L ，只关心最后的两个字符是不是 LL
//   因为前面的要是有 LLL ，那早就应该剪枝了，不会再到这一层
//   如果前面有 LL ，但又被其他字符间隔，那也不影响当前的关注
// 初始状态
// dp[1][1][0] = 1; // A，
// dp[1][0][1] = 1; // L，
// dp[1][0][0] = 1; // P，

const mod = 1000000007

func checkRecord(n int) int {
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
		}
	}

	dp[1][1][0] = 1
	dp[1][0][1] = 1
	dp[1][0][0] = 1

	for i := 2; i <= n; i++ {
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % mod
		dp[i][1][0] = (dp[i-1][1][0] + dp[i-1][1][1] + dp[i-1][1][2] + dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % mod
		dp[i][0][1] = dp[i-1][0][0]
		dp[i][0][2] = dp[i-1][0][1]
		dp[i][1][1] = dp[i-1][1][0]
		dp[i][1][2] = dp[i-1][1][1]
	}

	return (dp[n][0][0] + dp[n][0][1] + dp[n][0][2] + dp[n][1][0] + dp[n][1][1] + dp[n][1][2]) % mod
}

// 空间压缩一下
func checkRecordV1(n int) int {
	MOD := 1000000007
	dp00, dp01, dp10, dp02, dp11, dp12 := 1, 1, 1, 0, 0, 0

	for i := 0; i < n-1; i++ {
		dp00, dp10, dp01, dp02, dp11, dp12 = (dp00 + dp01 + dp02) % MOD, (dp00 + dp01 + dp02 + dp10 + dp11 + dp12) % MOD, dp00, dp01, dp10, dp11
	}

	return (dp00 + dp01 + dp02 + dp10 + dp11 + dp12) % MOD
}
