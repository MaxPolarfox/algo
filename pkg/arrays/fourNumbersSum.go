package arrays

func FourNumberSum(arr []int, target int) [][]int {
	memo := map[int][][]int{}
	quadruplets := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			diff := target - sum
			if pairs, found := memo[diff]; found {
				for _, pair := range pairs {
					newquad := append(pair, arr[i], arr[j])
					quadruplets = append(quadruplets, newquad)
				}
			}
		}
		for k := 0; k < i; k++ {
			sum := arr[i] + arr[k]
			memo[sum] = append(memo[sum], []int{arr[i], arr[k]})
		}
	}
	return quadruplets
}
