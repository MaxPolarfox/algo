package graphs

func HasSingleCycle(arr []int) bool {
	numVisited := 0
	curr := 0
	for numVisited < len(arr) {
		if numVisited > 0 && curr == 0 {
			return false
		}
		numVisited++
		curr = jumpToNext(arr, curr)
	}
	return curr == 0
}

func jumpToNext(arr []int, i int) int {
	num := arr[i]

	next := (i + num) % len(arr)
	if next >= 0 {
		return next
	}
	return next + len(arr)
}
