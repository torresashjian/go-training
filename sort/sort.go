// Package with sort interface
package sort

type Sort interface {
	Sort(arr []int) []int
	InPlaceSort(arr []int, left int, right int)
}
