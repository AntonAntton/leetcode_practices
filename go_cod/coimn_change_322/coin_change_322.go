package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for _, c := range coins {
		for a := 1; a <= amount; a++ {
			if a >= c {
				dp[a] = min(dp[a], 1+dp[a-c])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func main() {
	println(coinChange([]int{1, 2, 5}, 11))
	println(coinChange([]int{2}, 3))
	println(coinChange([]int{1}, 0))
	println(coinChange([]int{1}, 1))
	println(coinChange([]int{1}, 2))
}
