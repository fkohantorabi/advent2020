package main

import (
	"bufio"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type record struct {
	BYR string
	IYR string
	EYR string
	HGT string
	HCL string
	ECL string
	PID string
	CID string
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	recordLines := []string{}
	valid := 0

	for len(lines) > 0 {
		recordLines, lines = nextRecordLines(lines)

		if isRecordValid(recordLines) {
			valid++
			println()
		}
	}

	println("Valid passports:", valid)
}

func isRecordValid(recordLines []string) bool {
	r := readRecord(recordLines)

	return isBryValid(r.BYR) && isIyrValid(r.IYR) && isEyrValid(r.EYR) && isHgtValid(r.HGT) &&
		isHclValid(r.HCL) && isEclValid(r.ECL) && isPidValid(r.PID)
}

func nextRecordLines(lines []string) ([]string, []string) {
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

func readRecord(lines []string) record {
	result := record{}
	structType := reflect.ValueOf(&result).Elem()

	for i := range lines {
		for f := 0; f < structType.Type().NumField(); f++ {
			fieldName := structType.Type().Field(f).Name
			fieldValue := readField(lines[i], strings.ToLower(fieldName))

			if fieldValue != "" {
				structType.Field(f).SetString(fieldValue)
			}
		}
	}

	return result
}

func readField(line string, field string) string {
	r, _ := regexp.Compile(field + ":(\\S+)")

	if r.MatchString(line) {
		match := r.FindStringSubmatch(line)
		return match[1]
	}

	return ""
}

func isBryValid(input string) bool {
	if input == "" {
		return false
	}

	value, err := strconv.Atoi(input)
	return err == nil && value >= 1920 && value <= 2002
}

func isIyrValid(input string) bool {
	if input == "" {
		return false
	}

	value, err := strconv.Atoi(input)
	return err == nil && value >= 2010 && value <= 2020
}

func isEyrValid(input string) bool {
	if input == "" {
		return false
	}

	value, err := strconv.Atoi(input)
	return err == nil && value >= 2020 && value <= 2030
}

func isHgtValid(input string) bool {
	if input == "" {
		return false
	}

	cm, _ := regexp.Compile("^(\\d+)cm$")
	if cm.MatchString(input) {
		match := cm.FindStringSubmatch(input)
		value, err := strconv.Atoi(match[1])
		return err == nil && value >= 150 && value <= 193
	}

	in, _ := regexp.Compile("^(\\d+)in$")
	if in.MatchString(input) {
		match := in.FindStringSubmatch(input)
		value, err := strconv.Atoi(match[1])
		return err == nil && value >= 59 && value <= 76
	}

	return false
}

func isHclValid(input string) bool {
	if input == "" {
		return false
	}

	r, _ := regexp.Compile("^#[0-9a-f]{6}$")
	return r.MatchString(input)
}

func isEclValid(input string) bool {
	validValues := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, validValue := range validValues {
		if validValue == input {
			return true
		}
	}

	return false
}

func isPidValid(input string) bool {
	if input == "" {
		return false
	}

	r, _ := regexp.Compile("^\\d{9}$")
	return r.MatchString(input)
}
