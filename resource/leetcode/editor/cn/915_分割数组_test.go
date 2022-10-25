package main

import (
	"fmt"
	"testing"
)

func TestPartitionDisjoint(t *testing.T) {
	fmt.Println(PartitionDisjoint([]int{5, 0, 3, 8, 6}))
	fmt.Println(PartitionDisjoint([]int{1,1}))
	fmt.Println(PartitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
}
