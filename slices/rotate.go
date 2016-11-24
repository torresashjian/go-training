package slices

import (
	"fmt"
)

func Rotate(m [][]byte) [][]byte {
	var n [][]byte
	for x := range m {
		for j := range m[x] {
			fmt.Println(string(m[x][j]))
		}
	}

	return n
}
