package main

//给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
//
// 岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。 
//
// 
// 
// 你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。 
//
// 返回必须翻转的 0 的最小数目。 
// 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：grid = [[0,1],[1,0]]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
//输出：2
// 
//
// 示例 3： 
//
// 
//输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// n == grid.length == grid[i].length 
// 2 <= n <= 100 
// grid[i][j] 为 0 或 1 
// grid 中恰有两个岛 
// 
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 
// 👍 387 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func ShortestBridge(grid [][]int) int {
	n := len(grid)
	color := make([][]int, n)
	for x := 0; x < n; x++ {
		color[x] = make([]int, n)
	}
	// dfs找到一个岛的所有位置
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x >= n || x < 0 || y < 0 || y >= n || color[x][y] != 0 { // color[x][y] != 0 已访问过
			return
		}
		if grid[x][y] == 0 {
			color[x][y] = -1
			return
		}
		color[x][y] = 1
		dfs(x-1, y)
		dfs(x, y-1)
		dfs(x+1, y)
		dfs(x, y+1)
	}
A:
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 1 {
				dfs(x, y)
				break A
			}
		}
	}

	dis := 1

	for true {
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				if color[x][y] == dis {
					if x-1 >= 0 && color[x-1][y] <= 0 {
						if grid[x-1][y] == 1 {
							return dis - 1
						}
						color[x-1][y] = dis + 1
					}
					if x+1 < n && color[x+1][y] <= 0 {
						if grid[x+1][y] == 1 {
							return dis - 1
						}
						color[x+1][y] = dis + 1
					}
					if y-1 >= 0 && color[x][y-1] <= 0 {
						if grid[x][y-1] == 1 {
							return dis - 1
						}
						color[x][y-1] = dis + 1
					}
					if y+1 < n && color[x][y+1] <= 0 {
						if grid[x][y+1] == 1 {
							return dis - 1
						}
						color[x][y+1] = dis + 1
					}
				}
			}
		}
		dis++
	}
	return dis - 1
}

//leetcode submit region end(Prohibit modification and deletion)
