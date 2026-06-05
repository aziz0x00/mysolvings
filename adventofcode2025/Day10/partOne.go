package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Button []int

type Manual struct {
	Lights  int
	Buttons []Button
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

func CombinationFound(elements []Button, k, target int) bool {
	n := len(elements)

	var Search func(depth, maxDepth int, selected []Button) bool
	Search = func(depth, maxDepth int, selected []Button) bool {

		if maxDepth == 0 {
			return ApplyButtons(selected) == target
		}
		if depth == n {
			return false
		}
		return Search(depth+1, maxDepth, selected) ||
			Search(depth+1, maxDepth-1, append(selected, elements[depth]))
	}

	return Search(0, k, []Button{})
}

func main() {
	b, _ := os.ReadFile("./input.txt")
	input := string(b)

	var manuals []Manual

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			break
		}
		line = strings.ReplaceAll(line, "[", "")
		line = strings.ReplaceAll(line, "]", "")

		lights := 0
		var buttons []Button
		i := 0
		for ; line[i] != ' '; i++ {
			if line[i] == '#' {
				lights |= 1 << i
			}
		}

		for ; line[i] != '{'; i++ {
			if line[i-1] != '(' {
				continue
			}
			var b Button
			end := strings.Index(line[i:], ")")
			for v := range strings.SplitSeq(line[i:][:end], ",") {
				vInt, _ := strconv.Atoi(v)
				b = append(b, vInt)
			}
			buttons = append(buttons, b)
			i += end + 1
		}

		manuals = append(manuals, Manual{lights, buttons})
	}

	answer := 0

	for _, man := range manuals {
		minK := 1
		for ; minK < len(man.Buttons); minK++ {
			// toutes les `minK` parmi `len(man.Buttons)`
			if CombinationFound(man.Buttons, minK, man.Lights) {
				break
			}
		}
		answer += minK
	}

	fmt.Println(answer)
}
