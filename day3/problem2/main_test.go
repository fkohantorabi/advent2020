package main

import (
	"reflect"
	"testing"
)

func TestToMapLine(t *testing.T) {
	results := ToMapLine("##..#")
	expected := []bool{true, true, false, false, true}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expecting an array of 5 elements but got %d elements", len(results))
	}

}

func TestCountTrees3by1Slope(t *testing.T) {
	treeMap := [][]bool{
		{false, false, false, false},
		{false, false, false, true},
		{false, false, true, false},
		{false, true, false, false},
		{false, false, false, false},
		{false, false, false, true},
	}

	result := CountTrees(treeMap, 3, 1)

	if result != 4 {
		t.Errorf("Expecting 4 trees but found %d", result)
	}
}

func TestCountTrees2by3Slope(t *testing.T) {
	treeMap := [][]bool{
		{false, false, false, false},
		{false, false, false, true},
		{false, false, true, false},
		{false, true, true, false},
		{false, false, false, false},
		{false, false, false, true},
	}

	result := CountTrees(treeMap, 2, 3)

	if result != 1 {
		t.Errorf("Expecting 1 trees but found %d", result)
	}
}
