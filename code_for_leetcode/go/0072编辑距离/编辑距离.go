package leetcode

// dp[i][j] 代表 word1[0:1]与 word2[0:j]之间的编辑距离
//if word1[i] == word2[j] { 不用任何编辑
//dp[i][j] = dp[i-1][j-1]
//}
//else {
//dp[i][j] = getMin(
//dp[i-1][j-1]+1,
//dp[i-1][j] +1,
//dp[i][j-1]+1)
//}

// dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作。
// 注意⚠️：第一行和第一列引入""，即word1和word2变成空需要的步数
//操作一：word1删除一个元素，那么就是以下标i - 2为结尾的word1 与 j-1为结尾的word2的最近编辑距离 再加上一个操作。
//即 dp[i][j] = dp[i - 1][j] + 1;
//
//操作二：word2删除一个元素，那么就是以下标i - 1为结尾的word1 与 j-2为结尾的word2的最近编辑距离 再加上一个操作。
//即 dp[i][j] = dp[i][j - 1] + 1;

// 例子
// 'acd', 'abc'
// dp[1][1] = 0; dp[1][2] = dp[1][1]+1(插入);dp[1][3] = dp[1][2]+1(插入);
// dp[2][1] =dp[1][1] = 0+1(删除) ;dp[2][2] =dp[1][1] = 0+1(替换) ;

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// 初始化动态规划数组 dp，大小为 (m+1) x (n+1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化边界条件
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// 动态规划递推
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			// 三种操作方式：插入、删除、替换
			dp[i][j] = getMin(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
		}
	}

	// 返回动态规划数组右下角的值，即为最小编辑距离
	return dp[m][n]
}

func getMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else {
		if b < c {
			return b
		}
		return c
	}
}
