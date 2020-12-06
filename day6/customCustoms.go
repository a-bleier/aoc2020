package day6

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strings"
)

func Run() {
	lines := fileio.GetLinesFromFile("day6/input.txt")

	//Read groups
	var groups [][]string
	var group []string
	for _, line := range lines {

		if len(line) == 0 {
			groups = append(groups, group)
			group = make([]string, 0)
		} else {
			group = append(group, line)
		}
	}
	groups = append(groups, group)

	counterA := 0
	counterB := 0
	for _, g := range groups {
		//reset lookup before each group is checked

		s := strings.Join(g, "")
		for i := 0; i < 26; i++ {
			found := true
			for _, mem := range g {
				found = found && strings.Contains(mem, string('a'+i))
			}
			if found == true {
				counterB++
			}
			if strings.Count(s, string('a'+i)) != 0 {
				counterA++
			}

		}
	}

	fmt.Println("A: ", counterA)
	fmt.Println("B ", counterB)
}
