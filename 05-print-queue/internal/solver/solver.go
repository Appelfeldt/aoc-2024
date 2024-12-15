package solver

import (
	"log"
	"strconv"
	"strings"
)

func Calculate(input string) (int, int) {
	rows := strings.Split(input, "\n")

	splitIndex := 0
	for i, r := range rows {
		if strings.TrimSpace(r) == "" {
			splitIndex = i
			break
		}
	}
	rules := rows[:splitIndex]
	pages := rows[splitIndex+1:]

	ruleset := make(map[int]map[int]bool)
	for _, r := range rules {
		order := strings.Split(strings.TrimSpace(r), "|")
		o1, _ := strconv.Atoi(order[0])
		o2, err := strconv.Atoi(order[1])
		if err != nil {
			log.Fatal(err)
		}
		if ruleset[o1] == nil {
			ruleset[o1] = make(map[int]bool)
		}
		ruleset[o1][o2] = true
	}

	validSum := 0
	invalidSum := 0
	for _, p := range pages {
		valid := true
		str_values := strings.Split(p, ",")
		values := make([]int, len(str_values))
		for i := range str_values {
			values[i], _ = strconv.Atoi(str_values[i])
		}
		length := len(values)
		for i := range values {
			for j := i + 1; j < length; j++ {
				if !ruleset[values[i]][values[j]] && ruleset[values[j]][values[i]] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			v := values[length/2]
			validSum += v
		} else {
			fixed := make([]int, length)
			for i, v := range values {
				count := 0
				for j := 0; j < length; j++ {
					if i == j {
						continue
					}
					if !ruleset[values[i]][values[j]] {
						count++
					}
				}
				fixed[count] = v
			}
			invalidSum += fixed[length/2]
		}
	}

	return validSum, invalidSum
}
