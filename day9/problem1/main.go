package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	for i := 25; i < len(numbers); i++ {
		if !isValid(numbers, i, 25) {
			println("First invalid number is", numbers[i])
			break
		}
	}
}

func isValid(numbers []int, pos, preambleSize int) bool {
	for i := pos - preambleSize; i < pos-1; i++ {
		for j := i + 1; j < pos; j++ {
			if numbers[i]+numbers[j] == numbers[pos] {
				return true
			}
		}
	}

	return false
}
