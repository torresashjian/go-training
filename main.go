package main

import (
	"fmt"
	"github.com/migueltorreslopez/go-training/sort"
)

func main_quicksort() {
	unsorted := []int{3, 5, 6, 1, 2, 8, 9, 4, 7}
	// Call quicksort
	qs := sort.NewQuicksort()
	qs.InPlaceSort(unsorted, 0, len(unsorted)-1)
	fmt.Print(unsorted)
}

func main_mergesort2() {
	unsorted := []int{4, 3, 6, 7, 8, 9, 2, 1, 5}
	fmt.Print(sort.MergeSort(unsorted))
}

func main() {
	fmt.Printf("Here")
}
