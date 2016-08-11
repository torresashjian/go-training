package sort

type quicksort struct {
}

func NewQuicksort() Sort {
	return &quicksort{}
}

func (s *quicksort) Sort(arr []int) []int {
	return []int{}
}

func (s *quicksort) InPlaceSort(slice []int, low, high int) {
	if low >= high {
		return
	}
	p := s.partition(slice, low, high)
	s.InPlaceSort(slice, low, p-1)
	s.InPlaceSort(slice, p+1, high)

}

// Partitions the array and returns the pivot index
func (s *quicksort) partition(slice []int, low, high int) int {
	pivotIdx := s.FindPivotIndex(slice, low, high)

	pivot := slice[pivotIdx]
	i := low
	j := high

	for {
		for ; slice[i] < pivot; i++ {

		}
		for ; slice[j] > pivot; j-- {

		}
		if i >= j {
			return j
		}
		// Do swap
		temp := slice[i]
		slice[i] = slice[j]
		slice[j] = temp
	}
}

// FIXME improve performance of pivot, choose median of first mid last
// Returns the index of the pivot element to be chosen
func (s *quicksort) FindPivotIndex(slice []int, low int, high int) int {
	// Return mid index
	return (low + high) / 2
}
