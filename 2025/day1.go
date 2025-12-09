package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const LENGTH = 100

func main() {
	instructions := parse("./day1_intput.txt")
	fmt.Printf("Read %v instructions\n", len(instructions))
	password := decodePt2(50, instructions)
	fmt.Printf("Password is %v\n", password)
}

func decodePt1(start int, instructions []string) int {
	tickle := 0
	pos := start
	for _, instruction := range instructions {

		fmt.Printf("Read instruciont %v \n", instruction)
		// parse the instruction
		direction, steps := parseInstruction(instruction)

		fmt.Printf("Move %v steps towards direction: %q\n", steps, direction)
		switch direction {
		case 'L':
			pos = pos - steps%100
			if pos < 0 {
				pos = pos + LENGTH
			}
		case 'R':
			pos = pos + steps%100
			if pos >= LENGTH {
				pos = pos - LENGTH
			}
		}
		fmt.Printf("Needle is pointing to %v\n", pos)
		if pos == 0 {
			fmt.Printf("Increase the tickle from %v to %v\n", tickle, tickle+1)
			tickle++
		}
		fmt.Printf("Pointing to %v and tickle: %v at the moment\n", pos, tickle)
	}

	fmt.Printf("Password is %v\n", tickle)
	return tickle
}

func decodePt2(start int, instructions []string) int {
	tickle := 0
	pos := start
	for _, instruction := range instructions {
		direction, steps := parseInstruction(instruction)
		fmt.Printf("Move %v steps towards direction: %q from pos: %v\n", steps, direction, pos)

		loops := int(math.Floor(float64(steps / LENGTH)))
		fmt.Printf("%v times will pass through 0\n", loops)
		tickle = tickle + loops

		remaining := steps - loops*100
		if remaining == 0 {
			continue
		}

		prePos := pos
		if direction == 'L' {
			pos = pos - remaining
		} else {
			pos = pos + remaining
		}

		if pos < 0 {
			if prePos > 0 {
				tickle++
			}
			pos = pos + LENGTH
		} else if pos >= LENGTH {
			tickle++
			pos = pos - LENGTH
		} else if pos == 0 {
			tickle++
		}

		fmt.Printf("Pointing to %v and tickle: %v at the moment\n", pos, tickle)
	}
	return tickle
}

func parseInstruction(input string) (byte, int) {
	direction := input[0]
	stepString := input[1:]

	steps, _ := strconv.Atoi(stepString)

	return direction, steps
}

func parse(filename string) []string {
	instructions := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return instructions
}
