package main

import (
	"fmt"
	"os"
	"strings"
)

var graph map[string][]string
var answer int

func dfs(node string) {
	if node == "out" {
		answer++
		return
	}
	for _, c := range graph[node] {
		dfs(c)
	}
}

func main() {
	b, _ := os.ReadFile("./input.txt")
	graph = make(map[string][]string)

	for line := range strings.SplitSeq(string(b), "\n") {
		if len(line) == 0 {
			continue
		}

		i := strings.Index(line, ":")
		node := line[:i]
		fmt.Println(node)
		for child := range strings.SplitSeq(line[i+2:], " ") {
			graph[node] = append(graph[node], child)
		}
	}

	dfs("you")
	fmt.Println(answer)
}
