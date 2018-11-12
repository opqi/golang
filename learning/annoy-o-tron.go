package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1]
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("Hello!\t")
	}
	fmt.Println()
}
