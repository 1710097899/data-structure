package test

import (
	"data_structure/SelectionSort"
	"data_structure/SortingHelper"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 10000

	arr1 := SelectionSort.GenerateRandomArray(n, n)
	var arr2 = make([]int, len(arr1))
	var arr3 = make([]int, len(arr1))
	var arr4 = make([]int, len(arr1))
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	SortingHelper.SortTest("QuickSort", arr1, t)
	SortingHelper.SortTest("QuickSort2", arr2, t)
	SortingHelper.SortTest("QuickSort3", arr3, t)
	SortingHelper.SortTest("QuickSort4", arr4, t)

}
