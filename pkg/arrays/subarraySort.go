package arrays

import "math"

func SubarraySort(arr []int) []int {
	minOutOfOrder, maxOutOfOrder := math.MaxInt32, math.MinInt32

	for i, num := range arr {
		if isOutOfOrder(i, num, arr) {
			minOutOfOrder = min(minOutOfOrder, num)
			maxOutOfOrder = max(maxOutOfOrder, num)
		}
	}

	if minOutOfOrder == math.MaxInt32 {
		return []int{-1,-1}
	}
	subArrLeft := 0
	for minOutOfOrder >=  arr[subArrLeft] {
		subArrLeft++
	}

	subArrRight := len(arr)-1
	for maxOutOfOrder <=  arr[subArrRight] {
		subArrRight--
	}
	return []int{subArrLeft, subArrRight}
}


func isOutOfOrder(i int, num int, arr []int) bool {
	if i == 0 {
		return num > arr[i+1]
	}
	if i == len(arr) - 1 {
		return num < arr[i-1]
	}
	return num > arr[i+1] || num < arr[i-1]
}
