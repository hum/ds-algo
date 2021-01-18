package ds_test

import (
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestGetFunction(t *testing.T) {
	arr := ds.Array{}

	tests := []struct {
		expected interface{}
	}{
		{1},
		{"1"},
	}

	arr.Push(1)
	arr.Push("1")

	for i, tt := range tests {
		value, err := arr.Get(i)
		if err != nil {
			t.Fatalf("Index %d is out of range. Array length is: %d", i, arr.Length())
		}

		if value != tt.expected {
			t.Fatalf("Wrong value. Expected=%v, got %v", tt.expected, value)
		}
	}
}

func TestPopFunction(t *testing.T) {
	arr := ds.Array{}
	arr.Push(0)
	arr.Push(1)
	arr.Push(2)
	arr.Push("3")
	arr.Push(4)
	arr.Push(5)
	arr.Push("6")
	arr.Push(7)
	arr.Push("8")
	arr.Push(9)
	arr.Push(10)

	tests := []struct {
		expected interface{}
	}{
		{10},
		{9},
		{"8"},
		{7},
		{"6"},
		{5},
		{4},
		{"3"},
		{2},
		{1},
		{0},
	}

	for _, tt := range tests {
		value, err := arr.Pop()
		if err != nil {
			t.Fatalf("The array is empty. %v", err)
		}

		if value != tt.expected {
			t.Fatalf("Wrong popped value. Expected=%q, got %q", tt.expected, value)
		}
	}
}

func TestDeleteFunction(t *testing.T) {
	arr := ds.Array{}

	arr.Push(0)
	arr.Push("a")
	arr.Push("b")
	arr.Push("c")
	arr.Push("d")

	lengthBefore := arr.Length()
	value, err := arr.Delete(3)
	if err != nil {
		t.Fatalf("Index is out of range. Can't delete a value on index 3.")
	}

	v, err := arr.Get(3)
	if err != nil {
		t.Fatalf("Index is out of range. Can't get a value on index 3.")
	}

	if value == v && lengthBefore == arr.Length() {
		t.Fatalf("Error deleting element. Expected to delete the element %q from an array with length %d, but nothing changed.", value, arr.Length())
	}

	v, err = arr.Delete(500)
	if err == nil {
		t.Fatalf("Index 500 does not exist. The delete function should have thrown an error, but instead returned value %v", v)
	}
}
