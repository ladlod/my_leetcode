package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	trie := common_list.ConstructorTries()
	trie.Insert("apple")
	trie.Print()
	fmt.Println()
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
}
