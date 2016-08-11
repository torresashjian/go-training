package main

import (
	"fmt"
)

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	left := MergeSort(slice[:len(slice)/2])
	right := MergeSort(slice[len(slice)/2:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	all := len(left) + len(right)
	merged := make([]int, all)
	l := 0
	r := 0
	for i := 0; i < all; i++ {
		if len(left) == l {
			merged[i] = right[r]
			r++
		} else if len(right) == r {
			merged[i] = left[l]
			l++
		} else if left[l] <= right[r] {
			merged[i] = left[l]
			l++
		} else {
			merged[i] = right[r]
			r++
		}
	}
	return merged
}

func main() {
	unsorted := []int{4, 3, 6, 7, 8, 9, 2, 1, 5}
	fmt.Print(MergeSort(unsorted))

}
