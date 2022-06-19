package main

import "fmt"

//线性查找
func bfprt(a interface{}, arr []interface{}) interface{} {
	for i, k := range arr {
		if k == a {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 6, 5, 9, 3, 3}
	rst := bfprt(3, arr)
	fmt.Println(rst)
}
