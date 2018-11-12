package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		cntLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "doppelganger: %v\n", err)
				continue
			}
			cntLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func cntLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	if input.Err() != nil {
		fmt.Println(input.Err())
		os.Exit(2)
	}

	for input.Scan() {
		for _, args := range strings.Fields(input.Text()) {
			counts[args]++
		}
	}
}
