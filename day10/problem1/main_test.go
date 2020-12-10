package main

import (
	"testing"
)

func TestFindJoltDifferences(t *testing.T) {
	jolts := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45,
		19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}

	ones, threes := findJoltDifferences(jolts)

	if ones != 22 {
		t.Errorf("expected 22 ones but got %d", ones)
	}

	if threes != 10 {
		t.Errorf("expected 10 threes but got %d", threes)
	}
}
