package arrays

// Zigzag Traverse Write a function that takes in an n x m two-dimensional array (that can be square-shaped when n === m)
// and returns a one-dimensional array of all the array's elements in zigzag order.
// Zigzag order starts at the top left corner of the two-dimensional array, goes down by one element,
// and proceeds in a zigzag pattern all the way to the bottom right corner.
// Sample Input array = [
//					[1,  3,  4, 10],
//					[2,  5,  9, 11],
//					[6,  8, 12, 15],
//					[7, 13, 14, 16]
//				]
// Sample Output expected =  [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]

func ZigzagTraverse(arr [][]int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	col, row := 0, 0
	height, width := len(arr)-1, len(arr[0])-1
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
