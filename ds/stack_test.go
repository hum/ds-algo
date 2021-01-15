package ds_test

import (
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestStackPushPop(t *testing.T) {
	stack := ds.Stack{}

	values := [...]string{"google", "youtube", "discord"}

	stack.Push(values[0])
	stack.Push(values[1])
	stack.Push(values[2])

	if stack.IsEmpty() {
		t.Fatalf("Stack is empty while it's supposed to have values.")
	}

	value := stack.Peek()
	if value != values[2] {
		t.Fatalf("Stack pushed the values incorrectly, expected %v, got %v.", values[2], value)
	}

	stack.Pop()
	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Fatalf("Stack should be empty, but isn't.")
	}
}
