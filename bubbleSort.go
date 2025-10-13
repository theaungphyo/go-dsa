package main

import "fmt"

func main() {
	arr := []int{9, 5, 3, 4, 1, 11, 12, 6, 2}
	sort := BubbleSort(arr)
	fmt.Println("Bubble Sort", sort)
}

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
