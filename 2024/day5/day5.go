package main

import (
	"fmt"
	"strings"

	adventOfCode "github.com/phytopius/adventOfCode/helpers"
)

func part1(lines []string) {
	fmt.Println("## Part1")
	rules := [][]int{}
	inputs := [][]int{}

	for _, line := range lines {
		ruleString := strings.Split(line, "|")
		if len(ruleString) == 2 {
			rule, err := adventOfCode.ConvertStringSliceToIntSlice(ruleString)
			if err != nil {
				fmt.Println("something went wrong on rule parsing")
			}
			rules = append(rules, rule)
		}
		inputString := strings.Split(line, ",")
		if len(inputString) > 1 {
			input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)
			inputs = append(inputs, input)
		}
	}

	result := 0
	for _, input := range inputs {
		fmt.Printf("Input is %v\n", input)
		invalidRow := false
		for i := range len(rules) {

			// check index of first number
			// check index of second number
			// if index in input of first number is greater than second number the order is incorrect
			// flag it as invalid and go to next input row
			// it the input is in the correct oder get the middle number of the input and add it to the result
			indexValue1 := adventOfCode.FindIndex(input, rules[i][0])
			if indexValue1 != -1 {
				indexValue2 := adventOfCode.FindIndex(input, rules[i][1])
				if indexValue2 != -1 {
					if indexValue1 > indexValue2 {
						invalidRow = true
						continue
					}
				}
			}
		}
		if invalidRow == true {
			continue
		} else {
			middleIndex := len(input) / 2
			result += input[middleIndex]
		}
	}
	fmt.Println(result)
}

func part2(lines []string) {
	fmt.Println("## Part2")
	rules := [][]int{}
	inputs := [][]int{}

	for _, line := range lines {
		ruleString := strings.Split(line, "|")
		if len(ruleString) == 2 {
			rule, err := adventOfCode.ConvertStringSliceToIntSlice(ruleString)
			if err != nil {
				fmt.Println("something went wrong on rule parsing")
			}
			rules = append(rules, rule)
		}
		inputString := strings.Split(line, ",")
		if len(inputString) > 1 {
			input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)
			inputs = append(inputs, input)
		}
	}
	result := 0
	for _, input := range inputs {
		//fmt.Printf("Looking at input %v\n", input)
		//check every two numbers in reversed order and compre it to every rule
		//if there is  match change position of the those two numbers
		//not sure yet what a good way is to make sure newly created wrong pairs are also checked again
		//1) I could iterate len times over all inputs that should be enough
		//2) I could check every newly created pair again with their new neighbor
		//3) I could also just move backwards to the start if I switched something and stop the row if I reach the end without switching
		switched := false
		for k := 0; k < len(input); k++ { //Option 1 First I think this is inefficient but will work
			for j := 0; j < len(input)-1; j++ {
				num1 := input[j]
				num2 := input[j+1]
				//fmt.Printf("++Pair %d,%d\n", num1, num2)
				for _, rule := range rules {
					if num2 == rule[0] && num1 == rule[1] {
						switched = true
						temp := input[j]
						input[j] = input[j+1]
						input[j+1] = temp
						//fmt.Printf("+Input is now %v\n", input)
					}
				}
			}
		}
		if switched == true {
			middleIndex := len(input) / 2
			result += input[middleIndex]
		}
	}
	fmt.Printf("Result 2 is %d\n", result)
}

// testing out other things
// using a map instead of a list for the rules for O(1) access check
func part2_b(lines []string) {
	fmt.Println("## Part2_b")
	rules := make(map[[2]int]bool)
	inputs := [][]int{}

	for _, line := range lines {
		ruleString := strings.Split(line, "|")
		if len(ruleString) == 2 {
			rule, err := adventOfCode.ConvertStringSliceToIntSlice(ruleString)
			if err != nil {
				fmt.Println("something went wrong on rule parsing")
			}
			rules[[2]int{rule[0], rule[1]}] = true
		}
		inputString := strings.Split(line, ",")
		if len(inputString) > 1 {
			input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)
			inputs = append(inputs, input)
		}
	}
	result := 0
	for _, input := range inputs {
		//fmt.Printf("Looking at input %v\n", input)
		//check every two numbers in reversed order and compre it to every rule
		//if there is  match change position of the those two numbers
		//not sure yet what a good way is to make sure newly created wrong pairs are also checked again
		//1) I could iterate len times over all inputs that should be enough
		//2) I could check every newly created pair again with their new neighbor
		//3) I could also just decrease the index if I swapped something and increase if i did not swap anything
		switched := false
		for k := 0; k < len(input); k++ { //Option 1 First I think this is inefficient but will work
			for j := 0; j < len(input)-1; j++ {
				num1 := input[j]
				num2 := input[j+1]
				//fmt.Printf("++Pair %d,%d\n", num1, num2)
				if rules[[2]int{num2, num1}] {
					switched = true
					temp := input[j]
					input[j] = input[j+1]
					input[j+1] = temp
					// fmt.Printf("+Input is now %v\n", input)
				}
			}
		}
		if switched == true {
			middleIndex := len(input) / 2
			result += input[middleIndex]
		}
	}
	fmt.Printf("Result 2_b is %d\n", result)
}

// testing out other things
// using a map instead of a list for the rules for O(1) access check
// also doing the go back on swap go forth if not swap
// I think this is the fastest solution I can come up with
func part2_c(lines []string) {
	fmt.Println("## Part2_c")
	rules := make(map[[2]int]bool)
	inputs := [][]int{}

	for _, line := range lines {
		ruleString := strings.Split(line, "|")
		if len(ruleString) == 2 {
			rule, err := adventOfCode.ConvertStringSliceToIntSlice(ruleString)
			if err != nil {
				fmt.Println("something went wrong on rule parsing")
			}
			rules[[2]int{rule[0], rule[1]}] = true
		}
		inputString := strings.Split(line, ",")
		if len(inputString) > 1 {
			input, _ := adventOfCode.ConvertStringSliceToIntSlice(inputString)
			inputs = append(inputs, input)
		}
	}
	result := 0
	for _, input := range inputs {
		switched := false
		i := 0
		for i < len(input)-1 {
			if rules[[2]int{input[i+1], input[i]}] {
				switched = true
				input[i], input[i+1] = input[i+1], input[i] //I wonder if this works
				if i > 0 {
					i--
				}
			} else {
				i++
			}
		}
		if switched == true {
			middleIndex := len(input) / 2
			result += input[middleIndex]
		}
	}
	fmt.Printf("Result 2_c is %d\n", result)
}
func main() {
	fmt.Println("Starting day")
	lines := adventOfCode.ReadFileLineByLine("./input.txt")
	part1(lines)
	part2(lines)
	part2_b(lines)
	part2_c(lines)
}
