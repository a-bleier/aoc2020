package day1

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
)

func Run() {
	lines := fileio.GetLinesFromFile("day1/input.txt")
	var values []int
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		values = append(values, i)
	}

	//Find two numbers
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == 2020 {
				fmt.Println(values[i] * values[j])
				break
			}
		}
	}
	for i := 0; i < len(values)-2; i++ {
		for j := i + 1; j < len(values)-1; j++ {
			for k := j + 1; k < len(values); k++ {
				if values[i]+values[j]+values[k] == 2020 {
					fmt.Println(values[i] * values[j] * values[k])
					break
				}
			}

		}
	}

}
