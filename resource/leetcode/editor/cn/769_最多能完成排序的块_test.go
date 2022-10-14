package main

import (
	"fmt"
	"testing"
)

func TestMaxChunksToSorted(t *testing.T) {
	fmt.Println(MaxChunksToSorted([]int{4,3,2,1,0}))
	fmt.Println(MaxChunksToSorted([]int{0,1,2,3,4}))
	fmt.Println(MaxChunksToSorted([]int{0,1,3,4,2}))
	fmt.Println(MaxChunksToSorted([]int{2,3,1,3,0}))
	fmt.Println(MaxChunksToSorted([]int{4,0,2,1,3}))
}