package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		mid_val := matrix[mid/n][mid%n]

		if mid_val == target {
			return true
		} else if mid_val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 2, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	result := searchMatrix(matrix, target)
	fmt.Println(result) // Output: true
}
