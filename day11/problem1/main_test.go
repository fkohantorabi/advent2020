package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStep(t *testing.T) {
	seatMap := [][]int{
		{occupied, floor, occupied, occupied, floor, occupied, occupied, floor, occupied, occupied},
		{occupied, occupied, occupied, occupied, occupied, occupied, occupied, floor, occupied, occupied},
		{occupied, floor, occupied, floor, occupied, floor, floor, occupied, floor, floor},
		{occupied, occupied, occupied, occupied, floor, occupied, occupied, floor, occupied, occupied},
		{occupied, floor, occupied, occupied, floor, occupied, occupied, floor, occupied, occupied},
		{occupied, floor, occupied, occupied, occupied, occupied, occupied, floor, occupied, occupied},
		{floor, floor, occupied, floor, occupied, floor, floor, floor, floor, floor},
		{occupied, occupied, occupied, occupied, occupied, occupied, occupied, occupied, occupied, occupied},
		{occupied, floor, occupied, occupied, occupied, occupied, occupied, occupied, floor, occupied},
		{occupied, floor, occupied, occupied, occupied, occupied, occupied, floor, occupied, occupied}}

	expected := [][]int{
		{occupied, floor, empty, empty, floor, empty, occupied, floor, occupied, occupied},
		{occupied, empty, empty, empty, empty, empty, empty, floor, empty, occupied},
		{empty, floor, empty, floor, empty, floor, floor, empty, floor, floor},
		{occupied, empty, empty, empty, floor, empty, empty, floor, empty, occupied},
		{occupied, floor, empty, empty, floor, empty, empty, floor, empty, empty},
		{occupied, floor, empty, empty, empty, empty, occupied, floor, occupied, occupied},
		{floor, floor, empty, floor, empty, floor, floor, floor, floor, floor},
		{occupied, empty, empty, empty, empty, empty, empty, empty, empty, occupied},
		{occupied, floor, empty, empty, empty, empty, empty, empty, floor, empty},
		{occupied, floor, occupied, empty, empty, empty, empty, floor, occupied, occupied}}

	result, changes := step(seatMap)

	if !cmp.Equal(result, expected) {
		t.Error("Step is not producing the right result")
	}

	if changes != 51 {
		t.Errorf("Expected 51 changes but got %d", changes)
	}
}
