package main

import (
	"fmt"
	"strconv"
	"strings"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

func getDigitCount(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}

func part1(lines []string) {
	fmt.Println("## Part1")
	inputString := strings.Split(lines[0], " ")
	input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)

	fmt.Println(input)

	numBlinks := 25

	for i := 0; i < numBlinks; i++ {
		for j := 0; j < len(input); j++ {
			digitCount := getDigitCount(input[j])
			//0 case
			if input[j] == 0 {
				input[j] = 1
			} else if digitCount%2 == 0 {
				//fmt.Printf("Even number of digits at index %d\n", j)
				numberAsString := strconv.Itoa(input[j])
				midpoint := digitCount / 2
				num1, _ := strconv.Atoi(numberAsString[0:midpoint])
				num2, _ := strconv.Atoi(numberAsString[midpoint:])
				input[j] = num1
				input = append(input[0:j+1], append([]int{num2}, input[j+1:]...)...)

				j++
			} else {
				input[j] = input[j] * 2024
			}
		}
	}
	fmt.Printf("Number of Stones %d\n", len(input))
}

func part2(lines []string) {
	fmt.Println("## Part2")
}
func main() {
	fmt.Println("Starting day")

	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	part1(lines)
	part2(lines)
}
