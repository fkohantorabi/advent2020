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
	for i := range arr[:len(arr)-1] {
		remaining := arr[i+1 : len(arr)]
		for j := range remaining {
			if (arr[i] + remaining[j]) == 2020 {
				return arr[i] * remaining[j]
			}
		}
	}

	return -1
}
