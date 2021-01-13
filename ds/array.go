package ds 

import "fmt"

type Array struct {
	length int
	data []interface{}
}

func (a *Array) Get(index int) interface{} {
	return a.data[index]
}

func (a *Array) Push(value interface{}) int {
	a.data = append(a.data, value)
	a.length++
	return a.length
}

func (a *Array) Pop() interface{} {
	lastItem := a.data[a.length-1]
	a.data = a.data[:a.length-1]
	a.length--
	return lastItem
}

func (a *Array) Delete(index int) interface{} {
	deletedItem := a.Get(index)
	a.shiftItems(index)
	a.Pop()
	return deletedItem
}

func (a *Array) shiftItems(index int) {
	for i := index; i < a.length - 1; i++ {
		a.data[i] = a.data[i+1]
	}
}

func (a *Array) String() string {
	return fmt.Sprintf("%v", a.data)
}

func (a *Array) Length() int {
	return a.length	
}
