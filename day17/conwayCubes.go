package day17

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
)

type layer [][]bool
type space struct {
	pocket                     map[int]layer
	zLow, zHigh                int
	expXH, expXL, expYH, expYL bool
}

func Run() {
	lines := fileio.GetLinesFromFile("day17/input.txt")
	pocket := make(map[int]layer)
	startLayer := make(layer, 0)
	for _, l := range lines {
		row := make([]bool, 0)
		for _, c := range l {
			row = append(row, c == '#')
		}
		startLayer = append(startLayer, row)
	}
	pocket[0] = startLayer
	s := space{pocket: pocket, zLow: 0, zHigh: 0}
	for i := 0; i < 6; i++ {
		s.simulate()
	}

}

//Simulate a step
func (s *space) simulate() {
	//Save the current pocket
	oldPocket := make(map[int]layer, 0)
	for y, l := range s.pocket {
		ol := make(layer, 0)

		for _, row := range l {
			or := make([]bool, 0)
			for _, c := range row {
				or = append(or, c)
			}
			ol = append(ol, or)
		}
		oldPocket[y] = ol
	}
	fmt.Println(oldPocket)
	//look through neighbors in oldPocket, generate new space
	for z := s.zLow - 1; z <= s.zHigh+1; z++ {

		for y := -1; y < len(s.pocket[0])+1; y++ {

			for x := -1; x < len(s.pocket[0][0])+1; x++ {
				activeCounter := 0
				for d1 := -1; d1 < 2; d1++ {
					for d2 := -1; d2 < 2; d2++ {
						for d3 := -1; d3 < 2; d3++ {

							nz, ny, nx := z+d1, y+d2, y+d3
							if (nz < s.zLow || nz > s.zHigh || ny < 0 || ny > len(s.pocket[0])-1 ||
								nx < 0 || nx > len(s.pocket[0])-1) == false {

								if s.pocket[nz][ny][nx] == true {
									activeCounter++
								}
							}
						}
					}
				}
				isActive := true
				if z < s.zLow || z > s.zHigh || y < 0 || y > len(s.pocket[0])-1 ||
					x < 0 || x > len(s.pocket[0])-1 {
					isActive = false
				} else {
					isActive = s.pocket[z][y][x]
				}
				if (activeCounter == 2 || activeCounter == 3) == false && isActive == true {
					s.set(x, y, z, false, oldPocket)
				} else if activeCounter == 3 && isActive == false {
					s.set(x, y, z, true, oldPocket)
				}
			}
		}
	}

	//count active cubes

	counter := 0
	for _, l := range s.pocket {
		for _, row := range l {
			for _, el := range row {
				if el {
					counter++
				}
			}
		}
	}

	fmt.Println("A: ", counter)

}

func (s *space) set(x, y, z int, isActive bool, oldPocket map[int]layer) {
	//when not enough layers, add layers
	//when not enough space in a layer, resize all layers

	fmt.Printf("Setting %d %d %d active == %t\n", x, y, z, isActive)

	if z < s.zLow || z > s.zHigh {
		s.pocket[z] = make(layer, 0)
		for i := 0; i < len(s.pocket[0]); i++ {

			s.pocket[z] = append(s.pocket[z], make([]bool, len(s.pocket[0][0])))
		}
		if z < s.zLow {
			s.zLow = z
		} else {
			s.zHigh = z
		}
	}

	if y < 0 {
		if s.expYL == false {
			s.expYL = true

			for z, _ := range s.pocket {
				newL := make(layer, 0)
				newL = append(newL, make([]bool, len(s.pocket[0][0])))
				s.pocket[z] = append(newL, s.pocket[z]...)
			}
		}
	}

}
