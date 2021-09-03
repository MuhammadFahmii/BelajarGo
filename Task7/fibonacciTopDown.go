package Task7

func Memoize(n int) int {
	temp := map[int]int{}
	hasil := make([]int, n)
	for i := 0; i < n; i++ {
		hasil[i] = FibonacciTopDown(n, temp)
	}
	return hasil[n-1]
}

func FibonacciTopDown(n int, temp map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := temp[n-1]; !ok {
		temp[n-1] = FibonacciTopDown(n-1, temp)
	}
	if _, ok := temp[n-2]; !ok {
		temp[n-2] = FibonacciTopDown(n-2, temp)
	}
	return temp[n-1] + temp[n-2]
}
