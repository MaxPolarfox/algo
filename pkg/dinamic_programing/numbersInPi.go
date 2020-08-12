package dinamic_programing

import "math"

// Given a string representation of the first n digits of Pi and a list of positive integers (all in string format),
// write a function that returns the smallest number of spaces that can be added to the n digits of Pi such that all
// resulting numbers are found in the list of integers.
// Note that a single number can appear multiple times in the resulting numbers. For example,
// if Pi is "3141" and the numbers are "3" , the number "1" is allowed to appear twice in the list of resulting numbers after three spaces are added:
// "3 | 1 | 4 | 1" If no number of spaces to be added exists such that all resulting numbers are found in the list of integers,
// the function should return -1
// Sample Input pi = "3141592653589793238462643383279" numbers = [
													//    "314159265358979323846",
													//    "26433",
													//    "8",
													//    "3279",
													//    "314159265",
													//    "35897932384626433832",
													//    "79"
													//  ]
// Sample Output  2

func NumbersInPi(pi string, numbers []string) int {
	// Refactor numbers into map
	numbersMap := map[string]bool{}
	for _, number := range numbers {
		numbersMap[number] = true
	}

	// Create cache that stores min number of spaces on each index
	cache := map[int]int{}
	for i := len(pi) - 1; i >= 0; i-- {
		getMinSpaces(pi, numbersMap, cache, i)
	}
	if cache[0] == math.MaxInt32 {
		return -1
	}
	return cache[0]
}

func getMinSpaces(pi string, numbersMap map[string]bool, cache map[int]int, idx int) int {
	if idx == len(pi) {
		return -1
	} else if val, found := cache[idx]; found {
		return val
	}

	minSpaces := math.MaxInt32
	// get all the prefixes starting from the index until len(p)
	for i := idx; i < len(pi); i++ {
		prefix := pi[idx : i+1]
		// check if we have prefix in out numbersMap
		if _, found := numbersMap[prefix]; found {
			// call getMinSpaces recursevly to check if we have anything in cache for i+1
			minSpacesInSuffix := getMinSpaces(pi, numbersMap, cache, i+1)
			minSpaces = min(minSpaces, minSpacesInSuffix+1)
		}
	}
	cache[idx] = minSpaces
	return cache[idx]
}
