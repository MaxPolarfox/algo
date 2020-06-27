package recursion

func GetNthFib(n int) int {
	return helper(n, map[int]int{
		1: 0,
		2: 1,
	})
}

func helper(n int, memo map[int]int) int {
	if val, found := memo[n]; found {
		return val
	}
	memo[n] = helper(n-1, memo) + helper(n-2, memo)
	return memo[n]
}
