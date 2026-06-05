package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./input.txt")

	n := strings.Count(string(b), "\n")

	var ops []byte
	worksheet := make([][]int, n-1)

	for i, line := range strings.Split(string(b), "\n") {
		for item := range strings.SplitSeq(line, " ") {
			if len(item) == 0 {
				continue
			}

			if i == n-1 { // operators
				ops = append(ops, item[0])
				continue
			}

			num, _ := strconv.Atoi(item)
			worksheet[i] = append(worksheet[i], num)
		}
	}

	answer := 0

	unitMap := map[byte]int{'+': 0, '*': 1}
	opMap := map[byte](func(a, b int) int){
		'+': func(a, b int) int {
			return a + b
		},
		'*': func(a, b int) int {
			return a * b
		},
	}

	for i := range len(ops) {
		res := unitMap[ops[i]]
		op := opMap[ops[i]]
		for j := range n - 1 {
			res = op(res, worksheet[j][i])
		}
		answer += res
	}

	fmt.Println(answer)
}
