package main

import (
	"fmt"
	"testing"
)

func TestAdvantageCount(t *testing.T) {
	fmt.Println(AdvantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(AdvantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
	fmt.Println(AdvantageCount([]int{2,0,4,1,2}, []int{1,3,0,0,2}))

}
