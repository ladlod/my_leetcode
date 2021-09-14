package common_list

import "fmt"

// 给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。

// 思路
// 小顶堆记录当前最短板，
// 堆顶出堆，入堆堆四个方向的板子，res += 堆顶高度-入堆的板子高度
func TrapRainWater(heightMap [][]int) int {
	if len(heightMap) <= 0 {
		return 0
	}
	heap := &Heap{
		heap: make([]*Pos, 0),
	}

	// 记录节点是否遍历过
	flag := make([][]bool, len(heightMap))
	for i := 0; i < len(heightMap); i++ {
		flag[i] = make([]bool, len(heightMap[i]))
	}

	// 外圈入堆
	for i := 0; i < len(heightMap[0]); i++ {
		heap.insert(&Pos{
			val: heightMap[0][i],
			x:   0,
			y:   i,
		})
		heap.insert(&Pos{
			val: heightMap[len(heightMap)-1][i],
			x:   len(heightMap) - 1,
			y:   i,
		})
		flag[0][i] = true
		flag[len(heightMap)-1][i] = true
	}
	for j := 1; j < len(heightMap)-1; j++ {
		heap.insert(&Pos{
			val: heightMap[j][0],
			x:   j,
			y:   0,
		})
		heap.insert(&Pos{
			val: heightMap[j][len(heightMap[0])-1],
			x:   j,
			y:   len(heightMap[0]) - 1,
		})
		flag[j][0] = true
		flag[j][len(heightMap[0])-1] = true
	}

	ans := 0
	fmt.Println(flag, ans)
	for len(heap.heap) > 0 {
		top := heap.pop()
		if top.x > 0 && !flag[top.x-1][top.y] {
			flag[top.x-1][top.y] = true
			if heightMap[top.x-1][top.y] < top.val {
				ans += top.val - heightMap[top.x-1][top.y]
				heap.insert(&Pos{
					val: top.val,
					x:   top.x - 1,
					y:   top.y,
				})
			} else {
				heap.insert(&Pos{
					val: heightMap[top.x-1][top.y],
					x:   top.x - 1,
					y:   top.y,
				})
			}
		}

		if top.y > 0 && !flag[top.x][top.y-1] {
			flag[top.x][top.y-1] = true
			if heightMap[top.x][top.y-1] < top.val {
				ans += top.val - heightMap[top.x][top.y-1]
				heap.insert(&Pos{
					val: top.val,
					x:   top.x,
					y:   top.y - 1,
				})
			} else {
				heap.insert(&Pos{
					val: heightMap[top.x][top.y-1],
					x:   top.x,
					y:   top.y - 1,
				})
			}
		}

		if top.x+1 < len(heightMap) && !flag[top.x+1][top.y] {
			flag[top.x+1][top.y] = true
			if heightMap[top.x+1][top.y] < top.val {
				ans += top.val - heightMap[top.x+1][top.y]
				heap.insert(&Pos{
					val: top.val,
					x:   top.x + 1,
					y:   top.y,
				})
			} else {
				heap.insert(&Pos{
					val: heightMap[top.x+1][top.y],
					x:   top.x + 1,
					y:   top.y,
				})
			}
		}

		if top.y+1 < len(heightMap[0]) && !flag[top.x][top.y+1] {
			flag[top.x][top.y+1] = true
			if heightMap[top.x][top.y+1] < top.val {
				ans += top.val - heightMap[top.x][top.y+1]
				heap.insert(&Pos{
					val: top.val,
					x:   top.x,
					y:   top.y + 1,
				})
			} else {
				heap.insert(&Pos{
					val: heightMap[top.x][top.y+1],
					x:   top.x,
					y:   top.y + 1,
				})
			}
		}
		fmt.Println(flag, ans)
	}

	return ans
}

type Heap struct {
	heap []*Pos
}

type Pos struct {
	x   int
	y   int
	val int
}

func (h *Heap) insert(v *Pos) {
	h.heap = append(h.heap, v)
	h.shiftUp(len(h.heap) - 1)
}

func (h *Heap) put(v *Pos) {
	h.heap[0] = v
	h.shiftDown(0)
}

func (h *Heap) pop() *Pos {
	res := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.shiftDown(0)
	return res
}

func (h *Heap) shiftDown(i int) {
	l := i*2 + 1
	r := i*2 + 2
	if l < len(h.heap) && h.heap[l].val < h.heap[i].val {
		h.heap[i], h.heap[l] = h.heap[l], h.heap[i]
		h.shiftDown(l)
	}
	if r < len(h.heap) && h.heap[r].val < h.heap[i].val {
		h.heap[i], h.heap[r] = h.heap[r], h.heap[i]
		h.shiftDown(r)
	}
}

func (h *Heap) shiftUp(i int) {
	head := (i - 1) / 2
	if head >= 0 && h.heap[head].val > h.heap[i].val {
		h.heap[head], h.heap[i] = h.heap[i], h.heap[head]
		h.shiftUp(head)
	}
}
