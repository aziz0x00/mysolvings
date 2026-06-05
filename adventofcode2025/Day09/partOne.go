package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	b, _ := os.ReadFile("./input-test.txt")
	input := string(b)

	var points []Point

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var p Point
		fmt.Sscanf(line, "%d,%d", &p.X, &p.Y)
		points = append(points, p)
	}

	Abs := func(i int) int {
		if i < 0 {
			i = -i
		}
		return i
	}

	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			maxArea = max(maxArea, Abs(p1.X-p2.X+1)*Abs(p1.Y-p2.Y+1))
		}
	}

	fmt.Println(maxArea)
}
