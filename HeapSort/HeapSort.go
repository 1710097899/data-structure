package HeapSort

import "data_structure/MaxHeap"

func Sort(data []int) {
	maxHeap := MaxHeap.NewMaxHeap(1)
	for _, v := range data {
		maxHeap.Add(v)
	}
	for i := len(data) - 1; i >= 0; i-- {
		data[i] = maxHeap.ExtractMax()
	}
}
