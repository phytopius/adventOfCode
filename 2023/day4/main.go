package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
	z int
}
type Hailstone struct {
	position  Coordinate
	direction Coordinate
}

func main() {
	fmt.Println("Starting day 4")

	file, err := os.Open("./testinput.txt")
	if err != nil {
		log.Fatal("Error during file open")
	}

	hailstones := make([]Hailstone, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		line := scanner.Text()
		splitlineStrings := strings.Fields(line)
		splitlineInt := make([]int, len(splitlineStrings))
		for i, element := range splitlineStrings {
			if i == 3 {
				continue
			}
			element = strings.ReplaceAll(element, ",", "")
			num, err := strconv.Atoi(element)
			if err != nil {
				fmt.Printf("Error converting %s to int: %v\n", element, err)
				return
			}
			splitlineInt[i] = num
		}
		position := Coordinate{x: splitlineInt[0], y: splitlineInt[1], z: splitlineInt[2]}
		direction := Coordinate{x: splitlineInt[4], y: splitlineInt[5], z: splitlineInt[6]}

		h := Hailstone{position: position, direction: direction}
		hailstones = append(hailstones, h)
	}

	for _, element := range hailstones {
		fmt.Println(element)
	}

}
