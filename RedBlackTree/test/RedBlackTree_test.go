package test

import (
	"data_structure/RedBlackTree"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRedBlackTree(t *testing.T) {
	const count = 10
	redBlackTree := RedBlackTree.NewRedBlackTree()
	nums := []int{}
	for i := 0; i < count; i++ {
		nums = append(nums, rand.Intn(count))
	}
	fmt.Println("数据源:", nums)
	ti := time.Now()
	for _, v := range nums {
		redBlackTree.Add(v, v)
	}
	fmt.Println("RedBlackTree:", ti.Sub(time.Now()))
	redBlackTree.PrintPreOrder()
	fmt.Println("节点数量:", redBlackTree.GetSize())
}
