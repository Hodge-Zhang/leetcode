package main

//给定一个由 0 和 1 组成的数组 arr ，将数组分成 3 个非空的部分 ，使得所有这些部分表示相同的二进制值。 
//
// 如果可以做到，请返回任何 [i, j]，其中 i+1 < j，这样一来： 
//
// 
// arr[0], arr[1], ..., arr[i] 为第一部分； 
// arr[i + 1], arr[i + 2], ..., arr[j - 1] 为第二部分； 
// arr[j], arr[j + 1], ..., arr[arr.length - 1] 为第三部分。 
// 这三个部分所表示的二进制值相等。 
// 
//
// 如果无法做到，就返回 [-1, -1]。 
//
// 注意，在考虑每个部分所表示的二进制时，应当将其看作一个整体。例如，[1,1,0] 表示十进制中的 6，而不会是 3。此外，前导零也是被允许的，所以 [0,
//1,1] 和 [1,1] 表示相同的值。 
//
// 
//
// 示例 1： 
//
// 
//输入：arr = [1,0,1,0,1]
//输出：[0,3]
// 
//
// 示例 2： 
//
// 
//输入：arr = [1,1,0,1,1]
//输出：[-1,-1] 
//
// 示例 3: 
//
// 
//输入：arr = [1,1,0,0,1]
//输出：[0,2]
// 
//
// 
//
// 提示： 
// 
//
// 
// 3 <= arr.length <= 3 * 104 
// arr[i] 是 0 或 1 
// 
// Related Topics 数组 数学 
// 👍 102 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func ThreeEqualParts(arr []int) []int {
	// 每个part 1的数量应该相等
	// 如果1的数量相等后，再看值是否相同,注意尾部的0的数量也应该相等
	cnt1 := 0
	failed := []int{-1, -1}
	for _, i := range arr {
		if i == 1 {
			cnt1++
		}
	}
	if cnt1%3 != 0 {
		return failed
	}

	if cnt1 == 0 {
		return []int{0, 2}
	}
	// 计数结尾的0，每个part结尾0的数量相同
	var cnt0 int
	for idx := len(arr) - 1; idx >= 0; idx-- {
		if arr[idx] == 0 {
			cnt0++
			continue
		}
		break
	}

	var i, j int
	for cnt := 0; cnt != cnt1/3; {
		if arr[i] == 1 {
			cnt++
		}
		i++
	}
	for c := 0; c != cnt0; {
		if arr[i] == 0 {
			c++
			i++
			continue
		}
		return failed
	}
	j = i
	for cnt := 0; cnt != cnt1/3; {
		if arr[j] == 1 {
			cnt++
		}
		j++
	}

	for c := 0; c != cnt0; {
		if arr[j] == 0 {
			c++
			j++
			continue
		}
		return failed
	}

	si, sj, sk := 0, i, j
	for {
		if arr[si] == 0 {
			si++
			continue
		}
		break
	}

	for {
		if arr[sj] == 0 {
			sj++
			continue
		}
		break
	}

	for {
		if arr[sk] == 0 {
			sk++
			continue
		}
		break
	}

	if i-si != j-sj || len(arr)-sk != i-si {
		return failed
	}
	for ii := 0; ii < i-si; ii++ {
		if arr[si+ii] != arr[sj+ii] || arr[si+ii] != arr[sk+ii] {
			return failed
		}
	}

	return []int{i - 1, j}

}

//leetcode submit region end(Prohibit modification and deletion)
