package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	rounds []Round
}

func (g *Game) isGamevalid(validationRound Round) int {

	maxCounts := map[string]int{}
	for _, vc := range validationRound.colorCounts {
		maxCounts[vc.color] = vc.amount
	}

	for _, round := range g.rounds {
		for _, cc := range round.colorCounts {
			maxAllowed, exists := maxCounts[cc.color]
			if !exists {
				return 0
			}
			if cc.amount > maxAllowed {
				return 0
			}
		}
	}
	return g.id

}
func (g *Game) getMinimumCubePower() int {
	colors := [3]string{"red", "blue", "green"}
	power := 1
	for _, c := range colors {
		minPerColor := 1
		for _, round := range g.rounds {
			for _, cc := range round.colorCounts {
				minNeeded, exists := cc.getAmountForColor(c)
				if exists {
					minPerColor = max(minNeeded, minPerColor)
				}
			}
		}
		fmt.Printf("Min Per Color for %v is %v\n", c, minPerColor)
		power *= minPerColor
	}
	return power
}

type Round struct {
	colorCounts []ColorCount
}

type ColorCount struct {
	color  string
	amount int
}

func (c ColorCount) getAmountForColor(color string) (int, bool) {
	if color == c.color {
		return c.amount, true
	}
	return 1, false
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func parseGame(line string) (Game, error) {
	elements := strings.Split(line, ":")
	roundNumber, err := strconv.Atoi(strings.Split(elements[0], " ")[1])
	check(err)

	//fmt.Printf("Roundnumber: %v\n", roundNumber)

	rounds := []Round{}
	roundsString := strings.Split(elements[1], ";")
	for _, round := range roundsString {
		colorSet := strings.Split(round, ",")
		fmt.Printf("ColorSet: %v\n", colorSet)
		fmt.Println(len(colorSet))

		colorCounts := []ColorCount{}
		for _, color := range colorSet {
			fmt.Println(color)
			splitter := strings.Split(strings.TrimLeft(color, " "), " ")
			fmt.Println(splitter)
			fmt.Printf("number is : %v\n", splitter[0])
			colorCount, err := strconv.Atoi(splitter[0])

			check(err)
			colorCounts = append(colorCounts, ColorCount{amount: colorCount, color: splitter[1]})
		}
		rounds = append(rounds, Round{colorCounts: colorCounts})
	}
	game := Game{id: roundNumber, rounds: rounds}
	return game, nil
}

func main() {
	fmt.Println("Starting day 2")
	// Example usage
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//var games []Round
	scanner := bufio.NewScanner(file)

	games := []Game{}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		game, err := parseGame(scanner.Text())
		check(err)
		games = append(games, game)
	}

	validation := Round{
		colorCounts: []ColorCount{
			{color: "red", amount: 12},
			{color: "blue", amount: 14},
			{color: "green", amount: 13},
		},
	}

	fmt.Println(validation)

	validGames := 0
	totalPower := 0
	for i, game := range games {

		fmt.Println(game)
		validGames += game.isGamevalid(validation)
		fmt.Printf("Game Power is: %v\n", game.getMinimumCubePower())
		totalPower += game.getMinimumCubePower()
		if validGames > 0 {
			fmt.Printf("Game %v is valid\n", i+1)
		}
	}

	fmt.Printf("ValidGames Number %v\n", validGames)
	fmt.Printf("Power: %v\n", totalPower)
}
