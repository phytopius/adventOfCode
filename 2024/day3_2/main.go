package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func multiplier(input string) int {
	numFinder := regexp.MustCompile(`\d{1,3}`)
	matches := numFinder.FindAllString(input, -1)
	numbersSlice := make([]int, 2)

	for i, numberString := range matches {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal("Error during number parsing")
		}
		numbersSlice[i] = number
	}

	result := numbersSlice[0] * numbersSlice[1]
	return result
}
func main() {
	fmt.Println("Starting day 3")

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error during file read")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sum = 0
	mulEnabled := true
	for scanner.Scan() {
		line := scanner.Text()
		mulFinder := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
		matches := mulFinder.FindAllString(line, -1)

		fmt.Println(matches)
		for _, match := range matches {
			switch match {
			case "do()":
				//fmt.Println("Found a do")
				mulEnabled = true
			case "don't()":
				//fmt.Println("FOund a don't")
				mulEnabled = false
			default:
				if mulEnabled == true {
					sum += multiplier(match)
				}
			}
		}
	}
	fmt.Printf("Sum is %v\n", sum)
}
