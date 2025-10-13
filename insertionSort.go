package main

import "fmt"

func main() {
	arr := []int{9, 5, 3, 4, 1, 11, 12, 6, 2}
	sortArr := InsertionSort(arr)

	sortDescArr := InsertionSortDesc(arr)

	fmt.Println("Sorted array:", sortArr)
	fmt.Println("Sorted desc array:", sortDescArr)
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}

func InsertionSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] > arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}
