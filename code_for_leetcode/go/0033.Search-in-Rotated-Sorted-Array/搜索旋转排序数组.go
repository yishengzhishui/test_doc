package leetcode

// 二分查找
// 0~mid不包含旋转，且target在此范围内：nums[0] <= target <= nums[mid]
// 0~mid包含旋转，分两两种情况：
// 1，target <= nums[mid] < nums[0] （target 在旋转位置到 mid 之间）；
// 2，nums[mid] < nums[0] <= target（target 在 0 到旋转位置之间）；

// 上述三种情况可以总结如下：
// nums[0] <= target <= nums[mid]
// target <= nums[mid] < nums[0]
// nums[mid] < nums[0] <= target

// (nums[0] <= target)， (target <= nums[mid]) ，(nums[mid] < nums[0])，
// 现在我们想知道这三项中有哪两项为真（明显这三项不可能均为真（因为这三项可能已经包含了所有情况））
// 所以我们现在只需要区别出这三项中有两项为真还是只有一项为真。
// 使用 “异或” 操作可以轻松的得到上述结果（两项为真时异或结果为假，一项为真时异或结果为真

// 换句话说就是： nums[mid] < nums[0]，nums[0] > target，target > nums[mid] 三项均为真或者只有一项为真时向后规约。

func search(nums []int, target int) int {
	// 如果切片为空，返回 -1
	if len(nums) == 0 {
		return -1
	}

	// 获取切片长度
	size := len(nums)

	// 初始化左右指针
	left, right := 0, size-1

	// 开始二分查找
	for left <= right {
		mid := (left + right) / 2

		// 如果找到目标值，返回索引
		if nums[mid] == target {
			return mid
		}

		// 判断左半部分是否有序
		isLeftSorted := nums[left] <= nums[mid]

		// 判断目标值是否在有序的左半部分
		isTargetInLeft := target >= nums[left] && target <= nums[mid]

		// 判断目标值是否在有序的右半部分
		isTargetInRight := target >= nums[mid] && target <= nums[right]

		// 根据情况更新左右指针
		// 如果左半部分有序且目标值不在左半部分，
		//或者左半部分无序且目标值在右半部分，将 left 更新为 mid + 1；

		if (isLeftSorted && !isTargetInLeft ) ||
			(!isLeftSorted && isTargetInRight) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 如果未找到目标值，返回 -1
	return -1
}

func searchSimple(nums []int, target int) int {
	// 如果切片为空，返回 -1
	if len(nums) == 0 {
		return -1
	}
	// 初始化左右指针
	left, right := 0, len(nums)-1

	// 开始二分查找
	for left <= right {
		// 计算当前中间位置
		mid := (left + right) / 2

		// 如果目标值等于中间值，直接返回索引
		if target == nums[mid] {
			return mid
		}

		// 判断左半部分是否有序
		if nums[left] <= nums[mid] {
			// 如果目标值在左半部分有序范围内，更新右指针
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				// 否则，更新左指针
				left = mid + 1
			}
		} else {
			// 如果目标值在右半部分有序范围内，更新左指针
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				// 否则，更新右指针
				right = mid - 1
			}
		}
	}

	// 如果未找到目标值，返回 -1
	return -1
}
