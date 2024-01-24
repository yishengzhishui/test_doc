package leetcode

func validMountainArray(arr []int) bool {
	size := len(arr)
	if size < 3 {
		return false
	}

	left := 0

	// 第一阶段：上升
	for left < size-1 && arr[left] < arr[left+1] {
		left++
	}

	// 边界检查：如果山脉只有上坡或者只有下坡，返回 false
	if left == 0 || left == size-1 {
		return false
	}

	right := left

	// 第二阶段：下降
	for right < size-1 && arr[right] > arr[right+1] {
		right++
	}

	// 边界检查：如果山脉到达最右端，返回 true；否则返回 false
	return right == size-1
}
