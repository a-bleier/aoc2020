package day12

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"math"
	"strconv"
)

func Run() {
	lines := fileio.GetLinesFromFile("day12/input.txt")

	currDir := "E"
	dirIndex := 0
	manhattanDistance := 0

	dirMap := make(map[int]string)
	manhPerDim := make(map[string]int)
	dirMap[0] = "E"
	dirMap[1] = "N"
	dirMap[2] = "W"
	dirMap[3] = "S"

	for _, line := range lines {
		action := string(line[0])
		value, _ := strconv.Atoi(line[1:])

		if action == "W" || action == "N" || action == "E" || action == "S" {
			manhPerDim[action] += value
		} else if action == "F" {
			manhPerDim[currDir] += value
		} else {
			x := value / 90
			if action == "R" {
				x *= -1
			}
			dirIndex = mod(dirIndex+x, 4)
			currDir = dirMap[dirIndex]
		}
	}

	manhattanDistance = int(math.Abs(float64(manhPerDim["W"]-manhPerDim["E"]))) +
		int(math.Abs(float64(manhPerDim["N"]-manhPerDim["S"])))

	fmt.Println("A: ", manhattanDistance)

	navigateWithWaypoint(lines)

}

func navigateWithWaypoint(lines []string) {

	wpx, wpy := -10, 1
	sx, sy := 0, 0

	for _, line := range lines {
		action := string(line[0])
		value, _ := strconv.Atoi(line[1:])

		if action == "W" || action == "N" || action == "E" || action == "S" {
			switch action {
			case "W":
				wpx += value
				break
			case "E":
				wpx -= value
				break
			case "N":
				wpy += value
				break
			case "S":
				wpy -= value
				break
			}
		} else if action == "F" {
			sx += value * wpx
			sy += value * wpy
		} else {
			x := value / 90
			if action == "L" {
				x = 4 - x
			}
			switch x {
			case 1:
				wpx, wpy = -wpy, wpx
				break
			case 2:
				wpx, wpy = -wpx, -wpy
				break
			case 3:
				wpx, wpy = wpy, -wpx
				break
			case 4:
				break
			}
		}
	}

	fmt.Println("B: ", int(math.Abs(float64(sx)))+int(math.Abs(float64(sy))))
}
func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
