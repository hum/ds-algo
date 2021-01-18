package ds_test

import (
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestLinkedList(t *testing.T) {
	list := ds.LinkedList{}

	list.Prepend(8)
	list.Append(7)
	list.Prepend(6)
	list.Insert(1, 90)
	list.Insert(2, 91)
	list.Insert(1, 92)
	value, err := list.Remove(1)
	if err != nil {
		t.Fatalf("There was an error deleting an element, but the index should exist within the linked list.")
	}
	if value != 92 {
		t.Fatalf("The remove function has returned a wrong element. Expected %d, but got %d", 92, value)
	}

	value, err = list.Remove(100)
	if err == nil {
		t.Fatalf("Removing an item on the index of 100 did not return an error. Excepted error, but got %v", value)
	}

	tests := []struct {
		expected interface{}
	}{
		{6},
		{90},
		{91},
		{8},
		{7},
	}

	for i, tt := range tests {
		value, err := list.Get(i)
		if err != nil {
			t.Fatalf("List returned an out of bounds index error, but it's supposed to have data in it.")
		}

		if value != tt.expected {
			t.Fatalf("Unexpected value in the Linked List lookup. Expected %v, but got %v", tt.expected, value)
		}
	}

	value, err = list.Get(100)
	if err == nil {
		t.Fatalf("Getting the 100th element from a list did not return an error, but it should have.")
	}
}
