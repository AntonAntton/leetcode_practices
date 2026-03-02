package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var op [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			current_sum := nums[left] + nums[right]
			if current_sum == target {
				op = append(op, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if current_sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return op
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}
