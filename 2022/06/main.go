package main

import (
	aoc_common "advent-of-code-2022"
	"bufio"
	"fmt"
)

func main() {
	findFirstMarker("Part 1 Marker:", 4)
	findFirstMarker("Part 2 Marker:", 14)
}

func findFirstMarker(title string, markerLength int) {
	file := aoc_common.OpenFile("06/input")
	defer file.Close()
	markerWindow := make(map[rune]int)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		dataStream := []rune(line)

		for i := 0; i < len(dataStream); i++ {
			for j := 0; j < markerLength; j++ {
				if _, ok := markerWindow[dataStream[i+j]]; !ok {
					//If not in window add it
					markerWindow[dataStream[i+j]] = i
				} else {
					markerWindow = make(map[rune]int)
					break
				}
			}
			if len(markerWindow) == markerLength {
				fmt.Println(title, i+markerLength)
				break
			}
		}
	}
}
