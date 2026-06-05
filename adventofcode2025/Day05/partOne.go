package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Interval struct {
	Begin int
	End   int
}

func main() {
	b, _ := os.ReadFile("./input.txt")

	var intervals []Interval

	for line := range strings.SplitSeq(string(b), "\n") {
		if len(line) == 0 {
			break
		}

		var start, end int
		fmt.Sscanf(line, "%d-%d", &start, &end)

		intervals = append(intervals, Interval{start, end})
	}

	var ingredients []int

	b = b[bytes.Index(b, []byte{'\n', '\n'})+2:]

	for line := range strings.SplitSeq(string(b), "\n") {
		if len(line) == 0 {
			break
		}
		var k int
		fmt.Sscanf(line, "%d", &k)
		ingredients = append(ingredients, k)
	}

	answer := 0

	for _, k := range ingredients {
		for _, v := range intervals {
			if k >= v.Begin && k <= v.End {
				answer++
				break
			}
		}
	}

	fmt.Println(answer)
}
