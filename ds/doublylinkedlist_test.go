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
	list.Remove(1)

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
		value := list.Get(i)

		if value != tt.expected {
			t.Fatalf("Unexpected value in the Linked List lookup. Expected %v, but got %v", tt.expected, value)
		}
	}
}
