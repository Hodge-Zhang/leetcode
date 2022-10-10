package util

type TrieNode struct {
	val     int64
	charMap map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		val:     0,
		charMap: make(map[rune]*TrieNode),
	}
}

