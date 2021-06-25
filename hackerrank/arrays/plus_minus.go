package arrays

import "fmt"

func plusMinus(arr []int32) {

	var sumPositif, sumNegatif, sumZero float64
	n := float64(len((arr)))

	for i := 0; i < int(n); i++ {
		if arr[i] < 0 {
			sumNegatif += 1
		} else if arr[i] > 0 {
			sumPositif += 1
		} else {
			sumZero += 1
		}
	}

	fmt.Printf("%f \n", sumPositif/n)
	fmt.Printf("%f \n", sumNegatif/n)
	fmt.Printf("%f \n", sumZero/n)

}
