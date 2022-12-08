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
	file := aoc_common.OpenFile("08/input")
	defer file.Close()
	var forest [][]int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		rowRunes := []rune(fileScanner.Text())
		forestRow := make([]int, len(rowRunes))
		for i, r := range rowRunes {
			t := int(r - '0')
			forestRow[i] = t
		}
		forest = append(forest, forestRow)
	}

	var visibleCount int
	for row := 1; row <= len(forest)-2; row++ {
		for col := 1; col <= len(forest)-2; col++ {
			if isVisible(forest, row, col) {
				visibleCount += 1
			}
		}
	}
	fmt.Println("Part 1:", visibleCount+(2*(len(forest)))+(2*(len(forest)-2)))
}

func isVisible(forest [][]int, rowLoc int, colLoc int) bool {
	tree := forest[rowLoc][colLoc]
	var left, right, top, bottom bool
	left, right, top, bottom = true, true, true, true

	// Check Left
	for i := 0; i < colLoc; i++ {
		if forest[rowLoc][i] >= tree {
			left = false
			break
		}
	}
	// Check Right
	for i := colLoc + 1; i < len(forest); i++ {
		if forest[rowLoc][i] >= tree {
			right = false
			break
		}
	}
	// Check Top
	for i := 0; i < rowLoc; i++ {
		if forest[i][colLoc] >= tree {
			top = false
			break
		}
	}
	// Check Bottom
	for i := rowLoc + 1; i < len(forest); i++ {
		if forest[i][colLoc] >= tree {
			bottom = false
			break
		}
	}
	if left || right || top || bottom {
		return true
	} else {
		return false
	}
}

func part2() {
	file := aoc_common.OpenFile("08/input")
	defer file.Close()
	var forest [][]int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		rowRunes := []rune(fileScanner.Text())
		forestRow := make([]int, len(rowRunes))
		for i, r := range rowRunes {
			t := int(r - '0')
			forestRow[i] = t
		}
		forest = append(forest, forestRow)
	}

	var maxScenicScore int
	for row := 1; row <= len(forest)-2; row++ {
		for col := 1; col <= len(forest)-2; col++ {
			score := getScenicScore(forest, row, col)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}
	fmt.Println("Part 2:", maxScenicScore)
}

func getScenicScore(forest [][]int, rowLoc int, colLoc int) int {
	tree := forest[rowLoc][colLoc]
	var left, right, top, bottom int

	// Left
	for i := colLoc - 1; i >= 0; i-- {
		left++
		if tree <= forest[rowLoc][i] {
			break
		}
	}
	// Right
	for i := colLoc + 1; i < len(forest); i++ {
		right++
		if tree <= forest[rowLoc][i] {
			break
		}
	}
	// Top
	for i := rowLoc - 1; i >= 0; i-- {
		top++
		if tree <= forest[i][colLoc] {
			break
		}
	}
	// Bottom
	for i := rowLoc + 1; i < len(forest); i++ {
		bottom++
		if tree <= forest[i][colLoc] {
			break
		}
	}
	return left * right * top * bottom
}
