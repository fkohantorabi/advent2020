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
			} else if seatMap[y][x] == occupied && neighbors > 4 {
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
	neighbors := []int{0, 0, 0, 0, 0, 0, 0, 0}
	noEmptySeatInsight := []int{1, 1, 1, 1, 1, 1, 1, 1}
	maxX := len(seatMap[0]) - 1
	maxY := len(seatMap) - 1

	for i := 1; (px-i >= 0 || px+i <= maxX || py-i >= 0 || py+i <= maxY) && needToSearchMore(neighbors, noEmptySeatInsight); i++ {
		neighbors[0] = neighbors[0] | (noEmptySeatInsight[0] & isOccupied(seatMap, px, py, px-i, py))
		noEmptySeatInsight[0] = noEmptySeatInsight[0] & isNotEmpty(seatMap, px, py, px-i, py)

		neighbors[1] = neighbors[1] | (noEmptySeatInsight[1] & isOccupied(seatMap, px, py, px-i, py-i))
		noEmptySeatInsight[1] = noEmptySeatInsight[1] & isNotEmpty(seatMap, px, py, px-i, py-i)

		neighbors[2] = neighbors[2] | (noEmptySeatInsight[2] & isOccupied(seatMap, px, py, px, py-i))
		noEmptySeatInsight[2] = noEmptySeatInsight[2] & isNotEmpty(seatMap, px, py, px, py-i)

		neighbors[3] = neighbors[3] | (noEmptySeatInsight[3] & isOccupied(seatMap, px, py, px+i, py-i))
		noEmptySeatInsight[3] = noEmptySeatInsight[3] & isNotEmpty(seatMap, px, py, px+i, py-i)

		neighbors[4] = neighbors[4] | (noEmptySeatInsight[4] & isOccupied(seatMap, px, py, px+i, py))
		noEmptySeatInsight[4] = noEmptySeatInsight[4] & isNotEmpty(seatMap, px, py, px+i, py)

		neighbors[5] = neighbors[5] | (noEmptySeatInsight[5] & isOccupied(seatMap, px, py, px+i, py+i))
		noEmptySeatInsight[5] = noEmptySeatInsight[5] & isNotEmpty(seatMap, px, py, px+i, py+i)

		neighbors[6] = neighbors[6] | (noEmptySeatInsight[6] & isOccupied(seatMap, px, py, px, py+i))
		noEmptySeatInsight[6] = noEmptySeatInsight[6] & isNotEmpty(seatMap, px, py, px, py+i)

		neighbors[7] = neighbors[7] | (noEmptySeatInsight[7] & isOccupied(seatMap, px, py, px-i, py+i))
		noEmptySeatInsight[7] = noEmptySeatInsight[7] & isNotEmpty(seatMap, px, py, px-i, py+i)
	}

	return sum(neighbors)
}

func needToSearchMore(neighbors []int, noEmptySeatInSight []int) bool {

	for i := range neighbors {
		if neighbors[i] == 0 && noEmptySeatInSight[i] == 1 {
			return true
		}
	}

	return false
}

func isOccupied(seatMap [][]int, px, py, x, y int) int {
	if x == px && y == py {
		return 0
	}

	maxX := len(seatMap[0]) - 1
	maxY := len(seatMap) - 1

	if x < 0 || x > maxX || y < 0 || y > maxY {
		return 0
	}

	if seatMap[y][x] == occupied {
		return 1
	}

	return 0
}

func isNotEmpty(seatMap [][]int, px, py, x, y int) int {
	if x == px && y == py {
		return 1
	}

	maxX := len(seatMap[0]) - 1
	maxY := len(seatMap) - 1

	if x < 0 || x > maxX || y < 0 || y > maxY {
		return 1
	}

	if seatMap[y][x] == empty {
		return 0
	}

	return 1
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
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
