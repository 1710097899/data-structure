package test

import (
	"data_structure/Array"
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := Array.NewArray(10)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
	arr.Add(1, 100)
	fmt.Println(arr)
	arr.AddFirst(-1)
	fmt.Println(arr)
	res := arr.GetIndex(5)
	fmt.Println(res)
	arr.SetIndex(0, "9")
	fmt.Println(arr)
	res1 := arr.Contians("9")
	fmt.Println(res1)
	res2 := arr.Find(9)
	fmt.Println(res2)
	res3 := arr.Delect(2)
	fmt.Println(arr, res3)
	res4 := arr.DelectFirst()
	fmt.Println(arr, res4)
	res5 := arr.DelectLast()
	fmt.Println(arr, res5)
	res6 := arr.DelectE(7)
	fmt.Println(arr, res6)

}
