package fileio

import (
	"bufio"
	"log"
	"os"
)

//GetLinesFromFile returns the lines from a file in an array
func GetLinesFromFile(filename string) []string {

	var fileLines []string

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		fileLines = append(fileLines, s)

	}

	return fileLines

}
