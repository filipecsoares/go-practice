package main

import "testing"

func TestMinMax(t *testing.T) {
	expected := [2]int{1, 5}
	got := MinMax([]int{1, 2, 3, 4, 5})
	if expected != got {
		t.Errorf("Expected: %d, Got: %d", expected, got)
	}
	expected = [2]int{5, 2334454}
	got = MinMax([]int{2334454, 5})
	if expected != got {
		t.Errorf("Expected: %d, Got: %d", expected, got)
	}
}
