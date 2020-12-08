package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type rule struct {
	name string

	content map[string]int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := map[string]rule{}

	for scanner.Scan() {
		r := toRule(scanner.Text())
		rules[r.name] = r
	}

	count := 0
	for r := range rules {
		if canHold(r, "shiny gold", rules) {
			count++
		}
	}

	println("Total bad that can contain shiny gold:", count)
}

func toRule(input string) rule {
	result := rule{}

	nameRegex, _ := regexp.Compile("([\\w ]+) bags contain")
	result.name = nameRegex.FindStringSubmatch(input)[1]
	result.content = map[string]int{}

	contentRegex, _ := regexp.Compile("(?:, )?(\\d+) ([\\w ]+) bags?")
	matches := contentRegex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		size, _ := strconv.Atoi(match[1])
		result.content[match[2]] = result.content[match[2]] + size
	}

	return result
}

func canHold(outerBag, bag string, rules map[string]rule) bool {
	r := rules[outerBag]
	_, canHoldDirectly := r.content[bag]

	if canHoldDirectly {
		return true
	}

	for innerBag := range r.content {
		if canHold(innerBag, bag, rules) {
			return true
		}
	}

	return false
}
