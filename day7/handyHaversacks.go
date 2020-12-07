package day7

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
	"strings"
)

func Run() {

	lines := fileio.GetLinesFromFile("day7/input.txt")

	dim := len(lines)
	var adj [][]int
	sgB := 0
	bags := make(map[string]int, dim)
	for i := 0; i < dim; i++ {
		adj = append(adj, make([]int, dim))
	}

	for i, line := range lines {
		tokens := strings.Split(line, " contain ")
		b := tokens[0][:len(tokens[0])-1]
		if b == "shiny gold bag" {
			sgB = i
		}
		bags[b] = i
	}

	for k, line := range lines {
		tokens := strings.Split(line, " contain ")
		rBags := strings.Split(tokens[1], ", ")
		for i, rB := range rBags {
			var num int
			if rB == "no other bags." {
				num = 0
				for l := 0; l < dim; l++ {
					adj[bags[tokens[0][:len(tokens[0])-1]]][l] = num
					break
				}
			} else {

				numString := strings.Split(rB, " ")[0]
				num, _ = strconv.Atoi(numString)

				if i == len(rBags)-1 && num != 1 {
					rB = rB[:len(rB)-2]
				} else if i == len(rBags)-1 || num != 1 {
					rB = rB[:len(rB)-1]
				}

				rI := bags[rB[2:]]
				lI := k

				adj[rI][lI] = num
			}
		}
	}

	fmt.Println("A: ", BFS(adj, sgB))
	fmt.Println("B: ", DFS(adj, sgB))

}

/*

idea: Build a matrix; search if a path exists to golden shiny bag from each bag
*/

//adj must be n x n
func BFS(adj [][]int, s int) int {

	counter := 0
	n := len(adj)
	visited := make([]bool, n)
	q := make([]int, 0)

	visited[s] = true
	q = append(q, s)

	for len(q) != 0 {
		s = q[len(q)-1]
		q = q[:len(q)-1]

		for i, val := range adj[s] {
			if visited[i] == false && val != 0 {
				q = append(q, i)
				visited[i] = true
				counter++

			}
		}
	}

	return counter
}
func DFS(adj [][]int, s int) int {

	counter := 0
	for i := 0; i < len(adj); i++ {
		val := adj[i][s]
		if val != 0 {
			counter += val + val*DFS(adj, i)
		}
	}

	return counter
}
