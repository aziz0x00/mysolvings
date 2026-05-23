//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func mex(list []int) int {
	sort.Ints(list)
	for i := range len(list) {
		if list[i] != i {
			return i
		}
	}
	return len(list)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	mexMat := make([][]int, n)

	for i := range n {
		mexMat[i] = make([]int, n)
		for j := range n {
			var selection []int
			for k := range j {
				selection = append(selection, mexMat[i][k])
			}
			for k := range i {
				selection = append(selection, mexMat[k][j])
			}

			mexMat[i][j] = mex(selection)
		}
	}

	for i := range n {
		for j := range n {
			print(mexMat[i][j], " ")
		}
		println()
	}
}
