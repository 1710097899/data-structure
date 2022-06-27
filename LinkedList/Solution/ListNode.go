package Solution

type IListNode interface {
	NewListNode(arr []int)
}

//ListNode的构造函数
func NewNodeIsNil(e int) *ListNode {
	return &ListNode{
		Val:  e,
		Next: nil,
	}
}

func NewNode(e int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  e,
		Next: next,
	}
}
func NewLinkedList() *ListNode {
	return &ListNode{}
}

func (l *ListNode) NewListNode(arr []int) {
	if arr == nil || len(arr) == 0 {
		panic("arr can not be empty")
	}
	l.Val = arr[0]
	cur := l
	for i := 1; i < len(arr); i++ {
		cur.Next = NewNodeIsNil(arr[i])
		cur = cur.Next
	}
}
