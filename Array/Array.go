package Array

import "fmt"

type IArray interface {
	GetSize() int
	GetCap() int
	IsEmpty() bool
	AddLast(e interface{})
	Add(index int, e interface{})
	AddFirst(e interface{})
	String() string
	GetIndex(index int) interface{}
	SetIndex(index int, e interface{})
	Contians(e interface{}) bool
	Find(e interface{}) int
	Delect(index int) interface{}
	DelectFirst() interface{}
	DelectLast() interface{}
	DelectE(e interface{}) bool
	GetLast() interface{}
	GetFrist() interface{}
	GetArray() []interface{}
}

type Array struct {
	data []interface{}
	size int
}

//构造函数，传入数组的容量cap构造Array
func NewArray(cap int) IArray {
	return &Array{
		data: make([]interface{}, cap),
		size: 0,
	}
}

func (a *Array) GetArray() []interface{} {
	return a.data[:a.size]
}

//获取数组中元素的个数
func (a *Array) GetSize() int {
	return a.size
}

//获取数组的容量
func (a *Array) GetCap() int {
	return len(a.data)
}

//返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//向数组的最后添加元素
func (a *Array) AddLast(e interface{}) {
	/*a.data[a.size] = e
	a.size++*/
	a.Add(a.size, e) //等于上两行注释代码
}

//向数组的开头添加元素
func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

//向数组中指定位置添加元素
func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("Add failed.Required index>=0and<size")
	}
	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}
func (a *Array) String() string {
	var str string
	str = fmt.Sprintf("Array : size = %d,capacity = %d\n", a.size, len(a.data))
	str = fmt.Sprintf("%s%v\n", str, a.data[0:a.size])
	return str
}

//取出索引为index的元素
func (a *Array) GetIndex(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal")
	}
	return a.data[index]
}

//取出数组中最后一个元素
func (a *Array) GetLast() interface{} {
	return a.GetIndex(a.size - 1)
}

//取出数组中第一个元素
func (a *Array) GetFrist() interface{} {
	return a.GetIndex(0)
}

//修改索引为index的元素
func (a *Array) SetIndex(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed. Index is illegal")
	}
	a.data[index] = e
}

//查找数组中是否存在元素e
func (a *Array) Contians(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}

//寻找如果存在的元素e所对应的索引
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

//从数组中删除索引为index的元素，返回删除元素
func (a *Array) Delect(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Delect failed. Index is illegal")
	}
	res := a.data[index]
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}
	return res
}

//删除数组中第一个元素
func (a *Array) DelectFirst() interface{} {
	return a.Delect(0)
}

//删除数组中最后一个元素
func (a *Array) DelectLast() interface{} {
	return a.Delect(a.size - 1)
}

//查看数组中是否有元素e，如果有就删除元素e（第一个e）
func (a *Array) DelectE(e interface{}) bool {
	/*for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			a.data[i] = a.data[i+1]
		}
	}
	a.size--*/
	find := a.Find(e)
	if find != -1 {
		a.Delect(find)
		return true
	}
	return false
}

//实现动态数组使数组的容量变得可以动态伸缩
func (a *Array) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.data[i]
	}
	a.data = newArr
}
