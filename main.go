package main

import (
	"fmt"

	"github.com/ppichugin/manning-quick-sort/helper"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := helper.MakeRandomSlice(numItems, max)
	helper.PrintSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(slice)
	helper.PrintSlice(slice, 40)

	// Verify that it's sorted.
	helper.CheckSorted(slice)
}

func quicksort(arr []int) {
	quickSortCore(arr, 0, len(arr)-1)
}

func quickSortCore(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortCore(arr, low, p-1)
		arr = quickSortCore(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
