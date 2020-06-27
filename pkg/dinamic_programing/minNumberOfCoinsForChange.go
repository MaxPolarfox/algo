package dinamic_programing

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	numOfCoins := make([]int, n+1)
	for i := range numOfCoins {
		numOfCoins[i] = math.MaxInt32
	}
	numOfCoins[0] = 0

	for _, denom := range denoms {
		for amount := range numOfCoins {
			if denom <= amount {
				curr := numOfCoins[amount - denom] + 1
				if curr < numOfCoins[amount] {
					numOfCoins[amount] = curr
				}
			}
		}
	}
	if numOfCoins[n] != math.MaxInt32 {
		return numOfCoins[n]
	}
	return -1
}