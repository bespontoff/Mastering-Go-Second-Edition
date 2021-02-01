package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("give me numbers in parameters")
		os.Exit(1)
	}
	arguments := os.Args
	fmt.Println(arguments)
}
