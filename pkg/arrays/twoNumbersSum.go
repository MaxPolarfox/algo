package arrays

import (
	"sort"
)

func TwoNumberSum(arr []int, target int) []int {
	// Write your code here.
	sort.Ints(arr)

	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{arr[left], arr[right]}
		}
	}
	return []int{}
}
