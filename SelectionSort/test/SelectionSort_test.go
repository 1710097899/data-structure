package test

import (
	"data_structure/SelectionSort"
	"data_structure/SortingHelper"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	datasize := []int{10000, 100000}
	for _, v := range datasize {
		arr := SelectionSort.GenerateOrderedArray(v, v)
		SortingHelper.SortTest("SelectionSort", arr, t)
	}
}
