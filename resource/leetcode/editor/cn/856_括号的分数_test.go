package main

import (
	"fmt"
	"testing"
)

func TestScoreOfParentheses(t *testing.T) {
	fmt.Println(ScoreOfParentheses("()"))
	fmt.Println(ScoreOfParentheses("(())"))
	fmt.Println(ScoreOfParentheses("((())(()))"))
	fmt.Println(ScoreOfParentheses("(()(()))"))
}
