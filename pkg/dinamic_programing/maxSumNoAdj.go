package dinamic_programing

func MaxSubsetSumNoAdjacent(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}
	second, first := arr[0], max(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		current := max(first, second+arr[i])
		second = first
		first = current
	}
	return first
}
