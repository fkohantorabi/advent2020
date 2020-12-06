package main

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	result := getRow("BFFFBBFRRR")

	if result != 70 {
		t.Errorf("Expecting row 70 but got %d", result)
	}

	result = getRow("FFFBBBFRRR")

	if result != 14 {
		t.Errorf("Expecting row 14 but got %d", result)
	}

	result = getRow("BBFFBBFRLL")

	if result != 102 {
		t.Errorf("Expecting row 102 but got %d", result)
	}
}

func TestGetColumn(t *testing.T) {
	result := getColumn("BFFFBBFRRR")

	if result != 7 {
		t.Errorf("Expecting column 7 but got %d", result)
	}

	result = getColumn("FFFBBBFRRR")

	if result != 7 {
		t.Errorf("Expecting column 7 but got %d", result)
	}

	result = getColumn("BBFFBBFRLL")

	if result != 4 {
		t.Errorf("Expecting column 4 but got %d", result)
	}
}
