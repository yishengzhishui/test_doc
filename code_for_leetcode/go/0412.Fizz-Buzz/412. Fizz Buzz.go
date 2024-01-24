package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	// 创建一个长度为 n 的字符串切片，用于存储结果
	solution := make([]string, n)

	// 遍历从 1 到 n 的每个数字
	for i := 1; i <= n; i++ {
		// 初始化当前数字的字符串为空
		solution[i-1] = ""

		// 如果当前数字是3的倍数，将字符串追加 "Fizz"
		if i%3 == 0 {
			solution[i-1] += "Fizz"
		}

		// 如果当前数字是5的倍数，将字符串追加 "Buzz"
		if i%5 == 0 {
			solution[i-1] += "Buzz"
		}

		// 如果字符串为空，说明当前数字既不是3的倍数也不是5的倍数，将数字转为字符串存入结果
		if solution[i-1] == "" {
			solution[i-1] = strconv.Itoa(i)
		}
	}

	// 返回最终的结果字符串切片
	return solution
}
