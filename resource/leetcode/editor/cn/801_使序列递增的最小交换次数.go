package main

import "math"

//我们有两个长度相等且不为空的整型数组 nums1 和 nums2 。在一次操作中，我们可以交换 nums1[i] 和 nums2[i]的元素。 
//
// 
// 例如，如果 nums1 = [1,2,3,8] ， nums2 =[5,6,7,4] ，你可以交换 i = 3 处的元素，得到 nums1 =[1,2,3
//,4] 和 nums2 =[5,6,7,8] 。 
// 
//
// 返回 使 nums1 和 nums2 严格递增 所需操作的最小次数 。 
//
// 数组 arr 严格递增 且 arr[0] < arr[1] < arr[2] < ... < arr[arr.length - 1] 。 
//
// 注意： 
//
// 
// 用例保证可以实现操作。 
// 
//
// 
//
// 示例 1: 
//
// 
//输入: nums1 = [1,3,5,4], nums2 = [1,2,3,7]
//输出: 1
//解释: 
//交换 A[3] 和 B[3] 后，两个数组如下:
//A = [1, 3, 5, 7] ， B = [1, 2, 3, 4]
//两个数组均为严格递增的。 
//
// 示例 2: 
//
// 
//输入: nums1 = [0,3,5,8,9], nums2 = [2,1,4,6,9]
//输出: 1
// 
//
// 
//
// 提示: 
//
// 
// 2 <= nums1.length <= 105 
// nums2.length == nums1.length 
// 0 <= nums1[i], nums2[i] <= 2 * 105 
// 
// Related Topics 数组 动态规划 
// 👍 382 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func MinSwap(nums1 []int, nums2 []int) int {
	// 动态规划 当前 是否 需要/可以 交换 只与上一个位置的数有关系
	// dp[i][0] 为未交换i处最小交换次数
	// dp[i][1] 为交换i处的最小交换次数
	// dp[i][0]=min(dp[i-1][0], dp[i-1][1])
	// dp[i][1]=min(dp[i-1][0], dp[i-1][1])+1

	size := len(nums1)

	dp := make([][2]int, size)
	dp[0][1] = 1
	for i := 1; i < size; i++ {
		dp[i] = [2]int{math.MaxInt64, math.MaxInt64}
		// 本次&上次都未交换
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			if dp[i][0] > dp[i-1][0] {
				dp[i][0] = dp[i-1][0]
			}
		}
		// 上次交换
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			if dp[i][0] > dp[i-1][1] {
				dp[i][0] = dp[i-1][1]
			}
		}
		// 本次交换
		if nums2[i] > nums1[i-1] && nums1[i] > nums2[i-1] {
			if dp[i][1] > dp[i-1][0]+1 {
				dp[i][1] = dp[i-1][0] + 1
			}
		}
		// 本次&上次都交换
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			if dp[i][1] > dp[i-1][1]+1 {
				dp[i][1] = dp[i-1][1] + 1
			}
		}
	}
	if dp[size-1][0] > dp[size-1][1] {
		return dp[size-1][1]
	}
	return dp[size-1][0]

}

//leetcode submit region end(Prohibit modification and deletion)
