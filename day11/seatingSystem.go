package day11

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
)

func Run() {

	lines := fileio.GetLinesFromFile("day11/input.txt")

	model := make([][]rune, 0)
	for _, line := range lines {
		row := []rune(line)
		model = append(model, row)
	}
	partA(model)
	model = make([][]rune, 0)
	for _, line := range lines {
		row := []rune(line)
		model = append(model, row)
	}
	partB(model)

}

func partA(model [][]rune) {

	newModel := make([][]rune, 0)

	genCount := 0
	for true {
		genCount++
		newModel = make([][]rune, 0)

		for _, row := range model {
			nRow := make([]rune, 0)
			for _, r := range row {
				nRow = append(nRow, r)
			}
			newModel = append(newModel, nRow)
		}

		changes := 0
		occupied := 0
		for y := 0; y < len(model); y++ {
			for x := 0; x < len(model[y]); x++ {

				if model[y][x] == '.' {
					continue
				}
				//Iterate through mask
				counter := 0
				for i := 0; i < 9; i++ {
					if i == 4 {
						continue
					}
					dy, dx := getMaskIndex(i)
					ny, nx := y+dy, x+dx
					if nx == -1 || ny == -1 || nx == len(model[y]) || ny == len(model) {
						continue
					}
					if model[ny][nx] == '#' {
						counter++
					}
				}

				if model[y][x] == '#' {
					occupied++
				}

				if counter > 3 && model[y][x] == '#' {
					newModel[y][x] = 'L'
					changes++
				} else if counter == 0 && model[y][x] == 'L' {
					newModel[y][x] = '#'
					changes++
				}

			}
		}
		if changes == 0 {
			fmt.Println("A: ", occupied)
			break
		}

		model = newModel
	}

}

func partB(model [][]rune) {

	newModel := make([][]rune, 0)

	genCount := 0
	for true {
		genCount++
		newModel = make([][]rune, 0)

		for _, row := range model {
			nRow := make([]rune, 0)
			for _, r := range row {
				nRow = append(nRow, r)
			}
			newModel = append(newModel, nRow)
		}

		changes := 0
		occupied := 0
		for y := 0; y < len(model); y++ {
			for x := 0; x < len(model[y]); x++ {

				if model[y][x] == '.' {
					continue
				}
				//Iterate through mask
				counter := 0
				for i := 0; i < 9; i++ {
					if i == 4 {
						continue
					}
					dy, dx := getMaskIndex(i)

					ny, nx := y+dy, x+dx
					for ny >= 0 && nx >= 0 && ny < len(model) && nx < len(model[y]) {
						if model[ny][nx] == '#' {
							counter++
							break
						} else if model[ny][nx] == 'L' {
							break
						}
						ny, nx = ny+dy, nx+dx
					}
				}

				if model[y][x] == '#' {
					occupied++
				}

				if counter > 4 && model[y][x] == '#' {
					newModel[y][x] = 'L'
					changes++
				} else if counter == 0 && model[y][x] == 'L' {
					newModel[y][x] = '#'
					changes++
				}

			}
		}
		if changes == 0 {
			fmt.Println("B: ", occupied)
			break
		}

		model = newModel
	}

}

func getMaskIndex(pos int) (int, int) {
	switch pos {
	case 0:
		return -1, -1
		break
	case 1:
		return -1, 0
		break
	case 2:
		return -1, 1
		break
	case 3:
		return 0, -1
		break
	case 4:
		return 0, 0
		break
	case 5:
		return 0, 1
		break
	case 6:
		return 1, -1
		break
	case 7:
		return 1, 0
		break
	case 8:
		return 1, 1
		break
	}

	return 0, 0
}
