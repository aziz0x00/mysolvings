package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return max(b, -b)
}

func Sum(a []int) int {
	s := 0
	for _, i := range a {
		s += i
	}
	return s
}

type Matrix [][]int // []Row

func (mat Matrix) Ref() (pivots []int, freeVars []int) {
	// Calculate Row Echelon Form
	n := len(mat)
	m := len(mat[0])

	pivot := -1
	for pIdx := 0; pIdx < n; pIdx++ {
		// get pivot
		for c := pivot + 1; c < m-1; c++ {
			if mat[pIdx][c] != 0 {
				pivot = c
				break
			}
			for row := pIdx + 1; row < n; row++ {
				if mat[row][c] != 0 {
					pivot = c
					mat[row], mat[pIdx] = mat[pIdx], mat[row]
					break
				}
			}
			if mat[pIdx][c] == 0 {
				freeVars = append(freeVars, c)
			} else {
				break
			}
		}
		if len(freeVars)+len(pivots) == m-1 {
			break
		}

		for row := pIdx + 1; row < n; row++ {
			c2 := mat[row][pivot]
			c1 := mat[pIdx][pivot]
			c1, c2 = c1/Gcd(c1, c2), c2/Gcd(c1, c2)
			for col := pivot; col < m; col++ {
				mat[row][col] = c1*mat[row][col] - c2*mat[pIdx][col]
			}
		}
		pivots = append(pivots, pivot)
	}

	if len(freeVars)+len(pivots) != m-1 {
		for i := pivot + 1; i < m-1; i++ {
			freeVars = append(freeVars, i)
		}
	}
	return
}

func getSolution(freeVals []int, mat Matrix, freeVars, pivots []int) (sol []int, ok bool) {
	m := len(mat[0])
	sol = make([]int, m-1)
	for i, j := range freeVars {
		sol[j] = freeVals[i]
	}

	for row := len(pivots) - 1; row >= 0; row-- {
		rhs := mat[row][m-1]
		for col := pivots[row] + 1; col < m-1; col++ {
			rhs -= mat[row][col] * sol[col]
		}
		if rhs%mat[row][pivots[row]] != 0 {
			return
		}
		sol[pivots[row]] = rhs / mat[row][pivots[row]]
		if sol[pivots[row]] < 0 { // allow only non-negative integers sols
			return
		}
	}
	ok = true
	return
}

var ZigZagCache map[int]map[int][][]int // (n, level) -> diag

func ZigZag(n, level int) (diag [][]int) {

	if ZigZagCache[n] != nil && ZigZagCache[n][level] != nil {
		return ZigZagCache[n][level]
	}

	if n == 1 {
		diag = append(diag, []int{level})
		return
	}
	if level == 0 {
		diag = append(diag, make([]int, n))
		return
	}

	for i := range level + 1 {
		for _, j := range ZigZag(n-1, level-i) {
			diag = append(diag, append(j, i))
		}
	}

	if ZigZagCache[n] == nil {
		ZigZagCache[n] = make(map[int][][]int)
	}
	ZigZagCache[n][level] = diag
	return
}

func solve(man Manual) int {
	n := len(man.Requirements) // rows
	m := len(man.Buttons)      // cols

	var mat Matrix // Augmented matrix, so n x (m+1)

	for range n {
		mat = append(mat, make([]int, m+1))
	}

	for col, b := range man.Buttons {
		for _, row := range b {
			mat[row][col] = 1
		}
	}
	// add RHS
	for row, rhs := range man.Requirements {
		mat[row][m] = rhs
	}

	var cmat Matrix
	for _, r := range mat {
		cmat = append(cmat, slices.Clone(r))
	}

	pivots, freeVars := mat.Ref()

	if len(freeVars) == 0 {
		sol, ok := getSolution(nil, mat, freeVars, pivots)
		if !ok {
			panic("!!")
		}
		return Sum(sol)
	}

	best := 99999999
	for k := 0; k < 200; k++ {
		for _, freeVals := range ZigZag(len(freeVars), k) {
			sol, ok := getSolution(freeVals, mat, freeVars, pivots)
			if !ok {
				continue
			}
			s := 0
			for _, i := range sol {
				s += i
			}
			best = min(best, s)
		}
	}

	return best
}

type Button []int

type Manual struct {
	Buttons      []Button
	Requirements []int
}

func main() {
	b, _ := os.ReadFile("./input.txt")
	input := string(b)

	ZigZagCache = make(map[int]map[int][][]int)

	var manuals []Manual

	// parse input
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			break
		}
		line = line[strings.Index(line, "("):]
		i := 1

		// parse buttons into []Button
		var buttons []Button
		for ; line[i] != '{'; i++ {
			if line[i-1] != '(' {
				continue
			}
			end := strings.Index(line[i:], ")")

			var button Button
			for v := range strings.SplitSeq(line[i:][:end], ",") {
				vInt, _ := strconv.Atoi(v)
				button = append(button, vInt)
			}
			buttons = append(buttons, button)

			i += end + 1
		}

		// parse requirements
		var requirements []int
		i++
		end := strings.Index(line[i:], "}")
		for v := range strings.SplitSeq(line[i:][:end], ",") {
			vInt, _ := strconv.Atoi(v)
			requirements = append(requirements, vInt)
		}

		manuals = append(manuals, Manual{buttons, requirements})
	}

	answer := 0

	for _, man := range manuals {
		// fmt.Println(i, answer)
		answer += solve(man)
	}

	fmt.Println(answer)
}
