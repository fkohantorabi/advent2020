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
		if Matches(lower, upper, letter, text) {
			validCount++
		}
	}

	println("Total valid lines are", validCount)
}

func ParseLine(input string) (int, int, rune, string) {
	r, _ := regexp.Compile("(\\d+)-(\\d+) (\\w): (\\w+)")
	match := r.FindStringSubmatch(input)

	first, _ := strconv.Atoi(match[1])
	second, _ := strconv.Atoi(match[2])
	letter, _ := utf8.DecodeRuneInString(match[3])

	return first, second, letter, match[4]
}

func Matches(first, second int, letter rune, text string) bool {
	position := 0
	count := 0

	for _, char := range text {
		position++

		if (position == first || position == second) && char == letter {
			count++
		}
	}

	return count == 1
}
