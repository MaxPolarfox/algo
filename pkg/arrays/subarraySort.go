package arrays

import "math"

//Subarray Sort Write a function that takes in an array of at least two integers and that returns an array of the starting and ending indices of the smallest subarray in the input array that needs to be sorted in place in order for the entire input array to 7, 7, 7, be sorted (in ascending order). If the input array is already sorted, the function should return [-1, -1]
//Sample Input array
//Sample Output 10, 11 12, 16, 18, -1] 19]

// Time: O(n)  Space: O(1)
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
