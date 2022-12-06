package main

import (
	aoc_common "advent-of-code-2022"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file := aoc_common.OpenFile("04/input")
	defer file.Close()
	var fullyEnclosedPairCount, partiallyEnclosdePairCount int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		elvePair := strings.Split(line, ",")
		idStringPair1 := strings.Split(elvePair[0], "-")
		idStringPair2 := strings.Split(elvePair[1], "-")
		idPair1 := make([]int, len(idStringPair1))
		idPair2 := make([]int, len(idStringPair2))
		for i, s := range idStringPair1 {
			idPair1[i], _ = strconv.Atoi(s)
		}
		for i, s := range idStringPair2 {
			idPair2[i], _ = strconv.Atoi(s)
		}

		if isFullyEnclosed(idPair1, idPair2) {
			fullyEnclosedPairCount++
		}
		if isPartiallyEnclosed(idPair1, idPair2) {
			partiallyEnclosdePairCount++
		}
	}
	fmt.Println("Part 1 total pairs:", fullyEnclosedPairCount)
	fmt.Println("Part 2 total pairs:", partiallyEnclosdePairCount)
}

func isFullyEnclosed(idPair1 []int, idPair2 []int) bool {
	// Left in Right
	if idPair1[0] >= idPair2[0] && idPair1[1] <= idPair2[1] {
		return true
	}
	// Right in Left
	if idPair2[0] >= idPair1[0] && idPair2[1] <= idPair1[1] {
		return true
	}
	return false
}

func isPartiallyEnclosed(idPair1 []int, idPair2 []int) bool {
	// Left Start in Right
	if idPair1[0] >= idPair2[0] && idPair1[0] <= idPair2[1] {
		return true
	}
	// Right Start in Left
	if idPair2[0] >= idPair1[0] && idPair2[0] <= idPair1[1] {
		return true
	}
	return false
}
