package SelectionSort

import (
	"math/rand"
	"time"
)

func GenerateOrderedArray(n int, m int) []int {
	arr := make([]int, 0, 0)
	//加入随机数生产数组
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(m))
	}
	return arr
}
