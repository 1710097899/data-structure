package Map

import "fmt"

type BSTMap struct {
	root *BSTnode
	size int
}

type BSTnode struct {
	k, v  int
	Left  *BSTnode
	Right *BSTnode
}

func NewBSTnode(k, v int) *BSTnode {
	return &BSTnode{
		k:     k,
		v:     v,
		Left:  nil,
		Right: nil,
	}
}

func NewBSTMap() *BSTMap {
	return &BSTMap{
		root: nil,
		size: 0,
	}
}

func (b *BSTMap) Add(key, value int) {
	b.root = b.add(b.root, key, value)
}

func (b *BSTMap) add(Node *BSTnode, key, value int) *BSTnode {
	if Node == nil {
		b.size++
		return NewBSTnode(key, value)
	}
	if Node.v > value {
		Node.Left = b.add(Node.Left, key, value)
	} else if Node.v < value {
		Node.Right = b.add(Node.Right, key, value)
	} else {
		Node.v = value
	}
	return Node
}

//寻找二分搜索树中最小的元素
func (b *BSTMap) MiniMun() int {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return b.minMun(b.root).v
}

func (b *BSTMap) minMun(Node *BSTnode) *BSTnode {
	if Node.Left == nil {
		return Node
	}
	return b.minMun(Node.Left)

}

//寻找二分搜索树中最大的元素
func (b *BSTMap) MaxiMun() int {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return b.maxMun(b.root).v
}

func (b *BSTMap) maxMun(Node *BSTnode) *BSTnode {
	if Node.Right == nil {
		return Node
	}
	return b.maxMun(Node.Right)

}

//从二分搜索树中删除最小值所在的节点，返回最小值
func (b *BSTMap) RemoveMin() int {
	ret := b.MiniMun()
	b.root = b.removeMin(b.root)
	return ret

}

//删除以node为根的二分搜索树中的最小节点
//返回删除节点后新的二分搜索树的根
func (b *BSTMap) removeMin(Node *BSTnode) *BSTnode {
	if Node.Left == nil {
		rightNode := Node.Right
		Node.Right = nil
		b.size--
		return rightNode
	}
	Node.Left = b.removeMin(Node.Left)
	return Node
}

//从二分搜索树中删除最小值虽在的节点，返回最小值
func (b *BSTMap) RemoveMax() int {
	ret := b.MaxiMun()
	b.root = b.removeMax(b.root)
	return ret

}

//删除以node为根的二分搜索树中的最大节点
//返回删除节点后新的二分搜索树的根
func (b *BSTMap) removeMax(Node *BSTnode) *BSTnode {
	if Node.Right == nil {
		leftNode := Node.Left
		Node.Left = nil
		b.size--
		return leftNode
	}
	Node.Right = b.removeMin(Node.Right)
	return Node
}

func (b *BSTMap) Remove(key int) interface{} {
	Node := b.getBSTnode(b.root, key)
	if Node == nil {
		return nil
	} else {
		b.root = b.remove(b.root, key)
		return Node.v
	}
}

func (b *BSTMap) remove(Node *BSTnode, key int) *BSTnode {
	if Node == nil {
		return nil
	}
	if key < Node.k {
		Node.Left = b.remove(Node.Left, key)
		return Node
	} else if key > Node.k {
		Node.Right = b.remove(Node.Right, key)
		return Node
	} else {
		//待删除节点左子树为空的情况
		if Node.Left == nil {
			rightNode := Node.Right
			Node.Right = nil
			b.size--
			return rightNode
		}
		//待删除节点右子树为空的情况
		if Node.Right == nil {
			leftNode := Node.Left
			Node.Left = nil
			b.size--
			return leftNode
		}
		//待删除节点左右子树都不为空的情况
		//待着比待删除节点大的最小节点，即待删除节点右子树最小的节点（后继）
		//用这个节点顶替待删除节点的位置
		successor := b.minMun(Node.Right)
		successor.Right = b.removeMin(Node.Right)
		b.size++
		successor.Left = Node.Left
		Node = nil
		b.size--
		return successor
	}
}

func (b *BSTMap) Contains(key int) bool {
	return b.getBSTnode(b.root, key) != nil
}

func (b *BSTMap) Get(key int) interface{} {
	Node := b.getBSTnode(b.root, key)
	if Node == nil {
		return nil
	} else {
		return Node.v
	}
}

func (b *BSTMap) Set(key, value int) {
	Node := b.getBSTnode(b.root, key)
	if Node == nil {
		fmt.Println("node is not exsit")
	} else {
		Node.v = value
	}
}

func (b *BSTMap) GetSize() int {
	return b.size
}

func (b *BSTMap) IsEmpty() bool {
	return b.size == 0
}

func (b *BSTMap) getBSTnode(Node *BSTnode, key int) *BSTnode {
	if Node == nil {
		return nil
	}
	if Node.k == key {
		return Node
	} else if Node.k > key {
		return b.getBSTnode(Node.Left, key)
	} else {
		return b.getBSTnode(Node.Right, key)
	}
}
