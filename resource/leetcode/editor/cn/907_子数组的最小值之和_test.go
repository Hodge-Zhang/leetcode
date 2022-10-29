package main

import (
	"fmt"
	"testing"
)

func TestSumSubarrayMins(t *testing.T) {
	fmt.Println(SumSubarrayMins([]int{3, 1, 2, 4}))
	fmt.Println(SumSubarrayMins([]int{11, 81, 94, 43, 3}))
}
