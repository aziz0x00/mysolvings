package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidID(id int) bool {
	num := strconv.Itoa(id)

	return num[:len(num)/2] == num[len(num)/2:]
}

func main() {

	b, _ := os.ReadFile("./input.txt")

	answer := 0

	for rng := range strings.SplitSeq(string(b), ",") {
		var start, end int
		fmt.Sscanf(rng, "%d-%d", &start, &end)

		for i := start; i <= end; i++ {
			if isInvalidID(i) {
				answer += i
			}
		}
	}
	fmt.Println(answer)
}
