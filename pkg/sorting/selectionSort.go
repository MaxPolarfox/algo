package sorting

func SelectionSort(arr []int) []int {

	for i, _ := range arr {
		min := i
		for j := i+1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}