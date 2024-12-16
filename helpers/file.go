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

func ReadFileAsMap(path string) map[[2]int]string {
	lines := ReadFileLineByLine(path)

	mapGrid := make(map[[2]int]string)

	for i, line := range lines {
		for j, element := range line {
			mapGrid[[2]int{i, j}] = string(element)
		}
	}
	return mapGrid
}

func convertGridToCoordinateMap(grid [][]string) map[[2]int]string {
	mapGrid := make(map[[2]int]string)
	for i := range len(grid) {
		for j := range len(grid[i]) {
			mapGrid[[2]int{i, j}] = grid[i][j]
		}
	}
	return mapGrid
}
