package Union_Find

import "errors"

/*第四版union find
特点：加入rank数组，rank[ i ] 表示集合高度，每次执行union(p,q)，取出p，q所在集合的高度，
让高度低的指向高度高的。*/
// 实现
type unionSet4 struct {
	rank []int // rank[i]表示以i为根的树的高度
	set  []int
}

func NewUnionSet4(size int) UF {
	buf1 := make([]int, size)
	for i := 0; i < size; i++ {
		buf1[i] = i
	}
	buf2 := make([]int, size)
	for i := 0; i < size; i++ {
		buf2[i] = 1
	}

	return &unionSet4{
		rank: buf2,
		set:  buf1,
	}
}

func (set *unionSet4) GetSize4() int {
	return len(set.set)
}

func (set *unionSet4) GetID(p int) (int, error) {
	if p < 0 || p > len(set.set) {
		return 0, errors.New(
			"failed to get ID,index is illegal.")
	}

	return set.getRoot4(p), nil
}

func (set *unionSet4) getRoot4(p int) int {
	for p != set.set[p] {
		p = set.set[p]
	}
	return p
}

func (set *unionSet4) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return false, errors.New(
			"error: index is illegal.")
	}
	return set.getRoot4(set.set[p]) == set.getRoot4(set.set[q]), nil
}

func (set *unionSet4) Union(p, q int) error {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return errors.New(
			"error: index is illegal.")
	}

	pRoot := set.getRoot4(p)
	qRoot := set.getRoot4(q)

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
