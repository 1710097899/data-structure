package test

import (
	. "data_structure/InsertionSort"
	"data_structure/SortingHelper"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	datasize := []int{10000, 100000}
	for _, v := range datasize {
		arr := GenerateOrderedArray(v, v)
		//确保优化前后的算法所测试的数组一致
		arr2 := arr
		SortingHelper.SortTest("InsertionSort", arr, t)
		SortingHelper.SortTest("InsertionSort2", arr2, t)
	}
}
