package containert

import "testing"

// https://github.com/sakeven/RbTree
const RED bool = true
const BLACK bool = false

type Interface interface {
	Less(j Interface) bool
	Equal(j Interface) bool
}

type RBNode struct {
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
	Color  bool
	Value  Interface
}

func NewRBNode(value Interface) *RBNode {
	node := new(RBNode)
	node.Parent = nil
	node.Left = nil
	node.Right = nil
	node.Color = RED
	node.Value = value
	return node
}

func (rbnode *RBNode) getParent() *RBNode {
	return rbnode.Parent
}

func (rbnode *RBNode) getGrandParent() *RBNode {
	if rbnode.Parent != nil {
		return rbnode.Parent.Parent
	}
	return nil

}

func (rbnode *RBNode) getUncle() *RBNode {
	parentNode := rbnode.Parent
	if parentNode == nil {
		return nil
	}

	grandParent := parentNode.Parent
	if grandParent == nil {
		return nil
	}

	if grandParent.Left == parentNode {
		return grandParent.Right
	}
	return grandParent.Left
}

func (rbnode *RBNode) getSibling() *RBNode {
	parentNode := rbnode.Parent
	if parentNode == nil {
		return nil
	}
	if parentNode.Left == rbnode {
		return parentNode.Right
	}
	return parentNode.Left
}

type RBTree struct {
	root  *RBNode
	count uint64
}

func NewRBTree() *RBTree {
	rbtree := new(RBTree)
	rbtree.root = nil
	rbtree.count = 0
	return rbtree
}

func (rbtree *RBTree) Insert(node *RBNode) {

	if rbtree.root == nil {
		// case 1: insert root Node
		rbtree.root = node
		node.Color = BLACK
		rbtree.count++
		return
	}

	x := rbtree.root

	var y *RBNode
	for x != nil {
		y = x
		if node.Value.Less(x.Value) {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = y
	if node.Value.Less(y.Value) {
		y.Left = node
	} else {
		y.Right = node
	}

	rbtree.insertFixup(node)
	rbtree.count++
}

func (rbtree *RBTree) insertFixup(node *RBNode) {
	var parent, uncle, grandParent *RBNode
	if parent = node.getParent(); parent.Color == BLACK {
		// case 2: parent.Color is Black, do nothing
		return
	}
	for ; parent != nil && parent.Color == RED; parent = node.getParent() {
		grandParent = node.getGrandParent()
		if grandParent == nil {
			return
		}

		if parent == grandParent.Left {
			uncle = grandParent.Right
			if uncle != nil && uncle.Color == RED {
				// case 3: parent And uncle all is Red
				// parent And uncle turn to Black
				// grandParent turn to Red
				parent.Color, uncle.Color = BLACK, BLACK
				grandParent.Color = RED
				node = grandParent
				continue
			}

			if node == parent.Right {
				// case 4: parent is Red, uncle is Black
				// parent.Right is node
				// turn to case 5
				// note: node and arent is changed each other
				rbtree.leftRotate(parent)
			}
			// case 5: parent is Red, uncle is Black
			// parent.Left is node
			grandParent.Color = RED
			grandParent.Left.Color = BLACK // parent
			rbtree.rightRotate(grandParent)

		} else { // parent == grandParent.Right
			uncle = grandParent.Left
			if uncle != nil && uncle.Color == RED {
				// case 3: parent And uncle all is Red
				// parent And uncle turn to Black
				// grandParent turn to Red
				parent.Color, uncle.Color = BLACK, BLACK
				grandParent.Color = RED
				node = grandParent
				continue
			}

			if node == parent.Left {
				// case 4: parent is Red, uncle is Black
				// parent.Right is node
				// turn to case 5
				// note: node and arent is changed each other
				rbtree.rightRotate(parent)
			}
			// case 5: parent is Red, uncle is Black
			// parent.Left is node
			grandParent.Color = RED
			grandParent.Right.Color = BLACK // parent
			rbtree.leftRotate(grandParent)
		}
	}
	if rbtree.root != nil && rbtree.root.Color == RED {
		parent.Color = BLACK
	}
}

func (rbtree *RBTree) leftRotate(node *RBNode) {
	son := node.Right
	son.Parent, node.Right = node.Parent, son.Left

	if son.Left == nil {
		son.Left = node
	} else {
		grandSon := son.Left
		son.Left, grandSon.Parent = node, node
	}

	parent := node.Parent
	if parent == nil { // node is root
		rbtree.root, node.Parent = son, son
	} else if node == parent.Left {
		parent.Left, node.Parent = son, son
	} else { // node == parent.Right
		parent.Right, node.Parent = son, son
	}
}

func (rbtree *RBTree) rightRotate(node *RBNode) {
	son := node.Left
	son.Parent, node.Left = node.Parent, son.Right

	if son.Right == nil {
		son.Right = node
	} else {
		grandSon := son.Right
		son.Right, grandSon.Parent = node, node
	}

	parent := node.Parent
	if parent == nil { // node is root
		rbtree.root, node.Parent = son, son
	} else if node == parent.Left {
		parent.Left, node.Parent = son, son
	} else { // node == parent.Right
		parent.Right, node.Parent = son, son
	}
}

func (rbtree *RBTree) FindNode(j Interface) *RBNode {
	return rbtree.findNode(j)
}

func (rbtree *RBTree) findNode(j Interface) *RBNode {
	x := rbtree.root
	for x != nil {
		if x.Value.Equal(j) {
			return x
		} else if x.Value.Less(j) { // x.Value < j
			x = x.Left
		} else { // x.Value > j
			x = x.Right
		}
	}
	return nil
}

// ========== for Test
type ValueInt int

func (value ValueInt) Less(j Interface) bool {
	return value < j.(ValueInt)
}

func (value ValueInt) Equal(j Interface) bool {
	return value == j.(ValueInt)
}

func TestInsert(t *testing.T) {
	rbtree := NewRBTree()
	var v6 ValueInt = 6
	rbtree.Insert(NewRBNode(v6))

	var v2 ValueInt = 2
	rbtree.Insert(NewRBNode(v2))

	var v3 ValueInt = 3
	rbtree.Insert(NewRBNode(v3))

	var v4 ValueInt = 4
	rbtree.Insert(NewRBNode(v4))

	var v5 ValueInt = 5
	rbtree.Insert(NewRBNode(v5))

	t.Log(rbtree)
}
