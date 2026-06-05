//go:build ignore
package main

import (
	"fmt"
	"os"
	"strings"
)

var input string
var width int
var memo map[int]int

func goDown(i int, path string) {
	if i >= len(input) {
		memo[i] = 1
		return
	}
	if memo[i] != 0 {
		return
	}

	if input[i] == '^' {
		if i%width != 0 {
			goDown(i+width-1, path+"L")
		}
		if (i+1)%width != 0 {
			goDown(i+width+1, path+"R")
		}
		memo[i] = memo[i+width+1] + memo[i+width-1]
		return
	}

	goDown(i+width, path)
	memo[i] = memo[i+width]
}

func main() {
	var inputFile string
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	} else {
		inputFile = "./input.txt"
	}
	b, _ := os.ReadFile(inputFile)
	input = string(b)
	width = strings.Index(input, "\n") + 1
	memo = make(map[int]int)

	sIdx := 0
	for i := range width {
		if input[i] == 'S' {
			sIdx = i
			break
		}
	}

	goDown(sIdx, "")

	fmt.Println(memo[sIdx])
}
