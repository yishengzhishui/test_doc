package leetcode

//顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。
//如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。
func twoSum(nums []int, target int) []int {
	// 创建一个哈希表，用于存储数组元素和它们的索引
	hashTable := map[int]int{}

	// 遍历数组
	for index, value := range nums {
		// 计算目标值与当前值的差值
		complement := target - value

		// 在哈希表中查找差值是否已经在之前的元素中出现过
		if a, ok := hashTable[complement]; ok {
			// 如果找到了，返回当前元素的索引和之前出现的元素的索引
			return []int{index, a}
		}

		// 如果没有找到差值，将当前元素和索引存入哈希表中
		hashTable[value] = index
	}

	// 如果遍历完整个数组都没有找到满足条件的两个数，返回nil
	return nil
}
