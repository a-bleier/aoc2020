package day10

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"sort"
	"strconv"
)

func Run() {
	lines := fileio.GetLinesFromFile("day10/input.txt")

	var jolts []int

	for _, line := range lines {
		j, _ := strconv.Atoi(line)
		jolts = append(jolts, j)
	}

	sort.Ints(jolts)

	var diffCounter [3]int

	for i := 0; i < len(jolts)-1; i++ {
		diff := jolts[i+1] - jolts[i]
		if diff == 0 {
			continue
		}
		diffCounter[diff-1]++
	}
	diffCounter[jolts[0]-1]++
	diffCounter[2]++

	fmt.Println("A: ", diffCounter[0]*diffCounter[2])

	jolts = append([]int{0}, jolts...)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	start, end := 0, 0
	possibilities := 1

	for i := 0; i < len(jolts)-1; i++ {
		diff := jolts[i+1] - jolts[i]
		if diff == 3 {
			end = i + 1
			possibilities *= countPossibilities(jolts[start:end])
			start = i + 1
		}
	}

	fmt.Println("B: ", possibilities)

	//Nope, bad idea
	//fmt.Println("B: ", countPossibilities(jolts))

}

func countPossibilities(jolts []int) int {

	if len(jolts) == 1 {
		return 1
	}

	ps := powerSubSets(jolts)

	//Removing duplicates
	n := len(ps)
	for i := 0; i < n-1; i++ {
		for k := i + 1; k < n; k++ {
			if Equal(ps[i], ps[k]) {
				ps = append(ps[:k], ps[k+1:]...)
				n--
			}
		}
	}

	//counting valid sets
	counter := 0
	for _, set := range ps {
		sort.Ints(set)

		if len(set) == 0 {
			continue
		}

		if set[0] != jolts[0] || set[len(set)-1] != jolts[len(jolts)-1] {
			continue
		}
		flag := false
		for i := 0; i < len(set)-1; i++ {
			if set[i+1]-set[i] > 3 {
				flag = true
				break
			}
		}

		if flag == false {
			counter++
		}

	}

	if counter == 0 {
		return 1
	} else {
		return counter
	}

}

// finds all powersets; cpy pasted from stackoverflow
func powerSubSets(nums []int) [][]int {
	result := [][]int{}
	tmp := []int{}
	result = powerSubsetsDFS(nums, tmp, 0, result)
	return result
}

func powerSubsetsDFS(nums []int, tmp []int, idx int, result [][]int) [][]int {
	result = append(result, tmp)
	for i := idx; i < len(nums); i++ {
		tmp2 := append(tmp, nums[i]) // store in a new variable
		result = powerSubsetsDFS(nums, tmp2, i+1, result)
	}
	return result
}

//tests whether two slices are equal
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
