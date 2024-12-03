package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getResult(dataArray []string) int {
	result := 0
	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]
		numStr := ""
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			if _, err := strconv.Atoi(char); err == nil {
				numStr += char
				break
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			char := string(line[j])
			if _, err := strconv.Atoi(char); err == nil {
				numStr += char
				break
			}
		}
		if number, err := strconv.Atoi(numStr); err == nil {
			result += number

		}
	}
	return result
}
func main() {
	fmt.Println("Starting day 1")

	data, err := os.ReadFile("./input.txt")
	check(err)

	dataArray := strings.Fields(string(data))
	fmt.Printf("length is %v\n", len(dataArray))

	result := getResult(dataArray)

	fmt.Printf("Result is: %v\n", result)

	//1_2

	//numberWords := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbersMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	testing := dataArray
	fmt.Printf("Length of testing is %v\n", strconv.Itoa(len(testing)))
	replacedDataArray := []string{}
	for _, subStr := range testing {
		tempStr := replaceFirst(subStr, numbersMap)
		tempStr = replaceLast(tempStr, numbersMap)
		replacedDataArray = append(replacedDataArray, tempStr)
		// replacedDataArray = append(replacedDataArray, replacebasedOnMap(subStr, numbersMap))
	}

	// for i, value := range replacedDataArray {
	// 	fmt.Println(strconv.Itoa(i+1) + " " + value)
	// }

	// testing := replacedDataArray[61:62]

	for i, value := range replacedDataArray {
		fmt.Println(strconv.Itoa(i+1) + " " + value)
	}

	result2 := getResult(replacedDataArray)
	fmt.Printf("Result2 is %v\n", result2)
}

func replaceFirst(input string, replacements map[string]string) string {
	firstIndex := -1
	firstSubstring := ""
	for key := range replacements {
		index := strings.Index(input, key)
		if index != -1 && (index < firstIndex || firstIndex == -1) {
			firstIndex = index
			firstSubstring = key
		}
	}
	if firstIndex != -1 {
		fmt.Printf("Found first at %v which is %v\n", strconv.Itoa(firstIndex), firstSubstring)
		value, ok := replacements[firstSubstring]
		if ok {
			input = strings.ReplaceAll(input, firstSubstring, value)
		}
	}
	fmt.Printf("result string is %v\n", input)
	return input
}

func replaceLast(input string, replacements map[string]string) string {
	lastIndex := -1
	lastSubstring := ""
	for key := range replacements {
		index := strings.LastIndex(input, key)
		if index != -1 && (index > lastIndex) {
			lastIndex = index
			lastSubstring = key
		}
	}
	if lastIndex != -1 {
		fmt.Printf("Found last at %v which is %v\n", strconv.Itoa(lastIndex), lastSubstring)
		value, ok := replacements[lastSubstring]
		if ok {
			input = strings.ReplaceAll(input, lastSubstring, value)
		}
	}
	fmt.Printf("result string is %v\n", input)
	return input
}

func replacebasedOnMap(input string, replacements map[string]string) string {
	for key, value := range replacements {
		input = strings.Replace(input, key, value, 1)
		input = strings.Replace(input, key, value, -1)
	}
	return input
}
