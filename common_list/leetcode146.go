package common_list

import (
	"fmt"
	"sync"
)

// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：

// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/lru-cache
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type LRUCache struct {
	mutex    *sync.Mutex
	cache    map[int]*LRUListNode
	head     *LRUListNode
	end      *LRUListNode
	capacity int
}

type LRUListNode struct {
	Pre  *LRUListNode
	Key  int
	Val  int
	Next *LRUListNode
}

func ConstructorLRU(capacity int) LRUCache {
	head := &LRUListNode{
		Key: -1,
	}
	end := &LRUListNode{
		Key: -2,
	}
	head.Next = end
	end.Pre = head

	return LRUCache{
		mutex:    new(sync.Mutex),
		cache:    make(map[int]*LRUListNode),
		head:     head,
		end:      end,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.capacity <= 0 {
		return -1
	}

	fmt.Printf("Get key:%d\n", key)
	if node, ok := this.cache[key]; ok {
		this.mutex.Lock()
		defer this.mutex.Unlock()
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre

		this.head.Next.Pre = node
		node.Next = this.head.Next
		this.head.Next = node
		node.Pre = this.head

		this.Print()
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	fmt.Printf("Put key:%d, val:%d\n", key, value)
	if node, ok := this.cache[key]; ok {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre

		this.head.Next.Pre = node
		node.Next = this.head.Next
		this.head.Next = node
		node.Pre = this.head

		node.Val = value
	} else {
		if len(this.cache) < this.capacity {
			node := &LRUListNode{
				Key: key,
				Val: value,
			}
			this.head.Next.Pre = node
			node.Next = this.head.Next
			this.head.Next = node
			node.Pre = this.head

			this.cache[key] = node
		} else {
			node := &LRUListNode{
				Key: key,
				Val: value,
			}
			// insert
			this.head.Next.Pre = node
			node.Next = this.head.Next
			this.head.Next = node
			node.Pre = this.head
			this.cache[key] = node

			// delete
			delete(this.cache, this.end.Pre.Key)
			this.end.Pre.Pre.Next = this.end
			this.end.Pre = this.end.Pre.Pre
		}
	}
	this.Print()
	return
}

func (this LRUCache) Print() {
	for key, val := range this.cache {
		fmt.Printf("key:%d, val:%d\n", key, val.Val)
	}

	fmt.Printf("head, key:%d, val:%d\n", this.head.Key, this.head.Val)
	fmt.Printf("end, key:%d, val:%d\n", this.end.Key, this.end.Val)
	for node := this.head; node != nil; node = node.Next {
		pre := 0
		if node.Pre != nil {
			pre = node.Pre.Key
		}
		next := 0
		if node.Next != nil {
			next = node.Next.Key
		}

		fmt.Printf("key:%d, val:%d, pre:%d, next:%d\n", node.Key, node.Val, pre, next)
	}

	fmt.Println("=============================")
}
