package main

import (
	"fmt"
	"math"
	"strconv"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

func part1(lines []string) {
	totalSum := 0
	tempSum := 0
	for _, line := range lines {

		num, err := strconv.Atoi(line)
		if err != nil {
			if totalSum < tempSum {
				totalSum = tempSum
			}
			tempSum = 0
			continue
		}
		tempSum += num
	}

	fmt.Printf("Hight Calories: %v\n", totalSum)
}

func minInt(slice []int) (int, int) {
	min := math.MaxInt
	index := 0
	for i, val := range slice {
		if val < min {
			min = val
			index = i
		}
	}
	return min, index
}

func part2(lines []string) {
	topThreeSum := make([]int, 3)
	tempSum := 0
	for lineNumber, line := range lines {
		num, err := strconv.Atoi(line)
		fmt.Println(num)
		if err == nil {
			tempSum += num
		}
		if line == "" || lineNumber == len(lines)-1 {
			fmt.Printf("Current TempSum: %v\n", tempSum)
			currentMin, index := minInt(topThreeSum)
			fmt.Printf("Current min is %v at index: %v\n", currentMin, index)
			if tempSum > currentMin {
				topThreeSum[index] = tempSum
			}
			tempSum = 0
			fmt.Println(topThreeSum)
		}
	}

	totalSum := 0
	for _, element := range topThreeSum {
		totalSum += element
	}

	fmt.Printf("Total Calories is %v\n", totalSum)
}

func main() {
	fmt.Println("Starting day 1")

	lines := adventOfCode.ReadFileLineByLine("./input.txt")

	part1(lines)
	part2(lines)
}
