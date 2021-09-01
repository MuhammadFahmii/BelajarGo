package Task5

func Fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	} else {
		return Fibonacci(number-1) + Fibonacci(number-2)
	}
}
