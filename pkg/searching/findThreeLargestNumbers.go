package searching

import (
	"fmt"
	"math"
)

func FindThreeLargestNumbers(arr []int) []int {
	// Write your code here.
	larg := []int{math.MinInt32,math.MinInt32,math.MinInt32}
	for _, num := range arr {
		if num > larg[2] {
			shiftAndUpdate(larg, num, 2)
		} else if num > larg[1] {
			shiftAndUpdate(larg, num, 1)

		} else if num > larg[0] {
			shiftAndUpdate(larg, num, 0)

		}
	}
	return larg
}

func shiftAndUpdate (arr []int, num int, idx int) {
	for i := 0; i <= idx; i++ {
		fmt.Printf("in %v, %v, %v\n", arr, num, idx)
		if i == idx {
			arr[i] = num
		} else {
			arr[i] = arr[i+1]
		}
		fmt.Printf("out %v\n", arr)
	}
}