package leetcode

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	// zero是0的下标
	// 利用快排的思想
	// 以0与标杆，遍历数组，将不等于0的元素放到0的左侧

	zero := 0 // zero 是 0 的下标，表示当前可以放置非零元素的位置

	for index, value := range nums {
		if value != 0 { // 如果当前元素不是零
			// 交换非零元素和零元素的位置
			//将当前非零元素和零元素位置的元素交换，将非零元素放到零元素的左侧。
			nums[index], nums[zero] = nums[zero], nums[index]
			zero++ // 移动 zero 指针，表示下一个可以放置非零元素的位置
		}
	}
}
