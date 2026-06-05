package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
 * This is a slow solution, but it works for the sample input.
 *  It uses a recursive backtracking approach to try all possible distributions
 *  of button presses to satisfy the requirements.
 */

func MergeSort[K any](elements []K, key func(K) int) []K { // reused this implemented in day9, because golang sort library SUCKKKKSSS!!!!!!!
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

func Sum(a []int) (s int) {
	for _, k := range a {
		s += k
	}
	return
}

type Button []int

type Manual struct {
	Buttons      []Button
	Requirements []int
}

// to go through Requirements from least frequent index to most
func (m *Manual) reqOrder() (ret []int) {
	ctr := make(map[int]int)
	for _, b := range m.Buttons {
		for _, i := range b {
			ctr[i]++
		}
	}

	for i := range len(m.Requirements) {
		ret = append(ret, i)
	}
	// fmt.Println("ret>", ret)
	// sort.Slice(ret, func(i, j int) bool {
	// 	return ctr[ret[i]] <= ctr[ret[j]]
	// })
	ret = MergeSort(ret, func(i int) int {
		return ctr[i]
	})
	// fmt.Println("ret<", ret)
	return
}

var IntPartitionsCache map[struct{ n, maxTerms int }][][]int

func IntPartitions(n, maxTerms int) (ret [][]int) {
	if v := IntPartitionsCache[struct{ n, maxTerms int }{n, maxTerms}]; v != nil {
		return v
	}
	if maxTerms == 1 {
		return [][]int{{n}}
	}
	for i := 0; i <= n; i++ {
		for _, p := range IntPartitions(n-i, maxTerms-1) {
			p = append([]int{i}, p...)
			ret = append(ret, p)
		}
	}
	IntPartitionsCache[struct{ n, maxTerms int }{n, maxTerms}] = ret
	return
}

func solveReqs(target []int, bts []Button, order []int, k int, used [][]int) int {
	if len(bts) == 0 || k >= len(order) {
		s := 0
		for _, v := range used {
			s += Sum(v)
		}
		return s
	}

	currentIdx := order[k]

	if target[currentIdx] == 0 {
		return solveReqs(target, bts, order, k+1, used)
	}

	// gather buttons that contain currentIdx
	var selectedBts, nextBts []Button
	for _, b := range bts {
		l := len(selectedBts)
		for _, i := range b {
			if currentIdx == i {
				selectedBts = append(selectedBts, b)
			}
		}

		if l == len(selectedBts) {
			nextBts = append(nextBts, b)
		}
	}
	if len(selectedBts) == 0 {
		return 999999999
	}

	ret := 9999999999
	// fmt.Println("PPPPPPPPP", k, target[currentIdx], len(selectedBts), len(IntPartitions(target[currentIdx], len(selectedBts))))
	for _, partition := range IntPartitions(target[currentIdx], len(selectedBts)) {
		// fmt.Println("PPP", target[currentIdx], len(selectedBts), partition)

		nextTarget := slices.Clone(target)

		ok := true
		for i, button := range selectedBts {
			for _, idx := range button {
				if nextTarget[idx] < partition[i] {
					ok = false
					break
				}
				nextTarget[idx] -= partition[i]
			}
			if !ok {
				break
			}
		}
		if !ok {
			continue
		}

		// fmt.Println(target[currentIdx], len(selectedBts), nextBts, target, selectedBts)
		// fmt.Println("P", "k=", k, "p=", partition, "len_p=", len(IntPartitions(target[currentIdx], len(selectedBts))))

		ret = min(ret, solveReqs(nextTarget, nextBts, order, k+1, append(used, partition)))
	}

	return ret
}

func main() {
	b, _ := os.ReadFile("./input-test.txt")
	input := string(b)
	IntPartitionsCache = make(map[struct {
		n        int
		maxTerms int
	}][][]int)

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

	for i, man := range manuals {
		// man := manuals[54]
		order := man.reqOrder()
		// fmt.Println(order, len(order), len(man.Buttons))
		ans := solveReqs(man.Requirements, man.Buttons, order, 0, [][]int{})
		fmt.Println(i, ans)
		answer += ans
	}
	fmt.Println(answer)
}
