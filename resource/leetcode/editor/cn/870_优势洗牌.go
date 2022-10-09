package main

import (
	"sort"
)

//给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数
//目来描述。 
//
// 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
//输出：[2,11,7,15]
// 
//
// 示例 2： 
//
// 
//输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
//输出：[24,32,8,12]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums1.length <= 105 
// nums2.length == nums1.length 
// 0 <= nums1[i], nums2[i] <= 109 
// 
// Related Topics 贪心 数组 双指针 排序 
// 👍 231 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func AdvantageCount(nums1 []int, nums2 []int) []int {
	// 贪心，每次都比较num2中最大的值
	size := len(nums1)
	var sortedNums2 [][2]int
	for i, n := range nums2 {
		sortedNums2 = append(sortedNums2, [2]int{i, n})
	}

	sort.Ints(nums1)
	sort.Slice(sortedNums2, func(i, j int) bool {
		return sortedNums2[i][1] > sortedNums2[j][1]
	})

	l, r := 0, size-1
	res := make([]int, size)

	for i := 0; i < size; i++ {
		if nums1[r] > sortedNums2[i][1] {
			res[sortedNums2[i][0]] = nums1[r]
			r--
		} else {
			res[sortedNums2[i][0]] = nums1[l] // 赋为最小值
			l++
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
