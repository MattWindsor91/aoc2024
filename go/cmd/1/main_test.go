package main

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

var testFst = LocationList{3, 4, 2, 1, 3, 3}
var testSnd = LocationList{4, 3, 5, 3, 9, 3}

func TestReadLists(t *testing.T) {
	input := "3 4\n4 3\n2 5\n1 3\n3 9\n3 3"
	r := strings.NewReader(input)

	actualFst, actualSnd, err := ReadLists(r)
	if err != nil {
		t.Fatalf("ReadLists failed: %v", err)
	}

	if !reflect.DeepEqual(actualFst, testFst) {
		t.Errorf("first list: got %v, want %v", actualFst, testFst)
	}
	if !reflect.DeepEqual(actualSnd, testSnd) {
		t.Errorf("second list: got %v, want %v", actualSnd, testSnd)
	}
}

func TestSortedDistances(t *testing.T) {
	expected := []int{2, 1, 0, 1, 2, 5}
	actual := slices.Collect(sortedDistances(testFst, testSnd))

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}

func TestPartOne(t *testing.T) {
	expected := 11
	actual := PartOne(testFst, testSnd)

	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
