package ds_test

import (
  "testing"

  "github.com/hum/ds-algo/ds"
)

func TestGetFunction(t *testing.T) {
  arr := ds.Array{}

  tests := []struct {
    expected interface{}
  } {
    {1},
    {"1"},
  }

  arr.Push(1)
  arr.Push("1") 

  for i, tt := range tests {
    value := arr.Get(i)

    if value != tt.expected {
      t.Fatalf("Wrong value. Expected=%q, got %q", tt.expected, value)
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
  } {
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
    value := arr.Pop()

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
  value := arr.Delete(3)

  if value == arr.Get(3) && lengthBefore == arr.Length() {
    t.Fatalf("Error deleting element. Expected to delete the element %q from %s with length %d, but nothing changed.", value, arr.String(), arr.Length())
  }
}
