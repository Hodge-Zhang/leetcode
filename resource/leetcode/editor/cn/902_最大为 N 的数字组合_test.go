package main

import (
	"fmt"
	"testing"
)

func TestAtMostNGivenDigitSet(t *testing.T) {
	fmt.Println(AtMostNGivenDigitSet([]string{"1"}, 823))
	fmt.Println(AtMostNGivenDigitSet([]string{"3", "4", "5"}, 4)) //2
	fmt.Println(AtMostNGivenDigitSet([]string{"7", "9"}, 19))
	fmt.Println(AtMostNGivenDigitSet([]string{"7", "9"}, 7))

	fmt.Println(AtMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
}
