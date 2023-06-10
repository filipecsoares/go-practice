package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	// to take the input from the camel.in file, run the program with 'go run main.go < camel.in'
	fmt.Scanf("%s\n", &input)
	countWords := 1
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			countWords++
		}
	}
	fmt.Println(countWords)
}
