package ds

import (
	"fmt"
)

type StackArray struct {
	data []interface{}
}

func (s *StackArray) Peek() interface{} {
	return s.data[len(s.data)-1]
}

func (s *StackArray) Push(value interface{}) {
	if len(s.data) == 0 {
		s.data = make([]interface{}, 0)
	}
	s.data = append(s.data, value)
}

func (s *StackArray) Pop() {
	if len(s.data) == 0 {
		return
	}

	s.data = s.data[:len(s.data)-1]
}

func (s *StackArray) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *StackArray) String() string {
	result := ""
	for i := len(s.data) - 1; i > 0; i-- {
		result += fmt.Sprintf("%v -> ", s.data[i])
	}
	return result
}
