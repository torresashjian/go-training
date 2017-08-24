package limits

import "fmt"

// Use a uint64
func TestMemLimits() {
	maxint64 := 1<<63 - 1

	size, exponent := getSize(uint64(maxint64))

	fmt.Printf("Size: '%d', exponent: '%d' \n", size, exponent)

	fmt.Printf("size>>6: '%d' \n", size>>6)

	limitSet := make([]uint64, size>>6)

	for key, _ := range limitSet {
		limitSet[key] = 1
	}

}

func getSize(ui64 uint64) (size uint64, exponent uint64) {
	if ui64 < uint64(512) {
		ui64 = uint64(512)
	}
	size = uint64(1)
	for size < ui64 {
		size <<= 1
		exponent++
	}
	return size, exponent
}
