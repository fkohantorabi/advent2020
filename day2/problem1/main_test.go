package main

import "testing"

func TestParseLine(t *testing.T) {
	lower, upper, letter, text := ParseLine("1-2 a: abcd")

	t.Logf("lower=%d, upper=%d, letter=%U, text=%s", lower, upper, letter, text)

	if lower != 1 {
		t.Errorf("lower is %d but expected 1", lower)
	}

	if upper != 2 {
		t.Errorf("upper is %d but expected 2", upper)
	}

	if letter != []rune("a")[0] {
		t.Errorf("letter is %U but expected a", letter)
	}

	if text != "abcd" {
		t.Errorf("text is %s but expected abcd", text)
	}
}

func TestCount(t *testing.T) {
	result := Count([]rune("a")[0], "abac")

	if result != 2 {
		t.Errorf("Expected 2 but got %d", result)
	}
}
