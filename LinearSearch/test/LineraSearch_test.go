package test

import (
	"data_structure/LinearSearch"
	"fmt"
	"testing"
	"time"
)

func TestLinearGeneratorSearch(t *testing.T) {
	dataSize := []int{100000, 1000000}
	for _, n := range dataSize {
		tamp := LinearSrearch.GenerateOrderedArray(n)
		data := make([]interface{}, len(tamp))
		for i, v := range tamp {
			data[i] = v
		}
		start := time.Now()
		for k := 0; k < 10; k++ {
			LinearSrearch.Bfprt(n, data)
		}
		elapsed := time.Since(start)
		fmt.Println("n=", n, "10 runs LinearSearch  time of use：", elapsed)
	}
}
