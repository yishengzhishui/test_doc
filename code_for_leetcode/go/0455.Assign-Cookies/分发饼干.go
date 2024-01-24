package leetcode

import "sort"

// g,s都先要从小到大排序
// 排序后，依次遍历，如果s[j]>=g[i]则 i+1，否则i不动
// 贪心算法：尽量用较小的饼干去满足较小胃口的孩子，以最大化满足孩子的数量。
func findContentChildren(g []int, s []int) int {
	// 对孩子的胃口数组和饼干的大小数组进行升序排序
	sort.Ints(g)
	sort.Ints(s)

	// gi 表示孩子数组的当前位置，sj 表示饼干数组的当前位置
	gi, sj := 0, 0

	// 循环条件：孩子数组和饼干数组均未遍历完
	for gi < len(g) && sj < len(s) {
		// 判断当前饼干是否能够满足当前孩子的胃口
		if s[sj] >= g[gi] {
			// 如果能够满足，将孩子指针 gi 向后移动一位，表示当前孩子已经得到满足
			gi++
		}

		// 无论是否满足当前孩子，饼干指针 sj 都向后移动一位，准备尝试下一块饼干
		sj++
	}

	// 返回最终满足的孩子数量，即 gi 的值
	return gi
}
