package Task7

func FibonacciBottomTop(n int) int {
	arr := []int{0, 1}
	for i := 0; i < n; i++ {
		val := arr[i] + arr[i+1]
		arr = append(arr, val)
	}
	return arr[n]
}
