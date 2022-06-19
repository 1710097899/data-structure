package LinearSrearch

//自动生成数组构造器
func GenerateOrderedArray(n int) []int {
	arr := make([]int, n, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return arr
}
