package main

import (
	"fmt"
	"testing"
)

func TestDistinctSubseqII(t *testing.T) {
	fmt.Println(DistinctSubseqII("abc"))
	fmt.Println(DistinctSubseqII("abacad"))
	fmt.Println(DistinctSubseqII("aba"))
	fmt.Println(DistinctSubseqII("acac"))
}
