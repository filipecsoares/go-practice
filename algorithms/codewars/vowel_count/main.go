package main

import "strings"

func GetCount(str string) (count int) {
	vowels := "aeiou"
	for _, letter := range str {
		if strings.Contains(vowels, string(letter)) {
			count++
		}
	}
	return count
}

func GetCount2(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
