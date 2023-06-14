package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		checkAndPrint(i)
		fmt.Print(", ")
	}
	checkAndPrint(n)
	fmt.Println()
}

func checkAndPrint(n int) {
	// if n%3 == 0 && n%5 == 0 {
	// 	fmt.Print("Fizz Buzz")
	// } else if n%3 == 0 {
	// 	fmt.Print("Fizz")
	// } else if n%5 == 0 {
	// 	fmt.Print("Buzz")
	// } else {
	// 	fmt.Print(n)
	// }
	switch {
	// n%15 == 0 also resolve it, because any number divided by 3 and 5, is also divided by 15
	case n%3 == 0 && n%5 == 0:
		fmt.Print("Fizz Buzz")
	case n%3 == 0:
		fmt.Print("Fizz")
	case n%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}
