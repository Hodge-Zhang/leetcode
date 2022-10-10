package main

import (
	"fmt"
	"testing"
)

func TestMinSwap(t *testing.T) {
	fmt.Println(MinSwap([]int{1, 3, 3, 4}, []int{1, 2, 5, 7}))
	fmt.Println(MinSwap([]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}))
}
