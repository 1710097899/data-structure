package test

import (
	"data_structure/SelectionSort"
	"data_structure/SortingHelper"
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	n := 10000
	arr := SelectionSort.GenerateRandomArray(n, n)
	arr3 := SelectionSort.GenerateRandomArray(n, n)
	SortingHelper.SortTest("ShellSort", arr, t)
	SortingHelper.SortTest("MergeSort", arr3, t)
	fmt.Println()
	arr = SelectionSort.GenerateOrderedArray(n)
	arr3 = SelectionSort.GenerateOrderedArray(n)
	SortingHelper.SortTest("ShellSort", arr, t)
	SortingHelper.SortTest("MergeSort", arr3, t)
	fmt.Println()
	n = 10000
	arr = SelectionSort.GenerateRandomArray(n, 1)
	arr3 = SelectionSort.GenerateRandomArray(n, 1)
	SortingHelper.SortTest("ShellSort", arr, t)
	SortingHelper.SortTest("MergeSort", arr3, t)

}
