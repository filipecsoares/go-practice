package main

func Dominator(a []int) int {
	counts := make(map[int]int)
	for _, num := range a {
		counts[num]++
		if counts[num] > len(a)/2 {
			return num
		}
	}
	return -1
}
