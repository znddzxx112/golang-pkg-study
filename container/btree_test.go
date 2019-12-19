package containert

import (
	"fmt"
	"testing"
)

import "github.com/google/btree"

type TreeValueInt int

func (value TreeValueInt) Less(than btree.Item) bool {
	return value < than.(TreeValueInt)
}

func TestBtree(t *testing.T) {
	tree := btree.NewWithFreeList(2, btree.NewFreeList(4))
	var v6 TreeValueInt = 6
	fmt.Println(tree.ReplaceOrInsert(v6))
	fmt.Println(tree.ReplaceOrInsert(v6))

	var v2 TreeValueInt = 2
	fmt.Println(tree.ReplaceOrInsert(v2))

	var v3 TreeValueInt = 3
	fmt.Println(tree.ReplaceOrInsert(v3))

	var v4 TreeValueInt = 4
	fmt.Println(tree.ReplaceOrInsert(v4))

	var v5 TreeValueInt = 5
	fmt.Println(tree.ReplaceOrInsert(v5))
	fmt.Println(tree.ReplaceOrInsert(v5))

	fmt.Println("==========")
	tree.Ascend(func(i btree.Item) bool {
		fmt.Println(i.(TreeValueInt))
		return true
	})
	fmt.Println("==========")
	tree.AscendGreaterOrEqual(v3, func(item btree.Item) bool {
		fmt.Println(item.(TreeValueInt))
		return true
	})
	fmt.Println("==========")

	tree.AscendRange(v2, v5, func(item btree.Item) bool {
		fmt.Println(item.(TreeValueInt))
		return true
	})
	fmt.Println("==========")
	t.Log(tree.Get(v5))
	t.Log(tree.Has(v5))
	t.Log(tree.Delete(v5))
	t.Log(tree.Has(v5))
	t.Log(tree.Get(v5))
	fmt.Println("==========")
	tree.Descend(func(item btree.Item) bool {
		fmt.Println(item.(TreeValueInt))
		return true
	})
	t.Log("over")

}
