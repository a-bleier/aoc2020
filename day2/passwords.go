package day2

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
	"strings"
)

type entry struct {
	min      int
	max      int
	char     string
	password string
}

func Run() {
	lines := fileio.GetLinesFromFile("day2/input.txt")
	var entries []entry
	for _, line := range lines {
		temp := strings.Split(line, " ")
		minmax := strings.Split(temp[0], "-")

		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		char := strings.Split(temp[1], "")[0]
		password := temp[2]

		e := entry{min, max, char, password}

		entries = append(entries, e)

	}

	counter := 0

	for _, e := range entries {
		n := strings.Count(e.password, e.char)
		if n >= e.min && n <= e.max {
			counter++
		}
	}

	fmt.Println("A: ", counter)

	counter = 0
	for _, e := range entries {
		if (e.password[e.min-1] == e.char[0]) != (e.password[e.max-1] == e.char[0]) {
			counter++
		}
	}

	fmt.Println("B: ", counter)
}
