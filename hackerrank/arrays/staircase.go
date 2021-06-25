package arrays

import "fmt"

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		for j := 1; j <= int(n); j++ {
			if int(n)-i > j {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
