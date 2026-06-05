package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func divisors(n int) []int {
	var divisors []int

	for i := range n - 1 {
		if n%(i+1) == 0 {
			divisors = append(divisors, i+1)
		}
	}

	return divisors
}

func isInvalidID(id int) bool {
	idStr := strconv.Itoa(id)

	n := len(idStr)

	divs := divisors(n)

	for _, div := range divs {
		ok := true

		for i := 0; i < n-div; i += div {
			if idStr[i:i+div] != idStr[i+div:i+2*div] {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}

	return false
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
