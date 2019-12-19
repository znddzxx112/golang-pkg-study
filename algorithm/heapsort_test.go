package algorithm

import (
	"container/heap"
	"fmt"
	"testing"
)

func up(h heap.Interface, j int) {
	for {
		i0 := (j - 1) / 2
		if i0 < 0 {
			break
		}

		if i0 == j || !h.Less(j, i0) {
			break
		}

		h.Swap(j, i0)
		j = i0
	}
}

func Pop(h heap.Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Push(h heap.Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Remove(h heap.Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

func Fix(h heap.Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func down(h heap.Interface, i0, n int) bool {
	i := i0
	for {
		leftNode := i*2 + 1

		if leftNode >= n || leftNode < 0 {
			break
		}

		minNode := leftNode
		if rightNode := leftNode + 1; rightNode < n && h.Less(rightNode, leftNode) {
			minNode = rightNode
		}

		if !h.Less(minNode, i0) {
			break
		}

		h.Swap(minNode, i0)
		i = minNode
	}
	return i > i0
}

func Init(h heap.Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		fmt.Println(i)
	}
}

func TestHeapInit(t *testing.T) {
	h := &IntHeap{2, 7, 5, 6, 1}
	Init(h)
	t.Log(h)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// cd /usr/local
// mv go go1.12.1
// sudo curl -sSL "https://studygolang.com/dl/golang/go1.13.linux-amd64.tar.gz" -o go1.13.tar.gz --progress
// sudo tar -zxvf go1.13.tar.gz  -C /usr/local/
// go env -w GOSUMDB=off
// go env -w GOPROXY="https://goproxy.cn,direct"
//
// not find bin/godoc
// not find go/misc/git
//
// go help
// 	module-auth module authentication using go.sum
//	module-private module configuration for non-public modules
