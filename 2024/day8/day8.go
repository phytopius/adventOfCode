package main

import (
	"fmt"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

func isPositionValid(x, y, size int) bool {
	if x < 0 || x >= size {
		return false
	}
	if y < 0 || y >= size {
		return false
	}
	return true
}

func part1(grid [][]string) {
	fmt.Println("## Part1")

	antennaBuckets := make(map[string][][2]int)
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] != "." {
				antennaBuckets[grid[i][j]] = append(antennaBuckets[grid[i][j]], [2]int{i, j})
			}
		}
	}
	fmt.Println(antennaBuckets)

	antennaMap := make(map[[2]int]bool)
	gridSize := len(grid)
	for key, value := range antennaBuckets {
		fmt.Printf("Handling key %s\n", key)
		for i := 0; i < len(value); i++ {
			for j := 1; j < len(value); j++ {
				if i != j {
					ax, ay := value[i][0], value[i][1]
					bx, by := value[j][0], value[j][1]
					dx, dy := bx-ax, by-ay

					ex, ey := ax-dx, ay-dy
					fx, fy := bx+dx, by+dy
					fmt.Printf("Got a Pair\n")
					if isPositionValid(ex, ey, gridSize) {
						antennaMap[[2]int{ex, ey}] = true
						fmt.Printf("Adding antenna 1 at %d, %d\n", ex, ey)
					}
					if isPositionValid(fx, fy, gridSize) {
						antennaMap[[2]int{fx, fy}] = true
						fmt.Printf("Adding antenna 2 at %d, %d\n", fx, fy)
					}
				}
			}
		}
	}
	fmt.Println("########")
	fmt.Println(antennaMap)
	fmt.Printf("Found %d unique antenna positions\n", len(antennaMap))
}

func part2(lines []string) {
	fmt.Println("## Part2")
}
func main() {
	fmt.Println("Starting day")
	inputFilePath := "./input.txt"
	lines := adventOfCode.ReadFileLineByLine(inputFilePath)
	mapGrid := adventOfCode.ReadFileAsGrid(inputFilePath)
	part1(mapGrid)
	part2(lines)
}
