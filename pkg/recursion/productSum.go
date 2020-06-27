package recursion

func ProductSum(array []interface{}) int {
	return productSumHelper(array, 1)
}

func productSumHelper(arr []interface{}, m int) int {
	sum := 0
	for _, val := range arr {
		if cast, ok := val.([]interface{}); ok {
			sum += productSumHelper(cast, m+1)
		} else if cast, ok := val.(int); ok {
			sum += cast
		} else if cast, ok := val.(float64); ok { // For purpose of testing the int comes as float64
			sum += int(cast)
		}
	}
	return sum * m
}
