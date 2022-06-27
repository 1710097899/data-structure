package Solution

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = RemoveElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

func (l *ListNode) String() string {
	var str string
	cur := l
	for cur != nil {
		val := cur.Val
		str += fmt.Sprintf("%#v", val) + "->"
		cur = cur.Next
	}
	str += "nil"
	return str
}
