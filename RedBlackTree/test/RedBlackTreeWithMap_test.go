package test

import (
	"data_structure/RedBlackTree"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRedBlackTreeWithMap(t *testing.T) {
	//红黑树和go自带map性能对比
	const count = 1000000
	redBlackTree := RedBlackTree.NewRedBlackTree()
	nums := []int{}
	for i := 0; i < count; i++ {
		nums = append(nums, rand.Intn(count))
	}

	ti := time.Now()
	for _, v := range nums {
		redBlackTree.Add(v, v)
	}
	fmt.Println("redBlackTree:", ti.Sub(time.Now()))
	m := map[int]int{}
	ti = time.Now()
	for _, v := range nums {
		m[v] = v
	}
	fmt.Println("map:", ti.Sub(time.Now()))

	/*
		map在速度上还是很大优势，劣势就是无法顺序遍历。
		一种数据结构在某方面性能更优，往往是因为他多维护或少维护了某些东西，
		如哈希表和红黑树对比，哈希表牺牲了顺序，赢得了速度，而红黑树赢了顺序，输了速度。
		不存在说一定要使用某种数据结构，要具体问题具体分析。*/
}
