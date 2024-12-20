package solver

import (
	"fmt"
	"log"
	"strings"
)

func Calculate(input string) (int, int) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	height := len(rows)
	grid := make([][]rune, height)

	for y, r := range rows {
		grid[y] = []rune(strings.TrimSpace(r))
	}
	width := len(grid[0])
	startX := -1
	startY := -1
	pX := -1
	pY := -1
	done := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '^' {
				pX = x
				pY = y
				if startX == -1 {
					startX = x
					startY = y
				}
				done = true
				grid[y][x] = 'U'
				break
			}
		}
		if done {
			break
		}
	}

	done = false
	var d Direction = Up
	for !done {
		switch d {
		case Up:
			if pY-1 < 0 {
				done = true
			} else if grid[pY-1][pX] == '#' {
				d = Right
			} else {
				pY--
				grid[pY][pX] = 'U' //Up
			}
		case Down:
			if pY+1 >= height {
				done = true
			} else if grid[pY+1][pX] == '#' {
				d = Left
			} else {
				pY++
				grid[pY][pX] = 'D' //Down
			}
		case Left:
			if pX-1 < 0 {
				done = true
			} else if grid[pY][pX-1] == '#' {
				d = Up
			} else {
				pX--
				grid[pY][pX] = 'L' //Left
			}
		case Right:
			if pX+1 >= width {
				done = true
			} else if grid[pY][pX+1] == '#' {
				d = Down
			} else {
				pX++
				grid[pY][pX] = 'R' //Right
			}
		default:
			log.Fatal("unknown direction state")
		}
	}
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			switch grid[y][x] {
			case 'U':
				fallthrough
			case 'D':
				fallthrough
			case 'L':
				fallthrough
			case 'R':
				count++
			}
		}
	}

	obstructCount := 0
	done = false
	d = Up
	pX = startX
	pY = startY
	gridCopy := make([][]rune, len(grid))
	copy(gridCopy, grid)

	iterationCount := 0
	for !done {
		if canLoop(grid, pX, pY, width, height, startX, startY, d) {
			tX, tY := pX, pY
			switch d {
			case Up:
				tY--
			case Down:
				tY++
			case Left:
				tX--
			case Right:
				tX++
			}
			gridCopy[tY][tX] = 'O'
		}
		iterationCount++
		switch d {
		case Up:
			if pY-1 < 0 {
				done = true
			} else if grid[pY-1][pX] == '#' {
				d = Right
			} else {
				pY--
			}
		case Down:
			if pY+1 >= height {
				done = true
			} else if grid[pY+1][pX] == '#' {
				d = Left
			} else {
				pY++
			}
		case Left:
			if pX-1 < 0 {
				done = true
			} else if grid[pY][pX-1] == '#' {
				d = Up
			} else {
				pX--
			}
		case Right:
			if pX+1 >= width {
				done = true
			} else if grid[pY][pX+1] == '#' {
				d = Down
			} else {
				pX++
			}
		default:
			log.Fatal("unknown direction state")
		}
		fmt.Printf("\rIterations: %d", iterationCount)
	}

	fmt.Println()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if gridCopy[y][x] == 'O' {
				obstructCount++
			}
		}
	}

	return count, obstructCount
}

type Direction int

const (
	Up    Direction = 1
	Down  Direction = 2
	Left  Direction = 4
	Right Direction = 8
)

func canLoop(grid [][]rune, oX int, oY int, width int, height int, sX int, sY int, direction Direction) bool {

	visited := make([][]Direction, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]Direction, width)
	}
	visited[sY][sX] = Up

	switch direction {
	case Up:
		if oY-1 < 0 {
			return false
		}
		oY--
	case Down:
		if oY+1 >= height {
			return false
		}
		oY++
	case Left:
		if oX-1 < 0 {
			return false
		}
		oX--
	case Right:
		if oX+1 >= width {
			return false
		}
		oX++
	}

	done := false
	d := Up
	pX := sX
	pY := sY
	tX, tY := 0, 0
	for !done {
		switch d {
		case Up:
			tY = pY - 1
			tX = pX
			if tY < 0 {
				done = true
				return false
			} else if (visited[tY][tX] & d) > 0 {
				return true
			} else if grid[tY][tX] == '#' || (tY == oY && tX == oX) {
				d = Right
			} else {
				pY--
				visited[pY][pX] = visited[pY][pX] | d
			}

		case Down:
			tY = pY + 1
			tX = pX
			if tY >= height {
				done = true
				return false
			} else if (visited[tY][tX] & d) > 0 {
				return true
			} else if grid[tY][tX] == '#' || (tY == oY && tX == oX) {
				d = Left
			} else {
				pY++
				visited[pY][pX] = visited[pY][pX] | d
			}
		case Left:
			tY = pY
			tX = pX - 1
			if tX < 0 {
				done = true
				return false
			} else if (visited[tY][tX] & d) > 0 {
				return true
			} else if grid[tY][tX] == '#' || (tY == oY && tX == oX) {
				d = Up
			} else {
				pX--
				visited[pY][pX] = visited[pY][pX] | d
			}
		case Right:
			tY = pY
			tX = pX + 1
			if tX >= width {
				done = true
				return false
			} else if (visited[tY][tX] & d) > 0 {
				return true
			} else if grid[tY][tX] == '#' || (tY == oY && tX == oX) {
				d = Down
			} else {
				pX++
				visited[pY][pX] = visited[pY][pX] | d
			}
		default:
			log.Fatal("unknown direction state")
		}

	}
	return false
}
