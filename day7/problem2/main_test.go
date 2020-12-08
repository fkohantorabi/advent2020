package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToRule(t *testing.T) {
	input := "laid coral bags contain 1 bright gray bag"
	result := toRule(input)

	expected := rule{
		name:    "laid coral",
		content: map[string]int{"bright gray": 1}}

	if !cmp.Equal(result, expected, cmp.AllowUnexported(rule{})) {
		t.Errorf("Expected %+v rule but got %+v", expected, result)
	}
}

func TestToRuleSingleBagPlural(t *testing.T) {
	input := "laid coral bags contain 2 bright gray bags"
	result := toRule(input)

	expected := rule{
		name:    "laid coral",
		content: map[string]int{"bright gray": 2}}

	if !cmp.Equal(result, expected, cmp.AllowUnexported(rule{})) {
		t.Errorf("Expected %+v rule but got %+v", expected, result)
	}
}

func TestToRuleMultipleBags(t *testing.T) {
	input := "plaid crimson bags contain 2 muted black bags, 3 drab turquoise bags, 2 mirrored beige bags."
	result := toRule(input)

	expected := rule{
		name:    "plaid crimson",
		content: map[string]int{"muted black": 2, "drab turquoise": 3, "mirrored beige": 2}}

	if !cmp.Equal(result, expected, cmp.AllowUnexported(rule{})) {
		t.Errorf("Expected %+v rule but got %+v", expected, result)
	}
}

func TestCountBags(t *testing.T) {
	rules := map[string]rule{
		"silver": rule{"silver", map[string]int{"green": 1, "blue": 5}},
		"green":  rule{"green", map[string]int{"red": 2}},
		"blue":   rule{"blue", map[string]int{"red": 4}},
		"red":    rule{"red", map[string]int{}}}

	result := countBags("silver", rules, false)

	if result != 28 {
		t.Errorf("Bag count should be 22 but was %d", result)
	}
}

func TestCountBagsSiteExample(t *testing.T) {
	rules := map[string]rule{
		"shiny gold":  rule{"shiny gold", map[string]int{"dark red": 2}},
		"dark red":    rule{"dark red", map[string]int{"dark orange": 2}},
		"dark orange": rule{"dark orange", map[string]int{"dark yellow": 2}},
		"dark yellow": rule{"dark yellow", map[string]int{"dark green": 2}},
		"dark green":  rule{"dark green", map[string]int{"dark blue": 2}},
		"dark blue":   rule{"dark blue", map[string]int{"dark violet": 2}},
		"dark violet": rule{"dark violet", map[string]int{}}}

	result := countBags("shiny gold", rules, false)

	if result != 126 {
		t.Errorf("Bag count should be 126 but was %d", result)
	}
}
