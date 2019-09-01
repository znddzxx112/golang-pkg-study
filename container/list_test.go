package containert

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	list := list.New()
	t.Log(list.Front())

	list.PushFront(1)
	e := list.Back()
	t.Log(e.Value)
	e2 := e.Next()
	t.Log(e2)

	list.Init()
	mark1 := list.PushFront(1)
	list.PushBack(4)
	mark3 := list.InsertAfter(3, mark1)
	list.InsertBefore(2, mark3)

	t.Log(list.Len())

	element := list.Front()
	for element != nil {
		t.Log(element.Value)
		element = element.Next()
	}

	for e := list.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}

	for e2 := list.Back(); e2 != nil; e2 = e2.Prev() {
		t.Log(e2.Value)
	}
}
