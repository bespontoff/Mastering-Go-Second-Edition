package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("give me numbers in parameters")
		os.Exit(1)
	}
	arguments := os.Args
	sum := 0.0
	for i := 0; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err == nil {
			sum += n
		}
	}
	fmt.Println("Sum of args =", sum)
}
