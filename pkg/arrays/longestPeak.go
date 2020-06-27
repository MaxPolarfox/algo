package arrays

func LongestPeak(arr []int) int {
	longest := 0
	i := 1

	for i < len(arr)-1 {
		isPeak := arr[i] > arr[i-1] && arr[i] > arr[i+1]
		if !isPeak {
			i++
			continue
		}

		left := i - 2
		for left >= 0 && arr[left] < arr[left+1] {
			left--
		}

		right := i + 2
		for right < len(arr) && arr[right] < arr[right-1] {
			right++
		}

		current := right - left - 1
		if current > longest {
			longest = current
		}
		i = right
	}
	return longest
}
