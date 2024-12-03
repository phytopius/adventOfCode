package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitInput(input string) (arr1 []int, arr2 []int) {
	input = strings.ReplaceAll(input, "   ", " ")
	splitInput := strings.Fields(input)
	var input1 = []int{}
	var input2 = []int{}
	for i := 0; i < len(splitInput); i++ {
		number, err := strconv.Atoi(splitInput[i])
		check(err)
		if i%2 == 0 {

			input1 = append(input1, number)
		} else {
			input2 = append(input2, number)
		}
	}
	return input1, input2
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func calculateDistanceArray(arr1 []int, arr2 []int) []int {
	actualLength := min(len(arr1), len(arr2))

	distanceArray := make([]int, actualLength)
	for i := 0; i < actualLength; i++ {
		distanceArray[i] = absInt(arr1[i] - arr2[i])
	}
	return distanceArray
}

func sumOfArray(arr []int) int {
	distance := 0
	for i := 0; i < len(arr); i++ {
		distance += arr[i]
	}
	return distance
}

func calculateSimularity(arr1 []int, arr2 []int) int {
	similarities := []int{}
	for i := 0; i < len(arr1); i++ {
		counter := 0
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				counter++
			}
		}
		if counter != 0 {
			similarities = append(similarities, arr1[i]*counter)
		}

		//fmt.Println("Number is %v   Occurance is %v", arr1[i], counter, arr1[i]*counter)
	}

	similarity := sumOfArray(similarities)
	return similarity
}

func main() {
	fmt.Println("Starting with Problem 1_1")

	//read the data
	dat, err := os.ReadFile("./day1_input.txt")
	check(err)

	//Convert to an array of integers
	arr1, arr2 := splitInput(string(dat))
	sort.Ints(arr1)
	sort.Ints(arr2)
	distanceArray := calculateDistanceArray(arr1, arr2)
	distance := sumOfArray(distanceArray)
	fmt.Printf("Distance is: %v\n", (distance))

	similarity := calculateSimularity(arr1, arr2)
	fmt.Printf("Similarty is: %v\n", similarity)
}
