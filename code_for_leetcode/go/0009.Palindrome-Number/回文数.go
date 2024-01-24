package leetcode

import "strconv"

// IsPalindrome 是判断整数是否为回文数的函数
func IsPalindrome(x int) bool {
	// 判断特殊情况：负数或以0结尾的非零整数不是回文数
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	revert := 0

	// 反转一半的数字
	for x > revert {
		pop := x % 10
		x = (x - pop) / 10

		revert = revert*10 + pop
	}

	// 判断是否为回文数
	// x == (revert-revert%10)/10 这一部分是在处理数字位数为奇数的情况。
	return revert == x || x == (revert-revert%10)/10
}

// 解法二 数字转字符串

// isPalindrome 是判断整数是否为回文数的函数
func isPalindromeV1(x int) bool {
	// 判断是否为负数，负数不是回文数
	if x < 0 {
		return false
	}

	// 如果 x 是个位数，则一定是回文数
	if x < 10 {
		return true
	}

	// 将整数转换为字符串，以便进行字符比较
	s := strconv.Itoa(x)

	// 获取字符串的长度
	length := len(s)

	// 循环比较字符串的前半部分和后半部分
	for i := 0; i <= length/2; i++ {
		// 如果对应位置的字符不相等，则不是回文数
		if s[i] != s[length-1-i] {
			return false
		}
	}

	// 如果所有字符比较都相等，那么是回文数
	return true
}
