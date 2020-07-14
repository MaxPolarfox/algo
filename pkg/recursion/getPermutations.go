package recursion

func GetPermutations(arr []int) [][]int {
	perms := [][]int{}
	getPermutationsHelper(arr, []int{}, &perms)
	return perms
}

func getPermutationsHelper(arr []int, perm []int, perms *[][]int) {
	if len(arr) == 0 && len(perm) != 0 {
		*perms = append(*perms, perm)
		return
	}
	for i := range arr {
		newArr := make([]int, i)
		copy(newArr, arr[:i])
		newArr = append(newArr, arr[i+1:]...)
		newPerm := make([]int, len(perm))
		newPerm = append(newPerm, arr[i])
		copy(newPerm, perm)
		getPermutationsHelper(newArr, newPerm, perms)
	}
}
