package adventOfCode

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func ReadFileAsGrid(path string) [][]string {
	lines := ReadFileLineByLine(path)

	grid := [][]string{}

	for _, line := range lines {
		tempRow := []string{}
		for _, element := range line {
			tempRow = append(tempRow, string(element))
		}
		grid = append(grid, tempRow)
	}
	return grid
}
