package leetcode

// 完成所有任务的最短时间取决于出现次数最多的任务数量。
// A -> B -> (单位时间) -> A -> B -> (单位时间) -> A -> B
// 前面两个 A 任务一定会固定跟着 2 个单位时间的间隔；
// 最后一个 A 之后是否还有任务跟随取决于是否存在与任务 A 出现次数相同的任务。
// (任务 A 出现的次数 - 1) * (n + 1) + (出现次数为 3 的任务个数)
// 但是存在特殊的情况
// 在将所有的任务安排完成后，如果仍然有剩余的空闲时间，那么答案即为（任务的总数 + 剩余的空闲时间）；
// 如果在安排某一个任务时，遇到了剩余的空闲时间不够的情况，那么答案一定就等于任务的总数。

// leastInterval 函数用于计算任务调度的最小时间
func leastInterval(tasks []byte, n int) int {
	// 统计每个任务的出现次数
	taskFreq := make(map[byte]int)
	for _, task := range tasks {
		taskFreq[task]++
	}

	maxFreq := 0      // 记录出现次数最多的任务的次数
	countMaxFreq := 0 // 记录与最大次数相同的任务个数

	// 找到出现次数最多的任务以及相同最大次数的任务个数
	for _, freq := range taskFreq {
		if freq == maxFreq {
			countMaxFreq++
		} else if freq > maxFreq {
			maxFreq = freq
			countMaxFreq = 1
		}
	}

	// 计算最终结果，使用公式 (maxFreq-1)*(n+1)+countMaxFreq
	result := getMax(len(tasks), (maxFreq-1)*(n+1)+countMaxFreq)
	return result
}

// getMax 函数用于返回两个整数中的较大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
