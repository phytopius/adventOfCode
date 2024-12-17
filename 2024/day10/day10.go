package main

import (
	"fmt"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

var (
	UP    = []int{-1, 0}
	DOWN  = []int{1, 0}
	LEFT  = []int{0, -1}
	RIGHT = []int{0, 1}
)
var Directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func part1(grid [][]string) {
	fmt.Println("## Part1")

	intGrid := adventOfCode.ConvertStringGridToIntGrid(grid)

	trailheads := [][]int{}

	for r, row := range intGrid {
		for c, v := range row {
			if v == 0 {
				trailheads = append(trailheads, []int{r, c})
			}
		}
	}
	fmt.Println(trailheads)

	ends := findEnds(trailheads, intGrid)
	fmt.Println(ends)

	score := calculateScore(ends)
	fmt.Println("Score is %d\n", score)

}

func uniqueValues[T comparable](input [][]T) [][]T {
	uniqueMap := map[string][]T{}
	for _, element := range input {
		uniqueMap[fmt.Sprintf("%#v", element)] = element
	}

	output := [][]T{}
	for _, element := range uniqueMap {
		output = append(output, element)
	}

	return output
}
func calculateScore(ends map[string][][]int) int {
	score := 0
	for key, trailhead := range ends {
		uniqueEnds := uniqueValues(trailhead)
		fmt.Printf("Ends for %s: %v\n", key, uniqueEnds)
		score += len(uniqueEnds)
	}
	return score
}

func findEnds(trailheads [][]int, grid [][]int) map[string][][]int {
	trailEnds := map[string][][]int{}
	for _, trailhead := range trailheads {
		fmt.Printf("Looking at trailhead %d:%d\n", trailhead[0], trailhead[1])
		ends := [][]int{}
		for _, direction := range Directions {
			fmt.Printf("Going direction %d:%d\n", direction[0], direction[1])
			ends = append(ends, next(trailhead, direction, grid)...)
		}
		trailheadId := fmt.Sprintf("%d:%d", trailhead[0], trailhead[1])
		trailEnds[trailheadId] = ends
	}
	return trailEnds
}

func next(position, direction []int, grid [][]int) [][]int {
	currentValue := grid[position[0]][position[1]]
	//calculate new position
	newRow := position[0] + direction[0]
	newColumn := position[1] + direction[1]
	//check if inbounds
	if newRow < 0 || newRow >= len(grid) || newColumn < 0 || newColumn >= len(grid) {
		fmt.Println("Out of grid!")
		return [][]int{}
	}
	//check if we are increasing by 1 (trail rule)
	newValue := grid[newRow][newColumn]
	if newValue-currentValue != 1 {
		return [][]int{}
	}
	//we reached a trailend
	if newValue == 9 {
		fmt.Println("Reached a 9")
		return [][]int{{newRow, newColumn}}
	}

	ends := [][]int{}
	for _, newDirection := range Directions {
		ends = append(ends, next([]int{newRow, newColumn}, newDirection, grid)...)
	}
	return ends
}

func part2(lines []string) {
	fmt.Println("## Part2")
}
func main() {
	fmt.Println("Starting day")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	grid := adventOfCode.ReadFileAsGrid("./input.txt")
	part1(grid)
	part2(lines)

}
