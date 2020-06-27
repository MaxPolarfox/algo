package arrays

func IsMonotonic(arr []int) bool {
	isNonDecrising := true
	isNonIncreasing := true

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			isNonDecrising = false
		}
		if arr[i-1] < arr[i] {
			isNonIncreasing = false
		}
	}

	return isNonDecrising || isNonIncreasing
}
