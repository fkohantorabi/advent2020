package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStep(t *testing.T) {
	seatMap := [][]int{
		{occupied, floor, empty, empty, floor, empty, empty, floor, empty, occupied},
		{occupied, empty, empty, empty, empty, empty, empty, floor, empty, empty},
		{empty, floor, empty, floor, empty, floor, floor, empty, floor, floor},
		{empty, empty, empty, empty, floor, empty, empty, floor, empty, empty},
		{empty, floor, empty, empty, floor, empty, empty, floor, empty, empty},
		{empty, floor, empty, empty, empty, empty, empty, floor, empty, empty},
		{floor, floor, empty, floor, empty, floor, floor, floor, floor, floor},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty, occupied},
		{occupied, floor, empty, empty, empty, empty, empty, empty, floor, empty},
		{occupied, floor, empty, empty, empty, empty, empty, floor, empty, occupied}}

	expected := [][]int{
		{occupied, floor, empty, occupied, floor, occupied, occupied, floor, empty, occupied},
		{occupied, empty, occupied, occupied, occupied, occupied, occupied, floor, empty, empty},
		{empty, floor, occupied, floor, occupied, floor, floor, occupied, floor, floor},
		{occupied, occupied, empty, occupied, floor, occupied, occupied, floor, occupied, occupied},
		{occupied, floor, occupied, occupied, floor, occupied, empty, floor, occupied, occupied},
		{occupied, floor, occupied, occupied, occupied, occupied, occupied, floor, occupied, empty},
		{floor, floor, occupied, floor, occupied, floor, floor, floor, floor, floor},
		{empty, empty, empty, occupied, occupied, occupied, occupied, empty, empty, occupied},
		{occupied, floor, empty, occupied, occupied, occupied, occupied, occupied, floor, empty},
		{occupied, floor, empty, occupied, occupied, occupied, occupied, floor, empty, occupied}}

	result, changes := step(seatMap)

	if !cmp.Equal(result, expected) {
		t.Error("Step is not producing the right result")
	}

	if changes != 46 {
		t.Errorf("Expected 46 changes but got %d", changes)
	}
}
