package sorting

func QuickSort(arr []int) []int {
	return quickSortHelper(arr, 0, len(arr) - 1)
}

func quickSortHelper(arr []int, left, right int) []int {
	if left < right {
		pivot := pivotHelper(arr, left, right)
		// Left
		quickSortHelper(arr, left, pivot - 1)
		// Right
		quickSortHelper(arr, pivot + 1, right)
	}
	return arr
}

func pivotHelper(arr []int, start, end int) int {
	pivot := start
	swapIdx := start

	for i := start + 1; i <= end; i++ {
		if arr[pivot] > arr[i] {
			swapIdx++
			swap(arr, swapIdx, i)
		}
	}
	swap(arr, pivot, swapIdx)
	return swapIdx
}