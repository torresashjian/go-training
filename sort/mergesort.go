package sort

type mergeSort struct {
}

func NewMergeSort() Sort {
	return &mergeSort{}
}

func (ms *mergeSort) InPlaceSort(slice []int, low, high int) {
}

func (ms *mergeSort) Sort(arr []int) []int {
	return ms.doMergeSort(arr)
}

// Recursively divides the slice until there is just one element, and merges them once that is finished
func (ms *mergeSort) doMergeSort(arr []int) []int {
	arrLength := len(arr)
	if arrLength > 1 {
		// The array is not sorted
		left := ms.doMergeSort(arr[:arrLength/2])
		right := ms.doMergeSort(arr[arrLength/2:])
		// merge
		return ms.doMerge(left, right)
	}
	return arr
}

func (ms *mergeSort) doMerge(leftArr []int, rightArr []int) []int {
	totalLength := len(leftArr) + len(rightArr)
	counter := 0
	leftIndx := 0
	rightIndx := 0
	mergedArr := make([]int, totalLength)
	for counter < totalLength {
		if leftIndx >= len(leftArr) {
			mergedArr[counter] = rightArr[rightIndx]
			rightIndx = rightIndx + 1
		} else if rightIndx >= len(rightArr) {
			mergedArr[counter] = leftArr[leftIndx]
			leftIndx = leftIndx + 1
		} else if leftArr[leftIndx] < rightArr[rightIndx] {
			mergedArr[counter] = leftArr[leftIndx]
			leftIndx = leftIndx + 1
		} else {
			mergedArr[counter] = rightArr[rightIndx]
			rightIndx = rightIndx + 1
		}

		counter = counter + 1
	}
	return mergedArr
}
