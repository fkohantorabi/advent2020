package main

import (
	"bufio"
	"os"
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

	groupAnswers := []string{}
	totalAnswers := 0

	for len(lines) > 0 {
		groupAnswers, lines = nextGroupAnswers(lines)
		totalAnswers = totalAnswers + countSharedGroupAnswers(groupAnswers)
	}

	println("Total answers:", totalAnswers)
}

func nextGroupAnswers(lines []string) ([]string, []string) {
	position := 0
	for ; position < len(lines); position++ {
		if strings.TrimSpace(lines[position]) == "" {
			break
		}
	}

	if position < len(lines) {
		return lines[:position], lines[position+1:]
	}

	return lines[:position], []string{}
}

func countSharedGroupAnswers(answers []string) int {
	sharedAnswers := 0b11111111111111111111111111
	result := 0

	for _, answer := range answers {
		sharedAnswers = sharedAnswers & toBits(answer)
	}

	for i := 0; i < 26; i++ {
		if sharedAnswers&(int(1)<<i) > 0 {
			result++
		}
	}

	return result
}

func toBits(answer string) int {
	result := 0
	for i := 0; i < len(answer); i++ {
		bit := int(1) << (answer[i] - "a"[0])
		result = result | bit
	}

	return result
}
