package main

import (
	aoc_common "advent-of-code-2022"
	"bufio"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	file := aoc_common.OpenFile("06/input")
	defer file.Close()
	const MarkerLength = 4
	markerWindow := make(map[rune]int)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		dataStream := []rune(line)

		for i := 0; i < len(dataStream); i++ {
			for j := 0; j < MarkerLength; j++ {
				if _, ok := markerWindow[dataStream[i+j]]; !ok {
					//If not in window add it
					markerWindow[dataStream[i+j]] = i
				} else {
					markerWindow = make(map[rune]int)
					break
				}
			}
			if len(markerWindow) == MarkerLength {
				fmt.Println("Part 1 marker found:", i+MarkerLength)
				break
			}
		}
	}
}

func part2() {
	file := aoc_common.OpenFile("06/input")
	defer file.Close()
	const MarkerLength = 14
	markerWindow := make(map[rune]int)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		dataStream := []rune(line)

		for i := 0; i < len(dataStream); i++ {
			for j := 0; j < MarkerLength; j++ {
				if _, ok := markerWindow[dataStream[i+j]]; !ok {
					//If not in window add it
					markerWindow[dataStream[i+j]] = i
				} else {
					markerWindow = make(map[rune]int)
					break
				}
			}
			if len(markerWindow) == MarkerLength {
				fmt.Println("Part 2 marker found:", i+MarkerLength)
				break
			}
		}
	}
}
