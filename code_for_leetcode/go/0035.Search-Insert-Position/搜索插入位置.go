package leetcode

func searchInsertV1(nums []int, target int) int {
	// 初始化两个指针 left 和 right 分别指向数组的起始和末尾
	left, right := 0, len(nums)-1

	// 使用二分查找法
	for left <= right {
		// 计算中间位置 mid，避免溢出使用 left + (right-left)>>1
		mid := left + (right-left)>>1

		// 如果中间位置的值大于或等于目标值 target
		if nums[mid] >= target {
			// 将 right 指针移动到 mid 的左边
			right = mid - 1
		} else {
			// 如果 mid 已经是数组的最后一个元素，或者 mid 后一个元素大于等于 target
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				// 返回 mid+1，即为插入位置
				return mid + 1
			}
			// 将 left 指针移动到 mid 的右边
			left = mid + 1
		}
	}

	// 如果循环结束仍未找到插入位置，则返回 0，表示应插入到数组的起始位置
	return 0
}

func searchInsert(nums []int, target int) int {
	// 初始化两个指针 left 和 right 分别指向数组的起始和末尾
	left, right := 0, len(nums)-1

	// 使用二分查找法
	for left <= right {
		// 计算中间位置 mid
		mid := (left + right) / 2

		// 如果中间位置的值大于目标值 target
		if nums[mid] > target {
			// 将 right 指针移动到 mid 的左边
			right = mid - 1
		} else if nums[mid] < target {
			// 如果中间位置的值小于目标值 target
			// 将 left 指针移动到 mid 的右边
			left = mid + 1
		} else {
			// 如果中间位置的值等于目标值 target
			// 返回 mid，表示找到了目标值的位置
			return mid
		}
	}

	// 如果循环结束仍未找到插入位置，则返回 left，表示应插入到数组的起始位置

	return left

	//考虑两种情况：
	//1. 如果 `left > right` 是由于 `left = right + 1` 导致的，
	//说明原来的 `right` 位置大于目标值 `target`。因此，返回原来的 `right` 位置即为插入位置。

	//2. 如果 `left > right` 是由于 `right = left - 1` 导致的，
	//说明原来的 `left` 位置小于目标值 `target`。由于 `right` 能够移动到 `left` 位置，
	//说明此位置右侧的元素都大于目标值 `target`。因此，返回 `left` 位置即为插入位置。
	//
	//因此，直接返回 `left` 是因为 `left` 指针最终指向的位置就是应该插入目标值的位置。这样的返回逻辑可以覆盖两种情况，简化了代码，并有效找到插入位置。
}
