package sorting

func BubbleSort(arr []int) []int {
	sorted := false
	counter := 0
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1-counter; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				sorted = false
			}
		}
		counter++
	}
	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
