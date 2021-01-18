package ds_test

import (
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestDoublyLinkedList(t *testing.T) {
	list := ds.DoublyLinkedList{}

	list.Prepend(8)
	list.Append(7)
	list.Prepend(6)
	list.Insert(1, 90)
	list.Insert(2, 91)
	list.Insert(1, 92)
	value, err := list.Remove(1)
	if err != nil {
		t.Fatalf("Removing the 1st element returned an error %v", err)
	}
	if value != 92 {
		t.Fatalf("Returned value from the delete function does not match the expected value. Expected=%v, got=%v", 92, value)
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
			t.Fatalf("Index [%d] is out of range of the list.", i)
		}

		if value != tt.expected {
			t.Fatalf("Unexpected value in the Linked List lookup. Expected %v, but got %v", tt.expected, value)
		}
	}

	if value, err := list.Get(500); err == nil {
		t.Fatalf("Deleting the 500th element should have returned an error, but instead got value %v", value)
	}
}
