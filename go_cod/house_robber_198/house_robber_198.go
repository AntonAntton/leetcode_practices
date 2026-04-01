package main

import "fmt"

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums)) // Output: 4
}
