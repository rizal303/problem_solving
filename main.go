package main

import (
	"problem_solving/hackerrank/arrays"
)

func main() {
	values := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 2, 2, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	result := arrays.HourglassSum(values)
	print(result)
}
