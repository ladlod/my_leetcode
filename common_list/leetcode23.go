package common_list

import "fmt"

// 给你一个链表数组，每个链表都已经按升序排列。

// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

// 思路1：多路归并排序，时间复杂度O(mnlogn)
// 思路2：最小顶堆维护一个list头节点链表，时间复杂度O(mnlogn)
// m为单链表长度，n为链表数目

// 堆
type HeapListNode struct {
	heap []*ListNode
}

func (h *HeapListNode) insert(l *ListNode) {
	if h.heap != nil {
		h.heap = append(h.heap, l)
		h.shiftUp(len(h.heap) - 1)
	}
}

func (h *HeapListNode) put(l *ListNode) {
	if h.heap != nil && len(h.heap) > 0 {
		h.heap[0] = l
		h.shiftDown(0)
	}
}

func (h *HeapListNode) shiftUp(end int) {
	head := (end - 1) / 2
	if h.heap[head].Val > h.heap[end].Val {
		h.heap[head], h.heap[end] = h.heap[end], h.heap[head]
		h.shiftUp(head)
	}
}

func (h *HeapListNode) shiftDown(up int) {
	left := 2*up + 1
	right := 2 * (up + 1)
	if left < len(h.heap) && h.heap[left].Val < h.heap[up].Val {
		h.heap[left], h.heap[up] = h.heap[up], h.heap[left]
		h.shiftDown(left)
	}
	if right < len(h.heap) && h.heap[right].Val < h.heap[up].Val {
		h.heap[right], h.heap[up] = h.heap[up], h.heap[right]
		h.shiftDown(right)
	}
}

func (h *HeapListNode) pop() *ListNode {
	if h.heap != nil && len(h.heap) > 0 {
		res := h.heap[0]
		h.heap[0] = h.heap[len(h.heap)-1]
		h.heap = h.heap[:len(h.heap)-1]
		h.shiftDown(0)
		return res
	}
	return nil
}

func (h *HeapListNode) head() *ListNode {
	if h.heap == nil || len(h.heap) <= 0 {
		return nil
	}
	return h.heap[0]
}

func (h *HeapListNode) print() {
	fmt.Printf("[")
	for _, n := range h.heap {
		n.Print()
	}
	fmt.Printf("]\n")
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	heap := &HeapListNode{
		heap: make([]*ListNode, 0),
	}
	for _, list := range lists {
		if list != nil {
			heap.insert(list)
		}
	}
	heap.print()

	res := new(ListNode)

	for p := res; p != nil; p = p.Next {
		if heap.head() != nil && heap.head().Next != nil {
			p.Next = heap.head()
			heap.put(heap.head().Next)
		} else {
			p.Next = heap.pop()
		}
		heap.print()
	}

	return res.Next
}
