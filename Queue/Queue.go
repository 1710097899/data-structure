package Queue

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(e interface{}) //入队
	Dequeue() interface{}  //出队
	GetFront() interface{} //获取队首的元素
	Getcap() int
}
