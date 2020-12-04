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

	result := CountTrees(treeMap, 1, 1) *
		CountTrees(treeMap, 3, 1) *
		CountTrees(treeMap, 5, 1) *
		CountTrees(treeMap, 7, 1) *
		CountTrees(treeMap, 1, 2)

	println("Result:", result)
}

func CountTrees(treeMap [][]bool, slopeX, slopeY int) uint64 {
	result := uint64(0)
	x := 0

	for y := slopeY; y < len(treeMap); y = y + slopeY {
		x = x + slopeX

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
