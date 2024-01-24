package leetcode

import "fmt"

// 解法一

func subsets(nums []int) [][]int {
	// 初始化结果切片
	var result [][]int

	// 定义递归辅助函数
	var helper func(int, []int)
	helper = func(i int, tmp []int) {
		// 将当前子集添加到结果中
		result = append(result, append([]int{}, tmp...))
		fmt.Println(result)
		// 从当前位置 i 开始，递归生成子集
		for j := i; j < len(nums); j++ {
			helper(j+1, append(tmp, nums[j]))
		}
	}

	// 调用递归辅助函数
	helper(0, []int{})

	return result
}

//使用了迭代的方式生成子集，
//每次迭代都在之前的基础上添加新的数字。最终，result 中包含了所有可能的子集，包括空子集。
func subsetsV1(nums []int) [][]int {
	// 初始化结果切片，包含一个空子集
	result := [][]int{{}}

	// 遍历输入数组
	for _, num := range nums {
		// 遍历当前结果中的每个子集
		for _, subset := range result {
			// 将当前数字添加到子集中，并将新子集添加到结果中
			result = append(result, append([]int{num}, subset...))
		}
	}

	return result
}
