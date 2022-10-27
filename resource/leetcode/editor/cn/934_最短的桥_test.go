package main

import (
	"fmt"
	"testing"
)

func TestShortestBridge(t *testing.T) {
	fmt.Println(ShortestBridge([][]int{{0, 1}, {1, 0}}))
	fmt.Println(ShortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(ShortestBridge([][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}))
}
