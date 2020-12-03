package main

import "testing"

func TestParseLine(t *testing.T) {
	first, second, letter, text := ParseLine("1-2 a: abcd")

	t.Logf("first=%d, second=%d, letter=%U, text=%s", first, second, letter, text)

	if first != 1 {
		t.Errorf("first is %d but expected 1", first)
	}

	if second != 2 {
		t.Errorf("second is %d but expected 2", second)
	}

	if letter != []rune("a")[0] {
		t.Errorf("letter is %U but expected a", letter)
	}

	if text != "abcd" {
		t.Errorf("text is %s but expected abcd", text)
	}
}

func TestMatchesInFirstPosition(t *testing.T) {
	result := Matches(1, 3, []rune("a")[0], "aacd")

	if !result {
		t.Error("should match abcd")
	}
}

func TestMatchesInSecondPosition(t *testing.T) {
	result := Matches(1, 3, []rune("a")[0], "daad")

	if !result {
		t.Error("should match daad")
	}
}

func TestMatchesTooManyMatched(t *testing.T) {
	result := Matches(1, 3, []rune("a")[0], "abad")

	if result {
		t.Error("should not match abad")
	}
}

func TestMatchesStringTooShort(t *testing.T) {
	result := Matches(4, 9, []rune("a")[0], "abc")

	if result {
		t.Error("should not match abc")
	}
}
