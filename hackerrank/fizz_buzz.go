package hackerrank

import "fmt"

// FizzBuzz ...
func FizzBuzz(n int32) {
	var i int32 = 0
	for i = 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			print("FizzBuzz\n")
		} else if i%3 == 0 {
			print("Fizz\n")
		} else if i%5 == 0 {
			print("Buzz\n")
		} else {
			fmt.Printf("%v\n", i)
		}
	}
}
