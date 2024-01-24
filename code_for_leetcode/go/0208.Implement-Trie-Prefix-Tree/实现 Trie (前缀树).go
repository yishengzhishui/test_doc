package leetcode

type Trie struct {
	root map[rune]*Trie
	end  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: make(map[rune]*Trie), end: false}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, char := range word {
		if _, ok := node.root[char]; !ok {
			node.root[char] = &Trie{root: make(map[rune]*Trie), end: false}
		}
		node = node.root[char]
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t
	for _, char := range word {
		if _, ok := node.root[char]; !ok {
			return false
		}
		node = node.root[char]
	}
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, char := range prefix {
		if _, ok := node.root[char]; !ok {
			return false
		}
		node = node.root[char]
	}
	return true
}
