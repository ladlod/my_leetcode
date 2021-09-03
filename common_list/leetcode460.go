package common_list

import (
	"sync"
)

// 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

// 实现 LFUCache 类：

// LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
// int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。
// void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
// 注意「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。

// 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

// 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/lfu-cache
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 数据结构
// 使用一个map[key]node记录基础kv结构，一个map[cnt]list 记录访问次数到对应的双向链表，一个mincnt记录最低访问次数
// inc操作时，根据key找到对应的node和对应的双向链表，将该值从该双向链表中移除，移到访问次数+1的双向链表中，如果该链表移除后为空，则清空map，更新mincnt
// insert操作时，将对应的node放到cache以及访问次数为1的双向链表表头，同时更新最低访问次数为1
// put操作时，寻找mincnt对应的双向链表的最后一个节点并从两个map中删除，再进行insert操作

type LFUCache struct {
	cache    map[int]*LFUNode // 存储k-v
	relation map[int]*LFUSt   // 存储访问次数，到双向链表
	capacity int
	mincnt   int // 记录最小访问次数
	mutex    *sync.Mutex
}

type LFUSt struct {
	cnt  int
	head *LFUNode
	end  *LFUNode
}

func NewLFUSt(cnt int) *LFUSt {
	head := &LFUNode{}
	end := &LFUNode{}
	head.next, end.pre = end, head
	return &LFUSt{
		cnt:  cnt,
		head: head,
		end:  end,
	}
}

type LFUNode struct {
	key int
	val int
	cnt int

	next *LFUNode
	pre  *LFUNode
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*LFUNode),
		relation: make(map[int]*LFUSt),
		mincnt:   0,

		capacity: capacity,
		mutex:    new(sync.Mutex),
	}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.mutex.Lock()
		defer this.mutex.Unlock()
		this.inc(v)
		return v.val
	}

	return -1
}

func (this *LFUCache) inc(node *LFUNode) {
	// 寻找双向链表
	lfuSt := this.relation[node.cnt]

	// 从双向链表中删除节点
	node.pre.next, node.next.pre = node.next, node.pre

	// 如果双向链表已空，删除双向链表，更新最小值
	if lfuSt.head.next == lfuSt.end {
		delete(this.relation, lfuSt.cnt)
		if lfuSt.cnt == this.mincnt {
			this.mincnt++
		}
	}

	node.cnt++
	// 初始化map
	if _, ok := this.relation[node.cnt]; !ok {
		this.relation[node.cnt] = NewLFUSt(node.cnt)
	}

	// 将节点插入新双向链表
	newLfu := this.relation[node.cnt]
	newLfu.head.next, newLfu.head.next.pre, node.pre, node.next = node, node, newLfu.head, newLfu.head.next
}

func (this *LFUCache) insert(node *LFUNode) {
	// 插入cache
	this.cache[node.key] = node

	// 初始化map
	if _, ok := this.relation[1]; !ok {
		this.relation[node.cnt] = NewLFUSt(1)
		// 更新最小访问次数值
		this.mincnt = 1
	}

	// 将节点插入访问数为1的双向链表
	newLfu := this.relation[1]
	newLfu.head.next, newLfu.head.next.pre, node.pre, node.next = node, node, newLfu.head, newLfu.head.next
}

func (this *LFUCache) Put(key int, value int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if v, ok := this.cache[key]; ok {
		// update
		v.val = value
		this.inc(v)
		return
	} else {
		if len(this.cache) < this.capacity {
			// insert
			this.insert(&LFUNode{
				key: key,
				val: value,
				cnt: 1,
			})
		} else {
			// put
			if this.mincnt == 0 {
				return
			}
			lfuSt := this.relation[this.mincnt]
			last := lfuSt.end.pre
			delete(this.cache, last.key)
			last.pre.next, last.next.pre = last.next, last.pre
			if lfuSt.head.next == lfuSt.end {
				delete(this.relation, lfuSt.cnt)
			}
			this.insert(&LFUNode{
				key: key,
				val: value,
				cnt: 1,
			})
		}
	}
}
