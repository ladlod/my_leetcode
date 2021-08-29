package daily

// 堆
type HeapV interface {
	Less(val HeapV) bool
}

type Heap struct {
	heap []HeapV
}

func InitHeap(cap int) *Heap {
	return &Heap{
		heap: make([]HeapV, 0),
	}
}

func (h *Heap) Insert(val HeapV) {
	h.heap = append(h.heap, val)
	h.shiftUp(len(h.heap) - 1)
}

func (h *Heap) Put(val HeapV) {
	if len(h.heap) <= 0 {
		return
	}
	if h.heap[0].Less(val) { // 最小顶堆，put的数字比堆顶大时，替换该数字
		h.heap[0] = val
		h.shiftDown(0)
	}
}

func (h *Heap) shiftUp(end int) {
	head := (end - 1) / 2
	if h.heap[end].Less(h.heap[head]) {
		h.heap[head], h.heap[end] = h.heap[end], h.heap[head]
		h.shiftUp(head)
	}
}

func (h *Heap) shiftDown(head int) {
	left := 2*head + 1
	right := 2*head + 2
	if left < len(h.heap) && h.heap[left].Less(h.heap[head]) {
		h.heap[left], h.heap[head] = h.heap[head], h.heap[left]
		h.shiftDown(left)
	}
	if right < len(h.heap) && h.heap[right].Less(h.heap[head]) {
		h.heap[right], h.heap[head] = h.heap[head], h.heap[right]
		h.shiftDown(right)
	}
	return
}

func (h *Heap) GetHead() HeapV {
	if len(h.heap) > 0 {
		return h.heap[0]
	}
	return nil
}
