package sorting

func InsertionSort(arr []int) []int {

	for i:=1; i < len(arr); i++ {
		current := arr[i]
		for j := i-1; j >=0; j-- {
			if current < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}