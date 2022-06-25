package main

import (
	"data_structure/Stack/Solution"
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(Solution.IsValid("[]{}()"))
	fmt.Println(Solution.IsValid("[]}()"))
}
