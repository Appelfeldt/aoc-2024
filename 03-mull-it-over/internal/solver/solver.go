package solver

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Calculate(input string, enableToggle bool) int {

	r, err := regexp.Compile(`mul\((\d+),(\d+)\)|(do)\(\)|(don't)\(\)`)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failure compiling regex")
		return -1
	}
	if r == nil {
		fmt.Fprintln(os.Stderr, "no instructions found in input string")
		return -1
	}
	matches := r.FindAllStringSubmatch(input, -1)

	enabled := true
	sum := 0
	for _, m := range matches {
		if m[3] == "do" {
			if enableToggle {
				enabled = true
			}
		} else if m[4] == "don't" {
			if enableToggle {
				enabled = false
			}
		} else if enabled {
			v1, err := strconv.Atoi(m[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "failure parsing value: %v\n", m[1])
				return -1
			}
			v2, err := strconv.Atoi(m[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "failure parsing value: %v\n", m[2])
				return -1
			}
			sum += v1 * v2
		}
	}

	return sum
}
