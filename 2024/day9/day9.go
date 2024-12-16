package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

func findLastNonDot(s string) (int, error) {
	re, err := regexp.Compile(`[^.]`)
	if err != nil {
		return -1, err
	}
	matches := re.FindAllStringIndex(s, -1)
	if len(matches) == 0 {
		return -1, err
	}
	return matches[len(matches)-1][0], nil
}

func findLastNonDotInSlice(slice []rune) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if unicode.IsDigit(slice[i]) {
			return i
		}
	}
	return -1
}

func findFirstDotInSlice(slice []rune) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == '.' {
			return i
		}
	}
	return -1
}

func part1(lines []string) {
	fmt.Println("## Part1")

	//create the disk
	diskId := 0
	diskString := ""
	for i := 0; i < len(lines[0]); i++ {
		parsedNumber, err := strconv.Atoi(string(lines[0][i]))
		if err != nil {
			log.Fatal("Error during number conversion")
		}
		if i%2 == 0 {
			//disk size
			diskString = diskString + strings.Repeat(strconv.Itoa(diskId), parsedNumber)
			diskId++
		} else {
			diskString = diskString + strings.Repeat(".", parsedNumber)
		}
	}
	// for {
	// 	lastNumberIndex, _ := findLastNonDot(diskString)
	// 	firstDotIndex := strings.Index(diskString, ".")
	// 	if (lastNumberIndex < firstDotIndex) {
	// 		break
	// 	}
	// 	numberAtIndex :=
	// }

	//find last non dot char, replace with dot, find first dot char in string, replace with number
	diskSlice := []rune(diskString)
	for {
		//fmt.Println(string(diskSlice))
		lastDigitIndex := findLastNonDotInSlice(diskSlice)
		firstDotIndex := findFirstDotInSlice(diskSlice)
		if lastDigitIndex < firstDotIndex {
			break
		} else {
			tempRune := diskSlice[lastDigitIndex]
			diskSlice[lastDigitIndex] = '.'
			diskSlice[firstDotIndex] = tempRune
		}
	}
	checkSum := 0
	for i, elem := range diskSlice {
		if elem != '.' {
			val, _ := strconv.Atoi(string(elem))
			checkSum += val * i
		}
	}
	fmt.Printf("Checksum is %d\n", checkSum)
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
