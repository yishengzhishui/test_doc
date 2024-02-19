package interview

import (
	"fmt"
	"testing"
)

// 动态生成所有可能的子集，并记录每个子集的和，将和相等的子集存储在 sumMap 中。
//检查 sumMap 中和相等的子集是否有交集，如果有，则排除这些子集，取剩余子集的和的最大值。
func subsets(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	result := [][]int{{}}
	sumMap := make(map[int][][]int)

	for _, num := range nums {
		for _, subset := range result {
			newSubset := append([]int{num}, subset...)
			result = append(result, newSubset)
			sum := getSum(newSubset)
			sumMap[sum] = append(sumMap[sum], newSubset)
		}
	}
	fmt.Println(sumMap)

	var resultNum []int

	for key, value := range sumMap {
		if len(value) > 1 {
			for i := 0; i < len(value)-1; i++ {
				for j := i + 1; j < len(value); j++ {
					if !intersection(value[i], value[j]) {
						resultNum = append(resultNum, key)
					}
				}
			}
		}
	}
	if len(resultNum) < 1 {
		return 0
	}
	return getMax(resultNum)
}

// 当前切片最大值
func getMax(nums []int) int {
	res := nums[0]
	for _, value := range nums {
		if value > res {
			res = value
		}
	}
	return res
}

// 切片和
func getSum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 两个切片是否存在交集
func intersection(nums1 []int, nums2 []int) bool {
	m := map[int]bool{}
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			return true
		}
	}
	return false
}

func Test_Problem(t *testing.T) {
	// 示例1
	//nums1 := []int{2, 3, 4, 5, 9}
	//fmt.Println(subsets(nums1)) // 输出：9
	////// 示例2
	//nums2 := []int{2, 3}
	//fmt.Println(subsets(nums2)) // 输出：0

	//// 示例2
	nums3 := []int{1, 1}
	fmt.Println(subsets(nums3)) // 输出：0
}

//题目详情：
//设计一个函数，输入一个由正整数组成的数组 nums，该数组中可能存在不相交并且和相等的子集，如果存在，将其中最大的和返回，如果不存在，请返回0
//
//提示：
//1. 0 <= nums.length <= 20
//2. 1 <= nums[i] <= 1000
//3. sum(nums[i]) <= 5000
//
//示例1:
//输入：[2,3,4,5,9]
//输出：9
//解释: 有不相交的子集 {2,3,4} 和 {9}，他们的和相同，为9
//
//示例2:
//输入：[2,3]
//输出：0
//解释: 没有不相交并且和相等的子集
//
//1、请于 两天 内提交相关文件
//2、语言无限制
//3、请在1小时内完成该测试题
//4、代码源文件请另外贴在邮件正文中
