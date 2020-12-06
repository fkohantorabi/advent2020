package main

import (
	"bufio"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ids := []int{}

	for scanner.Scan() {
		boardingPass := scanner.Text()
		row := getRow(boardingPass)
		column := getColumn(boardingPass)
		id := row*8 + column
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] == 2 {
			println("ID=", ids[i]-1)
		}
	}
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
