package main

import (
	"testing"
)

func TestTurnRight(t *testing.T) {
	x, y := turn(-10, 4, -90)

	if x != -4 {
		t.Errorf("Expected -4 but got %d", x)
	}

	if y != -10 {
		t.Errorf("Expected -10 but got %d", y)
	}
}

func TestTurnLeft(t *testing.T) {
	x, y := turn(-10, 4, 90)

	if x != 4 {
		t.Errorf("Expected 4 but got %d", x)
	}

	if y != 10 {
		t.Errorf("Expected 10 but got %d", y)
	}
}
func TestTurnLeft180(t *testing.T) {
	x, y := turn(-10, 4, 180)

	if x != 10 {
		t.Errorf("Expected 10 but got %d", x)
	}

	if y != -4 {
		t.Errorf("Expected -4 but got %d", y)
	}
}

func TestTurnRight180(t *testing.T) {
	x, y := turn(-10, 4, -180)

	if x != 10 {
		t.Errorf("Expected 10 but got %d", x)
	}

	if y != -4 {
		t.Errorf("Expected -4 but got %d", y)
	}
}

func TestCalculateCoordinates(t *testing.T) {
	lines := []string{"F10", "N3", "F7", "R90", "F11"}

	x, y, vx, vy := calculateCoordinates(lines)

	if x != -214 || y != -72 {
		t.Errorf("Expecting ship to be at (-214, -72) but got (%d,%d)", x, y)
	}

	if vx != -4 || vy != -10 {
		t.Errorf("Expecting waypoint to be at (-4, -10) but got (%d,%d)", x, y)
	}
}
