package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid -1
		}
	}
	return -1
}

func main() {
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	result := binarySearch(sortedArray, target)

	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
