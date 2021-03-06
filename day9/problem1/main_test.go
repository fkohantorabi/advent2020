package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	numbers := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95}

	result := isValid(numbers, 5, 5)

	if !result {
		t.Errorf("expected valid response but got invalid")
	}
}
