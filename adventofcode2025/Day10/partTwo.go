package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Button []int

type Manual struct {
	Buttons      []Button
	Requirements []int
}

func ApplyButtons(buttons []Button) int {
	res := 0
	for _, b := range buttons {
		for _, v := range b {
			res ^= 1 << v
		}
	}
	return res
}

func FindCombination(elements []Button, k, target int) []Button {
	n := len(elements)

	var Search func(depth, maxDepth int, selected []Button) []Button
	Search = func(depth, maxDepth int, selected []Button) []Button {

		if maxDepth == 0 {
			if ApplyButtons(selected) == target {
				return selected
			} else {
				return nil
			}
		}
		if depth == n {
			return nil
		}
		b := Search(depth+1, maxDepth, selected)
		if b != nil {
			return b
		}
		return Search(depth+1, maxDepth-1, append(selected, elements[depth]))
	}

	return Search(0, k, []Button{})
}

func part1Solver(elements []Button, target int) []Button {

	for i := 1; i < len(elements); i++ {
		selectedBts := FindCombination(elements, i, target)
		if selectedBts != nil {
			return selectedBts
		}
	}
	return nil
}

func solve(m Manual) int {

	ans := 0

	reqs := slices.Clone(m.Requirements)

	// var allBts []Button

	for e := 0; ; e++ {

		target := 0
		for i, v := range reqs {
			target |= (v & 1) << i
		}

		if target == 0 {
			break
		}

		bts := part1Solver(m.Buttons, target)
		ans += len(bts) << e

		ctr := make(map[int]int)
		for _, b := range bts {
			for _, i := range b {
				ctr[i]++
			}
		}

		for i := range reqs {
			reqs[i] -= ctr[i]
			reqs[i] >>= 1
		}

		fmt.Println(bts, m.Requirements, reqs)

		// for range 1 << e {

		// 	allBts = append(allBts, bts...)
		// }
	}

	// testR := slices.Clone(m.Requirements)

	// fmt.Println(len(allBts))
	// for _, b := range allBts {
	// 	for _, i := range b {
	// 		testR[i]--
	// 	}
	// }
	// fmt.Println(testR, len(allBts))

	return ans
}

func main() {
	b, _ := os.ReadFile("./input-test.txt")
	input := string(b)

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
		// answer += solve(man)
		fmt.Println(solve(man))
		// break
	}

	fmt.Println(answer)
}
