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

	accumolator := 0
	visitedLines := []int{}

	for i := 0; i < len(lines); {
		instruction, num := toInstruction(lines[i])

		if contains(visitedLines, i) {
			break
		}

		visitedLines = append(visitedLines, i)

		if instruction == "jmp" {
			i = i + num
			continue
		}

		if instruction == "nop" {
		} else if instruction == "acc" {
			accumolator = accumolator + num
		} else {
			panic("Invalid instruction")
		}

		i++
	}

	println("Accumolator:", accumolator)
}

func toInstruction(line string) (string, int) {
	parts := strings.Split(line, " ")
	num, _ := strconv.Atoi(parts[1])

	return parts[0], num
}

func contains(arr []int, element int) bool {
	for _, a := range arr {
		if a == element {
			return true
		}
	}

	return false
}
