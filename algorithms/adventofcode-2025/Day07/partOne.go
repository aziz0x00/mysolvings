//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

var input string
var width int
var answer int
var visited map[int]bool

func goDown(i int) {
	if i >= len(input) || visited[i] {
		return
	}
	visited[i] = true

	if input[i] == '^' {
		answer++
		if i%width != 0 {
			goDown(i + width - 1)
		}
		if (i+1)%width != 0 {
			goDown(i + width + 1)
		}
		return
	}

	goDown(i + width)
}

func main() {
	b, _ := os.ReadFile("./input.txt")
	input = string(b)
	width = strings.Index(input, "\n") + 1
	visited = make(map[int]bool)

	sIdx := 0
	for i := range width {
		if input[i] == 'S' {
			sIdx = i
			break
		}
	}

	goDown(sIdx)

	fmt.Println(answer)
}
