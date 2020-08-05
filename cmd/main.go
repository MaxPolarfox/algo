package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/GO_algos/pkg/arrays"
	"github.com/GO_algos/pkg/dinamic_programing"
	"github.com/GO_algos/pkg/recursion"
)

func main() {
	algoPTR := flag.String("algo", "fibonacci", "algo name")

	valArrayIntPTR := flag.String("valArrayInt", "[]", "value []int")
	flag.Parse()

	algoName := *algoPTR
	valArrayIntArgument := *valArrayIntPTR
	arguments := flag.Args()

	switch algoName {
	case "fibonacci":
		n, _ := strconv.Atoi(arguments[0])
		result := recursion.GetNthFib(n)
		fmt.Printf("Result: %v", result)
	case "longest_peak":
		arr := []int{}
		for _, val := range arguments {
			valInt, _ := strconv.Atoi(val)
			arr = append(arr, valInt)
		}
		result := arrays.LongestPeak(arr)
		fmt.Printf("Result: %v", result)
	case "permutations":
		arr := []int{}
		for _, val := range arguments {
			valInt, _ := strconv.Atoi(val)
			arr = append(arr, valInt)
		}
		result := recursion.GetPermutations(arr)
		fmt.Printf("Result: %v", result)
	case "water_area":
		arr := []int{}
		for _, val := range arguments {
			valInt, _ := strconv.Atoi(val)
			arr = append(arr, valInt)
		}
		result := dinamic_programing.WaterArea(arr)
		fmt.Printf("Result: %v", result)
	case "is_monotonic":
		if valArrayIntArgument == "" {
			panic([]interface{}{`Not enough arguments, please pass -valArrayIntPTR="123456"`})
		}
		arrIntInput := StringToArrInt(valArrayIntArgument)
		result := arrays.IsMonotonic(arrIntInput)
		fmt.Printf("Result: %v", result)
	}
}

func StringToArrInt(s string) []int {
	result := []int{}
	arrString := strings.Split(s, "")
	for i := 0; i < len(arrString); i++ {
		num, err := strconv.Atoi(arrString[i])
		if err != nil {
			panic([]interface{}{"StringToArrInt.strings.Split", "err", err})
		}
		result = append(result, num)
	}
	return result
}
