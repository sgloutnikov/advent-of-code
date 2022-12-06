package main

import (
	aoc_common "advent-of-code-2022"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file := aoc_common.OpenFile("02/input")
	defer file.Close()

	moveLookup := map[string]string{"A": "R", "B": "P", "C": "S", "X": "R", "Y": "P", "Z": "S"}
	var totalScore int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		moves := strings.Split(line, " ")

		oppMove := moveLookup[moves[0]]
		myMove := moveLookup[moves[1]]

		totalScore += getScore(oppMove, myMove)
	}
	fmt.Println("Part 1 Total Score", totalScore)
}

func part2() {
	file := aoc_common.OpenFile("02/input")
	defer file.Close()

	moveLookup := map[string]string{"A": "R", "B": "P", "C": "S"}
	moveLookupPt2 := map[string]string{"RX": "S", "RY": "R", "RZ": "P",
		"PX": "R", "PY": "P", "PZ": "S",
		"SX": "P", "SY": "S", "SZ": "R"}
	var totalScore int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		moves := strings.Split(line, " ")

		oppMove := moveLookup[moves[0]]
		myMove := moveLookupPt2[oppMove+moves[1]]

		totalScore += getScore(oppMove, myMove)
	}
	fmt.Println("Part 2 Total Score", totalScore)
}

func getScore(oppMove string, myMove string) int {
	shapeScore := map[string]int{"R": 1, "P": 2, "S": 3}
	outcomeScore := map[string]int{"W": 6, "D": 3, "L": 0}
	var totalScore int

	// Draw
	if oppMove == myMove {
		totalScore += shapeScore[myMove] + outcomeScore["D"]
	} else if oppMove == "R" {
		// Rock
		if myMove == "P" {
			totalScore += shapeScore[myMove] + outcomeScore["W"]
		} else {
			totalScore += shapeScore[myMove] + outcomeScore["L"]
		}
	} else if oppMove == "P" {
		// Paper
		if myMove == "S" {
			totalScore += shapeScore[myMove] + outcomeScore["W"]
		} else {
			totalScore += shapeScore[myMove] + outcomeScore["L"]
		}
	} else if oppMove == "S" {
		// Scissor
		if myMove == "R" {
			totalScore += shapeScore[myMove] + outcomeScore["W"]
		} else {
			totalScore += shapeScore[myMove] + outcomeScore["L"]
		}
	}
	return totalScore
}
