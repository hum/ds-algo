package ds_test

import (
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestBinarySearchTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)

	searchedValue := 1
	_, err := tree.Lookup(searchedValue)
	if err != nil {
		t.Fatalf("There was an error looking up the value %d in the tree: %v", searchedValue, err)
	}

	searchedValue = 10000
	value, err := tree.Lookup(searchedValue)
	if err == nil {
		t.Fatalf("Tree should have returned an error, because the value %d is not in the tree. Expected=error, got=%v", searchedValue, value)
	}

}
