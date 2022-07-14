package Map

import "data_structure/AVLTree"

type AVLMap struct {
	k, v interface{}
	avl  AVLTree.AVLTree
}

func NewAVlTree() *AVLTree.AVLTree {
	return new(AVLTree.AVLTree)
}

func NewAVLMap() Map {
	return new(AVLMap)
}

func (a *AVLMap) Add(key interface{}, value interface{}) {
	a.Add(key, value)
}

func (a *AVLMap) Remove(key interface{}) interface{} {
	return a.Remove(key)
}

func (a *AVLMap) Contains(key interface{}) bool {
	return a.Contains(key)
}

func (a *AVLMap) Get(Key interface{}) interface{} {
	return a.Get(Key)
}

func (a *AVLMap) Set(key interface{}, value interface{}) {
	a.Set(key, value)
}

func (a *AVLMap) GetSize() int {
	return a.GetSize()
}

func (a *AVLMap) IsEmpty() bool {
	return a.IsEmpty()
}
