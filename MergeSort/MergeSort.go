package MergeSort

import "math"

//自顶向下实现归并排序
func Sort(arr []int) {
	var tamp = make([]int, len(arr))
	copy(tamp, arr)
	sort(arr, 0, len(arr)-1, tamp)
}

//自底向上实现归并排序
func SortBU(arr []int) {
	var tamp = make([]int, len(arr))
	copy(tamp, arr)

	var lenth int = len(arr)
	//遍历合并的区间长度
	for size := 1; size <= lenth; size += size {
		//遍历合并两个区间的初始位置
		//合并[i,i+size-1]和[i+size,i+2*size-1],arr l mid r	如果i+2*size>n了，越界了，就取n-1
		for i := 0; i+size < lenth; i += 2 * size {
			if i+size-1 > i+size {
				merge(arr, i, i+size-1, int(math.Min(float64(i+2*size-1), float64(lenth-1))), tamp)
			}

		}
	}
}

func sort(arr []int, l int, r int, tamp []int) {
	/*	当数组个数比较小的时候，使用插入排序更优
		if r-l <= 15 {
		insertionSort.Sort(arr, l, r)
		return
	}*/
	if l >= r {
		return
	}
	//防止长度越界
	mid := l + (r-l)/2
	sort(arr, l, mid, tamp)
	sort(arr, mid+1, r, tamp)
	//加上如下判断
	//优化：当数组本身为有序时(即前半数组的最后一个元素arr[mid]<后半数组第一个元素arr[mid+1])，这里的时间复杂度将会变成O(n)
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r, tamp)
	}
}

//合并两个有序的区间arr[l,mid],arr[mid+1,r]
func merge(arr []int, l int, mid int, r int, tamp []int) {
	copy(tamp[l:r+1], arr[l:r+1])
	i := l
	j := mid + 1
	//每轮循环为arr[k]赋值
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = tamp[j]
			j++
		} else if j > r {
			arr[k] = tamp[i]
			i++
		} else if tamp[i] < tamp[j] {
			arr[k] = tamp[i]
			i++
		} else {
			arr[k] = tamp[j]
			j++
		}
	}
}
