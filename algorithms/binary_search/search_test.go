package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 4
	result := BinarySearch(nums, target)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
