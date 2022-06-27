package test

import (
	"data_structure/LinkedList/Solution"
	"fmt"
	"testing"
)

func TestSolutiuon(t *testing.T) {
	LinkedList := Solution.NewLinkedList()
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	LinkedList.NewListNode(nums)
	fmt.Println(LinkedList)
	res := Solution.RemoveElements(LinkedList, 6)
	fmt.Println(res)
}
