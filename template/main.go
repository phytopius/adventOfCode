package main

import (
	"fmt"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

func part1([]string) {
	fmt.Println("## Part1")
}

func part2([]string) {
	fmt.Println("## Part2")
}
func main() {
	fmt.Println("Starting day")
	lines := adventOfCode.ReadFileLineByLine("./testinput.txt")
	part1(lines)
	part2(lines)
}
