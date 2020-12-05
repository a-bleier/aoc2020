package day5

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"sort"
)

func Run() {
	lines := fileio.GetLinesFromFile("day5/input.txt")

	maxSeatID := 0
	var seats []int
	for _, line := range lines {
		row := binarySearcher(0, 127, line[:len(line)-3], 'F', 'B')
		col := binarySearcher(0, 7, line[len(line)-3:], 'L', 'R')

		seatID := 8*row + col
		seats = append(seats, seatID)
		fmt.Println("row ", row, "col ", col, "seat-id", seatID)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Println("A: ", maxSeatID)

	sort.Ints(seats)

	lastSeat := -1

	for _, s := range seats {
		if s < 8 || s > 8*127-1 {
			lastSeat = s
			continue
		}
		if s-lastSeat == 2 {
			fmt.Println("B: ", s-1)
			break
		}
		lastSeat = s

	}
}

func binarySearcher(min, max int, line string, lChar, uChar byte) int {
	if min >= max {
		return min
	}
	if line[0] == lChar {
		return binarySearcher(min, (min+max)/2, line[1:], lChar, uChar)
	} else if line[0] == uChar {
		return binarySearcher((min+max)/2+1, max, line[1:], lChar, uChar)
	}

	return -1
}
