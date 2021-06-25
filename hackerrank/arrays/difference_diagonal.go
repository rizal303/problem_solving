package arrays

import "math"

func diagonalDifference(arr [][]int32) int32 {
	var sumD1, sumD2 float64
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == j {
				sumD1 += float64(arr[i][j])
			}
			if i == len(arr[i])-j-1 {
				sumD2 += float64(arr[i][j])
			}
		}
	}

	diff := math.Abs(sumD1 - sumD2)

	return int32(diff)
}
