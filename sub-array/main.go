package main

import (
	"fmt"
)

func getSubArrays(arr []int, n int) [][]int {
	var result [][]int
	findSubArray(arr, n, 0, []int{}, &result)
	return result
}

func findSubArray(arr []int, n int, index int, currentSubarray []int, result *[][]int) {
	if index == len(arr) {
		if sum(currentSubarray) == n {
			*result = append(*result, append([]int{}, currentSubarray...))
		}
		return
	}

	findSubArray(arr, n, index+1, append(currentSubarray, arr[index]), result)
	findSubArray(arr, n, index+1, currentSubarray, result)
}

func sum(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}

func main() {
	numbers := []int{0, 1, 6, -4, 2, 3}
	targetSum := 6

	subarrays := getSubArrays(numbers, targetSum)
	for _, subarray := range subarrays {
		fmt.Println(subarray)
	}
}
