package dinamic_programing

func WaterArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	water := 0
	leftMax := make ([]int, len(heights))
	leftMax[0] = heights[0]
	for i:=1; i < len(heights); i++ {
		leftMax[i] = max(heights[i], leftMax[i-1])
	}

	rightMax := make ([]int, len(heights))
	rightMax[len(rightMax) - 1] = heights[len(heights) -1]
	for i:= len(heights)-2; i >= 0 ; i-- {
		rightMax[i] = max(heights[i], rightMax[i+1])
	}

	for i := range heights {
		water += min(rightMax[i], leftMax[i]) - heights[i]
	}
	return water
}