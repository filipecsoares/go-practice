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

func FactorialRecursive(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return FactorialRecursive(num-1) * num
}
