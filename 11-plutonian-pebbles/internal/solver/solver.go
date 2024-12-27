package solver

import (
	"strconv"
	"strings"
)

var stoneCache map[int]map[int]int

func Calculate(input string) (int, int) {

	str_stones := strings.Fields(strings.TrimSpace(input))
	stones := make([]int, len(str_stones))
	for i, s := range str_stones {
		stones[i], _ = strconv.Atoi(s)
	}

	result25 := 0
	result75 := 0
	stoneCache = make(map[int]map[int]int)
	for _, s := range stones {
		result25 += blinkStone(s, 0, 25)
		result75 += blinkStone(s, 0, 75)
	}

	return result25, result75
}

func blinkStone(stone int, iteration int, limit int) int {
	if iteration >= limit {
		return 1
	}

	if _, ok := stoneCache[limit-iteration]; !ok {
		stoneCache[limit-iteration] = make(map[int]int)
	}

	if v, ok := stoneCache[limit-iteration][stone]; ok {
		return v
	}

	if stone == 0 {
		tmp := blinkStone(1, iteration+1, limit)
		stoneCache[limit-iteration][stone] = tmp
	} else if length := intLength(stone); length%2 == 0 {
		zeroes := length / 2
		m := 1
		for k := 0; k < zeroes; k++ {
			m *= 10
		}
		half1 := stone / m
		half2 := stone - (half1 * m)

		tmp := blinkStone(half1, iteration+1, limit) + blinkStone(half2, iteration+1, limit)
		stoneCache[limit-iteration][stone] = tmp
	} else {
		tmp := blinkStone(stone*2024, iteration+1, limit)
		stoneCache[limit-iteration][stone] = tmp
	}

	return stoneCache[limit-iteration][stone]
}

func intLength(i int) int {
	length := 0
	for i > 0 {
		i /= 10
		length++
	}
	return length
}
