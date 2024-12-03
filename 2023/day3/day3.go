package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type TYPE int

const (
	CHARACTER TYPE = iota
	EMPTY
	NUMBER
)

type Position struct {
	t     TYPE
	x     int
	y     int
	valid bool
	value rune
}

func newPosition(t TYPE, x int, y int, value rune) Position {
	p := Position{
		t:     t,
		x:     x,
		y:     y,
		valid: false,
		value: value,
	}
	return p
}

func (p *Position) checkIfValid(pos *Position) {
	//if it is a Chracter
	if p.t == CHARACTER {
		fmt.Printf("Found a character for %v\n", p)
		fmt.Printf("Setting valid for %v to true\n", pos)
		pos.valid = true
	}
}

type Schematic struct {
	positions []Position
}

func (p *Position) checkAdjacentPosition(pos *Position) {
	if p.t == NUMBER && pos.t == NUMBER {
		pos.valid = true
	}
}

func main() {
	fmt.Println("Starting day 3")

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	positions := []Position{}
	x := 0
	rowLength := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()

		rowLength = len(line)

		for y, char := range line {
			switch {
			case char == '.':
				positions = append(positions, Position{t: EMPTY, x: x, y: y, value: char})
			case char >= '0' && char <= '9':
				positions = append(positions, Position{t: NUMBER, x: x, y: y, value: char})
			default:
				positions = append(positions, Position{t: CHARACTER, x: x, y: y, value: char})
			}
		}
		x++
	}

	for i := range positions {
		// pos.checkIfValid()
		//if it is a number
		if positions[i].t == NUMBER {

			adjacentIndices := []int{
				i - 1,
				i + 1,
				i - rowLength - 1,
				i - rowLength,
				i - rowLength + 1,
				i + rowLength - 1,
				i + rowLength,
				i + rowLength + 1,
			}
			for _, adjIndex := range adjacentIndices {
				if adjIndex >= 0 && adjIndex < len(positions) {
					positions[adjIndex].checkIfValid(&positions[i])
				}

				//check other numbers which are concatenated
				if positions[i].valid == true {

					fmt.Println("Found a valid Number")

					innerLoop := i
					for {
						if innerLoop > 0 {
							positions[innerLoop].checkAdjacentPosition(&positions[innerLoop-1])
							if positions[innerLoop-1].valid == true {
								innerLoop--
							} else {
								break
							}
						} else {
							break
						}
					}
					innerLoop2 := innerLoop
					for {
						if innerLoop2 < len(positions) {
							positions[innerLoop2].checkAdjacentPosition(&positions[innerLoop2+1])
							if positions[innerLoop2+1].valid == true {
								innerLoop2++
							} else {
								break
							}
						} else {
							break
						}
					}

					// adjacentIndeces2 := []int{
					// 	i - 2,
					// 	i - 1,
					// 	i + 1,
					// 	i + 2,
					// }
					// currentRow := i / rowLength

					// for _, adjIndex2 := range adjacentIndeces2 {
					// 	if adjIndex2 >= 0 && adjIndex2 < len(positions) && adjIndex2/rowLength == currentRow {
					// 		if positions[adjIndex2].t == NUMBER {
					// 			fmt.Println("Setting adjacent Number to true")
					// 			positions[adjIndex2].valid = true
					// 		}
					// 	}
					// }
				}
			}
		}
	}

	// // Second pass: mark adjacent digits of valid numbers
	// for i := range positions {
	// 	if positions[i].t == NUMBER && positions[i].valid {
	// 		// Mark adjacent digits in same row
	// 		// Look left
	// 		for j := i - 1; j >= 0 && j/rowLength == i/rowLength && positions[j].t == NUMBER; j-- {
	// 			positions[j].valid = true
	// 		}
	// 		// Look right
	// 		for j := i + 1; j < len(positions) && j/rowLength == i/rowLength && positions[j].t == NUMBER; j++ {
	// 			positions[j].valid = true
	// 		}
	// 	}
	// }

	fmt.Println("###################################################")

	numbersArray := []int{}
	tempRunes := []rune{}
	for i, _ := range positions {
		// fmt.Printf("Index: %d val: %v\n", i, pos)
		if positions[i].valid == true {
			tempRunes = append(tempRunes, positions[i].value)
		} else if len(tempRunes) > 0 {
			str := string(tempRunes)
			fmt.Println("+++")
			fmt.Println(str)
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			numbersArray = append(numbersArray, num)
			tempRunes = nil
		}
	}

	sum := 0
	for _, num := range numbersArray {
		// println(num)
		sum += num
	}
	fmt.Printf("Final sum is %d\n", sum)

	//which ones to compare.
	// i-1, i+1, i-rowlen-1, i-rowlen, i-rowlen+1, i+rowlne-1, i+rowlen, i+rowlen-1

}
