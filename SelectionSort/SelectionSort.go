package SelectionSort

func swap(arr []int, i int, minIndex int) {
	arr[i], arr[minIndex] = arr[minIndex], arr[i]
}

func SelectionSort(arr []int) {
	//使arr[0,i)有序，arr[i,n)无序
	for i := 0; i < len(arr); i++ {
		//选择arr[]中最小的索引
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		swap(arr, i, min)
	}
}
