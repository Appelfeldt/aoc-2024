package solver

import (
	"strings"
	"unicode"
)

func Calculate(input string) (int, int) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	height := len(rows)
	width := len(strings.TrimSpace(rows[0]))

	//Create and instantiate grid and resonance grid. '.' denotes an empty space.
	grid := make([][]rune, height)
	r_grid := make([][]rune, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]rune, width)
		r_grid[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			grid[y][x] = '.'
			r_grid[y][x] = '.'
		}
	}

	//Parse and map each coordinate to its corresponding frequency and place each antenna on the grid.
	antennas := make(map[rune][][2]int)
	for y, r := range rows {
		for x, c := range r {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				if antennas[c] == nil {
					antennas[c] = make([][2]int, 0)
				}
				antennas[c] = append(antennas[c], [2]int{y, x})
				grid[y][x] = c
				r_grid[y][x] = c
			}
		}
	}

	for freq := range antennas {
		coordCount := len(antennas[freq])

		//Create and instantiate array containing antenna indices for the current frequency
		c_indices := make([]int, coordCount)
		for i := range antennas[freq] {
			c_indices[i] = i
		}

		//Create unique pairs for all indices, then calculate and place each pairs antinodes on the grid.
		channel := make(chan [2]int)
		go permutations(c_indices, channel)
		for p := range channel {
			_ = p

			c1 := antennas[freq][p[0]]
			c2 := antennas[freq][p[1]]
			dy := c1[0] - c2[0]
			dx := c1[1] - c2[1]
			an1 := [2]int{c1[0] + dy, c1[1] + dx}
			an2 := [2]int{c2[0] - dy, c2[1] - dx}
			if an1[0] >= 0 && an1[0] < height && an1[1] >= 0 && an1[1] < width {
				grid[an1[0]][an1[1]] = '#'
				r_grid[an1[0]][an1[1]] = '#'
			}
			if an2[0] >= 0 && an2[0] < height && an2[1] >= 0 && an2[1] < width {
				grid[an2[0]][an2[1]] = '#'
				r_grid[an2[0]][an2[1]] = '#'
			}

			for i := 2; true; i++ {
				an1 = [2]int{c1[0] + dy*i, c1[1] + dx*i}
				if an1[0] >= 0 && an1[0] < height && an1[1] >= 0 && an1[1] < width {
					r_grid[an1[0]][an1[1]] = '#'
				} else {
					break
				}
			}

			for i := 2; true; i++ {
				an2 := [2]int{c2[0] - dy*i, c2[1] - dx*i}
				if an2[0] >= 0 && an2[0] < height && an2[1] >= 0 && an2[1] < width {
					r_grid[an2[0]][an2[1]] = '#'
				} else {
					break
				}
			}

		}
	}

	anCount := 0
	for _, an := range grid {
		for _, c := range an {
			if c == '#' {
				anCount++
			}
		}
	}

	resCount := 0
	for _, an := range r_grid {
		for _, c := range an {
			if c != '.' {
				resCount++
			}
		}
	}

	return anCount, resCount
}

func permutations(indices []int, c chan [2]int) {
	n := len(indices)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			a := [2]int{indices[i], indices[j]}
			c <- a
		}
	}
	close(c)
}
