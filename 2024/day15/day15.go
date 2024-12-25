package main

import (
	"fmt"
	"strings"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

var Direction = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func parseInput(lines []string) ([][]string, []string) {
	gameField := [][]string{}
	inputs := []string{}
	for _, line := range lines {
		splitLine := strings.Split(line, "")
		if len(splitLine) > 0 {
			if splitLine[0] == "#" {
				gameField = append(gameField, splitLine)
			} else {
				inputs = append(inputs, splitLine...)
			}
		}
	}
	return gameField, inputs
}

func part1(lines []string) {
	fmt.Println("## Part1")
	gameGrid, orders := parseInput(lines)
	// store current position of robot
	// get next path / direction of robot
	// look at further in direction for objects and walls

	robotPosition := []int{}

	//find the initial position of the robot
outer:
	for r, row := range gameGrid {
		for c, col := range row {
			if col == "@" {
				robotPosition = []int{r, c}
				break outer
			}
		}
	}

	for _, order := range orders {
		direction := []int{}
		switch order {
		case "^":
			direction = Direction[0]
		case "v":
			direction = Direction[1]
		case "<":
			direction = Direction[2]
		case ">":
			direction = Direction[3]
		}

		//continue to go into the direction until we hit a wall
		//fmt.Println("Grid before")
		//adventOfCode.PrintGrid(gameGrid)
		nextPos := getNextPosition(robotPosition, direction)
		nextPosValue := gameGrid[nextPos[0]][nextPos[1]]
		//it it is a dot just do a simple switch update robot position
		if nextPosValue == "." {
			interchangeValues(robotPosition, nextPos, gameGrid)
			robotPosition = nextPos
			//if it is an O continue to look for the next one
		} else if nextPosValue == "O" {
			initialOPosition := nextPos
			walkingPosition := nextPos
			for {
				nextPosition := getNextPosition(walkingPosition, direction)
				nextPositionValue := gameGrid[nextPosition[0]][nextPosition[1]]
				if nextPositionValue == "#" {
					//do nothing
					break
					//walk until we hit a dot.
					//change the dot and the initial O
					//change the new dot and the robot
					//update robot position
				} else if nextPositionValue == "." {
					interchangeValues(initialOPosition, nextPosition, gameGrid)
					interchangeValues(robotPosition, initialOPosition, gameGrid)
					robotPosition = initialOPosition
					break
				}
				walkingPosition = nextPosition
			}
		} else {
			// do nothing we hit a wall "#"
		}

		//fmt.Println("Grid after")
		//adventOfCode.PrintGrid(gameGrid)
	}

	GPSSum := 0
	for r, row := range gameGrid {
		for c, col := range row {
			if col == "O" {
				GPSSum += 100*r + c
			}
		}
	}
	fmt.Printf("Sum is %d\n", GPSSum)
}
func interchangeValues(position1, position2 []int, gameGrid [][]string) {
	gameGrid[position1[0]][position1[1]], gameGrid[position2[0]][position2[1]] = gameGrid[position2[0]][position2[1]], gameGrid[position1[0]][position1[1]]
}

func getNextPosition(position, direction []int) []int {
	return []int{position[0] + direction[0], position[1] + direction[1]}
}

func part2(lines []string) {
	fmt.Println("## Part2")
}
func main() {
	fmt.Println("Starting day")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	part1(lines)
	part2(lines)
}
