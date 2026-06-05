package main

import (
	"fmt"
	"os"
	"strings"
)

var graph map[string][]string
var answer int

var mem map[string]map[int]int

func dfs(node string, ct int) {
	if mem[node] != nil && mem[node][ct] != 0 {
		answer += mem[node][ct] - 1
		return
	}

	if node == "out" {
		if ct >= 2 {
			answer++
		}
		return
	}

	for _, c := range graph[node] {
		i := 0
		if c == "dac" {
			i++
		}
		if c == "fft" {
			i++
		}
		old := answer
		dfs(c, ct+i)
		if mem[c] == nil {
			mem[c] = make(map[int]int)
		}
		mem[c][ct+i] = answer - old + 1
	}
	return
}

func main() {
	b, _ := os.ReadFile("./input.txt")
	graph = make(map[string][]string)
	mem = make(map[string]map[int]int)

	for line := range strings.SplitSeq(string(b), "\n") {
		if len(line) == 0 {
			continue
		}

		i := strings.Index(line, ":")
		node := line[:i]
		for child := range strings.SplitSeq(line[i+2:], " ") {
			graph[node] = append(graph[node], child)
		}
	}

	dfs("svr", 0)
	fmt.Println(answer)
}
