package main

import (
	"bufio"
	"os"
	"strconv"
)

const empty, floor, occupied = 0, 1, 2

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	direction := "E"
	x := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		command := line[0:1]
		digits, _ := strconv.Atoi(line[1:])

		if command == "L" {
			direction = turn(direction, digits)
		} else if command == "R" {
			direction = turn(direction, -digits)
		} else if command == "F" {
			x, y = move(direction, x, y, digits)
		} else {
			x, y = move(command, x, y, digits)
		}
	}

	println("Answer is", abs(x)+abs(y))
}

func move(command string, x, y, digits int) (int, int) {
	if command == "N" {
		y = y + digits
	}

	if command == "S" {
		y = y - digits
	}

	if command == "E" {
		x = x - digits
	}

	if command == "W" {
		x = x + digits
	}

	return x, y
}

func turn(direction string, degree int) string {
	directions := []string{"E", "N", "W", "S"}
	index := find(direction, directions)

	if index == -1 {
		panic("Direction is invalid")
	}

	newDirectionIndex := (4 + index + (degree / 90)) % 4
	return directions[newDirectionIndex]
}

func find(directionToFind string, directions []string) int {
	for i, direction := range directions {
		if direction == directionToFind {
			return i
		}
	}

	return -1
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
