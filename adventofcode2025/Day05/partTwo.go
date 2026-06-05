package main

import (
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

	// bubble sort ??
	i := 0
	updated := false
	for {
		if i+1 >= len(intervals) { // new turn
			if !updated {
				break
			}
			updated = false
			i = 0
			continue
		}

		if intervals[i].Begin > intervals[i+1].Begin {
			intervals[i], intervals[i+1] = intervals[i+1], intervals[i]
			updated = true
		}

		i++
	}

	answer := 0

	lb := intervals[0].Begin
	hb := intervals[0].End

	for i := 1; i < len(intervals)-1; i++ {
		if hb >= intervals[i].Begin {
			hb = max(hb, intervals[i].End)
			continue
		}
		answer += hb - lb + 1
		lb = intervals[i].Begin
		hb = intervals[i].End
	}

	i = len(intervals) - 1

	if hb >= intervals[i].Begin {
		hb = max(hb, intervals[i].End)
		answer += hb - lb + 1
	} else {
		// answer += hb - lb + 1
		answer += intervals[i].End - intervals[i].Begin + 1
	}

	fmt.Println(answer)
}
