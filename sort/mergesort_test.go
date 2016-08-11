package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	fmt.Println("Merge Sorting")
	// Create the arr to be sorted
	unsortedArr := []int{6, 5, 3, 1, 8, 7, 2, 4, 4}
	sortedArr := NewMergeSort().Sort(unsortedArr)
	for _, value := range sortedArr {
		fmt.Println(value)
	}
}

func TestMergeSortEmpty(t *testing.T) {
	fmt.Println("Merge Sorting Empty")
	// Create the arr to be sorted
	unsortedArr := []int{}
	sortedArr := NewMergeSort().Sort(unsortedArr)
	for _, value := range sortedArr {
		fmt.Println(value)
	}
}

func TestMergeSortOne(t *testing.T) {
	fmt.Println("Merge Sorting One")
	// Create the arr to be sorted
	unsortedArr := []int{3}
	sortedArr := NewMergeSort().Sort(unsortedArr)
	for _, value := range sortedArr {
		fmt.Println(value)
	}
}

func TestMergeSortTwo(t *testing.T) {
	fmt.Println("Merge Sorting Two")
	// Create the arr to be sorted
	unsortedArr := []int{7, 3}
	sortedArr := NewMergeSort().Sort(unsortedArr)
	for _, value := range sortedArr {
		fmt.Println(value)
	}
}
