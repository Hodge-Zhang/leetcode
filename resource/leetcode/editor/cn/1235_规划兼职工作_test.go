package main

import (
	"fmt"
	"testing"
)

func TestJobScheduling(t *testing.T) {
	fmt.Println(JobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 1000000000}, []int{50, 10, 40, 70})) // 这种情况需要优化内存分配
	fmt.Println(JobScheduling([]int{4, 2, 4, 8, 2}, []int{5, 5, 5, 10, 8}, []int{1, 2, 8, 10, 4}))
	fmt.Println(JobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
	fmt.Println(JobScheduling([]int{1}, []int{3}, []int{50}))
	fmt.Println(JobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
}

//[1,2,3,3]
//[3,4,5,1000000000]
//[50,10,40,70]
