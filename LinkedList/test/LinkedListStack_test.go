package test

import (
	"data_structure/LinkedList"
	"fmt"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	linkedListStack := LinkedList.NewLinkedListStarck()
	for i := 0; i < 5; i++ {
		linkedListStack.Push(i)
		fmt.Println(linkedListStack)
	}
	linkedListStack.Pop()
	fmt.Println(linkedListStack)
}
