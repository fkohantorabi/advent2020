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

	firstInvalidNumber := findFirstInvalidNumber(numbers)
	min, max := findMinMaxInContiguousSet(numbers, firstInvalidNumber)

	println("Sum of min and max is", min+max)
}

func findFirstInvalidNumber(numbers []int) int {
	for i := 25; i < len(numbers); i++ {
		if !isValid(numbers, i, 25) {
			return numbers[i]
		}
	}

	return -1
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

func findMinMaxInContiguousSet(numbers []int, expectedSum int) (int, int) {
	for i := range numbers {
		min, max := numbers[i], numbers[i]
		sum := numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] > max {
				max = numbers[j]
			}

			if numbers[j] < min {
				min = numbers[j]
			}

			sum = sum + numbers[j]

			if expectedSum == sum {
				return min, max
			}
		}
	}

	return -1, -1
}
