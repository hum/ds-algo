package ds

import (
	"fmt"
)

type BinarySearchTree struct {
	root *binaryTreeNode
}

type binaryTreeNode struct {
	value int
	left  *binaryTreeNode
	right *binaryTreeNode
}

func (t *BinarySearchTree) Insert(value int) {
	if t.root == nil {
		t.root = &binaryTreeNode{value: value}
		return
	}

	curr := t.root
	node := &binaryTreeNode{value: value}

	for {
		if value < curr.value {
			if curr.left == nil {
				curr.left = node
				break
			}
			curr = curr.left
		} else if value > curr.value {
			if curr.right == nil {
				curr.right = node
				break
			}
			curr = curr.right
		} else if value == curr.value {
			break
		}
	}

}

func (t *BinarySearchTree) Lookup(value int) (bool, error) {
	if t.root == nil {
		return false, fmt.Errorf("Value [%d] could not be searched, because the tree has 0 nodes.", value)
	}

	curr := t.root

	for curr != nil {
		if value < curr.value {
			curr = curr.left
		} else if value > curr.value {
			curr = curr.right
		} else {
			return true, nil
		}
	}
	return false, fmt.Errorf("Value [%d] was not found in the tree.", value)
}
