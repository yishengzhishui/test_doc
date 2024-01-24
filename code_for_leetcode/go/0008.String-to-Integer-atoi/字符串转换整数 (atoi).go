package leetcode

// myAtoi 将字符串转换为整数。
func myAtoi(str string) int {
	// 去除字符串两端的空格
	str = trimSpace(str)
	if str == "" {
		return 0
	}

	// 判断正负号
	sign := 1
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			sign = -1
		}
		// 去除正负号后的字符串
		str = str[1:]
	}

	result, i := 0, 0
	for i < len(str) && isDigit(str[i]) {
		// 检查溢出
		if result > (1<<31-1)/10 || (result == (1<<31-1)/10 && int(str[i]-'0') > 7) {
			if sign == 1 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}

		// 将字符转换为数字并累加
		result = result*10 + int(str[i]-'0')
		i++
	}

	// 结果乘以符号
	result *= sign

	return result
}

// trimSpace 去除字符串两端的空格
func trimSpace(str string) string {
	i, j := 0, len(str)-1

	// 向右找到第一个非空格字符
	for i <= j && str[i] == ' ' {
		i++
	}

	// 向左找到第一个非空格字符
	for i <= j && str[j] == ' ' {
		j--
	}

	// 返回去除空格后的子串
	return str[i : j+1]
}

// isDigit 判断字符是否是数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
