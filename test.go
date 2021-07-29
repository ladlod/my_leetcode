package main

import "github.com/ladlod/leetcode/common_list"

func testLRUCache() {
	cache := common_list.Constructor(1)
	cache.Put(2, 1)
	cache.Get(2)
	cache.Put(3, 2)
	cache.Get(2)
	cache.Get(3)
	cache.Get(4)
}
