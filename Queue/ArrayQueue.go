package Queue

import (
	"data_structure/Array"
	"fmt"
)

type arrayQueue struct {
	array Array.IArray
}

//构造函数
func NewArrayQueue(capacity int) Queue {
	return &arrayQueue{
		array: Array.NewArray(capacity),
	}
}

func (a *arrayQueue) GetSize() int {
	return a.array.GetSize()
}

func (a *arrayQueue) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *arrayQueue) Getcap() int {
	return a.array.GetCap()
}

func (a *arrayQueue) Enqueue(e interface{}) {
	a.array.AddLast(e)
}

func (a *arrayQueue) Dequeue() interface{} {
	return a.array.DelectFirst()
}

func (a *arrayQueue) GetFront() interface{} {
	return a.array.GetFrist()
}

func (a *arrayQueue) String() string {
	var str string
	str = fmt.Sprintf("Queue :front %v \n", a.array.GetArray())
	return str
}
