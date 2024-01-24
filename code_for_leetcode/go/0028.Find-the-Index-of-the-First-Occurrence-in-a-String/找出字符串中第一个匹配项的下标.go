package leetcode

import "strings"

// 解法一
func strStr(haystack string, needle string) int {
	// 外层循环，遍历 haystack 中的每一个可能的起始位置
	for i := 0; ; i++ {
		// 内层循环，遍历 needle 中的每一个字符
		for j := 0; ; j++ {
			// 如果 needle 中的所有字符都匹配，返回当前的起始位置 i
			if j == len(needle) {
				return i
			}
			// 到最后的了，但是已经不够继续的数量，就直接返回（也算是剪枝）
			// 如果 haystack 已经遍历完，但 needle 还没有匹配完全，返回 -1
			if i+j == len(haystack) {
				return -1
			}
			// 如果 needle[j] 与 haystack[i+j] 不相等，跳出内层循环
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

// 解法二
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
