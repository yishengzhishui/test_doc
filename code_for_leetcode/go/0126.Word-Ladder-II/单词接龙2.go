package leetcode

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordListSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordListSet[word] = true
	}

	// 如果目标单词不在单词列表中，直接返回 0
	if !wordListSet[endWord] {
		return [][]string{}
	}

	// 初始化层级映射，存储每个单词的路径
	layer := map[string][][]string{beginWord: {{beginWord}}}

	// 循环直到层级映射为空
	for len(layer) > 0 {
		// 初始化新的层级映射
		newLayer := make(map[string][][]string)

		// 遍历当前层级映射中的每个单词
		for word := range layer {
			// 如果当前单词等于目标单词，直接返回该单词的路径
			if word == endWord {
				return layer[word]
			}

			// 遍历当前单词的每个位置
			for i := range word {
				// 尝试变换当前单词的每个位置的字母
				for s := 'a'; s <= 'z'; s++ {
					newWord := word[:i] + string(s) + word[i+1:]

					// 如果新单词在单词列表中
					if wordListSet[newWord] {
						// 将新单词的路径添加到新的层级映射中
						for _, path := range layer[word] {
							curPath := append([]string{}, append(path, newWord)...)
							newLayer[newWord] = append(newLayer[newWord], curPath)

						}
					}
				}
			}
		}

		// 更新层级映射
		layer = newLayer

		// 从单词列表中移除已经访问过的单词
		for word := range layer {
			delete(wordListSet, word)
		}

	}

	// 如果循环结束仍未找到解，返回空切片
	return [][]string{}
}
func findLaddersV1(beginWord string, endWord string, wordList []string) [][]string {
	wordListSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordListSet[word] = true
	}

	// 如果目标单词不在单词列表中，直接返回 0
	if !wordListSet[endWord] {
		return [][]string{}
	}
	// 定义字母表
	ls := "abcdefghijklmnopqrstuvwxyz"

	// 初始化层级映射，存储每个单词的路径
	layer := map[string][][]string{beginWord: {{beginWord}}}

	// 循环直到层级映射为空
	for len(layer) > 0 {
		// 初始化新的层级映射
		newLayer := make(map[string][][]string)

		// 遍历当前层级映射中的每个单词
		fmt.Println("开始")
		fmt.Println(layer)
		for word := range layer {
			// 如果当前单词等于目标单词，直接返回该单词的路径
			if word == endWord {
				return layer[word]
			}

			// 遍历当前单词的每个位置
			for i := range word {
				// 尝试变换当前单词的每个位置的字母
				for _, s := range ls {
					newWord := word[:i] + string(s) + word[i+1:]

					// 如果新单词在单词列表中
					if wordListSet[newWord] {
						// 将新单词的路径添加到新的层级映射中
						for _, path := range layer[word] {
							curPath := append([]string{}, append(path, newWord)...)
							newLayer[newWord] = append(newLayer[newWord], curPath)

						}
					}
				}
			}
		}

		// 更新层级映射
		layer = newLayer
		fmt.Println("更新层级，放入新的word(layer = newLayer)")
		fmt.Println(layer)

		// 从单词列表中移除已经访问过的单词
		for word := range layer {
			delete(wordListSet, word)
		}

		fmt.Println("在wordlist中删除已经访问过的word，开始新的遍历")
		fmt.Println("next-----")

	}

	// 如果循环结束仍未找到解，返回空切片
	return [][]string{}
}
