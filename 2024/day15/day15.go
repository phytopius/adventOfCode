package main

import (
	"fmt"
	"strings"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

var Direction = map[string][]int{
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
	">": {0, 1},
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

func getRobotPosition(grid [][]string) []int {
	for r, row := range grid {
		for c, col := range row {
			if col == "@" {
				return []int{r, c}
			}
		}
	}
	return []int{}
}

func part1(lines []string) {
	fmt.Println("## Part1")
	gameGrid, orders := parseInput(lines)
	// store current position of robot
	// get next path / direction of robot
	// look at further in direction for objects and walls
	robotPosition := getRobotPosition(gameGrid)

	for _, order := range orders {
		direction := Direction[order]

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

func widenGrid(grid [][]string) [][]string {
	outputGrid := [][]string{}

	for _, row := range grid {
		newRow := []string{}
		for _, col := range row {
			switch col {
			case "#":
				newRow = append(newRow, []string{"#", "#"}...)
			case ".":
				newRow = append(newRow, []string{".", "."}...)
			case "O":
				newRow = append(newRow, []string{"[", "]"}...)
			case "@":
				newRow = append(newRow, []string{"@", "."}...)
			}
		}
		outputGrid = append(outputGrid, newRow)
	}
	return outputGrid
}

func part1_b(lines []string) {

}

func part2(lines []string) {
	fmt.Println("## Part2")

	gameGrid, orders := parseInput(lines)
	widerGameGrid := widenGrid(gameGrid)
	//widerGameGrid = gameGrid
	//adventOfCode.PrintGrid(widerGameGrid)
	//fmt.Println(orders)
	//this is getting a bit more complicated.
	//I cant use the same approach as in part 1

	//current Idea: check for all boxes in move direction
	//store those coordinates
	//check if a wall is in move direction for any box
	//if yes then abort movement
	//if no replace all old positions with "."
	//add direction to all box elements
	//draw new boxes

	robotPosition := getRobotPosition(widerGameGrid)

	for _, order := range orders {
		//adventOfCode.PrintGrid(widerGameGrid)
		//fmt.Printf("Moving %s\n", order)
		direction := Direction[order]
		nextPosition := getNextPosition(robotPosition, direction)
		nextPositionValue := widerGameGrid[nextPosition[0]][nextPosition[1]]
		visited := map[string]bool{}
		if nextPositionValue == "." {
			interchangeValues(robotPosition, nextPosition, widerGameGrid)
			robotPosition = nextPosition
		} else if nextPositionValue == "O" || nextPositionValue == "[" || nextPositionValue == "]" {
			//current problem is the recursion.
			//I am checking the same pieces over and over again because I add the 2nd part of the box
			boxPositions, err := getBoxPositions(widerGameGrid, nextPosition, direction, visited)
			if err != nil {
				fmt.Println(err)
			} else {
				//change robot to dot
				//nextPositions needs to be the robot
				//fmt.Println(boxPositions)
				//fmt.Println("No issues we can move")
				futureBoxPositions := map[string][][]int{}
				//clear and store elements
				for _, box := range boxPositions {
					value := widerGameGrid[box[0]][box[1]]
					futureBoxPositions[value] = append(futureBoxPositions[value], getNextPosition(box, Direction[order]))
					widerGameGrid[box[0]][box[1]] = "."
				}
				//fmt.Println(futureBoxPositions)
				interchangeValues(robotPosition, getNextPosition(robotPosition, Direction[order]), widerGameGrid)
				robotPosition = nextPosition
				//redraw
				for key, positions := range futureBoxPositions {
					for _, pos := range positions {
						widerGameGrid[pos[0]][pos[1]] = key
					}
				}
				//adventOfCode.PrintGrid(widerGameGrid)
				//replace all box positions with dots
				//redraw box positions in + direction
			}
		}

		//fmt.Scanf("doit")
	}
	GPSSum := 0
	for r, row := range widerGameGrid {
		for c, col := range row {
			if col == "O" || col == "[" {
				GPSSum += 100*r + c
			}
		}
	}
	fmt.Printf("Sum is %d\n", GPSSum)
}

func getBoxPositions(grid [][]string, position, direction []int, visited map[string]bool) ([][]int, error) {

	toKey := func(pos []int) string {
		return fmt.Sprintf("%d:%d", pos[0], pos[1])
	}

	key := toKey(position)
	if visited[key] {
		return nil, nil
	}
	visited[key] = true
	returnPositions := [][]int{}
	value := grid[position[0]][position[1]]
	checkPositions := [][]int{}
	switch value {
	case "O":
		returnPositions = append(returnPositions, position)
		nextPosition := getNextPosition(position, direction)
		checkPositions = append(returnPositions, nextPosition)
	case "]":
		returnPositions = append(returnPositions, position)
		position2 := getNextPosition(position, Direction["<"])
		checkPositions = append(checkPositions, position2)
		checkPositions = append(checkPositions, getNextPosition(position, direction))
	case "[":
		returnPositions = append(returnPositions, position)
		position2 := getNextPosition(position, Direction[">"])
		checkPositions = append(checkPositions, position2)
		checkPositions = append(checkPositions, getNextPosition(position, direction))
	case "#":
		return checkPositions, fmt.Errorf("we hit a wall")
	}

	for _, pos := range checkPositions {
		newPositions, err := getBoxPositions(grid, pos, direction, visited)
		if err != nil {
			return checkPositions, fmt.Errorf("we hit a wall")
		}
		returnPositions = append(returnPositions, newPositions...)
	}
	return returnPositions, nil
}
func main() {
	fmt.Println("Starting day")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	part1(lines)
	part2(lines)
}
