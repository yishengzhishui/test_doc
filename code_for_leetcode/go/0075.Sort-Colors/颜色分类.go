package leetcode

//双指针
func sortColorsV1(nums []int) {
	//红色(0)、白色(1)和蓝色(2)
	red, white, blue := 0, 0, len(nums)-1

	//red指向红色区域的下一个位置，white是当前遍历的指针，blue指向蓝色区域的下一个位置。
	for white <= blue {
		switch nums[white] {
		case 0:
			// 如果当前元素是红色(0)，将其交换到红色区域，并将红色和白色指针都向前移动
			nums[red], nums[white] = nums[white], nums[red]
			white++
			red++
		case 1:
			// 如果当前元素是白色(1)，只需将白色指针向前移动
			white++
		case 2:
			// 如果当前元素是蓝色(2)，将其交换到蓝色区域，并将蓝色指针向前移动
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		}
	}
}

// 稍微优化一下
func sortColorsV2(nums []int) {
	if len(nums) == 0 {
		return
	}

	sortedLeft, sortedRight := 0, len(nums)-1

	for i := 0; i <= sortedRight; {
		switch {
		case nums[i] == 0:
			nums[i], nums[sortedLeft] = nums[sortedLeft], nums[i]
			sortedLeft++
			i++
		case nums[i] == 2:
			nums[i], nums[sortedRight] = nums[sortedRight], nums[i]
			sortedRight--
		default:
			i++
		}
	}

	return
}

// 单指针
// sortColors 使用两次颜色交换，先将颜色为0的元素排到数组的最前面，然后将颜色为1的元素排到剩余部分的最前面
func sortColors(nums []int) {
	// 第一次颜色交换，将颜色为0的元素排到数组最前面，并获取0的个数
	count0 := swapColors(nums, 0)
	// 第二次颜色交换，将剩余部分的颜色为1的元素排到数组最前面
	swapColors(nums[count0:], 1)
}

// swapColors 将目标颜色移到切片前面，并返回目标颜色的个数
// colors 和 nums共用的底层数组
func swapColors(colors []int, target int) (countTarget int) {
	// 遍历切片，通过双指针将目标颜色移到前面
	for i, c := range colors {
		if c == target {

			// 交换元素，并将目标颜色个数加一
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}
