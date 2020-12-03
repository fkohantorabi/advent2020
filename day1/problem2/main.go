package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]int, 0, 200)

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, line)
	}

	answer := search(arr)
	println("Answer is", answer)
}

func search(arr []int) int {
	for i := range arr[:len(arr)-2] {
		second := arr[i+1 : len(arr)-1]
		for j := range second {
			third := second[j+1 : len(second)]
			for k := range third {
				if (arr[i] + second[j] + third[k]) == 2020 {
					return arr[i] * second[j] * third[k]
				}
			}
		}
	}

	return -1
}
