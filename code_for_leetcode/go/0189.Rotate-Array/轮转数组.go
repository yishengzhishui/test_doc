package leetcode

// 解法一 时间复杂度 O(n)，空间复杂度 O(1)
func rotate(nums []int, k int) {
	// 对 k 取余，防止 k 大于数组长度时进行多余的旋转
	k %= len(nums)

	// 整体翻转数组
	helper(0, len(nums)-1, nums)

	// 翻转前 k 个元素
	helper(0, k-1, nums)

	// 翻转剩余的元素
	helper(k, len(nums)-1, nums)
}

func helper(start, end int, nums []int) {
	// 使用双指针，将数组 nums 中从 start 到 end 的元素进行反转
	i, j := start, end

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		j--
		i++
	}
}

// 解法二 时间复杂度 O(n)，空间复杂度 O(n)
func rotate1(nums []int, k int) {
	// 创建一个新数组 newNums，用于存储旋转后的结果
	newNums := make([]int, len(nums))

	// 遍历原始数组 nums
	for i, v := range nums {
		// 计算新数组中的索引位置，采用 (i+k)%len(nums) 的方式实现循环移动
		newNums[(i+k)%len(nums)] = v
	}

	// 将新数组的元素复制回原始数组 nums
	copy(nums, newNums)
}
