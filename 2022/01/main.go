package main

import (
	aoc_common "advent-of-code-2022"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	file := aoc_common.OpenFile("01/input")
	defer file.Close()
	var calorieCounter int
	var elvesCalories []int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			itemCalories, err := strconv.Atoi(line)
			if err != nil {
			}
			calorieCounter += itemCalories
		} else {
			elvesCalories = append(elvesCalories, calorieCounter)
			calorieCounter = 0
		}
	}
	sort.Ints(elvesCalories)
	fmt.Println("Max calories:", elvesCalories[len(elvesCalories)-1])
	fmt.Println("Sum top 3:", elvesCalories[len(elvesCalories)-1]+
		elvesCalories[len(elvesCalories)-2]+
		elvesCalories[len(elvesCalories)-3])
}
