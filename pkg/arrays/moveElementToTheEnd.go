package arrays

func MoveElementToEnd(arr []int, toMove int) []int {
	i, j := 0, len(arr)-1

	for i < j {
		for i < j && arr[j] == toMove {
			j--
		}
		if arr[i] == toMove {
			arr[i], arr[j] = arr[j], arr[i]
		}
		i++
	}
	return arr
}
