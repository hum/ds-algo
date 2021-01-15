package ds

import (
	"fmt"
)

type Queue struct {
	first  *node
	last   *node
	length int
}

func (q *Queue) Peek() interface{} {
	return q.first.value
}

func (q *Queue) Enqueue(value interface{}) {
	node := &node{value: value}

	if q.length == 0 {
		q.first = node
	} else {
		q.last.next = node
	}
	q.last = node
	q.length++
}

func (q *Queue) Dequeue() {
	if q.length == 0 {
		return
	}
	q.first = q.first.next
	q.length--
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) String() string {
	result := ""
	current := q.first

	for current != nil {
		result += fmt.Sprintf("%v -> ", current.value)
		current = current.next
	}
	result += "\n"
	return result
}
