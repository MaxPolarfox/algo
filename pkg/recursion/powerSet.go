package recursion

func Powerset(arr []int) [][]int {
	subsets := [][]int{{}}
	for _, val := range arr {
		length := len(subsets)
		for i := 0; i < length; i++ {
			currentSubset := subsets[i]
			newSubset := append([]int{}, currentSubset...)
			newSubset = append(newSubset, val)
			subsets = append(subsets, newSubset)
		}
	}
	return subsets
}
