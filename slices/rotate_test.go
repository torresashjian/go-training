package slices

import (
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	m := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}

	Rotate(m)

}
