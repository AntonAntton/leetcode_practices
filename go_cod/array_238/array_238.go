package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	right := make([]int, n)
	res := make([]int, n)

	prod := 1
	for i := n - 1; i >= 0; i-- {
		prod *= nums[i]
		right[i] = prod
	}

	prod = 1
	for i := 0; i < n-1; i++ {
		lp := prod
		rp := right[i+1]

		res[i] = rp * lp
		prod *= nums[i]
	}
	res[n-1] = prod
	return res
}

func main() {
	nums := []int{5, 5, 5, 5}
	result := productExceptSelf(nums)
	fmt.Println(result)
}
