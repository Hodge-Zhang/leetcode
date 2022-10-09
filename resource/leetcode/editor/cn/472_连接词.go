package main

//给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。 
//
// 连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。 
//
// 
//
// 示例 1： 
//
// 
//输入：words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","ra
//t","ratcatdogcat"]
//输出：["catsdogcats","dogcatsdog","ratcatdogcat"]
//解释："catsdogcats" 由 "cats", "dog" 和 "cats" 组成; 
//     "dogcatsdog" 由 "dog", "cats" 和 "dog" 组成; 
//     "ratcatdogcat" 由 "rat", "cat", "dog" 和 "cat" 组成。
// 
//
// 示例 2： 
//
// 
//输入：words = ["cat","dog","catdog"]
//输出：["catdog"] 
//
// 
//
// 提示： 
//
// 
// 1 <= words.length <= 104 
// 0 <= words[i].length <= 30 
// words[i] 仅由小写字母组成 
// 0 <= sum(words[i].length) <= 105 
// 
// Related Topics 深度优先搜索 字典树 数组 字符串 动态规划 
// 👍 273 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var (
	wordsMap map[string]bool
)

func FindAllConcatenatedWordsInADict(words []string) []string {
	initWords(words)

	var catWords []string
	for _, word := range words {
		if dfs(word, 0, 0) > 1 {
			catWords = append(catWords, word)
		}
	}
	return catWords
}

func initWords(words []string) {
	wordsMap = make(map[string]bool, len(words))
	for _, word := range words {
		wordsMap[word] = true
	}
}

func dfs(s string, cnt int, pos int) int {
	if pos == len(s) {
		return cnt
	}

	// 快速结束
	//if cnt > 0 && wordsMap[s[pos:]] {
	//	return cnt + 1
	//}

	for i := pos + 1; i <= len(s); i++ {
		word := s[pos:i]
		if !wordsMap[word] {
			continue
		}
		cnt1 := dfs(s, cnt+1, i)
		if cnt1 > 1 {
			return cnt1
		}
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
