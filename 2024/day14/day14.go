package main

import (
	"fmt"
	"log"
	"strconv"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

const maxX = 101
const maxY = 103
const ITERATIONS = 100

type Coordinate struct {
	x int
	y int
}

type Robot struct {
	position Coordinate
	velocity Coordinate
}

func robotParserScanf(lines []string) ([]Robot, error) {
	robots := []Robot{}
	for _, line := range lines {
		var robot Robot
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.position.x, &robot.position.y, &robot.velocity.x, &robot.velocity.y)
		if err != nil {
			return nil, fmt.Errorf("failed to parse robot '%s', %v", line, err)
		}
		robots = append(robots, robot)
	}
	return robots, nil
}

func getSectorPositions(number int) (sector1End, sector2Start int) {
	numberMod := number % 2
	if numberMod == 0 {
		sector1End = number/2 - 1
		sector2Start = number / 2
	} else {
		sector1End = number/2 - 1
		sector2Start = number/2 + 1
	}
	return sector1End, sector2Start
}

func part1(robots []Robot) {
	fmt.Println("## Part1")

	sector1EndX, sector2StartX := getSectorPositions(maxX)
	sector1EndY, sector2StartY := getSectorPositions(maxY)

	for i := 0; i < ITERATIONS; i++ {
		//fmt.Println(robots)
		//drawRobots(robots, false)
		for j, robot := range robots {
			newX := robot.position.x + robot.velocity.x
			newY := robot.position.y + robot.velocity.y
			if newX < 0 {
				newX = maxX + newX
			}
			if newX >= maxX {
				newX = newX - maxX
			}
			if newY < 0 {
				newY = maxY + newY
			}
			if newY >= maxY {
				newY = newY - maxY
			}
			robot.position.x = newX
			robot.position.y = newY
			robots[j] = robot
		}
	}
	//fmt.Println(robots)
	sectors := []int{0, 0, 0, 0}
	for _, robot := range robots {
		if robot.position.x <= sector1EndX {
			if robot.position.y <= sector1EndY {
				sectors[0]++
			} else if robot.position.y >= sector2StartY {
				sectors[1]++
			}
		} else if robot.position.x >= sector2StartX {
			if robot.position.y <= sector1EndY {
				sectors[2]++
			} else if robot.position.y >= sector2StartY {
				sectors[3]++
			}
		}
	}
	fmt.Println(sectors)
	safetyFactor := 1
	for _, sector := range sectors {
		safetyFactor *= sector
	}
	fmt.Printf("Safety Factor is %d\n", safetyFactor)

	//drawRobots(robots, true)
}

func part2(lines []string) {
	fmt.Println("## Part2")
}

func drawRobots(robots []Robot, sectors bool) {
	grid := [][]string{}
	for i := 0; i < maxY; i++ {
		row := []string{}
		for j := 0; j < maxX; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}
	for _, robot := range robots {
		value := grid[robot.position.y][robot.position.x]
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			grid[robot.position.y][robot.position.x] = "1"
		}
		valueInt = valueInt + 1
		valueString := strconv.Itoa(valueInt)
		grid[robot.position.y][robot.position.x] = valueString
	}
	if !sectors {
		for _, row := range grid {
			fmt.Println(row)
		}
	} else {
		drawSectors(grid)
	}
}
func drawSectors(grid [][]string) {
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if i == maxY/2 || j == maxX/2 {
				grid[i][j] = " "
			}
		}
	}
	for _, row := range grid {
		fmt.Println(row)
	}

}
func main() {
	fmt.Println("Starting day")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	robots, err := robotParserScanf(lines)
	if err != nil {
		log.Panic("Error during robot parsing")
	}
	part1(robots)
	part2(lines)
}
