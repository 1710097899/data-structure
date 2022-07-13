package Union_Find

import "errors"

/*第二版union find
特点：union（p，q）操作时，找出p的根节点n，让set[q] = n*/
// 实现
type unionSet2 struct {
	set []int
}

func NewUnionSet2(size int) UF {
	buf := make([]int, size)
	for i := 0; i < size; i++ {
		buf[i] = i
	}
	return &unionSet2{set: buf}
}

func (set *unionSet2) GetSize2() int {
	return len(set.set)
}

func (set *unionSet2) GetID(p int) (int, error) {
	if p < 0 || p > len(set.set) {
		return 0, errors.New(
			"failed to get ID,index is illegal.")
	}

	return set.getRoot2(p), nil
}

// getRoot: 找出p的根节点,时间复杂度为O(h),h为树的高度
func (set *unionSet2) getRoot2(p int) int {
	if set.set[p] == p {
		return p
	}
	return set.getRoot2(set.set[p])
}

func (set *unionSet2) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return false, errors.New(
			"error: index is illegal.")
	}
	return set.getRoot2(set.set[p]) == set.getRoot2(set.set[q]), nil
}

func (set *unionSet2) Union(p, q int) error {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return errors.New(
			"error: index is illegal.")
	}

	pRoot := set.getRoot2(p)
	qRoot := set.getRoot2(q)

	if pRoot != qRoot {
		set.set[pRoot] = qRoot
	}
	return nil
}

/*时间复杂度分析： // h为树的高度
union()          O(h)
isConnected()    O(h)
相比第一版，这个时间复杂度更好接受，不过在极端的情况下，
比如执行union(0,1),union(0,2)...union(0,n),这样树变成单链，
时间复杂度都为O(n)，要继续优化*/
