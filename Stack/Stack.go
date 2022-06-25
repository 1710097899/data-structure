package Stack

type Stack interface {
	GetSize() int       //一共多少元素
	IsEmpty() bool      //栈是否为空
	Push(e interface{}) //入栈
	Pop() interface{}   //出栈
	Peek() interface{}  //栈顶的值
	String() string
}
