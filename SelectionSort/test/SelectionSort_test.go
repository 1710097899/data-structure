package test

import (
	"data_structure/SelectionSort"
	"fmt"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	datasize := []int{10000, 100000}
	for _, v := range datasize {
		arr := SelectionSort.GenerateOrderedArray(v, v)
		start := time.Now()
		SelectionSort.SelectionSort(arr)
		elapsed := time.Since(start)
		fmt.Println("SelectionSort", "n=", len(arr), " time of use：", elapsed)
	}
}
