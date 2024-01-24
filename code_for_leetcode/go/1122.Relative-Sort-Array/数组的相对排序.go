package leetcode

// 计数排序
// 对arr1中所有的元素进行计数
// 然后遍历arr2，根据这个顺序进行排列并列出重复的元素，最后将剩余的元素放在最后

func relativeSortArrayV1(arr1 []int, arr2 []int) []int {
	// 计数排序
	arrMax := getMax(arr1...)
	result := make([]int, 0, len(arr1))

	// 使用 map 统计 arr1 中每个元素的出现次数
	pre := make(map[int]int)
	for _, num := range arr1 {
		pre[num]++
	}

	// 遍历 arr2，按照 arr2 的顺序进行排列并列出重复的元素
	for _, num := range arr2 {
		result = append(result, repeat(pre[num], num)...)
		pre[num] = 0
	}

	// 将剩余的元素放在最后
	for i := 0; i <= arrMax; i++ {
		result = append(result, repeat(pre[i], i)...)
	}

	return result
}

func getMax(nums ...int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

// 辅助函数，返回重复 count 次的 num 切片
func repeat(count, num int) []int {
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = num
	}
	return result
}

// arr1 和 arr2 最大就1000
func relativeSortArray(arr1 []int, arr2 []int) []int {
	ans := make([]int, len(arr1))
	var n int

	// 计数数组
	count := make([]int, 1001)
	for i := 0; i < len(arr1); i++ {
		count[arr1[i]]++
	}

	// 按照arr2的顺序来取，先放arr2里面的元素
	for i := 0; i < len(arr2); i++ {
		for count[arr2[i]] > 0 {
			count[arr2[i]]--
			ans[n] = arr2[i]
			n++
		}
	}

	// 剩余的没在arr2出现过，按顺序取出
	for i := 0; i <= 1000; i++ {
		for count[i] > 0 {
			count[i]--
			ans[n] = i
			n++
		}
	}

	return ans
}
