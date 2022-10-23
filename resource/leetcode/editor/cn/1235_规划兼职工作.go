package main

import (
	"fmt"
	"sort"
)

//你打算利用空闲时间来做兼职工作赚些零花钱。 
//
// 这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。 
//
// 给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。 
//
// 注意，时间上出现重叠的 2 份工作不能同时进行。 
//
// 如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
//输出：120
//解释：
//我们选出第 1 份和第 4 份工作， 
//时间范围是 [1-3]+[3-6]，共获得报酬 120 = 50 + 70。
// 
//
// 示例 2： 
//
// 
//
// 输入：startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60
//]
//输出：150
//解释：
//我们选择第 1，4，5 份工作。 
//共获得报酬 150 = 20 + 70 + 60。
// 
//
// 示例 3： 
//
// 
//
// 输入：startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
//输出：6
// 
//
// 
//
// 提示： 
//
// 
// 1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4 
// 1 <= startTime[i] < endTime[i] <= 10^9 
// 1 <= profit[i] <= 10^4 
// 
// Related Topics 数组 二分查找 动态规划 排序 
// 👍 193 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func JobScheduling(startTime []int, endTime []int, profit []int) int {
	// 动态规划 dp[i]=max(dp[i-1], dp[j]+ profit[j]) 如果时间i有结束的工作
	size := len(startTime)
	jobs := make([][3]int, size)
	for i := 0; i < size; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	var timePoint []int
	timePoint = append(timePoint, startTime...)
	timePoint = append(timePoint, endTime...)
	sort.Ints(timePoint)

	timeMap := make(map[int]int, size*2)
	for i, v := range timePoint {
		timeMap[v] = i
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	fmt.Println(jobs)

	dp := make([]int, size*2) // dp[i] 对应的时间为timePoint[i]
	//dp[0]=0
	idx := 0
	for i := 1; i < size*2; i++ {
		if idx >= size || timePoint[i] != jobs[idx][1] { // end time
			dp[i] = dp[i-1]
			continue
		}
		for idx < size && timePoint[i] == jobs[idx][1] { // 可能存在多个end time相同的job
			dpi:=timeMap[jobs[idx][0]]
			if dp[i] < dp[dpi]+jobs[idx][2] {
				dp[i] = dp[dpi] + jobs[idx][2]
			}
			if dp[i] < dp[i-1] {
				dp[i] = dp[i-1]
			}
			idx++
		}
	}
	fmt.Println(dp)
	return dp[2*size-1]
}

//leetcode submit region end(Prohibit modification and deletion)
