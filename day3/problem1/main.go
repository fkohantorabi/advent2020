package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	treeMap := [][]bool{}

	for scanner.Scan() {
		mapLine := ToMapLine(scanner.Text())
		treeMap = append(treeMap, mapLine)
	}

	count := CountTrees(treeMap)

	println("Trees visited:", count)
}

func CountTrees(treeMap [][]bool) int {
	result := 0
	x := 0

	for y := 1; y < len(treeMap); y++ {
		x = x + 3

		if treeMap[y][x%len(treeMap[0])] {
			result++
		}
	}

	return result
}

func ToMapLine(line string) []bool {
	results := make([]bool, len(line))
	i := 0

	for _, char := range line {
		results[i] = char == []rune("#")[0]
		i++
	}

	return results
}
