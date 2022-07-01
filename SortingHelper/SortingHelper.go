package SortingHelper

import (
	insertionSort "data_structure/InsertionSort"
	"data_structure/MergeSort"
	"data_structure/SelectionSort"
	"errors"
	"fmt"
	"testing"
	"time"
)

func IsSorted(arr []int) error {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr[i], "----", arr[i+1])
			fmt.Println(i, "----", i+1)
			return errors.New("this sort is error ")
		}
	}
	return nil
}

func SortTest(sortName string, arr []int, t *testing.T) {
	start := time.Now()
	if sortName == "SelectionSort" {
		SelectionSort.SelectionSort(arr)
	} else if sortName == "InsertionSort" {
		insertionSort.InsertionSort(arr)
	} else if sortName == "InsertionSort2" {
		insertionSort.InsertionSort2(arr)
	} else if sortName == "MergeSort" {
		MergeSort.Sort(arr)
	} else if sortName == "MergeSortBU" {
		MergeSort.SortBU(arr)
	}
	elapsed := time.Since(start)
	err := IsSorted(arr)
	if err != nil {
		t.Error(sortName, " ", err)
	}
	fmt.Println(sortName+" n=", len(arr), " time of use：", elapsed)
}
