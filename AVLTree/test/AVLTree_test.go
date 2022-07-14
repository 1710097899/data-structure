package test

import (
	avltree "data_structure/AVLTree"
	"fmt"
	"math/rand"
	"testing"
)

//测试代码

func TestAVLTree(t *testing.T) {
	a := avltree.NewAvlTree()
	nums := []int{}
	a.PrintInOrder()
	fmt.Println("----")
	for i := 0; i < 100000; i++ {
		a.Add(100000, 0)
		if !a.IsBST() || !a.IsBalanced() {
			fmt.Println("代码有错误", a.IsBST(), a.IsBalanced())
			break
		}
	}

	for i := 0; i < 200000; i++ {
		nums = append(nums, rand.Intn(200000))
	}
	for _, v := range nums {
		a.Remove(v)
		if !a.IsBST() || !a.IsBalanced() {
			fmt.Println("代码有错误", a.IsBST(), a.IsBalanced())
			break
		}
	}

	fmt.Println("测试结束")
	fmt.Println("is BST:", a.IsBST())
	fmt.Println("is Balanced:", a.IsBalanced())
}
