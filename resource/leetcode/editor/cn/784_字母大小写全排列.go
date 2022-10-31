package main

//给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
//
// 返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "a1b2"
//输出：["a1b2", "a1B2", "A1b2", "A1B2"]
// 
//
// 示例 2: 
//
// 
//输入: s = "3z4"
//输出: ["3z4","3Z4"]
// 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length <= 12 
// s 由小写英文字母、大写英文字母和数字组成 
// 
// Related Topics 位运算 字符串 回溯 
// 👍 435 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func LetterCasePermutation(s string) []string {
	var ans []string
	var backtrace func(string)
	diff := 'a' - 'A'
	backtrace = func(path string) {
		if len(path) == len(s) {
			ans = append(ans, path)
			return
		}
		c := s[len(path)]
		path += string(c)
		backtrace(path)
		path = path[:len(path)-1]
		if c >= 'A' && c <= 'Z' {
			path += string(int32(c) + diff)
			backtrace(path)
		}

		if c >= 'a' && c <= 'z' {
			path += string(int32(c) - diff)
			backtrace(path)
		}
	}
	backtrace("")
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
