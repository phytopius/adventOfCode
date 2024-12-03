package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("Starting day2")

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	saveCounter := 0
	lineNumber := 1
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()
		arr := strings.Fields(line)

		positiveCounter := 0
		numArr := []int{}
		invalid := false
		for i := 0; i < len(arr); i++ {
			num, err := strconv.Atoi(arr[i])
			if err != nil {
				log.Fatal("Error parsing number")
			}

			numArr = append(numArr, num)
			if i > 0 {
				diff := numArr[i-1] - numArr[i]
				fmt.Printf("Diff is %v\n", diff)
				if Abs(diff) < 1 || Abs(diff) > 3 {
					invalid = true
					break
				}
				if diff > 0 {
					positiveCounter++
				}
			}
		}
		fmt.Printf("PositiveCounter is %v\n", positiveCounter)
		if invalid == false && (Abs(positiveCounter) == len(arr)-1 || positiveCounter == 0) {
			fmt.Printf("Line %v is safe\n", lineNumber)
			saveCounter++
		}
		lineNumber++
	}
	fmt.Printf("Save Counter is %v\n", saveCounter)

}
