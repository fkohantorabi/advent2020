package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadRecord(t *testing.T) {
	lines := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm"}

	expect := record{ECL: "gry",
		PID: "860033327",
		EYR: "2020",
		HCL: "#fffffd",
		BYR: "1937",
		IYR: "2017",
		CID: "147",
		HGT: "183cm"}

	result := readRecord(lines)

	if !cmp.Equal(result, expect) {
		t.Errorf("The structure recieved is not valid. %+v", result)
	}
}

func TestReadRecordMissingFields(t *testing.T) {
	lines := []string{
		"ecl:gry pid:860033327 hcl:#fffffd",
		"byr:1937 hgt:183cm"}

	expect := record{ECL: "gry",
		PID: "860033327",
		HCL: "#fffffd",
		BYR: "1937",
		HGT: "183cm"}

	result := readRecord(lines)

	if !cmp.Equal(result, expect) {
		t.Errorf("The structure recieved is not valid. %+v", result)
	}
}

func TestNextRecordLines(t *testing.T) {
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

	result, remaining := nextRecordLines(lines)

	if !cmp.Equal(result, expected) {
		t.Errorf("Expected %+v array but got %+v", expected, result)
	}

	if !cmp.Equal(remaining, expectedRemaining) {
		t.Errorf("Expected %+v remaining array but got %+v", expectedRemaining, remaining)
	}
}

func TestNextRecordLinesEndOfFile(t *testing.T) {
	lines := []string{
		"A",
		"B"}

	expected := []string{
		"A",
		"B"}

	expectedRemaining := []string{}
	result, remaining := nextRecordLines(lines)

	if !cmp.Equal(result, expected) {
		t.Errorf("Expected %+v array but got %+v", expected, result)
	}

	if !cmp.Equal(remaining, expectedRemaining) {
		t.Errorf("Expected %+v remaining array but got %+v", expectedRemaining, remaining)
	}
}
