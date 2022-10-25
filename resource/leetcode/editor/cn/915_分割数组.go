package main

import (
	"fmt"
)

//给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得： 
//
// 
// left 中的每个元素都小于或等于 right 中的每个元素。 
// left 和 right 都是非空的。 
// left 的长度要尽可能小。 
// 
//
// 在完成这样的分组后返回 left 的 长度 。 
//
// 用例可以保证存在这样的划分方法。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [5,0,3,8,6]
//输出：3
//解释：left = [5,0,3]，right = [8,6]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,1,1,0,6,12]
//输出：4
//解释：left = [1,1,1,0]，right = [6,12]
// 
//
// 
//
// 提示： 
//
// 
// 2 <= nums.length <= 105 
// 0 <= nums[i] <= 106 
// 可以保证至少有一种方法能够按题目所描述的那样对 nums 进行划分。 
// 
// Related Topics 数组 
// 👍 213 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func PartitionDisjoint(nums []int) int {
	// 右边维护一个单调递减的栈
	var s []int
	top := -1
	for i := len(nums) - 1; i > 0; i-- {
		if i == len(nums)-1 {
			s = append(s, nums[i])
			top++
			continue
		}
		if nums[i] <= s[top] {
			s = append(s, nums[i])
			top++
		}
	}
	fmt.Println(s)
	left, max := 1, nums[0]
	for ; left < len(nums); left++ {
		if max <= s[top] {
			break
		}
		if nums[left] == s[top] {
			top--
		}
		if max < nums[left] {
			max = nums[left]
		}
	}
	return left
}

//leetcode submit region end(Prohibit modification and deletion)
