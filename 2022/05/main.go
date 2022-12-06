package main

import (
	aoc_common "advent-of-code-2022"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	stacksMap := map[int][]string{
		1: {"J", "H", "G", "M", "Z", "N", "T", "F"},
		2: {"V", "W", "J"},
		3: {"G", "V", "L", "J", "B", "T", "H"},
		4: {"B", "P", "J", "N", "C", "D", "V", "L"},
		5: {"F", "W", "S", "M", "P", "R", "G"},
		6: {"G", "H", "C", "F", "B", "N", "V", "M"},
		7: {"D", "H", "G", "M", "R"},
		8: {"H", "N", "M", "V", "Z", "D"},
		9: {"G", "N", "F", "H"},
	}
	file := aoc_common.OpenFile("05/input")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// Ignore
		commandWords := strings.Split(line, " ")
		if commandWords[0] != "move" {
			continue
		}
		quantityMove, _ := strconv.Atoi(commandWords[1])
		moveFrom, _ := strconv.Atoi(commandWords[3])
		moveTo, _ := strconv.Atoi(commandWords[5])

		// Perform move command one crate at a time
		for i := 0; i < quantityMove; i++ {
			stacksMap[moveTo] = append(stacksMap[moveTo], stacksMap[moveFrom][len(stacksMap[moveFrom])-1])
			stacksMap[moveFrom] = removeLast(stacksMap[moveFrom])
		}
	}

	fmt.Print("Part 1: ")
	for i := 1; i <= 9; i++ {
		fmt.Print(stacksMap[i][len(stacksMap[i])-1])
	}
	fmt.Println()
}

func part2() {
	stacksMap := map[int][]string{
		1: {"J", "H", "G", "M", "Z", "N", "T", "F"},
		2: {"V", "W", "J"},
		3: {"G", "V", "L", "J", "B", "T", "H"},
		4: {"B", "P", "J", "N", "C", "D", "V", "L"},
		5: {"F", "W", "S", "M", "P", "R", "G"},
		6: {"G", "H", "C", "F", "B", "N", "V", "M"},
		7: {"D", "H", "G", "M", "R"},
		8: {"H", "N", "M", "V", "Z", "D"},
		9: {"G", "N", "F", "H"},
	}
	file := aoc_common.OpenFile("05/input")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// Ignore
		commandWords := strings.Split(line, " ")
		if commandWords[0] != "move" {
			continue
		}
		quantityMove, _ := strconv.Atoi(commandWords[1])
		moveFrom, _ := strconv.Atoi(commandWords[3])
		moveTo, _ := strconv.Atoi(commandWords[5])

		// Perform move command all crates at once
		stacksMap[moveTo] = append(stacksMap[moveTo],
			stacksMap[moveFrom][len(stacksMap[moveFrom])-quantityMove:]...)
		stacksMap[moveFrom] = stacksMap[moveFrom][:len(stacksMap[moveFrom])-quantityMove]
	}

	fmt.Print("Part 2: ")
	for i := 1; i <= 9; i++ {
		fmt.Print(stacksMap[i][len(stacksMap[i])-1])
	}
	fmt.Println()

}

func removeLast(stack []string) []string {
	return append(stack[:len(stack)-1], stack[len(stack):]...)
}
