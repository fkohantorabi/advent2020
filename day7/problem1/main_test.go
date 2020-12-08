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

func TestCanHoldDirectly(t *testing.T) {
	rules := map[string]rule{
		"silver": rule{"silver", map[string]int{"red": 2}}}

	result := canHold("silver", "red", rules)

	if !result {
		t.Errorf("Silver should be able to contain red")
	}
}

func TestCanHoldIndirectly(t *testing.T) {
	rules := map[string]rule{
		"silver": rule{"silver", map[string]int{"green": 1}},
		"green":  rule{"green", map[string]int{"red": 2}}}

	result := canHold("silver", "red", rules)

	if !result {
		t.Errorf("Silver should be able to contain red")
	}
}
