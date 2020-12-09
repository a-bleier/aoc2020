package day9

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
)

func Run() {
	lines := fileio.GetLinesFromFile("day9/input.txt")

	var numbers []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	var wrongNum int

	for x := 25; x < len(numbers); x++ {
		foundIt := false
		for i := x - 25; i < x-1; i++ {
			for k := i + 1; k < x; k++ {
				foundIt = foundIt || (numbers[i]+numbers[k] == numbers[x])
			}

		}

		if foundIt == false {
			wrongNum = numbers[x]
			break
		}
	}
	fmt.Println("A: ", wrongNum)

	low, high := -1, -1
	for i := 0; i < len(numbers)-1; i++ {
		sum := 0
		for k := i; k < len(numbers); k++ {
			sum += numbers[k]
			if sum == wrongNum {
				low, high = i, k
				break
			}
		}

		if low != -1 {
			break
		}

	}

	max, min := 0, 100000000

	for i := low; i <= high; i++ {

		if numbers[i] > max {
			max = numbers[i]
		}

		if numbers[i] < min {
			min = numbers[i]
		}

	}

	fmt.Println("B: ", max+min)

}
