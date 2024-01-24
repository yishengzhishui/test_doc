package leetcode

// dp[i][j] 表示子串 s[i..j] 是否为回文子串，
// 这里子串 s[i..j] 定义为左闭右闭区间，可以取到 s[i] 和 s[j]。
// 空的或长度为1的子串是回文串
// dp[i][j] = dp[i+1][j-1](去掉首尾后的子串) and  s[i] == s[j]
// j-i<2 即 i，j相同位置或两者间隔1个时；只要s[i] == s[j]肯定是回文串，
// 否则就要依据dp[i + 1][j - 1])进行判定了
func longestPalindromeV1(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}
	result := "" // 用于存储最长回文子串的结果
	// 创建一个二维布尔数组 dp，用于记录子串是否为回文
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
	}

	// 遍历字符串 s 中的所有可能的子串
	for j := 0; j < size; j++ {
		for i := 0; i <= j; i++ {
			// 判断当前子串是否为回文，首尾字符相等且去掉首尾字符后的子串也是回文
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])

			// 如果当前子串是回文 且长度（j-i+1）大于之前记录的最长回文串的长度，
			//则更新最长回文串
			if dp[i][j] && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}

	return result
}

func longestPalindrome(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}

	result := "" // 用于存储最长回文子串的变量

	// 遍历字符串 s 中的所有可能的中心位置
	// 为了处理奇偶性的问题
	// 如果i是奇数， 如i=1，left=0，right=0+1; 如i=3，left=1，right=1+1;
	// 如果i是偶数， 如i=2，left=1，right=1+0; 如i=4，left=2，right=2+0;
	for i := 0; i < 2*size-1; i++ {
		left := i / 2       // 当前中心位置的左边界
		right := left + i%2 // 当前中心位置的右边界

		// 在当前中心位置向两边扩展，直到找到不满足回文条件的位置
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		// 如果当前找到的回文子串长度大于之前记录的最长回文串长度，则更新最长回文串
		//left 和 right 分别表示回文子串的左右边界，它们指向的位置不属于回文子串。
		//因此，通过计算 right - left - 1，我们得到了回文子串的实际长度。
		if len(result) < right-left-1 {
			result = s[left+1 : right]
		}
	}

	return result
}
