package arrays

func SpiralTraverse(arr [][]int) []int {
	result := []int{}
	spiralFill(arr, 0, len(arr)-1, 0, len(arr[0])-1, &result)
	return result
}

func spiralFill(arr [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
	if startRow > endRow || startCol > endCol {
		return
	}

	// Top
	for col := startCol; col <= endCol; col++ {
		*result = append(*result, arr[startRow][col])
	}

	// Right
	for row := startRow + 1; row <= endRow; row++ {
		*result = append(*result, arr[row][endCol])
	}

	// bottom
	for col := endCol - 1; col >= startCol; col-- {
		if startRow == endRow {
			break
		}
		*result = append(*result, arr[endRow][col])
	}

	// left
	for row := endRow - 1; row >= startRow+1; row-- {
		if startCol == endCol {
			break
		}
		*result = append(*result, arr[row][startCol])
	}

	spiralFill(arr, startRow+1, endRow-1, startCol+1, endCol-1, result)
}
