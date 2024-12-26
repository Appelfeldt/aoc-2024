package solver

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calculate(input string) (int, int) {

	rows := strings.Split(strings.TrimSpace(input), "\n")
	height := len(rows)
	width := len(strings.TrimSpace(rows[0]))

	//Parse file, create topographic map and find all trailheads
	trailheads := make([][2]int, 0)
	grid := make([][]int, height)
	for y, row := range rows {
		grid[y] = make([]int, width)
		for x, r := range strings.TrimSpace(row) {
			value, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				log.Fatal()
			}
			grid[y][x] = value
			if value == 0 {
				trailheads = append(trailheads, [2]int{y, x})
			}
		}
	}

	score := 0
	rating := 0
	for _, t := range trailheads {
		t_score := 0
		visited := make([][]bool, height)
		for y := range visited {
			visited[y] = make([]bool, width)
		}
		rating += search(t[0], t[1], t[0], t[1], true, width, height, grid, visited)

		for _, row := range visited {
			for _, v := range row {
				if v {
					score++
					t_score++
				}
			}
		}

	}

	return score, rating
}

func search(y int, x int, origY int, origX int, trailhead bool, width int, height int, grid [][]int, visited [][]bool) int {
	//Stop search if the slope gradient is anything but 1 and the origin is not a trailhead
	if !trailhead && grid[y][x]-grid[origY][origX] != 1 {
		return 0
	}

	//If a peak is reached, mark the coordinate as visited
	if grid[y][x] == 9 {
		visited[y][x] = true
		return 1
	}

	rating := 0
	if y-1 >= 0 && y-1 != origY {
		rating += search(y-1, x, y, x, false, width, height, grid, visited)
	}
	if y+1 < height && y+1 != origY {
		rating += search(y+1, x, y, x, false, width, height, grid, visited)
	}
	if x-1 >= 0 && x-1 != origX {
		rating += search(y, x-1, y, x, false, width, height, grid, visited)
	}
	if x+1 < width && x+1 != origX {
		rating += search(y, x+1, y, x, false, width, height, grid, visited)
	}
	return rating
}
