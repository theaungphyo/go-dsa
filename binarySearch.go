package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 7, 9, 11, 13, 15}
	target := 9

	index := BinarySearch(arr, target)

	fmt.Printf("Found %d at index %d\n", target, index)
}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
