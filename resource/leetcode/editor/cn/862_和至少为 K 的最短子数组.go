package main

//给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回
//-1 。 
//
// 子数组 是数组中 连续 的一部分。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：nums = [1], k = 1
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2], k = 4
//输出：-1
// 
//
// 示例 3： 
//
// 
//输入：nums = [2,-1,2], k = 3
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 105 
// -105 <= nums[i] <= 105 
// 1 <= k <= 109 
// 
// Related Topics 队列 数组 二分查找 前缀和 滑动窗口 单调队列 堆（优先队列）
// 👍 586 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func ShortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSumArr := make([]int, n+1)
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num
	}
	ans := n + 1
	q := []int{}
	for i, curSum := range preSumArr {
		for len(q) > 0 && curSum-preSumArr[q[0]] >= k {
			if ans > i-q[0] {
				ans = i - q[0]
			}
			q = q[1:]
		}
		for len(q) > 0 && preSumArr[q[len(q)-1]] >= curSum {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans < n+1 {
		return ans
	}
	return -1

	// timeout
	//sum := make([]int, len(nums))
	//
	//for i, num := range nums {
	//	if i == 0 {
	//		sum[0] = nums[0]
	//		continue
	//	}
	//	sum[i] = sum[i-1] + num
	//}
	//for i := 1; i <= len(nums); i++ {
	//	v := sum[i-1]
	//	for j := i; j <= len(nums); j++ {
	//		if v >= k {
	//			return i
	//		}
	//		if j != len(nums) {
	//			v = sum[j] - sum[j-i]
	//		}
	//	}
	//}
	//return -1

}

//leetcode submit region end(Prohibit modification and deletion)
