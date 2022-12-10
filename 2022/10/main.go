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
}

func part1() {
	file := aoc_common.OpenFile("10/input")
	defer file.Close()
	var instruction []string
	var instructionOp []int

	var instrCount int
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		instr := strings.Split(fileScanner.Text(), " ")
		instruction = append(instruction, instr[0])
		instructionOp = append(instructionOp, 0)
		if len(instr) > 1 {
			instrOp, _ := strconv.Atoi(instr[1])
			instructionOp[instrCount] = instrOp
		}
		instrCount++
	}

	var register, cycleCount int
	register = 1
	signalStrengths := make(map[int]int)
	for i, _ := range instruction {
		if instruction[i] == "addx" {
			for addCycle := 1; addCycle <= 2; addCycle++ {
				cycleCount++
				if trapCycle(cycleCount) {
					signalStrengths[cycleCount] = register
				}
			}
			register += instructionOp[i]
		} else {
			//noop
			cycleCount++
			if trapCycle(cycleCount) {
				signalStrengths[cycleCount] = register
			}
		}
	}

	var sumSignalStrength int
	for c, r := range signalStrengths {
		sumSignalStrength += c * r
	}
	fmt.Println("Part 1:", sumSignalStrength)
}

func trapCycle(cycleCount int) bool {
	if cycleCount == 20 || (cycleCount+20)%40 == 0 {
		return true
	}
	return false
}
