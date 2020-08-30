package arrays

import (
	"reflect"
)

// MinimumSwaps ...
func MinimumSwaps(q []int32) int32 {

	swapF := reflect.Swapper(q)
	countOfSwaps := 0

	for i := 0; i < len(q); i++ {
		if int(q[i]) != i+1 {
			countOfSwaps++
			swapF(i, int(q[i]-1))
			i--
		}
	}
	return int32(countOfSwaps)
}
