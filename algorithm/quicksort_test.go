package algorithm

import (
	"sort"
	"testing"
)

func partition(data sort.Interface, head, tail int) int {
	data.Swap(head, tail)
	for i := head; i < tail; i++ {
		if data.Less(tail, i) {
			data.Swap(i, head)
			head++
		}
	}
	data.Swap(head, tail)
	return head
}

func quickSort(data sort.Interface, head, tail int) {
	if head >= tail {
		return
	}

	pivot := partition(data, head, tail)
	if pivot > 0 {
		quickSort(data, head, pivot-1)
	}
	if pivot+1 < data.Len() {
		quickSort(data, pivot+1, tail)
	}
}

func TestQuickSort(t *testing.T) {
	ints := IntSlice{3, 5, 4, 7, 6, 4, 8, 12, 10}
	quickSort(ints, 0, ints.Len()-1)
	t.Log(ints)
}
