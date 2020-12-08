package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	accumolator := correctProgram(lines)
	println("Accumolator:", accumolator)

}

func correctProgram(lines []string) int {
	accumolator, lastLineProccessed := run(lines, -1)
	positionToFlip := lastLineProccessed

	for lastLineProccessed != len(lines)-1 && positionToFlip >= 0 {
		accumolator, lastLineProccessed = run(lines, positionToFlip)
		positionToFlip--
	}

	if lastLineProccessed == len(lines)-1 {
		return accumolator
	}

	return -1
}

func run(lines []string, lineToFlip int) (int, int) {
	accumolator := 0
	visitedLines := []int{}
	lineNum := 0

	for lineNum < len(lines) {
		if contains(visitedLines, lineNum) {
			break
		}

		visitedLines = append(visitedLines, lineNum)
		instruction, num := toInstruction(lines[lineNum], lineNum == lineToFlip)

		if instruction == "jmp" {
			lineNum = lineNum + num
			continue
		}

		if instruction == "nop" {
		} else if instruction == "acc" {
			accumolator = accumolator + num
		} else {
			panic("Invalid instruction")
		}

		lineNum++
	}

	return accumolator, max(visitedLines)
}

func toInstruction(line string, flip bool) (string, int) {
	parts := strings.Split(line, " ")
	num, _ := strconv.Atoi(parts[1])
	instruction := parts[0]

	if flip {
		if parts[0] == "nop" {
			instruction = "jmp"
		}
		if parts[0] == "jmp" {
			instruction = "nop"
		}
	}

	return instruction, num
}

func max(arr []int) int {
	result := 0

	for _, num := range arr {
		if result < num {
			result = num
		}
	}

	return result
}

func contains(arr []int, element int) bool {
	for _, a := range arr {
		if a == element {
			return true
		}
	}

	return false
}
