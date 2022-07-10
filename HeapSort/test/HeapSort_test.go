package test

import (
	"data_structure/SelectionSort"
	"data_structure/SortingHelper"
	"testing"
)

func TestHeapSort(t *testing.T) {
	n := 1000000
	arr := SelectionSort.GenerateRandomArray(n, n)
	arr2 := SelectionSort.GenerateRandomArray(n, n)
	arr3 := SelectionSort.GenerateRandomArray(n, n)
	arr4 := SelectionSort.GenerateRandomArray(n, n)

	SortingHelper.SortTest("MergeSort", arr, t)
	SortingHelper.SortTest("QuickSort3", arr2, t)
	SortingHelper.SortTest("QuickSort4", arr3, t)
	SortingHelper.SortTest("HeapSort", arr4, t)

}
