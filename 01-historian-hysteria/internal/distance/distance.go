package distance

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Calculate(input string) (int, int) {

	//Separate, parse, store and sort values in two arrays.
	left := make([]int, 1000)
	right := make([]int, 1000)

	rows := strings.Split(input, "\n")
	for i, r := range rows {
		columns := strings.Fields(r)
		if len(columns) == 0 {
			break
		}
		err := error(nil)
		left[i], err = strconv.Atoi(columns[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failure parsing value: %v", err)
		}

		right[i], err = strconv.Atoi(columns[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failure parsing value: %v", err)
		}
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	//Calculate distance sum
	distance := 0
	for i := range left {
		distance += iAbs(left[i] - right[i])
	}

	//Calculate similarity score
	id_count := make(map[int]int)
	for _, j := range right {
		id_count[j]++
	}

	similarity := 0
	for _, j := range left {
		similarity += j * id_count[j]
	}

	return distance, similarity
}

func iAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
