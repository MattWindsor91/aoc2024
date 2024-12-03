package main

import (
	"reflect"
	"testing"
)

var example = Input{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestInput_NumSafe(t *testing.T) {
	t.Parallel()

	expected := 2
	actual := example.NumSafe()

	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}

func TestInput_NumSafeWithDampening(t *testing.T) {
	t.Parallel()

	expected := 4
	actual := example.NumSafeWithDampening()

	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
func TestReport_IsSafe(t *testing.T) {
	t.Parallel()

	expected := []bool{true, false, false, false, false, true}
	actual := make([]bool, len(expected))

	for i, l := range example {
		actual[i] = l.IsSafe()
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}

func TestReport_IsSafeWithDampening(t *testing.T) {
	t.Parallel()

	expected := []bool{true, false, false, true, true, true}
	actual := make([]bool, len(expected))

	for i, l := range example {
		actual[i] = l.IsSafeWithDampening()
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}

func TestReport_Dampen(t *testing.T) {
	if _, ok := example[1].Dampen(); ok {
		t.Errorf("row 1 should not be safe")
	}
	if _, ok := example[2].Dampen(); ok {
		t.Errorf("row 2 should not be safe")
	}
	if i, ok := example[3].Dampen(); !ok {
		t.Errorf("row 3 should be safe")
	} else if i != 1 {
		t.Errorf("unexpected dampening for row 3: %d", i)
	}
	if i, ok := example[4].Dampen(); !ok {
		t.Errorf("row 4 should be safe")
	} else if i != 2 {
		t.Errorf("unexpected dampening for row 4: %d", i)
	}
}
