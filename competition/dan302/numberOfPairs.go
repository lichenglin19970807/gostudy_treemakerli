package dan302

import "fmt"

func numberOfPairs(nums []int) []int {
	pairs := 0
	for {
		ok, v := removePairs(nums)
		nums = v
		if ok {
			pairs++
		} else {
			break
		}
	}
	res := make([]int, 0)
	res = append(res, pairs)
	res = append(res, len(nums))
	return res
}

func removePairs(nums []int) (bool, []int) {
	set := make(map[int]int) // key:elem int:index
	for i, v := range nums {
		if j, ok := set[v]; !ok {
			set[v] = i
		} else {
			nums = removeElemByIndex(nums, i, j)
			fmt.Println(nums)
			return true, nums
		}
	}
	return false, nums
}

func removeElemByIndex(nums []int, x, y int) []int {
	if x < len(nums)-1 {
		for i := x + 1; i < len(nums); i++ {
			nums[i-1] = nums[i]
		}
	}
	nums = nums[:len(nums)-1]
	for i := y + 1; i < len(nums); i++ {
		nums[i-1] = nums[i]
	}
	nums = nums[:len(nums)-1]
	return nums
}
