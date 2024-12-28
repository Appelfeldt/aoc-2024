package solver

import (
	"sort"
	"strings"
)

type cell struct {
	X     int
	Y     int
	Plant rune
	Fence fence
}

type fence int

const (
	Up fence = 1 << iota
	Down
	Left
	Right
)

var grid []cell
var open []cell
var visited []bool

var height int
var width int

func Calculate(input string) (int, int) {

	//Init
	rows := strings.Split(strings.TrimSpace(input), "\n")
	height = len(rows)
	width = len(strings.TrimSpace(rows[0]))

	grid = make([]cell, width*height)
	open = make([]cell, width*height)
	visited = make([]bool, width*height)

	for y, row := range rows {
		for x, r := range row {
			grid[y*width+x] = cell{X: x, Y: y, Plant: r}
			open[y*width+x] = grid[y*width+x]
		}
	}

	//Iterates through visited cells.
	//If it's not visited, collect all connected cells with the same plant and mark them as visited.
	//The collected cells are saved as a region instance.
	regions := make(map[int][]*cell)
	_ = regions
	r_id := 0
	result1 := 0
	for i, v := range visited {
		if v {
			continue
		}
		regions[r_id] = make([]*cell, 0)
		perimeter := search(&grid[i], r_id, regions)
		result1 += perimeter * len(regions[r_id])
		r_id++
	}

	result2 := 0
	//Loop through all regions
	for _, cells := range regions {
		sides := 0

		//Create groups each fence direction and loop through them
		fences := make([][]*cell, 4)
		for i := 0; i < 4; i++ {
			fences[i] = make([]*cell, 0)

			//Add cells to group if they have a fence in that direction
			for _, c := range cells {
				if (c.Fence & (1 << i)) > 0 {
					fences[i] = append(fences[i], c)
				}
			}

			//Sort fences into rows/columns
			lines := make(map[int][]int, 0)
			for _, c := range fences[i] {
				if i < 2 { //Row
					if _, ok := lines[c.Y]; !ok {
						lines[c.Y] = make([]int, 0)
					}
					lines[c.Y] = append(lines[c.Y], c.X)
				} else { //Column
					if lines[c.X] == nil {
						lines[c.X] = make([]int, 0)
					}
					lines[c.X] = append(lines[c.X], c.Y)
				}
			}

			//Sort cells in rows/columns
			for _, v := range lines {
				if len(v) < 2 {
					continue
				}
				sort.Ints(v)
			}

			//Calculate sides based on continuity of rows/columns
			for _, c := range lines {
				sides++
				if len(c) == 1 {
					continue
				}
				l_sides := 1
				for k := 0; k < len(c)-1; k++ {
					if c[k+1]-c[k] != 1 {
						sides++
						l_sides++
					}
				}
			}

		}
		result2 += sides * len(cells)
	}

	return result1, result2
}

func search(current *cell, r_id int, regions map[int][]*cell) int {
	if visited[current.Y*width+current.X] {
		return 0
	} else {
		regions[r_id] = append(regions[r_id], current)
		visited[current.Y*width+current.X] = true
	}

	p := 0

	if current.Y+1 < height {
		if c := &grid[(current.Y+1)*width+current.X]; current.Plant == c.Plant {
			p += search(c, r_id, regions)
		} else {
			current.Fence = current.Fence | Down
			p++
		}
	} else {
		current.Fence = current.Fence | Down
		p++
	}

	if current.Y-1 >= 0 {
		if c := &grid[(current.Y-1)*width+current.X]; current.Plant == c.Plant {
			p += search(c, r_id, regions)
		} else {
			current.Fence = current.Fence | Up
			p++
		}
	} else {
		current.Fence = current.Fence | Up
		p++
	}

	if current.X+1 < width {
		if c := &grid[current.Y*width+current.X+1]; current.Plant == c.Plant {
			p += search(c, r_id, regions)
		} else {
			current.Fence = current.Fence | Right
			p++
		}
	} else {
		current.Fence = current.Fence | Right
		p++
	}

	if current.X-1 >= 0 {
		if c := &grid[current.Y*width+current.X-1]; current.Plant == c.Plant {
			p += search(c, r_id, regions)
		} else {
			current.Fence = current.Fence | Left
			p++
		}
	} else {
		current.Fence = current.Fence | Left
		p++
	}

	return p
}
