package SelectionSort

import (
	"math/rand"
	"time"
)

//数组生成器
//生成有序数组
func GenerateOrderedArray(n int) []int {
	arr := make([]int, n, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return arr
}

//生成随机数组
func GenerateRandomArray(n int, bound int) []int {
	arr := make([]int, 0, 0)
	//加入随机数生产数组
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		//生成一个长度为n（i）的数组，数字范围在[0,bound)的数组
		arr = append(arr, rand.Intn(bound))
	}
	return arr
}
