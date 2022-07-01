package main

import (
	"data_structure/SelectionSort"
	"data_structure/SortingHelper"
	"testing"
)

func TestMergeSort(t *testing.T) {
	n := 100000
	array := SelectionSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("MergeSort", array, t)
	array2 := SelectionSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("SelectionSort", array2, t)
	array3 := SelectionSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("InsertionSort", array3, t)
}
