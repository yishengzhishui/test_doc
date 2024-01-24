package main

import "fmt"
//快排：先调配出左右子数组，然后对左右子数组进行排序
//数组取标杆pivot，将小元素放到pivot左边，大元素放到右侧，
//然后一次对右边和左边的子数组继续快排；最后达到整体有序



// 快速排序的主要函数
func quickSort(begin, end int, arr []int) {
	// 如果起始索引大于等于结束索引，说明已经完成排序，直接返回
	if begin >= end {
		return
	}
	// 进行分区操作，得到基准元素的索引
	pivot := partition(begin, end, arr)
	// 递归对基准元素左右两侧的子数组进行排序
	quickSort(begin, pivot-1, arr)
	quickSort(pivot+1, end, arr)
}

// 分区函数，用于将数组的一部分进行分区操作
// partition 使用 counter 记录小于基准的元素个数，
//每次小于基准的元素被找到时，counter 递增。
//最后，将基准元素放到 counter 的位置。
func partition(begin, end int, arr []int) int {
	// pivot为基准元素的索引，counter为小于pivot的元素的计数器
	pivot, counter := end, begin
	// 遍历数组，将小于pivot的元素与counter索引位置的元素交换，并递增counter
	for i := begin; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	// 将pivot元素放到最终位置，即counter的位置
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	// 返回pivot的最终位置
	return counter
}

// 快速排序的主要函数
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

// 主函数，入口点
func main() {
	// 待排序的数组
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("原始数组:", arr)

	// 调用快速排序函数对数组进行排序
	quickSort(0, len(arr)-1, arr)

	// 打印排序后的数组
	fmt.Println("排序后数组:", arr)

	arr = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("原始数组:", arr)

	// 调用快速排序函数对数组进行排序
	quickSortV2(0, len(arr)-1, arr)

	// 打印排序后的数组
	fmt.Println("排序后数组:", arr)
}
