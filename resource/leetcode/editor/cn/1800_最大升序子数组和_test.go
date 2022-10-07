package main

import (
	"fmt"
	"testing"
)

func TestMaxAscendingSum(t *testing.T) {
	fmt.Println(MaxAscendingSum([]int{10,20,30,5,10,50}))
	fmt.Println(MaxAscendingSum([]int{10,20,30,40,50}))
	fmt.Println(MaxAscendingSum([]int{12,17,15,13,10,11,12}))
	fmt.Println(MaxAscendingSum([]int{100,10,1}))
	fmt.Println(MaxAscendingSum([]int{10}))
}