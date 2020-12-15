package day15

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
	"strings"
)

func Run() {
	lines := fileio.GetLinesFromFile("day15/input.txt")

	stringValues := strings.Split(lines[0], ",")

	values := make([]int, 0)
	vMap := make(map[int]int)
	for _, sV := range stringValues {
		value, _ := strconv.Atoi(sV)
		values = append(values, value)
	}

	for i, v := range values {
		if i == len(values)-1 {
			break
		}
		vMap[v] = i + 1
	}

	turnCounter := uint64(len(values))
	last := 0
	toBeAdd := values[len(values)-1]

	for true {
		last = values[len(values)-1]
		if turnCounter == uint64(2020) {
			fmt.Println("A: ", last)
		} else if turnCounter == uint64(30000000) {
			fmt.Println("B: ", last)
			break
		}
		turnCounter++
		//Find last
		found := false
		diff := 0

		if vMap[last] != 0 {
			diff = len(values) - vMap[last]
			vMap[last] = len(values)
			found = true
		}

		vMap[toBeAdd] = len(values)

		if found == true {
			values = append(values, diff)
			toBeAdd = diff
		} else {
			values = append(values, 0)
			toBeAdd = 0
		}

	}

}
