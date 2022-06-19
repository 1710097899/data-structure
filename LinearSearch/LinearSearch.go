package LinearSrearch

import "reflect"

//线性查找
func Bfprt(a interface{}, arr []interface{}) int {
	for i := 0; i < len(arr); i++ {
		if reflect.DeepEqual(a, arr[i]) {
			return i
		}
	}
	return -1
}
