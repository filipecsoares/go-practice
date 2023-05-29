package fibonacci

func NthFibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return NthFibonacci(num-1) + NthFibonacci(num-2)
}

func IsFibonacci(num uint) bool {
	if num <= 1 {
		return true
	}
	var first uint = 0
	var second uint = 1
	for i := 0; i < int(num); i++ {
		first, second = second, first+second
		if num == uint(second) {
			return true
		}
	}
	return false
}
