package ds

import "fmt"

type Array struct {
	length int
	data   []interface{}
}

func (a *Array) Get(index int) (interface{}, error) {
	if index > a.length-1 {
		return nil, fmt.Errorf("Index out of range.")
	}
	return a.data[index], nil
}

func (a *Array) Push(value interface{}) int {
	a.data = append(a.data, value)
	a.length++
	return a.length
}

func (a *Array) Pop() (interface{}, error) {
	if a.length == 0 {
		return nil, fmt.Errorf("There is nothing left to pop. The length of the array is 0.")
	}
	lastItem := a.data[a.length-1]
	a.data = a.data[:a.length-1]
	a.length--
	return lastItem, nil
}

func (a *Array) Delete(index int) (interface{}, error) {
	deletedItem, err := a.Get(index)
	if err != nil {
		return nil, err
	}
	a.shiftItems(index)
	a.Pop()
	return deletedItem, nil
}

func (a *Array) shiftItems(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
}

func (a *Array) Length() int {
	return a.length
}
