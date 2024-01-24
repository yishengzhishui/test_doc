package leetcode

import (
	"sort"
)

// 解法一 双指针
func fourSumV1(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

// 先排序
// 递归方法+双指针 => N数之和
func fourSum(nums []int, target int) [][]int {
	// 定义结果集
	var result [][]int

	// 排序数组
	sort.Ints(nums)

	// 递归辅助函数，用于寻找 N 数之和
	var helper func(N, target, start, end int, cur []int)
	helper = func(N, target, start, end int, cur []int) {
		// 判断是否符合递归的条件
		if 2 <= N && N <= end-start+1 && nums[start]*N <= target && target <= nums[end]*N {
			// 当 N 为 2 时，使用双指针方法找到和为 target 的两个数
			if N == 2 {
				for start < end {
					twoSum := nums[start] + nums[end]
					if twoSum < target {
						start++
					} else if twoSum > target {
						end--
					} else {
						// 找到符合条件的两个数，加入结果集
						result = append(result, append(cur, nums[start], nums[end]))

						// 跳过相同的元素
						for start < end && nums[start] == nums[start+1] {
							start++
						}
						for start < end && nums[end] == nums[end-1] {
							end--
						}

						// 移动指针
						start++
						end--
					}
				}
			} else {
				// 当 N 不为 2 时，递归调用自身，寻找 N-1 数之和
				for i := start; i <= end; i++ {
					// 跳过相同的元素
					if i > start && nums[i] == nums[i-1] {
						continue
					}
					helper(N-1, target-nums[i], i+1, end, append(cur, nums[i]))
				}
			}
		}
	}

	// 调用递归辅助函数
	helper(4, target, 0, len(nums)-1, []int{})

	// 返回结果集
	return result
}
