package main

import (
	"fmt"
	"testing"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	fmt.Println(FindAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}
