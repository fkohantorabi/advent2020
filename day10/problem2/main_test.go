package main

import (
	"testing"
)

func TestCountPermutations(t *testing.T) {
	jolts := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45,
		19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}

	results := countPermutations(jolts)

	if results != 19208 {
		t.Errorf("expected 19208 permutations but got %d", results)
	}
}

func TestCountPermutationsSmallSet(t *testing.T) {
	jolts := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

	results := countPermutations(jolts)

	if results != 8 {
		t.Errorf("expected 8 permutations but got %d", results)
	}
}
