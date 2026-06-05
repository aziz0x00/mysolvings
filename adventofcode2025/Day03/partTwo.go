package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./input.txt")

	answer := 0

	for bank := range strings.SplitSeq(string(b), "\n") {
		if len(bank) == 0 {
			continue
		}

		var joltage []int
		size := len(bank)

		lastIdx := 0
		for i := range 12 {
			m := 0
			for j := lastIdx; j < size-(12-i-1); j++ {
				d := int(bank[j]) - '0'
				if d > m {
					m = d
					lastIdx = j + 1
				}
			}
			joltage = append(joltage, m)
		}

		acc := 0
		for _, v := range joltage {
			acc = 10*acc + v
		}

		answer += acc
	}

	fmt.Println(answer)
}
