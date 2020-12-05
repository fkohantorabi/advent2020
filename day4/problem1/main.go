package main

import (
	"bufio"
	"os"
	"reflect"
	"regexp"
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
		r := readRecord(recordLines)

		if r.BYR != "" && r.IYR != "" && r.EYR != "" && r.HGT != "" &&
			r.HCL != "" && r.ECL != "" && r.PID != "" {
			valid++
		}
	}

	println("Valid passports:", valid)
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
