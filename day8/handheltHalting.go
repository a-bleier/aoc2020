package day8

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
	"strings"
)

func Run() {
	instructions := fileio.GetLinesFromFile("day8/input.txt")

	_, acc := testProgram(instructions)

	fmt.Println("A: ", acc)

	for i, line := range instructions {
		instr := strings.Split(line, " ")
		if instr[0] == "nop" {
			instructions[i] = "jmp" + " " + instr[1]
		} else if instr[0] == "jmp" {
			instructions[i] = "nop" + " " + instr[1]
		} else {
			continue
		}

		isRepaired, accu := testProgram(instructions)
		if isRepaired == true {
			fmt.Println("B: ", accu)
			break
		}

		if instr[0] == "nop" {
			instructions[i] = "nop" + " " + instr[1]
		} else if instr[0] == "jmp" {
			instructions[i] = "jmp" + " " + instr[1]
		}
	}

}

func testProgram(instructions []string) (bool, int) {
	instrMap := make([]bool, len(instructions))
	ip := 0
	acc := 0
	ret := false

	for instrMap[ip] == false {
		instruction := strings.Split(instructions[ip], " ")
		instrMap[ip] = true
		if instruction[0] == "nop" {
			ip++
		} else if instruction[0] == "acc" {
			op, _ := strconv.Atoi(instruction[1])
			acc += op
			ip++
		} else if instruction[0] == "jmp" {
			op, _ := strconv.Atoi(instruction[1])
			ip += op
		}

		if ip >= len(instructions) {
			ret = true
			break
		}
	}

	return ret, acc
}
