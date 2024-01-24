package leetcode

// 找到最小值，mid与right比较，
// 如果mid大，则说明mid到right之间一定发生了旋转，那么最小值在mid+1到right之间
// 如果right大，说明mid到right之间是单调递增的，最小值会在 left到mid之间
func findMin(nums []int) int {
	// 初始化左右指针
	left, right := 0, len(nums)-1

	// 进行二分查找
	for left < right {
		// 计算中间位置
		mid := (left + right) / 2

		// 如果中间值大于右边界值，说明最小值在右半部分
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// 否则，最小值在左半部分或就是中间值
			right = mid
		}
	}

	// 最终 left 指向最小值
	return nums[left]
}

// 找到最大值，往后退一格就是最小值，mid与left比较，
// 如果left大，说明mid和left之间发生旋转，那么最大值就在left到mid-1之间
// 如果mid大，说明left和mid之间是单调递增的，最大mid到right之间

// 与left比较要注意 mid要靠近right一点，否则可能会进入死循环
// 因为下方left=mid 后面计算的时候，left和right相邻的话，可能导致死循环，left一直等于mid
func findMinV2(nums []int) int {
	// 初始化左右指针
	left, right := 0, len(nums)-1

	// 使用二分查找找到旋转点
	for left < right {
		// 计算当前中间位置
		mid := (left + right + 1) / 2

		// 如果中间值小于左边界值，说明旋转点在左半部分
		if nums[mid] < nums[left] {
			right = mid - 1
		} else {
			// 否则，旋转点在右半部分
			left = mid
		}
	}

	// 返回最小值，考虑循环数组的情况
	return nums[(left+1)%len(nums)]
}
