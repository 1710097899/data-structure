package Union_Find

import "errors"

/*第五版union find
路径压缩*/
// 实现
type unionSet5 struct {
	rank []int // rank[i]表示以i为根的树的高度
	set  []int
}

func NewUnionSet5(size int) UF {
	buf1 := make([]int, size)
	for i := 0; i < size; i++ {
		buf1[i] = i
	}
	buf2 := make([]int, size)
	for i := 0; i < size; i++ {
		buf2[i] = 1
	}

	return &unionSet5{
		rank: buf2,
		set:  buf1,
	}
}

func (set *unionSet5) GetSize() int {
	return len(set.set)
}

func (set *unionSet5) GetID(p int) (int, error) {
	if p < 0 || p > len(set.set) {
		return 0, errors.New(
			"failed to get ID,index is illegal.")
	}

	return set.getRoot5(p), nil
}

func (set *unionSet5) getRoot5(p int) int {
	for p != set.set[p] {
		set.set[p] = set.set[set.set[p]] //路径压缩
		p = set.set[p]
	}
	return p
}

func (set *unionSet5) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return false, errors.New(
			"error: index is illegal.")
	}
	return set.getRoot5(set.set[p]) == set.getRoot5(set.set[q]), nil
}

func (set *unionSet5) Union(p, q int) error {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return errors.New(
			"error: index is illegal.")
	}

	pRoot := set.getRoot5(p)
	qRoot := set.getRoot5(q)

	if pRoot != qRoot {
		if set.rank[pRoot] < set.rank[qRoot] {
			set.set[pRoot] = qRoot
		} else if set.rank[qRoot] < set.rank[pRoot] {
			set.set[qRoot] = pRoot
		} else { //到这步说明高度相等,谁指向谁都可以,高度要+1
			set.set[pRoot] = qRoot
			set.rank[qRoot] += 1
		}
	}
	return nil
}

/*时间复杂度分析： // h为树的高度
union()          O(h)
isConnected()    O(h)
从理论上讲，是比第三版更好的优化方式*/
