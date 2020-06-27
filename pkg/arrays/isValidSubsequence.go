package arrays

func IsValidSubsequence(arr []int, sequence []int) bool {
	curr := 0
	for _, val := range arr {
		if val == sequence[curr] {
			curr++
		}
		if curr == len(sequence) {
			return true
		}
	}

	return false
}
