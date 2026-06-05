package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	b, _ := os.ReadFile("./input.txt")

	sum := 50
	password := 0

	for s := range strings.SplitSeq(string(b), "\n") {
		if len(s) == 0 {
			continue
		}

		val, err := strconv.Atoi(s[1:])

		if err != nil {
			panic(err)
		}

		if s[0] == 'L' {
			val *= -1
		}

		passes := 0
		if val >= 0 {
			passes = (sum + val) / 100
		} else {
			passes = ((100-sum)%100 - val) / 100
		}

		password += passes

		sum = (sum + val) % 100
		if sum < 0 {
			sum += 100
		}
	}

	fmt.Println(password)
}
