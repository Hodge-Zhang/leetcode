package main

import (
	"fmt"
	"testing"
)

func TestThreeEqualParts(t *testing.T) {
	fmt.Println(ThreeEqualParts([]int{0, 0, 0}))
	fmt.Println(ThreeEqualParts([]int{1, 0, 1, 0, 1}))
	fmt.Println(ThreeEqualParts([]int{1, 1, 0, 1, 0}))
	fmt.Println(ThreeEqualParts([]int{0, 1, 0, 0, 1, 0, 1, 0}))
}
