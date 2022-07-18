package backtracking

// 分割回文串
func partition(s string) [][]string {
	res := make([][]string, 0)
	track := make([]string, 0)
	var backtracking func(track []string, start int)
	backtracking = func(track []string, start int) {
		if start == len(s) { // 终止条件
			path := make([]string, len(track))
			copy(path, track)
			res = append(res, path)
		}
		for i := start; i < len(s); i++ {
			// 处理节点
			if isPalindrome(s, start, i) {
				track = append(track, s[start:i+1])
			} else {
				continue
			}
			backtracking(track, i+1)
			// 回溯
			track = track[:len(track)-1]
		}
	}
	backtracking(track, 0)
	return res
}

// 判断是否回文
func isPalindrome(s string, start, i int) bool {
	for start < i {
		if s[start] != s[i] {
			return false
		}
		start++
		i--
	}
	return true
}
