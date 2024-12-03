package main

import (
	"fmt"
	"strings"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

const (
	Win  = 6
	Loss = 0
	Draw = 3

	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func part1(lines []string) {
	fmt.Println("## Part1")

	choices := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	totalScore := 0
	for _, line := range lines {
		fmt.Println("------")
		plays := strings.Split(line, " ")
		matchResult := choices[plays[0]] - choices[plays[1]]
		fmt.Printf("Played %v\n", choices[plays[1]])
		switch {
		case matchResult == 0:
			//Draw
			fmt.Println("Draw")
			totalScore += Draw + choices[plays[1]]
		case matchResult == -2 || matchResult == 1:
			//P2 lose
			fmt.Println("Loss")
			totalScore += Loss + choices[plays[1]]
		case matchResult == 2 || matchResult == -1:
			//P2 Win
			fmt.Println("Win")
			totalScore += Win + choices[plays[1]]
		}
		fmt.Printf("Score: %d\n", totalScore)
	}
	fmt.Printf("Total score is %v\n", totalScore)
}

func part2(lines []string) {
	fmt.Println("## Part2")

	choices := map[string]int{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Loss,
		"Y": Draw,
		"Z": Win,
	}

	totalScore := 0
	for _, line := range lines {
		plays := strings.Split(line, " ")
		switch choices[plays[1]] {
		case Win:
			totalScore += Win
			switch choices[plays[0]] {
			case Rock:
				totalScore += Paper
			case Paper:
				totalScore += Scissors
			case Scissors:
				totalScore += Rock
			}
		case Loss:
			switch choices[plays[0]] {
			case Rock:
				totalScore += Scissors
			case Paper:
				totalScore += Rock
			case Scissors:
				totalScore += Paper
			}
		case Draw:
			totalScore += Draw + choices[plays[0]]
		}
	}
	fmt.Printf("TotalScore is %v\n", totalScore)
}

func main() {
	fmt.Println("Starting day 2")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	part1(lines)
	part2(lines)
}
