package solver

import (
	"log"
	"strconv"
	"strings"
)

func Calculate(input string) (int, int) {

	//Parse file and calculate disk block size
	str_diskmap := strings.TrimSpace(input)
	blockCount := 0
	diskmap := make([]int, len(str_diskmap))
	for i, v := range str_diskmap {

		b, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		blockCount += b
		diskmap[i] = b
	}

	//Instantiate two disks, one for each puzzle part
	disk1 := make([]int, blockCount)
	disk2 := make([]int, blockCount)

	cursor := 0                //Front cursor
	b_cursor := blockCount - 1 //Back cursor
	for i, d := range diskmap {
		end := cursor + d
		for cursor < end {
			if i%2 == 0 { //If even then it's a file
				disk1[cursor] = i / 2
				disk2[cursor] = i / 2
				cursor++
			} else { //else it's an empty block
				disk1[cursor] = -1
				disk2[cursor] = -1
				cursor++
			}
		}
	}

	//Compact memory of disk 1
	cursor = 0
	for cursor < b_cursor {

		//Loop until a an empty space is found
		if disk1[cursor] != -1 {
			cursor++
			continue
		}

		//Loop until a file block is found
		for disk1[b_cursor] == -1 {
			b_cursor--
		}

		//Exit early if cursor and b_cursor have met or passed each other
		if cursor >= b_cursor {
			break
		}

		//Swap memory of cursor locations.
		disk1[cursor], disk1[b_cursor] = disk1[b_cursor], disk1[cursor]
	}

	//Compact memory of disk 2
	b_cursor = blockCount - 1
	for b_cursor >= 0 {
		cursor = 0
		//Search through the disk from the back until a valid file id is found (not -1)
		for disk2[b_cursor] == -1 {
			b_cursor--
		}

		//Store found disk id and determine the files length by search additional blocks for matching id.
		block_id := disk2[b_cursor]
		block_length := 0
		for disk2[b_cursor] == block_id {
			block_length++
			b_cursor--
			if b_cursor < 0 {
				break
			}
		}

		if b_cursor < 0 {
			break
		}

		//Search for available space
		done := false
		for cursor < blockCount && cursor < b_cursor && !done {

			if disk2[cursor] == -1 { //If empty space found
				//Determine block length of empty space
				empty_length := 1
				for i := cursor + 1; i < blockCount && disk2[i] == -1; i++ {
					empty_length++
					if empty_length >= block_length {
						break
					}
				}

				//If the empty space can house the file, move
				if empty_length >= block_length {
					for i := 0; i < block_length; i++ {
						//Swap block data
						disk2[cursor+i], disk2[b_cursor+1+i] = disk2[b_cursor+1+i], disk2[cursor+i]
					}
					done = true
				} else {
					cursor += empty_length
					continue
				}
			}
			cursor++
		}
	}

	//Calculate checksums
	var result1 int = 0
	for i, v := range disk1 {
		if v == -1 {
			continue
		}
		result1 += i * v
	}

	var result2 int = 0
	for i, v := range disk2 {
		if v == -1 {
			continue
		}
		result2 += i * v
	}
	return result1, result2
}
