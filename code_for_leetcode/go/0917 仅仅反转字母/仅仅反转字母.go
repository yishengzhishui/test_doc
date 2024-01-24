package leetcode

import "unicode"

func reverseOnlyLettersV1(S string) string {
	// 将输入的字符串 S 转换为字符切片 SChars，以便能够修改字符串中的字符。
	SChars := []rune(S)
	// 初始化两个指针 i 和 j，分别指向字符串的开头和结尾。
	i, j := 0, len(SChars)-1

	// 使用循环，当 i 小于 j 时执行循环体。
	for i < j {
		// 如果 SChars[i] 不是字母，将 i 指针向右移动。
		if !unicode.IsLetter(SChars[i]) {
			i++
			// 如果 SChars[j] 不是字母，将 j 指针向左移动。
		} else if !unicode.IsLetter(SChars[j]) {
			j--
			// 如果 SChars[i] 和 SChars[j] 都是字母，交换它们的位置，并将 i 向右移动，j 向左移动。
		} else {
			SChars[i], SChars[j] = SChars[j], SChars[i]
			i++
			j--
		}
	}

	// 将最终的字符切片 SChars 转换为字符串并返回。
	return string(SChars)
}

func reverseOnlyLetters(S string) string {
	// 初始化一个切片用于模拟栈
	stack := make([]rune, 0)

	// 初始化一个切片用于存储最终结果
	result := make([]rune, 0)

	// 遍历字符串 S，将字母压入栈中
	for _, s := range S {
		if unicode.IsLetter(s) {
			stack = append(stack, s)
		}
	}

	// 再次遍历字符串 S，如果是字母则从栈中弹出，构建最终结果
	for _, i := range S {
		if unicode.IsLetter(i) {
			result = append(result, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			result = append(result, i)
		}
	}

	// 将结果切片转换为字符串并返回
	return string(result)
}
