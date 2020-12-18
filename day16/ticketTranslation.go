package day16

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
	"strings"
)

func Run() {
	input := fileio.GetLinesFromFile("day16/input.txt")

	rules := make([][]int, 0)
	ruleNames := make([]string, 0)
	tickets := make([][]int, 0)
	myTicket := make([]int, 0)

	var state int
	for _, line := range input {
		if line == "" || line == "your ticket:" || line == "nearby tickets:" {
			state++
			continue
		}

		switch state {
		case 0:
			toks := strings.Split(line, ": ")
			ruleNames = append(ruleNames, toks[0])
			sRanges := strings.Split(toks[1], " or ")
			rule := make([]int, 4)
			sN1 := strings.Split(sRanges[0], "-")
			sN2 := strings.Split(sRanges[1], "-")
			rule[0], _ = strconv.Atoi(sN1[0])
			rule[1], _ = strconv.Atoi(sN1[1])
			rule[2], _ = strconv.Atoi(sN2[0])
			rule[3], _ = strconv.Atoi(sN2[1])
			rules = append(rules, rule)
			break
		case 2:
			sN := strings.Split(line, ",")
			for _, s := range sN {
				num, _ := strconv.Atoi(s)
				myTicket = append(myTicket, num)
			}
			break
		case 4:
			t := make([]int, 0)
			sN := strings.Split(line, ",")
			for _, s := range sN {
				num, _ := strconv.Atoi(s)
				t = append(t, num)
			}
			tickets = append(tickets, t)
			break
		}
	}

	var errorRate int
	validTickets := make([][]int, 0)

	for _, t := range tickets {
		isValid := true
		for _, v := range t {
			flag := false
			for _, r := range rules {
				flag = flag || !(v < r[0] || (v > r[1] && v < r[2]) || v > r[3])
			}
			if flag == false {
				errorRate += v
				isValid = false
			}
		}
		if isValid == true {
			validTickets = append(validTickets, t)
		}
	}

	fmt.Println("A: ", errorRate)

	ruleCandidates := make([][]int, 0)

	for i := 0; i < len(validTickets[0]); i++ { //Values in ticket

		fieldCand := make([]int, 0)
		for k := 0; k < len(rules); k++ { //rules
			flag := true
			for m := 0; m < len(validTickets); m++ { //tickets
				v := validTickets[m][i]
				r := rules[k]
				flag = flag && !(v < r[0] || (v > r[1] && v < r[2]) || v > r[3])
				if flag == false {
					break
				}
			}
			if flag == true {
				fieldCand = append(fieldCand, k)
			}

		}
		ruleCandidates = append(ruleCandidates, fieldCand)
	}

	orderedRules := make([]int, len(rules))
	for i, _ := range orderedRules {
		orderedRules[i] = -1
	}

	for i := 0; i < len(rules); i++ {

		el := 0
		for k := 0; k < len(rules); k++ {
			if len(ruleCandidates[k]) == 1 && orderedRules[k] == -1 {
				el = ruleCandidates[k][0]
				orderedRules[k] = el
			}
		}
		for k := 0; k < len(rules); k++ {
			if len(ruleCandidates[k]) == 1 {
				continue
			}
			rmI := pos(el, ruleCandidates[k])
			if rmI == -1 {
				continue
			}
			ruleCandidates[k][rmI] = ruleCandidates[k][len(ruleCandidates[k])-1]
			ruleCandidates[k] = ruleCandidates[k][:len(ruleCandidates[k])-1]
		}

	}
	prod := 1
	for i, r := range orderedRules {
		toks := strings.Split(ruleNames[r], " ")
		if toks[0] == "departure" {
			prod *= myTicket[i]
		}
	}

	fmt.Println("B: ", prod)
}

func pos(el int, sl []int) int {
	for i, elem := range sl {
		if el == elem {
			return i
		}
	}
	return -1
}
