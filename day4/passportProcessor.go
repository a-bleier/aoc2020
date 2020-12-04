package day4

import (
	"encoding/hex"
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"strconv"
	"strings"
)

func Run() {
	lines := fileio.GetLinesFromFile("day4/input.txt")

	var passports []map[string]string
	passports = append(passports, make(map[string]string))
	var pI int
	for _, line := range lines {
		entries := strings.Split(line, " ")

		if entries[0] == "" {
			passports = append(passports, make(map[string]string))
			pI++
			continue
		}

		for _, e := range entries {
			key := strings.Split(e, ":")[0]
			val := strings.Split(e, ":")[1]
			passports[pI][key] = val
		}
	}

	numKeys := 8

	counter := 0
	counter2 := 0
	for _, p := range passports {
		if passportIsValid(p) {
			counter2++
			fmt.Println(p)
		}

		if len(p) == numKeys {
			counter++
			continue
		} else if len(p) == numKeys-1 {
			if p["cid"] == "" {
				counter++
			}
			continue
		}

	}

	fmt.Println("A: ", counter)
	fmt.Println("B: ", counter2)
}

func passportIsValid(p map[string]string) bool {

	numKeys := 8
	if len(p) == numKeys || (len(p) == numKeys-1 && p["cid"] == "") {
		for key, val := range p {
			switch key {
			case "byr":
				if len(val) != 4 {
					return false
				}
				year, err := strconv.Atoi(val)
				if err != nil {
					return false
				}
				if year < 1920 || year > 2002 {
					return false
				}
				break
			case "iyr":
				if len(val) != 4 {
					return false
				}
				year, err := strconv.Atoi(val)
				if err != nil {
					return false
				}
				if year < 2010 || year > 2020 {
					return false
				}
				break
			case "eyr":
				if len(val) != 4 {
					return false
				}
				year, err := strconv.Atoi(val)
				if err != nil {
					return false
				}
				if year < 2020 || year > 2030 {
					return false
				}
				break
			case "hgt":
				num, err := strconv.Atoi(val[:len(val)-2])
				if err != nil {
					return false
				}
				if strings.HasSuffix(val, "in") {
					if num < 59 || num > 76 {
						return false
					}
				} else if strings.HasSuffix(val, "cm") {
					if num < 150 || num > 193 {
						return false
					}
				} else {
					return false
				}
				break
			case "hcl":
				if !strings.HasPrefix(val, "#") || len(val) != 7 {
					return false
				}
				_, err := hex.DecodeString(val[1:])
				if err != nil {
					return false
				}
				break
			case "ecl":
				eyeColor := [...]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				foundit := false
				for _, ec := range eyeColor {
					if val == ec {
						foundit = true
						break
					}
				}
				if foundit == false {
					return false
				}
				break
			case "pid":
				if len(val) != 9 {
					return false
				}
				_, err := strconv.Atoi(val)
				if err != nil {
					return false
				}
				break
			case "cid":
				//Is ignored because it's optional
				break
			}
		}
		return true
	}
	return false
}
