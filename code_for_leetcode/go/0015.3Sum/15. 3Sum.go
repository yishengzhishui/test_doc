package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// 双指针法
	// 从小到大排序
	// 两个元素的和是另一个元素的负数
	// 注意要去除重复的数字

	if len(nums) < 3 {
		return nil // 如果数组长度小于3，无法形成三元组，直接返回nil
	}
	sort.Ints(nums) // 对数组从小到大进行排序

	if nums[0] > 0 {
		return nil // 如果数组中的最小元素都大于0，那么无法形成和为0的三元组，直接返回nil
	}

	ans := make([][]int, 0) // 用于存储结果的切片，初始化为空

	for first := 0; first < len(nums)-2; first++ {
		if nums[first] > 0 {
			return ans // 如果第一个元素大于0，后续元素也必然大于0，无法形成和为0的三元组，直接返回结果
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue // 避免重复，如果当前元素与前一个元素相同，跳过当前循环
		}

		left := first + 1      // 左指针，指向第一个元素的下一个元素
		right := len(nums) - 1 // 右指针，指向数组末尾

		for left < right {
			result := nums[first] + nums[left] + nums[right] // 计算当前三元组的和

			if result > 0 {
				right-- // 如果和大于0，右指针向左移动，减小和
			} else if result < 0 {
				left++ // 如果和小于0，左指针向右移动，增大和
			} else {
				// 和为0，将三元组添加到结果中
				ans = append(ans, []int{nums[first], nums[left], nums[right]})

				// 处理重复元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return ans // 返回结果
}
