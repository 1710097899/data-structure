package QuickSort

import "math/rand"

func Sort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

//第一版快排
func sort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	sort(arr, l, p-1)
	sort(arr, p+1, r)
}

//第一版快排核心
func partition(arr []int, l, r int) int {
	//arr[l+1...j]<v,arr[j+1...i]>=v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[l] > arr[i] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func Sort2(arr []int) {
	sort2(arr, 0, len(arr)-1)
}

//第二版快排
func sort2(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition2(arr, l, r)
	sort2(arr, l, p-1)
	sort2(arr, p+1, r)
}

//第二版快排核心(防止当数组本身就是有序时，造成负效果)
func partition2(arr []int, l, r int) int {
	//生成一个[l,r]之间的随机索引
	var p int = rand.Intn(r-l+1) + l
	//将标定点元素换到第一个
	arr[l], arr[p] = arr[p], arr[l]
	//arr[l+1...j]<v,arr[j+1...i]>=v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[l] > arr[i] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func Sort3(arr []int) {
	sort3(arr, 0, len(arr)-1)
}

//第三版快排[双路快排]
func sort3(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition3(arr, l, r)
	sort3(arr, l, p-1)
	sort3(arr, p+1, r)
}

//第三版快排核心（防止数组中所有的元素都相同）[双路快排]
func partition3(arr []int, l, r int) int {
	//生成一个[l,r]之间的随机索引
	var p int = rand.Intn(r-l+1) + l
	//将标定点元素换到第一个
	arr[l], arr[p] = arr[p], arr[l]
	//arr[l+1...i-1]<=v,arr[j+1...r]>=v
	i := l + 1
	j := r
	for {
		for i <= j && arr[i] < arr[l] {
			i++
		}
		for j >= i && arr[j] > arr[l] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//三路快排
func Sort4(arr []int) {
	sort4(arr, 0, len(arr)-1)
}

//三路快排
func sort4(arr []int, l int, r int) {
	if l >= r {
		return
	}
	var p int = rand.Intn(r-l+1) + l
	arr[l], arr[p] = arr[p], arr[l]
	lt := l
	i := l + 1
	gt := r + 1
	//arr[l+1,lt]<v,arr[lt+1,i-1]==v,arr[gt,r]>v,当i==gt跳出循环
	for i < gt {
		if arr[i] < arr[l] {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > arr[l] {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else { //arr[i]==arr[l]
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	sort4(arr, l, lt-1)
	sort4(arr, gt, r)
}
