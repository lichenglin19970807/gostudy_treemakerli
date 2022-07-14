// Package backtracking include my solution of backtracking problem.
package backtracking

// 组合
func combine(n, k int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtracking func(track []int, start int)
	backtracking = func(track []int, start int) {
		// 终止条件
		if len(track) == k {
			path := make([]int, k)
			copy(path, track)
			res = append(res, path)
			return
		}
		// 单层搜索
		for i := start; i <= n; i++ {
			// 处理节点
			track = append(track, i)
			backtracking(track, i+1)
			// 回溯
			track = track[:len(track)-1]
		}
	}
	backtracking(track, 1)
	return res
}
