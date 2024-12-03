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

		positiveOrNegative := 0
		elementRemoved := false
		positiveCounter := 0
		invalid := false

		previous := 0

		for i := 0; i < len(arr); {
			num, err := strconv.Atoi(arr[i])
			if err != nil {
				log.Fatal("Error parsing number")
			}
			if i == 0 {
				num2, err := strconv.Atoi(arr[i+1])
				if err != nil {
					log.Fatal("error parsing number")
				}
				if Abs(num-num2) < 1 || Abs(num-num2) > 3 {
					elementRemoved = true
					arr = append(arr[:i], arr[i+1:]...)
				}
			}

			if i > 0 {
				diff := previous - num
				fmt.Printf("Diff is %v\n", diff)
				if Abs(diff) < 1 || Abs(diff) > 3 {
					if elementRemoved == false {
						fmt.Println("Broke the rules: Difference too big")
						elementRemoved = true
						arr = append(arr[:i], arr[i+1:]...)
						continue
						//remove element somehow
					} else {
						fmt.Println("DOUBLE BROKE: RIP")
						invalid = true
						break
					}
				}
				//set initial direction
				if i == 1 {
					if diff > 0 {
						positiveOrNegative = 1
					} else {
						positiveOrNegative = -1
					}
				} else {
					//postivie direction
					if positiveOrNegative == 1 {
						if diff < 0 && elementRemoved == false {
							fmt.Println("Broke the rules: NEGATIVE diff")
							elementRemoved = true
							//remove element somehow
							arr = append(arr[:i], arr[i+1:]...)
							continue
						} else if diff < 0 && elementRemoved == true {
							fmt.Println("DOUBLE BROKE: RIP")
							break
						}
						//negative direction
					} else {
						if diff > 0 && elementRemoved == false {
							fmt.Println("Broke the rules: POSITIVE diff")
							elementRemoved = true
							elementRemoved = true
							arr = append(arr[:i], arr[i+1:]...)
							continue
							//remove element somehow
						} else if diff > 0 && elementRemoved == true {
							fmt.Println("DOUBLE BROKE: RIP")
							break
						}
					}
				}
			}
			previous = num
			i++
		}
		fmt.Printf("PositiveCounter is %v\n", positiveCounter)
		if invalid == false && (Abs(positiveCounter) == len(arr)-1 || positiveCounter == 0) {
			fmt.Printf("Line %v is safe\n", lineNumber)
			saveCounter++
		} else if invalid == false && (Abs(positiveCounter) == len(arr)-2 || positiveCounter == 1) {
			fmt.Printf("Line %v is safe with exception\n", lineNumber)
			saveCounter++
		}
		lineNumber++
	}
	fmt.Printf("Save Counter is %v\n", saveCounter)

}
