//go:build ignore
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Sum(a []int) (s int) {
	for _, i := range a {
		s += i
	}
	return
}

type Configuration struct {
	DimX   int
	DimY   int
	Counts []int
}

type State [][]int
type Shape int

func (shape Shape) String() string {
	out := ""
	for i := range 9 {
		c := "."
		if (shape>>i)&1 == 1 {
			c = "#"
		}
		out += c
		if i%3 == 2 {
			out += "\n"
		}
	}
	out += "\n"
	// fmt.Println()
	return out
}

func (state State) String() string {
	dimY := len(state)
	dimX := len(state[0])
	out := ""
	for i := range dimY * dimX {
		c := "."
		if state[i%dimY][i/dimY] != 0 {
			c = strconv.Itoa(state[i%dimY][i/dimY])
		}
		out += c
		if (i+1)%dimY == 0 {
			out += "\n"
		}
	}
	out += "\n"
	return out
}

func CanFit(y, x int, shape Shape, state State) bool {
	// fmt.Println(n, m)
	// fmt.Println(state)
	// fmt.Println(y, x)
	for i := range 9 {
		if y+(i%3) >= len(state) || x+(i/3) >= len(state[0]) {
			return false
		}
		if ((shape>>i)&1 == 1) && state[y+(i%3)][x+(i/3)] != 0 {
			return false
		}
	}
	return true
}

func UnFit(y, x int, shape Shape, state State) {
	for i := range 9 {
		if (shape>>i)&1 == 1 {
			state[y+(i%3)][x+(i/3)] = 0
		}
	}
}

func Fit(y, x int, shape Shape, state State) {
	// fmt.Println(len(state), y, x)
	m := 0
	for i := range len(state) {
		for j := range len(state[0]) {
			m = max(m, state[i][j])
		}
	}

	for i := range 9 {
		if (shape>>i)&1 == 1 {
			state[y+(i%3)][x+(i/3)] = m + 1
		}
	}
}


func GetTransforms(shape Shape) (shapes []Shape) {
	shapes = append(shapes, shape)

	// make all rots with flip (rots, flips)
	for kind := range 7 {
		var transform Shape
		for j := range 3 {
			for i := range 3 {
				b := (shape >> (i + 3*j)) & 1

				coords := [][]int{
					{2 - j, i},     // 1 rot
					{2 - i, 2 - j}, // 2 rot
					{j, 2 - i},     // 3 rot

					{2 - j, 2 - i}, // 1 rot + flip
					{i, 2 - j},     // 2 rot + flip
					{j, i},         // 3 rot + flip

					{2 - i, j}, // flip
				}[kind]
				transform |= b << (coords[0] + 3*coords[1])
			}
		}
		shapes = append(shapes, transform)
	}
	return
}

var visited map[string]bool

func Possible(shapes []Shape, conf Configuration, used []int, state State) bool {
	done := true
	for i := range len(used) {
		if used[i] != conf.Counts[i] {
			done = false
			break
		}
	}
	if done {
		return true
	}
	if visited[state.String()] {
		return false
	}

	for i, shape := range shapes {
		if conf.Counts[i] == used[i] {
			continue
		}
		if Sum(conf.Counts)-Sum(used) == 1 {
			fmt.Println(used)
			fmt.Println(state)
		}

		for _, transf := range GetTransforms(shape) {
			for y := range conf.DimY - 2 {
				for x := range y + 1 {
					used[i]++
					if CanFit(y-x, x, transf, state) {
						Fit(y-x, x, transf, state)
						fmt.Println(state)
						if !visited[state.String()] && Possible(shapes, conf, used, state) {
							return true
						}
						visited[state.String()] = true
						UnFit(y-x, x, transf, state)
					}
					used[i]--
				}
			}
		}

	}
	return false
}

func main() {
	b, _ := os.ReadFile("./input-test.txt")
	input := string(b)

	var shapes []Shape

	i := 0
	for ; ; i++ {
		if input[i] != '#' && input[i] != '.' {
			if i+4*3 > strings.Index(input, "x") { // cpu abuse
				break
			}
			continue
		}
		var shape Shape
		for i, c := range strings.ReplaceAll(input[i:i+4*3], "\n", "") {
			if c == '#' {
				shape |= 1 << i
			}
		}
		shapes = append(shapes, shape)
		// fmt.Println(strconv.FormatInt(int64(shape), 2), input[i:i+4*3])
		i += 4 * 3
	}

	var confs []Configuration

	for line := range strings.SplitSeq(input[i:], "\n") {
		// fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		idx := strings.Index(line, ":")
		conf := Configuration{}
		fmt.Sscanf(line[:idx], "%dx%d", &conf.DimY, &conf.DimX)
		for c := range strings.SplitSeq(line[idx+2:], " ") {
			i, _ := strconv.Atoi(c)
			conf.Counts = append(conf.Counts, i)
		}
		confs = append(confs, conf)
	}

	// answer := 0
	conf := confs[1]
	fmt.Println(conf)
	var state [][]int
	for range conf.DimY {
		state = append(state, make([]int, conf.DimX))
	}
	used := make([]int, len(conf.Counts))
	visited = make(map[string]bool)
	if Possible(shapes, conf, used, state) {
		// answer++
		fmt.Println(conf.Counts)
	}
	// break
	// }
	answer := 0

	for _, conf := range confs {
		sumArea := 0
		for i, count := range conf.Counts {
			if count == 0 {
				continue
			}
			for j := range 9 {
				if (shapes[i]>>j)&1 == 1 {
					sumArea += count
				}
			}
		}
		fmt.Println(sumArea, conf.DimX*conf.DimY)
		if sumArea < conf.DimX*conf.DimY {
			answer++
		}
	}

	fmt.Println(answer)
}
