package dinamic_programing

// You're given a non-empty array of positive integers where each integer represents the maximum number of steps you can take forward in the array.
// For example, if the element at index 1 is 3, you can go from index 1 to index  2, 3, or 4.  Write a function that returns
// the minimum number of jumps needed to reach the final index.
// Note that jumping from index  i to index i + x always constitutes one jump, no matter how large x is.
//
// Sample Input: array = [3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3]
// Sample Output 4 = 2 -> (4 or 2) -> (2 or 3) -> 7 -> 3

// Time: O(n) Space: O(1)
func MinNumberOfJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	jumps := 0
	maxReach := arr[0]
	steps := arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if i+arr[i] > maxReach {
			maxReach = i + arr[i]
		}
		steps--
		if steps == 0 {
			jumps++
			steps = maxReach - i
		}
	}
	return jumps + 1
}
