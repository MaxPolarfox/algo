package main

import (
	"flag"
	"fmt"
	"github.com/algos/pkg/recursion"
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
	}
}
