package LinkedList

import "fmt"

type ILinkedListStack interface {
	GetSize() int
	Empty() bool
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
}

type LinkedListStack struct {
	linkedList ILinkedList
}

//构造函数
func NewLinkedListStarck() ILinkedListStack {
	return &LinkedListStack{
		linkedList: NewLinkedList(),
	}
}

func (l *LinkedListStack) GetSize() int {
	return l.linkedList.Getsize()
}

func (l *LinkedListStack) Empty() bool {
	return l.linkedList.IsEmpty()
}
func (l *LinkedListStack) Push(e interface{}) {
	l.linkedList.AddFirst(e)
}
func (l *LinkedListStack) Pop() interface{} {
	return l.linkedList.RemoveFirst()
}
func (l *LinkedListStack) Peek() interface{} {
	return l.linkedList.GetFirst()
}

func (l *LinkedListStack) String() string {
	str := "Stack: top " + fmt.Sprint(l.linkedList)
	return str
}
