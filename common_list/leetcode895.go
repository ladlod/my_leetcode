package common_list

// 实现 FreqStack，模拟类似栈的数据结构的操作的一个类。

// FreqStack 有两个函数：

// push(int x)，将整数 x 推入栈中。
// pop()，它移除并返回栈中出现最频繁的元素。
// 如果最频繁的元素不只一个，则移除并返回最接近栈顶的元素。

// 示例：

// 输入：
// ["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
// [[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
// 输出：[null,null,null,null,null,null,null,5,7,5,4]
// 解释：
// 执行六次 .push 操作后，栈自底向上为 [5,7,5,7,4,5]。然后：

// pop() -> 返回 5，因为 5 是出现频率最高的。
// 栈变成 [5,7,5,7,4]。

// pop() -> 返回 7，因为 5 和 7 都是频率最高的，但 7 最接近栈顶。
// 栈变成 [5,7,5,4]。

// pop() -> 返回 5 。
// 栈变成 [5,7,4]。

// pop() -> 返回 4 。
// 栈变成 [5,7]。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-frequency-stack
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 设计思路
// 考虑题目要求，若仅用一个栈来实现，那么pop的时间复杂度为O(n)即可暴力实现
// 若需要优化时间复杂度，则至少需要一个结构来记录每个元素出现的次数，所以我们记录一个freMap[int]int来记录每个元素出现的次数，并且记录一个出现过的最大访问次数
// 对于同一出现频率的元素，我们需要记录每个元素入栈的先后顺序，即记录一个出现频率层级的map[int]stack
type FreqStack struct {
	freMap map[int]int
	stack  map[int][]int
	maxFre int
}

func ConstructorFreStack() FreqStack {
	return FreqStack{
		freMap: make(map[int]int),
		stack:  make(map[int][]int),
		maxFre: 0,
	}
}

func (this *FreqStack) Push(val int) {
	this.freMap[val]++ // 频率++
	if this.freMap[val] > this.maxFre {
		this.maxFre = this.freMap[val] // 更新最大频率
	}
	if _, ok := this.stack[this.freMap[val]]; !ok {
		this.stack[this.freMap[val]] = make([]int, 0)
	}
	this.stack[this.freMap[val]] = append(this.stack[this.freMap[val]], val) // 新增至频率层级

	return
}

func (this *FreqStack) Pop() int {
	if this.maxFre <= 0 {
		return 0
	}

	v := this.stack[this.maxFre][len(this.stack[this.maxFre])-1] // 从最高频率层级出栈元素
	this.stack[this.maxFre] = this.stack[this.maxFre][:len(this.stack[this.maxFre])-1]

	this.freMap[v]-- // 更新元素访问频率
	if len(this.stack[this.maxFre]) == 0 {
		delete(this.stack, this.maxFre)
		this.maxFre-- // 更新当前最大访问频率
	}

	return v
}
