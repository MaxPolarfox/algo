package arrays


func ZigzagTraverse(arr [][]int) []int {
	if len(arr) ==  0 {
		return []int{}
	}

	col, row := 0, 0
	height, width := len(arr) - 1, len(arr[0]) - 1
	result := []int{}
	goingDown := true

	for !isOutOfBound(row, col, height, width) {
		result = append(result, arr[row][col])
		if goingDown {
			if col == 0 || row == height {
				goingDown = false
				if row == height {
					col++
				} else {
					row++
				}
			} else {
				row++
				col--
			}
		} else {
			if row == 0 || col == width {
				goingDown = true
				if col == width {
					row++
				} else {
					col++
				}
			} else {
				col++
				row--
			}
		}
	}
	return result
}


func isOutOfBound(row, col, height, width int) bool {
	return row < 0 || row > height || col < 0 || col > width
}