package main

import (
	"fmt"
	"github.com/torresashjian/go-training/arrays"
	"github.com/torresashjian/go-training/p2p"
	"github.com/torresashjian/go-training/sort"
	"github.com/torresashjian/go-training/strings"
	"github.com/torresashjian/go-training/types"
)

func main(){
	fmt.Println("In TestTypes")
	types.TestTypes()
}

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

func main_discoverupnp() {
	fmt.Println("In Main discoverupnp")
	p2p.DiscoverAllUPnP()
	p2p.DiscoverAllPMPs()
}

func main_palindrome() {
	fmt.Println("In Main is palindrome")
	fmt.Println("Is palindrome 1:", strings.IsPalindrome("abc cba"))
}

func main_arrays_basic() {
	fmt.Println("In main arrays_basic")
	arrays.CreateArray()
}
