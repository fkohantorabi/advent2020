package main

import (
	"testing"
)

func TestInstruction(t *testing.T) {
	instruction, num := toInstruction("acc +10")

	if instruction != "acc" {
		t.Errorf("Expected acc but got %+v", instruction)
	}

	if num != 10 {
		t.Errorf("Expected 10 but got %d", num)
	}
}

func TestInstructionNegtive(t *testing.T) {
	instruction, num := toInstruction("acc -10")

	if instruction != "acc" {
		t.Errorf("Expected acc but got %+v", instruction)
	}

	if num != -10 {
		t.Errorf("Expected -10 but got %d", num)
	}
}
