package dinamic_programing

func max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}

func min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if num < curr {
			curr = num
		}
	}
	return curr
}
