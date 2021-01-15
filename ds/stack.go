package ds

import (
	"fmt"
)

// the node struct is already implemented in the linkedlist.go file
// there's no need to reimplement it

type Stack struct {
	top    *node
	bottom *node
	length int
}

func (s *Stack) Peek() interface{} {
	return s.top.value
}

func (s *Stack) Push(value interface{}) {
	node := &node{value: value}

	if s.length == 0 {
		s.top = node
		s.bottom = node
	} else {
		tmp := s.top
		s.top = node
		s.top.next = tmp
	}
	s.length++
}

func (s *Stack) Pop() {
	if s.length == 0 {
		return
	}

	if s.top == s.bottom {
		s.bottom = nil
	} else {
		s.top = s.top.next
	}
	s.length--
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) String() string {
	result := ""
	current := s.top

	for current != nil {
		result += fmt.Sprintf("%v -> ", current.value)
		current = current.next
	}
	return result
}
