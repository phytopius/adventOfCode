package main

import (
	"fmt"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

func part1(input []string) {
	fmt.Println("## Part1")

	dd := [][]int{}
	//create coordinates
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			dd = append(dd, []int{dx, dy})
		}
	}

	grid := [][]string{}
	for _, row := range input {
		tempRow := []string{}
		for _, col := range row {
			tempRow = append(tempRow, string(col))
		}
		grid = append(grid, tempRow)
	}

	xmasCounter := 0
	numRows := len(input)
	numChars := len(input[0])
	for i := range numRows {
		for j := range numChars {
			for _, dir := range dd {
				//fmt.Printf("Direction %v\n", dir)
				var word = ""
				for n := 0; n < 4; n++ {
					dx := i + dir[0]*n
					dy := j + dir[1]*n
					if dx >= 0 && dx < numRows && dy >= 0 && dy < numChars {
						//fmt.Printf("(%d, %d) - %s\n", dx, dy, grid[dx][dy])
						word += grid[dx][dy]
						if word == "XMAS" {
							xmasCounter++
						}
					} else {
						break
					}
				}
				//fmt.Println(word)
			}
		}
	}
	fmt.Println(xmasCounter)
}

func part2(input []string) {
	fmt.Println("## Part2")

	grid := [][]string{}
	for _, row := range input {
		tempRow := []string{}
		for _, col := range row {
			tempRow = append(tempRow, string(col))
		}
		grid = append(grid, tempRow)
	}

	diagonalX := [][]int{}
	diagonalX = append(diagonalX, []int{-1, -1})
	diagonalX = append(diagonalX, []int{0, 0})
	diagonalX = append(diagonalX, []int{1, 1})

	diagonalY := [][]int{}
	diagonalY = append(diagonalY, []int{-1, 1})
	diagonalY = append(diagonalY, []int{0, 0})
	diagonalY = append(diagonalY, []int{1, -1})

	masCounter := 0
	numRows := len(input)
	numChars := len(input[0])
	for i := range numRows {
		for j := range numChars {
			//fmt.Printf(grid[i][j])
			if grid[i][j] != "A" {
				continue
			}

			word3 := ""
			word4 := ""
			//diagonal
			for _, elem := range diagonalX {
				dx := i + elem[0]
				dy := j + elem[1]
				if dx >= 0 && dx < numRows && dy >= 0 && dy < numChars {
					//fmt.Printf("(%d, %d) - %s\n", dx, dy, grid[dx][dy])
					word3 += grid[dx][dy]
				}
			}

			for _, elem := range diagonalY {
				dx := i + elem[0]
				dy := j + elem[1]
				if dx >= 0 && dx < numRows && dy >= 0 && dy < numChars {
					//fmt.Printf("(%d, %d) - %s\n", dx, dy, grid[dx][dy])
					word4 += grid[dx][dy]
				}
			}

			if (word3 == "MAS" || word3 == "SAM") && (word4 == "MAS" || word4 == "SAM") {
				masCounter++
			}
		}
	}
	fmt.Println(masCounter)
}
func main() {
	fmt.Println("Starting day 4")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	part1(lines)
	part2(lines)
}
