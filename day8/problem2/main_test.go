package main

import (
	"testing"
)

func TestInstruction(t *testing.T) {
	instruction, num := toInstruction("acc +10", false)

	if instruction != "acc" {
		t.Errorf("Expected acc but got %+v", instruction)
	}

	if num != 10 {
		t.Errorf("Expected 10 but got %d", num)
	}
}

func TestInstructionNegtive(t *testing.T) {
	instruction, num := toInstruction("acc -10", false)

	if instruction != "acc" {
		t.Errorf("Expected acc but got %+v", instruction)
	}

	if num != -10 {
		t.Errorf("Expected -10 but got %d", num)
	}
}
func TestCorrectProgram(t *testing.T) {
	lines := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6"}

	result := correctProgram(lines)

	if result != 8 {
		t.Errorf("Expecting accumolated value 8 but got %d", result)
	}
}
