package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validCount := 0

	for scanner.Scan() {
		lower, upper, letter, text := ParseLine(scanner.Text())
		count := Count(letter, text)

		if count >= lower && count <= upper {
			validCount++
		}
	}

	println("Total valid lines are", validCount)
}

func ParseLine(input string) (int, int, rune, string) {
	r, _ := regexp.Compile("(\\d+)-(\\d+) (\\w): (\\w+)")
	match := r.FindStringSubmatch(input)

	lower, _ := strconv.Atoi(match[1])
	upper, _ := strconv.Atoi(match[2])
	letter, _ := utf8.DecodeRuneInString(match[3])

	return lower, upper, letter, match[4]
}

func Count(letter rune, text string) int {
	result := 0

	for _, char := range text {
		if char == letter {
			result++
		}
	}

	return result
}
