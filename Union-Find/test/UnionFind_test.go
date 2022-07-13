package test

import (
	Union_Find "data_structure/Union-Find"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const count = 10000

func TestUnionFind(t *testing.T) {
	nums := []int{}
	for i := 0; i < 50*count; i++ {
		nums = append(nums, rand.Intn(count))
	}
	test1(nums)
	test2(nums)
	test3(nums)
	test4(nums)
	test5(nums)
}

func test1(nums []int) {
	s := Union_Find.NewUnionSet1(count)
	t := time.Now()
	for i := 1; i < 50*count; i++ {
		s.GetID(nums[i])
		s.Union(nums[i-1], nums[i])
	}
	fmt.Println("第一版时间: ", t.Sub(time.Now()))
}
func test2(nums []int) {
	s := Union_Find.NewUnionSet2(count)
	t := time.Now()
	for i := 1; i < 50*count; i++ {
		s.GetID(nums[i])
		s.Union(nums[i-1], nums[i])
	}
	fmt.Println("第二版时间: ", t.Sub(time.Now()))
}
func test3(nums []int) {
	s := Union_Find.NewUnionSet3(count)
	t := time.Now()
	for i := 1; i < 50*count; i++ {
		s.GetID(nums[i])
		s.Union(nums[i-1], nums[i])
	}
	fmt.Println("第三版时间: ", t.Sub(time.Now()))
}

func test4(nums []int) {
	s := Union_Find.NewUnionSet4(count)
	t := time.Now()
	for i := 1; i < 50*count; i++ {
		s.GetID(nums[i])
		s.Union(nums[i-1], nums[i])
	}
	fmt.Println("第四版时间: ", t.Sub(time.Now()))
}
func test5(nums []int) {
	s := Union_Find.NewUnionSet5(count)
	t := time.Now()
	for i := 1; i < 50*count; i++ {
		s.GetID(nums[i])
		s.Union(nums[i-1], nums[i])
	}
	fmt.Println("第五版时间: ", t.Sub(time.Now()))
}

// 测试结果
/*第一版时间:  -5.101884ms
第二版时间:  -423.963439ms
第三版时间:  -1.912094ms
第四版时间:  -1.788539ms
分析：
1）加入数量优化或者高度优化的并查集性能明显优于前两版，后续要加大数量级，不再对前两版测试；
2）第二版时间更慢是因为在go的底层中，访问连续的内存地址速度要快于通过索引访问，而第二版
寻找根节点是通过索引访问的
3）第三和第四版的优化明显抵消了通过索引访问慢的劣势；
当数量级增加到50万后，执行时间是：
第一版时间:  -216.2098ms
第二版时间:  -26.6249176s
第三版时间:  -2.0795ms
第四版时间:  -2.7512ms
两者性能差异非常小，不过第四版理论上能避免第三版最坏情况的出现，实际应用中使用第四版居多，
还能优化，加入路径压缩功能*/
