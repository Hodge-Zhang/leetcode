package main

import (
	"fmt"
	"testing"
)

func TestTotalFruit(t *testing.T) {
	fmt.Println(TotalFruit([]int{1, 2, 1}))
	fmt.Println(TotalFruit([]int{0, 1, 2, 2}))
	fmt.Println(TotalFruit([]int{1, 2, 3, 2, 2}))
	fmt.Println(TotalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}
