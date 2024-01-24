package leetcode

func inventoryManagement(stock []int, cnt int) []int {
	// 判断 stock 是否为空，如果为空则返回空切片
	if len(stock) == 0 {
		return []int{}
	}

	// 使用 Go 的排序函数对 stock 进行升序排序
	quickSortV2(0, len(stock)-1, stock)

	// 返回 stock 中前 cnt 个元素，即最小的 cnt 个元素
	return stock[:cnt]
}

func quickSortV2(begin, end int, nums []int) {
	// 如果起始索引大于等于结束索引，说明已经完成排序，直接返回
	if begin >= end {
		return
	}
	// 进行分区操作，得到基准元素的索引
	pivotIndex := partitionV2(begin, end, nums)
	// 递归对基准元素左右两侧的子数组进行排序
	quickSortV2(begin, pivotIndex-1, nums)
	quickSortV2(pivotIndex+1, end, nums)
}

// 分区函数，用于将数组的一部分进行分区操作
// partitionV2 使用 mark 记录小于基准的元素的最后一个索引，
//每次小于基准的元素被找到时，mark 递增。
//最后，将基准元素放到 mark 的位置。
func partitionV2(begin, end int, nums []int) int {
	// pivot为基准元素
	pivot := nums[begin]
	// mark为小于pivot的元素的最后一个索引
	mark := begin
	// 遍历数组，将小于pivot的元素与mark索引位置的元素交换，并递增mark
	for i := begin + 1; i <= end; i++ {
		if nums[i] < pivot {
			mark++
			nums[mark], nums[i] = nums[i], nums[mark]
		}
	}
	// 将pivot元素放到最终位置，即mark的位置
	nums[begin], nums[mark] = nums[mark], nums[begin]
	// 返回pivot的最终位置
	return mark
}
