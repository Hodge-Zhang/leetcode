package main

import (
	"fmt"
	"testing"
)

func Test_AreAlmostEqual(t *testing.T) {
	fmt.Println(AreAlmostEqual("bank","kanb"))
	fmt.Println(AreAlmostEqual("banky","kanby"))
}