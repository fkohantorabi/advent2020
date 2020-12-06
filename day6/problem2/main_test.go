package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextGroupAnswers(t *testing.T) {
	lines := []string{
		"A",
		"B",
		" ",
		"C",
		"D"}

	expected := []string{
		"A",
		"B"}

	expectedRemaining := []string{
		"C",
		"D"}

	result, remaining := nextGroupAnswers(lines)

	if !cmp.Equal(result, expected) {
		t.Errorf("Expected %+v array but got %+v", expected, result)
	}

	if !cmp.Equal(remaining, expectedRemaining) {
		t.Errorf("Expected %+v remaining array but got %+v", expectedRemaining, remaining)
	}
}

func TestNextGroupAnswersEndOfFile(t *testing.T) {
	lines := []string{
		"A",
		"B"}

	expected := []string{
		"A",
		"B"}

	expectedRemaining := []string{}
	result, remaining := nextGroupAnswers(lines)

	if !cmp.Equal(result, expected) {
		t.Errorf("Expected %+v array but got %+v", expected, result)
	}

	if !cmp.Equal(remaining, expectedRemaining) {
		t.Errorf("Expected %+v remaining array but got %+v", expectedRemaining, remaining)
	}
}

func TestCountSharedGroupAnswers(t *testing.T) {
	answers := []string{
		"abhz",
		"achzk",
		"ezah"}

	result := countSharedGroupAnswers(answers)

	if result != 3 {
		t.Errorf("Expected 3 answers but got %d", result)
	}
}
