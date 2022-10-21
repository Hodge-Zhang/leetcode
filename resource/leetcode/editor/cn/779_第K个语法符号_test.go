package main

import (
	"fmt"
	"testing"
)

func TestKthGrammar(t *testing.T) {
	fmt.Println(KthGrammar(1, 1))
	fmt.Println(KthGrammar(2, 1))
	fmt.Println(KthGrammar(2, 2))
	fmt.Println(KthGrammar(20, 1234))
}
