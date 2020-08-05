package dinamic_programing

// You're given an array of arrays where each subarray holds two integer values and represents an item; the first
// integer is the item's value, and the second integer is the item's weight.
// You're also given an integer representing the maximum capacity of a knapsack that you have.
// Your goal is to fit items in your knapsack without having the sum of their weights exceed the knapsack's capacity,
// all the while maximizing their combined value. Note that you only have one of each item at your disposal.
// Write a function that returns the maximized combined value of the items that you should pick as well as an array
// of the indices of each item picked. If there are multiple combinations of items that maximize the total value in
// the knapsack, your function can return any of them.
// Sample Input: items = [[1, 2], [4, 3], [5, 6], [6, 7]] capacity = 10
// Sample Output: [10, [1, 3]] // items [4, Hints 3] and [6, 7]

// Time: O(nc) Space: O(nc)
func KnapsackProblem(items [][]int, capacity int) []interface{} {
	values := make([][]int, len(items)+1)
	for i := range values {
		values[i] = make([]int, capacity+1)
	}

	for i := 1; i < len(items)+1; i++ {
		currentValue := items[i-1][0]
		currentWeight := items[i-1][1]
		for c := 0; c < capacity+1; c++ {
			if currentWeight > c {
				values[i][c] = values[i-1][c]
			} else {
				values[i][c] = max(values[i-1][c], values[i-1][c-currentWeight]+currentValue)
			}
		}
	}

	value := values[len(items)][capacity]
	sequence := getKnapsackItems(values, items)

	return []interface{}{
		value,          // total value
		sequence, // item indices
	}
}

func getKnapsackItems(values [][]int, items [][]int) []int {
	sequence := []int{}

	i, c := len(values)-1, len(values[0])-1
	for i > 0 {
		if values[i][c] == values[i-1][c] {
			i--
		} else {
			sequence = append([]int{i-1}, sequence...)
			c -= items[i-1][1]
			i--
		}
		if c == 0 {
			break
		}
	}
	return sequence
}

