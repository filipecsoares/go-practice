package module01

func GCD(a, b int) int {
	// smallest := a
	// if b < a {
	// 	smallest = b
	// }
	// for i := smallest; i >= 1; i-- {
	// 	if a%i == 0 && b%i == 0 {
	// 		return i
	// 	}
	// }
	// return 0

	// recursive
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return GCD(a, b)
}
