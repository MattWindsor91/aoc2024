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

func TestReport_IsSafe(t *testing.T) {
	expected := []bool{true, false, false, false, false, true}
	actual := make([]bool, len(expected))

	for i, l := range example {
		actual[i] = l.IsSafe()
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}

func TestInput_NumSafe(t *testing.T) {
	expected := 2
	actual := example.NumSafe()

	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
