package helper

import (
	"fmt"
	"math/rand"
	"time"
)

// MakeRandomSlice makes a slice containing pseudorandom numbers in [0, max).
func MakeRandomSlice(numItems, max int) []int {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

// PrintSlice prints at most numItems items.
func PrintSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

// CheckSorted verifies that the slice is sorted.
func CheckSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			println("The slice is NOT sorted!")
			return
		}
	}
	println("The slice is sorted")
}
