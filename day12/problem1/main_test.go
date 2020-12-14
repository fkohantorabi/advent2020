package main

import (
	"testing"
)

func TestTurnLeft(t *testing.T) {
	result := turn("E", 90)

	if result != "N" {
		t.Errorf("Expected to be facing north but got %s", result)
	}
}

func TestTurnRight(t *testing.T) {
	result := turn("E", -270)

	if result != "N" {
		t.Errorf("Expected to be facing north but got %s", result)
	}
}
