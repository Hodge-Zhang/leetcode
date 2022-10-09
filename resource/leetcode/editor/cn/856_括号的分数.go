package main

//给定一个平衡括号字符串 S，按下述规则计算该字符串的分数： 
//
// 
// () 得 1 分。 
// AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。 
// (A) 得 2 * A 分，其中 A 是平衡括号字符串。 
// 
//
// 
//
// 示例 1： 
//
// 输入： "()"
//输出： 1
// 
//
// 示例 2： 
//
// 输入： "(())"
//输出： 2
// 
//
// 示例 3： 
//
// 输入： "()()"
//输出： 2
// 
//
// 示例 4： 
//
// 输入： "(()(()))"
//输出： 6
// 
//
// 
//
// 提示： 
//
// 
// S 是平衡括号字符串，且只含有 ( 和 ) 。 
// 2 <= S.length <= 50 
// 
// Related Topics 栈 字符串 
// 👍 353 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func ScoreOfParentheses(s string) int {
	// 递归出口
	if s == "()" {
		return 1
	}

	// 判断算法模式
	lp := 0
	idx := 0 // 右括号位置
	for i, c := range s {
		if i == 0 {
			lp++
			continue
		}
		if c == '(' {
			lp++
		} else {
			lp--
			if lp == 0 {
				idx = i
				break
			}
		}
	}
	if idx == len(s)-1 {
		return 2 * ScoreOfParentheses(s[1:len(s)-1])
	}
	return ScoreOfParentheses(s[:idx+1]) + ScoreOfParentheses(s[idx+1:])
}

//leetcode submit region end(Prohibit modification and deletion)
