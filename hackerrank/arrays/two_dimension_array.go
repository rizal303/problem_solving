package arrays

// HourglassSum ...
func HourglassSum(arr [][]int32) int32 {
	count := 4
	var maxSum int32 = -63

	for r := 0; r < count; r++ {
		for c := 0; c < count; c++ {

			top := sum(arr[r][c : c+3])
			mid := arr[r+1][c+1]
			bottom := sum(arr[r+2][c : c+3])

			hourglass := top + mid + bottom

			if hourglass > maxSum {
				maxSum = hourglass
			}
		}
	}

	return maxSum
}

func sum(array []int32) int32 {
	var result int32 = 0
	for _, v := range array {
		result += v
	}
	return result
}
