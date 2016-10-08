package strings

import (
	//"fmt"
	"strconv"
)

func Compression(str string) string {
	var compressed string
	var previous string
	var previousNum int = 0

	for i, value := range str {
		if len(compressed) >= len(str)-2 {
			return str
		}
		if previous == "" {
			previous = string(value)
		}
		if previous == string(value) {
			previousNum++
		} else {
			compressed = compressed + previous + strconv.Itoa(previousNum)
			previous = string(value)
			previousNum = 1
		}
		if i == len(str)-1 {
			compressed = compressed + previous + strconv.Itoa(previousNum)
		}
	}
	return compressed
}
