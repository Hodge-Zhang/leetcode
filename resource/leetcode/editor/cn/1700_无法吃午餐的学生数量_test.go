package main

import (
	"fmt"
	"testing"
)

func TestCountStudents(t *testing.T) {
	fmt.Println(CountStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(CountStudents([]int{0, 0, 0, 1}, []int{1, 1, 0, 1}))
	fmt.Println(CountStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0,1, 1}))
}
