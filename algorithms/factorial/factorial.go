package factorial

func Factorial(num int) int {
	if num == 0 {
		return 0
	}
	result := num
	for i := num - 1; i >= 1; i-- {
		result *= i
	}
	return result
}
