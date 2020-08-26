package arrays

import (
	"fmt"
	"reflect"
)

// MinimumBribes ...
func MinimumBribes(q []int32) {

	countOfQueque := len(q)
	start := countOfQueque - 1
	countOfBribes := 0
	swapF := reflect.Swapper(q)
	for i := start; i >= 0; i-- {

		if int(q[i]) != (i + 1) {
			if ((i - 1) >= 0) && int(q[i-1]) == (i+1) {
				countOfBribes++
				swapF(i, i-1)
			} else if ((i - 2) >= 0) && int(q[i-2]) == (i+1) {
				countOfBribes += 2
				swapF(i-2, i-1)
				swapF(i-1, i)
			} else {
				fmt.Print("Too chaotic\n")
				return
			}
		}
	}
	fmt.Printf("%v\n", countOfBribes)
}
