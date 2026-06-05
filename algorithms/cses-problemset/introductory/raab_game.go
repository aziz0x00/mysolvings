//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

func Rot(p, k, n int) int {
	return (k + p) % n
}

func Solve(n, a, b int) {
	if (a == 0) != (b == 0) || a+b > n { // can't beat n, can't win with 1, and the game is in n turns.
		println("NO")
		return
	}
	println("YES")

	for i := range n {
		print(i+1, " ")
	}
	println()

	fixed := n - a - b
	for i := range fixed { // DRAWs
		print(1+i, " ")
	}
	for i := range n - fixed {
		print(1+fixed+Rot(a, i, a+b), " ")
	}
	println()
	// (n-a-b)      (a)              (b)
	// [DRAWs] [PlayerA wins] [PlayerB wins]
}

func main() {
	// b, _ := os.ReadFile("./raab_game_test/1.in")
	b, _ := os.ReadFile("./raab_game.txt")

	for _, line := range strings.Split(string(b), "\n")[1:] {
		if len(line) == 0 {
			continue
		}

		var n, a, b int
		fmt.Sscanf(line, "%d %d %d", &n, &a, &b)
		Solve(n, a, b)
	}
}
