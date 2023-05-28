package fibonacci

func NthFibonacci(num int) int {
    if num == 0 || num == 1 {
        return num
    }
    return NthFibonacci(num - 1) + NthFibonacci(num - 2)
}