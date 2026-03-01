package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int

	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			res = append(res, perm)
			return
		}

		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	dfs(0)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println(result) // Output: [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 2, 1], [3, 1, 2]]
}
