package BST

import (
	"fmt"
)

type Comparator func(interface{}, interface{}) bool

type BinTreeValue interface{}

type BinTreeNode struct {
	leftChild  *BinTreeNode
	rightChild *BinTreeNode

	value BinTreeValue
}

type BinTree struct {
	root       *BinTreeNode
	comparator Comparator
}

func New(cmp Comparator) *BinTree {
	return &BinTree{
		root:       nil,
		comparator: cmp,
	}
}

//
func (b *BinTree) find(value BinTreeValue) (*BinTreeNode, **BinTreeNode) {
	found_node := b.root
	found_node_ptr := &b.root

	for found_node != nil {
		if b.comparator(found_node.value, value) {
			found_node_ptr = &found_node.leftChild
			found_node = found_node.leftChild
		} else if found_node.value == value {
			return found_node, found_node_ptr
		} else {
			found_node_ptr = &found_node.rightChild
			found_node = found_node.rightChild
		}
	}

	return found_node, found_node_ptr
}

func (b *BinTree) Find(value BinTreeValue) bool {
	found_node, _ := b.find(value)
	return found_node != nil
}

func (b *BinTree) Add(value BinTreeValue) {
	found_node, found_node_ptr := b.find(value)

	if found_node == nil {
		*found_node_ptr = &BinTreeNode{
			value: value,
		}
	}

}

func (b *BinTree) Del(value interface{}) {
	found_node, found_node_ptr := b.find(value)

	if found_node == nil { // nothing to delete
		return
	}

	if found_node.leftChild != nil && found_node.rightChild != nil { //deleting hode has two children
		new_current := found_node.rightChild

		if new_current != nil {
			var new_current_parent *BinTreeNode

			for new_current.leftChild != nil {
				new_current_parent = new_current
				new_current = new_current.leftChild
			}

			new_current_parent.leftChild = new_current.rightChild
			new_current.leftChild = found_node.leftChild
			new_current.rightChild = found_node.rightChild
		}
		*found_node_ptr = new_current

	} else { //deleting node has one or less children
		if found_node.leftChild != nil {
			*found_node_ptr = found_node.leftChild
		} else if found_node.rightChild != nil {
			*found_node_ptr = found_node.rightChild
		} else {
			*found_node_ptr = nil
		}
	}
}

func (b *BinTree) Print() {
	b.print(b.root)
}

func (b *BinTree) print(node *BinTreeNode) {
	if node == nil {
		return
	}

	if node.leftChild != nil {
		fmt.Printf("%d -> %d\n", node.value.(int), node.leftChild.value.(int))
		b.print(node.leftChild)
	}

	if node.rightChild != nil {
		fmt.Printf("%d -> %d\n", node.value.(int), node.rightChild.value.(int))
		b.print(node.rightChild)
	}
}
