package rbtree

import "testing"

type ValueInt int

func (value ValueInt) LessThan(j interface{}) bool {
	return value < j.(ValueInt)
}

func TestInsert(t *testing.T) {
	rbtree := NewTree()
	t.Log(rbtree.Size())

	var v6 ValueInt = 6
	rbtree.Insert(v6, 6)

	var v2 ValueInt = 2
	rbtree.Insert(v2, 2)

	var v3 ValueInt = 3
	rbtree.Insert(v3, 3)

	var v4 ValueInt = 4
	rbtree.Insert(v4, 4)

	var v5 ValueInt = 5
	rbtree.Insert(v5, 5)

	t.Log(rbtree)

	t.Log(rbtree.FindIt(v2))

	t.Log(rbtree.Find(v2))

	t.Log(rbtree.Size())
	rbtree.Delete(v3)
	t.Log(rbtree.Size())
}
