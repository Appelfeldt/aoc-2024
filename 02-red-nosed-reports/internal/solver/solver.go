package solver

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculate(input string) (int, int) {

	//Separate, parse and store values.
	rows := strings.Split(input, "\n")
	reports := make([][]int, len(rows)-1)
	unsafe_reports := make([][]int, 0)
	// fmt.Printf("Rows: %d\n", len(rows))
	for i, r := range rows {
		str_levels := strings.Fields(r)
		if len(str_levels) == 0 {
			continue
		}
		err := error(nil)

		levels := make([]int, len(str_levels))
		// fmt.Printf("Row %d levels: %d\n", i, len(levels))
		for i := range str_levels {
			levels[i], err = strconv.Atoi(str_levels[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "failure parsing value: %v", str_levels[i])
				return -1, -1
			}
		}
		reports[i] = levels
	}
	// fmt.Print(reports)

	//Determine safety
	safe_reports := 0
	for _, r := range reports {
		safe := isReportSafe(r)
		if safe {
			safe_reports++
		} else {
			unsafe_reports = append(unsafe_reports, r)
		}
	}

	//Determine safety accounting for problem dampener
	safe_reports_dampener := 0
	for q, r := range unsafe_reports {
		tmp_report := make([]int, len(r)-1)
		safe := false
		for i := 0; i < len(r); i++ {
			j := 0
			for k := range r {
				if k == i {
					continue
				}
				tmp_report[j] = r[k]
				j++
			}
			if q == 0 {
				fmt.Println(tmp_report)
			}
			safe = isReportSafe(tmp_report)
			if safe {
				break
			}
		}

		if safe {
			safe_reports_dampener++
			continue
		}
	}

	return safe_reports, safe_reports + safe_reports_dampener
}

func isReportSafe(levels []int) bool {
	sign := 0
	for i := 0; i < len(levels)-1; i++ {
		delta := levels[i] - levels[i+1]
		if delta == 0 {
			return false
		}
		if sign == 0 {
			sign = delta
		} else if sign < 0 && delta >= 0 {
			return false
		} else if sign > 0 && delta <= 0 {
			return false
		}
		if iAbs(delta) < 1 || iAbs(delta) > 3 {
			return false
		}
	}
	return true
}

func iAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
