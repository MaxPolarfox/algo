package arrays

func LargestRange(arr []int) []int {
	best := []int{}
	longestLength := 0
	memo := map[int]bool{}
	for _, val := range arr {
		memo[val] = true
	}
	for _, num := range arr {
		if !memo[num] {
			continue
		}
		memo[num] = false
		currentLength, left, right := 1, num-1, num+1
		for memo[left] {
			memo[left] = false
			currentLength++
			left--
		}
		for memo[right] {
			memo[right] = false
			currentLength++
			right++
		}
		if currentLength > longestLength {
			longestLength = currentLength
			best = []int{left + 1, right - 1}
		}
	}
	return best
}
