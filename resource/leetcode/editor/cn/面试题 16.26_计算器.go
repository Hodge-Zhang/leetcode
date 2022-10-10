package main

//给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
//
// 表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。 
//
// 示例 1: 
//
// 输入: "3+2*2"
//输出: 7
// 
//
// 示例 2: 
//
// 输入: " 3/2 "
//输出: 1 
//
// 示例 3: 
//
// 输入: " 3+5 / 2 "
//输出: 5
// 
//
// 说明： 
//
// 
// 你可以假设所给定的表达式都是有效的。 
// 请不要使用内置的库函数 eval。 
// 
// Related Topics 栈 数学 字符串 
// 👍 85 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func Calculate(s string) int {
	// 栈
	//exp := strings.ReplaceAll(s, " ", "") // 干掉空格
	size := len(s)
	stackNum := make([]int, size)
	stackOp := make([]rune, size)
	sizeNum := 0
	sizeOp := 0

	i := 0

	for i < size {
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			lastNum := 0
			for i < size && s[i] >= '0' && s[i] <= '9' {
				lastNum *= 10
				lastNum += int(s[i] - '0')
				i++
			}
			if sizeOp > 0 && stackOp[sizeOp-1] == '*' {
				stackNum[sizeNum-1] *= lastNum
				sizeOp--
			} else if sizeOp > 0 && stackOp[sizeOp-1] == '/' {
				stackNum[sizeNum-1] /= lastNum
				sizeOp--
			} else {
				stackNum[sizeNum] = lastNum
				sizeNum++
			}
			continue
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			stackOp[sizeOp] = rune(s[i])
			sizeOp++
			i++
		}
	}

	// 计算加减 从左向右
	for i := 0; i < sizeOp; i++ {
		if stackOp[i] == '+' {
			stackNum[i+1] += stackNum[i]
		}
		if stackOp[i] == '-' {
			stackNum[i+1] = stackNum[i] - stackNum[i+1]
		}
	}
	return stackNum[sizeNum-1]
}

//leetcode submit region end(Prohibit modification and deletion)
