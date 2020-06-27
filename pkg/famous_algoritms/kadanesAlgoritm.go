package famous_algoritms

func KadanesAlgorithm(arr []int) int {
	maxEndingHere := arr[0]
	maxSoFar := arr[0]

	for i := 1; i < len(arr); i++ {
		num := arr[i]
		maxEndingHere = max(num, maxEndingHere+num)
		maxSoFar = max(maxEndingHere, maxSoFar)
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
