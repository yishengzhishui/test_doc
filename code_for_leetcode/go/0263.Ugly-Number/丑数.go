package leetcode

// 2, 3, 5 轮流相除，最后的结果不为1，则false
func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for _, i := range []int{2, 3, 5} {
		for num%i == 0 {
			num /= i
		}
	}
	return num == 1
}
