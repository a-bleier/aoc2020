package day3

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
)

func Run() {
	lines := fileio.GetLinesFromFile("day3/input.txt")

	var terrain [][]int
	for _, line := range lines {
		var row []int
		for pos, char := range line {
			if char == '#' {
				row = append(row, pos)
			}
		}
		terrain = append(terrain, row)
	}

	rowLen := len(lines[0])

	fmt.Println("A: ", checkSlope(3, 1, terrain, rowLen))

	product := checkSlope(1, 1, terrain, rowLen) *
		checkSlope(3, 1, terrain, rowLen) *
		checkSlope(5, 1, terrain, rowLen) *
		checkSlope(7, 1, terrain, rowLen) *
		checkSlope(1, 2, terrain, rowLen)
	fmt.Println("B: ", product)
}

func checkSlope(x int, y int, terrain [][]int, rowLen int) int {
	counter := 0

	pos := -x
	for rowI := 0; rowI < len(terrain); rowI += y {
		pos += x
		for _, el := range terrain[rowI] {
			if el == pos%rowLen {
				counter++
				break
			}
		}
	}

	return counter
}
