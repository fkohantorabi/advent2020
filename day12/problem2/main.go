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
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	x, y, _, _ := calculateCoordinates(lines)
	println("Answer is", abs(x)+abs(y))
}

func calculateCoordinates(lines []string) (int, int, int, int) {
	x := 0
	y := 0
	vx, vy := -10, 1

	for _, line := range lines {
		command := line[0:1]
		digits, _ := strconv.Atoi(line[1:])

		if command == "L" {
			vx, vy = turn(vx, vy, digits)
		} else if command == "R" {
			vx, vy = turn(vx, vy, -digits)
		} else if command == "F" {
			x = x + vx*digits
			y = y + vy*digits
		} else {
			vx, vy = move(command, vx, vy, digits)
		}
	}

	return x, y, vx, vy
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

	if x == 0 && y == 0 {
		panic("I yet to code this scenario!")
	}

	return x, y
}

func turn(vx, vy int, degree int) (int, int) {
	turns := int(degree / 90)

	if turns%2 == 0 {
		x := abs(turns) - 3
		y := abs(turns) - 3

		return x * vx, y * vy
	}

	directions := []string{"E", "N", "W", "S"}
	indexX := find(getDirectionX(vx), directions)
	indexY := find(getDirectionY(vy), directions)
	newDirectionIndexX := (4 + indexX + turns) % 4
	newDirectionIndexY := (4 + indexY + turns) % 4

	x := abs(vx)
	y := abs(vy)

	if directions[newDirectionIndexX] == "S" {
		x = -x
	}

	if directions[newDirectionIndexY] == "E" {
		y = -y
	}

	return y, x
}

func getDirectionX(x int) string {
	if x < 0 {
		return "E"
	}

	return "W"
}

func getDirectionY(y int) string {
	if y < 0 {
		return "S"
	}

	return "N"
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
