package solver

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func Calculate(input string) (int, int) {
	rows := strings.Split(strings.TrimSpace(input), "\n")

	result1 := 0
	result2 := 0
	for _, r := range rows {
		row := strings.Split(r, ": ")
		goal, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal(err)
		}

		//Separate and parse numbers
		str_values := strings.Fields(row[1])
		values := make([]int, len(str_values))
		for i, s := range str_values {
			values[i], err = strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
		}

		//Test two operators
		total := 0
		vCount := len(values)
		combinations := int(math.Pow(2, float64(vCount-1)))
		for i := 0; i < combinations; i++ {
			total = values[0]
			for j := 1; j < vCount; j++ {
				mask := (1 << (j - 1))
				if (i & mask) != 0 {
					total += values[j]
				} else {
					total *= values[j]
				}
			}
			if total == goal {
				result1 += goal
				break
			}
		}

		//Test three operators
		operators := []rune{'+', '*', '|'}
		buffer := make([]rune, 0)
		out := make(chan []rune)
		combinations = int(math.Pow(3, float64(vCount-1)))
		p_array := make([][]rune, combinations)

		go permutations(vCount-1, len(operators), operators, buffer, out)

		for i := 0; i < combinations; i++ {
			p_array[i] = <-out
		}

		for _, p := range p_array {
			total := values[0]
			for i := 1; i < vCount; i++ {
				switch p[i-1] {
				case '+':
					total += values[i]
				case '*':
					total *= values[i]
				case '|':
					total *= int(math.Pow(10, float64(digitCount(values[i]))))
					total += values[i]
				}
			}
			if total == goal {
				result2 += total
				break
			}
		}
	}

	return result1, result2
}

func digitCount(number int) (count int) {
	for number > 0 {
		number /= 10
		count++
	}
	return
}

func permutations(k int, tCount int, tokens []rune, buffer []rune, out chan []rune) {
	if k == 0 {
		out <- buffer
		return
	}

	for i := 0; i < tCount; i++ {
		bufferCopy := make([]rune, len(buffer))
		bufferCopy = append(bufferCopy, tokens[i])
		copy(bufferCopy, buffer)
		permutations(k-1, tCount, tokens, bufferCopy, out)
	}
}
