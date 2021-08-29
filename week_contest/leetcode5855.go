package week_contest

import (
	"github.com/ladlod/leetcode/daily"
)

type NumStr string

func (this NumStr) Less(val daily.HeapV) bool {
	that := val.(NumStr)
	if len(this) != len(that) {
		return len(this) < len(that)
	}
	return string(this) < string(that)
}

func kthLargestNumber(nums []string, k int) string {
	heap := daily.InitHeap(k)
	for i := 0; i < k; i++ {
		heap.Insert(NumStr(nums[i]))
	}
	for j := k; j < len(nums); j++ {
		heap.Put(NumStr(nums[j]))
	}
	return string(heap.GetHead().(NumStr))
}
