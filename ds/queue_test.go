package ds_test

import (
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestQueue(t *testing.T) {
	queue := ds.Queue{}

	values := [...]string{"google", "youtube", "discord"}

	queue.Enqueue(values[0])
	queue.Enqueue(values[1])
	queue.Enqueue(values[2])

	if queue.IsEmpty() {
		t.Fatalf("Queue is empty while it's supposed to have values.")
	}

	value := queue.Peek()
	if value != values[0] {
		t.Fatalf("Queue pushed the values incorrectly, expected %v, got %v.", values[0], value)
	}

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	if !queue.IsEmpty() {
		t.Fatalf("Queue should be empty, but isn't.")
	}
}
