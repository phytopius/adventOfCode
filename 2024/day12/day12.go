package main

import (
	"fmt"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

var Direction = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func inBounds(row, col, size int) bool {
	if row < 0 || col < 0 || row >= size || col >= size {
		return false
	}
	return true
}

// first approach which was done in a few minutes.
// when running testinput 2 I learned that this will not work if multiple plots have the same plant :(
// this one is wrong
func part1(grid [][]string) {
	fmt.Println("## Part1")
	resultMap := map[string][]int{}
	for r, row := range grid {
		for c, col := range row {
			fences := 0
			for _, dir := range Direction {
				checkRow := r + dir[0]
				checkCol := c + dir[1]
				if !inBounds(checkRow, checkCol, len(row)) {
					fences += 1
				} else if grid[checkRow][checkCol] != col {
					fences += 1
				}
			}
			resultMap[col] = append(resultMap[col], fences)
		}
	}
	price := 0
	for plant, plot := range resultMap {
		fences := 0
		for _, fence := range plot {
			fences += fence
		}
		fmt.Printf("A region of %s plants with price %d * %d = %d\n", plant, len(plot), fences, fences*len(plot))
		price += fences * len(plot)
	}
	fmt.Printf("Final price is %d\n", price)
}

// solved with a resursive function. Was good practice and works :)
// I still kind of don't like this though. was hoping for something simpler
func part1_a(grid [][]string) {
	fmt.Println("## Part1_a")

	visited := make(map[string]bool)

	price := 0

	regions := [][]int{}
	for r, row := range grid {
		for c, _ := range row {
			region := next(r, c, grid, visited)
			if len(region) > 0 {
				regions = append(regions, region)
			}
		}
	}
	for _, region := range regions {
		fences := 0
		for _, plant := range region {
			fences += plant
		}
		price += fences * len(region)
	}
	fmt.Printf("Final price is %d\n", price)
}

func next(row, col int, grid [][]string, visited map[string]bool) []int {
	region := []int{}
	fences := 0
	if visited[fmt.Sprintf("%d:%d", row, col)] {
		return []int{}
	}
	visited[fmt.Sprintf("%d:%d", row, col)] = true
	for _, dir := range Direction {
		checkRow := row + dir[0]
		checkCol := col + dir[1]
		if !inBounds(checkRow, checkCol, len(grid)) {
			fences += 1
		} else if grid[checkRow][checkCol] != grid[row][col] {
			fences += 1
		} else {
			region = append(region, next(checkRow, checkCol, grid, visited)...)
		}
	}
	region = append(region, []int{fences}...)
	return region
}

func part2(grid [][]string) {
	fmt.Println("## Part2")
}
func main() {
	fmt.Println("Starting day")
	grid := adventOfCode.ReadFileAsGrid("./input.txt")
	part1(grid)
	part1_a(grid)
	part2(grid)
}
