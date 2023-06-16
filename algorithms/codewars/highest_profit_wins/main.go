package main

import "sort"

func MinMax(arr []int) [2]int {
	var result [2]int
	for i, item := range arr {
		if item < result[0] || i == 0 {
			result[0] = item
		}
		if item > result[1] || i == 0 {
			result[1] = item
		}
	}
	return result
}

func MinMaxShortVersion(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}
