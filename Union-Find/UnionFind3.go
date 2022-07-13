package Union_Find

import "errors"

/*第三版union find
特点：加入numNode数组，numNode[ i ] 表示集合 i 中节点数量，每次执行union时，
取出numNode的节点数量，让节点数量小的指向节点数量大的。*/

// 实现
type unionSet3 struct {
	numNode []int // numNode[i]表示集合i中是节点数量
	set     []int
}

func NewUnionSet3(size int) UF {
	buf1 := make([]int, size)
	for i := 0; i < size; i++ {
		buf1[i] = i
	}
	buf2 := make([]int, size)
	for i := 0; i < size; i++ {
		buf2[i] = 1 // 初始节点数量均为1
	}

	return &unionSet3{
		numNode: buf2,
		set:     buf1,
	}
}

func (set *unionSet3) GetSize3() int {
	return len(set.set)
}

func (set *unionSet3) GetID(p int) (int, error) {
	if p < 0 || p > len(set.set) {
		return 0, errors.New(
			"failed to get ID,index is illegal.")
	}
	return set.getRoot3(p), nil
}

func (set *unionSet3) getRoot3(p int) int {
	for p != set.set[p] {
		p = set.set[p]
	}
	return p
}

func (set *unionSet3) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return false, errors.New(
			"error: index is illegal.")
	}
	return set.getRoot3(set.set[p]) == set.getRoot3(set.set[q]), nil
}

func (set *unionSet3) Union(p, q int) error {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return errors.New(
			"error: index is illegal.")
	}

	pRoot := set.getRoot3(p)
	qRoot := set.getRoot3(q)

	if pRoot != qRoot {
		if set.numNode[pRoot] < set.numNode[qRoot] {
			set.set[pRoot] = qRoot
			set.numNode[qRoot] += set.numNode[pRoot]
		} else {
			set.set[qRoot] = pRoot
			set.numNode[pRoot] += set.numNode[qRoot]
		}
	}
	return nil
}

/*时间复杂度分析： // h为树的高度
union()          O(h)
isConnected()    O(h)
相比第二版，这版能有效避免树变为单链的情况，不过还能继续优化，*/
