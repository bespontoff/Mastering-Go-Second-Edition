package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("give me more arguments")
		os.Exit(1)
	}
	arguments := os.Args
	count := 0
	sum := 0.0

	for i := 0; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += n
			count++
		}
	}
	fmt.Println("AVG of arguments =", sum/float64(count))
}
