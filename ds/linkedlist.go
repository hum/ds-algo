package ds

import (
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

type LinkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *LinkedList) Append(value interface{}) {
	node := &node{value: value}

	if l.length == 0 {
		l.head = node
	} else {
		l.tail.next = node
	}

	l.tail = node
	l.length++
}

func (l *LinkedList) Prepend(value interface{}) {
	if l.length == 0 {
		l.Append(value)
		return
	}
	node := &node{value: value}

	node.next = l.head
	l.head = node
	l.length++
}

func (l *LinkedList) Insert(index int, value interface{}) {
	if index == 0 {
		l.Prepend(value)
		return
	} else if index+1 >= l.length {
		l.Append(value)
		return
	}

	node := &node{value: value}
	currentNode := l.traverseToIndex(index)

	node.next, currentNode.next = currentNode.next, node
	l.length++
}

func (l *LinkedList) Remove(index int) (interface{}, error) {
	if index > l.length-1 {
		return nil, fmt.Errorf("Index [%d] out of bounds.", index)
	}
	curr := l.traverseToIndex(index)
	deletedItem := curr.next
	curr.next = curr.next.next
	l.length--
	return deletedItem.value, nil
}

func (l *LinkedList) Get(index int) (interface{}, error) {
	if index > l.length-1 {
		return nil, fmt.Errorf("Index [%d] out of bounds.", index)
	}
	if index == 0 {
		return l.head.value, nil
	}
	return l.traverseToIndex(index).next.value, nil
}

func (l *LinkedList) traverseToIndex(index int) *node {
	if index == 0 {
		return l.head
	}

	if index >= l.length {
		return l.tail
	}

	current := l.head
	i := 0

	for i != index-1 {
		current = current.next
		i++
	}
	return current
}

func (l *LinkedList) Reverse() {
	if l.length == 1 {
		return
	}

	first := l.head
	second := first.next

	for second != nil {
		second.next, first, second = first, second, second.next
	}
	l.head.next = nil
	l.head = first
}

func (l *LinkedList) String() string {
	result := ""
	current := l.head

	for current != nil {
		result += fmt.Sprintf("%v -> ", current.value)
		current = current.next
	}

	result += "\n"
	return result
}
