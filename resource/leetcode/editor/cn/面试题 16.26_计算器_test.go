package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	fmt.Println(Calculate("3+2*2"))
	fmt.Println(Calculate("3/2"))
	fmt.Println(Calculate("3+5/2"))
	fmt.Println(Calculate("1-1+1-3"))
}
