package test

import (
	"data_structure/Trie"
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Trie.NewTrie()
	fmt.Println(trie)
	trie.Add2("dasd")
	fmt.Println(trie)
	res := trie.Contains("dasd")
	fmt.Println(res)
	res1 := trie.IsPrefix("da")
	fmt.Println(res1)
}
