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

	permutations := countPermutations(jolts)
	println("Answer is", permutations)
}

func countPermutations(jolts []int) int {
	sort.Ints(jolts)
	paths := make([]int, len(jolts), len(jolts))
	paths[len(paths)-1] = 1

	for i := len(jolts) - 2; i >= 0; i-- {
		for j := i + 1; j < len(jolts); j++ {
			if jolts[j]-jolts[i] <= 3 {
				paths[i] = paths[i] + paths[j]
			} else {
				break
			}
		}
	}

	result := 0
	for i, jolt := range jolts {
		if jolt <= 3 {
			result = result + paths[i]
		}
	}

	return result
}
