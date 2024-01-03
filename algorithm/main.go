package main

import (
	"fmt"
)

// getSubArrays takes an array and a target sum
// and return all subarrays with the given sum.
func getSubArrays(arr []int, n int) [][]int {
	var result [][]int

	findSubArray(arr, n, 0, []int{}, &result)
	return result
}

// all possible combinations of the including or excluding each element
// in the subarray to find those with the target sum.
func findSubArray(arr []int, n int, index int, currentSubarray []int, result *[][]int) {
	if index == len(arr) {
		// check if the current subarray has the target sum and add it to the result.
		if sum(currentSubarray) == n {
			*result = append(*result, append([]int{}, currentSubarray...))
		}
		return
	}

	// Include the current element in the subarray and explore further.
	findSubArray(arr, n, index+1, append(currentSubarray, arr[index]), result)

	// Exclude the current element from the subarray and explore further.
	findSubArray(arr, n, index+1, currentSubarray, result)
}

// sym calculate the sum of elements in an array.
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

	// Find the print subarrays with the target sum.
	subarrays := getSubArrays(numbers, targetSum)
	for _, subarray := range subarrays {
		fmt.Println(subarray)
	}
}
