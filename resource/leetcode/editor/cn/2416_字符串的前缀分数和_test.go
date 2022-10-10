package main

import (
	"fmt"
	"testing"
)

func TestSumPrefixScores(t *testing.T) {
	fmt.Println(SumPrefixScores([]string{"abc", "ab", "bc", "b"}))
	fmt.Println(SumPrefixScores([]string{"abcd"}))
}
