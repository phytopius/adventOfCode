package main

import (
	"fmt"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

type DIRECTION int

const (
	UP DIRECTION = iota
	RIGHT
	DOWN
	LEFT
)

func convertGridToCoordinateMap(grid [][]string) map[[2]int]string {
	mapGrid := make(map[[2]int]string)
	for i := range len(grid) {
		for j := range len(grid[i]) {
			mapGrid[[2]int{i, j}] = grid[i][j]
		}
	}
	return mapGrid
}

func findStartPosition(grid [][]string, target string) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == target {
				return i, j
			}
		}
	}
	return -1, -1
}
func part1(grid [][]string) {
	fmt.Println("## Part1")

	//lets try to use a map for faster lookup

	gameMap := convertGridToCoordinateMap(grid)

	startRow, startCol := findStartPosition(grid, "^")
	fmt.Printf("Start Position is %d, %d\n", startRow, startCol)

	charOnMap := true
	directions := ([4]string{"up", "right", "down", "left"})
	directionPosition := 0
	direction := "up"
	nextRow, nextCol := startRow, startCol

	uniqueFieldsVisited := 1
	gameMap[[2]int{startRow, startCol}] = "X"
	for charOnMap {
		currentRow := nextRow
		currentCol := nextCol
		nextFieldValue := ""
		switch direction {
		case "up":
			nextRow--
		case "down":
			nextRow++
		case "left":
			nextCol--
		case "right":
			nextCol++
		}
		if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
			charOnMap = false
			break
		}
		nextFieldValue = gameMap[[2]int{nextRow, nextCol}]
		//fmt.Printf("Moving %s: New Position is %d,%d with value %s\n", direction, nextRow, nextCol, nextFieldValue)
		switch nextFieldValue {
		case ".":
			uniqueFieldsVisited++
			gameMap[[2]int{nextRow, nextCol}] = "X"
		case "#":
			directionPosition++
			if directionPosition == 4 {
				directionPosition = 0
			}
			direction = directions[directionPosition]
			nextRow = currentRow
			nextCol = currentCol
		}
	}
	fmt.Printf("Number of visited unique fields %d\n", uniqueFieldsVisited)
}

func part2(gird [][]string) {
	fmt.Println("## Part2")
}
func main() {
	fmt.Println("Starting day")
	grid := adventOfCode.ReadFileAsGrid("./input.txt")
	part1(grid)
	part2(grid)
}
