package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

type Coordinate struct {
	x int
	y int
}

type Game struct {
	ButtonA Coordinate
	ButtonB Coordinate
	Prize   Coordinate
}

const MaxButton = 100

func part1(games []Game) {
	fmt.Println("## Part1")

	cost := 0
	for _, game := range games {
		currentPos := Coordinate{0, 0}
		currentCost := math.MaxInt
		for i := 0; i < MaxButton; i++ {
			for j := 0; j <= MaxButton; j++ {
				currentPos.x = i*game.ButtonA.x + j*game.ButtonB.x
				currentPos.y = i*game.ButtonA.y + j*game.ButtonB.y
				if (currentPos.x > game.Prize.x) || currentPos.y > game.Prize.y {
					break
				}
				if currentPos.x == game.Prize.x && currentPos.y == game.Prize.y {
					currentCost = min(currentCost, i*3+j)
					fmt.Printf("%v: has Solution for A:%d and B:%d Cost is:%d\n ", game, i, j, i*3+j)
				}
			}
		}
		if currentCost != math.MaxInt {
			cost += currentCost
		}
	}

	fmt.Printf("Cost for is %d\n", cost)
}

// Cramer's rule
func part1_b(games []Game) {
	fmt.Println("## Part1 b")
	cost := 0
	for _, game := range games {
		nA := (game.Prize.x*game.ButtonB.y - game.Prize.y*game.ButtonB.x) / (game.ButtonA.x*game.ButtonB.y - game.ButtonA.y*game.ButtonB.x)
		nB := (game.ButtonA.x*game.Prize.y - game.ButtonA.y*game.Prize.x) / (game.ButtonA.x*game.ButtonB.y - game.ButtonA.y*game.ButtonB.x)
		if nA*game.ButtonA.x+nB*game.ButtonB.x == game.Prize.x && nA*game.ButtonA.y+nB*game.ButtonB.y == game.Prize.y {
			fmt.Printf("%v: has Solution for A:%d and B:%d Cost is:%d\n ", game, nA, nB, nA*3+nB)
			cost += nA*3 + nB
		}
	}
	fmt.Printf("Cost is %d\n", cost)
}

func part2(lines []string) {
	fmt.Println("## Part2")
}

func gameParser(lines []string) []Game {
	coordinateRule := regexp.MustCompile(`\d+`)

	g := Game{}
	games := []Game{}
	for _, line := range lines {
		matches := coordinateRule.FindAllString(line, -1)
		coordinatePoints := []int{}
		if len(matches) == 2 {
			for _, number := range matches {
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Panic("Could not convert string to int")
				}
				coordinatePoints = append(coordinatePoints, n)
			}
			if strings.Contains(line, "Button A") {
				g.ButtonA = Coordinate{coordinatePoints[0], coordinatePoints[1]}
			} else if strings.Contains(line, "Button B") {
				g.ButtonB = Coordinate{coordinatePoints[0], coordinatePoints[1]}
			} else if strings.Contains(line, "Prize") {
				g.Prize = Coordinate{coordinatePoints[0], coordinatePoints[1]}
				games = append(games, g)
			}
		}
	}
	return games
}

func main() {
	fmt.Println("Starting day")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")

	games := gameParser(lines)
	fmt.Println(len(games))

	part1(games)
	part1_b(games)
}
