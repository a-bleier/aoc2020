package day13

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"math/big"
	"strconv"
	"strings"
)

var one = big.NewInt(1)

func Run() {
	lines := fileio.GetLinesFromFile("day13/input.txt")

	timestamp, _ := strconv.Atoi(lines[0])
	ids := make([]int, 0)
	idLine := strings.Split(lines[1], ",")

	for _, r := range idLine {
		if r == "x" {
			continue
		}
		id, _ := strconv.Atoi(string(r))
		ids = append(ids, id)
	}

	bestId := 0
	bestTime := 0
	for i, id := range ids {
		diff := (timestamp - (timestamp % id) + id) - timestamp
		if i == 0 {
			bestId = id
			bestTime = diff
			continue
		}
		if diff < bestTime {
			bestTime = diff
			bestId = id
		}
	}

	fmt.Println("A: ", bestId*bestTime)

	busIds := make([]int, 0)
	for _, r := range idLine {
		if r == "x" {
			busIds = append(busIds, 1)
			continue
		}
		id, _ := strconv.Atoi(string(r))
		busIds = append(busIds, id)
	}

	a := make([]*big.Int, 0)
	m := make([]*big.Int, 0)

	//highestIndex := 0
	for i := 0; i < len(busIds); i++ {
		if busIds[i] == 1 {
			continue
		}

		m = append(m, big.NewInt(int64(busIds[i])))
		if i == 0 {
			a = append(a, big.NewInt(int64(0)))
		} else {
			a = append(a, big.NewInt(int64(busIds[i]-i)))
		}

	}

	result, _ := crt(a, m)

	fmt.Println("B: ", result.Uint64())

}

//Thank to Rosetta Code
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
