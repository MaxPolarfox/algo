package arrays

import "sort"

func ThreeNumberSum(arr []int, targ int) [][]int {
	sort.Ints(arr)
	triplets := [][]int{}

	for i, _ := range arr {
		left, right := i+1, len(arr) - 1

		for left < right {
			sum := arr[left] + arr[right] + arr[i]
			if sum == targ {
				triplets =	append(triplets, []int{arr[i], arr[left], arr[right]})
				left++
				right--
			} else if sum < targ {
				left++
			} else {
				right--
			}
		}
	}

	return triplets
}