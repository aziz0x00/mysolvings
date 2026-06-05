package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluateWorksheet(worksheet [][]int, ops []byte) int {

	ans := 0

	unitMap := map[byte]int{'+': 0, '*': 1}
	opMap := map[byte](func(a, b int) int){
		'+': func(a, b int) int {
			return a + b
		},
		'*': func(a, b int) int {
			return a * b
		},
	}

	for i := range worksheet {
		res := unitMap[ops[i]]
		op := opMap[ops[i]]
		for j := range worksheet[i] {
			res = op(res, worksheet[i][j])
		}
		ans += res
	}

	return ans
}

func main() {
	b, _ := os.ReadFile("./input.txt")
	input := string(b)

	n := strings.Count(input, "\n")
	m := strings.Index(input, "\n") + 1

	// operators
	var ops []byte
	for item := range strings.SplitSeq(input[m*(n-1):m*n-1], " ") {
		if len(item) == 0 {
			continue
		}
		ops = append(ops, item[0])
	}

	var worksheet [][]int
	worksheet = append(worksheet, []int{})
	idx := 0
	for i := 0; i < m-1; i++ {

		digits := ""
		for j := range n - 1 {
			if input[i+m*j] == ' ' {
				continue
			}
			digits += string(input[i+m*j])
		}
		if digits == "" {
			worksheet = append(worksheet, []int{})
			idx++
		} else {
			num, _ := strconv.Atoi(digits)
			worksheet[idx] = append(worksheet[idx], num)
		}
	}

	answer := evaluateWorksheet(worksheet, ops)
	fmt.Println(answer)
}
