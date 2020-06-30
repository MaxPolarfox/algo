package main

import (
	"flag"
	"fmt"
	"github.com/algos/GO_algos/pkg/arrays"
	"github.com/algos/GO_algos/pkg/recursion"
	"strconv"
)

func main() {
	algoPTR := flag.String("algo", "fibonacci", "algo name")
	flag.Parse()

	algoName := *algoPTR
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
	}
}
