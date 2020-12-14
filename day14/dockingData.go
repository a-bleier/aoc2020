package day14

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func Run() {
	lines := fileio.GetLinesFromFile("day14/input.txt")

	partA(lines)
	partB(lines)

}

func partA(lines []string) {
	memory := make(map[int]*big.Int)
	var mask string

	for _, line := range lines {
		if line[:4] == "mask" {
			mask = strings.Split(line, " = ")[1]
		} else {
			vals := strings.Split(line[4:], "] = ")
			addr, _ := strconv.Atoi(vals[0])
			val, _ := strconv.Atoi(vals[1])

			memory[addr] = big.NewInt(0)
			memory[addr].SetUint64(uint64(val))

			for i, r := range mask {
				if r == 'X' {
					continue
				} else if r == '0' {
					memory[addr].SetBit(memory[addr], len(mask)-1-i, 0)
				} else if r == '1' {
					memory[addr].SetBit(memory[addr], len(mask)-1-i, 1)
				}
			}

		}
	}
	var sum uint64
	for _, val := range memory {
		sum += val.Uint64()
	}

	fmt.Println("A: ", sum)
}

func partB(lines []string) {
	memory := make(map[uint64]uint64)
	var mask string

	for _, line := range lines {
		if line[:4] == "mask" {
			mask = strings.Split(line, " = ")[1]
		} else {
			vals := strings.Split(line[4:], "] = ")
			val, _ := strconv.Atoi(vals[1])
			addr, _ := strconv.Atoi(vals[0])
			addrSpace := generateAddresses(mask, uint64(addr))

			for _, a := range addrSpace {
				memory[uint64(a)] = uint64(val)
			}

		}
	}
	var sum uint64
	for _, val := range memory {
		sum += val
	}

	fmt.Println("B: ", sum)
}

func generateAddresses(mask string, refAddr uint64) []uint64 {

	var addrSpace []uint64
	bigRef := &big.Int{}
	bigRef.SetUint64(refAddr)

	floatCount := strings.Count(mask, "X")
	for i := uint64(0); i < uint64(math.Pow(float64(2), float64(floatCount))); i++ {
		x := floatCount - 1

		addr := big.NewInt(0)
		addr.SetUint64(0)

		for m := 0; m < len(mask); m++ {
			if mask[m] == 'X' {
				addr.SetBit(addr, len(mask)-1-m, (uint(i)>>x)%2)
				x--
			} else if mask[m] == '0' {
				addr.SetBit(addr, len(mask)-1-m, bigRef.Bit(len(mask)-1-m))
			} else if mask[m] == '1' {
				addr.SetBit(addr, len(mask)-1-m, 1)
			}
		}

		addrSpace = append(addrSpace, addr.Uint64())

	}

	return addrSpace
}
