package solver

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Calculate(input string) (int, int) {
	grid := strings.Split(strings.TrimSpace(input), "\n")
	xmas := horizontal(grid, "XMAS") + vertical(grid, "XMAS") + diagonal(grid, "XMAS")
	x_mas := cross(grid)
	return xmas, x_mas
}

func horizontal(grid []string, value string) int {
	total := 0
	rvalue := reverse(value)

	for _, s := range grid {
		total += strings.Count(s, value)
		total += strings.Count(s, rvalue)
	}
	fmt.Printf("Horizontal: %d\n", total)
	return total
}

func vertical(grid []string, value string) int {
	total := 0
	rvalue := reverse(value)
	width := utf8.RuneCountInString(strings.TrimSpace(grid[0]))
	height := len(grid)
	for x := 0; x < width; x++ {
		r := make([]rune, height)
		for y := 0; y < height; y++ {
			r[y] = rune(grid[y][x])
		}
		column := string(r)
		total += strings.Count(column, value)
		total += strings.Count(column, rvalue)
	}
	fmt.Printf("Vertical: %d\n", total)
	return total
}

func diagonal(grid []string, value string) int {
	total := 0
	rvalue := reverse(value)
	length := utf8.RuneCountInString(strings.TrimSpace(value))
	width := utf8.RuneCountInString(strings.TrimSpace(grid[0]))
	height := len(grid)

	for x := 0; x < width-length+1; x++ {
		for y := 0; y < height-length+1; y++ {
			r1 := make([]rune, length)
			r2 := make([]rune, length)
			for i := 0; i < length; i++ {
				r1[i] = rune(grid[y+i][x+i])
				r2[i] = rune(grid[height-y-i-1][x+i])
			}
			diagonal1 := string(r1)
			diagonal2 := string(r2)
			total += strings.Count(diagonal1, value)
			total += strings.Count(diagonal1, rvalue)
			total += strings.Count(diagonal2, value)
			total += strings.Count(diagonal2, rvalue)

			// if x <= width-length && y <= height-length {

			// 	r1 := make([]rune, length)
			// 	for i := 0; i < length; i++ {
			// 		r1[i] = rune(grid[y+i][x+i])
			// 	}
			// 	diagonal1 := string(r1)
			// 	if diagonal1 == value {
			// 		total++
			// 	}
			// }

			// if x >= length && y <= height-length {
			// 	r2 := make([]rune, length)
			// 	for i := 0; i < length; i++ {
			// 		r2[i] = rune(grid[y+i][x-i])
			// 	}
			// 	diagonal2 := string(r2)
			// 	if diagonal2 == value {
			// 		total++
			// 	}
			// }

			// if x <= width-length && y >= length {
			// 	r3 := make([]rune, length)
			// 	for i := 0; i < length; i++ {
			// 		r3[i] = rune(grid[y-i][x+i])
			// 	}
			// 	diagonal3 := string(r3)
			// 	if diagonal3 == value {
			// 		total++
			// 	}
			// }

			// if x >= length && y >= length {
			// 	r4 := make([]rune, length)
			// 	for i := 0; i < length; i++ {
			// 		r4[i] = rune(grid[y-i][x-i])
			// 	}
			// 	diagonal4 := string(r4)
			// 	if diagonal4 == value {
			// 		total++
			// 	}
			// }

		}
	}
	fmt.Printf("Diagonal: %d\n", total)
	return total
}

func cross(grid []string) int {
	total := 0
	width := utf8.RuneCountInString(strings.TrimSpace(grid[0]))
	height := len(grid)

	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if grid[y][x] == 'A' &&
				((grid[y+1][x+1] == 'M' && grid[y-1][x-1] == 'S') || (grid[y+1][x+1] == 'S' && grid[y-1][x-1] == 'M')) &&
				((grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S') || (grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M')) {
				total++
			}
		}
	}

	return total
}

func reverse(s string) string {
	l := len(s)
	r := make([]rune, len(s))
	for i, j := range s {
		r[l-i-1] = j
	}

	return string(r)
}
