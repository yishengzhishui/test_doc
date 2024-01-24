package leetcode

import (
	"strings"
)

func checkRecordV1(s string) bool {
	record := []rune(s)

	countA := 0
	for _, char := range record {
		if char == 'A' {
			countA++
		}
	}

	countLLL := strings.Count(s, "LLL")

	return countA < 2 && countLLL == 0
}

func checkRecord(s string) bool {
	// 初始化缺席次数和连续迟到次数
	absentCount, lateCount := 0, 0

	// 遍历字符串中的每个字符
	for _, ch := range s {
		// 根据字符类型进行不同的处理
		switch ch {
		// 如果是 'A'，表示学生缺席
		case 'A':
			// 缺席次数增加
			absentCount++
			// 如果缺席次数超过了规定的次数（2次），直接返回 false
			if absentCount >= 2 {
				return false
			}
			// 缺席时重置连续迟到次数为 0
			lateCount = 0
		// 如果是 'L'，表示学生迟到
		case 'L':
			// 迟到次数增加
			lateCount++
			// 如果连续迟到次数超过了规定的次数（3次），直接返回 false
			if lateCount >= 3 {
				return false
			}
		// 如果是其他字符，重置连续迟到次数为 0
		default:
			lateCount = 0
		}
	}

	// 如果遍历完整个字符串都没有违反规定，返回 true，表示学生的考勤是合规的
	return true
}
