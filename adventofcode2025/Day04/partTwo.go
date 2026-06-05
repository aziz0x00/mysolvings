package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func removable(n int, grid [][]int, iternation int, i int, j int) bool {

	counter := 0

	// iternating with direction [d1, d2]
	// i.e. [i, j] + [d1, d2]
	for d1 := -1; d1 <= 1; d1++ {
		for d2 := -1; d2 <= 1; d2++ {

			if d1 == 0 && d2 == 0 {
				continue
			}

			if d1 < 0 && i == 0 || d1 > 0 && i == n-1 {
				continue
			}
			if d2 < 0 && j == 0 || d2 > 0 && j == n-1 {
				continue
			}

			if grid[i+d1][j+d2] == '@' || grid[i+d1][j+d2] == iternation {
				counter++
				if counter >= 4 {
					break
				}
			}
		}
		if counter >= 4 {
			break
		}
	}

	return counter < 4
}

func main() {
	b, _ := os.ReadFile("./input.txt")

	n := strings.Index(string(b), "\n")

	grid := make([][]int, n)

	// loads the grid into a 2D array
	for i, line := range bytes.Split(b, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		grid[i] = make([]int, n)
		for j, char := range line {
			grid[i][j] = int(char)
		}
	}

	answer := 0

	iteration := 1000 // starting from a high value to avoid conflict with ASCII values
	for {
		iteration++

		over := true

		for i := range n {
			for j := range n {
				if grid[i][j] != '@' {
					continue
				}

				if removable(n, grid, iteration, i, j) {
					answer++
					grid[i][j] = iteration
					over = false
				}
			}
		}

		if over {
			break
		}
	}

	fmt.Println(answer)
}
