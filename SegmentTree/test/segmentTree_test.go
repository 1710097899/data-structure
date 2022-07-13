package test

import (
	"data_structure/SegmentTree"
	"fmt"
	"testing"
)

//测试线段树
func multiplication(v1, v2 int) int {
	return v1 + v2
}

func TestSegmentTree(t *testing.T) {
	data := []int{-2, 0, 3, -5, 2, -1}
	segment := SegmentTree.NewSegmentTree(data, multiplication)
	fmt.Println(segment)
	queryres := segment.Query(0, 2)
	fmt.Println(queryres)
	segment.Set(0, 1)
	fmt.Println(segment)
}
