package leetcode

// https://leetcode.cn/problems/number-of-burgers-with-no-waste-of-ingredients/solutions/101702/bu-lang-fei-yuan-liao-de-yi-bao-zhi-zuo-fang-an-2/

// 4x+2y=tomatoSlices
// x+y=cheeseSlices

// 解方程：
// tomatoSlices/2 - cheeseSlices=x
// cheeseSlices*2 - tomatoSlices/2=y

// 边界条件
// tomatoSlices=2k
// tomatoSlices≥2×cheeseSlices
// 4×cheeseSlices≥tomatoSlices
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	// 判断是否存在可能的解，若不存在则返回空切片
	if tomatoSlices%2 != 0 || tomatoSlices < 2*cheeseSlices || cheeseSlices*4 < tomatoSlices {
		return []int{}
	}

	// 计算两种汉堡的数量并返回
	return []int{tomatoSlices/2 - cheeseSlices, cheeseSlices*2 - tomatoSlices/2}
}
