package main

import "fmt"

func SelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}

	return arr
}

func main() {
	arr := []int{9, 5, 3, 4, 1, 11, 12, 6, 2}
	fmt.Println("Before sorting:", arr)

	sorted := SelectionSort(arr)
	fmt.Println("After sorting:", sorted)
}
