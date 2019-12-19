package algorithm

import (
	"sort"
	"testing"
)

func SelectionSort(data sort.Interface) {
	length := data.Len()
	var i, j, t int
	for i = 0; i < length; i++ {
		t = i
		for j = i + 1; j < length; j++ {
			if data.Less(i, j) {
				t = j
			}
		}
		if t != i {
			data.Swap(i, t)
		}
	}
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestSelectionSort(t *testing.T) {
	ints := IntSlice{3, 5, 4, 7}
	SelectionSort(ints)
	t.Log(ints)
}

func BubbleSort(data sort.Interface) {
	length := data.Len()
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	ints := IntSlice{3, 5, 4, 7}
	BubbleSort(ints)
	t.Log(ints)
}

func InsertionSort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		for j := i + 1; j > 0; j-- {
			if data.Less(j-1, j) {
				data.Swap(j-1, j)
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {
	ints := IntSlice{3, 5, 4, 7}
	InsertionSort(ints)
	t.Log(ints)
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func TestMaxDepth(t *testing.T) {
	t.Log(maxDepth(1))
	t.Log(maxDepth(2))
	t.Log(maxDepth(7))
	t.Log(maxDepth(8))
	t.Log(maxDepth(9))
	t.Logf("%d", 8>>1)
	t.Logf("%d", 4>>1)
	t.Logf("%d", 2>>1)
	t.Logf("%d", 1>>1)
}
