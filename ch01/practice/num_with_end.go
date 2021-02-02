package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		println("give me more arguments")
		os.Exit(1)
	}
	arguments := os.Args
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "END" {
			os.Exit(0)
		}
		n, err := strconv.ParseInt(arguments[i], 10, 64)
		if err == nil {
			println(n)
		}
	}
}
