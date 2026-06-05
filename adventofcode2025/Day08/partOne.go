package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
}

type Circuit []Point

func Sq(i int) int { // seriously go??!!
	return i * i
}

func NormSq(p []Point) int {
	return Sq(p[0].X-p[1].X) + Sq(p[0].Y-p[1].Y) + Sq(p[0].Z-p[1].Z)
}

func Cardinal(p Circuit) int {
	return len(p)
}

func MergeSort[K any](elements []K, key func(K) int) []K { // side quest !!
	n := len(elements)

	if n <= 1 {
		return elements
	}

	left := elements[:n/2]
	right := elements[n/2:]

	left = MergeSort(left, key)
	right = MergeSort(right, key)

	var i, j int

	var merged []K

	for i < len(left) && j < len(right) {

		if key(left[i])-key(right[j]) <= 0 {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

func main() {
	problem := map[string]struct {
		MaxConnections int
		Filename       string
	}{
		"test": {
			MaxConnections: 10,
			Filename:       "input-test.txt",
		},
		"chall": {
			MaxConnections: 1000,
			Filename:       "input.txt",
		},
	}["chall"]

	b, _ := os.ReadFile(problem.Filename)
	input := string(b)
	var points []Point

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var p Point
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		points = append(points, p)
	}

	var pairs [][]Point

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, []Point{points[i], points[j]})
		}
	}

	// PleaseSubscribe -- hommage to ippsec
	pairs = MergeSort(pairs, NormSq)

	memo := make(map[Point]int) // value is index+1 (leave 0 for unassigned)

	var circuits []Circuit

	count := 0
	for _, pair := range pairs {
		if count >= problem.MaxConnections {
			break
		}

		// loop over pair elements
		// and add `pair` to circuit if one its items is already took
		for i := range 2 { // for [i, j] in [[0, 1], [1, 0]]
			j := (i + 1) % 2

			if memo[pair[i]] != 0 {
				idx := memo[pair[i]] - 1

				// both already in circuits or not
				if memo[pair[j]] != 0 {
					idx2 := memo[pair[j]] - 1
					// fmt.Println("INFO [both] idx:", idx, idx2, "current:", pair)
					if idx != idx2 {
						circuits[idx] = append(circuits[idx], circuits[idx2]...)
						// fmt.Println("UPDATE", idx, circuits[idx])
						count++
						for _, p := range circuits[idx2] {
							memo[p] = idx + 1
						}
						// fmt.Println("DELETE", idx2, circuits[idx2])
						circuits[idx2] = nil

					} else {
						count++
					}

				} else {
					// fmt.Println("INFO [only1] idx:", idx, "current:", pair)
					// fmt.Println("ADD", j, pair)
					circuits[idx] = append(circuits[idx], pair[j])
					count++
					memo[pair[j]] = idx + 1
				}
				// fmt.Println("RES", circuits)
				break
			}
		}

		if memo[pair[0]] == 0 {
			circuits = append(circuits, pair)
			memo[pair[0]] = len(circuits)
			memo[pair[1]] = len(circuits)
			count++
		}
		// fmt.Println(">>>COUNT>>>", count, circuits)
	}

	circuits = MergeSort(circuits, Cardinal)

	answer := 1
	found := 0
	for i := len(circuits) - 1; i >= 0 && found < 3; i-- {
		if len(circuits[i]) != 0 {
			found++
			answer *= len(circuits[i])
		}
	}

	fmt.Println(answer)
}
