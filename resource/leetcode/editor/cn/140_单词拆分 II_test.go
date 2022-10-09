package main

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	fmt.Println(WordBreak("catsanddog", []string{"cat","cats","and","sand","dog"}))
	fmt.Println(WordBreak("catsandog", []string{"cat","cats","and","sand","dog"}))
}
