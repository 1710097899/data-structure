package Stack

import (
	"data_structure/Array"
)

type arraystack struct {
	array Array.IArray
}

func ArrayStack(cap int) Stack {
	return &arraystack{
		array: Array.NewArray(10),
	}
}
