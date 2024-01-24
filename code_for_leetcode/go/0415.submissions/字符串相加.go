package leetcode

import "strconv"

func addStrings(num1 string, num2 string) string {
	// 初始化变量
	carry := 0
	res := ""

	// 初始化索引
	i, j := len(num1)-1, len(num2)-1

	// 循环直到两个数字的所有位都处理完毕
	for i >= 0 || j >= 0 || carry > 0 {
		// 获取当前位上的数字，如果已经没有数字则默认为0
		n1 := 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}

		n2 := 0
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}

		// 计算当前位上的和以及进位
		tmp := n1 + n2 + carry
		carry = tmp / 10

		// 在结果字符串前面插入当前位的数字
		res = strconv.Itoa(tmp%10) + res
	}

	return res
}
