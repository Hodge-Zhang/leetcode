package main

import (
	"fmt"
	"testing"
)

func TestShortestSubarray(t *testing.T) {
	fmt.Println(ShortestSubarray([]int{56, -21, 56, 35, -9}, 61))
	fmt.Println(ShortestSubarray([]int{1, 2}, 4))
	fmt.Println(ShortestSubarray([]int{2, -1, 2}, 3))
}
