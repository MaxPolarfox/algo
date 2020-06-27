package arrays

import (
	"math"
	"sort"
)

func SmallestDifference(arr1, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	a, b := 0, 0
	current, smallest := math.MaxInt32, math.MaxInt32
	smallestPair := []int{}
	for a < len(arr1) && b < len(arr2) {
		first, second := arr1[a], arr2[b]
		if first < second {
			current = second - first
			a++
		} else if second < first {
			current = first - second
			b++
		} else {
			return []int{first, second}
		}

		if current < smallest {
			smallest = current
			smallestPair = []int{first, second}
		}
	}
	return smallestPair
}
