package leetcode

//和 433 最小基因的解法基本一致
// BFS 26个字母挨个试
// 时间较长，但是思路简单
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordListSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordListSet[word] = true
	}

	// 如果目标单词不在单词列表中，直接返回 0
	if !wordListSet[endWord] {
		return 0
	}
	// 使用队列进行广度优先搜索，初始状态包含起始单词和步数
	type wordStep struct {
		word string
		step int
	}
	queue := []wordStep{{beginWord, 1}}

	// 开始广度优先搜索
	for len(queue) > 0 {
		cur, step := queue[0].word, queue[0].step
		queue = queue[1:]

		// 如果当前单词等于目标单词，找到了最短变换路径，返回步数
		if cur == endWord {
			return step
		}

		// 遍历当前单词的每个位置，尝试变换每个字母
		for i := 0; i < len(cur); i++ {
			for s := 'a'; s <= 'z'; s++ {
				newWord := cur[:i] + string(s) + cur[i+1:]

				// 如果新单词在单词列表中，加入队列，并从列表中移除，避免重复遍历
				if wordListSet[newWord] {
					queue = append(queue, wordStep{newWord, step + 1})
					delete(wordListSet, newWord)
				}
			}
		}
	}

	// 如果循环结束仍未找到解，表示无法从起始单词变换到目标单词，返回 0
	return 0
}

// 双向BFS
func ladderLengthV1(beginWord string, endWord string, wordList []string) int {

	// 初始化步数和单词列表
	step, wordListSet := 1, make(map[string]bool)
	for _, word := range wordList {
		wordListSet[word] = true
	}

	// 如果目标单词不在单词列表中，直接返回 0
	if !wordListSet[endWord] {
		return 0
	}
	// 定义字母表
	ls := "abcdefghijklmnopqrstuvwxyz"

	// 初始化前向和后向集合，将起始单词加入前向集合
	front, back := map[string]bool{beginWord: true}, map[string]bool{endWord: true}

	// 开始双向广度优先搜索
	// 只写front，因为后面front与back会互换，将较小的那个赋为front
	for len(front) > 0 {
		// 新的前向集合
		newFront := make(map[string]bool)
		step++

		// 遍历当前前向集合中的每个单词
		for word := range front {
			// 尝试每个位置的每个字母变换
			for i := 0; i < len(word); i++ {
				for _, s := range ls {
					newWord := word[:i] + string(s) + word[i+1:]

					// 前后BFS相交了
					//如果新单词在后向集合中，找到了最短变换路径，返回步数
					if back[newWord] {
						return step
					}

					// 如果新单词在单词列表中，加入新的前向集合，并从列表中移除，避免重复遍历
					if wordListSet[newWord] {
						newFront[newWord] = true
						delete(wordListSet, newWord)
					}
				}
			}
		}

		// 更新前向集合
		front = newFront

		// 如果前向集合的大小超过后向集合，交换两者，以保持前向集合较小
		if len(front) > len(back) {
			front, back = back, front
		}
	}

	// 如果循环结束仍未找到解，表示无法从起始单词变换到目标单词，返回 0
	return 0
}

// BFS+ 预处理
func ladderLengthV2(beginWord string, endWord string, wordList []string) int {
	// 构建预处理的映射关系
	dic := preConstruct(wordList)
	// 调用 BFS 函数进行搜索并返回最短转换序列长度
	return BFS(beginWord, endWord, dic)
}

// 预处理函数，构建哈希表，键为中间状态的模式，值为拥有相同模式的单词列表
// 例子： begin:dug, end: dog, wordList:[dog, dig]
// 预处理 ：hash[d*g: [dog,dig], *og: [dog], do*:[dog], *ig:[dig]......]
func preConstruct(words []string) map[string][]string {
	dicWord := make(map[string][]string)

	for _, word := range words {
		for i := range word {
			newWord := word[:i] + "#" + word[i+1:]
			dicWord[newWord] = append(dicWord[newWord], word)
		}
	}

	return dicWord
}

// BFS 函数，搜索最短转换序列
func BFS(begin, end string, hash map[string][]string) int {
	// 初始化队列，存储当前单词及其转换步数

	type wordStep struct {
		word string
		step int
	}
	queue := []wordStep{{begin, 1}}
	// 初始化已访问的单词集合
	visited := make(map[string]bool)

	// 循环直到队列为空
	for len(queue) > 0 {
		// 出队列，获取当前单词及其转换步数
		cur, step := queue[0].word, queue[0].step
		queue = queue[1:]

		// 如果当前单词等于目标单词，返回转换步数
		if cur == end {
			return step
		}

		// 遍历当前单词的每个位置
		for i := range cur {
			// 构建新的中间状态模式
			newWord := cur[:i] + "#" + cur[i+1:]
			// 获取具有相同中间状态的单词列表
			referWord := hash[newWord]

			// 遍历每个具有相同中间状态的单词
			for _, word := range referWord {
				// 如果单词未被访问过
				if !visited[word] {
					// 将单词及其步数添加到队列
					queue = append(queue, wordStep{word, step + 1})
					// 标记单词为已访问
					visited[word] = true
				}
			}
		}
	}

	// 如果队列为空且未找到解，返回 0
	return 0
}
