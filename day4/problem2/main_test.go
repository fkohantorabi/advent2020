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

func TestIsBryValid(t *testing.T) {
	result := isBryValid("2000")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsBryInvalidUpper(t *testing.T) {
	result := isBryValid("2003")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsBryInvalidLower(t *testing.T) {
	result := isBryValid("1919")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsBryInvalidEmpty(t *testing.T) {
	result := isBryValid("")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsIyrValid(t *testing.T) {
	result := isIyrValid("2015")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsIyInvalidUpper(t *testing.T) {
	result := isIyrValid("2021")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsIyrInvalidLower(t *testing.T) {
	result := isIyrValid("2009")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsIyrInvalidEmpty(t *testing.T) {
	result := isIyrValid("")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsEyrValid(t *testing.T) {
	result := isEyrValid("2025")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsEyrInvalidUpper(t *testing.T) {
	result := isEyrValid("2031")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsEyrInvalidLower(t *testing.T) {
	result := isEyrValid("2019")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsEyrInvalidEmpty(t *testing.T) {
	result := isEyrValid("")

	if result {
		t.Error("Should be invalid")
	}
}
func TestIsHgtCmValid(t *testing.T) {
	result := isHgtValid("160cm")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsHgtCmInvalidUpper(t *testing.T) {
	result := isHgtValid("194cm")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHgtCmInvalidLower(t *testing.T) {
	result := isHgtValid("149cm")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHgtInValid(t *testing.T) {
	result := isHgtValid("60in")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsHgtInInvalidUpper(t *testing.T) {
	result := isHgtValid("77in")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHgtInInvalidLower(t *testing.T) {
	result := isHgtValid("58in")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHgtOtherInvalid(t *testing.T) {
	result := isHgtValid("100pt")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHgtOtherInvalidNumberOnly(t *testing.T) {
	result := isHgtValid("100")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHgtEmptyInvalid(t *testing.T) {
	result := isHgtValid("")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHclValid(t *testing.T) {
	result := isHclValid("#452de5")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsHclInvalidChar(t *testing.T) {
	result := isHclValid("#452ke5")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsHclEmptyInvalid(t *testing.T) {
	result := isHclValid("")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsEclValid(t *testing.T) {
	result := isEclValid("blu")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsEclEmptyInvalid(t *testing.T) {
	result := isEclValid("")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsPidValid(t *testing.T) {
	result := isPidValid("000000001")

	if !result {
		t.Error("Should be valid")
	}
}

func TestIsPidEmptyInvalid(t *testing.T) {
	result := isPidValid("")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsPidAlphabeticalInvalid(t *testing.T) {
	result := isPidValid("1234a6789")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsPidTooShortInvalid(t *testing.T) {
	result := isPidValid("6789")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsPidTooLongInvalid(t *testing.T) {
	result := isPidValid("1234567890")

	if result {
		t.Error("Should be invalid")
	}
}

func TestIsValidInvalidSample1(t *testing.T) {
	lines := []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"}

	result := isRecordValid(lines)

	if result {
		t.Errorf("Input should be invalid")
	}
}

func TestIsValidInvalidSample2(t *testing.T) {
	lines := []string{
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946"}

	result := isRecordValid(lines)

	if result {
		t.Errorf("Input should be invalid")
	}
}

func TestIsValidInvalidSample3(t *testing.T) {
	lines := []string{
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277"}

	result := isRecordValid(lines)

	if result {
		t.Errorf("Input should be invalid")
	}
}

func TestIsValidInvalidSample4(t *testing.T) {
	lines := []string{
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007"}

	result := isRecordValid(lines)

	if result {
		t.Errorf("Input should be invalid")
	}
}
