package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./input.txt")

	n := strings.Index(string(b), "\n")

	grid := make([][]byte, n)

	// loads the grid into a 2D array
	for i, line := range bytes.Split(b, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		grid[i] = make([]byte, n)
		for j, char := range line {
			grid[i][j] = char
		}
	}

	answer := 0

	for i := range n {
		for j := range n {
			if grid[i][j] != byte('@') {
				continue
			}
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

					if grid[i+d1][j+d2] == byte('@') {
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

			if counter < 4 {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
