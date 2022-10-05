package main

import (
	"fmt"
	"testing"
)

func TestSubdomainVisits(t *testing.T) {
	fmt.Println(SubdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(SubdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}