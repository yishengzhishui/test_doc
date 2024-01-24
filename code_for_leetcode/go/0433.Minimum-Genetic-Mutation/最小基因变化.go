package leetcode

//BFS
func minMutation(start string, end string, bank []string) int {

	// 将基因库转为集合，以提高检索效率
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}
	// 如果目标基因不在基因库中，直接返回 -1
	if !bankSet[end] {
		return -1
	}

	// 定义基因库中的碱基
	ls := "ACGT"

	// 初始化队列，用于广度优先搜索
	type GeneStep struct {
		Gene string
		Step int
	}

	queue := []GeneStep{{start, 0}}

	// 开始广度优先搜索
	for len(queue) > 0 {
		// 出队列
		current := queue[0]
		queue = queue[1:]

		// 如果当前基因等于目标基因，返回步数
		if current.Gene == end {
			return current.Step
		}

		// 遍历当前基因的每个碱基，尝试变异
		for i := 0; i < len(current.Gene); i++ {
			for _, s := range ls {
				newGene := current.Gene[:i] + string(s) + current.Gene[i+1:]

				// 如果新基因在基因库中，加入队列，并从基因库中移除(避免重复)
				if bankSet[newGene] {
					queue = append(queue, GeneStep{newGene, current.Step + 1})
					delete(bankSet, newGene)
				}
			}
		}
	}

	// 如果循环结束仍未找到解，表示无法从起始基因变异到目标基因，返回 -1
	return -1
}

// 双向BFS
// 定义函数 minMutation，计算从起始基因到目标基因的最小变异步数
func minMutationV1(start string, end string, bank []string) int {
	// 将基因库转为集合，以提高检索效率
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}
	// 如果目标基因不在基因库中，直接返回 -1
	if !bankSet[end] {
		return -1
	}

	// 定义基因库中的碱基
	ls := []string{"A", "C", "G", "T"}

	// 初始化前向和后向集合，以及变异步数
	front, back := map[string]bool{start: true}, map[string]bool{end: true}
	step := 0

	// 开始广度优先搜索
	for len(front) > 0 {
		newFront := make(map[string]bool)
		step++

		// 遍历当前前向集合中的基因
		for gene := range front {
			// 遍历基因的每个碱基
			for i := 0; i < len(gene); i++ {
				// 尝试变异每个碱基为 A、C、G、T
				for _, s := range ls {
					newGene := gene[:i] + s + gene[i+1:]

					// 如果新基因在后向集合中，表示找到了最小变异步数，直接返回步数
					// 因为无论是从前向后还是从后向前， step是同一个，直接返回就行
					if back[newGene] {
						return step
					}

					// 如果新基因在基因库中，加入新的前向集合，并从基因库中移除，避免重复遍历
					if bankSet[newGene] {
						newFront[newGene] = true
						delete(bankSet, newGene)
					}
				}
			}
		}

		// 更新前向集合
		front = newFront

		// 如果前向集合的大小超过后向集合，交换两者，以保持前向集合较小
		// 以较小的进行遍历
		if len(front) > len(back) {
			front, back = back, front
		}
	}

	// 如果循环结束仍未找到解，表示无法从起始基因变异到目标基因，返回 -1
	return -1
}
