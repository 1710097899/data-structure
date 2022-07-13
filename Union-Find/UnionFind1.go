package Union_Find

import "errors"

/*第一版union find
特点：union（p，q）操作时，将set[p] = set[q]，之后遍历一遍set，
找出所有set[i] == set[p]的索引，执行 set[i] = set[q]*/

// 实现
type unionSet1 struct {
	set []int
}

func NewUnionSet1(size int) UF {
	buf := make([]int, size)
	for i := 0; i < size; i++ {
		buf[i] = i // 初始时，所有节点均指向自己
	}
	return &unionSet1{set: buf}
}

func (set *unionSet1) GetSize1() int {
	return len(set.set)
}

func (set *unionSet1) GetID(p int) (int, error) {
	if p < 0 || p > len(set.set) {
		return 0, errors.New(
			"failed to get ID,index is illegal.")
	}
	return set.set[p], nil
}

func (set *unionSet1) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
		return false, errors.New(
			"failed to get ID,index is illegal.")
	}
	return set.set[p] == set.set[q], nil
}

func (set *unionSet1) Union(p, q int) error {
	b, err := set.IsConnected(p, q)
	if err != nil {
		return err
	}

	if !b {
		pID := set.set[p]
		qID := set.set[q]
		for k, v := range set.set {
			if v == pID {
				set.set[k] = qID
			}
		}
	}
	return nil
}

/*时间复杂度分析：
union()          O(n)
isConnected()    O(1)
*/
