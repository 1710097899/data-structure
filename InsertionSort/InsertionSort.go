package insertionSort

func swap(arr []int, i int, minIndex int) {
	arr[i], arr[minIndex] = arr[minIndex], arr[i]
}

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j-1 >= 0; j-- {
			//如果arr[j]比arr[j-1]小，就交换位置
			if arr[j-1] > arr[j] {
				swap(arr, j-1, j)
			}
		}
	}
}

//插入排序算法的优化
func InsertionSort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		//将arr[i]的值先进行暂存
		k := arr[i]
		for j := i; j-1 >= 0; j-- {
			//将比arr[i]大的数先向后进行平移，当arr[i]的值比arr[j-1]的值大的时候
			//就说明arr[j]当前的位置就是k(arr[i])应该待的位置
			if arr[j-1] > k {
				arr[j] = arr[j-1]
			}
			arr[j] = k
		}
	}
}

//插入排序协助优化归并排序
func Sort(arr []int, l int, r int) {
	for i := l; i <= r; i++ {
		k := arr[i]
		for j := i; j-1 > l; j-- {
			if arr[j-1] > k {
				arr[j] = arr[j-1]
			}
			arr[j] = k
		}
	}
}
