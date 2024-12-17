package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

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

func part1(lines []string, blinks int) {
	fmt.Println("## Part1")
	inputString := strings.Split(lines[0], " ")
	input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)
	fmt.Println(input)
	for i := 0; i < blinks; i++ {
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
				input = append(input[0:j], append([]int{num1, num2}, input[j+1:]...)...)

				j++
			} else {
				input[j] = input[j] * 2024
			}
		}
	}
	fmt.Printf("Number of Stones %d\n", len(input))
}

func part2(lines []string, blinks int) {
	fmt.Println("## Part2")

	inputString := strings.Split(lines[0], " ")
	input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)

	fmt.Println(input)

	numberOfStones := 0
	for _, element := range input {
		everyNumberArray := []int{element}
		for i := 0; i < blinks; i++ {
			for j := 0; j < len(everyNumberArray); j++ {
				digitCount := getDigitCount(everyNumberArray[j])
				if everyNumberArray[j] == 0 {
					everyNumberArray[j] = 1
				} else if digitCount%2 == 0 {
					//fmt.Printf("Even number of digits at index %d\n", j)
					numberAsString := strconv.Itoa(everyNumberArray[j])
					midpoint := digitCount / 2
					num1, _ := strconv.Atoi(numberAsString[0:midpoint])
					num2, _ := strconv.Atoi(numberAsString[midpoint:])
					everyNumberArray = append(everyNumberArray[0:j], append([]int{num1, num2}, everyNumberArray[j+1:]...)...)

					j++
				} else {
					everyNumberArray[j] = everyNumberArray[j] * 2024
				}
			}
		}
		numberOfStones += len(everyNumberArray)
	}
	fmt.Printf("Number of stones %d\n", numberOfStones)
}

func part2_b(lines []string, blinks int) {
	fmt.Println("## Part2")
	var wg sync.WaitGroup

	inputString := strings.Split(lines[0], " ")
	input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)
	resultChan := make(chan int, len(input))

	fmt.Println(input)

	for _, element := range input {
		wg.Add(1)
		go processElement(element, blinks, &wg, resultChan)
	}

	wg.Wait()
	close(resultChan)

	numberOfStones := 0
	for result := range resultChan {
		numberOfStones += result
	}
	fmt.Printf("Total number of Stones %d\n", numberOfStones)

}

// first thought was to calculate the number of stones for each element. This would mean less memory stuff for the entire input
// this failed though.
func processElement(element int, blinks int, wg *sync.WaitGroup, resultChan chan<- int) {
	fmt.Printf("Started task for %d\n ", element)
	defer wg.Done()

	everyNumberArray := []int{element}
	for i := 0; i < blinks; i++ {
		fmt.Printf("I am on blink %d\n", i)
		for j := 0; j < len(everyNumberArray); j++ {
			digitCount := getDigitCount(everyNumberArray[j])
			if everyNumberArray[j] == 0 {
				everyNumberArray[j] = 1
			} else if digitCount%2 == 0 {
				//fmt.Printf("Even number of digits at index %d\n", j)
				numberAsString := strconv.Itoa(everyNumberArray[j])
				midpoint := digitCount / 2
				num1, _ := strconv.Atoi(numberAsString[0:midpoint])
				num2, _ := strconv.Atoi(numberAsString[midpoint:])
				everyNumberArray = append(everyNumberArray[0:j], append([]int{num1, num2}, everyNumberArray[j+1:]...)...)

				j++
			} else {
				everyNumberArray[j] = everyNumberArray[j] * 2024
			}
		}
	}
	fmt.Printf("Finished task for %d\n", element)
	resultChan <- len(everyNumberArray)
}

// using memoization
func part2_c(lines []string, blinks int) {
	fmt.Println("##Part2c")
	inputString := strings.Split(lines[0], " ")
	input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)
	cache := map[string]int{}
	score := 0

	for _, element := range input {
		score += next(element, blinks, cache)
	}
	fmt.Printf("Number of stones is %d\n", score)
}

func next(element, remainingSteps int, cache map[string]int) int {
	//we are done
	if remainingSteps == 0 {
		return 1
	}
	remainingSteps--
	//check if this is already in cache
	value, ok := cache[fmt.Sprintf("%d:%d", element, remainingSteps)]
	if ok {
		return value
	}

	//rule for value 0
	if element == 0 {
		resultValue := next(1, remainingSteps, cache)
		cache[fmt.Sprintf("%d:%d", element, remainingSteps)] = resultValue
		return resultValue
	}
	//rule for even digit numbers
	digitCount := getDigitCount(element)
	if digitCount%2 == 0 {
		numberAsString := strconv.Itoa(element)
		num1, _ := strconv.Atoi(numberAsString[0 : digitCount/2])
		num2, _ := strconv.Atoi(numberAsString[digitCount/2:])
		resultValue := next(num1, remainingSteps, cache) + next(num2, remainingSteps, cache)

		cache[fmt.Sprintf("%d:%d", element, remainingSteps)] = resultValue
		return resultValue
	}
	//default rule
	val := element * 2024
	resultValue := next(val, remainingSteps, cache)

	cache[fmt.Sprintf("%d:%d", element, remainingSteps)] = resultValue
	return resultValue
}

func runPart1(lines []string, blinks int) func() {
	return func() {
		part1(lines, blinks)
	}
}

func runPart2(lines []string, blinks int) func() {
	return func() {
		part2_c(lines, blinks)
	}
}
func main() {
	fmt.Println("Starting day")
	blinks := 25
	lines := adventOfCode.ReadFileLineByLine("./testinput.txt")
	//lines = []string{"125 17 12 1 0"}
	//adventOfCode.MeasureTime("part1", runPart1(lines, blinks))
	adventOfCode.MeasureTime("part2", runPart2(lines, blinks))
	//part2(lines)
}
