package main

//给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
//
// 由于答案可能很大，因此 返回答案模 10^9 + 7 。 
//
// 
//
// 示例 1： 
//
// 
//输入：arr = [3,1,2,4]
//输出：17
//解释：
//子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。 
//最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。 
//
// 示例 2： 
//
// 
//输入：arr = [11,81,94,43,3]
//输出：444
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr.length <= 3 * 104 
// 1 <= arr[i] <= 3 * 104 
// 
//
// 
// Related Topics 栈 数组 动态规划 单调栈 
// 👍 583 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func SumSubarrayMins(arr []int) int {
	dp := make([]int, len(arr))
	sums := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for j := 0; j+i < len(arr); j++ {
			if i == 0 {
				dp[j] = j
			} else {
				if arr[dp[j]] > arr[j+i] {
					dp[j] = j + i
				}

			}

			sums[dp[j]]++
		}
	}
	total := 0
	for i, sum := range sums {
		total += arr[i] * sum
	}
	return total % (1e9 + 7)
}

//leetcode submit region end(Prohibit modification and deletion)
