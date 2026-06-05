package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Point struct {
	X int
	Y int
}

func Abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func Area(pair []Point) int {
	p1 := pair[0]
	p2 := pair[1]
	return (Abs(p1.X-p2.X) + 1) * (Abs(p1.Y-p2.Y) + 1)
}

var altitudes [][]int

func IsInterior(p Point) bool {
	// Check whether `p` is interior to admissible space
	// The idea is that if you draw a vertical semiline starting
	// from the point, it will hit the boundry an
	// odd number of times.

	alts := altitudes[p.X]
	// fmt.Println(p, alts)

	for i, v := range alts {
		if v >= p.Y {
			return alts[i] == p.Y || (len(alts)-i)%2 != 0
		}
	}
	return false
}

func IsAdmissible(p1, p2 Point) bool {
	// verify that the perimeter is interior to the admissible space

	maxX, minX := max(p1.X, p2.X), min(p1.X, p2.X)
	maxY, minY := max(p1.Y, p2.Y), min(p1.Y, p2.Y)

	for _, y := range []int{p1.Y, p2.Y} {
		for i := minX; i <= maxX; i++ {
			if !IsInterior(Point{i, y}) {
				// fmt.Println("Y", i, y, []int{minX, maxX})
				return false
			}
		}
	}

	for _, x := range []int{p1.X, p2.X} {
		for i := minY; i <= maxY; i++ {
			if !IsInterior(Point{x, i}) {
				// fmt.Println("X", x, i)
				return false
			}
		}
	}

	return true
}

func main() {
	b, _ := os.ReadFile("./input.txt")
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

	var pairs [][]Point

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, []Point{points[i], points[j]})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return Area(pairs[i]) > Area(pairs[j])
	})

	maxX := 0
	for _, p := range points {
		maxX = max(maxX, p.X+1)
	}
	altitudes = make([][]int, maxX)

	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		if p1.Y == p2.Y {
			for j := min(p1.X, p2.X); j < max(p1.X, p2.X); j++ {
				altitudes[j] = append(altitudes[j], p1.Y)
			}
		}
	}
	for _, alts := range altitudes {
		sort.Slice(alts, func(i, j int) bool { return alts[i] < alts[j] })
	}

	for _, p := range pairs {
		if IsAdmissible(p[0], p[1]) {
			fmt.Println(p, Area(p))
			break
		}
	}
}
