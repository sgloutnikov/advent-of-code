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
	file := aoc_common.OpenFile("03/input")
	defer file.Close()

	var sumPriorities int
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		compartment1 := line[0 : len(line)/2]
		compartment2 := line[len(line)/2:]
		c1Items := make([]int, 52)
		c2Items := make([]int, 52)

		// Identify compartment item priorities and count
		for _, r := range compartment1 {
			c1Items[getItemPriority(r)-1]++
		}
		for _, r := range compartment2 {
			c2Items[getItemPriority(r)-1]++
		}

		// Identify item in both compartments
		for i, _ := range c1Items {
			// If item found in both compartments
			if c1Items[i] > 0 && c2Items[i] > 0 {
				sumPriorities += i + 1
			}
		}
	}
	fmt.Println("Part 1 Sum Priorities:", sumPriorities)
}

func part2() {
	file := aoc_common.OpenFile("03/input")
	defer file.Close()

	var elfNumberTracker, sumPriorities int
	groupRutsacks := make(map[int][]int)
	groupRutsacks[0] = make([]int, 52)
	groupRutsacks[1] = make([]int, 52)
	groupRutsacks[2] = make([]int, 52)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		for _, r := range line {
			groupRutsacks[elfNumberTracker][getItemPriority(r)-1]++
		}
		elfNumberTracker++

		// Reset for new group
		if elfNumberTracker == 3 {
			// Find common item
			for i, _ := range groupRutsacks[0] {
				if groupRutsacks[0][i] > 0 && groupRutsacks[1][i] > 0 && groupRutsacks[2][i] > 0 {
					sumPriorities += i + 1
				}
			}
			// Reset
			groupRutsacks[0] = make([]int, 52)
			groupRutsacks[1] = make([]int, 52)
			groupRutsacks[2] = make([]int, 52)
			elfNumberTracker = 0
		}
	}
	fmt.Println("Part 2 Sum Priorities:", sumPriorities)
}

func getItemPriority(c rune) int {
	var p int
	// Lowercase
	if c >= 97 && c <= 122 {
		p = int(c) - 96
	} else if c >= 65 && c <= 90 {
		// Uppercase
		p = int(c) - 38
	} else {
		fmt.Println("Unsupported character")
	}
	return p
}
