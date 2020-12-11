package main

import (
	"bufio"
	"os"
)

const empty, floor, occupied = 0, 1, 2

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seatMap := [][]int{}

	for scanner.Scan() {
		row := toRow(scanner.Text())
		seatMap = append(seatMap, row)
	}

	stabilizedSeatMap := stabilize(seatMap)
	println("# of occupied seats", countOccupied(stabilizedSeatMap))
}

func countOccupied(seatMap [][]int) int {
	result := 0

	for _, row := range seatMap {
		for _, seat := range row {
			if seat == occupied {
				result++
			}
		}
	}

	return result
}

func stabilize(seatMap [][]int) [][]int {
	changes := 1
	for changes > 0 {
		seatMap, changes = step(seatMap)
	}

	return seatMap
}

func step(seatMap [][]int) ([][]int, int) {
	changes := 0
	resultMap := copyArray(seatMap)

	for y := range seatMap {
		for x := range seatMap[y] {
			neighbors := countNeighbors(seatMap, x, y)

			if seatMap[y][x] == empty && neighbors == 0 {
				resultMap[y][x] = occupied
				changes++
			} else if seatMap[y][x] == occupied && neighbors > 3 {
				resultMap[y][x] = empty
				changes++
			}
		}
	}

	return resultMap, changes
}

func copyArray(arr [][]int) [][]int {
	result := make([][]int, 0, len(arr))

	for _, inner := range arr {
		copiedInner := make([]int, len(inner))
		copy(copiedInner, inner)
		result = append(result, copiedInner)
	}

	return result
}

func countNeighbors(seatMap [][]int, px, py int) int {
	result := 0

	for x := max(px-1, 0); x < min(px+2, len(seatMap[0])); x++ {
		for y := max(py-1, 0); y < min(py+2, len(seatMap)); y++ {
			if x == px && y == py {
				continue
			}

			if seatMap[y][x] == occupied {
				result++
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func toRow(row string) []int {
	result := make([]int, len(row))

	for i, char := range row {
		if char == []rune("L")[0] {
			result[i] = empty
		} else {
			result[i] = floor
		}
	}

	return result
}
