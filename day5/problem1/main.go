package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxID := 0

	for scanner.Scan() {
		boardingPass := scanner.Text()
		row := getRow(boardingPass)
		column := getColumn(boardingPass)
		id := row*8 + column

		if id > maxID {
			maxID = id
		}
	}

	println("Max id is", maxID)
}

func getRow(boardingPass string) int {
	lower := 0
	upper := 127
	for _, char := range boardingPass {
		if char == []rune("B")[0] {
			lower = (lower+upper)/2 + 1
		}

		if char == []rune("F")[0] {
			upper = (lower + upper) / 2
		}
	}

	return lower
}

func getColumn(boardingPass string) int {
	left := 0
	right := 7
	for _, char := range boardingPass {
		if char == []rune("R")[0] {
			left = (left+right)/2 + 1
		}

		if char == []rune("L")[0] {
			right = (left + right) / 2
		}
	}

	return left
}
