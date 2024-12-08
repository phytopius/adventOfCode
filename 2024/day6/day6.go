package main

import (
	"fmt"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

type DIRECTION int

const (
	UP = iota
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
	directions := ([4]int{UP, RIGHT, DOWN, LEFT})
	directionPosition := 0
	direction := UP
	nextRow, nextCol := startRow, startCol

	uniqueFieldsVisited := 1
	gameMap[[2]int{startRow, startCol}] = "X"
	for charOnMap {
		//store wher we currently are
		currentRow := nextRow
		currentCol := nextCol
		nextFieldValue := ""
		switch direction {
		case UP:
			nextRow--
		case DOWN:
			nextRow++
		case LEFT:
			nextCol--
		case RIGHT:
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
			directionPosition = (directionPosition + 1) % len(directions)
			direction = directions[directionPosition]
			nextRow = currentRow
			nextCol = currentCol
		}
	}
	fmt.Printf("Number of visited unique fields %d\n", uniqueFieldsVisited)
}

func part2(grid [][]string) {
	fmt.Println("## Part2")
	gameMap := convertGridToCoordinateMap(grid)

	startRow, startCol := findStartPosition(grid, "^")
	fmt.Printf("Start Position is %d, %d\n", startRow, startCol)

	charOnMap := true
	directions := ([4]int{UP, RIGHT, DOWN, LEFT})
	directionMarkers := [4]string{"U", "R", "D", "L"}
	directionPosition := 0
	direction := UP
	nextRow, nextCol := startRow, startCol
	loopPositions := [][2]int{}

	uniqueFieldsVisited := 1
	gameMap[[2]int{startRow, startCol}] = "U"
	grid[startRow][startCol] = "U"
	for charOnMap {
		adventOfCode.PrintGrid(grid)
		fmt.Printf("\n\n")
		//store wher we currently are
		currentRow := nextRow
		currentCol := nextCol
		nextFieldValue := ""
		switch direction {
		case UP:
			nextRow--
		case DOWN:
			nextRow++
		case LEFT:
			nextCol--
		case RIGHT:
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
			gameMap[[2]int{nextRow, nextCol}] = directionMarkers[directionPosition]
			grid[nextRow][nextCol] = directionMarkers[directionPosition]
		case "#":
			directionPosition = (directionPosition + 1) % len(directions)
			direction = directions[directionPosition]
			nextRow = currentRow
			nextCol = currentCol

			gameMap[[2]int{nextRow, nextCol}] = "+"
			grid[nextRow][nextCol] = "+"
			//highly experimental: If i am on a field where I already was. Put an obstacle in front of it
			//check if the next field is the same than the one I would put on it i assume its a loop
		case "U":
			directionAfterObstacle := directions[(directionPosition+1)%len(directions)]
			testRow := nextRow
			testCol := nextCol
			switch directionAfterObstacle {
			case UP:
				testRow--
			case DOWN:
				testRow++
			case LEFT:
				testCol--
			case RIGHT:
				testCol++
			}
			testFieldValue := gameMap[[2]int{testRow, testCol}]
			//found a loop position.. Fully reset and start again.
			//Ignore already found loop position
			//Do this until i exit the arena
			if testFieldValue == directionMarkers[(directionPosition+1)%len(directions)] || testFieldValue == "+" {
				loopPositions = append(loopPositions, [2]int{nextRow, nextCol})
				fmt.Printf("Found a loop Position at %d, %d\n", testRow, testCol)
			}
			gameMap[[2]int{nextRow, nextCol}] = "+"
			grid[nextRow][nextCol] = "+"
		case "D":
			directionAfterObstacle := directions[(directionPosition+1)%len(directions)]
			testRow := nextRow
			testCol := nextCol
			switch directionAfterObstacle {
			case UP:
				testRow--
			case DOWN:
				testRow++
			case LEFT:
				testCol--
			case RIGHT:
				testCol++
			}
			testFieldValue := gameMap[[2]int{testRow, testCol}]
			//found a loop position.. Fully reset and start again.
			//Ignore already found loop position
			//Do this until i exit the arena
			if testFieldValue == directionMarkers[(directionPosition+1)%len(directions)] || testFieldValue == "+" {
				loopPositions = append(loopPositions, [2]int{nextRow, nextCol})
				fmt.Printf("Found a loop Position at %d, %d\n", testRow, testCol)
			}
			gameMap[[2]int{nextRow, nextCol}] = "+"
			grid[nextRow][nextCol] = "+"
		case "L":
			directionAfterObstacle := directions[(directionPosition+1)%len(directions)]
			testRow := nextRow
			testCol := nextCol
			switch directionAfterObstacle {
			case UP:
				testRow--
			case DOWN:
				testRow++
			case LEFT:
				testCol--
			case RIGHT:
				testCol++
			}
			testFieldValue := gameMap[[2]int{testRow, testCol}]
			//found a loop position.. Fully reset and start again.
			//Ignore already found loop position
			//Do this until i exit the arena
			if testFieldValue == directionMarkers[(directionPosition+1)%len(directions)] || testFieldValue == "+" {
				loopPositions = append(loopPositions, [2]int{nextRow, nextCol})
				fmt.Printf("Found a loop Position at %d, %d\n", testRow, testCol)
			}
			gameMap[[2]int{nextRow, nextCol}] = "+"
			grid[nextRow][nextCol] = "+"
		case "R":
			directionAfterObstacle := directions[(directionPosition+1)%len(directions)]
			testRow := nextRow
			testCol := nextCol
			switch directionAfterObstacle {
			case UP:
				testRow--
			case DOWN:
				testRow++
			case LEFT:
				testCol--
			case RIGHT:
				testCol++
			}
			testFieldValue := gameMap[[2]int{testRow, testCol}]
			//found a loop position.. Fully reset and start again.
			//Ignore already found loop position
			//Do this until i exit the arena
			if testFieldValue == directionMarkers[(directionPosition+1)%len(directions)] || testFieldValue == "+" {
				loopPositions = append(loopPositions, [2]int{nextRow, nextCol})
				fmt.Printf("Found a loop Position at %d, %d\n", testRow, testCol)
			}
			gameMap[[2]int{nextRow, nextCol}] = "+"
			grid[nextRow][nextCol] = "+"
		}
	}
	fmt.Printf("Number of visited unique fields %d\n", uniqueFieldsVisited)
	fmt.Println(loopPositions)
	fmt.Printf("Length of loop positions is %d\n", len(loopPositions))
}

func checkIfIntersectsWithWallAndPlusInFront(gameMap map[[2]int]string, currentPos [2]int, dir int, maxRows int, maxCols int) int {
	directions := map[int][2]int{
		UP:    {-1, 0},
		RIGHT: {0, 1},
		DOWN:  {1, 0},
		LEFT:  {0, -1},
	}

	//fmt.Printf("Starting Position Row:%d, Col:%d. Direction is %d\n", currentPos[0], currentPos[1], dir)
	doIt := true
	i := 1
	previousChar := ""
	for doIt {
		currentChar := ""
		newRow := currentPos[0] + directions[dir][0]*i
		newCol := currentPos[1] + directions[dir][1]*i
		if newCol >= 0 && newRow >= 0 && newCol < maxCols && newRow < maxRows {
			currentChar = gameMap[[2]int{newRow, newCol}]
			//fmt.Printf("I am checking: Row:%d, Col:%d. Content is %s: Previous Content is %s\n", newRow, newCol, currentChar, previousChar)
			if currentChar == "#" {
				//fmt.Printf("Path intersects with a wall\n")
				if previousChar == "+" {
					//fmt.Printf("Also Previous char was a +. So this is loop")
					return 1
				} else {
					break
				}
			}
		} else {
			break
		}
		i++
		previousChar = currentChar
	}
	return 0
}

func part2_a(grid [][]string) {
	fmt.Println("## Part2")
	gameMap := convertGridToCoordinateMap(grid)
	maxRows := len(grid)
	maxCols := len(grid[0])

	startRow, startCol := findStartPosition(grid, "^")
	fmt.Printf("Start Position is %d, %d\n", startRow, startCol)

	charOnMap := true
	directions := ([4]int{UP, RIGHT, DOWN, LEFT})
	directionMarkers := [4]string{"|", "-", "|", "-"}
	directionPosition := 0
	direction := UP
	nextRow, nextCol := startRow, startCol
	loopPositionCounter := 0

	uniqueFieldsVisited := 1
	gameMap[[2]int{startRow, startCol}] = "|"
	grid[startRow][startCol] = "|"
	for charOnMap {
		//adventOfCode.PrintGrid(grid)
		//fmt.Printf("\n\n")
		//store wher we currently are
		currentRow := nextRow
		currentCol := nextCol

		turnDirection := (direction + 1) % 4
		loopPositionCounter += checkIfIntersectsWithWallAndPlusInFront(gameMap, [2]int{currentRow, currentCol}, turnDirection, maxRows, maxCols)

		nextFieldValue := ""
		switch direction {
		case UP:
			nextRow--
		case DOWN:
			nextRow++
		case LEFT:
			nextCol--
		case RIGHT:
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
			gameMap[[2]int{nextRow, nextCol}] = directionMarkers[directionPosition]
			grid[nextRow][nextCol] = directionMarkers[directionPosition]
		case "#":
			directionPosition = (directionPosition + 1) % len(directions)
			direction = directions[directionPosition]
			nextRow = currentRow
			nextCol = currentCol

			gameMap[[2]int{nextRow, nextCol}] = "+"
			grid[nextRow][nextCol] = "+"
		}

	}
	adventOfCode.PrintGrid(grid)
	fmt.Printf("Found %d loop positions\n", loopPositionCounter)
}

func main() {
	fmt.Println("Starting day")
	grid := adventOfCode.ReadFileAsGrid("./input.txt")
	part1(grid)
	//part2(grid)
	part2_a(grid)
}
