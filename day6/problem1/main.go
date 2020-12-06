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
		totalAnswers = totalAnswers + countAnsweredQuestions(groupAnswers)
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

func countAnsweredQuestions(answers []string) int {
	alreadyAnswered := 0
	result := 0

	for _, answer := range answers {
		for i := 0; i < len(answer); i++ {
			bit := int(1) << (answer[i] - "a"[0])
			if alreadyAnswered&bit == 0 {
				result++
				alreadyAnswered = alreadyAnswered | bit
			}
		}
	}

	return result
}
