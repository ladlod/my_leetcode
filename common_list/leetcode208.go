package common_list

import "fmt"

// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

// 请你实现 Trie 类：

// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type Trie struct {
	v     byte
	nodes []*Trie
	end   bool
}

/** Initialize your data structure here. */
func ConstructorTries() Trie {
	return Trie{
		v:     0,
		nodes: make([]*Trie, 0),
		end:   false,
	}
}

func (this *Trie) Print() {
	fmt.Printf("%c ", this.v)
	for _, node := range this.nodes {
		node.Print()
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) <= 0 {
		return
	}
	for _, node := range this.nodes {
		if node.v == word[0] {
			if len(word) == 1 {
				node.end = true
			} else {
				node.Insert(word[1:])
			}
			return
		}
	}
	newNode := &Trie{
		v:     word[0],
		nodes: make([]*Trie, 0),
		end:   false,
	}
	if len(word) == 1 {
		newNode.end = true
	} else {
		newNode.Insert(word[1:])
	}
	this.nodes = append(this.nodes, newNode)
	return
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) <= 0 {
		return false
	}
	for _, node := range this.nodes {
		if node.v == word[0] {
			if len(word) == 1 && node.end {
				return true
			} else {
				if node.Search(word[1:]) {
					return true
				}
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) <= 0 {
		return true
	}
	for _, node := range this.nodes {
		if node.v == prefix[0] {
			if len(prefix) == 1 && !node.end {
				return true
			} else {
				if node.StartsWith(prefix[1:]) {
					return true
				}
			}
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
