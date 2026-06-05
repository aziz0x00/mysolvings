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

		n := 0
		nIdx := 0
		m := 0

		for i := range bank[:len(bank)-1] {
			d := int(bank[i]) - '0'

			if n < d {
				n = d
				nIdx = i
			}
		}

		for _, bat := range bank[nIdx+1:] {
			m = max(int(bat)-'0', m)
		}

		answer += n*10 + m
	}
	fmt.Println(answer)
}
