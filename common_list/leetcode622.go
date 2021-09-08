package common_list

import "sync"

// 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

// 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

// 你的实现应该支持如下操作：

// MyCircularQueue(k): 构造器，设置队列长度为 k 。
// Front: 从队首获取元素。如果队列为空，返回 -1 。
// Rear: 获取队尾元素。如果队列为空，返回 -1 。
// enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
// deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
// isEmpty(): 检查循环队列是否为空。
// isFull(): 检查循环队列是否已满。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/design-circular-queue
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MyCircularQueue struct {
	queue []int
	head  int
	end   int
	mutex *sync.Mutex
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		head:  0,
		end:   0,
		mutex: new(sync.Mutex),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.mutex.Lock()
		defer this.mutex.Unlock()

		this.queue[this.end%len(this.queue)] = value
		this.end++
		return true
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.mutex.Lock()
		defer this.mutex.Unlock()

		this.head++
		return true
	}
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.queue[this.head%len(this.queue)]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.queue[(this.end-1)%len(this.queue)]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return (this.end - this.head) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.end - this.head) == len(this.queue)
}
