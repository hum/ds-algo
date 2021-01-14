package ds

import (
	"fmt"
)

type node2 struct {
	value interface{}
	next  *node2
	prev  *node2
}

type DoublyLinkedList struct {
	head   *node2
	tail   *node2
	length int
}

func (l *DoublyLinkedList) Append(value interface{}) {
	node := &node2{value: value}

	if l.length == 0 {
		l.head = node
	} else {
		l.tail.next = node
	}

	node.prev = l.tail
	l.tail = node
	l.length++
}

func (l *DoublyLinkedList) Prepend(value interface{}) {
	if l.length == 0 {
		l.Append(value)
		return
	}
	node := &node2{value: value}

	l.head.prev = node
	node.next = l.head

	l.head = node
	l.length++
}

func (l *DoublyLinkedList) Insert(index int, value interface{}) {
	if index == 0 {
		l.Prepend(value)
		return
	} else if index >= l.length {
		l.Append(value)
		return
	}

	node := &node2{value: value}
	leader := l.traverseToIndex(index)

	node.prev = leader
	node.next, leader.next = leader.next, node
	node.next.prev = node

	l.length++
}

func (l *DoublyLinkedList) Remove(index int) {
	curr := l.traverseToIndex(index)
	curr.next, curr.next.prev = curr.next.next, curr
	l.length--
}

func (l *DoublyLinkedList) Get(index int) interface{} {
	if index == 0 {
		return l.head.value
	}

	return l.traverseToIndex(index).next.value
}

func (l *DoublyLinkedList) traverseToIndex(index int) *node2 {
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

func (l *DoublyLinkedList) String() string {
	result := ""
	current := l.head

	for current != nil {
		result += fmt.Sprintf("%v -> ", current.value)
		current = current.next
	}

	result += "\n"
	return result
}
