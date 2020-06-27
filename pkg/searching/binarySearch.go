package searching

func BinarySearch(arr []int, target int) int {
	return helper(arr, target, 0, len(arr)-1)
}

func helper (arr []int, target int, left int, right int) int {
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if target < arr[mid]  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}