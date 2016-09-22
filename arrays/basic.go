package arrays

import (
	"fmt"
)

func CreateArray() {
	// Create a new array
	var my_arr [2]string
	my_arr[0] = "1"
	my_arr[1] = "2"
	fmt.Printf("My Array Values: %s , %s", my_arr[0], my_arr[1])
}
