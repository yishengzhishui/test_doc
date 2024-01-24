package leetcode

import (
	"strings"
)

func reverseWords151(s string) string {
	ss := strings.Fields(s)
	reverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse151(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	// 使用 strings.Fields 函数拆分字符串 s，得到一个字符串切片
	//字符串按照空白字符（包括空格、制表符和换行符等）进行分割，
	//返回一个包含所有分割部分的字符串切片。
	//该函数会自动去除字符串开头和结尾的空白字符，并将中间的空白字符作为分隔符，
	//将原始字符串分割成多个子字符串。这在处理单词、句子等需要分割的文本时非常有用。
	words := strings.Fields(s)

	// 使用 for 循环和双指针对字符串切片进行反转
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// 使用 strings.Join 函数将反转后的字符串切片拼接成一个字符串，以空格分隔
	result := strings.Join(words, " ")

	return result
}
