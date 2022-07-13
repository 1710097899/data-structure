package SegmentTree

import "fmt"

type SegmentTree struct {
	data   []int
	tree   []int
	merger mergerFunc
}

type mergerFunc func(a, b int) int

//构造函数
func NewSegmentTree(arr []int, merger mergerFunc) *SegmentTree {
	length := len(arr)
	data := make([]int, 4*length)
	// 将线段树中中需要存储的数据，存储到data中
	for i := 0; i < len(arr); i++ {
		data[i] = arr[i]
	}
	tree := &SegmentTree{
		data:   arr,
		tree:   data,
		merger: merger,
	}
	tree.BuildSegmentTree(0, 0, length-1)
	return tree
}

//在treeindex的位置创建【l....r】的线段树
func (s *SegmentTree) BuildSegmentTree(treeindex int, l, r int) int {
	if l == r {
		s.tree[treeindex] = s.data[l]
		return s.data[l]
	}
	leftTreeIndex := s.leftChild(treeindex)
	rigthTreeIndex := s.rightChild(treeindex)
	mid := l + (r-l)/2
	//在leftTreeIndex的位置创建【l，mid】的线段树，在rigthTreeIndex的位置创建【mid+1，r】的线段树
	s.BuildSegmentTree(leftTreeIndex, l, mid)
	s.BuildSegmentTree(rigthTreeIndex, mid+1, r)

	//最后根据业务条件加上tree[]的条件，例如：求和，tree[treeindex]=tree[leftTreeIndex]+tree[rigthTreeIndex]
	s.tree[treeindex] = s.merger(s.tree[leftTreeIndex], s.tree[rigthTreeIndex])

	return s.tree[treeindex]
}

func (s *SegmentTree) getsize() int {
	return len(s.data)
}

func (s *SegmentTree) get(index int) int {
	if index < 0 || index >= len(s.data) {
		fmt.Println("Index is illegal")
	}
	return s.data[index]
}

//返回完全二叉树的数组表示中，一个索引表示的元素的左孩子的节点的索引
func (s *SegmentTree) leftChild(index int) int {
	return 2*index + 1
}

//返回完全二叉树的数组表示中，一个索引表示的元素的右孩子的节点的索引
func (s *SegmentTree) rightChild(index int) int {
	return 2*index + 2
}

//返回[queryL, queryR]区间所存储的值
func (s *SegmentTree) Query(queryL, queryR int) int {

	return s.query(0, 0, len(s.data)-1, queryL, queryR)
}

//在以treeindex为根的线段树中[l,r]的范围里，搜索[queryL, queryR]的值
func (s *SegmentTree) query(treeindex, l, r int, queryL, queryR int) int {
	if queryL < 0 || queryL >= len(s.data) || queryR < 0 || queryR >= len(s.data) || queryL > queryR {
		fmt.Println("Index is illegal")
	}
	if l == queryL && r == queryR {
		return s.tree[treeindex]
	}
	mid := l + (r-l)/2
	leftTreeIndex := s.leftChild(treeindex)
	rigthTreeIndex := s.rightChild(treeindex)

	if queryL >= mid+1 {
		return s.query(rigthTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return s.query(leftTreeIndex, l, mid, queryL, queryR)
	}
	leftres := s.query(rigthTreeIndex, mid+1, r, mid+1, queryR)
	rightres := s.query(leftTreeIndex, l, mid, queryL, mid)
	return s.merger(leftres, rightres)
}

//将index位置的值换成e
func (s *SegmentTree) Set(index int, e int) {
	if index < 0 || index >= len(s.data) {
		fmt.Println("index is illegel")
	}
	s.data[index] = e
	s.set(0, 0, len(s.data)-1, index, e)
}

func (s *SegmentTree) set(treeindex, l, r int, index, e int) {
	if l == r {
		s.tree[treeindex] = e
		return
	}
	mid := l + (r-l)/2
	leftTreeIndex := s.leftChild(treeindex)
	rigthTreeIndex := s.rightChild(treeindex)
	if index >= mid+1 {
		s.set(rigthTreeIndex, mid+1, r, index, e)
	} else {
		s.set(leftTreeIndex, l, mid, index, e)
	}
	s.tree[treeindex] = s.merger(s.tree[leftTreeIndex], s.tree[rigthTreeIndex])
}
