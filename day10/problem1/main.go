package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	jolts := []int{}

	for scanner.Scan() {
		jolt, _ := strconv.Atoi(scanner.Text())
		jolts = append(jolts, jolt)
	}

	ones, threes := findJoltDifferences(jolts)

	println("Answer is", ones*threes)
}

func findJoltDifferences(jolts []int) (int, int) {
	sort.Ints(jolts)
	ones, threes := 0, 1
	prev := 0

	for _, jolt := range jolts {
		if jolt-prev > 3 {
			panic("I didn't expect this")
		}

		if jolt-prev == 3 {
			threes++
		}

		if jolt-prev == 1 {
			ones++
		}

		prev = jolt
	}

	return ones, threes
}
